package api

import (
	"TeleChan/pkg/log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Carrier represents the item of goroutine channels
type Carrier struct {
	user string       // the user id
	data *MessageData // the message
}

// handle the user's connect message
func (e *apiExecutor) handleUp(conn *Connection, up *Connect) error {
	e.userMux.Lock()
	defer e.userMux.Unlock()

	// update the user's online status
	conn.online = true
	// check if the user connection is existed
	if _, ok := e.users[conn.user]; ok {
		return status.Errorf(codes.PermissionDenied, "user connection is existed")
	}
	// buffer the user's connection
	e.users[conn.user] = conn

	return nil
}

// 1. handle the left data in queue
// 2. push the buffered messages to users
func (e *apiExecutor) afterCare() bool {
	delay := e.settings.Server.PushDelayTime

	select {
	case m := <-e.messages:
		e.send(m)

	case <-time.After(time.Second * delay):
		return false
	}
	return true
}

// send the messsages to user
func (e *apiExecutor) send(carriers []*Carrier) {
	for _, carrier := range carriers {
		// if the user is online
		conn, ok := e.users[carrier.user]
		if !ok || !conn.online {
			log.Infof("user %s is offline", carrier.user)
			continue
		}

		res := &ServerUpdate{
			Message: &ServerUpdate_Data{
				Data: carrier.data,
			},
		}
		// send the message to user
		if err := conn.stream.Send(res); err != nil {
			log.Infof("stream send {%v}: %v", carrier, err)
		}
	}
}

// create goroutines to push messages to users
func (e *apiExecutor) handlePush(wg *sync.WaitGroup) {
	defer wg.Done()

	// mark the goroutine started
	e.ready <- struct{}{}
	for {
		select {
		case m := <-e.messages:
			log.Infof("messages: %v", m)
			e.send(m)
		case <-e.closed:
			// if the close signal is received
			// try to handle the left data, and push the messages to user
			for {
				if e.afterCare() == false {
					log.Info("push goroutine exit")
					return
				}
			}
		}
	}
}

// handle the user's disconnect message
func (e *apiExecutor) handleDown(conn *Connection, down *Disconnect) error {
	e.userMux.Lock()
	defer e.userMux.Unlock()

	// update the user's online status
	conn.online = false
	// check if the user connection is existed
	if _, ok := e.users[conn.user]; ok {
		delete(e.users, conn.user)
	}

	return nil
}

// handle the user's ping message
func (e *apiExecutor) handlePing(conn *Connection, ping *Ping) error {
	// calculate the one trip for the ping
	latency := time.Now().UnixNano()/1000000 - int64(ping.Timestamp)
	log.Infof("Ping from user %s, latency: %dms", conn.user, latency)

	// send the pong message back to user
	res := &ServerUpdate{
		Message: &ServerUpdate_Pong{
			Pong: &Pong{
				Timestamp: ping.Timestamp,
			},
		},
	}
	return conn.stream.Send(res)
}

// Communicate - push the messages to user by the stream
func (e *apiExecutor) Communicate(stream TeleChan_CommunicateServer) error {
	// resolve the user's address
	peer := peerAddr(stream.Context())

	user := e.getUserId(stream.Context())
	// new a connection for the user
	conn := newConnection(user, peer, stream)

	var recvErr error
	// create a channel for the user's requests
	reqChan := make(chan *ClientUpdate, 1)
	// start a goroutine for read the requests
	go e.receiveThread(conn, reqChan, &recvErr)

	for {
		select {
		case <-e.closed:
			log.Infof("stream got close signal")
			return nil

		// read and handle the user's request
		case req, ok := <-reqChan:
			if !ok {
				// handle user's connection down
				e.handleDown(conn, &Disconnect{})

				return recvErr
			}

			// handle the user's connect message
			up := req.GetUp()
			if up != nil {
				if err := e.handleUp(conn, up); err != nil {
					return err
				}
			}

			// handle the user's ping message
			ping := req.GetPing()
			if ping != nil {
				if err := e.handlePing(conn, ping); err != nil {
					return err
				}
			}

			// handle the user's disconnect message
			down := req.GetDown()
			if down != nil {
				if err := e.handleDown(conn, down); err != nil {
					return err
				}
			}
		}
	}
}

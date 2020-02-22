package api

import (
	"TeleChan/pkg/log"
	"io"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// CommunicateStream - the communicate stream interface
type CommunicateStream interface {
	Send(*ServerUpdate) error
	Recv() (*ClientUpdate, error)
	grpc.ServerStream
}

// Connection defines the user's connection
type Connection struct {
	stream CommunicateStream // the user's connection stream
	user   string            // the user's id
	peer   string            // the user's address
	online bool              // if the user is online
	done   chan struct{}     // close and release the connection
}

// new a connection with the peer and stream
func newConnection(user string, peer string, stream CommunicateStream) *Connection {
	return &Connection{
		done:   make(chan struct{}),
		stream: stream,
		user:   user,
		peer:   peer,
	}
}

// start a new goroutine for receive messages from the user's connection
func (e *apiExecutor) receiveThread(conn *Connection, reqChan chan *ClientUpdate, recvErr *error) {
	defer close(reqChan)
	for {
		req, err := conn.stream.Recv()
		if err != nil {
			// if the connection or request is cancelled or terminated by the client
			if status.Code(err) == codes.Canceled || err == io.EOF {
				log.Errorf("%q is terminated: %v", conn.peer, err)
				return
			}

			*recvErr = err
			log.Errorf("%q is terminated with errors %v", conn.peer, err)
			return
		}

		// send the received message to channel
		reqChan <- req
	}
}

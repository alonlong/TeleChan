package api

import (
	"TeleChan/pkg/db/model"
	"TeleChan/pkg/log"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Publish - user publish a new message to channel
func (e *apiExecutor) Publish(ctx context.Context, request *PublishRequest) (*SuccessReply, error) {
	if request.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Channel name is empty")
	}
	if request.Content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Content is empty")
	}

	// get the user id from context
	user := e.getUserId(ctx)

	// query the channel from database by name
	c, err := e.handler.FindChannelByName(request.Name)
	if err != nil {
		log.Errorf("Find channel by name %s: %v", request.Name, err)
		return nil, status.Errorf(codes.Internal, "Find channel: %v", err)
	}
	if c == nil {
		return nil, status.Errorf(codes.NotFound, "Channel is not found")
	}
	if user != c.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "Channel is not owned")
	}

	// insert the message to database
	m := &model.Message{
		Content: request.Content,
	}
	if err := e.handler.CreateMessage(m); err != nil {
		log.Errorf("Create message: %v", err)
		return nil, status.Errorf(codes.Internal, "Create message: %v", err)
	}

	// insert the channel-message to database
	cm := &model.ChannelMessage{
		ChannelID: c.ChannelID,
		MessageID: m.ID,
		Brief:     m.Content,
	}
	if err := e.handler.CreateChannelMessage(cm); err != nil {
		log.Errorf("Create message: %v", err)
		return nil, status.Errorf(codes.Internal, "Create message: %v", err)
	}

	// query the users from database who joined the channel by channel id
	users, err := e.handler.FindJoinedUsersByChannel(c.ChannelID)
	if err != nil {
		log.Errorf("Find users by channel %s: %v", c.ChannelID, err)
		return nil, status.Errorf(codes.Internal, "Find users: %v", err)
	}
	// if there is no other users joined this channel
	if len(users) == 0 {
		return &SuccessReply{}, nil
	}

	carriers := []*Carrier{}
	// prepare the messages for pushing
	for _, u := range users {
		carriers = append(carriers, &Carrier{
			user: u.UserID,
			data: &MessageData{
				Name:    u.Name,
				Content: cm.Brief,
				Seq:     cm.MessageID,
			},
		})
	}

	timeout := e.settings.Server.PublishTimeout
	// send the messages to queue
	select {
	case e.messages <- carriers:
	case <-time.After(time.Second * timeout):
		return nil, status.Errorf(codes.DeadlineExceeded, "Publish timeout")
	}

	return &SuccessReply{}, nil
}

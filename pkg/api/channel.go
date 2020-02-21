package api

import (
	"TeleChan/pkg/db/model"
	"TeleChan/pkg/log"
	context "context"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewChan - create a new channel with name
func (e *apiExecutor) NewChan(ctx context.Context, request *NewChanRequest) (*SuccessReply, error) {
	if request.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Channel name is empty")
	}
	// fetch the user id from context
	userID := e.getUserId(ctx)

	// query the channel from database by name
	u, err := e.handler.FindChannelByName(request.Name)
	if err != nil {
		log.Errorf("Find channel by name %s: %v", request.Name, err)
		return nil, status.Errorf(codes.Internal, "Find channel: %v", err)
	}
	if u != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Channel is existed")
	}

	u = &model.Channel{
		ChannelID: newUUID(),
		UserID:    userID,
		Name:      request.Name,
	}
	// insert the channel to database
	if err := e.handler.CreateChannel(u); err != nil {
		log.Errorf("Create channel: %v", err)
		return nil, status.Errorf(codes.Internal, "Create channel: %v", err)
	}

	uc := &model.UserChannel{
		UserID:    u.UserID,
		ChannelID: u.ChannelID,
		Name:      u.Name,
		Owned:     true,
	}
	// insert the user channel to database
	if err := e.handler.CreateUserChannel(uc); err != nil {
		log.Errorf("Create user-channel: %v", err)
		return nil, status.Errorf(codes.Internal, "Create user-channel: %v", err)
	}

	return &SuccessReply{}, nil
}

// JoinChan - user joins the channel with name
func (e *apiExecutor) JoinChan(ctx context.Context, request *JoinChanRequest) (*SuccessReply, error) {
	if request.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Channel name is empty")
	}
	// fetch the user id from context
	userID := e.getUserId(ctx)

	// query the channel from database by name
	u, err := e.handler.FindChannelByName(request.Name)
	if err != nil {
		log.Errorf("Find channel by name %s: %v", request.Name, err)
		return nil, status.Errorf(codes.Internal, "Find channel: %v", err)
	}
	if u == nil {
		return nil, status.Errorf(codes.NotFound, "Channel is not found")
	}
	if userID == u.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "Not allow to join own channel")
	}

	// query the user channel from database by unique key
	uc, err := e.handler.FindUserChannelByUK(userID, u.ChannelID)
	if err != nil {
		log.Errorf("Find user-channel by user_id %s, channel_id %s: %v", userID, u.ChannelID, err)
		return nil, status.Errorf(codes.Internal, "Find user-channel: %v", err)
	}
	if uc != nil {
		return nil, status.Errorf(codes.AlreadyExists, "User-Channel is existed")
	}

	uc = &model.UserChannel{
		UserID:    userID,
		ChannelID: u.ChannelID,
		Name:      u.Name,
		Owned:     false,
	}
	// insert the user channel to database
	if err := e.handler.CreateUserChannel(uc); err != nil {
		log.Errorf("Create user-channel: %v", err)
		return nil, status.Errorf(codes.Internal, "Create user-channel: %v", err)
	}

	return &SuccessReply{}, nil
}

// GetChans - query the user's owned and joined channels
func (e *apiExecutor) GetChans(ctx context.Context, request *GetChansRequest) (*ChansReply, error) {
	// fetch the user id from context
	userID := e.getUserId(ctx)

	// get the user's owned and joined channels
	ucs, err := e.handler.FindUserChannelsByUser(userID)
	if err != nil {
		log.Errorf("Find user-channel by user_id %s: %v", userID, err)
		return nil, status.Errorf(codes.Internal, "Find user-channel: %v", err)
	}

	// prepare the channels
	chans := []*Channel{}
	for _, uc := range ucs {
		chans = append(chans, &Channel{
			Id:    uc.ChannelID,
			Name:  uc.Name,
			Owned: uc.Owned,
		})
	}
	return &ChansReply{Channels: chans}, nil
}

// SearchChans - search the channels by input
func (e *apiExecutor) SearchChans(ctx context.Context, request *SearchChansRequest) (*ChansReply, error) {
	if request.Input == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Search input is empty")
	}

	// fetch the user id from context
	userID := e.getUserId(ctx)

	// get the channels by like input
	us, err := e.handler.FindChannelsByLikeName(request.Input)
	if err != nil {
		log.Errorf("Find channel by like %s: %v", request.Input, err)
		return nil, status.Errorf(codes.Internal, "Find channel: %v", err)
	}

	// prepare the channels
	chans := []*Channel{}
	for _, u := range us {
		chans = append(chans, &Channel{
			Id:    u.ChannelID,
			Name:  u.Name,
			Owned: (u.UserID == userID),
		})
	}
	return &ChansReply{Channels: chans}, nil
}

// LeaveChan - leave the joined channel
func (e *apiExecutor) LeaveChan(ctx context.Context, request *LeaveChanRequest) (*SuccessReply, error) {
	if request.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Channel name is empty")
	}

	userID := e.getUserId(ctx)
	// query the channel from database by name
	u, err := e.handler.FindChannelByName(request.Name)
	if err != nil {
		log.Errorf("Find channel by name %s: %v", request.Name, err)
		return nil, status.Errorf(codes.Internal, "Find channel: %v", err)
	}
	if u == nil {
		return nil, status.Errorf(codes.NotFound, "Channel is not found")
	}

	// query the user channel by user id and channel id
	uc, err := e.handler.FindUserChannelByUK(userID, u.ChannelID)
	if err != nil {
		log.Errorf("Find user-channel by user_id %s, channel_id %s: %v", userID, u.ChannelID, err)
		return nil, status.Errorf(codes.Internal, "Find user-channel: %v", err)
	}
	if uc == nil {
		return nil, status.Errorf(codes.NotFound, "User-channel is not found")
	}
	if uc.Owned {
		return nil, status.Errorf(codes.PermissionDenied, "Delete owned user-channel")
	}

	// delete the user channel by user id and channel id
	if err := e.handler.DeleteUserChannelByUK(userID, u.ChannelID); err != nil {
		log.Errorf("Delete user-channel by user_id %s, channel_id %s: %v", userID, u.ChannelID, err)
		return nil, status.Errorf(codes.Internal, "Delete user-channel: %v", err)
	}

	return &SuccessReply{}, nil
}

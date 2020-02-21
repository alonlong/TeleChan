package api

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/metadata"
)

var (
	headerAuthorize = "authorization"
	headerUser      = "user-id"
)

// get the jwt token from context
func (s *apiExecutor) getJwtToken(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	res := md.Get(headerAuthorize)
	if len(res) > 0 {
		return res[0]
	}
	return ""
}

// get the user id 'email' from context
func (s *apiExecutor) getUserId(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	res := md.Get(headerUser)
	if len(res) == 0 {
		panic("user id is not found")
	}
	return res[0]
}

// create a new uuid
func newUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

package client

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
)

const AuthContextKey = "authorization"

type ContextBasedJWT struct{}

// GetRequestMetadata retrieves authentication information from the Context of each request. This
// implements the PerRPCCredentials interface, which tells the gRPC framework to supply the
// "authorization" metadata on each request and set it to the value of the token stored in the Context.
func (ContextBasedJWT) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	token, ok := ctx.Value(AuthContextKey).(oauth2.Token)
	if !ok {
		return nil, errors.New("context value for authorization was not oauth2.Token")
	}

	return map[string]string{
		AuthContextKey: token.Type() + " " + token.AccessToken,
	}, nil
}

func (ContextBasedJWT) RequireTransportSecurity() bool {
	return false
}

// Auth is a helper that adds a token with the provided username into a request Context.
func Auth(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, AuthContextKey, oauth2.Token{TokenType: "Bearer", AccessToken: user})
}

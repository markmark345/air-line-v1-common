package contexts

import (
	"context"
	"errors"
)

type ContextKey string

const (
	ContextKeyUserId ContextKey = "userId"
)

func SetUseeId(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserId, userID)
}

func GetUserId(ctx context.Context) (string, error) {
	if str := ctx.Value(ContextKeyUserId); str != nil && str != "" {
		return str.(string), nil
	}
	return "", errors.New("user id not found")
}

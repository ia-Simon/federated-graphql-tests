// Package userid implements functions to deal with the
// system header X-User-ID.
package userid

import (
	"context"
	"net/http"
)

type key struct{}

var userIDKey key

const header = "X-User-ID"

func ToHeader(h *http.Header, userID string) {
	h.Set(header, userID)
}

func FromHeader(h *http.Header) string {
	return h.Get(header)
}

func ToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func FromContext(ctx context.Context) string {
	return ctx.Value(userIDKey).(string)
}

package middleware

import (
	"net/http"

	"asap.local/sys-headers/userid"
)

func HTTPSystemHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := userid.ToContext(r.Context(), userid.FromHeader(&r.Header))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

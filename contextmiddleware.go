package terse

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
)

func WithContext(handler http.Handler, baseContext context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.WithContext(baseContext, r)
		request := r.WithContext(ctx)
		handler.ServeHTTP(w, request)
	})
}

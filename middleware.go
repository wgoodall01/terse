package terse

import (
	"context"
	"net/http"
	"os"
	"strings"
	"sync"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func WithAppengineContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		request := r.WithContext(ctx)
		handler.ServeHTTP(w, request)
	})
}

type contextKey string

func (ck *contextKey) String() string {
	return "middleware context " + string(*ck)
}

const claimsContextKey = contextKey("claims")

func WithAuthentication(handler http.Handler) http.Handler {
	var setupOnce sync.Once
	var verifyJwt JwtVerifier

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		setupOnce.Do(func() {
			clientId := os.Getenv("OAUTH_CLIENT_ID")
			var err error
			verifyJwt, err = NewGoogleJwtVerifier(ctx, clientId)
			if err != nil {
				// log and blow up
				log.Criticalf(ctx, "Can't load OAuth verifier keys, exiting")
				os.Exit(10)
			}
		})

		// get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// If it's not there, fail the request
			http.Error(w, `No "Authorization" header present.`, 401)
			return
		}

		// Split header into "Bearer" "<JWT>"
		splitHeader := strings.Split(authHeader, " ")
		if len(splitHeader) != 2 {
			// If the header has more parts, fail the request
			http.Error(w, `Malformed "Authorization" header.`, 401)
		}
		if splitHeader[0] != "Bearer" {
			// If first part isn't "Bearer", fail the request
			http.Error(w, `No "Authorization: Bearer xxxxxxxxxx" header.`, 401)
			return
		}

		// get the JWT and verify it
		rawJwt := splitHeader[1]
		claims, jwtErr := verifyJwt(ctx, rawJwt)
		if jwtErr != nil {
			// if the JWT isn't valid, fail the request.
			http.Error(w, "Invalid auth token", 401)
			log.Warningf(ctx, "Invalid auth token: %v", jwtErr)
			return
		}

		// put the claims in the request context
		newReq := r.WithContext(context.WithValue(ctx, claimsContextKey, claims))

		// run the handler
		handler.ServeHTTP(w, newReq)
	})
}

// GetClaims returns the Claims from the context of a JWT-authenticated endpoint
func GetClaims(ctx context.Context) *Claims {
	return ctx.Value(claimsContextKey).(*Claims)
}

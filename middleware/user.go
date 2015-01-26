package middleware

import (
	"log"
	"net/http"

	"github.com/dbongo/hackapp/session"
	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

// SetUser ...
func SetUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.FromC(*c)
		log.Printf("ctx: %v", ctx)
		log.Println()
		var user = session.GetUser(ctx, r)
		log.Printf("user: %v", user)
		log.Println()
		if user != nil && user.ID != "" {
			UserToC(c, user)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// RequireUser ...
func RequireUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if ToUser(c) == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

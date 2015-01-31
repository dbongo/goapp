package middleware

import (
	"net/http"

	"github.com/dbongo/hackapp/session"
	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

// SetUser ...
func SetUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.FromC(*c)
		user := session.GetUser(ctx, r)
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

// RequireAdminUser ...
func RequireAdminUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := ToUser(c)
		switch {
		case user == nil:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case user != nil && !user.Admin:
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

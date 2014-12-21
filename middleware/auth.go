package middleware

import (
	"fmt"
	"net/http"

	"github.com/dbongo/goapp/auth"
	"github.com/dbongo/goapp/logger"

	"github.com/zenazn/goji/web"
)

// Auth ...
func Auth(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token := r.Header.Get("Authorization")
		fmt.Sscanf(token, "Bearer %s", &token)
		user := auth.UserFromJWT(string(token))
		logger.Info.Printf("current user is %v", user)
		if user == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		UserToC(c, user)
		logger.Info.Printf("current user from c.Env[user] is %v", c.Env["user"])
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// SetUser is a middleware function that retrieves the currently authenticated
// user from the request and stores in the context.
// func SetUser(c *web.C, h http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		ctx := context.FromC(*c)
// 		user := auth.GetUser(ctx, r)
// 		logger.Info.Printf("current context.Context is %v", ctx)
// 		logger.Info.Printf("current user is %v", user)
// 		if user != nil {
// 			UserToC(c, user)
// 		}
// 		logger.Info.Printf("current goji.Context is %v", c)
// 		h.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(fn)
// }

// RequireUser is a middleware function that verifies
// there is a currently authenticated user stored in
// the context.
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

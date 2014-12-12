package api

import (
	"errors"

	"github.com/zenazn/goji/web"
)

const (
	errRequestKey string = "RequestErr"
	currentUser   string = "CurrentUser"
)

// ErrUserNotSigned ...
var ErrUserNotSigned = errors.New("User is not signed in.")

// AddRequestErr ...
func AddRequestErr(c *web.C, error *HTTPResponse) {
	c.Env[errRequestKey] = error
}

// GetRequestErr ...
func GetRequestErr(c *web.C) (*HTTPResponse, bool) {
	val, ok := c.Env[errRequestKey].(*HTTPResponse)
	if !ok {
		return nil, false
	}
	return val, true
}

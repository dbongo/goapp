package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	"github.com/zenazn/goji/web"
)

// Key to use when setting the request ID
const RequestIDKey = "reqID"

var prefix string
var reqid uint64

func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

// RequestID ...
func RequestID(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if c.Env == nil {
			c.Env = make(map[string]interface{})
		}
		myid := atomic.AddUint64(&reqid, 1)
		c.Env[RequestIDKey] = fmt.Sprintf("%s-%06d", prefix, myid)

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// GetReqID ...
func GetReqID(c web.C) string {
	if c.Env == nil {
		return ""
	}
	v, ok := c.Env[RequestIDKey]
	if !ok {
		return ""
	}
	if reqID, ok := v.(string); ok {
		return reqID
	}
	return ""
}

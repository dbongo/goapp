package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Config is a struct for specifying configuration parameters for the
// Logger middleware.
type Config struct {
	Prefix               string
	DisableAutoBrackets  bool
	RemoteAddressHeaders []string
}

// Logger ...
type Logger struct {
	http.Handler

	ch   chan *record
	conf Config
}

// HTTPLogger ...
func HTTPLogger(h http.Handler) http.Handler {
	l := newHTTPLogger(h)
	fn := func(rw http.ResponseWriter, req *http.Request) {
		l.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	addr := req.RemoteAddr
	for _, headerKey := range l.conf.RemoteAddressHeaders {
		if val := req.Header.Get(headerKey); len(val) > 0 {
			addr = val
			break
		}
	}
	r := &record{
		ResponseWriter: rw,
		start:          time.Now().UTC(),
		ip:             addr,
		method:         req.Method,
		rawpath:        req.RequestURI,
		responseStatus: http.StatusOK,
		proto:          req.Proto,
		userAgent:      req.UserAgent(),
	}
	l.Handler.ServeHTTP(r, req)
	l.ch <- r
}

func (r *record) Write(b []byte) (int, error) {
	written, err := r.ResponseWriter.Write(b)
	r.responseBytes += int64(written)
	return written, err
}

func (r *record) WriteHeader(status int) {
	r.responseStatus = status
	r.ResponseWriter.WriteHeader(status)
}

type record struct {
	http.ResponseWriter

	start               time.Time
	ip, method, rawpath string
	responseStatus      int
	responseBytes       int64
	proto, userAgent    string
}

func newHTTPLogger(h http.Handler) http.Handler {
	c := Config{"hackapp", false, []string{"X-Forwarded-Proto"}}
	l := &Logger{
		Handler: h,
		ch:      make(chan *record, 1000),
		conf:    c,
	}
	go l.logResponse()
	return l
}

// [hackapp] 2014/12/30 20:41:41 [::1]:62629 GET /api/hello 200 13b 437.126µs HTTP/1.1 curl/7.37.1
// [hackapp] 2014/12/30 20:47:04 [::1]:62930 POST /login 200 490b 224.032597ms HTTP/1.1 curl/7.37.1
func (l *Logger) logResponse() {
	for {
		lr := <-l.ch
		timeStamp := fmt.Sprintf(
			"%04d/%02d/%02d %02d:%02d:%02d",
			lr.start.Year(),
			lr.start.Month(),
			lr.start.Day(),
			lr.start.Hour(),
			lr.start.Minute(),
			lr.start.Second(),
		)
		prefix := l.conf.Prefix
		if len(prefix) > 0 && l.conf.DisableAutoBrackets == false {
			prefix = "[" + prefix + "]"
		}
		logRecord := fmt.Sprintf(
			"%s %s %s %s %s %d %db %v %s %s\n",
			prefix,
			timeStamp,
			lr.ip,
			lr.method,
			lr.rawpath,
			lr.responseStatus,
			lr.responseBytes,
			time.Since(lr.start),
			lr.proto,
			lr.userAgent,
		)
		os.Stdout.WriteString(logRecord)
	}
}

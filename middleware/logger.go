package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type logRecord struct {
	http.ResponseWriter

	time                      time.Time
	ip, method, rawpath       string
	responseStatus            int
	responseBytes             int64
	proto, userAgent, referer string
}

type logHandler struct {
	ch      chan *logRecord
	handler http.Handler
}

// log format
// 2014/12/27 08:43:01 127.0.0.1 POST /login 200 HTTP/1.1 289 Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)
func (lh *logHandler) logFromChannel() {
	for {
		lr := <-lh.ch

		dateString := fmt.Sprintf(
			"%04d/%02d/%02d %02d:%02d:%02d",
			lr.time.Year(),
			lr.time.Month(),
			lr.time.Day(),
			lr.time.Hour(),
			lr.time.Minute(),
			lr.time.Second())

		logLine := fmt.Sprintf(
			"%s %s %s %s %d %s %d %s %s\n",
			dateString,
			lr.ip,
			lr.method,
			lr.rawpath,
			lr.responseStatus,
			lr.proto,
			lr.responseBytes,
			lr.userAgent,
			lr.referer)

		os.Stdout.WriteString(logLine)
	}
}

// HTTPLogger ...
func HTTPLogger(h http.Handler) http.Handler {
	hl := NewHTTPLogger(h)
	fn := func(rw http.ResponseWriter, req *http.Request) {
		hl.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}

// NewHTTPLogger ...
func NewHTTPLogger(handler http.Handler) http.Handler {
	lh := &logHandler{
		ch:      make(chan *logRecord, 1000),
		handler: handler,
	}
	go lh.logFromChannel()
	return lh
}

func (lh *logHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	addr := req.RemoteAddr
	if colon := strings.LastIndex(addr, ":"); colon != -1 {
		addr = addr[:colon]
	}
	lr := &logRecord{
		ResponseWriter: rw,
		time:           time.Now(),
		ip:             addr,
		method:         req.Method,
		rawpath:        req.RequestURI,
		responseStatus: http.StatusOK,
		proto:          req.Proto,
		userAgent:      req.UserAgent(),
		referer:        req.Referer(),
	}
	lh.handler.ServeHTTP(lr, req)
	lh.ch <- lr
}

func (lr *logRecord) Write(b []byte) (int, error) {
	written, err := lr.ResponseWriter.Write(b)
	lr.responseBytes += int64(written)
	return written, err
}

func (lr *logRecord) WriteHeader(status int) {
	lr.responseStatus = status
	lr.ResponseWriter.WriteHeader(status)
}

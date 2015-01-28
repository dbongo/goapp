package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Logger ...
type Logger struct {
	http.Handler
	ch chan *Record
}

// Record ...
type Record struct {
	http.ResponseWriter
	start     time.Time
	ip        string
	method    string
	rawpath   string
	status    int
	bytes     int64
	proto     string
	userAgent string
}

// HTTPLogger ...
func HTTPLogger(h http.Handler) http.Handler {
	l := logHTTP(h)
	fn := func(rw http.ResponseWriter, req *http.Request) {
		l.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}

func (log *Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	addr := req.RemoteAddr
	record := &Record{
		ResponseWriter: w,
		start:          time.Now().Local(),
		ip:             addr,
		method:         req.Method,
		rawpath:        req.RequestURI,
		status:         http.StatusOK,
		proto:          req.Proto,
		userAgent:      req.UserAgent(),
	}
	log.Handler.ServeHTTP(record, req)
	log.ch <- record
}

func (r *Record) Write(b []byte) (int, error) {
	written, err := r.ResponseWriter.Write(b)
	r.bytes += int64(written)
	return written, err
}

// WriteHeader ...
func (r *Record) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func logHTTP(h http.Handler) http.Handler {
	log := &Logger{
		Handler: h,
		ch:      make(chan *Record, 1000),
	}
	go log.response()
	return log
}

func (log *Logger) response() {
	for {
		var buf bytes.Buffer
		res := <-log.ch
		timeStamp := fmt.Sprintf(
			"%04d/%02d/%02d %02d:%02d:%02d",
			res.start.Year(),
			res.start.Month(),
			res.start.Day(),
			res.start.Hour(),
			res.start.Minute(),
			res.start.Second(),
		)
		fmt.Fprintf(&buf, "%s %s - %s %s ", timeStamp, res.ip, res.method, res.rawpath)
		if res.status < 200 {
			writeColor(&buf, bBlue, "%03d ", res.status)
		} else if res.status < 300 {
			writeColor(&buf, bGreen, "%03d ", res.status)
		} else if res.status < 400 {
			writeColor(&buf, bCyan, "%03d ", res.status)
		} else if res.status < 500 {
			writeColor(&buf, bYellow, "%03d ", res.status)
		} else {
			writeColor(&buf, bRed, "%03d ", res.status)
		}
		fmt.Fprintf(&buf, "%v - %d %s %s\n", time.Since(res.start), res.bytes, res.proto, res.userAgent)
		os.Stdout.WriteString(buf.String())
	}
}

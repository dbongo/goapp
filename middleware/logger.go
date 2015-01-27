package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Config ...
type Config struct {
	Prefix               string
	DisableAutoBrackets  bool
	RemoteAddressHeaders []string
}

// Logger ...
type Logger struct {
	http.Handler

	ch   chan *Record
	conf Config
}

// Record ...
type Record struct {
	http.ResponseWriter

	start               time.Time
	ip, method, rawpath string
	responseStatus      int
	responseBytes       int64
	proto, userAgent    string
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
	for _, headerKey := range log.conf.RemoteAddressHeaders {
		if val := req.Header.Get(headerKey); len(val) > 0 {
			addr = val
			break
		}
	}
	record := &Record{
		ResponseWriter: w,

		start:          time.Now().UTC(),
		ip:             addr,
		method:         req.Method,
		rawpath:        req.RequestURI,
		responseStatus: http.StatusOK,
		proto:          req.Proto,
		userAgent:      req.UserAgent(),
	}
	log.Handler.ServeHTTP(record, req)
	log.ch <- record
}

func (r *Record) Write(b []byte) (int, error) {
	written, err := r.ResponseWriter.Write(b)
	r.responseBytes += int64(written)
	return written, err
}

// WriteHeader ...
func (r *Record) WriteHeader(status int) {
	r.responseStatus = status
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
		writeColor(&buf, bWhite, "%s ", timeStamp)
		writeColor(&buf, bWhite, "%s - ", res.ip)
		writeColor(&buf, bWhite, "%s ", res.method)
		writeColor(&buf, bWhite, "%s ", res.rawpath)
		status := res.responseStatus
		if status < 200 {
			writeColor(&buf, bBlue, "%03d ", status)
		} else if status < 300 {
			writeColor(&buf, bGreen, "%03d ", status)
		} else if status < 400 {
			writeColor(&buf, bCyan, "%03d ", status)
		} else if status < 500 {
			writeColor(&buf, bYellow, "%03d ", status)
		} else {
			writeColor(&buf, bRed, "%03d ", status)
		}
		writeColor(&buf, bWhite, "%v - ", time.Since(res.start))
		writeColor(&buf, bWhite, "%d ", res.responseBytes)
		writeColor(&buf, bWhite, "%s ", res.proto)
		writeColor(&buf, bWhite, "%s\n", res.userAgent)
		os.Stdout.WriteString(buf.String())
	}
}

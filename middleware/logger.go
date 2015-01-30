package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

var urlmap = make(map[string]map[string]*Response)

// Logger ...
type Logger struct {
	http.Handler
	ch chan *Response
}

// Response ...
type Response struct {
	http.ResponseWriter
	start     time.Time
	ip        string
	method    string
	rawpath   string
	counter   int64
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

func (l *Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	addr := req.RemoteAddr
	numreqs := count(req.RequestURI, req.Method)
	res := &Response{
		ResponseWriter: w,
		start:          time.Now().Local(),
		ip:             addr,
		method:         req.Method,
		rawpath:        req.RequestURI,
		status:         http.StatusOK,
		counter:        numreqs,
		proto:          req.Proto,
		userAgent:      req.UserAgent(),
	}
	l.Handler.ServeHTTP(res, req)
	l.ch <- res
}

func (r *Response) Write(b []byte) (int, error) {
	written, err := r.ResponseWriter.Write(b)
	r.bytes += int64(written)
	return written, err
}

// WriteHeader ...
func (r *Response) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func logHTTP(h http.Handler) http.Handler {
	log := &Logger{
		Handler: h,
		ch:      make(chan *Response, 1000),
	}
	go log.response()
	return log
}

func (l *Logger) response() {
	for {
		var buf bytes.Buffer
		res := <-l.ch
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
		fmt.Fprintf(&buf, "%v - %dx %d %s %s\n", time.Since(res.start), res.counter, res.bytes, res.proto, res.userAgent)
		os.Stdout.WriteString(buf.String())
	}
}

func count(requrl, reqmethod string) int64 {
	var reqnum int64
	if method, ok := urlmap[requrl]; ok {
		if r, ok := method[reqmethod]; ok {
			r.counter++
			reqnum = r.counter
		} else {
			first := &Response{
				rawpath: requrl,
				method:  reqmethod,
				counter: 1,
			}
			reqnum = first.counter
			urlmap[requrl][reqmethod] = first
		}
	} else {
		methodmap := make(map[string]*Response)
		first := &Response{
			rawpath: requrl,
			method:  reqmethod,
			counter: 1,
		}
		reqnum = first.counter
		methodmap[reqmethod] = first
		urlmap[requrl] = methodmap
	}
	return reqnum
}

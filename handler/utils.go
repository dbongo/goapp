package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dbongo/goapp/logger"
)

type Response map[string]interface{}

func (r *Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func ServeJSON(w http.ResponseWriter, json *Response, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, json)
}

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func writeBody(w http.ResponseWriter, body []byte, status int, contentType string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", contentType+";charset=utf8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.Write(body)
}

// write a plain text message
func writeString(w http.ResponseWriter, body string, status int) {
	writeBody(w, []byte(body), status, "text/plain")
}

func writeJSON(w http.ResponseWriter, value interface{}, status int) {
	body, err := json.Marshal(value)
	if err != nil {
		logger.Error.Print(err)
	}
	writeBody(w, body, status, "application/json")
}

func writeError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func parseBody(body io.ReadCloser, r interface{}) error {
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &r); err != nil {
		return errors.New("The request was bad-formed.")
	}
	return nil
}

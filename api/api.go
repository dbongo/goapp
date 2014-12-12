package api

import (
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/tsuru/config"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

// API ...
type API struct{}

// New ...
func New() *API {
	return &API{}
}

// Init ...
func (api *API) Init() {
	err := config.ReadConfigFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err.Error())
	}
}

// Router ...
func (api *API) Router() {

	// middleware
	goji.Use(RequestIDHandler)
	goji.NotFound(NotFoundHandler)

	// handlers
	servicesHandler := &ServicesHandler{}
	//debugHandler := &DebugHandler{}

	// public routes
	goji.Get("/", api.Route(servicesHandler, "Index"))
}

// Route ...
func (api *API) Route(handler interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		methodValue := reflect.ValueOf(handler).MethodByName(route)
		methodInterface := methodValue.Interface()
		method := methodInterface.(func(c *web.C, w http.ResponseWriter, r *http.Request) *HTTPResponse)
		response := method(&c, w, r)

		// Need to check if there's any error.
		_, err := GetRequestErr(&c)
		if !err {
			w.WriteHeader(response.StatusCode)
			if _, exists := c.Env["Content-Type"]; exists {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			}
			io.WriteString(w, response.Payload)
		}
	}
	return fn
}

package main

import (
	"github.com/dbongo/goapp/api"
	"github.com/zenazn/goji"
)

func main() {
	api := api.New()

	api.Init()
	api.Router()

	goji.Serve()
}

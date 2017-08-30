package main

import (
	"net/http"

	"github.com/joshadambell/commercialdata/routes"

	"goji.io"

	"github.com/joshadambell/commercialdata/internal/mongo"
)

func main() {
	mongo.ConnectOrDie("localhost:27017")

	mux := goji.NewMux()
	routes.Register(mux)

	http.ListenAndServe("localhost:8000", mux)
}

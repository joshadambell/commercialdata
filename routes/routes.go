package routes

import (
	"github.com/joshadambell/commercialdata/routes/handlers"
	goji "goji.io"
	"goji.io/pat"
)

func Register(mux *goji.Mux) {

	mux.HandleFunc(pat.Get("/hello/:name"), handlers.Hello)
	mux.HandleFunc(pat.Post("/hello/:name"), handlers.Hello)
}

package main

import (
	"github.com/go-chi/chi"
)

func createRoutes() chi.Router {
	// We're using chi as the router. You'll want to read
	// the documentation https://github.com/go-chi/chi
	// so that you can capture parameters like /events/5
	// or /api/events/4 -- where you want to get the
	// event id (5 and 4, respectively).

	r := chi.NewRouter()
	r.Get("/", indexController)
	r.Get("/about", aboutController)
	addStaticFileServer(r, "/static/", "staticfiles")
	return r
}

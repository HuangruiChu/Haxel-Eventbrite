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
	r.Get("/events/new", CreateEventController)
	r.Post("/events/new", CreateEventController)
	r.Get("/events/{id}", eventController)
	r.Get("/api/events", EventsAPIController)
	// Create a route for the API endpoint with a dynamic 'id' parameter
	r.Get("/api/events/{id}", EventAPIController)
	addStaticFileServer(r, "/static/", "staticfiles")
	return r
}

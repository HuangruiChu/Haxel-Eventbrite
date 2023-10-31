package main

import (
	"net/http"
)

func aboutController(w http.ResponseWriter, r *http.Request) {

	// Serve the "about" template with custom content
	data := struct {
		Title   string
		Content string
	}{
		"Ahh, the about page!",
		"Custom content for the about page.",
	}
	tmpl["about"].ExecuteTemplate(w, "layout", data)
}

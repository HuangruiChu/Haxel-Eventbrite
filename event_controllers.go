package main

import (
	"net/http"
	"time"
	"github.com/go-chi/chi"
	"strconv"
)


func eventController(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id") // Retrieve the "id" parameter from the URL as a string

    // Convert "idStr" to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Now, you have "id" as an integer

	type eventContextData struct {
		ID        int
		Title     string
		Date      time.Time
		Image     string
		Location  string
		Attending []string
	}

	theEvent, bool := getEventByID(id)
	if bool != true {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	contextData := eventContextData{
		ID:        theEvent.ID,
		Title:     theEvent.Title,
		Date:      theEvent.Date,
		Image:     theEvent.Image,
		Location:  theEvent.Location,
		Attending: theEvent.Attending,
	}

	tmpl["event"].Execute(w, contextData)
}

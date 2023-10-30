package main

import (
	"net/http"
	"time"
)

func indexController(w http.ResponseWriter, r *http.Request) {

	type indexContextData struct {
		Events []Event
		Today  time.Time
	}

	theEvents, err := getAllEvents()
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	contextData := indexContextData{
		Events: theEvents,
		Today:  time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}

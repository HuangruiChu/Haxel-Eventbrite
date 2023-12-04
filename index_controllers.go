package main

import (
	"net/http"
	"time"
	"sort"
)

func sortEventsByDate(events []Event) {
    sort.Slice(events, func(i, j int) bool {
        return events[i].Date.Before(events[j].Date)
    })
}

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
	
	// Sort events by date
	sortEventsByDate(theEvents)

	contextData := indexContextData{
		Events: theEvents,
		Today:  time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}

package main

import (
	"net/http"
	"time"
	"sort"
)

func sortEvents(events []Event) {
    now := time.Now()
    sort.Slice(events, func(i, j int) bool {
        event1IsPast := events[i].Date.Before(now)
        event2IsPast := events[j].Date.Before(now)

        if event1IsPast != event2IsPast {
            return event2IsPast // Puts future events first
        }
        // If both are either in the past or future, sort by date
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

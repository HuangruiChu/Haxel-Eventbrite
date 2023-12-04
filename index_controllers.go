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

        if event1IsPast && event2IsPast {
            // Both events are in the past, sort from newest to oldest
            return events[i].Date.After(events[j].Date)
        } else if !event1IsPast && !event2IsPast {
            // Both events are in the future, sort from oldest to newest
            return events[i].Date.Before(events[j].Date)
        } else {
            // One event is in the past and the other in the future,
            // the future event should come first
            return event2IsPast
        }
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

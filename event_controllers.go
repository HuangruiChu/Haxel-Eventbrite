package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"strconv"
	"time"
	"encoding/json"
)


func eventController(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id") // Retrieve the "id" parameter from the URL as a string

    // Convert "idStr" to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil{
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Now, you have "id" as an integer


	theEvent, bool := getEventByID(id)
	if bool != true {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	contextData := Event{
		ID:        theEvent.ID,
		Title:     theEvent.Title,
		Date:      theEvent.Date,
		Image:     theEvent.Image,
		Location:  theEvent.Location,
		Attending: theEvent.Attending,
	}

	tmpl["event"].Execute(w, contextData)
}

func NewEventController(w http.ResponseWriter, r *http.Request) {
	tmpl["create_event"].ExecuteTemplate(w, "layout", nil)
}

func CreateEventController(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // It's a POST request, process the form data
        dateStr := r.FormValue("date")
        if dateStr == "" {
            http.Error(w, "Date cannot be empty", http.StatusBadRequest)
            return
        }
        // Parse the date
		timelayout := "2006-01-02T15:04"
		date, err := time.Parse(timelayout, dateStr)
        if err != nil {
            http.Error(w, "Invalid date format", http.StatusBadRequest)
			Log.Println("looks like a bad date??",date)
            return
        }
		// Check if the date is in the past
		if date.Before(time.Now()) {
			http.Error(w, "Date cannot be in the past", http.StatusBadRequest)
			return
		}
		// Get the other form values
		Title := r.FormValue("title")
		// Check if the title is valid [TODO]
		if Title == "" {
			http.Error(w, "Title cannot be empty", http.StatusBadRequest)
			return
		}
		Image := r.FormValue("image")
		// Check if the image is valid [TODO]
		if Image == "" {
			http.Error(w, "Image cannot be empty", http.StatusBadRequest)
			return	
		}
		Location := r.FormValue("location")	
		// Check if the location is valid [TODO]
		if Location == "" {
			http.Error(w, "Location cannot be empty", http.StatusBadRequest)
			return
		}

		// Create a new Event
		newEvent := Event{
			Title:    Title, // Assign the title from the form
			Date:     date, // Assign the parsed time.Time
			Image:    Image, // Assign the image from the form
			Location: Location, // Assign the location from the form
		}
	
        // Call the addEvent function to add the new event
        id := addEvent(newEvent)

        // Redirect to the event page
		http.Redirect(w, r, "/events/"+strconv.Itoa(id), http.StatusFound)
		
    } else {
        // Render the form for non-POST requests
        tmpl["create_event"].ExecuteTemplate(w, "layout", nil)
    }
}

// EventResponse is a struct to define the JSON response format.
type EventResponse struct {
	Events []Event `json:"events"`
}

// EventsAPIController is the controller for the API returning all events.
func EventsAPIController(w http.ResponseWriter, r *http.Request) {
	// Get all events
	events, err := getAllEvents()
	if err != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	// Create a response containing the events
	response := EventResponse{Events: events}

	// Marshal the response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response writer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

// EventAPIController is the controller for the API returning a single event.
func EventAPIController(w http.ResponseWriter, r *http.Request) {
	// Extract the 'id' from the URL path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Get the event by ID
	event, found := getEventByID(id)
	if !found {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Marshal the event into JSON
	jsonResponse, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response writer
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
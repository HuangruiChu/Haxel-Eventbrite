package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"strconv"
	"time"
	"encoding/json"
	"net/url"
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

func isValidImageURL(url string) bool {
	//Check if the image is valid [TODO]
	return true
func isValidImageURL(url string) bool {
	//Check if the image is valid [TODO]
	return true
}

// CreateEventController is the controller for the events/new page.
func CreateEventController(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodPost {
		Errors := "" // Create a string to store errors
        // It's a POST request, process the form data
        dateStr := r.FormValue("date")
        if dateStr == "" {
			Errors += "Date cannot be empty! "
        }else{
			// Parse the date
			timelayout := "2006-01-02T15:04"
			date, err := time.Parse(timelayout, dateStr)
			if err != nil {
				Errors += "Invalid date format! " + dateStr
			}
			// Check if the date is in the past
			if date.Before(time.Now()) {
				Errors += "Date cannot be in the past! "
			}
		}
		// Get the Title
		title := r.FormValue("title")
		// Check if the title is valid
		if len(title) <= 5 || len(title) >= 50 {
			Errors += "Bad Title! Title must be between 5 and 50 unicode characters! "
		}

		//Get the image URL
		image := r.FormValue("image")
		// Check if the image is valid
		_, iamgeerr := url.ParseRequestURI(image)
		if iamgeerr != nil {
			Errors += "Invalid image URL! "
		}
		if !isValidImageURL(image){
			Errors += "Image should have file types of \".png\", \".jpg\", \".jpeg\", \".gif\", or \".gifv\". "
		}
		//Get the location
		location := r.FormValue("location")	
		// Check if the location is valid [TODO]
		if len(location) <=5 || len(location) >= 50{
			Errors += "Bad Location! Location must be between 5 and 50 unicode characters! "
		}
		if Errors == "" {
			// Create a new Event
			newEvent := Event{
				Title:    title, // Assign the title from the form
				Date:     date, // Assign the parsed time.Time
				Image:    image, // Assign the image from the form
				Location: location, // Assign the location from the form
			}
			// Call the addEvent function to add the new event
			id := addEvent(newEvent)
			// Redirect to the event page
			http.Redirect(w, r, "/events/"+strconv.Itoa(id), http.StatusFound)
		} else {
			// If there are errors, render the form again with error messages
			tmplData := struct {
                Errors string
            }{
                Errors: Errors,
            }
            tmpl["create_event_with_error"].ExecuteTemplate(w, "layout", tmplData)
		}
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
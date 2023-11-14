package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"strconv"
	"time"
	"encoding/json"
	"net/url"
	"crypto/sha256"
	"encoding/hex"
)

func hashToLength7(input string) string {
	// Hash the input string using SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	// Convert the hash to a hex-encoded string
	hashString := hex.EncodeToString(hashBytes)

	// Truncate the hash string to the first 7 characters
	if len(hashString) > 7 {
		hashString = hashString[:7]
	}

	return hashString
}


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
	// Render the template
	tmpl["event"].Execute(w, contextData)
}

//Check whether the email is valid
func isValidEmail(email string) bool {
	//Check if the email is valid 
	if len(email) <= 8 || len(email) >= 50 {
		return false
	}
	//Email should end with "yale.edu"
	if email[len(email)-8:] != "yale.edu" {
		return false
	}
	return true
}


func eventRSVPController(w http.ResponseWriter, r *http.Request) {
	Errors := "" // Create a string to store errors
	Confirmation_Code := "" // Create a string to store confirmation code
	//RSVP_Response - encapsulates information about an event rsvp
	type RSVP_Response struct {
		Errors string
		Confirmation_Code string
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Location  string    `json:"location"`
		Image     string    `json:"image"`
		Date      time.Time `json:"date"`
		Attending []string  `json:"attending"`
	}
	
	idStr := chi.URLParam(r, "id") // Retrieve the "id" parameter from the URL as a string
    // Convert "idStr" to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil{
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    // Now, you have "id" as an integer
	theEvent, bool := getEventByID(id)
	//bool should always be true because we have already checked it in eventController
	if bool != true {
		//show error message that the event is not found
		Errors += "Event not found! "
	}
	contextData := RSVP_Response{
		ID:        theEvent.ID,
		Title:     theEvent.Title,
		Date:      theEvent.Date,
		Image:     theEvent.Image,
		Location:  theEvent.Location,
		Attending: theEvent.Attending,
		Confirmation_Code: Confirmation_Code,
		Errors: Errors,
	}

	//get the attendee's email
	email := r.FormValue("email")
	// Check if the email is valid
	if !isValidEmail(email){
		//show error message that the event is open to yale students only
		Errors += "This event is open to Yale students only! "
	}
	// if the email is not valid, render the form again with error messages
	if Errors != "" {
		// If there are errors, render the form again with error messages
		contextData.Errors = Errors
		tmpl["event_with_error"].Execute(w, contextData)
		return
	}
	// Adds an attendee to an event
	err = addAttendee(id, email )
	//RSVPerr should always be nil because we have already checked it in eventController
	//Maybe we should return an error message if the event is full, rather than just show the event is not found
	if err != nil {
		//show error message that the event is full
		Errors += "This event is full! "
	}else{
		contextData.Confirmation_Code = hashToLength7(email)
	}
	//Refech the event information for updating the attending list
	theEvent, bool = getEventByID(id)
	//bool should always be true because we have already checked it in eventController
	if bool != true {
		//show error message that the event is not found
		Errors += "Event not found! "
	}
	contextData.Attending = theEvent.Attending

	//TODO: check if the email is already in the list
	//TODO: check if the event is full
	//TODO: check if the event is in the past
	//TODO: check if the event is open to yale students only
	//TODO: render error message if any of the above is true
	tmpl["event_with_confirmation"].Execute(w, contextData)
}

func isValidImageURL(url string) bool {
	//Check if the image is valid 
	//Image should have file types of ".png", ".jpg", ".jpeg", ".gif", or ".gifv"
	if url[len(url)-4:] != ".png" && url[len(url)-4:] != ".jpg" && url[len(url)-5:] != ".jpeg" && url[len(url)-4:] != ".gif" && url[len(url)-5:] != ".gifv" {
		return false
	}
	return true
}

// CreateEventController is the controller for the events/new page.
func CreateEventController(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodPost {
		Errors := "" // Create a string to store errors
        // It's a POST request, process the form data
        dateStr := r.FormValue("date")
		// Parse the date
		timelayout := "2006-01-02T15:04"
		date, err := time.Parse(timelayout, dateStr)
        if dateStr == "" {
			Errors += "Date cannot be empty! "
        }else{
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
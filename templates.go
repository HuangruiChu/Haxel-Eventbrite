package main

import (
	"html/template"
)

var tmpl = make(map[string]*template.Template)

func init() {
	m := template.Must
	p := template.ParseFiles
	tmpl["index"] = m(p("templates/index.gohtml", "templates/layout.gohtml"))
	tmpl["about"] = m(p("templates/about.gohtml", "templates/layout.gohtml"))
	tmpl["event"] = m(p("templates/event.gohtml", "templates/layout.gohtml"))
	tmpl["event_with_error"] = m(p("templates/event_with_error.gohtml", "templates/layout.gohtml"))
	tmpl["event_with_confirmation"] = m(p("templates/event_with_confirmation.gohtml", "templates/layout.gohtml"))
	tmpl["create_event"] = m(p("templates/create_event.gohtml", "templates/layout.gohtml"))
	tmpl["create_event_with_error"] = m(p("templates/create_event_with_error.gohtml", "templates/layout.gohtml"))
	tmpl["event_donation"] = m(p("templates/event_donation.gohtml", "templates/layout.gohtml"))
}

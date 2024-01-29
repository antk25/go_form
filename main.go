package main

import (
	"html/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	//	ToDo - load templates here
}

func main() {
	loadTemplates()
}

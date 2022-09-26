package models

import "github.com/hamza-starcevic/bookings/pkg/forms"

// * Templatedata holds data sent from handlers to templates
type Templatedata struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

package forms

import (
	"net/http"
	"net/url"
)

// Form creates a function form structure that embeds URL values
type Form struct {
	url.Values
	Errors errors
}

// New initialize the Form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

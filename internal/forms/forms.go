package forms

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Forms creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

// Required checks all the required given fields to see if they are populated
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("The field %s is required", field))
		}
	}
}

// Required checks all the required given fields to see if they are populated
func (f *Form) IsNumber(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		number, err := strconv.Atoi(value)
		if err != nil {
			f.Errors.Add(field, fmt.Sprintf("The field %s must be a number, but received %d", field, number))
		}
	}
}

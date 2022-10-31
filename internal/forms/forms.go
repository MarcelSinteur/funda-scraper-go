package forms

import (
	"fmt"
	"net/url"
	"strings"
)

// Forms creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
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

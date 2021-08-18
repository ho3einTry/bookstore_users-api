package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

type Address struct {
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.Required, validation.Length(5, 20)),
		validation.Field(&a.City, validation.Required, validation.Length(5, 20)),
		validation.Field(&a.State, validation.Required, validation.Length(5, 20)),
		validation.Field(&a.PhoneNumber, validation.Required, validation.Length(11, 11),
			is.Digit, validation.Match(regexp.MustCompile("^09[0-9]{9}$"))))
}

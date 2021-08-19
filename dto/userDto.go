package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/ho3einTry/bookstore_users-api/domain"
	"regexp"
)

const (
	passwordRegex = "^(?=.*[A-Z].*[A-Z])(?=.*[!@#$&*])(?=.*[0-9].*[0-9])(?=.*[a-z].*[a-z].*[a-z]).{8,}$"
	statusRegex   = "^(active|inactive)$"
)

type UserDto struct {
	Id          int64            `json:"id,omitempty"`
	FirstName   string           `json:"first_name,omitempty"`
	LastName    string           `json:"last_name,omitempty"`
	Email       string           `json:"email,omitempty"`
	Addresses   []domain.Address `json:"addresses,omitempty"`
	DateCreated string           `json:"date_created,omitempty"`
	Status      string           `json:"status,omitempty"`
	Password    string           `json:"password,omitempty"`
}

func (ud *UserDto) ToUser() *domain.User {
	return &domain.User{
		Id:          ud.Id,
		FirstName:   ud.FirstName,
		LastName:    ud.LastName,
		Email:       ud.Email,
		Addresses:   ud.Addresses,
		DateCreated: ud.DateCreated,
		Status:      ud.Status,
		Password:    ud.Password,
	}
}

//func (ud *domain.User) FromUser()  *UserDto {
//
//}

func (ud UserDto) Validate() error {
	err := validation.ValidateStruct(&ud,
		validation.Field(&ud.Id, is.Digit),
		validation.Field(&ud.FirstName, validation.Required, validation.Length(3, 20)),
		validation.Field(&ud.LastName, validation.Required, validation.Length(3, 20)),
		validation.Field(&ud.Email, validation.Required, is.Email),
		validation.Field(&ud.Addresses),
		validation.Field(&ud.DateCreated, validation.Date("YYYY-MM-DD")),
		validation.Field(&ud.Status, validation.Match(regexp.MustCompile(statusRegex))),
		//validation.Field(ud.Password, validation.Match(regexp.MustCompile(passwordRegex))),

	)
	return err
}

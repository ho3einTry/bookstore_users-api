package viewModel

import "github.com/ho3einTry/bookstore_users-api/domain"

type UserDto struct {
	Id          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	Status      string `json:"status,omitempty"`
}

func (ud *UserDto) FromUser(user *domain.User) {
	ud.Id = user.Id
	ud.FirstName = user.FirstName
	ud.LastName = user.LastName
	ud.Email = user.Email
	ud.DateCreated = user.DateCreated
	ud.Status = user.Status
}

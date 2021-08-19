package domain

type User struct {
	Id          int64     `json:"id,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Addresses   []Address `json:"addresses,omitempty"`
	DateCreated string    `json:"date_created,omitempty"`
	Status      string    `json:"status,omitempty"`
	Password    string    `json:"password,omitempty"`
}

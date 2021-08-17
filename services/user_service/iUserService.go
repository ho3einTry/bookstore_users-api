package user_service

type iUserService interface {
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
	SearchUser()
}

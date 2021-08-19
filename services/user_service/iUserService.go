package user_service

import (
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/viewModel"
)

type IUserService interface {
	GetUser(id *int64) (*viewModel.UserDto, *exceptions.AppException)
	CreateUser(userDto *dto.UserDto) (int64, *exceptions.AppException)
	UpdateUser()
	DeleteUser()
	SearchUser()
}

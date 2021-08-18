package user_service

import (
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/viewModel"
)

type UserService struct {
}

func (us *UserService) GetUser(id *int64) (*viewModel.UserDto, *exceptions.AppException) {

	panic("implement me")
}

func (us *UserService) CreateUser(userDto *dto.UserDto) (*viewModel.UserDto, *exceptions.AppException) {

	//todo validation

	//todo call repository

	//todo map user to user view model

	//todo return view model
	panic("implement me")
}

func (us *UserService) UpdateUser() {
	panic("implement me")
}

func (us *UserService) DeleteUser() {
	panic("implement me")
}

func (us *UserService) SearchUser() {
	panic("implement me")
}

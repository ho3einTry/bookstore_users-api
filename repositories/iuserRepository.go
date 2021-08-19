package repositories

import (
	"github.com/ho3einTry/bookstore_users-api/domain"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
)

type IUserRepository interface {
	Create(user *domain.User) (int64, *exceptions.AppException)
	Get(id *int64) (*domain.User, *exceptions.AppException)
	Find(email *string) (*domain.User, *exceptions.AppException)
}

//GetUser(id *int64) (*viewModel.UserDto, *exceptions.AppException)
//CreateUser(userDto *dto.UserDto) (int64, *exceptions.AppException)

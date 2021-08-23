package user_service

import (
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/repositories"
	"github.com/ho3einTry/bookstore_users-api/viewModel"
)

var userRepository repositories.IUserRepository = &repositories.UserRepository{}

type UserService struct {
}

func (us *UserService) GetUser(id *int64) (*viewModel.UserDto, *exceptions.AppException) {

	panic(id)
}

func (us *UserService) CreateUser(userDto *dto.UserDto) (int64, *exceptions.AppException) {

	//todo validation

	//todo call find repository
	user, err := userRepository.Find(&userDto.Email)
	if err != nil {
		//appErr := exceptions.NewAppException("db_error_user_find", "an error occurred inside user repository find method")
		return 0, err
	}
	if user != nil {
		err := exceptions.NewAppException("user_already_exists", "user is already exists")
		return 0, &err
	}

	//todo map user dto to user
	user = userDto.ToUser()

	userId, err := userRepository.Create(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
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

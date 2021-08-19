package repositories

import (
	"fmt"
	"github.com/ho3einTry/bookstore_users-api/domain"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
)

var ()

type UserRepository struct {
}

func (ur *UserRepository) Create(user *domain.User) (int64, *exceptions.AppException) {
	fmt.Println("Create")
	return 0, nil
}

func (ur *UserRepository) Get(id *int64) (*domain.User, *exceptions.AppException) {
	fmt.Println("get")
	return nil, nil
}
func (ur *UserRepository) Find(email *string) (*domain.User, *exceptions.AppException) {
	fmt.Println("find")
	return nil, nil
}

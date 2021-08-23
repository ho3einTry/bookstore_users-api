package repositories

import (
	"github.com/ho3einTry/bookstore_users-api/Logger"
	"github.com/ho3einTry/bookstore_users-api/data_sources/mysql"
	"github.com/ho3einTry/bookstore_users-api/domain"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created, status,password FROM users WHERE id=?;"
	queryFindByEmail = "SELECT id, first_name, last_name, email, date_created, status,password FROM users WHERE email=?;"
)

type UserRepository struct {
}

func (ur *UserRepository) Create(user *domain.User) (int64, *exceptions.AppException) {
	statement, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		Logger.Error("error when trying to prepare save user statement", err)
		exception := exceptions.NewAppException("db_error", "error when tying to create user")
		return 0, &exception

	}
	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		Logger.Error("error when trying to save user", saveErr)
		exception := exceptions.NewAppException("db_error", "error when tying to create user")
		return 0, &exception
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		Logger.Error("error when trying to get last insert id after creating a new user", err)
		exception := exceptions.NewAppException("db_error", "error when tying to create user")
		return 0, &exception
	}
	return userId, nil
}
func (ur *UserRepository) Get(id *int64) (*domain.User, *exceptions.AppException) {
	statement, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		Logger.Error("error when trying to prepare get user statement", err)
		exception := exceptions.NewAppException("db_error", "error accrued on get user repository")
		return nil, &exception
	}

	defer statement.Close()

	result := statement.QueryRow(id)
	var user domain.User
	if getErr := result.Scan(
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	); getErr != nil {
		Logger.Error("error when trying to get user by id", getErr)
		exception := exceptions.NewAppException("db_error", "error accrued on get user repository")
		return nil, &exception
	}
	return &user, nil
}
func (ur *UserRepository) Find(email *string) (*domain.User, *exceptions.AppException) {
	statement, err := users_db.Client.Prepare(queryFindByEmail)
	if err != nil {
		Logger.Error("error when trying to prepare get user by email and password statement", err)
		exception := exceptions.NewAppException("db_error", "error when tying to find user")
		return nil, &exception
	}
	defer statement.Close()
	result := statement.QueryRow(&email)
	var user domain.User
	if getErr := result.Scan(
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	); getErr != nil {
		Logger.Error("error when trying to get user by email", getErr)
		exception := exceptions.NewAppException("db_error", "error when tying to find user")
		return nil, &exception
	}
	return &user, nil
}

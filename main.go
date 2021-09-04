package main

import (
	"github.com/ho3einTry/bookstore_users-api/app"
	users_db "github.com/ho3einTry/bookstore_users-api/data_sources/mysql"
)

func main() {
	users_db.Init()
	app.StartApplication()
}

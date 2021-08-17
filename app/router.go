package app

import (
	"github.com/gin-gonic/gin"
	ping_controller "github.com/ho3einTry/bookstore_users-api/controllers/ping"
	user_controller "github.com/ho3einTry/bookstore_users-api/controllers/user"
)

var (
	router = gin.Default()
	//router.Static("/assets", "./assets")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")
)

func urlMappers() {
	//Ping controller
	router.GET("/ping", ping_controller.Ping)

	//User Controller
	router.POST("/users", user_controller.Add)
	router.GET("/users/:user_id", user_controller.Get)
	router.PUT("/users/:user_id", user_controller.Update)
	router.PATCH("/users/:user_id", user_controller.Update)
	router.DELETE("/users/:user_id", user_controller.Delete)
	router.GET("/internal/users/search", user_controller.Search)
	router.POST("/users/login", user_controller.Login)
}

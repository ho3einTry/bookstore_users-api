package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/services/user_service"
)

var userService user_service.IUserService = &user_service.UserService{}

func Add(c *gin.Context) {
	var userDto dto.UserDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		err := exceptions.NewAppException("bad_request", "invalid json body")
		c.JSON(400, err.Message)
	}

	if err := userDto.Validate(); err != nil {
		c.JSON(400, err)
	}

	id, err := userService.CreateUser(&userDto)
	if err != nil {
		c.JSON(400, err.Message)
	}
	c.JSON(200, id)

}
func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
func Get(c *gin.Context)    {}
func Search(c *gin.Context) {}
func Login(c *gin.Context)  {}

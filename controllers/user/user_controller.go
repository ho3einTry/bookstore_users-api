package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/services/user_service"
	"strconv"
)

var userService user_service.IUserService = &user_service.UserService{}

func Add(c *gin.Context) {
	var userDto dto.UserDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		err := exceptions.NewAppException("bad_request", "invalid json body")
		c.JSON(400, err.Message)
		return
	}

	if err := userDto.Validate(); err != nil {
		c.JSON(400, err)
		return
	}

	id, err := userService.CreateUser(&userDto)

	if err != nil {
		c.JSON(400, err.Message)
		return
	}
	c.JSON(200, id)

}
func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
func Get(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(400, "User Id must int")
		return
	}

	user, appException := userService.GetUser(&userId)

	if appException != nil {
		c.JSON(400, appException)
		return
	}
	c.JSON(200, user)
}
func Search(c *gin.Context) {}
func Login(c *gin.Context)  {}

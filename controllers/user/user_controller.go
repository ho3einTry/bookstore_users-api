package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ho3einTry/bookstore_users-api/dto"
	"github.com/ho3einTry/bookstore_users-api/exceptions"
	"github.com/ho3einTry/bookstore_users-api/services/user_service"
)

var UserService user_service.IUserService = &user_service.UserService{}

func Add(c *gin.Context) {
	var userDto dto.UserDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		err := exceptions.NewAppException("bad_request", "invalid json body")
		c.JSON(400, err.Message)
	}

	if err := userDto.Validate(); err != nil {
		//b,_:=json.Marshal(err)
		//err:=exceptions.NewAppException("bad_request",string(b))
		c.JSON(400, err)
	}

	//UserService.CreateUser()
}
func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
func Get(c *gin.Context)    {}
func Search(c *gin.Context) {}
func Login(c *gin.Context)  {}

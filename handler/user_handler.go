package handler

import (
	"miniproject/common/dto"
	"miniproject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var reg *dto.RegisterUserDTO

	if err := c.ShouldBindJSON(&reg); err != nil {
		_ = c.Error(err)
		return
	}

	if err := u.userService.Register(c, reg); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, "register succeed")
}

func (u *UserHandler) Login(c *gin.Context) {
	var login *dto.UserLoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		_ = c.Error(err)
		return
	}

	tokenDTO, err := u.userService.Login(c, login)

	if err != nil {
		_ = c.Error(err)
		return
	}

	if tokenDTO.ID != 0 {
		c.Set("user-id", tokenDTO.ID)
	}

	c.JSON(http.StatusAccepted, tokenDTO)
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

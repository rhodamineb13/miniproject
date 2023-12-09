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

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

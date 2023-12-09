package shared

import (
	"miniproject/common/database"
	"miniproject/handler"
	"miniproject/repository"
	"miniproject/service"
)

type Handler struct {
	UserHandler *handler.UserHandler
}

func NewHandler() (*Handler, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepo(db)))

	return &Handler{
		UserHandler: userHandler,
	}, nil
}

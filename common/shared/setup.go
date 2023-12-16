package shared

import (
	"miniproject/handler"
	"miniproject/repository"
	"miniproject/service"
)

type Handler struct {
	UserHandler *handler.UserHandler
	BookHandler *handler.BookHandler
}

func NewHandler() (*Handler, error) {

	userHandler := handler.NewUserHandler(service.NewUserService(DB.r, repository.NewUserRepo(DB.db)))
	bookHandler := handler.NewBookHandler(service.NewBookService(repository.NewBookRepo(DB.db)))

	return &Handler{
		UserHandler: userHandler,
		BookHandler: bookHandler,
	}, nil
}

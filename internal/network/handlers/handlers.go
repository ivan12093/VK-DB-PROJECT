package handlers

import "github.com/ivan12093/VK-DB-PROJECT/internal/usecase"

type Handlers struct {
	UserHandler    *HandlerUsers
	PostHandler    *HandlerPosts
	ForumHandler   *HandlerForum
	ThreadHandler  *HandlerThreads
	ServiceHandler *HandlerServices
}

func NewHandlers(usecases *usecase.Usecases) *Handlers {
	return &Handlers{
		UserHandler:    NewUsersHandler(usecases.UserUsecase),
		PostHandler:    NewPostsHandler(usecases.PostUsecase),
		ForumHandler:   NewForumsHandler(usecases.ForumUsecase),
		ThreadHandler:  NewThreadsHandler(usecases.ThreadUsecase),
		ServiceHandler: NewServicesHandler(usecases.ServiceUsecase),
	}
}

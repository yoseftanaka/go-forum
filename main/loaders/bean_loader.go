package loaders

import (
	"forum/main/config"
	"forum/main/controllers"
	"forum/main/handlers"
	"forum/main/repositories"
)

var (
	UserController *controllers.UserController
	UserHandler    *handlers.UserHandler
	UserRepository *repositories.UserRepository
	PostController *controllers.PostController
	PostHandler    *handlers.PostHandler
	PostRepository *repositories.PostRepository
	AuthHandler    *handlers.AuthHandler
	AuthController *controllers.AuthController
)

func LoadRepositories() {
	UserRepository = repositories.NewUserRepository(config.DB)
	PostRepository = repositories.NewPostRepository(config.DB)
}

func LoadHandlers() {
	UserHandler = handlers.NewUserHandler(UserRepository)
	PostHandler = handlers.NewPostHandler(PostRepository)
	AuthHandler = handlers.NewAuthHandler(UserRepository)
}

func LoadControllers() {
	UserController = controllers.NewUserController(UserHandler)
	PostController = controllers.NewPostController(PostHandler)
	AuthController = controllers.NewAuthController(AuthHandler)
}

func LoadBean() {
	LoadRepositories()
	LoadHandlers()
	LoadControllers()
}

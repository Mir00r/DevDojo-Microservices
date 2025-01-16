package containers

import (
	database "github.com/Mir00r/user-service/db"
	"github.com/Mir00r/user-service/internal/api/controllers"
	"github.com/Mir00r/user-service/internal/repositories"
	"github.com/Mir00r/user-service/internal/services"
)

// Container struct holds all application dependencies
type Container struct {
	UserRepository          repositories.UserRepository
	PublicUserController    *controllers.PublicUserController
	ProtectedUserController *controllers.ProtectedUserController
	InternalUserController  *controllers.InternalUserController
}

// NewContainer initializes all dependencies and returns a Container instance
func NewContainer() *Container {

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	publicUserController := controllers.NewPublicUserController(userService)
	protectedUserController := controllers.NewProtectedUserController(userService)
	internalUserController := controllers.NewInternalUserController(userService)

	return &Container{
		UserRepository: userRepo,

		PublicUserController:    publicUserController,
		ProtectedUserController: protectedUserController,
		InternalUserController:  internalUserController,
	}
}

package containers

import (
	database "github.com/Mir00r/auth-service/db"
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/services"
)

// Container struct holds all application dependencies
type Container struct {
	UserRepository          repositories.UserRepository
	TokenRepository         repositories.TokenRepository
	MFARepository           repositories.MFARepository
	AuthService             services.AuthService
	TokenService            services.TokenServiceInterface
	MFAService              services.MFAService
	PublicAuthController    *controllers.PublicAuthController
	ProtectedAuthController *controllers.ProtectedAuthController
	InternalAuthController  *controllers.InternalAuthController
}

// NewContainer initializes all dependencies and returns a Container instance
func NewContainer() *Container {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)
	tokenRepo := repositories.NewTokenRepository(database.DB)
	mfaRepo := repositories.NewMFARepository(database.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo, tokenRepo)
	internalAuthService := services.NewInternalAuthService(userRepo)
	tokenService := services.NewTokenService(tokenRepo, userRepo)
	mfaService := services.NewMFAService(mfaRepo, userRepo)

	// Initialize controllers
	publicAuthController := controllers.NewPublicAuthController(authService, tokenService)
	protectedAuthController := controllers.NewProtectedAuthController(authService, tokenService, mfaService)
	internalAuthController := controllers.NewInternalAuthController(internalAuthService)

	return &Container{
		UserRepository:          userRepo,
		TokenRepository:         tokenRepo,
		MFARepository:           mfaRepo,
		AuthService:             authService,
		TokenService:            tokenService,
		MFAService:              mfaService,
		PublicAuthController:    publicAuthController,
		ProtectedAuthController: protectedAuthController,
		InternalAuthController:  internalAuthController,
	}
}

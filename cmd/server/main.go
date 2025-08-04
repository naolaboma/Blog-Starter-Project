package main

import (
	"log"
	"net/http"

	"Blog-API/internal/delivery/controllers"
	"Blog-API/internal/delivery/router"
	"Blog-API/internal/infrastructure/database"
	"Blog-API/internal/infrastructure/jwt"
	"Blog-API/internal/infrastructure/middleware"
	"Blog-API/internal/infrastructure/password"
	"Blog-API/internal/repository"
	"Blog-API/internal/usecase"
	"Blog-API/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.Server.GinMode)

	mongoDB, err := database.NewMongoDB(cfg.MongoDB.URI, cfg.MongoDB.Database)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer mongoDB.Close()

	passwordService := password.NewPasswordService()
	jwtService := jwt.NewJWTService(cfg.JWT.Secret, cfg.JWT.AccessExpiry, cfg.JWT.RefreshExpiry)

	userRepo := repository.NewUserRepository(mongoDB)
	sessionRepo := repository.NewSessionRepository(mongoDB)

	userUseCase := usecase.NewUserUseCase(userRepo, passwordService, jwtService, sessionRepo)

	userHandler := controllers.NewUserHandler(userUseCase)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, sessionRepo)

	router := router.SetupRouter(userHandler, authMiddleware)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("MongoDB connected to: %s", cfg.MongoDB.URI)
	log.Printf("Gin mode: %s", cfg.Server.GinMode)

	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 
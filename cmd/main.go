package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"go-scaffold.iserranodev.net/api/routes"
	"go-scaffold.iserranodev.net/infrastructure/db"
	"go-scaffold.iserranodev.net/infrastructure/handler"
	"go-scaffold.iserranodev.net/internal/service"
	"go-scaffold.iserranodev.net/internal/usecase"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Conexión a la base de datos
	conn := db.Connect()
	userRepo := db.NewUserRepository(conn)
	roleRepo := db.NewRoleRepository(conn)
	jwtSvc := service.NewJWTService(os.Getenv("JWT_SECRET"))

	// Casos de uso
	authUC := usecase.NewAuthUseCase(userRepo, roleRepo)
	userUC := usecase.NewUserUseCase(userRepo, roleRepo)
	profileUC := usecase.NewProfileUseCase(userRepo)
	roleUC := usecase.NewRoleUseCase(roleRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUC, jwtSvc)
	userHandler := handler.NewUserHandler(userUC)
	profileHandler := handler.NewProfileHandler(profileUC)
	roleHandler := handler.NewRoleHandler(roleUC)

	// Router
	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
	if allowedOriginsEnv == "" {
		allowedOriginsEnv = "http://localhost:3000" // Default allowed origins
	}

	allowedOrigins := strings.Split(allowedOriginsEnv, ",")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true, // para ver logs de CORS
	})

	r.Use(corsMiddleware.Handler)

	// Rutas
	routes.RegisterAuthRoutes(r, authHandler)
	routes.RegisterUserRoutes(r, userHandler)
	routes.RegisterProfileRoutes(r, profileHandler)
	routes.RegisterRoleRoutes(r, roleHandler)

	// Arrancar servidor
	log.Println("API corriendo en :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

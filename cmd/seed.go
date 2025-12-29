package main

import (
	"fmt"
	"go-react-backend.iserranodev.net/infrastructure/db"
	"go-react-backend.iserranodev.net/internal/domain"
	"go-react-backend.iserranodev.net/internal/helper"
	"log"
)

func main() {
	conn := db.Connect()
	userRepo := db.NewUserRepository(conn)
	roleRepo := db.NewRoleRepository(conn)

	adminRole, err := roleRepo.GetByName("admin")
	if err != nil || adminRole == nil {
		adminRole = &domain.Role{ID: 1, Name: "admin", Label: "Administrador"}
		conn.Create(adminRole)
	}

	passHash, _ := helper.HashPassword("admin123")

	user := &domain.User{
		Email:    "admin@admin.com",
		Password: passHash,
		Roles:    []*domain.Role{adminRole},
	}

	err = userRepo.Create(user)
	if err != nil {
		log.Fatal("Error creando usuario admin:", err)
	}

	fmt.Println("Usuario admin creado con éxito")
}

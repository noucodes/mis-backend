package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/noucodes/mis-backend/config"
	"github.com/noucodes/mis-backend/models"
)

func GetUsers(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	return c.JSON(users)
}
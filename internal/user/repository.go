package user

import (
	"context"

	"github.com/adhtanjung/bored/database"
	"github.com/adhtanjung/bored/internal/model"
	"github.com/gofiber/fiber"
)

type Repository interface {
	Create(ctx context.Context, user model.User) error
}

// repository persists albums in database
type repository struct {
	db *database.Dbinstance
}

func (r repository) Create(ctx context.Context, user model.User) error {
	// return r.db.With(ctx).Model(&album).Insert()
	// Check if username is already taken, if true return error
	if err := r.db.Db.Where("username= ?", &user.Username).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Username already taken", "data": err})
	}
}

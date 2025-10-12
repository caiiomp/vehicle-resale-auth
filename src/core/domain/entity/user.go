package entity

import (
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/valueObjects"
)

type User struct {
	ID           string
	Name         string
	Email        string
	Password     string
	PasswordHash string
	Role         valueObjects.RoleType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

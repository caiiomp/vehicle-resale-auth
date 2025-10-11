package model

import (
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
)

type User struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string    `json:"name,omitempty" bson:"name,omitempty"`
	Email        string    `json:"email,omitempty" bson:"email,omitempty"`
	PasswordHash string    `json:"password_hash,omitempty" bson:"password_hash,omitempty"`
	Role         string    `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func UserFromDomain(user entity.User) User {
	return User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func (ref User) ToDomain() *entity.User {
	return &entity.User{
		ID:           ref.ID,
		Name:         ref.Name,
		Email:        ref.Email,
		Role:         ref.Role,
		PasswordHash: ref.PasswordHash,
		CreatedAt:    ref.CreatedAt,
		UpdatedAt:    ref.UpdatedAt,
	}
}

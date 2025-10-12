package userApi

import (
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	valueoObjects "github.com/caiiomp/vehicle-resale-auth/src/core/domain/valueObjects"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func (ref createUserRequest) ToDomain() *entity.User {
	return &entity.User{
		Name:     ref.Name,
		Email:    ref.Email,
		Password: ref.Password,
		Role:     valueoObjects.RoleType{Value: ref.Role},
	}
}

type userResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func userResponseFromDomain(user entity.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role.Value,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type userURI struct {
	ID string `uri:"user_id"`
}

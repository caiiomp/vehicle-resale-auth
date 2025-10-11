package userApi

import (
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (ref createUserRequest) ToDomain() *entity.User {
	return &entity.User{
		Name:     ref.Name,
		Email:    ref.Email,
		Password: ref.Password,
		Role:     ref.Role,
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
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type userURI struct {
	ID string `uri:"user_id"`
}

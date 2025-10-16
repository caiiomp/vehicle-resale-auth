package userApi

import (
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ref createUserRequest) ToDomain() *entity.User {
	return &entity.User{
		Name:     ref.Name,
		Email:    ref.Email,
		Password: ref.Password,
	}
}

type userURI struct {
	ID string `uri:"user_id"`
}

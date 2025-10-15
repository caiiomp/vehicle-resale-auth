package userApi

import (
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

type userURI struct {
	ID string `uri:"user_id"`
}

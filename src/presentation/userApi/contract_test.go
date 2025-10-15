package userApi

import (
	"testing"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/valueObjects"
	"github.com/stretchr/testify/assert"
)

func TestToDomain(t *testing.T) {
	request := createUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@email.com",
		Password: "123",
		Role:     "ADMIN",
	}

	expected := &entity.User{
		Name:     "John Doe",
		Email:    "john.doe@email.com",
		Password: "123",
		Role:     valueObjects.RoleTypeAdmin,
	}

	actual := request.ToDomain()

	assert.Equal(t, expected, actual)
}

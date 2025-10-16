package userApi

import (
	"testing"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestToDomain(t *testing.T) {
	request := createUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@email.com",
		Password: "123",
	}

	expected := &entity.User{
		Name:     "John Doe",
		Email:    "john.doe@email.com",
		Password: "123",
	}

	actual := request.ToDomain()

	assert.Equal(t, expected, actual)
}

package userApi

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Test_userResponseFromDomain(t *testing.T) {
	userID := primitive.NewObjectID().Hex()

	now := time.Now()

	user := entity.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john.doe@email.com",
		Role:      valueObjects.RoleTypeAdmin,
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := userResponse{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john.doe@email.com",
		Role:      "ADMIN",
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := userResponseFromDomain(user)

	assert.Equal(t, expected, actual)
}

package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserFromDomain(t *testing.T) {
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

	expected := User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john.doe@email.com",
		Role:      "ADMIN",
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := UserFromDomain(user)

	assert.Equal(t, expected, actual)
}

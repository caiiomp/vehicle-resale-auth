package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
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
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john.doe@email.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := UserFromDomain(user)

	assert.Equal(t, expected, actual)
}

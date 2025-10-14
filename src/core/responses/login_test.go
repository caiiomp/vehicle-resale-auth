package responses

import (
	"testing"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestLoginResponseFromDomain(t *testing.T) {
	auth := entity.Auth{
		AccessToken: "123",
		ExpiresIn:   123,
	}

	expected := LoginResponse{
		AccessToken: "123",
		ExpiresIn:   123,
	}

	actual := LoginResponseFromDomain(auth)

	assert.Equal(t, expected, actual)
}

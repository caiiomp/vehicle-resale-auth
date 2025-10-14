package valueObjects

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRoleTypeValidation(t *testing.T) {
	t.Run("should be valid role type", func(t *testing.T) {
		validate := validator.New()

		roleType := RoleType{
			Value: "ADMIN",
		}

		validate.RegisterStructValidation(RoleTypeValidation, RoleType{})

		err := validate.Struct(roleType)

		assert.Nil(t, err)
	})

	t.Run("should be valid role type when it is empty", func(t *testing.T) {
		validate := validator.New()

		roleType := RoleType{
			Value: "",
		}

		validate.RegisterStructValidation(RoleTypeValidation, RoleType{})

		err := validate.Struct(roleType)

		assert.Nil(t, err)
	})

	t.Run("should be invalid role type", func(t *testing.T) {
		validate := validator.New()

		roleType := RoleType{
			Value: "INVALID",
		}

		validate.RegisterStructValidation(RoleTypeValidation, RoleType{})

		err := validate.Struct(roleType)

		assert.NotNil(t, err)
	})
}

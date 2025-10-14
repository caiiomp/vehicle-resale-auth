package auth

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/caiiomp/vehicle-resale-auth/src/core/_mocks"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not login when failed to get user by email", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(nil, unexpectedError)

		service := NewAuthService(userRepositoryMocked, "xxx")

		actual, err := service.Login(ctx, "john.doe@email.com", "123")

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should not login when user does not exist", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(nil, nil)

		service := NewAuthService(userRepositoryMocked, "xxx")

		actual, err := service.Login(ctx, "john.doe@email.com", "123")

		assert.Nil(t, actual)
		assert.ErrorContains(t, err, "does not exist")
	})

	t.Run("should not login when password does not match", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("111"), bcrypt.DefaultCost)

		user := entity.User{
			Email:        "john.doe@email.com",
			PasswordHash: string(passwordHash),
		}

		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(&user, nil)

		service := NewAuthService(userRepositoryMocked, "xxx")

		actual, err := service.Login(ctx, "john.doe@email.com", "123")

		assert.Nil(t, actual)
		assert.Nil(t, err)
	})

	t.Run("should login successfully", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

		user := entity.User{
			Email:        "john.doe@email.com",
			PasswordHash: string(passwordHash),
		}

		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(&user, nil)

		service := NewAuthService(userRepositoryMocked, "xxx")

		actual, err := service.Login(ctx, "john.doe@email.com", "123")

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

package user

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks "github.com/caiiomp/vehicle-resale-auth/src/core/_mocks"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not create an user when failed to get user by email", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "john.doe@email.com",
			Password: "123",
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(nil, unexpectedError)

		service := NewUserService(validator.New(), userRepositoryMocked)

		actual, err := service.Create(ctx, user)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
		userRepositoryMocked.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should not create an user when email already exists", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "john.doe@email.com",
			Password: "123",
		}

		userRepositoryMocked := mocks.NewUserRepository(t)
		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(&user, nil)

		service := NewUserService(validator.New(), userRepositoryMocked)

		actual, err := service.Create(ctx, user)

		assert.Nil(t, actual)
		assert.ErrorContains(t, err, "already exists")
		userRepositoryMocked.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("should create an user successfully", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "john.doe@email.com",
			Password: "123",
		}

		userRepositoryMocked := mocks.NewUserRepository(t)

		userRepositoryMocked.On("GetByEmail", ctx, "john.doe@email.com").
			Return(nil, nil)
		userRepositoryMocked.On("Create", ctx, mock.AnythingOfType("entity.User")).
			Return(&user, nil)

		service := NewUserService(validator.New(), userRepositoryMocked)

		actual, err := service.Create(ctx, user)

		assert.NotNil(t, actual)
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	ctx := context.TODO()
	userID := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not get an user by id when failed to get user by id", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		userRepositoryMocked.On("GetByID", ctx, userID).
			Return(nil, unexpectedError)

		service := NewUserService(validator.New(), userRepositoryMocked)

		actual, err := service.GetByID(ctx, userID)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should get an user by id successfully", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		now := time.Now()

		user := entity.User{
			ID:           userID,
			Name:         "John Doe",
			Email:        "john.doe@email.com",
			Password:     "123",
			PasswordHash: "xxx",
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		userRepositoryMocked.On("GetByID", ctx, userID).
			Return(&user, nil)

		service := NewUserService(validator.New(), userRepositoryMocked)

		expected := user

		actual, err := service.GetByID(ctx, userID)

		assert.Equal(t, &expected, actual)
		assert.Nil(t, err)
	})
}

func TestSearch(t *testing.T) {
	ctx := context.TODO()
	userID := primitive.NewObjectID().Hex()
	unexpectedError := errors.New("unexpected error")

	t.Run("should not search users when failed to search users", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		userRepositoryMocked.On("Search", ctx).
			Return(nil, unexpectedError)

		service := NewUserService(validator.New(), userRepositoryMocked)

		actual, err := service.Search(ctx)

		assert.Nil(t, actual)
		assert.Equal(t, unexpectedError, err)
	})

	t.Run("should search users successfully", func(t *testing.T) {
		userRepositoryMocked := mocks.NewUserRepository(t)

		now := time.Now()

		users := []entity.User{
			{
				ID:           userID,
				Name:         "John Doe",
				Email:        "john.doe@email.com",
				Password:     "123",
				PasswordHash: "xxx",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		}

		userRepositoryMocked.On("Search", ctx).
			Return(users, nil)

		service := NewUserService(validator.New(), userRepositoryMocked)

		expected := users

		actual, err := service.Search(ctx)

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})
}

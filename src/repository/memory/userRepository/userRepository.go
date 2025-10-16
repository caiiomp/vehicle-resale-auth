package userRepository

import (
	"context"
	"time"

	interfaces "github.com/caiiomp/vehicle-resale-auth/src/core/_interfaces"
	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/model"
	"github.com/google/uuid"
)

type userRepository struct {
	users []model.User
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepository{
		users: []model.User{},
	}
}

func (ref *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	record := model.UserFromDomain(user)

	record.ID = uuid.NewString()

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	ref.users = append(ref.users, record)

	for _, user := range ref.users {
		if user.ID == record.ID {
			return user.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	for _, user := range ref.users {
		if user.ID == id {
			return user.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, user := range ref.users {
		if user.Email == email {
			return user.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *userRepository) Search(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, len(ref.users))

	for i, user := range ref.users {
		users[i] = *user.ToDomain()
	}

	return users, nil
}

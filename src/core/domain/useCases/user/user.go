package user

import (
	"context"
	"fmt"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/userRepository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Search(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, id string, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepository userRepository.UserRepository
}

func NewUserRepository(userRepository userRepository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (ref *userService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	existingUser, err := ref.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("email %s already exists", user.Email)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = string(passwordHash)

	return ref.userRepository.Create(ctx, user)
}

func (ref *userService) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return ref.userRepository.GetByID(ctx, id)
}

func (ref *userService) Search(ctx context.Context) ([]entity.User, error) {
	return ref.userRepository.Search(ctx)
}

func (ref *userService) Update(ctx context.Context, id string, user entity.User) (*entity.User, error) {
	return ref.userRepository.Update(ctx, id, user)
}

func (ref *userService) Delete(ctx context.Context, id string) error {
	return ref.userRepository.Delete(ctx, id)
}

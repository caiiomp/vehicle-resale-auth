package user

import (
	"context"
	"fmt"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/userRepository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Search(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	validate       *validator.Validate
	userRepository userRepository.UserRepository
}

func NewUserService(validate *validator.Validate, userRepository userRepository.UserRepository) UserService {
	return &userService{
		validate:       validate,
		userRepository: userRepository,
	}
}

func (ref *userService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	if err := ref.validate.Struct(user); err != nil {
		return nil, err
	}

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

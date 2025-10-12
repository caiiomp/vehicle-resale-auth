package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/userRepository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
}

type authService struct {
	userRepository userRepository.UserRepository
	jwtSecretKey   string
}

func NewAuthService(userRepository userRepository.UserRepository, jwtSecretKey string) AuthService {
	return &authService{
		userRepository: userRepository,
		jwtSecretKey:   jwtSecretKey,
	}
}

func (ref *authService) Login(ctx context.Context, email string, password string) (*entity.Auth, error) {
	user, err := ref.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("email '%s' does not exist", email)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, nil
		}
		return nil, err
	}

	expiresIn := time.Now().Add(time.Hour * 24).Unix()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Value,
		"exp":     expiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ref.jwtSecretKey))
	if err != nil {
		return nil, err
	}

	return &entity.Auth{
		AccessToken: tokenString,
		ExpiresIn:   expiresIn,
	}, nil
}

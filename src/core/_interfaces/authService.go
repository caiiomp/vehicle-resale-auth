package interfaces

import (
	"context"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*entity.Auth, error)
}

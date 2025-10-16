package interfaces

import (
	"context"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"
)

type UserService interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Search(ctx context.Context) ([]entity.User, error)
}

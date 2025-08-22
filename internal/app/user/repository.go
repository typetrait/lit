package user

import (
	"context"

	"github.com/typetrait/lit/internal/domain/user"
)

type Repository interface {
	FindAll(ctx context.Context) ([]user.User, error)
	FindByID(ctx context.Context, id int64) (user.User, error)
	Create(ctx context.Context, user user.User) (user.User, error)
	Update(ctx context.Context, user user.User) error
	Delete(ctx context.Context, user user.User) error
}

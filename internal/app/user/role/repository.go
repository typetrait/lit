package role

import (
	"context"

	"github.com/typetrait/lit/internal/domain/user"
)

type Repository interface {
	FindAll(ctx context.Context) ([]user.Role, error)
	FindByNames(ctx context.Context, roleNames []string) ([]user.Role, error)
	FindByName(ctx context.Context, roleName string) (user.Role, error)
}

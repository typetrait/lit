package settings

import (
	"context"

	"github.com/typetrait/lit/internal/domain/settings"
)

type Repository interface {
	FindAll(ctx context.Context) (settings.Settings, error)
}

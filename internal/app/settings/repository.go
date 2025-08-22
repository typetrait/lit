package settings

import "github.com/typetrait/lit/internal/domain/settings"

type Repository interface {
	FindAll() ([]settings.Settings, error)
}

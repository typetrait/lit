package app

import "github.com/typetrait/lit/internal/domain/settings"

type SettingsProvider interface {
	Settings() settings.Settings
}

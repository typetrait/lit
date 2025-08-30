package settings

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/settings"
)

type Provider struct {
	repository Repository
}

func NewProvider(repository Repository) *Provider {
	return &Provider{
		repository: repository,
	}
}

func (p *Provider) Settings(ctx context.Context) (settings.Settings, error) {
	s, err := p.repository.FindAll(ctx)
	if err != nil {
		return settings.Settings{}, fmt.Errorf("loading settings: %w", err)
	}
	return s, nil
}

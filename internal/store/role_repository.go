package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/typetrait/lit/internal/domain/user"
	"github.com/typetrait/lit/internal/store/model"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (rr *RoleRepository) FindAll(ctx context.Context) ([]user.Role, error) {
	var roleModels []model.Role
	result := rr.db.WithContext(ctx).Find(&roleModels)
	if result.Error != nil {
		return nil, fmt.Errorf("finding roles in repository: %w", result.Error)
	}
	roles := make([]user.Role, len(roleModels))
	for _, r := range roleModels {
		roles = append(roles, r.ToDomainRole())
	}
	return roles, nil
}

func (rr *RoleRepository) FindByNames(ctx context.Context, roleNames []string) ([]user.Role, error) {
	if len(roleNames) == 0 {
		return []user.Role{}, nil
	}

	var models []model.Role
	if err := rr.db.WithContext(ctx).
		Where("name IN ?", roleNames).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("finding roles by names in repository: %w", err)
	}

	byName := make(map[string]user.Role, len(models))
	for _, m := range models {
		d := m.ToDomainRole()
		byName[d.Name] = d
	}
	out := make([]user.Role, 0, len(models))
	for _, name := range roleNames {
		if r, ok := byName[name]; ok {
			out = append(out, r)
		}
	}
	return out, nil
}

func (rr *RoleRepository) FindByName(ctx context.Context, roleName string) (user.Role, error) {
	var m model.Role
	err := rr.db.WithContext(ctx).
		Where("name = ?", roleName).
		First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.Role{}, err
		}
		return user.Role{}, fmt.Errorf("finding role %q in repository: %w", roleName, err)
	}
	return m.ToDomainRole(), nil
}

package store

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/user"
	"github.com/typetrait/lit/internal/store/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindAll(ctx context.Context) ([]user.User, error) {
	var userModels []model.User
	result := ur.db.WithContext(ctx).Order("created_at DESC").Find(&userModels)
	if result.Error != nil {
		return nil, fmt.Errorf("finding users in repository: %w", result.Error)
	}
	users := make([]user.User, len(userModels))
	for _, u := range userModels {
		users = append(users, u.ToDomainUser())
	}
	return users, nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id int64) (user.User, error) {
	var userModel model.User
	if err := ur.db.WithContext(ctx).First(&userModel, "id = ?", id).Error; err != nil {
		return user.User{}, fmt.Errorf("finding user by ID in repository: %w", err)
	}
	return userModel.ToDomainUser(), nil
}

func (ur *UserRepository) Create(ctx context.Context, userToCreate user.User) (user.User, error) {
	userModel := model.FromDomainUser(userToCreate)
	if err := ur.db.WithContext(ctx).Create(&userModel).Error; err != nil {
		return user.User{}, fmt.Errorf("creating user in repository: %w", err)
	}
	return userModel.ToDomainUser(), nil
}

func (ur *UserRepository) Update(ctx context.Context, user user.User) error {
	if err := ur.db.WithContext(ctx).Save(&user).Error; err != nil {
		return fmt.Errorf("updating user in repository: %w", err)
	}
	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, user user.User) error {
	if err := ur.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return fmt.Errorf("deleting user in repository: %w", err)
	}
	return nil
}

package model

import (
	"time"

	"github.com/typetrait/lit/internal/domain/user"
)

type User struct {
	ID          int64     `gorm:"primaryKey"`
	Email       string    `gorm:"unique;not null"`
	DisplayName string    `gorm:"unique;not null"`
	Roles       []Role    `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE;"`
	IsActive    bool      `gorm:"not null;default:true"`
	CreatedAt   time.Time `gorm:"not null"`
}

func FromDomainUser(user user.User) User {
	var roles []Role
	for _, role := range user.Roles {
		roles = append(roles, FromDomainRole(role))
	}
	return User{
		ID:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Roles:       roles,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt,
	}
}

func (u User) ToDomainUser() user.User {
	var roles []user.Role
	for _, role := range u.Roles {
		roles = append(roles, role.ToDomainRole())
	}
	return user.User{
		ID:          u.ID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		Roles:       roles,
		IsActive:    u.IsActive,
		CreatedAt:   u.CreatedAt,
	}
}

package model

import (
	"github.com/typetrait/lit/internal/domain/user"
)

type Role struct {
	ID          int64        `gorm:"primaryKey;auto_increment"`
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;"`
}

func FromDomainRole(role user.Role) Role {
	var permissions []Permission
	for _, p := range role.Permissions {
		permissions = append(permissions, Permission{
			Key: p.Key,
		})
	}
	return Role{
		Name:        role.Name,
		Permissions: permissions,
	}
}

func (r Role) ToDomainRole() user.Role {
	var permissions []user.Permission
	for _, p := range r.Permissions {
		permissions = append(permissions, user.Permission{
			Key: p.Key,
		})
	}
	return user.Role{
		Name:        r.Name,
		Permissions: permissions,
	}
}

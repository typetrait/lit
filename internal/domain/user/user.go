package user

import (
	"time"
)

type User struct {
	ID          int64
	Email       string
	DisplayName string
	Roles       []Role
	IsActive    bool
	CreatedAt   time.Time
}

func (u *User) HasPermission(required Permission) bool {
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if perm == required {
				return true
			}
		}
	}
	return false
}

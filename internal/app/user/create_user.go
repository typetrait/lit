package user

import (
	"context"
	"fmt"
	"time"

	"github.com/typetrait/lit/internal/app/user/role"
	"github.com/typetrait/lit/internal/domain/user"
)

type CreateUser struct {
	userRepository Repository
	roleRepository role.Repository
}

func NewCreateUser(userRepository Repository) *CreateUser {
	return &CreateUser{
		userRepository: userRepository,
	}
}

func (cu *CreateUser) CreateUser(ctx context.Context, createUserCommand CreateUserCommand) (user.User, error) {
	email, err := user.NewEmail(createUserCommand.Email)
	if err != nil {
		return user.User{}, fmt.Errorf("deriving user email: %w", err)
	}

	roles, err := cu.roleRepository.FindByNames(ctx, createUserCommand.Roles)
	if err != nil {
		return user.User{}, fmt.Errorf("finding user roles: %w", err)
	}

	userToCreate := user.User{
		Email:       email.String(),
		DisplayName: createUserCommand.DisplayName,
		Roles:       roles,
		CreatedAt:   time.Now(),
	}

	createdUser, err := cu.userRepository.Create(ctx, userToCreate)
	if err != nil {
		return user.User{}, fmt.Errorf("creating user: %w", err)
	}

	return createdUser, nil
}

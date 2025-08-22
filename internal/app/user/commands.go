package user

type CreateUserCommand struct {
	Email       string
	DisplayName string
	Roles       []string
}

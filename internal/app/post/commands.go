package post

import "github.com/typetrait/lit/internal/domain/user"

type CreatePostCommand struct {
	Title   string
	Content string
	Author  user.User
	Tags    []string
}

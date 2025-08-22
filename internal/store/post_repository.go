package store

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/store/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) Create(ctx context.Context, postToCreate post.Post) (post.Post, error) {
	postModel := model.FromDomainPost(postToCreate)
	if err := pr.db.WithContext(ctx).Create(&postModel).Error; err != nil {
		return post.Post{}, fmt.Errorf("creating post in repository: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

func (pr *PostRepository) FindAll(ctx context.Context) ([]post.Post, error) {
	var postModels []model.Post
	if err := pr.db.WithContext(ctx).Find(&postModels).Error; err != nil {
		return []post.Post{}, fmt.Errorf("finding all posts: %w", err)
	}

	var posts []post.Post
	for _, m := range postModels {
		posts = append(posts, m.ToDomainPost())
	}

	return posts, nil
}

func (pr *PostRepository) FindBySlug(ctx context.Context, slug string) (post.Post, error) {
	var postModel model.Post
	if err := pr.db.WithContext(ctx).First(&postModel, "slug = ?", slug).Error; err != nil {
		return post.Post{}, fmt.Errorf("finding post by slug: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

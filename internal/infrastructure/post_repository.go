package infrastructure

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/infrastructure/model"
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

func (pr *PostRepository) Update(ctx context.Context, postToUpdate post.Post) (post.Post, error) {
	postModel := model.FromDomainPost(postToUpdate)
	if err := pr.db.WithContext(ctx).Save(&postModel).Error; err != nil {
		return post.Post{}, fmt.Errorf("updating post in repository: %w", err)
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

func (pr *PostRepository) FindByID(ctx context.Context, id int64) (post.Post, error) {
	var postModel model.Post
	if err := pr.db.WithContext(ctx).First(&postModel, "id = ?", id).Error; err != nil {
		return post.Post{}, fmt.Errorf("finding post by ID: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

func (pr *PostRepository) FindBySlug(ctx context.Context, slug string) (post.Post, error) {
	var postModel model.Post
	if err := pr.db.WithContext(ctx).First(&postModel, "slug = ?", slug).Error; err != nil {
		return post.Post{}, fmt.Errorf("finding post by slug: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

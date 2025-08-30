package repository

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

func (r *PostRepository) Create(ctx context.Context, postToCreate post.Post) (post.Post, error) {
	postModel := model.FromDomainPost(postToCreate)
	if err := r.db.WithContext(ctx).Create(&postModel).Error; err != nil {
		return post.Post{}, fmt.Errorf("creating post in repository: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

func (r *PostRepository) Update(ctx context.Context, postToUpdate post.Post) (post.Post, error) {
	postModel := model.FromDomainPost(postToUpdate)
	if err := r.db.WithContext(ctx).Save(&postModel).Error; err != nil {
		return post.Post{}, fmt.Errorf("updating post in repository: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

func (r *PostRepository) FindAll(ctx context.Context) ([]post.Post, error) {
	var postModels []model.Post
	if err := r.db.WithContext(ctx).Find(&postModels).Error; err != nil {
		return []post.Post{}, fmt.Errorf("finding all posts: %w", err)
	}

	var posts []post.Post
	for _, m := range postModels {
		posts = append(posts, m.ToDomainPost())
	}

	return posts, nil
}

func (r *PostRepository) FindAllPublished(ctx context.Context) ([]post.Post, error) {
	var postModels []model.Post
	if err := r.db.WithContext(ctx).Find(&postModels, "status = ?", "published").Error; err != nil {
		return []post.Post{}, fmt.Errorf("finding all published posts: %w", err)
	}

	var posts []post.Post
	for _, m := range postModels {
		posts = append(posts, m.ToDomainPost())
	}

	return posts, nil
}

func (r *PostRepository) FindByID(ctx context.Context, id int64) (post.Post, error) {
	var postModel model.Post
	if err := r.db.WithContext(ctx).First(&postModel, "id = ?", id).Error; err != nil {
		return post.Post{}, fmt.Errorf("finding post by ID: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

func (r *PostRepository) FindBySlug(ctx context.Context, slug string) (post.Post, error) {
	var postModel model.Post
	if err := r.db.WithContext(ctx).First(&postModel, "slug = ?", slug).Error; err != nil {
		return post.Post{}, fmt.Errorf("finding post by slug: %w", err)
	}
	return postModel.ToDomainPost(), nil
}

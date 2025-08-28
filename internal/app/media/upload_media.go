package media

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/typetrait/lit/internal/app/post"
	domain "github.com/typetrait/lit/internal/domain/post"
)

var (
	ErrInvalidMedia = errors.New("invalid media")
)

type UploadMedia struct {
	storage         Storage
	postRepository  post.Repository
	mediaRepository Repository
	detector        Detector
}

func NewUploadMedia(
	storage Storage,
	postRepository post.Repository,
	mediaRepository Repository,
	detector Detector,
) *UploadMedia {
	return &UploadMedia{
		storage:         storage,
		postRepository:  postRepository,
		mediaRepository: mediaRepository,
		detector:        detector,
	}
}

func (um *UploadMedia) Upload(ctx context.Context, cmd UploadMediaCommand) (domain.Media, error) {
	contentType, err := um.detector.DetectType(cmd.Upload.Reader)
	if err != nil {
		return domain.Media{}, ErrInvalidMedia
	}

	if !um.isValidMediaType(contentType) {
		return domain.Media{}, ErrInvalidMedia
	}

	associatedPost, err := um.postRepository.FindByID(ctx, cmd.Upload.PostID)
	if err != nil {
		return domain.Media{}, fmt.Errorf("finding associated post: %w", err)
	}

	objectKey := uuid.New()
	err = um.storage.Put(ctx, objectKey.String(), cmd.Upload.Reader)
	if err != nil {
		return domain.Media{}, fmt.Errorf("storing media: %w", err)
	}

	var caption *string
	if cmd.Caption != nil {
		if tmp := strings.TrimSpace(*cmd.Caption); tmp != "" {
			caption = &tmp
		}
	}

	writeTime := time.Now().UTC()
	m := domain.Media{
		Post:      associatedPost,
		Key:       objectKey,
		Mime:      contentType,
		Alt:       strings.TrimSpace(cmd.Alt),
		Caption:   caption,
		CreatedAt: writeTime,
		UpdatedAt: writeTime,
	}

	createdMedia, err := um.mediaRepository.Create(ctx, m)
	if err != nil {
		rollbackErr := um.storage.Delete(ctx, objectKey.String())
		if rollbackErr != nil {
			return domain.Media{},
				fmt.Errorf(
					"creating media: %w (cleanup failed: %v)",
					err,
					rollbackErr,
				)
		}
		return domain.Media{}, fmt.Errorf("creating media: %w", err)
	}
	return createdMedia, nil
}

func (um *UploadMedia) isValidMediaType(contentType string) bool {
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	return mediaType == "image/jpeg" ||
		mediaType == "image/jpg" ||
		mediaType == "image/png" ||
		mediaType == "image/gif" ||
		mediaType == "image/webp"
}

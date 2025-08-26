package media

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID        int64
	Key       uuid.UUID
	Mime      string
	Alt       string
	Caption   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

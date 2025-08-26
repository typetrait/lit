package post

type DraftPostResponse struct {
	PostID int64 `json:"post_id"`
}

type PublishPostResponse struct {
	PostID int64  `json:"post_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

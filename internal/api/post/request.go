package post

type PublishPostRequest struct {
	Title         string   `json:"title"`
	ContentFormat string   `json:"content_format"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
}

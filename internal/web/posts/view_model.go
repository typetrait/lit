package posts

import "html/template"

type PostViewModel struct {
	Title       string
	HTMLContent template.HTML
	Slug        string
	DatePosted  string
}

type ViewModel struct {
	Posts []PostViewModel
}

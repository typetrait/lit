package home

import "html/template"

type PostPreviewViewModel struct {
	Title       string
	HTMLContent template.HTML
	Slug        string
}

type ViewModel struct {
	Posts []PostPreviewViewModel
}

package rendering

import (
	"fmt"
	"html/template"

	"github.com/typetrait/lit/internal/domain/post"
)

type ContentRenderer struct {
	markdownRenderer *MarkdownRenderer
}

func NewContentRenderer(markdownRenderer *MarkdownRenderer) *ContentRenderer {
	return &ContentRenderer{
		markdownRenderer: markdownRenderer,
	}
}

func (r *ContentRenderer) ContentToHTML(content post.Content) (template.HTML, error) {
	var rawHTML string
	var err error
	switch content.Format {
	case post.FormatMarkdown:
		rawHTML, err = r.markdownRenderer.Render(content.Source)
		if err != nil {
			return "", fmt.Errorf("rendering content to HTML: %w", err)
		}
	}
	return template.HTML(rawHTML), nil
}

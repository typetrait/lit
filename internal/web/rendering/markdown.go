package rendering

import (
	"bytes"

	"github.com/yuin/goldmark"
)

type MarkdownRenderer struct {
}

func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{}
}

func (m *MarkdownRenderer) Render(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

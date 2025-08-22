package post

type ContentFormat uint8

const (
	FormatMarkdown ContentFormat = iota
)

type Content struct {
	Format ContentFormat
	Source string
}

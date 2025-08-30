package settings

const (
	KeyBlogName     string = "lit__BlogName"
	KeyBlogSubtitle string = "lit__BlogSubtitle"
	KeyBlogAbout    string = "lit__BlogAbout"
)

type Settings struct {
	BlogName           string
	BlogSubtitle       string
	BlogAbout          string
	AdditionalSettings map[string]any
}

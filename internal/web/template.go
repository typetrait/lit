package web

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/app"
	"github.com/typetrait/lit/internal/domain/settings"
)

const (
	ViewDataModel    string = "Model"
	ViewDataSettings string = "Settings"
)

type Template struct {
	sets     map[string]*template.Template
	settings settings.Settings
}

func NewTemplate(settingsProvider app.SettingsProvider) *Template {
	base := template.Must(template.ParseFiles("public/views/base.gohtml"))

	base = template.Must(base.ParseGlob("public/views/partials/*.gohtml"))

	pageFiles, err := filepath.Glob("public/views/pages/*.gohtml")
	if err != nil {
		panic(err)
	}

	sets := make(map[string]*template.Template, len(pageFiles))
	for _, file := range pageFiles {
		name := strings.TrimSuffix(filepath.Base(file), ".gohtml")

		cl := template.Must(base.Clone())
		template.Must(cl.ParseFiles(file))

		sets[name] = cl
	}

	t := &Template{sets: sets}
	t.settings = settingsProvider.Settings()
	return t
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	set, ok := t.sets[name]
	if !ok {
		return echo.NewHTTPError(404, "template not found: "+name)
	}

	viewData := map[string]any{
		ViewDataSettings: t.settings,
	}

	switch v := data.(type) {
	case nil:
	case map[string]any:
		for k, val := range v {
			viewData[k] = val
		}
	default:
		viewData[ViewDataModel] = v
	}

	return set.ExecuteTemplate(w, name, viewData)
}

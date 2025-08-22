package about

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	aboutTemplate string = "about"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, aboutTemplate, map[string]any{})
	}
}

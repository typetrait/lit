package sign_in

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	signInTemplate string = "sign-in"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, signInTemplate, map[string]any{})
	}
}

func (h *Handler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, signInTemplate, map[string]any{})
	}
}

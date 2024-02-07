package handlers

import (
	usecase "simple-bank/internal/usecase/account"

	"github.com/labstack/echo/v4"
)

type ResetHandler struct {
	resetUseCase *usecase.ResetUseCase
}

func NewResetHandler(resetUseCase *usecase.ResetUseCase) *ResetHandler {
	return &ResetHandler{resetUseCase: resetUseCase}
}

func (h *ResetHandler) Reset(c echo.Context) error {
	err := h.resetUseCase.Execute()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.String(200, "OK")
}

func (h *ResetHandler) Setup(e *echo.Echo) {
	e.POST("/reset", h.Reset)
}

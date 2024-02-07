package handlers

import (
	"errors"
	"net/http"
	usecase "simple-bank/internal/usecase/account"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BalanceHandler struct {
	getBalanceUseCase *usecase.GetBalanceUseCase
}

func NewBalanceHandler(getBalanceUseCase *usecase.GetBalanceUseCase) *BalanceHandler {
	return &BalanceHandler{getBalanceUseCase: getBalanceUseCase}
}

func (h *BalanceHandler) GetBalance(c echo.Context) error {
	accountID := c.QueryParam("account_id")

	balance, err := h.getBalanceUseCase.Execute(usecase.GetBalanceInputDTO{ID: accountID})
	if err != nil {
		if errors.Is(err, usecase.ErrGetBalanceAccountNotExists) {
			return c.String(404, "0")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(200, strconv.Itoa(balance.Balance))
}

func (h *BalanceHandler) Setup(e *echo.Echo) {
	e.GET("/balance", h.GetBalance)
}

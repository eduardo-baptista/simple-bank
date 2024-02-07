package handlers

import (
	"errors"
	"net/http"
	"simple-bank/internal/shared/dto"
	usecase "simple-bank/internal/usecase/account"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	depositUseCase  *usecase.DepositUseCase
	withdrawUseCase *usecase.WithdrawUseCase
	transferUseCase *usecase.TransferUseCase
}

type HandleEventRequest struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
	Origin      string `json:"origin"`
}

type HandleEventResponse struct {
	Destination *dto.AccountDTO `json:"destination,omitempty"`
	Origin      *dto.AccountDTO `json:"origin,omitempty"`
}

func NewEventHandler(
	depositUseCase *usecase.DepositUseCase,
	withdrawUseCase *usecase.WithdrawUseCase,
	transferUseCase *usecase.TransferUseCase,
) *EventHandler {
	return &EventHandler{
		depositUseCase:  depositUseCase,
		withdrawUseCase: withdrawUseCase,
		transferUseCase: transferUseCase,
	}
}

func (h *EventHandler) HandleEvent(c echo.Context) error {
	var request HandleEventRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	switch request.Type {
	case "deposit":
		output, err := h.depositUseCase.Execute(usecase.DepositInputDTO{
			Destination: request.Destination,
			Amount:      request.Amount,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, HandleEventResponse{Destination: &output.Destination})
	case "withdraw":
		output, err := h.withdrawUseCase.Execute(usecase.WithdrawInputDTO{
			Origin: request.Origin,
			Amount: request.Amount,
		})
		if err != nil {
			if errors.Is(err, usecase.ErrWithdrawAccountNotExists) {
				return c.String(http.StatusNotFound, "0")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, HandleEventResponse{Origin: &output.Origin})
	case "transfer":
		output, err := h.transferUseCase.Execute(usecase.TransferInputDTO{
			Origin:      request.Origin,
			Destination: request.Destination,
			Amount:      request.Amount,
		})

		if err != nil {
			if errors.Is(err, usecase.ErrTransferOriginAccountNotExists) {
				return c.String(http.StatusNotFound, "0")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, HandleEventResponse{
			Origin:      &output.Origin,
			Destination: &output.Destination,
		})
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid event type")
	}
}

func (h *EventHandler) Setup(e *echo.Echo) {
	e.POST("/event", h.HandleEvent)
}

package support

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	appHttp "simple-bank/internal/infrastructure/http"
	handlers "simple-bank/internal/infrastructure/http/handler"
	"simple-bank/internal/infrastructure/repository/inmemory"
	"simple-bank/internal/usecase/account"

	"github.com/labstack/echo/v4"
)

type TestApp struct {
	AccountRepository *inmemory.AccountRepository
	HTTPServer        *appHttp.HTTPServer
}

func NewTestApp() *TestApp {
	accountRepository := inmemory.NewAccountRepository()

	getBalanceUseCase := account.NewGetBalanceUseCase(accountRepository)
	resetUseCase := account.NewResetUseCase(accountRepository)
	depositUseCase := account.NewDepositUseCase(accountRepository)
	withdrawUseCase := account.NewWithdrawUseCase(accountRepository)
	transferUseCase := account.NewTransferUseCase(accountRepository)

	balanceHandler := handlers.NewBalanceHandler(getBalanceUseCase)
	resetHandler := handlers.NewResetHandler(resetUseCase)
	eventHandler := handlers.NewEventHandler(
		depositUseCase,
		withdrawUseCase,
		transferUseCase,
	)

	httpServer := appHttp.NewHTTPServer(
		"3000",
		balanceHandler,
		resetHandler,
		eventHandler,
	)

	return &TestApp{
		AccountRepository: accountRepository,
		HTTPServer:        httpServer,
	}
}

func (a *TestApp) PerformRequest(res *httptest.ResponseRecorder, req *http.Request) {
	a.HTTPServer.Engine.ServeHTTP(res, req)
}

func (a *TestApp) NewJSONRequest(method, path string, body any) *http.Request {
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, path, bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return req
}

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

	resetUseCase := account.NewResetUseCase(accountRepository)
	resetHandler := handlers.NewResetHandler(resetUseCase)
	httpServer := appHttp.NewHTTPServer(
		"3000",
		resetHandler,
	)

	return &TestApp{
		AccountRepository: accountRepository,
		HTTPServer:        httpServer,
	}
}

func (a *TestApp) PerformRequest(res *httptest.ResponseRecorder, req *http.Request) {
	a.HTTPServer.Engine.ServeHTTP(res, req)
}


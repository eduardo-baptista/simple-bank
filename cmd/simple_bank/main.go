package main

import (
	"flag"
	"simple-bank/internal/infrastructure/http"
	handlers "simple-bank/internal/infrastructure/http/handler"
	"simple-bank/internal/infrastructure/repository/inmemory"
	usecase "simple-bank/internal/usecase/account"
)

func main() {
	port := flag.String("port", "3000", "server port, default is 3000")
	flag.Parse()

	accountRepository := inmemory.NewAccountRepository()

	getBalanceUseCase := usecase.NewGetBalanceUseCase(accountRepository)
	resetUseCase := usecase.NewResetUseCase(accountRepository)
	depositUseCase := usecase.NewDepositUseCase(accountRepository)
	withdrawUseCase := usecase.NewWithdrawUseCase(accountRepository)
	transferUseCase := usecase.NewTransferUseCase(accountRepository)

	balanceHandler := handlers.NewBalanceHandler(getBalanceUseCase)
	resetHandler := handlers.NewResetHandler(resetUseCase)
	eventHandler := handlers.NewEventHandler(
		depositUseCase,
		withdrawUseCase,
		transferUseCase,
	)

	httpServer := http.NewHTTPServer(
		*port,
		balanceHandler,
		resetHandler,
		eventHandler,
	)

	httpServer.Start()
}

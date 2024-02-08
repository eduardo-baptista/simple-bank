package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"simple-bank/internal/infrastructure/http"
	handlers "simple-bank/internal/infrastructure/http/handler"
	"simple-bank/internal/infrastructure/repository/inmemory"
	usecase "simple-bank/internal/usecase/account"
	"syscall"
	"time"
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

	go httpServer.Start()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		panic(err)
	}
}

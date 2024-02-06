package repository

import "simple-bank/internal/domain/entity"

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source=${GOFILE} -destination=mocks/${GOFILE} -package=mocks AccountRepository

type AccountRepository interface {
	GetAccountByID(id string) (*entity.Account, error)
	UpdateAccount(account *entity.Account) error
	SaveAccount(account *entity.Account) error
	DeleteAllAccounts() error
}

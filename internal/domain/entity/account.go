package entity

import (
	domainErrs "simple-bank/internal/domain/errors"
)

type Account struct {
	ID      string
	Balance int
}

func NewAccount(id string, balance int) *Account {
	return &Account{
		ID:      id,
		Balance: balance,
	}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance < amount {
		return domainErrs.ErrAccountInsufficientBalance
	}

	a.Balance -= amount
	return nil
}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	t.Run("Should create new account", func(t *testing.T) {
		account := NewAccount("ID", 100)

		assert.Equal(t, "ID", account.ID)
		assert.Equal(t, 100, account.Balance)
	})
}

func TestAccount_Deposit(t *testing.T) {
	t.Run("Should deposit to account", func(t *testing.T) {
		account := NewAccount("ID", 100)
		account.Deposit(50)

		assert.Equal(t, 150, account.Balance)
	})
}

func TestAccount_Withdraw(t *testing.T) {
	t.Run("Should withdraw from account", func(t *testing.T) {
		account := NewAccount("ID", 100)
		account.Withdraw(50)

		assert.Equal(t, 50, account.Balance)
	})

	t.Run("Should return error when insufficient balance", func(t *testing.T) {
		account := NewAccount("ID", 100)
		err := account.Withdraw(150)

		assert.NotNil(t, err)
		assert.Equal(t, "Account as insufficient balance", err.Error())
	})
}

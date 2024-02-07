package inmemory

import "simple-bank/internal/domain/entity"

type AccountRepository struct {
	Accounts map[string]entity.Account
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		Accounts: make(map[string]entity.Account),
	}
}

func (r *AccountRepository) GetAccountByID(id string) (*entity.Account, error) {
	account, ok := r.Accounts[id]
	if !ok {
		return nil, nil
	}
	return &account, nil
}

func (r *AccountRepository) UpdateAccount(account *entity.Account) error {
	r.Accounts[account.ID] = *account
	return nil
}

func (r *AccountRepository) SaveAccount(account *entity.Account) error {
	r.Accounts[account.ID] = *account
	return nil
}

func (r *AccountRepository) DeleteAllAccounts() error {
	r.Accounts = make(map[string]entity.Account)
	return nil
}

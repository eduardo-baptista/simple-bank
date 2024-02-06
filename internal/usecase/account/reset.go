package usecase

import (
	"errors"
	"simple-bank/internal/domain/repository"
)

var ErrResetFailToDeleteAllAccounts = errors.New("[ResetUseCase] Fail to delete all accounts")

type ResetUseCase struct {
	accountRepository repository.AccountRepository
}

func NewResetUseCase(accountRepository repository.AccountRepository) *ResetUseCase {
	return &ResetUseCase{accountRepository: accountRepository}
}

func (uc *ResetUseCase) Execute() error {
	if err := uc.accountRepository.DeleteAllAccounts(); err != nil {
		return errors.Join(ErrResetFailToDeleteAllAccounts, err)
	}
	return nil
}

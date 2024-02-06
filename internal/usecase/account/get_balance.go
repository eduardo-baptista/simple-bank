package usecase

import (
	"errors"
	"simple-bank/internal/domain/repository"
)

var (
	ErrGetBalanceFailToRetrieveAccount = errors.New("[GetBalanceUseCase] Fail to retrieve account")
	ErrGetBalanceAccountNotExists      = errors.New("[GetBalanceUseCase] Account not exists")
)

type GetBalanceInputDTO struct {
	ID string
}

type GetBalanceOutputDTO struct {
	Balance int
}

type GetBalanceUseCase struct {
	accountRepository repository.AccountRepository
}

func NewGetBalanceUseCase(accountRepository repository.AccountRepository) *GetBalanceUseCase {
	return &GetBalanceUseCase{accountRepository: accountRepository}
}

func (uc *GetBalanceUseCase) Execute(input GetBalanceInputDTO) (*GetBalanceOutputDTO, error) {
	account, err := uc.accountRepository.GetAccountByID(input.ID)
	if err != nil {
		return nil, errors.Join(ErrGetBalanceFailToRetrieveAccount, err)
	}

	if account == nil {
		return nil, ErrGetBalanceAccountNotExists
	}

	return &GetBalanceOutputDTO{Balance: account.Balance}, nil
}

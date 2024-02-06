package usecase

import (
	"errors"
	"simple-bank/internal/domain/entity"
	"simple-bank/internal/domain/repository"
	"simple-bank/internal/shared/dto"
)

var (
	ErrDepositFailToRetrieveAccount = errors.New("[DepositUseCase] Fail to retrieve account")
	ErrDepositFailToUpdateAccount   = errors.New("[DepositUseCase] Fail to update account")
	ErrDepositFailToSaveAccount     = errors.New("[DepositUseCase] Fail to save account")
)

type DepositInputDTO struct {
	Destination string
	Amount      int
}

type DepositOutputDTO struct {
	Destination dto.AccountDTO
}

type DepositUseCase struct {
	accountRepository repository.AccountRepository
}

func NewDepositUseCase(accountRepository repository.AccountRepository) *DepositUseCase {
	return &DepositUseCase{accountRepository: accountRepository}
}

func (uc *DepositUseCase) Execute(input DepositInputDTO) (*DepositOutputDTO, error) {
	account, err := uc.accountRepository.GetAccountByID(input.Destination)
	if err != nil {
		return nil, errors.Join(ErrDepositFailToRetrieveAccount, err)
	}

	if account == nil {
		account = entity.NewAccount(input.Destination, input.Amount)
		if err = uc.accountRepository.SaveAccount(account); err != nil {
			return nil, ErrDepositFailToSaveAccount
		}
		return &DepositOutputDTO{
			Destination: dto.AccountDTO{
				ID:      account.ID,
				Balance: account.Balance,
			},
		}, nil
	}

	account.Deposit(input.Amount)

	if err = uc.accountRepository.UpdateAccount(account); err != nil {
		return nil, ErrDepositFailToUpdateAccount
	}

	return &DepositOutputDTO{
		Destination: dto.AccountDTO{
			ID:      account.ID,
			Balance: account.Balance,
		},
	}, nil
}

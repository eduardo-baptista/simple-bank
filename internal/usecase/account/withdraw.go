package account

import (
	"errors"
	"simple-bank/internal/domain/repository"
	"simple-bank/internal/shared/dto"
)

var (
	ErrWithdrawFailToRetrieveAccount = errors.New("[WithdrawUseCase] Fail to retrieve account")
	ErrWithdrawAccountNotExists      = errors.New("[WithdrawUseCase] Account not exists")
	ErrWithdrawFailToWithdraw        = errors.New("[WithdrawUseCase] Fail to withdraw")
	ErrWithdrawFailToUpdateAccount   = errors.New("[WithdrawUseCase] Fail to update account")
)

type WithdrawInputDTO struct {
	Origin string
	Amount int
}

type WithdrawOutputDTO struct {
	Origin dto.AccountDTO
	Amount int
}

type WithdrawUseCase struct {
	accountRepository repository.AccountRepository
}

func NewWithdrawUseCase(accountRepository repository.AccountRepository) *WithdrawUseCase {
	return &WithdrawUseCase{accountRepository: accountRepository}
}

func (uc *WithdrawUseCase) Execute(input WithdrawInputDTO) (*WithdrawOutputDTO, error) {
	account, err := uc.accountRepository.GetAccountByID(input.Origin)
	if err != nil {
		return nil, errors.Join(ErrWithdrawFailToRetrieveAccount, err)
	}

	if account == nil {
		return nil, ErrWithdrawAccountNotExists
	}

	err = account.Withdraw(input.Amount)
	if err != nil {
		return nil, errors.Join(ErrWithdrawFailToWithdraw, err)
	}

	if err = uc.accountRepository.UpdateAccount(account); err != nil {
		return nil, errors.Join(ErrWithdrawFailToUpdateAccount, err)
	}

	return &WithdrawOutputDTO{
		Origin: dto.AccountDTO{
			ID:      account.ID,
			Balance: account.Balance,
		},
		Amount: input.Amount,
	}, nil
}

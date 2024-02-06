package usecase

import (
	"errors"
	"simple-bank/internal/domain/entity"
	"simple-bank/internal/domain/repository"
	"simple-bank/internal/shared/dto"
)

var (
	ErrTransferFailToRetrieveOriginAccount     = errors.New("[TransferUseCase] fail to retrieve origin account")
	ErrTransferOriginAccountNotExists          = errors.New("[TransferUseCase] origin account not exists")
	ErrTransferFailToWithdrawOriginAccount     = errors.New("[TransferUseCase] fail to withdraw from origin account")
	ErrWithdrawFailDestinationAccountNotExists = errors.New("[TransferUseCase] fail to retrieve destination account")
	ErrTransferFailToUpdateOriginAccount       = errors.New("[TransferUseCase] fail to update origin account")
	ErrTransferFailToCreateDestinationAccount  = errors.New("[TransferUseCase] fail to create destination account")
	ErrTransferFailToRollbackOriginAccount     = errors.New("[TransferUseCase] fail to rollback origin account")
	ErrTransferFailToDepositDestinationAccount = errors.New("[TransferUseCase] fail to deposit destination account")
)

type TransferInputDTO struct {
	Origin      string
	Destination string
	Amount      int
}

type TransferOutputDTO struct {
	Origin      dto.AccountDTO
	Destination dto.AccountDTO
}

type TransferUseCase struct {
	accountRepository repository.AccountRepository
}

func NewTransferUseCase(repo repository.AccountRepository) *TransferUseCase {
	return &TransferUseCase{accountRepository: repo}
}

func (uc *TransferUseCase) Execute(input TransferInputDTO) (*TransferOutputDTO, error) {
	origin, err := uc.accountRepository.GetAccountByID(input.Origin)
	if err != nil {
		return nil, errors.Join(ErrTransferFailToRetrieveOriginAccount, err)
	}

	if origin == nil {
		return nil, ErrTransferOriginAccountNotExists
	}

	destination, err := uc.accountRepository.GetAccountByID(input.Destination)
	if err != nil {
		return nil, errors.Join(ErrWithdrawFailDestinationAccountNotExists, err)
	}

	err = origin.Withdraw(input.Amount)
	if err != nil {
		return nil, errors.Join(ErrTransferFailToWithdrawOriginAccount, err)
	}

	err = uc.accountRepository.UpdateAccount(origin)
	if err != nil {
		return nil, errors.Join(ErrTransferFailToUpdateOriginAccount, err)
	}

	if destination == nil {
		destination, err = uc.createDestinationAccount(input.Destination, input.Amount)
	} else {
		err = uc.depositOnDestinationAccount(destination, input.Amount)
	}

	if err != nil {
		rollbackErr := uc.rollbackOriginAccount(origin, input.Amount)
		return nil, errors.Join(rollbackErr, err)
	}

	return &TransferOutputDTO{
		Origin:      dto.AccountDTO(*origin),
		Destination: dto.AccountDTO(*destination),
	}, nil
}

func (uc *TransferUseCase) createDestinationAccount(
	id string,
	amount int,
) (*entity.Account, error) {
	destination := entity.NewAccount(id, amount)
	err := uc.accountRepository.SaveAccount(destination)

	if err != nil {
		return nil, errors.Join(ErrTransferFailToCreateDestinationAccount, err)
	}
	return destination, nil
}

func (uc *TransferUseCase) depositOnDestinationAccount(
	destination *entity.Account,
	amount int,
) error {
	destination.Deposit(amount)
	if err := uc.accountRepository.UpdateAccount(destination); err != nil {
		return errors.Join(ErrTransferFailToDepositDestinationAccount, err)
	}

	return nil
}

func (uc *TransferUseCase) rollbackOriginAccount(
	origin *entity.Account,
	amount int,
) error {
	origin.Deposit(amount)
	if err := uc.accountRepository.UpdateAccount(origin); err != nil {
		return errors.Join(ErrTransferFailToRollbackOriginAccount, err)
	}

	return nil
}

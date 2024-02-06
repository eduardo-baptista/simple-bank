package usecase

import (
	"errors"
	"simple-bank/internal/domain/entity"
	"simple-bank/internal/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestTransferUseCaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockAccountRepository
	sut  *TransferUseCase
}

func (suite *TestTransferUseCaseSuite) SetupSubTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAccountRepository(suite.ctrl)
	suite.sut = NewTransferUseCase(suite.repo)
}

func (suite *TestTransferUseCaseSuite) TearDownSubTest() {
	suite.ctrl.Finish()
}

func (suite *TestTransferUseCaseSuite) TestTransfer() {
	suite.Run("Should transfer amount from origin to destination", func() {
		origin := entity.NewAccount("ID1", 100)
		destination := entity.NewAccount("ID2", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.EXPECT().GetAccountByID("ID2").Return(destination, nil)
		suite.repo.EXPECT().UpdateAccount(origin).Return(nil)
		suite.repo.EXPECT().UpdateAccount(destination).Return(nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.NoError(err)
		suite.Equal(output.Origin.Balance, 50)
		suite.Equal(output.Destination.Balance, 150)
		suite.Equal(output.Origin.ID, "ID1")
		suite.Equal(output.Destination.ID, "ID2")
	})

	suite.Run("Should return error when fails to retrieve origin account", func() {
		amount := 50

		suite.repo.
			EXPECT().
			GetAccountByID("ID1").
			Return(nil, errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToRetrieveOriginAccount)
		suite.Nil(output)
	})

	suite.Run("Should return error when origin account does not exists", func() {
		amount := 50

		suite.repo.
			EXPECT().
			GetAccountByID("ID1").
			Return(nil, nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferOriginAccountNotExists)
		suite.Nil(output)
	})

	suite.Run("Should return error when fails to retrieve destination account", func() {
		origin := entity.NewAccount("ID1", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.
			EXPECT().
			GetAccountByID("ID2").
			Return(nil, errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrWithdrawFailDestinationAccountNotExists)
		suite.Nil(output)
	})

	suite.Run("Should return error when fails to withdraw from origin account", func() {
		origin := entity.NewAccount("ID1", 100)
		destination := entity.NewAccount("ID2", 100)
		amount := 150

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.EXPECT().GetAccountByID("ID2").Return(destination, nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToWithdrawOriginAccount)
		suite.Nil(output)
	})

	suite.Run("Should create destination account when it does not exists", func() {
		origin := entity.NewAccount("ID1", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.
			EXPECT().
			GetAccountByID("ID2").
			Return(nil, nil)
		suite.repo.
			EXPECT().
			SaveAccount(&entity.Account{ID: "ID2", Balance: amount}).
			Return(nil)
		suite.repo.EXPECT().UpdateAccount(origin).Return(nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.NoError(err)
		suite.Equal(output.Origin.Balance, 50)
		suite.Equal(output.Destination.Balance, 50)
		suite.Equal(output.Origin.ID, "ID1")
		suite.Equal(output.Destination.ID, "ID2")
	})

	suite.Run("Should return error when fails to update origin account", func() {
		origin := entity.NewAccount("ID1", 100)
		destination := entity.NewAccount("ID2", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.EXPECT().GetAccountByID("ID2").Return(destination, nil)
		suite.repo.
			EXPECT().
			UpdateAccount(origin).
			Return(errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToUpdateOriginAccount)
		suite.Nil(output)
	})

	suite.Run("Should rollback origin account when fails to create destination account", func() {
		origin := entity.NewAccount("ID1", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.
			EXPECT().
			GetAccountByID("ID2").
			Return(nil, nil)
		suite.repo.
			EXPECT().
			SaveAccount(&entity.Account{ID: "ID2", Balance: amount}).
			Return(errors.New("[AccountRepository] internal error"))
		suite.repo.
			EXPECT().
			UpdateAccount(origin).
			Return(nil)
		suite.repo.
			EXPECT().
			UpdateAccount(&entity.Account{ID: origin.ID, Balance: 100}).
			Return(nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToCreateDestinationAccount)
		suite.Nil(output)
	})

	suite.Run("Should return error when fails to rollback account after fails to create destination account", func() {
		origin := entity.NewAccount("ID1", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.
			EXPECT().
			GetAccountByID("ID2").
			Return(nil, nil)
		suite.repo.
			EXPECT().
			SaveAccount(&entity.Account{ID: "ID2", Balance: amount}).
			Return(errors.New("[AccountRepository] internal error"))
		suite.repo.
			EXPECT().
			UpdateAccount(origin).
			Return(nil)
		suite.repo.
			EXPECT().
			UpdateAccount(&entity.Account{ID: origin.ID, Balance: 100}).
			Return(errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToCreateDestinationAccount)
		suite.ErrorIs(err, ErrTransferFailToRollbackOriginAccount)
		suite.Nil(output)
	})

	suite.Run("Should rollback origin account when fails to deposit destination account", func() {
		origin := entity.NewAccount("ID1", 100)
		destination := entity.NewAccount("ID2", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.EXPECT().GetAccountByID("ID2").Return(destination, nil)
		suite.repo.EXPECT().UpdateAccount(origin).Return(nil)
		suite.repo.
			EXPECT().
			UpdateAccount(destination).
			Return(errors.New("[AccountRepository] internal error"))
		suite.repo.
			EXPECT().
			UpdateAccount(&entity.Account{ID: origin.ID, Balance: 100}).
			Return(nil)

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToDepositDestinationAccount)
		suite.Nil(output)
	})

	suite.Run("Should return error when fails to rollback account after fails to deposit destination account", func() {
		origin := entity.NewAccount("ID1", 100)
		destination := entity.NewAccount("ID2", 100)
		amount := 50

		suite.repo.EXPECT().GetAccountByID("ID1").Return(origin, nil)
		suite.repo.EXPECT().GetAccountByID("ID2").Return(destination, nil)
		suite.repo.EXPECT().UpdateAccount(origin).Return(nil)
		suite.repo.
			EXPECT().
			UpdateAccount(destination).
			Return(errors.New("[AccountRepository] internal error"))
		suite.repo.
			EXPECT().
			UpdateAccount(&entity.Account{ID: origin.ID, Balance: 100}).
			Return(errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(TransferInputDTO{
			Origin:      "ID1",
			Destination: "ID2",
			Amount:      amount,
		})

		suite.ErrorIs(err, ErrTransferFailToDepositDestinationAccount)
		suite.ErrorIs(err, ErrTransferFailToRollbackOriginAccount)
		suite.Nil(output)
	})
}

func TestTransfer(t *testing.T) {
	suite.Run(t, new(TestTransferUseCaseSuite))
}

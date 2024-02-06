package account

import (
	"errors"
	"simple-bank/internal/domain/entity"
	"simple-bank/internal/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestDepositUseCaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockAccountRepository
	sut  *DepositUseCase
}

func (suite *TestDepositUseCaseSuite) SetupSubTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAccountRepository(suite.ctrl)
	suite.sut = NewDepositUseCase(suite.repo)
}

func (suite *TestDepositUseCaseSuite) TearDownSubTest() {
	suite.ctrl.Finish()
}

func (suite *TestDepositUseCaseSuite) TestDeposit() {
	suite.Run("Should deposit amount to account", func() {
		account := entity.NewAccount("ID", 100)

		suite.repo.EXPECT().GetAccountByID("ID").Return(account, nil)
		suite.repo.EXPECT().UpdateAccount(account).Return(nil)

		output, err := suite.sut.Execute(DepositInputDTO{
			Destination: "ID",
			Amount:      100,
		})

		suite.NoError(err)
		suite.Equal(output.Destination.Balance, 200)
		suite.Equal(output.Destination.ID, "ID")
	})

	suite.Run("Should return error when fails to retrieve account", func() {
		suite.repo.EXPECT().GetAccountByID("ID").Return(nil, errors.New("[AccountRepository] internal error"))

		_, err := suite.sut.Execute(DepositInputDTO{
			Destination: "ID",
			Amount:      100,
		})

		suite.ErrorIs(err, ErrDepositFailToRetrieveAccount)
	})

	suite.Run("Should return error when fails to update account", func() {
		account := entity.NewAccount("ID", 100)
		suite.repo.EXPECT().GetAccountByID("ID").Return(account, nil)
		suite.repo.EXPECT().UpdateAccount(account).Return(errors.New("[AccountRepository] internal error"))

		_, err := suite.sut.Execute(DepositInputDTO{
			Destination: "ID",
			Amount:      100,
		})

		suite.ErrorIs(err, ErrDepositFailToUpdateAccount)
	})

	suite.Run("Should create an account when not exists", func() {
		suite.repo.EXPECT().GetAccountByID("ID").Return(nil, nil)
		suite.repo.EXPECT().
			SaveAccount(&entity.Account{ID: "ID", Balance: 100}).
			Return(nil)

		output, err := suite.sut.Execute(DepositInputDTO{
			Destination: "ID",
			Amount:      100,
		})

		suite.NoError(err)
		suite.Equal(output.Destination.Balance, 100)
		suite.Equal(output.Destination.ID, "ID")
	})

	suite.Run("Should return error when fails to save account", func() {
		suite.repo.EXPECT().GetAccountByID("ID").Return(nil, nil)
		suite.repo.EXPECT().
			SaveAccount(&entity.Account{ID: "ID", Balance: 100}).
			Return(errors.New("[AccountRepository] internal error"))

		_, err := suite.sut.Execute(DepositInputDTO{
			Destination: "ID",
			Amount:      100,
		})

		suite.ErrorIs(err, ErrDepositFailToSaveAccount)
	})
}

func TestDeposit(t *testing.T) {
	suite.Run(t, new(TestDepositUseCaseSuite))
}

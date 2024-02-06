package account

import (
	"errors"
	"simple-bank/internal/domain/entity"
	domainErrs "simple-bank/internal/domain/errors"
	"simple-bank/internal/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestWithdrawUseCaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockAccountRepository
	sut  *WithdrawUseCase
}

func (suite *TestWithdrawUseCaseSuite) SetupSubTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAccountRepository(suite.ctrl)
	suite.sut = NewWithdrawUseCase(suite.repo)
}

func (suite *TestWithdrawUseCaseSuite) TearDownSubTest() {
	suite.ctrl.Finish()
}

func (suite *TestWithdrawUseCaseSuite) TestWithdraw() {
	suite.Run("Should withdraw amount from account", func() {
		account := entity.NewAccount("1", 100)
		suite.repo.EXPECT().GetAccountByID("1").Return(account, nil)
		suite.repo.EXPECT().UpdateAccount(account).Return(nil)

		output, err := suite.sut.Execute(WithdrawInputDTO{
			Origin: "1",
			Amount: 50,
		})

		suite.NoError(err)
		suite.Equal(output.Origin.Balance, 50)
		suite.Equal(output.Origin.ID, "1")
	})

	suite.Run("Should return error when fail to retrieve account", func() {
		suite.repo.EXPECT().GetAccountByID("1").Return(nil, errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(WithdrawInputDTO{
			Origin: "1",
			Amount: 50,
		})

		suite.ErrorIs(err, ErrWithdrawFailToRetrieveAccount)
		suite.Nil(output)
	})

	suite.Run("Should return error when withdraw account without balance", func() {
		account := entity.NewAccount("1", 100)
		suite.repo.EXPECT().GetAccountByID("1").Return(account, nil)

		output, err := suite.sut.Execute(WithdrawInputDTO{
			Origin: "1",
			Amount: 150,
		})

		suite.ErrorIs(err, ErrWithdrawFailToWithdraw)
		suite.ErrorIs(err, domainErrs.ErrAccountInsufficientBalance)
		suite.Nil(output)
	})

	suite.Run("Should return error when fail to update account", func() {
		account := entity.NewAccount("1", 100)
		suite.repo.EXPECT().GetAccountByID("1").Return(account, nil)
		suite.repo.EXPECT().UpdateAccount(account).Return(errors.New("[AccountRepository] internal error"))

		output, err := suite.sut.Execute(WithdrawInputDTO{
			Origin: "1",
			Amount: 50,
		})

		suite.ErrorIs(err, ErrWithdrawFailToUpdateAccount)
		suite.Nil(output)
	})

	suite.Run("Should return error when account not exists", func() {
		suite.repo.EXPECT().GetAccountByID("1").Return(nil, nil)

		output, err := suite.sut.Execute(WithdrawInputDTO{
			Origin: "1",
			Amount: 50,
		})

		suite.ErrorIs(err, ErrWithdrawAccountNotExists)
		suite.Nil(output)
	})
}

func TestWithdraw(t *testing.T) {
	suite.Run(t, new(TestWithdrawUseCaseSuite))
}

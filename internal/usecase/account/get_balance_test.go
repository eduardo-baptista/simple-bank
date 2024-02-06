package account

import (
	"errors"
	"simple-bank/internal/domain/entity"
	"simple-bank/internal/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestGetBalanceUseCaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockAccountRepository
	sut  *GetBalanceUseCase
}

func (suite *TestGetBalanceUseCaseSuite) SetupSubTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAccountRepository(suite.ctrl)
	suite.sut = NewGetBalanceUseCase(suite.repo)
}

func (suite *TestGetBalanceUseCaseSuite) TearDownSubTest() {
	suite.ctrl.Finish()
}

func (suite *TestGetBalanceUseCaseSuite) TestGetBalance() {
	suite.Run("Should return account balance when account exists", func() {
		account := entity.NewAccount("ID", 100)

		suite.repo.EXPECT().GetAccountByID("ID").Return(account, nil)

		output, err := suite.sut.Execute(GetBalanceInputDTO{ID: "ID"})

		suite.NoError(err)
		suite.Equal(account.Balance, output.Balance)
	})

	suite.Run("Should return error when fails to retrieve account", func() {
		suite.repo.EXPECT().GetAccountByID("ID").Return(nil, errors.New("[AccountRepository] internal error"))

		_, err := suite.sut.Execute(GetBalanceInputDTO{ID: "ID"})

		suite.ErrorIs(err, ErrGetBalanceFailToRetrieveAccount)
	})

	suite.Run("Should return error when account does not exist", func() {
		suite.repo.EXPECT().GetAccountByID("ID").Return(nil, nil)

		_, err := suite.sut.Execute(GetBalanceInputDTO{ID: "ID"})

		suite.ErrorIs(err, ErrGetBalanceAccountNotExists)
	})
}

func TestGetBalance(t *testing.T) {
	suite.Run(t, new(TestGetBalanceUseCaseSuite))
}

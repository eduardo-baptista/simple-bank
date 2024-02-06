package account

import (
	"errors"
	"simple-bank/internal/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestResetUseCaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockAccountRepository
	sut  *ResetUseCase
}

func (suite *TestResetUseCaseSuite) SetupSubTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAccountRepository(suite.ctrl)
	suite.sut = NewResetUseCase(suite.repo)
}

func (suite *TestResetUseCaseSuite) TearDownSubTest() {
	suite.ctrl.Finish()
}

func (suite *TestResetUseCaseSuite) TestReset() {
	suite.Run("Should reset all accounts", func() {
		suite.repo.EXPECT().DeleteAllAccounts().Return(nil)

		err := suite.sut.Execute()

		suite.NoError(err)
	})

	suite.Run("Should return error when fails to reset all accounts", func() {
		suite.repo.
			EXPECT().
			DeleteAllAccounts().
			Return(errors.New("[AccountRepository] internal error"))

		err := suite.sut.Execute()

		suite.ErrorIs(err, ErrResetFailToDeleteAllAccounts)
	})
}

func TestReset(t *testing.T) {
	suite.Run(t, new(TestResetUseCaseSuite))
}

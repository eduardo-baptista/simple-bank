package integration

import (
	"net/http"
	"net/http/httptest"
	"simple-bank/internal/domain/entity"
	"simple-bank/test/support"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestBalanceHandlerSuite struct {
	suite.Suite
	app *support.TestApp
}

func (suite *TestBalanceHandlerSuite) SetupSubTest() {
	suite.app = support.NewTestApp()
}

func (suite *TestBalanceHandlerSuite) Test_GET_Balance() {
	suite.Run("Should return balance when account exists", func() {
		account := entity.NewAccount("ID1", 100)
		suite.app.AccountRepository.SaveAccount(account)

		req := httptest.NewRequest(http.MethodGet, "/balance?account_id=ID1", nil)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusOK, rec.Code)
		suite.Equal("100", rec.Body.String())
	})

	suite.Run("Should return 404 when account does not exist", func() {
		req := httptest.NewRequest(http.MethodGet, "/balance?account_id=ID1", nil)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusNotFound, rec.Code)
		suite.Equal("0", rec.Body.String())
	})
}

func TestBalanceHandler(t *testing.T) {
	suite.Run(t, new(TestBalanceHandlerSuite))
}

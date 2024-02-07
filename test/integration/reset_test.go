package integration

import (
	"net/http"
	"net/http/httptest"
	"simple-bank/internal/domain/entity"
	"simple-bank/test/support"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestResetHandlerSuite struct {
	suite.Suite
	app *support.TestApp
}

func (suite *TestResetHandlerSuite) SetupSubTest() {
	suite.app = support.NewTestApp()
}

func (suite *TestResetHandlerSuite) Test_POST_Reset() {
	suite.Run("Should reset the app removing all accounts", func() {
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("ID1", 100))
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("ID2", 200))

		req := httptest.NewRequest(http.MethodPost, "/reset", nil)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(200, rec.Code)
		suite.Equal("OK", rec.Body.String())
		suite.Equal(0, len(suite.app.AccountRepository.Accounts))
	})
}

func TestResetHandler(t *testing.T) {
	suite.Run(t, new(TestResetHandlerSuite))
}

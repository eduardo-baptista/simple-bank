package integration

import (
	"net/http"
	"net/http/httptest"
	"simple-bank/internal/domain/entity"
	"simple-bank/test/support"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestEventHandlerSuite struct {
	suite.Suite
	app *support.TestApp
}

func (suite *TestEventHandlerSuite) SetupSubTest() {
	suite.app = support.NewTestApp()
}

func (suite *TestEventHandlerSuite) Test_POST_Event_Deposit() {
	suite.Run("Should create account when account does not exist", func() {
		body := map[string]interface{}{
			"type":        "deposit",
			"destination": "100",
			"amount":      100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusCreated, rec.Code)
		suite.JSONEq(`{"destination": {"id": "100", "balance": 100}}`, rec.Body.String())
	})

	suite.Run("Should deposit amount when account exists", func() {
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("100", 100))

		body := map[string]interface{}{
			"type":        "deposit",
			"destination": "100",
			"amount":      100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusCreated, rec.Code)
		suite.JSONEq(`{"destination": {"id": "100", "balance": 200}}`, rec.Body.String())
	})
}

func (suite *TestEventHandlerSuite) Test_POST_Event_Withdraw() {
	suite.Run("Should return 404 when account does not exist", func() {
		body := map[string]interface{}{
			"type":   "withdraw",
			"origin": "100",
			"amount": 100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusNotFound, rec.Code)
		suite.Equal("0", rec.Body.String())
	})

	suite.Run("Should withdraw amount when account exists", func() {
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("100", 100))

		body := map[string]interface{}{
			"type":   "withdraw",
			"origin": "100",
			"amount": 100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusCreated, rec.Code)
		suite.JSONEq(`{"origin": {"id": "100", "balance": 0}}`, rec.Body.String())
	})
}

func (suite *TestEventHandlerSuite) Test_POST_Event_Transfer() {
	suite.Run("Should return 404 when origin account does not exist", func() {
		body := map[string]interface{}{
			"type":        "transfer",
			"origin":      "100",
			"destination": "200",
			"amount":      100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusNotFound, rec.Code)
		suite.Equal("0", rec.Body.String())
	})

	suite.Run("Should create destination account when does not exist", func() {
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("100", 100))

		body := map[string]interface{}{
			"type":        "transfer",
			"origin":      "100",
			"destination": "200",
			"amount":      100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusCreated, rec.Code)
		suite.JSONEq(`{"origin": {"id": "100", "balance": 0}, "destination": {"id": "200", "balance": 100}}`, rec.Body.String())
	})

	suite.Run("Should transfer amount when accounts exist", func() {
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("100", 100))
		suite.app.AccountRepository.SaveAccount(entity.NewAccount("200", 100))

		body := map[string]interface{}{
			"type":        "transfer",
			"origin":      "100",
			"destination": "200",
			"amount":      100,
		}
		req := suite.app.NewJSONRequest(http.MethodPost, "/event", body)
		rec := httptest.NewRecorder()

		suite.app.PerformRequest(rec, req)

		suite.Equal(http.StatusCreated, rec.Code)
		suite.JSONEq(`{"origin": {"id": "100", "balance": 0}, "destination": {"id": "200", "balance": 200}}`, rec.Body.String())
	})
}

func TestEventHandler(t *testing.T) {
	suite.Run(t, new(TestEventHandlerSuite))
}

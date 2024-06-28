package test

import (
	"bytes"
	"encoding/json"
	"goproject/internal/handlers"
	"goproject/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) CreateTransaction(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionService) GetTransactionByID(id string) (*models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetAllTransactions() ([]models.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockTransactionService)
	handler := &handlers.TransactionHandler{}

	t.Run("Success", func(t *testing.T) {
		transaction := &models.Transaction{KonsumenID: 1, Amount: 100000}
		mockService.On("CreateTransaction", transaction).Return(nil)

		reqBody, _ := json.Marshal(transaction)
		req, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/transaction", handler.CreateTransaction)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"status": "transaction successful"}`, w.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer([]byte(`{}`)))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/transaction", handler.CreateTransaction)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})
}

func TestGetTransactionByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockTransactionService)
	handler := &handlers.TransactionHandler{}

	t.Run("Success", func(t *testing.T) {
		transaction := &models.Transaction{ID: 1, KonsumenID: 1, Amount: 100000}
		mockService.On("GetTransactionByID", "1").Return(transaction, nil)

		req, err := http.NewRequest(http.MethodGet, "/transaction/1", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/transaction/:id", handler.GetTransactionByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody, _ := json.Marshal(gin.H{"transaction": transaction})
		assert.JSONEq(t, string(expectedBody), w.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		errR := error.Error(nil)
		mockService.On("GetTransactionByID", "2").Return(nil, errR)

		req, err := http.NewRequest(http.MethodGet, "/transaction/2", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/transaction/:id", handler.GetTransactionByID)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}

func TestGetAllTransactions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockTransactionService)
	handler := &handlers.TransactionHandler{}

	t.Run("Success", func(t *testing.T) {
		transactions := []models.Transaction{
			{ID: 1, KonsumenID: 1, Amount: 100000},
			{ID: 2, KonsumenID: 2, Amount: 200000},
		}
		mockService.On("GetAllTransactions").Return(transactions, nil)

		req, err := http.NewRequest(http.MethodGet, "/transactions", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/transactions", handler.GetAllTransactions)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody, _ := json.Marshal(transactions)
		assert.JSONEq(t, string(expectedBody), w.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		errR := error.Error(nil)
		mockService.On("GetAllTransactions").Return(nil, errR)

		req, err := http.NewRequest(http.MethodGet, "/transactions", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/transactions", handler.GetAllTransactions)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}
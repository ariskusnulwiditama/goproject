package handlers

import (
    "net/http"
    "goproject/internal/services"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
	"goproject/internal/models"
)

type TransactionHandler struct {
    transactionService *services.TransactionService
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
    transactionService := services.NewTransactionService(db)
    return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.transactionService.CreateTransaction(&transaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "transaction successful"})
}


func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
    id := c.Param("id")
    transaction, err := h.transactionService.GetTransactionByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
    transactions, err := h.transactionService.GetAllTransactions()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, transactions)
}
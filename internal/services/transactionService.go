package services

import (
    "github.com/jinzhu/gorm"
    "sync"
	"goproject/internal/models"
)



type TransactionService struct {
    db  *gorm.DB
    mux sync.Mutex
}

func NewTransactionService(db *gorm.DB) *TransactionService {
    return &TransactionService{db: db}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
    s.mux.Lock()
    defer s.mux.Unlock()

    if err := s.db.Create(transaction).Error; err != nil {
        return err
    }
    return nil
}

func (s *TransactionService) GetTransaction() ([]models.Transaction, error) {
    s.mux.Lock()
    defer s.mux.Unlock()

    var transactions []models.Transaction
    if err := s.db.Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}

func (s *TransactionService) GetTransactionByID(id string) (*models.Transaction, error) {
    var transaction models.Transaction
    if err := s.db.Preload("Konsumen").First(&transaction, id).Error; err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
    var transactions []models.Transaction
    if err := s.db.Preload("Konsumen").Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}
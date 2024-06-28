package server

import (
    "github.com/gin-gonic/gin"
    "goproject/internal/handlers"
)

type Server struct {
    router *gin.Engine
}

func NewServer() *Server {
    router := gin.Default()
    return &Server{router: router}
}

func (s *Server) RegisterRoutes(transactionHandler *handlers.TransactionHandler, konsumenHandler *handlers.KonsumenHandler) {
    s.router.POST("/transaction", transactionHandler.CreateTransaction)
    s.router.GET("/transaction", transactionHandler.GetAllTransactions)
    s.router.GET("/transaction/:id", transactionHandler.GetTransactionByID)
    s.router.POST("/konsumen", konsumenHandler.CreateKonsumen)
    s.router.GET("/konsumen", konsumenHandler.GetAllKonsumens)
    s.router.GET("/konsumen/:id", konsumenHandler.GetKonsumenByID)
}



func (s *Server) Run() {
    s.router.Run(":8080")
}
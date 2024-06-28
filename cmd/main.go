package main

import (
    "goproject/internal/handlers"
    "goproject/pkg/db"
    "goproject/pkg/server"
    "goproject/internal/models"
)

func main() {
    db := db.SetupDatabase()
    srv := server.NewServer()

    db.AutoMigrate(&models.Konsumen{}, &models.Transaction{})
    transactionHandler := handlers.NewTransactionHandler(db)
    konsumenHandler := handlers.NewKonsumenHandler(db)

     srv.RegisterRoutes(transactionHandler, konsumenHandler)
    srv.Run()
}
package models

type Transaction struct {
    ID        uint `json:"id" gorm:"primary_key"`
    KonsumenID uint `json:"konsumen_id"`
    Amount    int  `json:"amount"`
    Konsumen   Konsumen  `gorm:"foreignkey:KonsumenID"`
}
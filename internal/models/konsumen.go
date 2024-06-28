package models

type Konsumen struct {
    ID          uint   `gorm:"primary_key"`
    NIK         string `gorm:"size:16;not null" json:"nik"`
    FullName    string `gorm:"size:100;not null" json:"full_name"`
    LegalName   string `gorm:"size:100;not null" json:"legal_name"`
    TempatLahir string `gorm:"size:100;not null" json:"tempat_lahir"`
    TanggalLahir string `gorm:"type:date;not null" json:"tanggal_lahir"`
    Gaji        float64 `gorm:"type:decimal(15,2);not null" json:"gaji"`
    FotoKTP     []byte `gorm:"type:blob" json:"foto_ktp"`
    FotoSelfie  []byte `gorm:"type:blob" json:"foto_selfie"`
}

func (Transaction) TableName() string {
    return "transaction"
}
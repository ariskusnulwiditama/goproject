package handlers

import (
	"goproject/internal/services"
	"io/ioutil"
	"strconv"

	"goproject/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type KonsumenHandler struct {
	konsumenService *services.KonsumenService
}

func NewKonsumenHandler(db *gorm.DB) *KonsumenHandler {
	konsumenService := services.NewKonsumenService(db)
	return &KonsumenHandler{konsumenService: konsumenService}
}

func (h *KonsumenHandler) CreateKonsumen(c *gin.Context)  {
	var konsumen models.Konsumen

    konsumen.NIK = c.PostForm("nik")
    konsumen.FullName = c.PostForm("full_name")
    konsumen.LegalName = c.PostForm("legal_name")
    konsumen.TempatLahir = c.PostForm("tempat_lahir")
    konsumen.TanggalLahir = c.PostForm("tanggal_lahir")
    gaji, _ := c.GetPostForm("gaji")
    konsumen.Gaji, _ = strconv.ParseFloat(gaji, 64)

    // Upload KTP file
    ktpFile, err := c.FormFile("foto_ktp")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid KTP file"})
        return
    }
    ktpFileData, err := ktpFile.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read KTP file"})
        return
    }
    defer ktpFileData.Close()
    ktpBytes, err := ioutil.ReadAll(ktpFileData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read KTP file"})
        return
    }
    konsumen.FotoKTP = ktpBytes

    // Upload Selfie file
    selfieFile, err := c.FormFile("foto_selfie")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid selfie file"})
        return
    }
    selfieFileData, err := selfieFile.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read selfie file"})
        return
    }
    defer selfieFileData.Close()
    selfieBytes, err := ioutil.ReadAll(selfieFileData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read selfie file"})
        return
    }
    konsumen.FotoSelfie = selfieBytes

    if err := h.konsumenService.CreateKonsumen(&konsumen); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "konsumen created"})
}

func (h *KonsumenHandler) GetKonsumenByID(c *gin.Context) {
	id := c.Param("id")
	konsumen, err := h.konsumenService.GetKonsumenByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"konsumen": konsumen})
}

func (h *KonsumenHandler) GetAllKonsumens(c *gin.Context) {
    konsumens, err := h.konsumenService.GetAllKonsumens()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, konsumens)
}
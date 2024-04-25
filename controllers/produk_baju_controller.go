package controllers

import (
	"net/http"
	"zakki-store/models"

	"github.com/gin-gonic/gin"
)

func IndexProdukBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi GetAllTokoBaju untuk mengambil semua data TokoBaju dari database
	TokoBaju, err := models.GetAllProdukBaju(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data TokoBaju, kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get TokoBaju data"})
		return
	}

	// Kirim data TokoBaju sebagai response
	c.JSON(http.StatusOK, TokoBaju)
}

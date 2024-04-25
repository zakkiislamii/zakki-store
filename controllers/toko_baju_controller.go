package controllers

import (
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexTokoBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi GetAllTokoBaju untuk mengambil semua data TokoBaju dari database
	TokoBaju, err := models.GetAllTokoBaju(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data TokoBaju, kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get TokoBaju data"})
		return
	}

	// Kirim data TokoBaju sebagai response
	c.JSON(http.StatusOK, TokoBaju)
}

func ShowTokoBaju(c *gin.Context) {}

func CreateTokoBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Mendapatkan data dari request
	var TokoBaju structs.TokoBaju
	err := c.ShouldBindJSON(&TokoBaju)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	err = models.InsertTokoBaju(db, TokoBaju)
	if err != nil {
		panic(err)
	}

	// Kirim respons ke klien
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdateTokoBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB
	var TokoBaju structs.TokoBaju
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&TokoBaju)
	if err != nil {
		panic(err)
	}
	TokoBaju.IdTokoBaju = int(id)
	err = models.UpdateTokoBaju(db, TokoBaju)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update TokoBaju",
	})
}
func DeleteTokoBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID TokoBaju dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TokoBaju ID"})
		return
	}

	// Hapus TokoBaju dari database berdasarkan ID
	err = models.DeleteTokoBaju(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete TokoBaju"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete TokoBaju"})
}

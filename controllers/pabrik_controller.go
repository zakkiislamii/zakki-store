package controllers

import (
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexPabrik(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi GetAllPabrik untuk mengambil semua data Pabrik dari database
	Pabrik, err := models.GetAllPabrik(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data Pabrik, kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Pabrik data"})
		return
	}

	// Kirim data Pabrik sebagai response
	c.JSON(http.StatusOK, Pabrik)
}

func ShowPabrik(c *gin.Context) {}

func CreatePabrik(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Mendapatkan data dari request
	var Pabrik structs.Pabrik
	err := c.ShouldBindJSON(&Pabrik)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	err = models.InsertPabrik(db, Pabrik)
	if err != nil {
		panic(err)
	}

	// Kirim respons ke klien
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatePabrik(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB
	var Pabrik structs.Pabrik
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&Pabrik)
	if err != nil {
		panic(err)
	}
	Pabrik.IdPabrik = int(id)
	err = models.UpdatePabrik(db, Pabrik)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Pabrik",
	})
}
func DeletePabrik(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID Pabrik dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pabrik ID"})
		return
	}

	// Hapus Pabrik dari database berdasarkan ID
	err = models.DeletePabrik(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Pabrik"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete Pabrik"})
}

package controllers

import (
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexProdukBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi GetAllProdukBaju untuk mengambil semua data ProdukBaju dari database
	ProdukBaju, err := models.GetAllProdukBaju(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data ProdukBaju, kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ProdukBaju data"})
		return
	}

	// Kirim data ProdukBaju sebagai response
	c.JSON(http.StatusOK, ProdukBaju)
}

func CreateProdukBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Mendapatkan data dari request
	var ProdukBaju structs.ProdukBaju
	err := c.ShouldBindJSON(&ProdukBaju)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	err = models.InsertProdukBaju(db, ProdukBaju)
	if err != nil {
		panic(err)
	}

	// Kirim respons ke klien
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdateProdukBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB
	var ProdukBaju structs.ProdukBaju
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&ProdukBaju)
	if err != nil {
		panic(err)
	}
	ProdukBaju.IdProduk = int(id)
	err = models.UpdateProdukBaju(db, ProdukBaju)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update ProdukBaju",
	})
}
func DeleteProdukBaju(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID ProdukBaju dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ProdukBaju ID"})
		return
	}

	// Hapus ProdukBaju dari database berdasarkan ID
	err = models.DeleteProdukBaju(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ProdukBaju"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete ProdukBaju"})
}

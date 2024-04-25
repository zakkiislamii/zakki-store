package controllers

import (
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexPelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi GetAllPelanggan untuk mengambil semua data pelanggan dari database
	pelanggan, err := models.GetAllPelanggan(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data pelanggan, kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pelanggan data"})
		return
	}

	// Kirim data pelanggan sebagai response
	c.JSON(http.StatusOK, pelanggan)
}

func ShowPelanggan(c *gin.Context) {}

func CreatePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Mendapatkan data dari request
	var pelanggan structs.Pelanggan
	err := c.ShouldBindJSON(&pelanggan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	err = models.InsertPelanggan(db, pelanggan)
	if err != nil {
		panic(err)
	}

	// Kirim respons ke klien
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB
	var pelanggan structs.Pelanggan
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&pelanggan)
	if err != nil {
		panic(err)
	}
	pelanggan.IdPelanggan = int(id)
	err = models.UpdatePelanggan(db, pelanggan)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Pelanggan",
	})
}
func DeletePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID pelanggan dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pelanggan ID"})
		return
	}

	// Hapus pelanggan dari database berdasarkan ID
	err = models.DeletePelanggan(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pelanggan"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Success Delete Pelanggan"})
}

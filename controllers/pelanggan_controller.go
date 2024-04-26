package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"zakki-store/helper"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexPelanggan(c *gin.Context) {
	db := models.DB

	pelanggan, err := models.GetAllPelanggan(db)
	if err != nil {
		log.Println("Error getting pelanggan data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pelanggan data"})
		return
	}

	c.JSON(http.StatusOK, pelanggan)
}

func CreatePelanggan(c *gin.Context) {
	db := models.DB

	var pelanggan structs.Pelanggan
	if err := c.ShouldBindJSON(&pelanggan); err != nil {
		log.Println("Invalid JSON format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.InsertPelanggan(db, pelanggan); err != nil {
		log.Println("Failed to create pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pelanggan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan created successfully"})
}

func UpdatePelanggan(c *gin.Context) {
	db := models.DB

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid pelanggan ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pelanggan ID"})
		return
	}

	var pelanggan structs.Pelanggan
	if err := c.ShouldBindJSON(&pelanggan); err != nil {
		log.Println("Invalid JSON format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	pelanggan.IdPelanggan = id

	if err := models.UpdatePelanggan(db, pelanggan); err != nil {
		log.Println("Failed to update pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pelanggan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan updated successfully"})
}

func DeletePelanggan(c *gin.Context) {
	db := models.DB

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid pelanggan ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pelanggan ID"})
		return
	}

	if err := models.DeletePelanggan(db, id); err != nil {
		log.Println("Failed to delete pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pelanggan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan deleted successfully"})
}

func ViewUlasan(c *gin.Context) {
	db := models.DB

	ulasan, err := models.ViewUlasan(db)
	if err != nil {
		log.Println("Failed to get ulasan data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ulasan data"})
		return
	}

	c.JSON(http.StatusOK, ulasan)
}

func BeliProduk(c *gin.Context) {
	db := models.DB

	var request structs.PembelianRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("Failed to bind request data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request data"})
		return
	}

	namaProduk := request.NamaProduk
	jumlahBarang := request.JumlahBarang
	namaPelanggan := request.NamaPelanggan

	idPelanggan, err := models.GetIDPelangganByNamaPelanggan(db, namaPelanggan)
	if err != nil {
		log.Println("Failed to get ID pelanggan:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get ID pelanggan"})
		return
	}

	err = models.BeliProduk(db, namaProduk, jumlahBarang, idPelanggan)
	if err != nil {
		log.Println("Failed to buy product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hargaProduk, err := models.GetHargaProdukByNamaProduk(db, namaProduk)
	if err != nil {
		log.Println("Failed to get product price:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalHarga := jumlahBarang * hargaProduk

	transaksi := structs.TransaksiResponse{
		TanggalTransaksi: time.Now(),
		JumlahBarang:     jumlahBarang,
		TotalHarga:       totalHarga,
		NamaProduk:       namaProduk,
		NamaPelanggan:    namaPelanggan,
	}

	c.JSON(http.StatusOK, transaksi)
}

func ViewTransaksi(c *gin.Context) {
	db := models.DB

	transaksis, err := models.ViewTransaksi(db)
	if err != nil {
		log.Println("Failed to get transaction data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction data"})
		return
	}

	c.JSON(http.StatusOK, transaksis)
}

func GetRiwayatTransaksiByUser(c *gin.Context) {
	db := models.DB

	username := c.Param("username")

	riwayat, err := models.GetRiwayatTransaksiByUsername(db, username)
	if err != nil {
		log.Println("Failed to get transaction history:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction history"})
		return
	}

	c.JSON(http.StatusOK, riwayat)
}

func GetProdukInfo(c *gin.Context) {
	db := models.DB

	produkInfo, err := models.GetProdukInfo(db)
	if err != nil {
		log.Println("Failed to get product information:", err)
		helper.ResponseJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to get product information"})
		return
	}

	helper.ResponseJSON(c, http.StatusOK, produkInfo)
}

func BeriUlasan(c *gin.Context) {
	db := models.DB

	var ulasan structs.PelangganUlasan
	if err := c.BindJSON(&ulasan); err != nil {
		log.Println("Invalid request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := models.BeriUlasan(db, ulasan)
	if err != nil {
		log.Println("Failed to add ulasan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add ulasan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ulasan added successfully"})
}

func GetUlasanPelangganByUsername(c *gin.Context) {
	db := models.DB

	username := c.Param("username")

	ulasanPelanggan, err := models.GetUlasanPelangganByUsername(db, username)
	if err != nil {
		log.Println("Failed to get ulasan pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ulasan pelanggan"})
		return
	}

	c.JSON(http.StatusOK, ulasanPelanggan)
}

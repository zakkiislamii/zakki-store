package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
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
		log.Println("Error getting pelanggan data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pelanggan data"})
		return
	}

	// Kirim data pelanggan sebagai response
	c.JSON(http.StatusOK, pelanggan)
}

func CreatePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Mendapatkan data dari request
	var pelanggan structs.Pelanggan
	if err := c.ShouldBindJSON(&pelanggan); err != nil {
		log.Println("Invalid JSON format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Insert pelanggan ke dalam database
	err := models.InsertPelanggan(db, pelanggan)
	if err != nil {
		log.Println("Failed to create pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pelanggan"})
		return
	}

	// Kirim respons ke klien
	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan created successfully"})
}

func UpdatePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID pelanggan dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid pelanggan ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pelanggan ID"})
		return
	}

	// Mendapatkan data dari request
	var pelanggan structs.Pelanggan
	if err := c.ShouldBindJSON(&pelanggan); err != nil {
		log.Println("Invalid JSON format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	pelanggan.IdPelanggan = id

	// Update pelanggan di dalam database
	if err := models.UpdatePelanggan(db, pelanggan); err != nil {
		log.Println("Failed to update pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pelanggan"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan updated successfully"})
}

func DeletePelanggan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Dapatkan ID pelanggan dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid pelanggan ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pelanggan ID"})
		return
	}

	// Hapus pelanggan dari database berdasarkan ID
	if err := models.DeletePelanggan(db, id); err != nil {
		log.Println("Failed to delete pelanggan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pelanggan"})
		return
	}

	// Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"result": "Pelanggan deleted successfully"})
}

func ViewUlasan(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi ViewUlasan untuk mengambil semua data ulasan dari database
	ulasan, err := models.ViewUlasan(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data ulasan, kembalikan respons error
		log.Println("Failed to get ulasan data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ulasan data"})
		return
	}

	// Kirim data ulasan sebagai response
	c.JSON(http.StatusOK, ulasan)
}

func BeliProduk(c *gin.Context) {
	db := models.DB

	// Buat variabel untuk menyimpan data permintaan
	var request structs.PembelianRequest

	// Gunakan c.Bind untuk mengikat data permintaan ke struct PembelianRequest
	if err := c.Bind(&request); err != nil {
		log.Println("Failed to bind request data:", err)
		c.JSON(400, gin.H{"error": "Failed to bind request data"})
		return
	}

	namaProduk := request.NamaProduk
	jumlahBarang := request.JumlahBarang
	namaPelanggan := request.NamaPelanggan

	// Mendapatkan ID pelanggan berdasarkan nama pelanggan
	idPelanggan, err := models.GetIDPelangganByNamaPelanggan(db, namaPelanggan)
	if err != nil {
		log.Println("Failed to get ID pelanggan:", err)
		c.JSON(400, gin.H{"error": "Failed to get ID pelanggan"})
		return
	}

	// Memanggil fungsi untuk melakukan pembelian produk
	err = models.BeliProduk(db, namaProduk, jumlahBarang, idPelanggan)
	if err != nil {
		log.Println("Failed to buy product:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan harga produk berdasarkan nama produk
	hargaProduk, err := models.GetHargaProdukByNamaProduk(db, namaProduk)
	if err != nil {
		log.Println("Failed to get product price:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	totalHarga := jumlahBarang * hargaProduk

	// Membuat objek transaksi dengan informasi yang diperoleh
	transaksi := structs.TransaksiResponse{
		TanggalTransaksi: time.Now(),
		JumlahBarang:     jumlahBarang,
		TotalHarga:       totalHarga,
		NamaProduk:       namaProduk,
		NamaPelanggan:    namaPelanggan,
	}

	// Mengirim respons JSON dengan informasi transaksi
	c.JSON(200, transaksi)
}

func ViewTransaksi(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	// Panggil fungsi ViewTransaksi untuk mendapatkan data transaksi dari database
	transaksis, err := models.ViewTransaksi(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data transaksi, kirim respons error
		log.Println("Failed to get transaction data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction data"})
		return
	}

	// Kirim data transaksi sebagai respons JSON
	c.JSON(http.StatusOK, transaksis)
}

func GetRiwayatTransaksiByUser(c *gin.Context) {
	db := models.DB

	username := c.Param("username")

	riwayat, err := models.GetRiwayatTransaksiByUsername(db, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transaction history"})
		return
	}

	c.JSON(http.StatusOK, riwayat)
}

func GetProdukInfo(c *gin.Context) {
	db := models.DB

	produkInfo, err := models.GetProdukInfo(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product information"})
		return
	}

	c.JSON(http.StatusOK, produkInfo)
}

func BeriUlasan(c *gin.Context) {
	db := models.DB

	var ulasan structs.PelangganUlasan
	if err := c.BindJSON(&ulasan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := models.BeriUlasan(db, ulasan)
	if err != nil {
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
		log.Printf("Failed to get ulasan pelanggan: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ulasan pelanggan"})
		return
	}

	c.JSON(http.StatusOK, ulasanPelanggan)
}

package routers

import (
	"zakki-store/controllers"
	"zakki-store/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/logout", controllers.Logout)
	r.GET("/pelanggan/semua-data", controllers.IndexPelanggan)

	// Rute-rute yang memerlukan autentikasi sebagai admin
	adminRoutes := r.Group("/")
	adminRoutes.Use(middlewares.JWTMiddlewareAdmin())
	{
		adminRoutes.GET("/pelanggan/semua-data", controllers.IndexPelanggan)
		adminRoutes.POST("/pelanggan/tambah", controllers.CreatePelanggan)
		adminRoutes.DELETE("/pelanggan/:id", controllers.DeletePelanggan)
		adminRoutes.PUT("/pelanggan/:id", controllers.UpdatePelanggan)
		adminRoutes.GET("/toko-baju/semua-data", controllers.IndexTokoBaju)
		adminRoutes.POST("/toko-baju/tambah", controllers.CreateTokoBaju)
		adminRoutes.DELETE("/toko-baju/:id", controllers.DeleteTokoBaju)
		adminRoutes.PUT("/toko-baju/:id", controllers.UpdateTokoBaju)
		adminRoutes.GET("/pabrik/semua-data", controllers.IndexPabrik)
		adminRoutes.POST("/pabrik/tambah", controllers.CreatePabrik)
		adminRoutes.DELETE("/pabrik/:id", controllers.DeletePabrik)
		adminRoutes.PUT("/pabrik/:id", controllers.UpdatePabrik)
		adminRoutes.GET("/produk-baju/semua-data", controllers.IndexProdukBaju)
		adminRoutes.POST("/produk-baju/tambah", controllers.CreateProdukBaju)
		adminRoutes.DELETE("/produk-baju/:id", controllers.DeleteProdukBaju)
		adminRoutes.PUT("/produk-baju/:id", controllers.UpdateProdukBaju)
		adminRoutes.GET("/transaksi", controllers.ViewTransaksi)
	}

	// Rute-rute yang memerlukan autentikasi sebagai pelanggan
	PelangganRoutes := r.Group("/")
	PelangganRoutes.Use(middlewares.JWTMiddleware())
	{
		PelangganRoutes.GET("/ulasan-semua-produk", controllers.ViewUlasan)
		PelangganRoutes.POST("/beli-produk", controllers.BeliProduk)
		PelangganRoutes.GET("/lihat-riwayat-anda/:username", controllers.GetRiwayatTransaksiByUser)
		PelangganRoutes.GET("/produk-info", controllers.GetProdukInfo)
		PelangganRoutes.POST("/beri-ulasan", controllers.BeriUlasan)
		PelangganRoutes.GET("/riwayat-ulasan-anda/:username", controllers.GetUlasanPelangganByUsername)
	}
	return r
}

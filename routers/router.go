package routers

import (
	"zakki-store/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	//role user
	r.POST("/register", controllers.Register)
	//role admin dan user
	r.POST("/login", controllers.Login)

	//role admin
	r.GET("/pelanggan/semua-data", controllers.IndexPelanggan)
	r.GET("/pelanggan/berdasarkan-id/:id", controllers.ShowPelanggan)
	r.POST("/pelanggan/tambah", controllers.CreatePelanggan)
	r.DELETE("/pelanggan/:id", controllers.DeletePelanggan)
	//role user
	r.PUT("/pelanggan/:id", controllers.UpdatePelanggan)

	//role admin
	r.GET("/toko-baju/semua-data", controllers.IndexTokoBaju)
	r.GET("/toko-baju/berdasarkan-id/:id", controllers.ShowTokoBaju)
	r.POST("/toko-baju/tambah", controllers.CreateTokoBaju)
	r.DELETE("/toko-baju/:id", controllers.DeleteTokoBaju)
	//role user
	r.PUT("/toko-baju/:id", controllers.UpdateTokoBaju)

	//role admin
	r.GET("/pabrik/semua-data", controllers.IndexPabrik)
	r.GET("/pabrik/berdasarkan-id/:id", controllers.ShowPabrik)
	r.POST("/pabrik/tambah", controllers.CreatePabrik)
	r.DELETE("/pabrik/:id", controllers.DeletePabrik)
	//role user
	r.PUT("/pabrik/:id", controllers.UpdatePabrik)

	//role admin
	r.GET("/produk-baju/semua-data", controllers.IndexProdukBaju)
	//role user
	return r
}

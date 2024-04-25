package controllers

import (
	"net/http"

	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB
	var user structs.Pelanggan

	// Bind JSON request ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi input
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Eksekusi query untuk create data pelanggan tanpa enkripsi kata sandi
	err := models.InsertPelanggan(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pelanggan"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

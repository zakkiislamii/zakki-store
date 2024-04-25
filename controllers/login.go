package controllers

import (
	"net/http"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Ambil koneksi database dari models.DB
	db := models.DB

	var user structs.Pelanggan
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validasi input
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Periksa apakah pengguna ada dalam database
	var storedPassword string
	err := db.QueryRow("SELECT password FROM pelanggan WHERE username = $1", user.Username).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

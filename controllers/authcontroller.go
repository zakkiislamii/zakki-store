package controllers

import (
	"log"
	"net/http"
	"time"
	"zakki-store/config"
	"zakki-store/helper"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *gin.Context) {
	var userInput structs.Pelanggan
	if err := c.BindJSON(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(c, http.StatusBadRequest, response)
		return
	}

	// Cari pengguna berdasarkan username
	err := models.LoginUser(userInput.Username, userInput.Password)
	if err != nil {
		// Jika username dan password tidak ditemukan dalam database,
		// cek apakah pengguna adalah admin
		if userInput.Username == "admin" && userInput.Password == "admin123" {
			// Jika pengguna adalah admin, buat token JWT
			expTime := time.Now().Add(time.Minute * 1)
			claims := &config.JWTAdmin{
				Username: userInput.Username,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "go-jwt-mux",
					ExpiresAt: jwt.NewNumericDate(expTime),
				},
			}

			tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			token, err := tokenAlgo.SignedString(config.JWT_KEY)
			if err != nil {
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(c, http.StatusInternalServerError, response)
				return
			}

			// Set cookie token
			c.SetCookie("token", token, 600, "/", "", false, true)

			log.Printf("Admin logged in: %s", userInput.Username)
			response := map[string]string{"message": "Login berhasil sebagai admin"}
			helper.ResponseJSON(c, http.StatusOK, response)
			return
		}

		// Jika bukan admin dan tidak ditemukan dalam database, kirim pesan error
		log.Printf("Failed login attempt: %s", err.Error())
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(c, http.StatusUnauthorized, response)
		return
	}

	// Jika ditemukan dalam database, buat token JWT seperti sebelumnya
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: userInput.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		log.Printf("Error creating JWT token: %s", err.Error())
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(c, http.StatusInternalServerError, response)
		return
	}

	c.SetCookie("token", token, 600, "/", "", false, true)

	log.Printf("User logged in: %s", userInput.Username)
	response := map[string]string{"message": "Login berhasil"}
	helper.ResponseJSON(c, http.StatusOK, response)
}

func Register(c *gin.Context) {
	var userInput structs.Pelanggan
	if err := c.BindJSON(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(c, http.StatusBadRequest, response)
		return
	}

	if err := models.InsertPelanggan(models.DB, userInput); err != nil {
		log.Printf("Failed to register user: %s", err.Error())
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(c, http.StatusInternalServerError, response)
		return
	}

	log.Printf("New user registered: %s", userInput.Username)
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(c, http.StatusOK, response)
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(c, http.StatusOK, response)
}

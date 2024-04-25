package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {
	// Membaca file environment
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Tidak sukses membaca file environment:", err)
		return // Keluar dari fungsi jika gagal membaca file environment
	}
	fmt.Println("Sukses membaca file environment")

	// Mendapatkan nilai-nilai koneksi dari environment
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Membuat string koneksi
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Membuka koneksi ke database
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Gagal membuka koneksi:", err)
		return // Keluar dari fungsi jika gagal membuka koneksi
	}

	// Memeriksa koneksi ke database
	err = DB.Ping()
	if err != nil {
		fmt.Println("Gagal melakukan ping ke database:", err)
		return // Keluar dari fungsi jika ping gagal
	}
	fmt.Println("Koneksi ke database berhasil")
}

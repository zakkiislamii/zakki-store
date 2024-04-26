package models

import (
	"database/sql"
	"errors"
	"log"
	"zakki-store/structs"
)

func GetAllPelanggan(db *sql.DB) ([]structs.Pelanggan, error) {
	query := "SELECT * FROM Pelanggan"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var pelanggan []structs.Pelanggan
	for rows.Next() {
		var p structs.Pelanggan
		if err := rows.Scan(&p.IdPelanggan, &p.NamaPelanggan, &p.NoHp, &p.Username, &p.Password); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		pelanggan = append(pelanggan, p)
	}
	log.Println("Successfully retrieved all pelanggan data")
	return pelanggan, nil
}

func ViewUlasan(db *sql.DB) ([]structs.UlasanPelanggan, error) {
	query := `
		SELECT pb.nama_produk, pl.nama_pelanggan, u.rating, u.ulasan 
		FROM Ulasan u 
		JOIN Produk_Baju pb ON u.id_produk = pb.id_produk 
		JOIN Pelanggan pl ON u.id_pelanggan = pl.id_pelanggan
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var ulasan []structs.UlasanPelanggan
	for rows.Next() {
		var dataUlasan structs.UlasanPelanggan
		if err := rows.Scan(&dataUlasan.NamaProduk, &dataUlasan.NamaPelanggan, &dataUlasan.Rating, &dataUlasan.Ulasan); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		ulasan = append(ulasan, dataUlasan)
	}
	log.Println("Successfully retrieved ulasan data")
	return ulasan, nil
}

func InsertPelanggan(db *sql.DB, pelanggan structs.Pelanggan) error {
	query := "INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, pelanggan.NamaPelanggan, pelanggan.NoHp, pelanggan.Username, pelanggan.Password)
	if err != nil {
		log.Printf("Error inserting pelanggan: %v\n", err)
		return err
	}
	log.Println("Pelanggan inserted successfully")
	return nil
}

func UpdatePelanggan(db *sql.DB, pelanggan structs.Pelanggan) error {
	query := "UPDATE Pelanggan SET nama_pelanggan = $1, no_hp = $2, username = $3, password = $4  WHERE id_pelanggan = $5"
	_, err := db.Exec(query, pelanggan.NamaPelanggan, pelanggan.NoHp, pelanggan.Username, pelanggan.Password, pelanggan.IdPelanggan)
	if err != nil {
		log.Printf("Error updating pelanggan: %v\n", err)
		return err
	}
	log.Println("Pelanggan updated successfully")
	return nil
}

func DeletePelanggan(db *sql.DB, id int) error {
	query := "DELETE FROM Pelanggan WHERE id_pelanggan = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting pelanggan: %v\n", err)
		return err
	}
	log.Println("Pelanggan deleted successfully")
	return nil
}

func BeriUlasan(db *sql.DB, ulasan structs.PelangganUlasan) error {
	query := `
		INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan)
		SELECT
			$1,  -- Ulasan
			$2,  -- Rating
			(SELECT id_produk FROM Produk_Baju WHERE nama_produk = $3), -- ID Produk
			(SELECT id_pelanggan FROM Pelanggan WHERE username = $4)     -- ID Pelanggan
		WHERE EXISTS (
			SELECT 1 FROM Produk_Baju WHERE nama_produk = $3
		) AND EXISTS (
			SELECT 1 FROM Pelanggan WHERE username = $4
		)
	`
	_, err := db.Exec(query, ulasan.Ulasan, ulasan.Rating, ulasan.NamaProduk, ulasan.Username)
	if err != nil {
		log.Printf("Error inserting ulasan: %v\n", err)
		return err
	}
	log.Println("Ulasan added successfully")
	return nil
}

func GetUlasanPelangganByUsername(db *sql.DB, username string) ([]structs.RiwayatUlasan, error) {
	query := `
		SELECT u.ulasan, u.rating, pb.nama_produk, tb.nama_toko
		FROM Ulasan u
		JOIN Produk_Baju pb ON u.id_produk = pb.id_produk
		JOIN Toko_Baju tb ON pb.id_toko = tb.id_toko
		JOIN Pelanggan p ON u.id_pelanggan = p.id_pelanggan
		WHERE p.username = $1
	`

	rows, err := db.Query(query, username)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var ulasanPelanggan []structs.RiwayatUlasan

	for rows.Next() {
		var ulasan structs.RiwayatUlasan
		err := rows.Scan(&ulasan.Ulasan, &ulasan.Rating, &ulasan.NamaProduk, &ulasan.NamaToko)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		ulasanPelanggan = append(ulasanPelanggan, ulasan)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return ulasanPelanggan, nil
}

func LoginUser(username, password string) error {
	var user structs.Pelanggan
	row := DB.QueryRow("SELECT * FROM Pelanggan WHERE username = $1", username)
	err := row.Scan(&user.IdPelanggan, &user.NamaPelanggan, &user.NoHp, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("username atau password salah")
		}
		return err
	}

	if user.Password != password {
		return errors.New("username atau password salah")
	}

	return nil
}

package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetAllPelanggan(db *sql.DB) ([]structs.Pelanggan, error) {
	sql := "SELECT * FROM Pelanggan"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []structs.Pelanggan
	for rows.Next() {
		var pelanggan structs.Pelanggan
		if err := rows.Scan(&pelanggan.IdPelanggan, &pelanggan.NamaPelanggan, &pelanggan.NoHp, &pelanggan.Username, &pelanggan.Password); err != nil {
			return nil, err
		}
		p = append(p, pelanggan)
	}
	return p, nil
}

func InsertPelanggan(db *sql.DB, pelanggan structs.Pelanggan) error {
	sql := "INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sql, pelanggan.NamaPelanggan, pelanggan.NoHp, pelanggan.Username, pelanggan.Password)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePelanggan(db *sql.DB, pelanggan structs.Pelanggan) error {
	sql := "UPDATE Pelanggan SET nama_pelanggan = $1, no_hp = $2, username = $3, password = $4  WHERE id_pelanggan = $5"
	_, err := db.Exec(sql, pelanggan.NamaPelanggan, pelanggan.NoHp, pelanggan.Username, pelanggan.Password, pelanggan.IdPelanggan)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting pelanggan:", err)
		return err
	}
	log.Println("Pelanggan deleted successfully")
	return nil
}

func DeletePelanggan(db *sql.DB, id int) error {
	sql := "DELETE FROM Pelanggan WHERE id_pelanggan = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting pelanggan:", err)
		return err
	}
	log.Println("Pelanggan deleted successfully")
	return nil
}

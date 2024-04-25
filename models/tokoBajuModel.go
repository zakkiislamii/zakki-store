package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetAllTokoBaju(db *sql.DB) ([]structs.TokoBaju, error) {
	sql := "SELECT * FROM Toko_Baju"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []structs.TokoBaju
	for rows.Next() {
		var TokoBaju structs.TokoBaju
		if err := rows.Scan(&TokoBaju.IdTokoBaju, &TokoBaju.NamaTokoBaju); err != nil {
			return nil, err
		}
		p = append(p, TokoBaju)
	}
	return p, nil
}

func InsertTokoBaju(db *sql.DB, TokoBaju structs.TokoBaju) error {
	sql := "INSERT INTO Toko_Baju (nama_toko) VALUES ($1)"
	_, err := db.Exec(sql, TokoBaju.NamaTokoBaju)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTokoBaju(db *sql.DB, TokoBaju structs.TokoBaju) error {
	sql := "UPDATE Toko_Baju SET nama_toko = $1 WHERE id_TokoBaju = $2"
	_, err := db.Exec(sql, TokoBaju.NamaTokoBaju, TokoBaju.IdTokoBaju)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting TokoBaju:", err)
		return err
	}
	log.Println("TokoBaju deleted successfully")
	return nil
}

func DeleteTokoBaju(db *sql.DB, id int) error {
	sql := "DELETE FROM TokoBaju WHERE id_TokoBaju = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting TokoBaju:", err)
		return err
	}
	log.Println("TokoBaju deleted successfully")
	return nil
}

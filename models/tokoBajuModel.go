package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetAllTokoBaju(db *sql.DB) ([]structs.TokoBaju, error) {
	sqlQuery := "SELECT * FROM Toko_Baju"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Println("Error retrieving TokoBaju data:", err)
		return nil, err
	}
	defer rows.Close()

	var tokoBajus []structs.TokoBaju
	for rows.Next() {
		var tokoBaju structs.TokoBaju
		if err := rows.Scan(&tokoBaju.IdTokoBaju, &tokoBaju.NamaTokoBaju); err != nil {
			log.Println("Error scanning TokoBaju row:", err)
			return nil, err
		}
		tokoBajus = append(tokoBajus, tokoBaju)
	}
	return tokoBajus, nil
}

func InsertTokoBaju(db *sql.DB, tokoBaju structs.TokoBaju) error {
	sqlQuery := "INSERT INTO Toko_Baju (nama_toko) VALUES ($1)"
	_, err := db.Exec(sqlQuery, tokoBaju.NamaTokoBaju)
	if err != nil {
		log.Println("Error inserting TokoBaju:", err)
		return err
	}
	log.Println("TokoBaju inserted successfully")
	return nil
}

func UpdateTokoBaju(db *sql.DB, tokoBaju structs.TokoBaju) error {
	sqlQuery := "UPDATE Toko_Baju SET nama_toko = $1 WHERE id_TokoBaju = $2"
	_, err := db.Exec(sqlQuery, tokoBaju.NamaTokoBaju, tokoBaju.IdTokoBaju)
	if err != nil {
		log.Println("Error updating TokoBaju:", err)
		return err
	}
	log.Println("TokoBaju updated successfully")
	return nil
}

func DeleteTokoBaju(db *sql.DB, id int) error {
	sqlQuery := "DELETE FROM Toko_Baju WHERE id_TokoBaju = $1"
	_, err := db.Exec(sqlQuery, id)
	if err != nil {
		log.Println("Error deleting TokoBaju:", err)
		return err
	}
	log.Println("TokoBaju deleted successfully")
	return nil
}

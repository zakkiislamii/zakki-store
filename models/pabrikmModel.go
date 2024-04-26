package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

// GetAllPabrik retrieves all pabrik records from the database.
func GetAllPabrik(db *sql.DB) ([]structs.Pabrik, error) {
	sql := "SELECT * FROM Pabrik"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pabriks []structs.Pabrik
	for rows.Next() {
		var pabrik structs.Pabrik
		if err := rows.Scan(&pabrik.IdPabrik, &pabrik.NamaPabrik); err != nil {
			return nil, err
		}
		pabriks = append(pabriks, pabrik)
	}
	return pabriks, nil
}

// InsertPabrik inserts a new pabrik record into the database.
func InsertPabrik(db *sql.DB, pabrik structs.Pabrik) error {
	sql := "INSERT INTO Pabrik (nama_Pabrik) VALUES ($1)"
	_, err := db.Exec(sql, pabrik.NamaPabrik)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePabrik updates an existing pabrik record in the database.
func UpdatePabrik(db *sql.DB, pabrik structs.Pabrik) error {
	sql := "UPDATE Pabrik SET nama_Pabrik = $1 WHERE id_Pabrik = $2"
	_, err := db.Exec(sql, pabrik.NamaPabrik, pabrik.IdPabrik)
	if err != nil {
		log.Println("Error updating Pabrik:", err)
		return err
	}
	log.Println("Pabrik updated successfully")
	return nil
}

// DeletePabrik deletes a pabrik record from the database by its ID.
func DeletePabrik(db *sql.DB, id int) error {
	sql := "DELETE FROM Pabrik WHERE id_Pabrik = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		log.Println("Error deleting Pabrik:", err)
		return err
	}
	log.Println("Pabrik deleted successfully")
	return nil
}

package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetAllPabrik(db *sql.DB) ([]structs.Pabrik, error) {
	sql := "SELECT * FROM Pabrik"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []structs.Pabrik
	for rows.Next() {
		var Pabrik structs.Pabrik
		if err := rows.Scan(&Pabrik.IdPabrik, &Pabrik.NamaPabrik); err != nil {
			return nil, err
		}
		p = append(p, Pabrik)
	}
	return p, nil
}

func InsertPabrik(db *sql.DB, Pabrik structs.Pabrik) error {
	sql := "INSERT INTO Pabrik (nama_Pabrik) VALUES ($1)"
	_, err := db.Exec(sql, Pabrik.NamaPabrik)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePabrik(db *sql.DB, Pabrik structs.Pabrik) error {
	sql := "UPDATE Pabrik SET nama_Pabrik = $1,  WHERE id_Pabrik = $2"
	_, err := db.Exec(sql, Pabrik.NamaPabrik, Pabrik.IdPabrik)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting Pabrik:", err)
		return err
	}
	log.Println("Pabrik deleted successfully")
	return nil
}

func DeletePabrik(db *sql.DB, id int) error {
	sql := "DELETE FROM Pabrik WHERE id_Pabrik = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		// Log kesalahan jika terjadi
		log.Println("Error deleting Pabrik:", err)
		return err
	}
	log.Println("Pabrik deleted successfully")
	return nil
}

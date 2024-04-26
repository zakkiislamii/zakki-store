package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetAllProdukBaju(db *sql.DB) ([]structs.ProdukBaju, error) {
	sql := "SELECT pb.id_produk, pb.nama_produk, pb.harga, pb.stok, pb.id_toko, pb.id_pabrik, tb.nama_toko, p.nama_pabrik FROM Produk_Baju pb JOIN Toko_Baju tb ON pb.id_toko = tb.id_toko JOIN Pabrik p ON pb.id_pabrik = p.id_pabrik;"
	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var produkBajus []structs.ProdukBaju
	for rows.Next() {
		var produkBaju structs.ProdukBaju
		if err := rows.Scan(&produkBaju.IdProduk, &produkBaju.NamaProduk, &produkBaju.Harga, &produkBaju.Stok, &produkBaju.IdTokoBaju, &produkBaju.IdPabrik, &produkBaju.NamaTokoBaju, &produkBaju.NamaPabrik); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		produkBajus = append(produkBajus, produkBaju)
	}
	log.Println("Successfully retrieved all product data")
	return produkBajus, nil
}

func InsertProdukBaju(db *sql.DB, ProdukBaju structs.ProdukBaju) error {
	sql := "INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(sql, ProdukBaju.NamaProduk, ProdukBaju.Harga, ProdukBaju.Stok, ProdukBaju.IdTokoBaju, ProdukBaju.IdPabrik)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProdukBaju(db *sql.DB, ProdukBaju structs.ProdukBaju) error {
	sql := "UPDATE Produk_Baju SET nama_produk = $1, harga = $2, stok = $3, id_toko = $4, id_pabrik = $5 WHERE id_produk = $6"
	_, err := db.Exec(sql, ProdukBaju.NamaProduk, ProdukBaju.Harga, ProdukBaju.Stok, ProdukBaju.IdTokoBaju, ProdukBaju.IdPabrik, ProdukBaju.IdProduk)
	if err != nil {
		log.Println("Error updating ProdukBaju:", err)
		return err
	}
	log.Println("ProdukBaju updated successfully")
	return nil
}

func DeleteProdukBaju(db *sql.DB, id int) error {
	sql := "DELETE FROM Produk_Baju WHERE id_produk = $1"
	_, err := db.Exec(sql, id)
	if err != nil {
		log.Println("Error deleting ProdukBaju:", err)
		return err
	}
	log.Println("ProdukBaju deleted successfully")
	return nil
}

func GetProdukInfo(db *sql.DB) ([]structs.ProdukInfo, error) {
	query := `
		SELECT pb.nama_produk, pb.harga, pb.stok, t.nama_toko
		FROM Produk_Baju pb
		JOIN Toko_Baju t ON pb.id_toko = t.id_toko
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var produkInfo []structs.ProdukInfo

	for rows.Next() {
		var info structs.ProdukInfo
		err := rows.Scan(&info.NamaProduk, &info.Harga, &info.Stok, &info.NamaToko)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		produkInfo = append(produkInfo, info)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return produkInfo, nil
}

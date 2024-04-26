package models

import (
	"database/sql"
	"log"
	"zakki-store/structs"
)

func GetNamaProdukByID(db *sql.DB, idProduk int) (string, error) {
	var namaProduk string
	err := db.QueryRow("SELECT nama_produk FROM Produk_Baju WHERE id_produk = $1", idProduk).Scan(&namaProduk)
	if err != nil {
		log.Printf("Error getting product name by ID: %v\n", err)
		return "", err
	}
	log.Printf("Successfully retrieved product name for ID %d\n", idProduk)
	return namaProduk, nil
}

func GetNamaPelangganByID(db *sql.DB, idPelanggan int) (string, error) {
	var namaPelanggan string
	err := db.QueryRow("SELECT nama_pelanggan FROM Pelanggan WHERE id_pelanggan = $1", idPelanggan).Scan(&namaPelanggan)
	if err != nil {
		log.Printf("Error getting customer name by ID: %v\n", err)
		return "", err
	}
	log.Printf("Successfully retrieved customer name for ID %d\n", idPelanggan)
	return namaPelanggan, nil
}

func GetIDProdukByNamaProduk(db *sql.DB, namaProduk string) (int, error) {
	var idProduk int
	err := db.QueryRow("SELECT id_produk FROM Produk_Baju WHERE nama_produk = $1", namaProduk).Scan(&idProduk)
	if err != nil {
		log.Printf("Error getting product ID by product name: %v\n", err)
		return 0, err
	}
	log.Printf("Successfully retrieved product ID for product name %s\n", namaProduk)
	return idProduk, nil
}

func GetHargaProdukByNamaProduk(db *sql.DB, namaProduk string) (int, error) {
	var hargaProduk int
	err := db.QueryRow("SELECT harga FROM Produk_Baju WHERE nama_produk = $1", namaProduk).Scan(&hargaProduk)
	if err != nil {
		log.Printf("Error getting product price by product name: %v\n", err)
		return 0, err
	}
	log.Printf("Successfully retrieved product price for product name %s\n", namaProduk)
	return hargaProduk, nil
}

func BeliProduk(db *sql.DB, namaProduk string, jumlahBarang int, idPelanggan int) error {
	var idProduk int
	err := db.QueryRow("SELECT id_produk FROM Produk_Baju WHERE nama_produk = $1", namaProduk).Scan(&idProduk)
	if err != nil {
		log.Printf("Error getting product ID by product name: %v\n", err)
		return err
	}

	err = kurangiStokProduk(db, idProduk, jumlahBarang)
	if err != nil {
		return err
	}

	existingTransaksi, err := cekTransaksi(db, idProduk, idPelanggan)
	if err != nil {
		return err
	}

	if existingTransaksi {
		err = perbaruiTransaksi(db, idProduk, idPelanggan, jumlahBarang)
		if err != nil {
			return err
		}
	} else {
		err = buatTransaksiBaru(db, idProduk, idPelanggan, jumlahBarang)
		if err != nil {
			return err
		}
	}

	log.Println("Product purchase successful")
	return nil
}

func kurangiStokProduk(db *sql.DB, idProduk, jumlahBarang int) error {
	_, err := db.Exec("UPDATE Produk_Baju SET stok = stok - $1 WHERE id_produk = $2", jumlahBarang, idProduk)
	if err != nil {
		log.Printf("Error updating product stock: %v\n", err)
		return err
	}
	log.Printf("Product stock updated: Product ID %d, Quantity %d\n", idProduk, jumlahBarang)
	return nil
}

func cekTransaksi(db *sql.DB, idProduk, idPelanggan int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Transaksi WHERE id_produk = $1 AND id_pelanggan = $2", idProduk, idPelanggan).Scan(&count)
	if err != nil {
		log.Printf("Error checking existing transaction: %v\n", err)
		return false, err
	}
	return count > 0, nil
}

func perbaruiTransaksi(db *sql.DB, idProduk, idPelanggan, jumlahBarang int) error {
	_, err := db.Exec("UPDATE Transaksi SET jumlah_barang = jumlah_barang + $1, total_harga = total_harga + (SELECT harga FROM Produk_Baju WHERE id_produk = $2) * $1 WHERE id_produk = $2 AND id_pelanggan = $3", jumlahBarang, idProduk, idPelanggan)
	if err != nil {
		log.Printf("Error updating existing transaction: %v\n", err)
		return err
	}
	log.Printf("Transaction updated: Product ID %d, Customer ID %d, Quantity %d\n", idProduk, idPelanggan, jumlahBarang)
	return nil
}

func buatTransaksiBaru(db *sql.DB, idProduk, idPelanggan, jumlahBarang int) error {
	_, err := db.Exec("INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES (CURRENT_DATE, $1, (SELECT harga FROM Produk_Baju WHERE id_produk = $2) * $1, $2, $3)", jumlahBarang, idProduk, idPelanggan)
	if err != nil {
		log.Printf("Error inserting new transaction: %v\n", err)
		return err
	}
	log.Printf("New transaction created: Product ID %d, Customer ID %d, Quantity %d\n", idProduk, idPelanggan, jumlahBarang)
	return nil
}

func ViewTransaksi(db *sql.DB) ([]structs.ViewTransaksi, error) {
	rows, err := db.Query("SELECT t.tanggal_transaksi, t.jumlah_barang, t.total_harga, pb.nama_produk, p.nama_pelanggan FROM Transaksi t JOIN Produk_Baju pb ON t.id_produk = pb.id_produk JOIN Pelanggan p ON t.id_pelanggan = p.id_pelanggan")
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var transaksis []structs.ViewTransaksi

	for rows.Next() {
		var transaksi structs.ViewTransaksi
		err := rows.Scan(&transaksi.TanggalTransaksi, &transaksi.JumlahBarang, &transaksi.TotalHarga, &transaksi.NamaProduk, &transaksi.NamaPelanggan)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		transaksis = append(transaksis, transaksi)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	log.Println("Successfully retrieved transaction data")
	return transaksis, nil
}

func GetTransaksiByNamaPelanggan(db *sql.DB, namaPelanggan string) ([]structs.ViewTransaksi, error) {
	query := `
		SELECT t.tanggal_transaksi, t.jumlah_barang, t.total_harga, pb.nama_produk, p.nama_pelanggan 
		FROM Transaksi t 
		JOIN Produk_Baju pb ON t.id_produk = pb.id_produk 
		JOIN Pelanggan p ON t.id_pelanggan = p.id_pelanggan
		WHERE p.nama_pelanggan = $1
	`

	rows, err := db.Query(query, namaPelanggan)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var transaksis []structs.ViewTransaksi

	for rows.Next() {
		var transaksi structs.ViewTransaksi
		err := rows.Scan(&transaksi.TanggalTransaksi, &transaksi.JumlahBarang, &transaksi.TotalHarga, &transaksi.NamaProduk, &transaksi.NamaPelanggan)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		transaksis = append(transaksis, transaksi)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return transaksis, nil
}

func GetIDPelangganByNamaPelanggan(db *sql.DB, namaPelanggan string) (int, error) {
	var idPelanggan int
	err := db.QueryRow("SELECT id_pelanggan FROM Pelanggan WHERE nama_pelanggan = $1", namaPelanggan).Scan(&idPelanggan)
	if err != nil {
		log.Printf("Error getting customer ID by customer name: %v\n", err)
		return 0, err
	}
	log.Printf("Successfully retrieved customer ID for customer name %s\n", namaPelanggan)
	return idPelanggan, nil
}

func GetRiwayatTransaksiByUsername(db *sql.DB, username string) ([]structs.RiwayatTransaksi, error) {
	query := `
		SELECT t.tanggal_transaksi, t.jumlah_barang, t.total_harga, pb.nama_produk
		FROM Transaksi t 
		JOIN Produk_Baju pb ON t.id_produk = pb.id_produk 
		JOIN Pelanggan p ON t.id_pelanggan = p.id_pelanggan
		WHERE p.username = $1
	`

	rows, err := db.Query(query, username)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var riwayat []structs.RiwayatTransaksi

	for rows.Next() {
		var transaksi structs.RiwayatTransaksi
		err := rows.Scan(&transaksi.TanggalTransaksi, &transaksi.JumlahBarang, &transaksi.TotalHarga, &transaksi.NamaProduk)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		riwayat = append(riwayat, transaksi)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return riwayat, nil
}

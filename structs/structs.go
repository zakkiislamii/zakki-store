package structs

import "time"

type Pelanggan struct {
	IdPelanggan   int    `json:"id_pelanggan"`
	NamaPelanggan string `json:"nama_pelanggan"`
	NoHp          string `json:"no_hp"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type Pabrik struct {
	IdPabrik   int    `json:"id_pabrik"`
	NamaPabrik string `json:"nama_pabrik"`
}

type TokoBaju struct {
	IdTokoBaju   int    `json:"id_toko"`
	NamaTokoBaju string `json:"nama_toko"`
}

type ProdukBaju struct {
	IdProduk     int    `json:"id_produk"`
	NamaProduk   string `json:"nama_produk"`
	Harga        int    `json:"harga"`
	Stok         int    `json:"stok"`
	IdTokoBaju   int    `json:"id_toko"`
	IdPabrik     int    `json:"id_pabrik"`
	NamaTokoBaju string `json:"nama_toko"`
	NamaPabrik   string `json:"nama_pabrik"`
}

type Ulasan struct {
	IdUlasan      int    `json:"id_ulasan"`
	Ulasan        string `json:"ulasan"`
	Rating        int    `json:"rating"`
	IdProduk      int    `json:"id_produk"`
	IdPelanggan   int    `json:"id_pelanggan"`
	NamaProduk    string `json:"nama_produk"`
	NamaPelanggan string `json:"nama_pelanggan"`
}

type UlasanPelanggan struct {
	Ulasan        string `json:"ulasan"`
	Rating        int    `json:"rating"`
	NamaProduk    string `json:"nama_produk"`
	NamaPelanggan string `json:"nama_pelanggan"`
}

type Transaksi struct {
	IdTransaksi      int       `json:"id_transaksi"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	JumlahBarang     int       `json:"jumlah_barang"`
	TotalHarga       int       `json:"total_harga"`
	NamaProduk       string    `json:"nama_produk"`
	NamaPelanggan    string    `json:"nama_pelanggan"`
}
type PembelianRequest struct {
	JumlahBarang  int    `json:"jumlah_barang"`
	NamaProduk    string `json:"nama_produk"`
	NamaPelanggan string `json:"nama_pelanggan"`
}

type ProfilePelanggan struct {
	IdPelanggan   int    `json:"id_pelanggan"`
	NamaPelanggan string `json:"nama_pelanggan"`
	NoHp          string `json:"no_hp"`
	Username      string `json:"username"`
}

type ViewTransaksi struct {
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	JumlahBarang     int       `json:"jumlah_barang"`
	TotalHarga       int       `json:"total_harga"`
	NamaProduk       string    `json:"nama_produk"`
	NamaPelanggan    string    `json:"nama_pelanggan"`
}

type TransaksiResponse struct {
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	JumlahBarang     int       `json:"jumlah_barang"`
	TotalHarga       int       `json:"total_harga"`
	NamaProduk       string    `json:"nama_produk"`
	NamaPelanggan    string    `json:"nama_pelanggan"`
}

type RiwayatTransaksi struct {
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	JumlahBarang     int       `json:"jumlah_barang"`
	TotalHarga       int       `json:"total_harga"`
	NamaProduk       string    `json:"nama_produk"`
}

type ProdukInfo struct {
	NamaProduk string `json:"nama_produk"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	NamaToko   string `json:"nama_toko"`
}

type PelangganUlasan struct {
	Ulasan     string `json:"ulasan"`
	Rating     int    `json:"rating"`
	NamaToko   string `json:"nama_toko"`
	NamaProduk string `json:"nama_produk"`
	Username   string `json:"username"`
}

type RiwayatUlasan struct {
	Ulasan     string `json:"ulasan"`
	Rating     int    `json:"rating"`
	NamaProduk string `json:"nama_produk"`
	NamaToko   string `json:"nama_toko"`
}

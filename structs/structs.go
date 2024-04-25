package structs

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
	NamaTokoBaju string `json:"nama_toko"`
	NamaPabrik   string `json:"nama_pabrik"`
	IdTokoBaju   int    `json:"id_toko"`
	IdPabrik     int    `json:"id_pabrik"`
}

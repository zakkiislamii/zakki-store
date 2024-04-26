-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE Pabrik (
  id_pabrik SERIAL PRIMARY KEY,
  nama_pabrik VARCHAR(100) NOT NULL
);

CREATE TABLE Toko_Baju (
  id_toko SERIAL PRIMARY KEY,
  nama_toko VARCHAR(100) NOT NULL
);

CREATE TABLE Pelanggan (
  id_pelanggan SERIAL PRIMARY KEY,
  nama_pelanggan VARCHAR(100) NOT NULL,
  no_hp VARCHAR(15) NOT NULL,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL
);

CREATE TABLE Produk_Baju (
  id_produk SERIAL PRIMARY KEY,
  nama_produk VARCHAR(100) NOT NULL,
  harga INT NOT NULL,
  stok INT NOT NULL,
  id_toko INT NOT NULL,
  id_pabrik INT NOT NULL,
  FOREIGN KEY (id_toko) REFERENCES Toko_Baju(id_toko),
  FOREIGN KEY (id_pabrik) REFERENCES Pabrik(id_pabrik)
);

CREATE TABLE Ulasan (
  id_ulasan SERIAL PRIMARY KEY,
  ulasan TEXT NOT NULL,
  rating INT NOT NULL,
  id_produk INT NOT NULL,
  id_pelanggan INT NOT NULL,
  FOREIGN KEY (id_produk) REFERENCES Produk_Baju(id_produk),
  FOREIGN KEY (id_pelanggan) REFERENCES Pelanggan(id_pelanggan)
);

CREATE TABLE Transaksi (
  id_transaksi SERIAL PRIMARY KEY,
  tanggal_transaksi DATE NOT NULL,
  jumlah_barang INT NOT NULL,
  total_harga INT NOT NULL,
  id_produk INT NOT NULL,
  id_pelanggan INT NOT NULL,
  FOREIGN KEY (id_produk) REFERENCES Produk_Baju(id_produk),
  FOREIGN KEY (id_pelanggan) REFERENCES Pelanggan(id_pelanggan),
  UNIQUE (id_produk, id_pelanggan)
);
-- +migrate StatementEnd
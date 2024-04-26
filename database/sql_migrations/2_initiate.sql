-- +migrate Up
-- +migrate StatementBegin
INSERT INTO Pabrik (nama_pabrik) VALUES ('Textile World');
INSERT INTO Pabrik (nama_pabrik) VALUES ('Fashion Garment');
INSERT INTO Pabrik (nama_pabrik) VALUES ('Clothing Manufacturer');
INSERT INTO Pabrik (nama_pabrik) VALUES ('Garment Factory');
INSERT INTO Pabrik (nama_pabrik) VALUES ('Apparel Industry');
INSERT INTO Toko_Baju (nama_toko) VALUES ('Fashionista');
INSERT INTO Toko_Baju (nama_toko) VALUES ('Style Trend');
INSERT INTO Toko_Baju (nama_toko) VALUES ('Chic Boutique');
INSERT INTO Toko_Baju (nama_toko) VALUES ('Trendy Wear');
INSERT INTO Toko_Baju (nama_toko) VALUES ('Urban Outfit');
INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ('John Doe', '08123456789', 'johndoe', 'password123');
INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ('Jane Smith', '08765432109', 'janesmith', 'abc123');
INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ('David Lee', '08567894561', 'davidlee', 'qwerty');
INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ('Emily Brown', '08123987654', 'emilybrown', 'ilovecoding');
INSERT INTO Pelanggan (nama_pelanggan, no_hp, username, password) VALUES ('Michael Johnson', '08976543210', 'michaelj', 'letmein');
INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ('T-Shirt', 100000, 50, 1, 1);
INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ('Hoodie', 150000, 30, 2, 2);
INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ('Jeans', 200000, 40, 3, 3);
INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ('Dress', 180000, 25, 4, 4);
INSERT INTO Produk_Baju (nama_produk, harga, stok, id_toko, id_pabrik) VALUES ('Jacket', 220000, 35, 5, 5);
INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan) VALUES ('Produk bagus, nyaman dipakai.', 4, 1, 1);
INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan) VALUES ('Sangat puas dengan kualitasnya.', 5, 2, 2);
INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan) VALUES ('Ukuran sesuai dengan yang diharapkan.', 4, 3, 3);
INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan) VALUES ('Warnanya bagus, cocok untuk acara formal.', 4, 4, 4);
INSERT INTO Ulasan (ulasan, rating, id_produk, id_pelanggan) VALUES ('Desainnya keren, sangat memuaskan.', 5, 5, 5);
INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES ('2024-04-01', 2, 200000, 1, 1);
INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES ('2024-04-02', 1, 150000, 2, 2);
INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES ('2024-04-03', 3, 600000, 3, 3);
INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES ('2024-04-04', 2, 360000, 4, 4);
INSERT INTO Transaksi (tanggal_transaksi, jumlah_barang, total_harga, id_produk, id_pelanggan) VALUES ('2024-04-05', 1, 220000, 5, 5);

-- +migrate StatementEnd
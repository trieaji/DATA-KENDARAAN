package dataKendaraandto

type DatakendaraanResponse struct {
	ID             int    `json:"id"`
	NoRegistrasi   string `json:"no_regis"`
	NamaPemilik    string `json:"name_pemilik"`
	MerkKendaraan  string `json:"merk_kendaraan"`
	TahunPembuatan int    `json:"tahun_pembuatan"`
	Kapasitas      int    `json:"kapasitas"`
	Warna          string `json:"warna"`
	BahanBakar     string `json:"bahan_bakar"`
}

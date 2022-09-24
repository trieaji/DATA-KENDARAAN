package models

type DataKendaraan struct {
	ID             int    `json:"id" gorm:"primary_key:auto_increment"`
	NoRegistrasi   string `json:"no_regis" form:"no_regis" gorm:"type: varchar(255)" `
	NamaPemilik    string `json:"name_pemilik" form:"name_pemilik" gorm:"type: varchar(255)"`
	MerkKendaraan  string `json:"merk_kendaraan" form:"merk_kendaraan" gorm:"type: varchar(255)"`
	TahunPembuatan int    `json:"tahun_pembuatan" form:"tahun_pembuatan" gorm:"type: int"`
	Kapasitas      int    `json:"kapasitas" form:"kapasitas" gorm:"type: int"`
	Warna          string `json:"warna" form:"warna" gorm:"type: varchar(255)" `
	BahanBakar     string `json:"bahan_bakar" form:"bahan_bakar" gorm:"type: varchar(255)" `
}

type DatakendaraanRequest struct {
	NoRegistrasi   string `json:"no_regis" form:"no_regis" gorm:"type: varchar(255)" `
	NamaPemilik    string `json:"name_pemilik" form:"name_pemilik" gorm:"type: varchar(255)"`
	MerkKendaraan  string `json:"merk_kendaraan" form:"merk_kendaraan" gorm:"type: varchar(255)"`
	TahunPembuatan int    `json:"tahun_pembuatan" form:"tahun_pembuatan" gorm:"type: int"`
	Kapasitas      int    `json:"kapasitas" form:"kapasitas" gorm:"type: int"`
	Warna          string `json:"warna" form:"warna" gorm:"type: varchar(255)" `
	BahanBakar     string `json:"bahan_bakar" form:"bahan_bakar" gorm:"type: varchar(255)" `
}

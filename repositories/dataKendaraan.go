package repositories

import (
	"prokdrn/models"

	"gorm.io/gorm"
)

type DataKendaraanRepository interface {
	FindDatas() ([]models.DataKendaraan, error)
	GetData(ID int) (models.DataKendaraan, error)
	CreateData(datakndrn models.DataKendaraan) (models.DataKendaraan, error)
	UpdateData(datakndrn models.DataKendaraan, ID int) (models.DataKendaraan, error)
	DeleteData(datakndrn models.DataKendaraan, ID int) (models.DataKendaraan, error)
}

// Create RepositoryUser function here ...
func RepositoryDataKendaraan(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDatas() ([]models.DataKendaraan, error) {
	var datas []models.DataKendaraan
	err := r.db.Find(&datas).Error

	return datas, err
}

func (r *repository) GetData(ID int) (models.DataKendaraan, error) {
	var data models.DataKendaraan
	err := r.db.First(&data, ID).Error

	return data, err
}

func (r *repository) CreateData(datakndrn models.DataKendaraan) (models.DataKendaraan, error) {
	err := r.db.Create(&datakndrn).Error

	return datakndrn, err
}

func (r *repository) UpdateData(datakndrn models.DataKendaraan, ID int) (models.DataKendaraan, error) {
	err := r.db.Save(&datakndrn).Error

	return datakndrn, err
}

func (r *repository) DeleteData(datakndrn models.DataKendaraan, ID int) (models.DataKendaraan, error) {
	err := r.db.Delete(&datakndrn).Error

	return datakndrn, err
}

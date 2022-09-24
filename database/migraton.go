package database

import (
	"fmt"
	"prokdrn/models"
	"prokdrn/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.DataKendaraan{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Berhasil Luur")
}

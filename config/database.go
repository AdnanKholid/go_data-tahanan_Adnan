package config

import (
	"fmt"
	"mini_project/helper"
	msel "mini_project/models/m_sel"
	msipir "mini_project/models/m_sipir"
	mtahanan "mini_project/models/m_tahanan"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	var dsn string = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GetConfig("DB_USERNAME"),
		helper.GetConfig("DB_PASSWORD"),
		helper.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	AutoMigration()
	fmt.Println("Database Connection")
}

func InitTestDB() {
	var err error
	var dsn string = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GetConfig("DB_USERNAME"),
		helper.GetConfig("DB_PASSWORD"),
		helper.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	AutoMigration()
	fmt.Println("Database Connection")
}

func SeedTahanan() mtahanan.Tahanan {

	var tahanan mtahanan.Tahanan = mtahanan.Tahanan{
		SelID:         1,
		Nama:          "sample",
		Usia:          19,
		MasaTahanan:   "5 tahun",
		Pelanggaran:   "penculikan",
		TanggalMasuk:  "20-08-2005",
		TanggalKeluar: "21-02-2010",
	}

	if err := DB.Create(&tahanan).Error; err != nil {
		panic(err)
	}
	var storage mtahanan.Tahanan
	DB.Last(&storage)
	return storage
}

func SeedSel() msel.Sel {
	var tahanan []mtahanan.Tahanan
	var sel msel.Sel = msel.Sel{
		NoSel:   021,
		SipirID: 12,
		Tahanan: tahanan,
	}

	if err := DB.Create(&sel).Error; err != nil {
		panic(err)
	}
	var storage msel.Sel
	DB.Last(&storage)
	return storage
}

func SeedSipir() msipir.Sipir {
	var sel []msel.Sel
	var tahanan msipir.Sipir = msipir.Sipir{
		Nama:    "tes",
		Sel:     sel,
		Jabatan: "perwira",
	}

	if err := DB.Create(&tahanan).Error; err != nil {
		panic(err)
	}
	var storage msipir.Sipir
	DB.Last(&storage)
	return storage
}

package mtahanan

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Tahanan struct {
	ID            uint           `json:"id"`
	SelID         uint           `json:"sel_id"`
	Nama          string         `json:"nama"`
	Usia          uint           `json:"usia"`
	MasaTahanan   string         `json:"hukuman"`
	Pelanggaran   string         `json:"pelanggaran"`
	TanggalMasuk  string         `json:"tanggal_masuk"`
	TanggalKeluar string         `json:"tanggal_keluar"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type Response struct {
	ID            uint           `json:"id"`
	SelID         uint           `json:"sel_id"`
	Nama          string         `json:"nama"`
	Usia          uint           `json:"usia"`
	MasaTahanan   string         `json:"hukuman"`
	Pelanggaran   string         `json:"pelanggaran"`
	TanggalMasuk  string         `json:"tanggal_masuk"`
	TanggalKeluar string         `json:"tanggal_keluar"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type Input struct {
	SelID         uint   `json:"sel_id"`
	Nama          string `json:"nama"`
	Usia          uint   `json:"usia"`
	MasaTahanan   string `json:"hukuman"`
	Pelanggaran   string `json:"pelanggaran"`
	TanggalMasuk  string `json:"tanggal_masuk"`
	TanggalKeluar string `json:"tanggal_keluar"`
}

func (input *Input) Validator() error {
	validator := validator.New()

	err := validator.Struct(input)

	return err
}

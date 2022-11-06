package msel

import (
	mtahanan "mini_project/models/m_tahanan"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Sel struct {
	ID        uint               `json:"id"`
	NoSel     uint               `json:"no_sel"`
	SipirID   uint               `json:"sipir_id"`
	Tahanan   []mtahanan.Tahanan `json:"tahanan"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at"`
}

type Response struct {
	ID      uint `json:"id"`
	NoSel   uint `json:"no_sel"`
	SipirID uint `json:"sipir_id"`
	Tahanan []struct {
		ID            uint   `json:"id"`
		Nama          string `json:"nama"`
		Usia          uint   `json:"usia"`
		MasaTahanan   string `json:"hukuman"`
		Pelanggaran   string `json:"pelanggaran"`
		TanggalMasuk  string `json:"tanggal_masuk"`
		TanggalKeluar string `json:"tanggal_keluar"`
	}
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Input struct {
	NoSel   uint `json:"no_sel" validate:"required"`
	SipirID uint `json:"sipir_id" validate:"required"`
}

func (input *Input) Validator() error {
	validator := validator.New()

	err := validator.Struct(input)

	return err
}

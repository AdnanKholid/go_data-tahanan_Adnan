package msipir

import (
	msel "mini_project/models/m_sel"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Sipir struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Sel       []msel.Sel
	Jabatan   string         `json:"jabatan"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Response struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
	Sel  []struct {
		ID    uint `json:"id"`
		NoSel uint `json:"no_sel"`
	}
	Jabatan   string         `json:"jabatan"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Input struct {
	Nama    string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

func (input *Input) Validator() error {
	validator := validator.New()

	err := validator.Struct(input)

	return err
}

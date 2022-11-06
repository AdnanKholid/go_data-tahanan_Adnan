package rsipir

import (
	"mini_project/config"
	msipir "mini_project/models/m_sipir"

	"github.com/jinzhu/copier"
)

type SipirRepoImpl struct{}

func (r *SipirRepoImpl) GetAll() []msipir.Response {
	var sipir []msipir.Sipir
	var response []msipir.Response

	config.DB.Preload("Sel").Find(&sipir)
	copier.Copy(&response, sipir)

	return response

}

func (r *SipirRepoImpl) GetByID(id string) msipir.Response {
	var sipir msipir.Sipir
	var response msipir.Response

	config.DB.Preload("Sel").First(&sipir, "id=?", id)
	copier.Copy(&response, sipir)

	return response

}

func (r *SipirRepoImpl) Create(input msipir.Input) msipir.Response {
	var storage msipir.Sipir = msipir.Sipir{}
	var response msipir.Response = msipir.Response{}

	var sipir msipir.Sipir = msipir.Sipir{
		Nama:    input.Nama,
		Jabatan: input.Jabatan,
	}
	result := config.DB.Preload("Sel").Create(&sipir)
	result.Last(&storage)
	copier.Copy(&response, storage)

	return response

}

func (r *SipirRepoImpl) Update(id string, input msipir.Input) msipir.Response {
	var response msipir.Response

	var sipir msipir.Sipir
	config.DB.Preload("Sel").First(&sipir, "id=?", id)

	sipir.Nama = input.Nama
	sipir.Jabatan = input.Jabatan

	config.DB.Preload("Sel").Save(&sipir)
	copier.Copy(&response, sipir)

	return response

}

func (r *SipirRepoImpl) Delete(id string) bool {
	var sipir msipir.Sipir
	config.DB.First(&sipir, "id=?", id)

	result := config.DB.Delete(&sipir)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

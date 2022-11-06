package rtahanan

import (
	"mini_project/config"
	mtahanan "mini_project/models/m_tahanan"

	"github.com/jinzhu/copier"
)

type TahananRepoImpl struct{}

func (r *TahananRepoImpl) GetAll() []mtahanan.Response {
	var tahanan []mtahanan.Tahanan
	var response []mtahanan.Response

	config.DB.Find(&tahanan)
	copier.Copy(&response, tahanan)

	return response

}

func (r *TahananRepoImpl) GetByID(id string) mtahanan.Response {
	var tahanan mtahanan.Tahanan
	var response mtahanan.Response

	config.DB.First(&tahanan, "id=?", id)
	copier.Copy(&response, tahanan)

	return response

}

func (r *TahananRepoImpl) Create(input mtahanan.Input) mtahanan.Response {
	var storage mtahanan.Tahanan = mtahanan.Tahanan{}
	var response mtahanan.Response = mtahanan.Response{}

	var tahanan mtahanan.Tahanan = mtahanan.Tahanan{
		SelID:         input.SelID,
		Nama:          input.Nama,
		Usia:          input.Usia,
		MasaTahanan:   input.MasaTahanan,
		Pelanggaran:   input.Pelanggaran,
		TanggalMasuk:  input.TanggalMasuk,
		TanggalKeluar: input.TanggalKeluar,
	}
	result := config.DB.Create(&tahanan)
	result.Last(&storage)
	copier.Copy(&response, storage)

	return response

}

func (r *TahananRepoImpl) Update(id string, input mtahanan.Input) mtahanan.Response {
	var response mtahanan.Response

	var tahanan mtahanan.Tahanan
	config.DB.First(&tahanan, "id=?", id)

	tahanan.SelID = input.SelID
	tahanan.Nama = input.Nama
	tahanan.Usia = input.Usia
	tahanan.MasaTahanan = input.MasaTahanan
	tahanan.Pelanggaran = input.Pelanggaran
	tahanan.TanggalMasuk = input.TanggalMasuk
	tahanan.TanggalKeluar = input.TanggalKeluar

	config.DB.Save(&tahanan)
	copier.Copy(&response, tahanan)

	return response

}

func (r *TahananRepoImpl) Delete(id string) bool {
	var tahanan mtahanan.Tahanan
	config.DB.First(&tahanan, "id=?", id)

	result := config.DB.Unscoped().Delete(&tahanan)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

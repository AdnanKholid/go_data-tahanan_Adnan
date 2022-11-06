package rsel

import (
	"mini_project/config"
	msel "mini_project/models/m_sel"

	"github.com/jinzhu/copier"
)

type SelRepoImpl struct{}

func (r *SelRepoImpl) GetAll() []msel.Response {
	var sel []msel.Sel
	var response []msel.Response

	config.DB.Preload("Tahanan").Find(&sel)
	copier.Copy(&response, sel)

	return response

}

func (r *SelRepoImpl) GetByID(id string) msel.Response {
	var sel msel.Sel
	var response msel.Response

	config.DB.Preload("Tahanan").First(&sel, "id=?", id)
	copier.Copy(&response, sel)

	return response

}

func (r *SelRepoImpl) Create(input msel.Input) msel.Response {
	var storage msel.Sel = msel.Sel{}
	var response msel.Response = msel.Response{}

	var sel msel.Sel = msel.Sel{
		NoSel:   input.NoSel,
		SipirID: input.SipirID,
	}
	result := config.DB.Preload("Tahanan").Create(&sel)
	result.Last(&storage)
	copier.Copy(&response, storage)

	return response

}

func (r *SelRepoImpl) Update(id string, input msel.Input) msel.Response {
	var response msel.Response

	var sel msel.Sel
	config.DB.Preload("Tahanan").First(&sel, "id=?", id)

	sel.NoSel = input.NoSel
	sel.SipirID = input.SipirID

	config.DB.Preload("Tahanan").Save(&sel)
	copier.Copy(&response, sel)

	return response

}

func (r *SelRepoImpl) Delete(id string) bool {
	var sel msel.Sel
	config.DB.First(&sel, "id=?", id)

	result := config.DB.Delete(&sel)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

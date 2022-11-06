package config

import (
	msel "mini_project/models/m_sel"
	msipir "mini_project/models/m_sipir"
	mtahanan "mini_project/models/m_tahanan"
	muser "mini_project/models/m_user"
)

func AutoMigration() {
	DB.AutoMigrate(&msipir.Sipir{}, &muser.User{}, &msel.Sel{}, &mtahanan.Tahanan{})
}

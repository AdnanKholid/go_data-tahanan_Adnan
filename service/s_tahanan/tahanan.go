package service

import (
	mtahanan "mini_project/models/m_tahanan"
	"mini_project/repository"
	rtahanan "mini_project/repository/r_tahanan"
)

type TahananService struct {
	repo repository.TahananRepo
}

func NewTahanan() TahananService {
	return TahananService{
		repo: &rtahanan.TahananRepoImpl{},
	}
}

func (s *TahananService) GetAll() []mtahanan.Response {
	return s.repo.GetAll()
}

func (s *TahananService) GetByID(id string) mtahanan.Response {
	return s.repo.GetByID(id)
}

func (s *TahananService) Create(input mtahanan.Input) mtahanan.Response {
	return s.repo.Create(input)
}

func (s *TahananService) Update(id string, input mtahanan.Input) mtahanan.Response {
	return s.repo.Update(id, input)
}

func (s *TahananService) Delete(id string) bool {
	return s.repo.Delete(id)
}

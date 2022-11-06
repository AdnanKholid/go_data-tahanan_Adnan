package service

import (
	msel "mini_project/models/m_sel"
	"mini_project/repository"
	rsel "mini_project/repository/r_sel"
)

type SelService struct {
	repo repository.SelRepo
}

func NewSel() SelService {
	return SelService{
		repo: &rsel.SelRepoImpl{},
	}
}

func (s *SelService) GetAll() []msel.Response {
	return s.repo.GetAll()
}

func (s *SelService) GetByID(id string) msel.Response {
	return s.repo.GetByID(id)
}

func (s *SelService) Create(input msel.Input) msel.Response {
	return s.repo.Create(input)
}

func (s *SelService) Update(id string, input msel.Input) msel.Response {
	return s.repo.Update(id, input)
}

func (s *SelService) Delete(id string) bool {
	return s.repo.Delete(id)
}

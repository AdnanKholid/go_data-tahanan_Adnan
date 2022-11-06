package service

import (
	msipir "mini_project/models/m_sipir"
	"mini_project/repository"
	rsipir "mini_project/repository/r_sipir"
)

type SipirService struct {
	repo repository.SipirRepo
}

func NewSipir() SipirService {
	return SipirService{
		repo: &rsipir.SipirRepoImpl{},
	}
}

func (s *SipirService) GetAll() []msipir.Response {
	return s.repo.GetAll()
}

func (s *SipirService) GetByID(id string) msipir.Response {
	return s.repo.GetByID(id)
}

func (s *SipirService) Create(input msipir.Input) msipir.Response {
	return s.repo.Create(input)
}

func (s *SipirService) Update(id string, input msipir.Input) msipir.Response {
	return s.repo.Update(id, input)
}

func (s *SipirService) Delete(id string) bool {
	return s.repo.Delete(id)
}

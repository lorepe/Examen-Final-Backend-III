package odontologo

import (
	"Final/internal/domain"
	"Final/internal/odontologo"
)

type SeviceOdontologo interface {
	//FIXME revisar error
	GetAll() []domain.Odontologo
}

type service struct{
	repo RepositoryOdontologo
}

func NewService(repository RepositoryOdontologo)SeviceOdontologo{
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Odontologo){
	odontologos:=s.repo.GetAll()
	return odontologos
}

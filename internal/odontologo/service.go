package odontologo

import (
	"Final/internal/domain"
)

type SeviceOdontologo interface {
	GetAll() ([]domain.Odontologo, error)
	CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error)
	GetOdontologoById(id int) (domain.Odontologo, error)
	UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error)
}

type service struct {
	repo RepositoryOdontologo
}

func NewService(repository RepositoryOdontologo) SeviceOdontologo {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Odontologo, error) {
	odontologos, err := s.repo.GetAll()
	if err != nil {
		return []domain.Odontologo{}, err
	}
	return odontologos, nil
}

func (s *service) CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error) {

	odontologo, err := s.repo.CreateOdontologo(o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil

}

func (s *service) GetOdontologoById(id int) (domain.Odontologo, error) {
	odontologo, err := s.repo.GetOdontologoById(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *service) UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error) {
	odontologo, err:= s.repo.UpdateOdontologo(id,o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

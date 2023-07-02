package odontologo

import (
	"Final/internal/domain"
	"errors"
)

type SeviceOdontologo interface {
	GetAll() ([]domain.Odontologo, error)
	CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error)
	GetOdontologoById(id int) (domain.Odontologo, error)
	UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error)
	UpdateMatricula(id int, o domain.Odontologo) (domain.Odontologo, error)
	DeleteOdontologo(id int) error
}

type service struct {
	repo RepositoryOdontologo
}

func NewService(repository RepositoryOdontologo) SeviceOdontologo {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Odontologo, error) {
	return s.repo.GetAll()
}

func (s *service) CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error) {

	odontologo, err := s.repo.CreateOdontologo(o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil

}

func (s *service) GetOdontologoById(id int) (domain.Odontologo, error) {
	return s.repo.GetOdontologoById(id)
}

func (s *service) UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error) {
	od, err := s.repo.GetOdontologoById(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	if od.Matricula != o.Matricula {
		return domain.Odontologo{}, errors.New("Registration must be equal")
	}
	return s.repo.UpdateOdontologo(id, o)
}

func (s *service) UpdateMatricula(id int, o domain.Odontologo) (domain.Odontologo, error) {
	return s.repo.UpdateOdontologo(id, o)
}

func (s *service) DeleteOdontologo(id int) error {
	return s.repo.DeleteOdontologo(id)
}

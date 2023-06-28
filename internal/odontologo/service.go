package odontologo

import (
	"Final/internal/domain"
)

type SeviceOdontologo interface {
	GetAll() ([]domain.Odontologo, error)
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

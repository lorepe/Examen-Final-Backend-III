package odontologo

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryOdontologo interface {
	GetAll() ([]domain.Odontologo, error)
}
type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) RepositoryOdontologo {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Odontologo, error) {
	odontologos, err := r.storage.GetAllOdontologos()
	if err != nil {
		return []domain.Odontologo{}, errors.New("list not found")
	}
	return odontologos, nil
}

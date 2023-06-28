package odontologo

import (
	"Final/internal/domain"
	"Final/pkg/store"

)

type RepositoryOdontologo interface {
	
	GetAll() []domain.Odontologo
}
type repository struct{
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) RepositoryOdontologo {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Odontologo) {
	odontologos, err := r.storage.GetAllOdontologos()
	if err != nil {
		return []domain.Odontologo{}
	}
	return odontologos
}
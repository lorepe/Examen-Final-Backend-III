package odontologo

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryOdontologo interface {
	GetAll() ([]domain.Odontologo, error)
	CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error)
	GetOdontologoById(id int) (domain.Odontologo, error)
	UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error)
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

func (r *repository) CreateOdontologo(o domain.Odontologo) (domain.Odontologo, error) {
	//TODO Validate matricula

	err := r.storage.PostOdontologo(o)
	if err != nil {
		return domain.Odontologo{}, errors.New("Error creating dentist")
	}
	return o, nil

}

func (r *repository) GetOdontologoById(id int) (domain.Odontologo, error) {
	odontologo, err := r.storage.GetOdontologoById(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("dentist not found")
	}
	return odontologo, nil
}

func (r *repository) UpdateOdontologo(id int, o domain.Odontologo) (domain.Odontologo, error) {
	//TODO validate matricula

	err := r.storage.UpdateOdontologo(id, o)
	if err != nil {
		return domain.Odontologo{},errors.New("Error updating dentist")
	}
	return o, nil

}

package paciente

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryPaciente interface {
	GetAll() ([]domain.Paciente, error)
}
type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) RepositoryPaciente {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Paciente, error) {
	pacientes, err := r.storage.GetAllPacientes()
	if err != nil {
		return []domain.Paciente{}, errors.New("list not found")
	}
	return pacientes, nil
}

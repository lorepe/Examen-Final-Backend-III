package paciente

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryPaciente interface {
	GetAll() ([]domain.Paciente, error)
	CreatePaciente(p domain.Paciente) (domain.Paciente, error)
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

func (r *repository) CreatePaciente(p domain.Paciente) (domain.Paciente, error) {
	//TODO Validate dni

	err := r.storage.PostPaciente(p)
	if err != nil {
		return domain.Paciente{}, errors.New("Error creating paciente")
	}
	return p, nil

}

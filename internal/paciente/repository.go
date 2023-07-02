package paciente

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryPaciente interface {
	GetAll() ([]domain.Paciente, error)
	CreatePaciente(p domain.Paciente) (domain.Paciente, error)
	GetPacienteById(id int) (domain.Paciente, error)
	UpdatePaciente(int, domain.Paciente) (domain.Paciente, error)
	DeletePaciente(int) error
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
	validacion, err := r.storage.VerificarDni(p.Dni)
	if err != nil {
		return domain.Paciente{}, err
	}
	if validacion == true {
		return domain.Paciente{}, errors.New("Dni used")
	}
	err = r.storage.PostPaciente(p)
	if err != nil {
		return domain.Paciente{}, errors.New("Error creating paciente")
	}
	return p, nil
}

func (r *repository) GetPacienteById(id int) (domain.Paciente, error) {
	paciente, err := r.storage.GetPacienteById(id)
	if err != nil {
		return domain.Paciente{}, errors.New("patient not found")
	}
	return paciente, nil
}

func (r *repository) UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.UpdatePaciente(id, p)
	if err != nil {
		return domain.Paciente{}, errors.New("Error updating patient")
	}
	return p, nil

}

func (r *repository) DeletePaciente(id int) error {
	err := r.storage.DeletePaciente(id)
	if err != nil {
		return errors.New("id not found")
	}
	return nil

}

package paciente

import (
	"Final/internal/domain"
	"errors"
)

type ServicePaciente interface {
	GetAll() ([]domain.Paciente, error)
	CreatePaciente(domain.Paciente) (domain.Paciente, error)
	GetPacienteById(int) (domain.Paciente, error)
	UpdatePaciente(int, domain.Paciente) (domain.Paciente, error)
	DeletePaciente(int) error
	UpdateDni(int, domain.Paciente) (domain.Paciente, error)
}

type service struct {
	repo RepositoryPaciente
}

func NewService(repository RepositoryPaciente) ServicePaciente {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Paciente, error) {
	return s.repo.GetAll()
}

func (s *service) CreatePaciente(p domain.Paciente) (domain.Paciente, error) {
	return s.repo.CreatePaciente(p)
}

func (s *service) GetPacienteById(id int) (domain.Paciente, error) {
	return s.repo.GetPacienteById(id)
}


func (s *service) UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error) {
	pacienteID, err := s.repo.GetPacienteById(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if pacienteID.Dni != p.Dni {
		return domain.Paciente{}, errors.New("DNI must be equal")
	}
	return s.repo.UpdatePaciente(id, p)
}

func (s *service) UpdateDni(id int, p domain.Paciente) (domain.Paciente, error) {
	return s.repo.UpdatePaciente(id, p)
}

func (s *service) DeletePaciente(id int) error {
	return s.repo.DeletePaciente(id)
}

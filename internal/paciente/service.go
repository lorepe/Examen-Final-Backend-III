package paciente

import (
	"Final/internal/domain"
)

type ServicePaciente interface {
	GetAll() ([]domain.Paciente, error)
	CreatePaciente(domain.Paciente) (domain.Paciente, error)
	GetPacienteById(id int) (domain.Paciente, error)
	UpdatePaciente(int, domain.Paciente) (domain.Paciente, error)
}

type service struct {
	repo RepositoryPaciente
}

func NewService(repository RepositoryPaciente) ServicePaciente {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Paciente, error) {
	pacientes, err := s.repo.GetAll()
	if err != nil {
		return []domain.Paciente{}, err
	}
	return pacientes, nil
}

func (s *service) CreatePaciente(p domain.Paciente) (domain.Paciente, error) {

	paciente, err := s.repo.CreatePaciente(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil

}

func (s *service) GetPacienteById(id int) (domain.Paciente, error) {
	paciente, err := s.repo.GetPacienteById(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

// FIXME Reemplazar por valores predeterminados para el patch
func (s *service) UpdatePaciente(id int, p domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.repo.UpdatePaciente(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

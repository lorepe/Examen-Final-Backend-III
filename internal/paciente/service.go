package paciente

import (
	"Final/internal/domain"
)

type ServicePaciente interface {
	GetAll() ([]domain.Paciente, error)
	CreatePaciente(domain.Paciente) (domain.Paciente, error)
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

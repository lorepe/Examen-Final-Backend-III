package turno

import "Final/internal/domain"

type ServiceTurno interface {
	GetAll() ([]domain.Turno, error)
}

type service struct {
	repo RepositoryTurno
}

func NewService(repository RepositoryTurno) ServiceTurno {
	return &service{repository}

}

func (s *service) GetAll() ([]domain.Turno, error) {
	turnos, err := s.repo.GetAll()
	if err != nil {
		return []domain.Turno{}, err
	}
	return turnos, nil
}

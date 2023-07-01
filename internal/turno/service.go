package turno

import "Final/internal/domain"

type ServiceTurno interface {
	GetAll() ([]domain.Turno, error)
	CreateTurno(domain.Turno) (domain.Turno, error)
	GetTurnoById(int) (domain.Turno, error)
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

func (s *service) CreateTurno(t domain.Turno) (domain.Turno, error) {

	turno, err := s.repo.CreateTurno(t)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *service) GetTurnoById(id int) (domain.Turno, error) {
	return s.repo.GetTurnoById(id)
}

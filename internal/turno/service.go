package turno

import "Final/internal/domain"

type ServiceTurno interface {
	GetAll() ([]domain.Turno, error)
	CreateTurno(domain.Turno) (domain.Turno, error)
	GetTurnoById(int) (domain.Turno, error)
	UpdateTurno(int, domain.Turno) (domain.Turno, error)
	DeleteTurno(int) error
	CreateTurnoaAuxiliar(ta domain.TurnoAuxiliar) error
	GetAllByDni(int)([]domain.Turno,error)
}

type service struct {
	repo RepositoryTurno
}

func NewService(repository RepositoryTurno) ServiceTurno {
	return &service{repository}

}

func (s *service) GetAll() ([]domain.Turno, error) {
	return s.repo.GetAll()
}

func (s *service) CreateTurno(t domain.Turno) (domain.Turno, error) {
	return s.repo.CreateTurno(t)
}

func (s *service) GetTurnoById(id int) (domain.Turno, error) {
	return s.repo.GetTurnoById(id)
}

// FIXME Reemplazar por valores predeterminados para el patch
func (s *service) UpdateTurno(id int, t domain.Turno) (domain.Turno, error) {
	return s.repo.UpdateTurno(id, t)
	
}

func (s *service) DeleteTurno(id int) error {
	return s.repo.DeleteTurno(id)
	
}

func (s *service) CreateTurnoaAuxiliar(ta domain.TurnoAuxiliar) error {

	err := s.repo.CreateTurnoDNIMat(ta)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAllByDni(dni int) ([]domain.Turno, error) {
	turnos, err := s.repo.GetAllByDni(dni)
	if err != nil {
		return []domain.Turno{}, err
	}
	return turnos, nil
}

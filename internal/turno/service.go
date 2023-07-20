package turno

import "Final/internal/domain"

type ServiceTurno interface {
	GetAll() ([]domain.Turno, error)
	CreateTurno(domain.Turno) (domain.Turno, error)
	GetTurnoById(int) (domain.Turno, error)
	UpdateTurno(int, domain.Turno) (domain.Turno, error)
	UpdateCitaTurno(int, string) (domain.Turno, error)
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


func (s *service) UpdateTurno(id int, t domain.Turno) (domain.Turno, error) {
	_, err := s.repo.GetTurnoById(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return s.repo.UpdateTurno(id, t)
	
}

func (s *service) DeleteTurno(id int) error {
	_, err := s.repo.GetTurnoById(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteTurno(id)
	
}

func (s *service) CreateTurnoaAuxiliar(ta domain.TurnoAuxiliar) error {
	
	return s.repo.CreateTurnoDNIMat(ta)
}

func (s *service) GetAllByDni(dni int) ([]domain.Turno, error) {
	return s.repo.GetAllByDni(dni)
}

func(s *service) UpdateCitaTurno(id int, fechaHora string) (domain.Turno, error){
	t, err := s.repo.GetTurnoById(id)
	if err != nil {
		return domain.Turno{}, err
	}
	t.FechaHora = fechaHora
	return s.repo.UpdateTurno(id, t)
}

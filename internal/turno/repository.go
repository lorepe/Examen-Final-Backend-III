package turno

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryTurno interface {
	GetAll() ([]domain.Turno, error)
	CreateTurno(domain.Turno) (domain.Turno, error)
	GetTurnoById(int) (domain.Turno, error)
	UpdateTurno(int, domain.Turno) (domain.Turno, error)
	DeleteTurno(int) error
	CreateTurnoDNIMat(domain.TurnoAuxiliar) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) RepositoryTurno {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Turno, error) {
	turnos, err := r.storage.GetAllTurnos()
	if err != nil {
		return []domain.Turno{}, errors.New("list not found")
	}
	return turnos, nil
}

func (r *repository) CreateTurno(t domain.Turno) (domain.Turno, error) {

	err := r.storage.PostTurno(t)

	if err != nil {
		return domain.Turno{}, errors.New("Error creating appointment")
	}
	return t, nil
}

func (r *repository) GetTurnoById(id int) (domain.Turno, error) {
	turno, err := r.storage.GetTurnoById(id)
	if err != nil {
		// return domain.Turno{}, errors.New("appointment not found")

		return domain.Turno{}, err
	}
	return turno, nil
}

func (r *repository) UpdateTurno(id int, t domain.Turno) (domain.Turno, error) {
	//TODO validate dni
	err := r.storage.UpdateTurno(id, t)
	if err != nil {
		return domain.Turno{}, errors.New("Error updating appointment")
	}
	return t, nil
}

func (r *repository) DeleteTurno(id int) error {

	err := r.storage.DeleteTurno(id)
	//FIXME ESTABLECER UN ERROR PARA ID INEXISTENTE
	//FIXME ESTABLECER ERROR PARA PACIENTE CON TURNO
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateTurnoDNIMat(ta domain.TurnoAuxiliar) error {
	//FIXME RETURN
	err := r.storage.PostTurnoDNIMat(ta)

	if err != nil {
		return errors.New("Error creating appointment")
	}
	return nil
}

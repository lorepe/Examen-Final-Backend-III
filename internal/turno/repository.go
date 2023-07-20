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
	GetAllByDni(int) ([]domain.Turno, error)
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
		return domain.Turno{}, errors.New("appointment not found")
	}
	return turno, nil
}

func (r *repository) UpdateTurno(id int, t domain.Turno) (domain.Turno, error) {
	err := r.storage.UpdateTurno(id, t)
	if err != nil {
		return domain.Turno{}, errors.New("Error updating appointment")
	}
	return t, nil
}

func (r *repository) DeleteTurno(id int) error {

	err := r.storage.DeleteTurno(id)
	if err != nil {
		return errors.New("id not found")
	}
	return nil
}

func (r *repository) CreateTurnoDNIMat(ta domain.TurnoAuxiliar) error {
	listaPacientes, err := r.storage.GetAllPacientes()
	if err != nil {
		return err
	}
	var paciente domain.Paciente
	for _, p := range listaPacientes {
		if (p.Dni == ta.PacienteDni) {
			paciente.Id = p.Id
			paciente.Nombre = p.Nombre
			paciente.Apellido = p.Apellido
			paciente.Dni = p.Dni
			paciente.FechaAlta = p.FechaAlta
			
			}
		} 
		
	listaOdontologos, err := r.storage.GetAllOdontologos()
	if err != nil {
		return err
	}
	var odontologo domain.Odontologo
	for _, o := range listaOdontologos {
		if o.Matricula == ta.OdontologoMatricula {
				odontologo.Id = o.Id
				odontologo.Matricula = o.Matricula
				odontologo.Nombre = o.Nombre
				odontologo.Apellido =o.Apellido
			}
		}

	var newTurno = domain.Turno{
		Paciente: paciente,
		Odontologo: odontologo,
		FechaHora: ta.FechaHora,
		Descripcion: ta.Descripcion,
	}
	err = r.storage.PostTurno(newTurno)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllByDni(dni int) ([]domain.Turno, error) {
	return r.storage.GetTurnosByDni(dni)
}

package store

import "Final/internal/domain"

type StoreInterface interface {
	GetAllPacientes() ([]domain.Paciente, error)
	GetAllTurnos() ([]domain.Turno, error)
	GetAllOdontologos() ([]domain.Odontologo, error)

	PostOdontologo(domain.Odontologo) error
	PostPaciente(domain.Paciente) error
	PostTurno(domain.Turno) error

	GetOdontologoById(id int) (domain.Odontologo, error)
	GetPacienteById(id int) (domain.Paciente, error)
	GetTurnoById(id int) (domain.Turno, error)

	UpdateOdontologo(id int, o domain.Odontologo) error
	UpdatePaciente(id int, p domain.Paciente) error
	UpdateTurno(id int, t domain.Turno) error

	DeleteOdontologo(id int) error
	DeletePaciente(id int) error
}

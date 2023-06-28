package store

import "Final/internal/domain"

type StoreInterface interface {
	GetAllPacientes() ([]domain.Paciente, error)
	GetAllTurnos() ([]domain.Turno, error)

	GetAllOdontologos() ([]domain.Odontologo, error)

	PostOdontologo(domain.Odontologo) error

	GetOdontologoById(id int) (domain.Odontologo, error)

	UpdateOdontologo(id int, o domain.Odontologo) error

	DeleteOdontologo(id int) error
}

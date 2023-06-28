package store

import (
	"Final/internal/domain"
	"database/sql"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(database *sql.DB) StoreInterface {
	return &sqlStore{
		db: database,
	}
}

func (db *sqlStore) GetAllOdontologos() ([]domain.Odontologo, error) {
	var o domain.Odontologo
	var listaO []domain.Odontologo
	query := "SELECT * FROM odontologo;"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Nombre, &o.Apellido, &o.Matricula)
		if err != nil {
			return nil, err
		} else {
			listaO = append(listaO, o)
		}

	}
	return listaO, nil
}

func (db *sqlStore) GetAllPacientes() ([]domain.Paciente, error) {
	var p domain.Paciente
	var listaP []domain.Paciente
	query := "SELECT * FROM paciente;"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Nombre, &p.Apellido, &p.Domicilio, &p.Dni, &p.FechaAlta)
		if err != nil {
			return nil, err
		} else {
			listaP = append(listaP, p)
		}

	}
	return listaP, nil
}

func (db *sqlStore) GetAllTurnos() ([]domain.Turno, error) {
	var t domain.Turno
	var listaT []domain.Turno
	query := "SELECT * FROM turno;"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Paciente, &t.Odontologo, &t.FechaHora, &t.Descripcion)
		if err != nil {
			return nil, err
		} else {
			listaT = append(listaT, t)
		}

	}
	return listaT, nil
}

func (db *sqlStore) GetOdontologoById(id int) (domain.Odontologo, error) {
	var odontologo domain.Odontologo
	rows, err := db.db.Query("SELECT * FROM odontologo WHERE id = ?", id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	for rows.Next() {
		err := rows.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
		if err != nil {
			return domain.Odontologo{}, err
		}

	}
	return odontologo, nil
}

func (db *sqlStore) UpdateOdontologo(id int, o domain.Odontologo) error{
	query:= "UPDATE odontologo SET nombre = ?, apellido =?, matricula=? WHERE id=?"
	stmt,err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	resultado, err := stmt.Exec(o.Nombre, o.Apellido, o.Matricula,id)
	if err != nil {
		return err
	}
	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (db *sqlStore) PostOdontologo(o domain.Odontologo) error {
	query := "INSERT INTO odontologo (nombre, apellido, matricula) VALUES(?,?,?)"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(o.Nombre, o.Apellido, o.Matricula)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (db *sqlStore) DeleteOdontologo(id int) error {

	query := "DELETE FROM products WHERE id = ?"
	_, err := db.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
	
}

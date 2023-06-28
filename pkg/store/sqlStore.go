package store

import (
	"Final/internal/domain"
	"database/sql"
	"log"
)

type sqlStore struct {
	DB *sql.DB
}

func NewSqlStore(database *sql.DB) StoreInterface {
	return &sqlStore{
		DB: database,
	}
}

func (db *sqlStore) GetAllOdontologos() ([]domain.Odontologo, error) {
	var o domain.Odontologo
	var listaO []domain.Odontologo
	query := "SELECT * FROM odontologo;"
	rows, err := db.DB.Query(query)
	//FIXME VER ESTE ERROR
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Nombre, &o.Apellido, &o.Matricula)
		//FIXME VER ESTE ERROR
		if err != nil {
			log.Println(err.Error())
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
	rows, err := db.DB.Query(query)
	//FIXME VER ESTE ERROR
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Nombre, &p.Apellido, &p.Domicilio, &p.Dni, &p.FechaAlta)
		//FIXME VER ESTE ERROR
		if err != nil {
			log.Println(err.Error())
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
	rows, err := db.DB.Query(query)
	//FIXME VER ESTE ERROR
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Paciente, &t.Odontologo, &t.FechaHora, &t.Descripcion)
		//FIXME VER ESTE ERROR
		if err != nil {
			log.Println(err.Error())
			return nil, err
		} else {
			listaT = append(listaT, t)
		}

	}
	return listaT, nil
}

/*func (s *sqlStore) GetAll() ([]domain.Product, error) {

	var p domain.Product
	var products []domain.Product
	query := "SELECT * FROM products;"
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		} else {
			products = append(products, p)
		}
	}
	return products, nil
}*/

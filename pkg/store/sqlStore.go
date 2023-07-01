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
	query := `select t.id, 
	p.id, 
	p.nombre, 
	p.apellido, 
	p.domicilio, 
	p.dni, 
	p.fecha_alta, 
	o.id, 
	o.apellido, 
	o.nombre, 
	o.matricula, 
	t.fecha_hora, 
	t.descripcion  
	from turno as t 
	inner join paciente as p 
	on p.id = t.id_paciente 
	inner join odontologo as o 
	on o.id = t.id_odontologo;`
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&t.Id,
			&t.Paciente.Id,
			&t.Paciente.Nombre,
			&t.Paciente.Apellido,
			&t.Paciente.Domicilio,
			&t.Paciente.Dni,
			&t.Paciente.FechaAlta,
			&t.Odontologo.Id,
			&t.Odontologo.Apellido,
			&t.Odontologo.Nombre,
			&t.Odontologo.Matricula,
			&t.FechaHora,
			&t.Descripcion)
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

	//FIXME LIKE PACIENTE
	// if err := rows.Err(); err != nil {
	// 	return domain.Odontologo{}, err

	// }
	for rows.Next() {
		err := rows.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
		if err != nil {
			return domain.Odontologo{}, err
		}

	}
	return odontologo, nil
}

func (db *sqlStore) UpdateOdontologo(id int, o domain.Odontologo) error {
	query := "UPDATE odontologo SET nombre = ?, apellido =?, matricula=? WHERE id=?"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	resultado, err := stmt.Exec(o.Nombre, o.Apellido, o.Matricula, id)
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

	query := "DELETE FROM odontologo WHERE id = ?"
	_, err := db.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}

func (db *sqlStore) PostPaciente(p domain.Paciente) error {
	query := "INSERT INTO paciente (nombre, apellido, dni, domicilio, fecha_alta) VALUES(?,?,?,?,?)"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(p.Nombre, p.Apellido, p.Dni, p.Domicilio, p.FechaAlta)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (db *sqlStore) GetPacienteById(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	rows := db.db.QueryRow("SELECT * FROM paciente WHERE id = ?", id)
	// if err != nil {
	// 	return domain.Paciente{}, err
	// }
	// for rows.Next() {
	err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Dni, &paciente.Domicilio, &paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
		// }

	}
	return paciente, nil
}

func (db *sqlStore) UpdatePaciente(id int, p domain.Paciente) error {
	query := "UPDATE paciente SET nombre = ?, apellido =?, dni=?, domicilio=?, fecha_alta =? WHERE id=?"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	resultado, err := stmt.Exec(p.Nombre, p.Apellido, p.Dni, p.Domicilio, p.FechaAlta, id)
	if err != nil {
		return err
	}
	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (db *sqlStore) DeletePaciente(id int) error {

	query := "DELETE FROM paciente WHERE id = ?"
	_, err := db.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}

func (db *sqlStore) PostTurno(t domain.Turno) error {
	query := "INSERT INTO turno (id_paciente, id_odontologo, fecha_hora, descripcion) VALUES(?,?,?,?)"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(t.Paciente.Id, t.Odontologo.Id, t.FechaHora, t.Descripcion)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (db *sqlStore) GetTurnoById(id int) (domain.Turno, error) {
	var t domain.Turno
	rows := db.db.QueryRow(
		`select t.id, 
		p.id, 
		p.nombre, 
		p.apellido, 
		p.domicilio, 
		p.dni, 
		p.fecha_alta, 
		o.id, 
		o.apellido, 
		o.nombre, 
		o.matricula, 
		t.fecha_hora, 
		t.descripcion  
		from turno as t 
		inner join paciente as p 
		on p.id = t.id_paciente 
		inner join odontologo as o 
		on o.id = t.id_odontologo
		where t.id = ?;`, id)
	
		err:= rows.Scan(&t.Id,
				&t.Paciente.Id,
				&t.Paciente.Nombre,
				&t.Paciente.Apellido,
				&t.Paciente.Domicilio,
				&t.Paciente.Dni,
				&t.Paciente.FechaAlta,
				&t.Odontologo.Id,
				&t.Odontologo.Apellido,
				&t.Odontologo.Nombre,
				&t.Odontologo.Matricula,
				&t.FechaHora,
				&t.Descripcion)
			if err != nil {
				return domain.Turno{}, err
			}
	return t, nil
}

func (db *sqlStore) UpdateTurno(id int, t domain.Turno) error {
	query := "UPDATE turno SET id_paciente = ?, id_odontologo =?, fecha_hora=?, descripcion=? WHERE id=?"
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	resultado, err := stmt.Exec(t.Paciente.Id, t.Odontologo.Id, t.FechaHora,t.Descripcion, id)
	if err != nil {
		return err
	}
	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}


func (db *sqlStore) DeleteTurno(id int) error {

	query := "DELETE FROM turno WHERE id = ?"
	_, err := db.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}


func (db *sqlStore) PostTurnoDNIMat(ta domain.TurnoAuxiliar) (error) {
	var paciente domain.Paciente
	rows := db.db.QueryRow("SELECT * FROM paciente WHERE id = ?", ta.PacienteId)
	var odontologo domain.Odontologo
	rows2 := db.db.QueryRow("SELECT * FROM odontologo WHERE id = ?", ta.OdontologoId)

	err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Dni, &paciente.Domicilio, &paciente.FechaAlta)
	if err != nil {
		return err
	}
	err2 := rows2.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
	if err2 != nil {
		return nil
	}
	query := "INSERT INTO turno (id_paciente, id_odontologo, fecha_hora, descripcion) VALUES(?,?,?,?)"
	stmt, err3 := db.db.Prepare(query)
	if err3 != nil {
		return err3
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(paciente.Id, odontologo.Id, ta.FechaHora, ta.Descripcion)
	if err != nil {
		return err
	}

	_, err = resultado.RowsAffected()
	if err != nil {
		return err
	}
	return nil
	

}
func (db *sqlStore) GetTurnosByDni(dni int) ([]domain.Turno, error) {
	var t domain.Turno
	var listaT []domain.Turno
	query := `select t.id, 
	p.id, 
	p.nombre, 
	p.apellido, 
	p.domicilio, 
	p.dni, 
	p.fecha_alta, 
	o.id, 
	o.apellido, 
	o.nombre, 
	o.matricula, 
	t.fecha_hora, 
	t.descripcion  
	from turno as t 
	inner join paciente as p 
	on p.id = t.id_paciente 
	inner join odontologo as o 
	on o.id = t.id_odontologo
	where p.dni =? ;`
	rows := db.db.QueryRow(query,dni)
	// if err != nil {
	// 	return nil, err
	// }
// 
// 	for rows.Next() {
		err := rows.Scan(&t.Id,
			&t.Paciente.Id,
			&t.Paciente.Nombre,
			&t.Paciente.Apellido,
			&t.Paciente.Domicilio,
			&t.Paciente.Dni,
			&t.Paciente.FechaAlta,
			&t.Odontologo.Id,
			&t.Odontologo.Apellido,
			&t.Odontologo.Nombre,
			&t.Odontologo.Matricula,
			&t.FechaHora,
			&t.Descripcion)
		if err != nil {
			return nil, err
		} else {
			listaT = append(listaT, t)
		}

	// }
	return listaT, nil
}
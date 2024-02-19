package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"pat-api/internal/models"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newUserPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *User) error {
	const psqlInsert = `INSERT INTO auth.user (name, lastname, email, cellphone, password, age, city, department) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		m.Lastname,
		m.Email,
		m.Cellphone,
		m.Password,
		m.Age,
		m.City,
		m.Department,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *User) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE auth.user SET name = :name, lastname = :lastname, email = :email, cellphone = :cellphone, password = :password, age = :age, city = :city, department = :department, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id int64) error {
	const psqlDelete = `DELETE FROM auth.user WHERE id = :id `
	m := User{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id int64) (*User, error) {
	const psqlGetByID = `SELECT id , name, lastname, email, cellphone, password, age, city, department, created_at, updated_at FROM auth.user WHERE id = $1 `
	mdl := User{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*User, error) {
	var ms []*User
	const psqlGetAll = ` SELECT id , name, lastname, email, cellphone, password, age, city, department, created_at, updated_at FROM auth.user `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) GetByEmail(email string) (*User, error) {
	const psqlGetByEmail = `SELECT id , name, lastname, email, cellphone, password, age, city, department, created_at, updated_at FROM auth.user WHERE email = $1 limit 1;`
	mdl := User{}
	err := s.DB.Get(&mdl, psqlGetByEmail, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil

}

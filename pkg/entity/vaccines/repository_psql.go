package vaccines

import (
	"database/sql"
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

func newVaccinesPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Vaccines) error {
	const psqlInsert = `INSERT INTO entity.vaccines ("name", veterinary, doctor, pet) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		m.Veterinary,
		m.Doctor,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Vaccines) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.vaccines SET "name" = :name, veterinary = :veterinary, doctor = :doctor, pet = :pet, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM entity.vaccines WHERE id = :id `
	m := Vaccines{ID: id}
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
func (s *psql) getByID(id int64) (*Vaccines, error) {
	const psqlGetByID = `SELECT id , "name", veterinary, doctor, pet, created_at, updated_at FROM entity.vaccines WHERE id = $1 `
	mdl := Vaccines{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll(pet int64) ([]*Vaccines, error) {
	var ms []*Vaccines
	const psqlGetAll = ` SELECT id , "name", veterinary, doctor, pet, created_at, updated_at FROM entity.vaccines where pet = $1;`

	err := s.DB.Select(&ms, psqlGetAll, pet)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

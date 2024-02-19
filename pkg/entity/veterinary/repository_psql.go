package veterinary

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

func newVeterinaryPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Veterinary) error {
	const psqlInsert = `INSERT INTO entity.veterinary (name, description, email, address, cellphone, "user", web_page) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		m.Description,
		m.Email,
		m.Address,
		m.Cellphone,
		m.User,
		m.WebPage,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Veterinary) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.veterinary SET name = :name, description = :description, email = :email, address = :address, cellphone = :cellphone, "user" = :user, web_page = :web_page, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM entity.veterinary WHERE id = :id `
	m := Veterinary{ID: id}
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
func (s *psql) getByID(id int64) (*Veterinary, error) {
	const psqlGetByID = `SELECT id , name, description, email, address, cellphone, "user", web_page, created_at, updated_at FROM entity.veterinary WHERE id = $1 `
	mdl := Veterinary{}
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
func (s *psql) getAll() ([]*Veterinary, error) {
	var ms []*Veterinary
	const psqlGetAll = ` SELECT id , name, description, email, address, cellphone, "user", web_page, created_at, updated_at FROM entity.veterinary `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

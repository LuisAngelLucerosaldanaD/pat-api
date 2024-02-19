package pet

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

func newPetPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Pet) error {
	const psqlInsert = `INSERT INTO entity.pet (name, category, age, weight, sexo, "user", type) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		m.Category,
		m.Age,
		m.Weight,
		m.Sex,
		m.User,
		m.TypePet,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Pet) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.pet SET name = :name, category = :category, age = :age, weight = :weight, sexo = :sexo, "user" = :user, type = :type, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM entity.pet WHERE id = :id `
	m := Pet{ID: id}
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
func (s *psql) getByID(id int64) (*Pet, error) {
	const psqlGetByID = `SELECT id , name, category, age, weight, sexo, "user", type, created_at, updated_at FROM entity.pet WHERE id = $1 `
	mdl := Pet{}
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
func (s *psql) getAll(id int64) ([]*Pet, error) {
	var ms []*Pet
	const psqlGetAll = ` SELECT id , name, category, age, weight, sexo, "user", type, created_at, updated_at FROM entity.pet where "user" = $1`

	err := s.DB.Select(&ms, psqlGetAll, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

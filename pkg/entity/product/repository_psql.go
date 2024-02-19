package product

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

func newProductPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Product) error {
	const psqlInsert = `INSERT INTO entity.product (name, description, stock, veterinary, category, price, type_product) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	stmt, err := s.DB.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		m.Description,
		m.Stock,
		m.Veterinary,
		m.Category,
		m.Price,
		m.TypeProduct,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Product) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE entity.product SET name = :name, description = :description, stock = :stock, veterinary = :veterinary, category = :category, price = :price, type_product = :type_product, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM entity.product WHERE id = :id `
	m := Product{ID: id}
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
func (s *psql) getByID(id int64) (*Product, error) {
	const psqlGetByID = `SELECT id , name, description, stock, veterinary, category, price, type_product, created_at, updated_at FROM entity.product WHERE id = $1 `
	mdl := Product{}
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
func (s *psql) getAll() ([]*Product, error) {
	var ms []*Product
	const psqlGetAll = ` SELECT id , name, description, stock, veterinary, category, price, type_product, created_at, updated_at FROM entity.product `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

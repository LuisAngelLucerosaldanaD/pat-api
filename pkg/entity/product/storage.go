package product

import (
	"github.com/jmoiron/sqlx"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesProductRepository interface {
	create(m *Product) error
	update(m *Product) error
	delete(id int64) error
	getByID(id int64) (*Product, error)
	getAll() ([]*Product, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesProductRepository {
	var s ServicesProductRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newProductPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

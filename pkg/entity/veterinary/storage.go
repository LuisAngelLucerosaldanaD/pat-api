package veterinary

import (
	"github.com/jmoiron/sqlx"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesVeterinaryRepository interface {
	create(m *Veterinary) error
	update(m *Veterinary) error
	delete(id int64) error
	getByID(id int64) (*Veterinary, error)
	getAll() ([]*Veterinary, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesVeterinaryRepository {
	var s ServicesVeterinaryRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newVeterinaryPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

package vaccines

import (
	"github.com/jmoiron/sqlx"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesVaccinesRepository interface {
	create(m *Vaccines) error
	update(m *Vaccines) error
	delete(id int64) error
	getByID(id int64) (*Vaccines, error)
	getAll(pet int64) ([]*Vaccines, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesVaccinesRepository {
	var s ServicesVaccinesRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newVaccinesPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

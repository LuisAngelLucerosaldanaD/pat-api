package pet

import (
	"github.com/jmoiron/sqlx"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesPetRepository interface {
	create(m *Pet) error
	update(m *Pet) error
	delete(id int64) error
	getByID(id int64) (*Pet, error)
	getAll(id int64) ([]*Pet, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesPetRepository {
	var s ServicesPetRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newPetPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

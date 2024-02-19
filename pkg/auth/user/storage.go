package user

import (
	"github.com/jmoiron/sqlx"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesUserRepository interface {
	create(m *User) error
	update(m *User) error
	delete(id int64) error
	getByID(id int64) (*User, error)
	getAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserRepository {
	var s ServicesUserRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newUserPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

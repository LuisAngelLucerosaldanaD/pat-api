package auth

import (
	"github.com/jmoiron/sqlx"
	"pat-api/internal/models"
	"pat-api/pkg/auth/user"
)

type Server struct {
	SrvUser user.PortsServerUser
}

func NewServerAuth(db *sqlx.DB, usr *models.User, txID string) *Server {

	repoUser := user.FactoryStorage(db, usr, txID)
	srvUser := user.NewUserService(repoUser, usr, txID)

	return &Server{
		SrvUser: srvUser,
	}
}

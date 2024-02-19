package entity

import (
	"github.com/jmoiron/sqlx"
	"pat-api/internal/models"
	"pat-api/pkg/entity/pet"
	"pat-api/pkg/entity/product"
	"pat-api/pkg/entity/vaccines"
	"pat-api/pkg/entity/veterinary"
)

type Server struct {
	SrvPet        pet.PortsServerPet
	SrvVeterinary veterinary.PortsServerVeterinary
	SrvProduct    product.PortsServerProduct
	SrvVaccines   vaccines.PortsServerVaccines
}

func NewServerEntity(db *sqlx.DB, usr *models.User, txID string) *Server {

	repoPet := pet.FactoryStorage(db, usr, txID)
	srvPet := pet.NewPetService(repoPet, usr, txID)

	repoVeterinary := veterinary.FactoryStorage(db, usr, txID)
	srvVeterinary := veterinary.NewVeterinaryService(repoVeterinary, usr, txID)

	repoProduct := product.FactoryStorage(db, usr, txID)
	srvProduct := product.NewProductService(repoProduct, usr, txID)

	repoVaccines := vaccines.FactoryStorage(db, usr, txID)
	srvVaccines := vaccines.NewVaccinesService(repoVaccines, usr, txID)

	return &Server{
		SrvPet:        srvPet,
		SrvVeterinary: srvVeterinary,
		SrvProduct:    srvProduct,
		SrvVaccines:   srvVaccines,
	}
}

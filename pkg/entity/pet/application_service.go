package pet

import (
	"fmt"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

type PortsServerPet interface {
	CreatePet(name string, category string, age int, weight float64, sexo string, user int64, typePet string) (*Pet, int, error)
	UpdatePet(id int64, name string, category string, age int, weight float64, sexo string, user int64, typePet string) (*Pet, int, error)
	DeletePet(id int64) (int, error)
	GetPetByID(id int64) (*Pet, int, error)
	GetAllPet(id int64) ([]*Pet, error)
}

type service struct {
	repository ServicesPetRepository
	user       *models.User
	txID       string
}

func NewPetService(repository ServicesPetRepository, user *models.User, TxID string) PortsServerPet {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePet(name string, category string, age int, weight float64, sexo string, user int64, typePet string) (*Pet, int, error) {
	m := NewCreatePet(name, category, age, weight, sexo, user, typePet)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Pet :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePet(id int64, name string, category string, age int, weight float64, sexo string, user int64, typePet string) (*Pet, int, error) {
	m := NewPet(id, name, category, age, weight, sexo, user, typePet)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Pet :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePet(id int64) (int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetPetByID(id int64) (*Pet, int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllPet(id int64) ([]*Pet, error) {
	return s.repository.getAll(id)
}

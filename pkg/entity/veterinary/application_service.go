package veterinary

import (
	"fmt"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

type PortsServerVeterinary interface {
	CreateVeterinary(name string, description string, email string, address string, cellphone string, user int64, webPage string) (*Veterinary, int, error)
	UpdateVeterinary(id int64, name string, description string, email string, address string, cellphone string, user int64, webPage string) (*Veterinary, int, error)
	DeleteVeterinary(id int64) (int, error)
	GetVeterinaryByID(id int64) (*Veterinary, int, error)
	GetAllVeterinary() ([]*Veterinary, error)
}

type service struct {
	repository ServicesVeterinaryRepository
	user       *models.User
	txID       string
}

func NewVeterinaryService(repository ServicesVeterinaryRepository, user *models.User, TxID string) PortsServerVeterinary {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateVeterinary(name string, description string, email string, address string, cellphone string, user int64, webPage string) (*Veterinary, int, error) {
	m := NewCreateVeterinary(name, description, email, address, cellphone, user, webPage)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Veterinary :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateVeterinary(id int64, name string, description string, email string, address string, cellphone string, user int64, webPage string) (*Veterinary, int, error) {
	m := NewVeterinary(id, name, description, email, address, cellphone, user, webPage)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Veterinary :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteVeterinary(id int64) (int, error) {
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

func (s *service) GetVeterinaryByID(id int64) (*Veterinary, int, error) {
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

func (s *service) GetAllVeterinary() ([]*Veterinary, error) {
	return s.repository.getAll()
}

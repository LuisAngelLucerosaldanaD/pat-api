package vaccines

import (
	"fmt"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

type PortsServerVaccines interface {
	CreateVaccines(name string, veterinary int64, doctor string, pet int64) (*Vaccines, int, error)
	UpdateVaccines(id int64, name string, veterinary int64, doctor string, pet int64) (*Vaccines, int, error)
	DeleteVaccines(id int64) (int, error)
	GetVaccinesByID(id int64) (*Vaccines, int, error)
	GetAllVaccines(pet int64) ([]*Vaccines, error)
}

type service struct {
	repository ServicesVaccinesRepository
	user       *models.User
	txID       string
}

func NewVaccinesService(repository ServicesVaccinesRepository, user *models.User, TxID string) PortsServerVaccines {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateVaccines(name string, veterinary int64, doctor string, pet int64) (*Vaccines, int, error) {
	m := NewCreateVaccines(name, veterinary, doctor, pet)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Vaccines :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateVaccines(id int64, name string, veterinary int64, doctor string, pet int64) (*Vaccines, int, error) {
	m := NewVaccines(id, name, veterinary, doctor, pet)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Vaccines :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteVaccines(id int64) (int, error) {
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

func (s *service) GetVaccinesByID(id int64) (*Vaccines, int, error) {
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

func (s *service) GetAllVaccines(pet int64) ([]*Vaccines, error) {
	return s.repository.getAll(pet)
}

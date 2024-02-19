package user

import (
	"fmt"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

type PortsServerUser interface {
	CreateUser(name string, lastname string, email string, cellphone string, password string, age int, city string, department string) (*User, int, error)
	UpdateUser(id int64, name string, lastname string, email string, cellphone string, password string, age int, city string, department string) (*User, int, error)
	DeleteUser(id int64) (int, error)
	GetUserByID(id int64) (*User, int, error)
	GetAllUser() ([]*User, error)
	GetUserByEmail(email string) (*User, int, error)
}

type service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
}

func NewUserService(repository ServicesUserRepository, user *models.User, TxID string) PortsServerUser {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUser(name string, lastname string, email string, cellphone string, password string, age int, city string, department string) (*User, int, error) {
	m := NewCreateUser(name, lastname, email, cellphone, password, age, city, department)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create User :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUser(id int64, name string, lastname string, email string, cellphone string, password string, age int, city string, department string) (*User, int, error) {
	m := NewUser(id, name, lastname, email, cellphone, password, age, city, department)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update User :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUser(id int64) (int, error) {
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

func (s *service) GetUserByID(id int64) (*User, int, error) {
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

func (s *service) GetAllUser() ([]*User, error) {
	return s.repository.getAll()
}

func (s *service) GetUserByEmail(email string) (*User, int, error) {
	m, err := s.repository.GetByEmail(email)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t GetByEmail row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

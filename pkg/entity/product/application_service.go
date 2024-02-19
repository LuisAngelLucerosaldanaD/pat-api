package product

import (
	"fmt"

	"pat-api/internal/logger"
	"pat-api/internal/models"
)

type PortsServerProduct interface {
	CreateProduct(name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) (*Product, int, error)
	UpdateProduct(id int64, name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) (*Product, int, error)
	DeleteProduct(id int64) (int, error)
	GetProductByID(id int64) (*Product, int, error)
	GetAllProduct() ([]*Product, error)
}

type service struct {
	repository ServicesProductRepository
	user       *models.User
	txID       string
}

func NewProductService(repository ServicesProductRepository, user *models.User, TxID string) PortsServerProduct {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateProduct(name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) (*Product, int, error) {
	m := NewCreateProduct(name, description, stock, veterinary, category, price, typeProduct)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Product :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateProduct(id int64, name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) (*Product, int, error) {
	m := NewProduct(id, name, description, stock, veterinary, category, price, typeProduct)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Product :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteProduct(id int64) (int, error) {
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

func (s *service) GetProductByID(id int64) (*Product, int, error) {
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

func (s *service) GetAllProduct() ([]*Product, error) {
	return s.repository.getAll()
}

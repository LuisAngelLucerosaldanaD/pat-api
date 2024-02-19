package product

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Product  Model struct Product
type Product struct {
	ID          int64     `json:"id" db:"id" valid:"-"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Stock       int       `json:"stock" db:"stock" valid:"required"`
	Veterinary  int64     `json:"veterinary" db:"veterinary" valid:"required"`
	Category    string    `json:"category" db:"category" valid:"required"`
	Price       float64   `json:"price" db:"price" valid:"required"`
	TypeProduct string    `json:"type_product" db:"type_product" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(id int64, name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Stock:       stock,
		Veterinary:  veterinary,
		Category:    category,
		Price:       price,
		TypeProduct: typeProduct,
	}
}

func NewCreateProduct(name string, description string, stock int, veterinary int64, category string, price float64, typeProduct string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Stock:       stock,
		Veterinary:  veterinary,
		Category:    category,
		Price:       price,
		TypeProduct: typeProduct,
	}
}

func (m *Product) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

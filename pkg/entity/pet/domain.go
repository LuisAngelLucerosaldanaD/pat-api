package pet

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Pet  Model struct Pet
type Pet struct {
	ID        int64     `json:"id" db:"id" valid:"-"`
	Name      string    `json:"name" db:"name" valid:"required"`
	Category  string    `json:"category" db:"category" valid:"required"`
	Age       int       `json:"age" db:"age" valid:"required"`
	Weight    float64   `json:"weight" db:"weight" valid:"required"`
	Sex       string    `json:"sexo" db:"sexo" valid:"required"`
	User      int64     `json:"user" db:"user" valid:"required"`
	TypePet   string    `json:"type" db:"type" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewPet(id int64, name string, category string, age int, weight float64, sexo string, user int64, typePet string) *Pet {
	return &Pet{
		ID:       id,
		Name:     name,
		Category: category,
		Age:      age,
		Weight:   weight,
		Sex:      sexo,
		User:     user,
		TypePet:  typePet,
	}
}

func NewCreatePet(name string, category string, age int, weight float64, sexo string, user int64, typePet string) *Pet {
	return &Pet{
		Name:     name,
		Category: category,
		Age:      age,
		Weight:   weight,
		Sex:      sexo,
		User:     user,
		TypePet:  typePet,
	}
}

func (m *Pet) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

package vaccines

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Vaccines  Model struct Vaccines
type Vaccines struct {
	ID         int64     `json:"id" db:"id" valid:"-"`
	Name       string    `json:"name" db:"name" valid:"required"`
	Veterinary int64     `json:"veterinary" db:"veterinary" valid:"required"`
	Doctor     string    `json:"doctor" db:"doctor" valid:"required"`
	Pet        int64     `json:"pet" db:"pet" valid:"-"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func NewVaccines(id int64, name string, veterinary int64, doctor string, pet int64) *Vaccines {
	return &Vaccines{
		ID:         id,
		Name:       name,
		Veterinary: veterinary,
		Doctor:     doctor,
		Pet:        pet,
	}
}

func NewCreateVaccines(name string, veterinary int64, doctor string, pet int64) *Vaccines {
	return &Vaccines{
		Name:       name,
		Veterinary: veterinary,
		Doctor:     doctor,
		Pet:        pet,
	}
}

func (m *Vaccines) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

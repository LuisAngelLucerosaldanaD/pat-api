package veterinary

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Veterinary  Model struct Veterinary
type Veterinary struct {
	ID          int64     `json:"id" db:"id" valid:"-"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Email       string    `json:"email" db:"email" valid:"required"`
	Address     string    `json:"address" db:"address" valid:"required"`
	Cellphone   string    `json:"cellphone" db:"cellphone" valid:"required"`
	User        int64     `json:"user" db:"user" valid:"required"`
	WebPage     string    `json:"web_page" db:"web_page" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewVeterinary(id int64, name string, description string, email string, address string, cellphone string, user int64, webPage string) *Veterinary {
	return &Veterinary{
		ID:          id,
		Name:        name,
		Description: description,
		Email:       email,
		Address:     address,
		Cellphone:   cellphone,
		User:        user,
		WebPage:     webPage,
	}
}

func NewCreateVeterinary(name string, description string, email string, address string, cellphone string, user int64, webPage string) *Veterinary {
	return &Veterinary{
		Name:        name,
		Description: description,
		Email:       email,
		Address:     address,
		Cellphone:   cellphone,
		User:        user,
		WebPage:     webPage,
	}
}

func (m *Veterinary) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

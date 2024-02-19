package user

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// User  Model struct User
type User struct {
	ID         int64     `json:"id" db:"id" valid:"-"`
	Name       string    `json:"name" db:"name" valid:"required"`
	Lastname   string    `json:"lastname" db:"lastname" valid:"required"`
	Email      string    `json:"email" db:"email" valid:"required"`
	Cellphone  string    `json:"cellphone" db:"cellphone" valid:"required"`
	Password   string    `json:"password" db:"password" valid:"required"`
	Age        int       `json:"age" db:"age" valid:"required"`
	City       string    `json:"city" db:"city" valid:"required"`
	Department string    `json:"department" db:"department" valid:"required"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(id int64, name string, lastname string, email string, cellphone string, password string, age int, city string, department string) *User {
	return &User{
		ID:         id,
		Name:       name,
		Lastname:   lastname,
		Email:      email,
		Cellphone:  cellphone,
		Password:   password,
		Age:        age,
		City:       city,
		Department: department,
	}
}

func NewCreateUser(name string, lastname string, email string, cellphone string, password string, age int, city string, department string) *User {
	return &User{
		Name:       name,
		Lastname:   lastname,
		Email:      email,
		Cellphone:  cellphone,
		Password:   password,
		Age:        age,
		City:       city,
		Department: department,
	}
}

func (m *User) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

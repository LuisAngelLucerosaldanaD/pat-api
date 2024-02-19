package models

import "time"

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	Msg   string      `json:"msg"`
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Cellphone  string    `json:"cellphone"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	City       string    `json:"city"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

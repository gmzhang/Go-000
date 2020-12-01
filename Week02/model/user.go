package model

type User struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  uint   `json:"age" db:"age"`
}

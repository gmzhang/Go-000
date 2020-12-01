package dao

import "github.com/gmzhang/Go-000/Week02/model"

type Dao interface {
	GetUserById(id uint) (user model.User, err error)
}

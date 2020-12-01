package service

import "github.com/gmzhang/Go-000/Week02/model"

type Service interface {
	GetUserById(id uint) (user model.User, err error)
}

package service

import (
	"github.com/gmzhang/Go-000/Week02/dao"
	"github.com/jmoiron/sqlx"
	"github.com/gmzhang/Go-000/Week02/model"
)

type implService struct {
	dao dao.Dao
}

func NewService(db *sqlx.DB) Service {
	return &implService{dao: dao.NewDao(db)}
}

func (i *implService) GetUserById(id uint) (user model.User, err error) {
	return i.dao.GetUserById(id)
}

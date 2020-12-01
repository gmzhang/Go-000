package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/gmzhang/Go-000/Week02/model"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/gmzhang/Go-000/Week02/errs"
)

type implDao struct {
	db *sqlx.DB
}

func NewDao(db *sqlx.DB) Dao {
	return &implDao{db: db}
}

func (i *implDao) GetUserById(id uint) (user model.User, err error) {
	query := "SELECT `id`,`name`,`age` FROM `users` WHERE id=?;"
	err = i.db.Get(&user, query, id)

	if err == sql.ErrNoRows {
		return user, errors.Wrapf(errs.ErrUserNotFound, "user id: %d", id)
	}

	if err != nil {
		return user, errors.Wrapf(err, "get user by id error, user id: %d", id)
	}

	return
}

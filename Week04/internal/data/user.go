package data

import "github.com/gmzhang/Go-000/Week04/internal/biz"

var _ biz.UserRepo = new(userRepo)

type userRepo struct {
}

func NewUserReop() biz.UserRepo {
	return &userRepo{}
}

func (u *userRepo) GetUser(id int32) biz.User {

	return biz.User{ID: 1, Name: "foo", Avatar: "https://bar.com/foo.png"}

}

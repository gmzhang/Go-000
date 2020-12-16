//+build wireinject

package main

import (
	"github.com/gmzhang/Go-000/Week04/internal/biz"
	"github.com/google/wire"
	"github.com/gmzhang/Go-000/Week04/internal/data"
)

func InitUsercase() *biz.UserUsecase {
	wire.Build(biz.NewUserUsecase, data.NewUserReop())

	return &biz.UserUsecase{}
}

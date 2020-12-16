//Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/gmzhang/Go-000/Week04/internal/biz"
	"github.com/gmzhang/Go-000/Week04/internal/data"
)

// Injectors from wire.go:

func InitUserUsecase() *biz.UserUsecase {
	userRepo := data.NewUserReop()
	userUsecase := biz.NewUserUsecase(userRepo)
	return userUsecase
}
package service

import (
	"github.com/gmzhang/Go-000/Week04/internal/biz"

	v1 "github.com/gmzhang/Go-000/Week04/api/user/v1"

	"context"
)

type UserService struct {
	u *biz.UserUsecase
}

func NewUserService(u *biz.UserUsecase) v1.UserServer {
	return &UserService{u: u}
}

func (s *UserService) GetUser(ctx context.Context, r *v1.GetUserRequest) (*v1.UserResponse, error) {

	user := s.u.GetUser(r.Id)

	return &v1.UserResponse{Id: user.ID, Name: user.Name, Avatar: user.Avatar}, nil
}

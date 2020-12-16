## 学习笔记

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。


### 文件结构

```
├── README.md
├── api
│   └── user
│       └── v1
│           ├── user.pb.go
│           └── user.proto
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
└── internal
    ├── biz
    │   └── user.go
    ├── data
    │   └── user.go
    ├── pkg
    │   └── grpc
    │       └── service.go
    └── service
        └── user.go

```

### 数据层

```go
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

```

### 业务层

```go

package biz

type User struct {
	ID     int32
	Name   string
	Avatar string
}

type UserRepo interface {
	GetUser(id int32) User
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUser(id int32) User {
	return u.repo.GetUser(id)
}

```

### API 层

```go

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

```

### main 启动

```go

package main

import (
	"github.com/gmzhang/Go-000/Week04/internal/service"
	"github.com/gmzhang/Go-000/Week04/internal/pkg/grpc"
	pb "github.com/gmzhang/Go-000/Week04/api/user/v1"
	"context"
	"fmt"
)

func main() {

	address := ":8081"

	fmt.Println(address)

	us := InitUserUsecase()
	server := service.NewUserService(us)

	grpcSvc := grpc.NewServer(address)
	pb.RegisterUserServer(grpcSvc.Server, server)

	grpcSvc.Start(context.TODO())

}

```

### proto 定义

```text

syntax = "proto3";

package user.v1;


service User {
  rpc GetUser (GetUserRequest) returns (UserResponse) {}
}

message GetUserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string avatar = 3;
}
```
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

package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gmzhang/Go-000/Week02/api"
	"github.com/gmzhang/Go-000/Week02/service"
)

var db *sqlx.DB

func main() {

	initDB()

	runHttpServer()

}

func runHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())

	api.NewApiHandler(service.NewService(db), e)

	err := e.Start(":8080")
	if err != nil {
		logrus.Fatalln(err)
	}
}

func initDB() {
	dsn := fmt.Sprintf("user:pwd@tcp(host:port)/db?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci&loc=Local")

	var err error
	db, err = sqlx.Open("mysql", dsn)

	if err != nil {
		err = db.Ping()
	}

	if err != nil {
		logrus.Fatalf("db connect error: %#v", err)
	}

	logrus.Infoln("connect db success")
}

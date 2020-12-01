package api

import (
	"github.com/gmzhang/Go-000/Week02/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gmzhang/Go-000/Week02/errs"
)

type handler struct {
	service service.Service
}

func NewApiHandler(service service.Service, e *echo.Echo) {
	handler := &handler{service: service}
	e.GET("/user/:id", handler.getUserById)
}

func (h handler) getUserById(c echo.Context) error {
	id := c.Get("id").(uint)
	user, err := h.service.GetUserById(id)
	if err != nil {
		logrus.Errorf("get user by id error: %+v", err)
	}
	result := errs.GetErrorMap(err)
	result["data"] = user
	return c.JSON(http.StatusOK, result)
}

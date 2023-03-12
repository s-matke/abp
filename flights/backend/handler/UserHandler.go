package handler

import (
	"encoding/json"
	"flight/model"
	"flight/service"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello!")
}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {

	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || user.Email == "" || user.Password == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.Create(&user)

	if err != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

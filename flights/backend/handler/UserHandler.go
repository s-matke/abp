package handler

import (
	"encoding/json"
	"flight/dto"
	"flight/model"
	"flight/service"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func (handler *UserHandler) SignIn(writer http.ResponseWriter, req *http.Request) {
	var signInDTO dto.SignInDTO
	err := json.NewDecoder(req.Body).Decode(&signInDTO)

	if err != nil || signInDTO.Email == "" || signInDTO.Password == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.UserService.FindByEmail(signInDTO.Email)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInDTO.Password))

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)

}

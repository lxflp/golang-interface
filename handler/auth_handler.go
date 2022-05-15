package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lxflp/golang-interface/models"
	"github.com/lxflp/golang-interface/storage"
	"net/http"
)

type AuthHandler struct {
	storage storage.UserStorager
}

func NewAuthHandler(storage storage.UserStorager) *AuthHandler {
	return &AuthHandler{storage: storage}
}

func (a *AuthHandler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	var user models.UserData
	err := Decode(request, &user)
	if err != nil {
		fmt.Fprintf(writer, "неверные данные")
		return
	}
	err = a.storage.Save(user)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}
	fmt.Fprintf(writer, "пользователь успешно зарегистрирован")
}
func (a *AuthHandler) LoginUser(writer http.ResponseWriter, request *http.Request) {
	var user models.UserData
	var err error
	err = Decode(request, &user)
	if err != nil {
		fmt.Fprintf(writer, "неверные данные")
		return
	}
	var user2 models.UserData
	user2, err = a.storage.Get(user.Phone)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}
	if user.Password != user2.Password {
		fmt.Fprintf(writer, "неверный пароль")
		return
	}
	fmt.Fprintf(writer, "вы успешно вошли в систему")
}

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(val); err != nil {
		return err
	}
	return nil
}

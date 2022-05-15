package main

import (
	"github.com/lxflp/golang-interface/handler"
	"github.com/lxflp/golang-interface/models"
	"github.com/lxflp/golang-interface/storage"
	"net/http"
)

func main() {
	var sstorage storage.UserStorager
	var auth *handler.AuthHandler
	//sstorage = storage.NewUserStorageSlice(make([]models.UserData, 0, 10))
	sstorage = storage.NewUserStorageMap(make(map[string]models.UserData))
	auth = handler.NewAuthHandler(sstorage)
	http.HandleFunc("/register", auth.RegisterUser)
	http.HandleFunc("/login", auth.LoginUser)
	http.ListenAndServe(":8888", nil)
}

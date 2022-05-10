package main

import (
	"fmt"
	"net/http"
)

type usersStorager interface {
	save(user userData) error
	get(phone string) (userData, error)
	delete(phone string) error
}

type userData struct {
	Phone    string
	Password string
	Name     string
}

type userStorage struct {
	Data []userData //здесь хранятся пользователи
}

//методы:
func (u userStorage) save(user userData) error {
	for _, usr := range u.Data {
		if usr.Phone == user.Phone {
			return fmt.Errorf("данный номер уже зарегестрирован")
		}
	}
	u.Data = append(u.Data, user)
	return nil
}

func (u userStorage) get(phone string) (userData, error) {
	for _, usr := range u.Data {
		if usr.Phone == phone {
			return usr, nil
		}
	}
	return userData{}, fmt.Errorf("пользователь не найден")
}

func (u userStorage) delete(phone string) error {
	for i, usr := range u.Data {
		if usr.Phone == phone {
			u.Data = append(u.Data[:i], u.Data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("пользователь не найден")
}

func main() {
	http.HandleFunc("/register", registerUser)

}
func registerUser(writer http.ResponseWriter, request *http.Request) {

}

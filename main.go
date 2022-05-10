package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type usersStorager interface {
	save(user userData) error
	get(phone string) (userData, error)
	delete(phone string) error
}

type userData struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type userStorage struct {
	Data []userData //здесь хранятся пользователи
}

//методы:
func (u *userStorage) save(user userData) error {
	for _, usr := range u.Data {
		if usr.Phone == user.Phone {
			return fmt.Errorf("данный номер уже зарегестрирован")
		}
	}
	u.Data = append(u.Data, user)
	return nil
}

func (u *userStorage) get(phone string) (userData, error) {
	for _, usr := range u.Data {
		if usr.Phone == phone {
			return usr, nil
		}
	}
	return userData{}, fmt.Errorf("пользователь не найден")
}

func (u *userStorage) delete(phone string) error {
	for i, usr := range u.Data {
		if usr.Phone == phone {
			u.Data = append(u.Data[:i], u.Data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("пользователь не найден")
}

var storage = userStorage{Data: make([]userData, 0, 10)}

func main() {
	http.HandleFunc("/register", registerUser)
	http.ListenAndServe(":8888", nil)
}
func registerUser(writer http.ResponseWriter, request *http.Request) {
	var user userData
	err := Decode(request, &user)
	if err != nil {
		fmt.Fprintf(writer, "Неверные данные")
		return
	}
	err = storage.save(user)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return

	}
	fmt.Fprintf(writer, "Пользователь успешно зарегистрирован")
}

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

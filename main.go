package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type usersStorager interface {
	save(user userData) error           //функция сохранения пользователей
	get(phone string) (userData, error) //функция получения пользователя по номеру телефона
	delete(phone string) error          //функция удаления пользователя
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
	http.HandleFunc("/login", loginUser)
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
func loginUser(writer http.ResponseWriter, request *http.Request) {
	var user userData
	var err error
	err = Decode(request, &user)
	if err != nil {
		fmt.Fprintf(writer, "Неверные данные")
		return
	}
	var user2 userData
	user2, err = storage.get(user.Phone)
	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}
	if user.Password != user2.Password {
		fmt.Fprintf(writer, "неверный пароль")
		return
	}
	fmt.Fprintf(writer, "Вы успешно вошли в систему")
}

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

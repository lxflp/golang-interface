package storage

import (
	"fmt"
	"github.com/lxflp/golang-interface/models"
)

var storage = UserStorageSlice{Data: make([]models.UserData, 0, 10)}

type UserStorager interface {
	Save(user models.UserData) error
	Get(phone string) (models.UserData, error)
	Delete(phone string) error
}

type UserStorageSlice struct {
	Data []models.UserData
}

func NewUserStorageSlice(data []models.UserData) *UserStorageSlice {
	return &UserStorageSlice{Data: data}
}

func (u *UserStorageSlice) Save(user models.UserData) error {
	for _, usr := range u.Data {
		if usr.Phone == user.Phone {
			return fmt.Errorf("данный номер уже зарегистрирован")
		}
	}
	u.Data = append(u.Data, user)
	return nil
}

func (u *UserStorageSlice) Get(phone string) (models.UserData, error) {
	for _, usr := range u.Data {
		if usr.Phone == phone {
			return usr, nil
		}
	}
	return models.UserData{}, fmt.Errorf("пользователь не найден")
}

func (u *UserStorageSlice) Delete(phone string) error {
	for i, usr := range u.Data {
		if usr.Phone == phone {
			u.Data = append(u.Data[:i], u.Data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("пользователь не найден")
}

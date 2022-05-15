package storage

import (
	"fmt"
	"github.com/lxflp/golang-interface/models"
)

type UserStorageMap struct {
	Data map[string]models.UserData
}

func (u *UserStorageMap) Save(user models.UserData) error {
	_, ok := u.Data[user.Phone]
	if ok {
		return fmt.Errorf("пользователь уже зарегистрирован")
	}
	u.Data[user.Phone] = user
	return nil
}

func (u *UserStorageMap) Get(phone string) (models.UserData, error) {
	user, ok := u.Data[phone]
	if !ok {
		return models.UserData{}, fmt.Errorf("пользователь не найден")
	}
	return user, nil
}

func (u *UserStorageMap) Delete(phone string) error {
	_, ok := u.Data[phone]
	if !ok {
		return fmt.Errorf("пользователь не найден")
	}
	delete(u.Data, phone)
	return nil

}

func NewUserStorageMap(data map[string]models.UserData) *UserStorageMap {
	return &UserStorageMap{Data: data}
}

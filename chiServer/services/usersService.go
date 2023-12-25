package services

import (
	"fmt"
)

type UserId string

type User struct {
	name string
}

type UserDto struct {
	Id   UserId `json:"id"`
	Name string `json:"name"`
}

type UsersService struct {
	storage KeyValueStorage[UserId, User]
}

func NewUsersService(storage KeyValueStorage[UserId, User]) *UsersService {
	return &UsersService{
		storage: storage,
	}
}

func (ur *UsersService) CreateUser(name string) (UserId, error) {
	userId := UserId(randomString(12))
	if _, exists := ur.storage.GetValue(userId); exists {
		return "", fmt.Errorf("error: user already exists")
	}

	ur.storage.SetValue(userId, User{name: name})

	return userId, nil
}

func (ur *UsersService) GetUser(id UserId) (string, error) {
	user, exists := ur.storage.GetValue(id)
	if !exists {
		return "", fmt.Errorf("error: user does not exist")
	}

	return user.name, nil
}

func (ur *UsersService) GetAllUsers() ([]UserDto, error) {
	all := ur.storage.All()
	userNames := []UserDto{}

	for k, v := range all {
		userNames = append(userNames, UserDto{
			Id:   k,
			Name: v.name,
		})
	}

	return userNames, nil
}

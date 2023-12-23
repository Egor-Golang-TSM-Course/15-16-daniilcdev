package services

import (
	"fmt"
	"math/rand"
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
	users map[UserId]User
}

func NewUsersService() *UsersService {
	return &UsersService{
		users: make(map[UserId]User),
	}
}

func (ur *UsersService) CreateUser(name string) (UserId, error) {
	userId := UserId(randomString(12))
	if _, exists := ur.users[userId]; exists {
		return "", fmt.Errorf("error: user already exists")
	}

	ur.users[userId] = User{name: name}

	return userId, nil
}

func (ur *UsersService) GetUser(id UserId) (string, error) {
	user, exists := ur.users[id]
	if !exists {
		return "", fmt.Errorf("error: user does not exist")
	}

	return user.name, nil
}

func (ur *UsersService) GetAllUsers() ([]UserDto, error) {
	userNames := make([]UserDto, 0, len(ur.users))

	for k, v := range ur.users {
		userNames = append(userNames, UserDto{
			Id:   k,
			Name: v.name,
		})
	}

	return userNames, nil
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

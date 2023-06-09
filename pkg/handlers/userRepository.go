package handlers

import (
	"fmt"
	. "playground/pkg/models"
	"strconv"
)

type UserRepository interface {
	GetAllUsers() []User
	GetUserByID(id string) (User, error)
	AddUser(user User) (string, error)
	UpdateUser(id string, newUser User) (User, error)
	DeleteUser(id string) error
}

type UserRepositoryImpl struct {
	users  []User
	nextID int
}

func (u UserRepositoryImpl) GetAllUsers() []User {
	return u.users
}

func (u UserRepositoryImpl) GetUserByID(id string) (User, error) {
	for _, user := range u.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

func (u UserRepositoryImpl) AddUser(user User) (string, error) {
	user.ID = strconv.Itoa(u.nextID)
	u.nextID++
	u.users = append(u.users, user)
	return user.ID, nil
}

func (u UserRepositoryImpl) UpdateUser(id string, newUser User) (User, error) {
	for i, user := range u.users {
		if user.ID == id {
			newUser.ID = id
			u.users[i] = newUser
			return newUser, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

func (u UserRepositoryImpl) DeleteUser(id string) error {
	for i, user := range u.users {
		if user.ID == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

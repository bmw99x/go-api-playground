package repository

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
	GenerateRepository() UserRepository
}

type UserRepositoryImpl struct {
	Users  []User
	NextID int
}

func GenerateRepository() UserRepositoryImpl {
	return UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) GetAllUsers() []User {
	return u.Users
}

func (u *UserRepositoryImpl) GetUserByID(id string) (User, error) {
	for _, user := range u.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

func (u *UserRepositoryImpl) AddUser(user User) (string, error) {
	user.ID = strconv.Itoa(u.NextID)
	u.NextID++
	u.Users = append(u.Users, user)
	return user.ID, nil
}

func (u *UserRepositoryImpl) UpdateUser(id string, newUser User) (User, error) {
	for i, user := range u.Users {
		if user.ID == id {
			newUser.ID = id
			u.Users[i] = newUser
			return newUser, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

func (u *UserRepositoryImpl) DeleteUser(id string) error {
	for i, user := range u.Users {
		if user.ID == id {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

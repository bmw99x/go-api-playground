package repository

import (
	"fmt"
	. "playground/pkg/models"
	"strconv"
)

// User repositories must implement these methods
type UserRepository interface {
	GetAllUsers() []User
	//GetUserByID(id string) (User, error)
	//AddUser(user User) (string, error)
	//UpdateUser(id string, newUser User) (User, error)
	//DeleteUser(id string) error
}

func GenerateRepository(useDb bool) UserRepository {
	if useDb {
		return DBUserRepositoryImpl{}
	}
	return UserRepositoryImpl{}
}

type UserRepositoryImpl struct {
	Users  []User
	NextID int
}

// Database definitions
type DBUserRepositoryImpl struct {}

func (u DBUserRepositoryImpl) GetAllUsers() []User {
	//TODO implement me
	panic("implement me")
}

func (u *DBUserRepositoryImpl) GetUserByID(id string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DBUserRepositoryImpl) AddUser(user interface{}) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DBUserRepositoryImpl) UpdateUser(id string, newUser interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (u *DBUserRepositoryImpl) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}


// In memory definitions
func (u UserRepositoryImpl) GetAllUsers() []User {
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

package models

import (
	"fmt"
	"strconv"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

var (
	users  []User
	nextID = 1
)

// GetAllUsers returns all users
func GetAllUsers() []User {
	return users
}

// GetUserByID retrieves a user by ID
func GetUserByID(id string) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

// AddUser adds a new user
func AddUser(user User) (string, error) {
	user.ID = strconv.Itoa(nextID)
	nextID++
	users = append(users, user)
	return user.ID, nil
}

// UpdateUser updates an existing user
func UpdateUser(id string, newUser User) (User, error) {
	for i, user := range users {
		if user.ID == id {
			newUser.ID = id
			users[i] = newUser
			return newUser, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

// DeleteUser deletes a user
func DeleteUser(id string) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

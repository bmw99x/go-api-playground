package tests

import (
	"fmt"
	. "playground/pkg/models"
	. "playground/pkg/repository"
	"testing"
)

func TestAddUser(t *testing.T) {
	/**
	Given a user
	When AddUser is called
	Then the user is added to the repository
	And the user is returned
	*/
	repo := GenerateRepository()
	user := User{
		ID:       "1",
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test",
	}
	id, err := repo.AddUser(user)
	if err != nil {
		t.Errorf("Error adding user: %v", err)
	}
	if repo.Users[0].ID != id {
		t.Errorf("Expected user ID to be %v, got %v", id, repo.Users[0].ID)
	}

}

func TestGetAllUsers(t *testing.T) {
	/**
	Given a list of users created by
		- AddUser
	When GetAllUsers is called
	Then the list of users is returned
	*/
	repo := GenerateRepository()
	user := User{
		Name:     "Test",
		Email:    "example@gmail.com",
		Password: "test",
	}
	var id, _ = repo.AddUser(user)
	users := repo.GetAllUsers()
	fmt.Println(id)
	if users[0].ID != id {
		t.Errorf("Expected user ID to be %v, got %v", id, "0")
	}

}

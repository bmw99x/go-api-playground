package tests

import (
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
	t.Parallel()
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
	t.Parallel()
	repo := GenerateRepository()
	user := User{
		Name:     "Test",
		Email:    "example@gmail.com",
		Password: "test",
	}
	var id, _ = repo.AddUser(user)
	users := repo.GetAllUsers()
	if users[0].ID != id {
		t.Errorf("Expected user ID to be %v, got %v", id, "0")
	}
}

func TestGetUserByID(t *testing.T) {
	/**
	Given a list of users created by
		- AddUser
	When GetUserByID is called with a valid ID
	Then the user with the given ID is returned
	*/
	t.Parallel()
	repo := GenerateRepository()
	expectedUser := User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test",
	}
	var id, _ = repo.AddUser(expectedUser)
	actualUser, _ := repo.GetUserByID(id)
	if actualUser.ID != id {
		t.Errorf("Expected user ID to be %v, got %v", id, actualUser.ID)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	/**
	Given no users created
	When GetUserByID is called with an invalid ID (i.e. any ID)
	Then an error is returned
	*/
	t.Parallel()
	repo := GenerateRepository()
	_ = User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test",
	}
	_, err := repo.GetUserByID("1")
	if err.Error() != "User not found" {
		t.Errorf("Expected error to be %v, got %v", "User not found", err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	/**
	Given a user created by
		- AddUser
	When UpdateUser is called with a valid ID
	and a modified user
	Then the user with the given ID is updated
	*/
	t.Parallel()
	repo := GenerateRepository()
	expectedUser := User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test",
	}
	id, _ := repo.AddUser(expectedUser)
	modifiedUser := User{
		ID:       id,
		Name:     "Test2",
		Email:    "test123@gmail.com",
		Password: "test",
	}
	actualUpdatedUser, _ := repo.UpdateUser(id, modifiedUser)
	if modifiedUser != actualUpdatedUser {
		t.Errorf("Expected user to be %v, got %v", modifiedUser, actualUpdatedUser)
	}
}

func TestDeleteUser(t *testing.T) {
	/**
	Given a user created by
		- AddUser
	When DeleteUser is called with a valid ID
	Then the user with the given ID is deleted
	*/
	t.Parallel()
	repo := GenerateRepository()
	expectedUser := User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test",
	}
	id, _ := repo.AddUser(expectedUser)
	_ = repo.DeleteUser(id)
	_, err := repo.GetUserByID(id)
	if err.Error() != "User not found" {
		t.Errorf("Expected error to be %v, got %v", "User not found", err.Error())
	}

}

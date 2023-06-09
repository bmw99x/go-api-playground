package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"playground/pkg/models"
	. "playground/pkg/repository"
)

var userRepo = UserRepositoryImpl{}

// ListUsers PingExample godoc
// @Summary List users
// @Schemes
// @Description List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func ListUsers(c *gin.Context) {
	users := userRepo.GetAllUsers() // Replace with your actual function to retrieve users

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// AddUser PingExample godoc
// @Summary Add a user
// @Description Add a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Router /users [post]
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := userRepo.AddUser(user) // Replace with your actual function to add a new user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "data": user})
}

// UpdateUser PingExample godoc
// @Summary Update a user
// @Description Update an existing user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userRepo.UpdateUser(id, user) // Replace with your actual function to update a user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser PingExample godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := userRepo.DeleteUser(id); err != nil { // Replace with your actual function to delete a user
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("Could not delete user due to:", err)})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

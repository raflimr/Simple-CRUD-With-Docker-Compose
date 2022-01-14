package controller

import (
	"net/http"

	"d-crud/models"
	"d-crud/utils"

	"github.com/gin-gonic/gin"
)

// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	utils.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Username: input.Username, Password: input.Password, Address: input.Address}
	utils.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Find a user by id
func FindUserById(c *gin.Context) { // Get model if exist
	var user models.User

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Find a user by username
func FindUserByUsername(c *gin.Context) { // Get model if exist
	var user models.User

	if err := utils.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update a user by id
func UpdateUserById(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Data
	var updatedInput models.User
	updatedInput.Username = input.Username
	updatedInput.Password = input.Password
	updatedInput.Address = input.Address

	utils.DB.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update a user by id
func UpdateUserByUsername(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := utils.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Data
	var updatedInput models.User
	updatedInput.Username = input.Username
	updatedInput.Password = input.Password
	updatedInput.Address = input.Address

	utils.DB.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Delete a user by id
func DeleteUserById(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	utils.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Delete a user by username
func DeleteUserByUsername(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := utils.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	utils.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

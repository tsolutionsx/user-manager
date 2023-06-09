package controllers

import (
	"ewc-backend-go/database"
	"ewc-backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check Exist User By Id
func checkIfUserExists(userId string) bool {
	var user models.User
	database.Instance.First(&user, userId)
	return user.ID != 0
}

// Register New User
func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

// Get All Users
func GetUsers(context *gin.Context) {
	var users []models.User
	record := database.Instance.Find(&users)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

// Get User by ID
func GetUserById(context *gin.Context) {
	userId := context.Param("id")
	if !checkIfUserExists(userId) {
		context.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
	}

	var user models.User
	database.Instance.First(&user, userId)
	context.JSON(http.StatusOK, gin.H{"user": user})
}

// Update User by Id
func UpdateUser(context *gin.Context) {
	userId := context.Param("id")
	if !checkIfUserExists(userId) {
		context.JSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
		return
	}

	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Instance.First(&user, userId).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.Instance.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})

}

// Delete User by ID
func DeleteUser(context *gin.Context) {
	userId := context.Param("id")
	if !checkIfUserExists(userId) {
		context.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	var user models.User
	database.Instance.Delete(&user, userId)
	context.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully!"})
}

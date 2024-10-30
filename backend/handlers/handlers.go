package handlers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpUserHandler(ctx *gin.Context) {
	var user models.User
	var registerData models.RegisterDataUser

	// parsing JSON request body into RegisterData structure
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	// hashing the password
	hashedPassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing error"})
		return
	}

	// fill in the data to be written to the database
	user.Name = registerData.Name
	user.Email = registerData.Email
	user.Hash = hashedPassword

	// save the user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the user"})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
	}
}

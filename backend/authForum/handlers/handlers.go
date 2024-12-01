package handlers

import (
	"authForum/database"
	"authForum/models"
	"authForum/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// user creation
func SignUpUserHandler(ctx *gin.Context) {
	var user models.User
	var registerData models.DataUser

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
		if strings.Contains(result.Error.Error(), "23505") { // unique constraint error
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the user"})
		}
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
	}
}

// User Authentication
func SignInUserHandler(ctx *gin.Context) {

	var registerData models.DataUser
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect data"})
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", registerData.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if !utils.CheckPasswordHash(registerData.Password, user.Hash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	tokens, err := utils.GenerateTokens(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tokens"})
		return
	}

	// Create an anonymous structure for the response
	userResponse := struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens, "user": userResponse})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValue := strings.Split(ctx.GetHeader("Authorization"), " ")
		// Checking if the Authorization header is present
		if len(tokenValue) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		accessToken, err := jwt.Parse(tokenValue[1],
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

		if err != nil || accessToken == nil || !accessToken.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		log.Println("AuthMiddleware - missed a token")
		ctx.Next()
	}
}

// getting all users
func GetUsersHandler(ctx *gin.Context) {
	var allUsers []models.User
	database.DB.Find(&allUsers)
	ctx.JSON(http.StatusOK, gin.H{"users": allUsers})
}

// user remove
func DelUeserHandler(ctx *gin.Context) {
	idUser := ctx.Param("id")

	var user models.User
	result := database.DB.Where("id = ?", idUser).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		}
		return
	}
	database.DB.Unscoped().Delete(&user)
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// user update
func UpdateUeserHandler(ctx *gin.Context) {

	var userUpdateData models.UpdDataUser
	if err := ctx.ShouldBindJSON(&userUpdateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect new data"})
		return
	}

	var user models.User
	result := database.DB.Where("id = ?", ctx.Param("id")).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if userUpdateData.Name == user.Name && userUpdateData.Email == user.Email && userUpdateData.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No new data"})
		return
	}

	var checkEmailUser models.User
	result = database.DB.Where("email = ?", userUpdateData.Email).First(&checkEmailUser)
	if result.RowsAffected > 0 && user.ID != checkEmailUser.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	refreshToken := ctx.GetHeader("Refresh-Token")
	userID, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	newTokens, err := utils.GenerateTokens(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	if len(userUpdateData.Password) != 0 {
		hashedPassword, err := utils.HashPassword(userUpdateData.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing error"})
			return
		}
		user.Hash = hashedPassword
	}

	user.Name = userUpdateData.Name
	user.Email = userUpdateData.Email
	database.DB.Save(&user)

	// Create an anonymous structure for the response
	userResponse := struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"tokens": newTokens, "user": userResponse})
}

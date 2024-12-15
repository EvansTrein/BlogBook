package handlers

import (
	"authForum/database"
	"authForum/models"
	"authForum/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUpUserHandler(ctx *gin.Context) {
	var userData models.DataUser

	if resultErrParseBody := ctx.BindJSON(&userData); resultErrParseBody != nil {
		ctx.JSON(400, models.ResponceError{ErrMess: "incorrect data in the body of the request", Err: resultErrParseBody.Error()})
		return
	}

	hashPass, err := utils.HashPassword(userData.Password)
	if err != nil {
		ctx.JSON(500, models.ResponceError{ErrMess: "password hashing error", Err: err.Error()})
		return
	}

	if createUserErrDB := database.SqlCreateUser(userData.Name, userData.Email, hashPass); createUserErrDB != nil {
		if strings.Contains(createUserErrDB.Error(), "pq: duplicate key value violates unique constraint") {
			ctx.JSON(400, models.ResponceError{ErrMess: "user with this email already exists", Err: createUserErrDB.Error()})
			return
		}
		ctx.JSON(500, models.ResponceError{ErrMess: "failed to create a record in the database", Err: createUserErrDB.Error()})
		return
	}

	ctx.JSON(201, models.ResponseMessage{Message: "user has been successfully created"})
}

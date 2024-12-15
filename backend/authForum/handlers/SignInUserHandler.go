package handlers

import (
	"authForum/database"
	"authForum/models"
	"authForum/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignInUserHandler(ctx *gin.Context) {
	var userData models.DataUser
	var activeUser models.User
	var respActiveUser models.RespActiveUser

	if resultErrParseBody := ctx.BindJSON(&userData); resultErrParseBody != nil {
		ctx.JSON(400, models.ResponceError{ErrMess: "incorrect data in the body of the request", Err: resultErrParseBody.Error()})
		return
	}

	if err := database.SqlFindUser(userData.Email, &activeUser); err != nil {
		if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(400, models.ResponceError{ErrMess: "user with this email address is not registered", Err: err.Error()})
			return
		}
		ctx.JSON(500, models.ResponceError{ErrMess: "database search failed", Err: err.Error()})
		return
	}

	if !utils.CheckPasswordHash(userData.Password, activeUser.PasswordHash) {
		ctx.JSON(400, models.ResponceError{ErrMess: "incorrect password"})
		return
	}

	tokens, err := utils.GenerateTokens(activeUser.Id)
	if err != nil {
		ctx.JSON(500, models.ResponceError{ErrMess: "Failed to create tokens", Err: err.Error()})
		return
	}

	respActiveUser.Id = activeUser.Id
	respActiveUser.UserName = activeUser.UserName
	respActiveUser.UserEmail = activeUser.UserEmail
	respActiveUser.UserTokens = tokens

	ctx.JSON(200, respActiveUser)
}

package controller

import (
	"net/http"

	"red-nigiri-api/internal/model"
	"red-nigiri-api/internal/validators"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var input validators.SignUpInput

	// Bind the request body into a type with `ShouldBindJSON`
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	newUser, err := user.Create()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": newUser})
}

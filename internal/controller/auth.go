package controller

import (
	"net/http"

	"red-nigiri-api/internal/helper"
	"red-nigiri-api/internal/model"
	"red-nigiri-api/internal/validators"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var input validators.SignUpInput

	// Bind the request body into a type with `ShouldBindJSON`.
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

func SignIn(context *gin.Context) {
	var input validators.SignInInput

	// Bind the request body into a type with `ShouldBindJSON`.
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByEmail(input.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.IsPasswordValid(input.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// maxAge: 3600 seconds equal 1 hour.
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", jwt, 3600*24*7, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{
		"message": "Authenticated",
	})

}

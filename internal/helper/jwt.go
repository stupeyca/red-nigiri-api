package helper

import (
	"os"
	"red-nigiri-api/internal/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtPrivateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	// iat: Issued at.
	// eat: Expires at.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	// `SignedString` creates and returns a complete, signed JWT.
	return token.SignedString(jwtPrivateKey)
}

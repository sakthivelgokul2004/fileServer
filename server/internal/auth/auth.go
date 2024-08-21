package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenJwt(id uuid.UUID) (string, error) {
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenSring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println("erros")
		log.Panicf("error %v", err)
		return "", err
	}
	return tokenSring, err
}


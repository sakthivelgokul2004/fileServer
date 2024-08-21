package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"server/internal/database"
	"server/internal/handlers"
)

type DBContex struct {
	DB *database.Queries
}

func (DBcontex DBContex) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("logiing")
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			handlers.RespondWithError(w, http.StatusInternalServerError, "can 't vaidate token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
			}
			id, ok := claims["sub"].(uuid.UUID)
			if ok == true {
				user, err := DBcontex.DB.GetUserById(r.Context(), id)
				if err != nil {
					handlers.RespondWithError(w, http.StatusUnauthorized, "User does not exsit")
				}
				ctx := context.WithValue(r.Context(), "user", user)
             r.WithContext(ctx)
			}
		} else {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
		}
	})
}

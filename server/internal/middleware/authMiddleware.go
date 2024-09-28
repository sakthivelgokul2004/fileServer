package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"server/internal/auth"
	"server/internal/database"
	"server/internal/handlers"
)

type DBContex struct {
	DB *database.Queries
}

func (DBcontex DBContex) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			log.Printf("%s at auth", err)
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
			return
		}
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			handlers.RespondWithError(w, http.StatusInternalServerError, "can 't vaidate token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
			}
			idstr, ok := claims["sub"].(string)
			id, _ := uuid.Parse(idstr)
			if ok {
				user, err := DBcontex.DB.GetUserById(r.Context(), id)
				if err != nil {
					handlers.RespondWithError(w, http.StatusUnauthorized, "User does not exsit")
				}
				ctx := context.WithValue(r.Context(), auth.UserContextKey, user)
				req := r.WithContext(ctx)
				next.ServeHTTP(w, req)
			}
		} else {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
		}
	})
}

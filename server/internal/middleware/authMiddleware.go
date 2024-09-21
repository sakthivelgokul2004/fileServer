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

	"server/internal/database"
	"server/internal/handlers"
)

type DBContex struct {
	DB *database.Queries
}

func (DBcontex DBContex) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hi")
		log.Println("ji")
		// if r.Method == http.MethodOptions {
		// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// 	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		// 	w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }
		cookie, err := r.Cookie("Authorization")
		cookies := r.Cookies()
		fmt.Println(len(cookies))
		fmt.Println(r.Header)
		if err != nil {
			log.Printf("%s at auth", err)
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
			return
		}
		fmt.Println(cookie.Value)
		fmt.Println(cookie)
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized  access")
		}
	})
}

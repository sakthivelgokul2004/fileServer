package middleware

// import (
// 	"net/http"

// 	"github.com/rs/cors"
// )

// // func CorsMiddleware(next http.Handler) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		c := cors.New(cors.Options{
// // 			AllowedOrigins:   []string{"http://localhost:8080/", "https://firebasestorage.googleapis.com/v0/b/fileserver-8c567.appspot.com/"},
// // 			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
// // 			AllowedHeaders:   []string{"Authorization", "Content-Type"},
// // 			AllowCredentials: true,
// // 		})
// // 	})
// // }

package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"os"
	"server/internal/database"
	"server/internal/handlers"
	"server/internal/middleware"
	sqlFs "server/sql"

	// firebase "firebase.google.com/go"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	// "github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var embedMigrations embed.FS = sqlFs.EmbedMigrations

func main() {

	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Error(err)
		log.Panicf("db connection failed %v", dbUrl)
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		println("goose")
		panic(err)
	}

	if err := goose.Up(db, "schema"); err != nil {
		panic(err)
	}
	dbConfig := handlers.DBConfig{
		DB: database.New(db),
	}
	dbcontex := middleware.DBContex{
		DB: database.New(db),
	}
	//
	// config := &firebase.Config{
	// StorageBucket: "fileserver-8c567.appspot.com",
	// }
	// app, err := firebase.NewApp(context.Background(), config)
	// if err != nil {
	// log.Fatalf("error initializing app: %v\n", err)
	// }
	// Storage, err := app.Storage(context.Background())
	// if err != nil {
	// log.Fatalln(err)
	// }
	// bucket, err := Storage.DefaultBucket()
	// if err != nil {
	// log.Fatalln(err)
	// }
	// it := bucket.Objects(context.Background(), nil)
	// arrs, err := it.Next()
	// if err != nil {
	// log.Fatalf("error listing objects: %v\n", err)
	// }
	// fmt.Println(arrs.Name, arrs.ContentType)
	router := http.NewServeMux()
	// fmt.Print("server  at on port 8080")
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /signup", dbConfig.SignupHandler)
	authRouter.HandleFunc("GET /getfile", dbConfig.GetFileUrl)
	authRouter.HandleFunc("POST /login", dbConfig.LoginHandler)
	router.Handle("/auth/", http.StripPrefix("/auth", middleware.Logging(authRouter)))
	authcatedRouter := http.NewServeMux()
	authcatedRouter.HandleFunc("/upload", handlers.Upload)
	authcatedRouter.HandleFunc("POST /addfile", dbConfig.Addfile)
	authcatedRouter.HandleFunc("GET /getfile", dbConfig.GetFileUrl)
	router.Handle("/user/", http.StripPrefix("/user", dbcontex.AuthMiddleware(authcatedRouter)))
	// corsHandler := cors.New(cors.Options{
	// 	AllowedOrigins:     []string{"http://localhost:8080", "http://localhost:5173"}, // Allow all origins
	// 	AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	OptionsPassthrough: true,
	// 	ExposedHeaders:     []string{"Authorization"},
	// 	MaxAge:             86400,
	// })
	server := http.Server{
		Addr: ":8080",
		// Handler: corsHandler.Handler(router),
		Handler: router,
	}

	fmt.Println("gdog")
	server.ListenAndServe()
}

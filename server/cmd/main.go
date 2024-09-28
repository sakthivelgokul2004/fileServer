package main

import (
	"context"
	"database/sql"
	"embed"
	"net/http"
	"os"
	"server/internal/database"
	"server/internal/handlers"
	"server/internal/middleware"
	sqlFs "server/sql"

	// firebase "firebase.google.com/go"
	firebase "firebase.google.com/go"
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
	config := &firebase.Config{
		StorageBucket: "fileserver-8c567.appspot.com",
	}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	Storage, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	//  bucket, err := Storage.DefaultBucket()
	//  if err != nil {
	//  	log.Fatalln(err)
	//  }

	dbConfig := handlers.DBConfig{
		DB:        database.New(db),
		Filestore: Storage,
	}
	dbcontex := middleware.DBContex{
		DB: database.New(db),
	}
	router := http.NewServeMux()
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
	server := http.Server{
		Addr: ":8080",
		// Handler: corsHandler.Handler(router),
		Handler: router,
	}

	server.ListenAndServe()
}

package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	log "github.com/sirupsen/logrus"

	"server/internal/database"
	"server/internal/handlers"
	"server/internal/middleware"
	sqlFs "server/sql"
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

	router := http.NewServeMux()
	// router.HandleFunc("/work", handlers.WellcomeHandler)
	log.SetReportCaller(true)
	router.HandleFunc("/", handlers.ServecFiles)
	fmt.Print("server  at on port 8080")

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /signup", dbConfig.SignupHandler)
	authRouter.HandleFunc("POST /login", dbConfig.LoginHandler)
	router.Handle("/auth/", http.StripPrefix("/auth", middleware.Logging(authRouter)))
	authcatedRouter := http.NewServeMux()
	authcatedRouter.HandleFunc("/upload", handlers.Upload)
	router.Handle("/user/", http.StripPrefix("/user", dbcontex.AuthMiddleware(authcatedRouter)))
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}

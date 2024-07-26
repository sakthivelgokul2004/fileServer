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

	"server/internal/handlers"
	sqlFs "server/sql"
)

func wellcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	println(r.Body)
	w.Write([]byte("hhi"))
}

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
	var mux http.ServeMux = *http.NewServeMux()
	mux.HandleFunc("/work", wellcomeHandler)
	log.SetReportCaller(true)
	mux.HandleFunc("/upload", handlers.Upload)
	mux.HandleFunc("/", handlers.ServecFiles)
	fmt.Print("server  at on port 8080")

	http.ListenAndServe(":8080", &mux)
}

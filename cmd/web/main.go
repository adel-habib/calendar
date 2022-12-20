package main

import (
	"github.com/adel-habib/calendar/pkg/calendar"
	"log"
	"net/http"
	"os"
	"text/template"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *calendar.Calendar
	templateCache map[string]*template.Template
}

func main() {
	addr := getEnv("address", ":4000")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

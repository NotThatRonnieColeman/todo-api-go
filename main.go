package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	dbType := flag.String("db", "", "Database to use. Options are: \"sqlite3\", \"mysql\" and \"mongo\"")
	flag.Parse()
	a := flag.Args()

	if len(a) != 0 {
		log.Fatalf("Uknown argument: %s", a[0])
	}

	var db Database
	if *dbType == "mongo" {
		connectionString := os.Getenv("CONNECTION_STRING")
		if connectionString == "" {
			log.Fatal("Missing value for CONNECTION_STRING environment variable.")
		}
		db = &mongodb{connectionString: connectionString}
	} else if *dbType == "sqlite3" || *dbType == "mysql" {
		connectionString := os.Getenv("CONNECTION_STRING")
		if connectionString == "" {
			log.Fatal("Missing value for CONNECTION_STRING environment variable.")
		}
		db = &gormdb{dialect: *dbType, connectionString: connectionString}
	} else {
		flag.Usage()
		log.Fatal("Please specify a valid database to use.")
	}

	db.init()
	defer db.close()
	router := mux.NewRouter()
	app := &Application{db: db, router: router}
	app.initRoutes()

	address := os.Getenv("HOST_ADDRESS")
	log.Printf("Starting web server on %s\n", address)
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

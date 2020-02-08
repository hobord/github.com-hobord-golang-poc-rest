package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpdelivery "github.com/hobord/golang-poc-rest/delivery/http"
	persistence "github.com/hobord/golang-poc-rest/infrastructure/mysql"
	"github.com/hobord/golang-poc-rest/usecase"
)

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	if dbConnection == "" {
		dbConnection = "dbuser:secret@tcp(mysql:3306)/testdb?multiStatements=true"
	}

	conn, err := sql.Open("mysql", dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	persistence.MigrationUp(conn)

	repository := persistence.NewFooMysqlRepository(conn)
	interactor := usecase.CreateFooInteractor(repository)

	r := mux.NewRouter()

	httpdelivery.MakeRouting(r, interactor)

	log.Print("App start, listen on port " + httpPort)
	log.Print("navigate to: http://localhost:" + httpPort + "/foo")
	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}

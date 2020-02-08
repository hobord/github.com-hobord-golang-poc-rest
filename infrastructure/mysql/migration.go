package persistence

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationUp(conn *sql.DB) {
	//conn, err := sql.Open("mysql", "postgres://localhost:5432/database?sslmode=enable")
	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://infrastructure/mysql/db",
		"mysql", driver)
	m.Steps(2)
	if err != nil {
		log.Fatal(err)
	}
	//if err := m.Up(); err != nil {
	//	log.Print(err)
	//}
}

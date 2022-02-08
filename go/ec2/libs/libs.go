package libs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"log"
)

type pg struct {
	db *sql.DB
}

func PostgresInit() (*pg, error) {

	var (
		host = os.Getenv("HOST_PG")
		port = os.Getenv("PORT_PG")
		dbname = os.Getenv("DB_PG")
		user = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return &pg{db}, nil

}

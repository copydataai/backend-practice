package libs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type pg struct {
	db *sql.DB
}

func PostgresInit() (*pg, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", "localhost", "user_read", "user_read", 5436, "reviews-camera")
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

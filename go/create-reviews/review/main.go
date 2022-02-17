package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Reviews struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	User    int32  `json:"author"`
	Product int32  `json:"product"`
}

func (this Reviews) createReviews(db *sql.DB) (bool, error) {
	fmt.Println(this)
	row := db.QueryRow(`INSERT INTO reviews(title, content, "user", product, created_at) VALUES($1, $2, $3, $4, $5)`, this.Title, this.Content, this.User, this.Product, time.Now().Format(time.RFC3339))
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

type pg struct {
	db *sql.DB
}

func initPg() (*pg, error) {
	var (
		host     =  os.Getenv("HOST_PG")
		port     =  os.Getenv("PORT_PG")
		dbname   =  os.Getenv("DB_PG")
		user     =  os.Getenv("USER_PG")
		password =  os.Getenv("PASSWORD_PG")
		sslmode  =  os.Getenv("SSLMODE_PG")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, dbname, sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &pg{db}, nil
}

func createReview(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	var review Reviews
	if err := json.Unmarshal([]byte(request.Body), &review); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	created, err := review.createReviews(db.db)
	if err != nil && !created {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(createReview)
}

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

// Singleton
type db struct {
	pg *sql.DB
}

type Products struct {
	Detail        int64  `json:"detail"`
	Trademark     string `json:"trademark"`
	Manufacturing string `json:"manufacturing"`
	SKU           string `json:"sku"`
}

func initPg() (*db, error) {
	var (
		host     = os.Getenv("HOST_PG")
		port     = os.Getenv("PORT_PG")
		dbname   = os.Getenv("DB_PG")
		user     = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	pg, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &db{pg}, nil
}

func (prod Products) createProducts(db *sql.DB) (bool, error) {
	row := db.QueryRow(`INSERT INTO products(detail, trademark, manufacturing, "SKU", created_at) VALUES($1, $2,$3, $4, $5)`, prod.Detail, prod.Trademark, prod.Manufacturing, prod.SKU, time.Now().Format(time.RFC3339))
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

func productHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	var product Products
	if err := json.Unmarshal([]byte(request.Body), &product); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	created, err := product.createProducts(db.pg)
	if err != nil && !created {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(productHandler)
}

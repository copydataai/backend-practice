package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

// Singleton
type db struct {
	pg *sql.DB
}

type Products struct {
	Detail        int64   `json:"detail"`
	Trademark     string  `json:"trademark"`
	Manufacturing string  `json:"manufacturing"`
	SKU           string  `json:"sku"`
	Price         int64 `json:"price"`
}

type VerifyClaims struct {
	jwt.StandardClaims
	User int64
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
	defer db.Close()
	row := db.QueryRow(`INSERT INTO products(detail, trademark, manufacturing, "SKU", price, created_at) VALUES($1, $2,$3, $4, $5, $6)`, prod.Detail, prod.Trademark, prod.Manufacturing, prod.SKU, prod.Price, time.Now().Format(time.RFC3339))
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

func verifyJWT(authToken string) error {
	onlyToken := strings.Replace(authToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(onlyToken, &VerifyClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("SECRET_KEY")
		return []byte(secret), nil
	})
	_, ok := token.Claims.(*VerifyClaims)
	if ok && token.Valid {
		return nil
	}
	return err
}

func productHandler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	authToken := request.Headers["Authorization"]
	if err := verifyJWT(authToken); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body:       err.Error(),
		}, nil
	}
	var product Products
	if err := json.Unmarshal([]byte(request.Body), &product); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}
	created, err := product.createProducts(db.pg)
	if err != nil && !created {
		return &events.APIGatewayProxyResponse{}, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(productHandler)
}

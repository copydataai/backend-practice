package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)
type db struct{
	pg *sql.DB
}

func initPg() (*db, error) {
	var (
		host = os.Getenv("HOST_PG")
		port = os.Getenv("PORT_PG")
		dbname = os.Getenv("DB_PG")
		user = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	pg, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &db{pg}, nil
}

type ProductType struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

// Create product-type
func (pType ProductType) createProductType(db *sql.DB) (bool, error){
	defer db.Close()
	row := db.QueryRow("INSERT INTO product_type(name, detail) VALUES($1, $2);", pType.Name, pType.Detail)
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

// JWT
type VerifyClaims struct{
	jwt.StandardClaims
	User int64
}

// validate JWT
func verifyJWT(authToken string)  error {
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

// Handler return Response
func productTypeHandler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error){
	db, err := initPg()
	defer db.pg.Close()
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	authToken, ok := request.Headers["Authorization"]
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body: `{"message": "Don't have Token"}`,
		}, nil
	}
	if err := verifyJWT(authToken); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body: err.Error(),
		}, nil
	}
	var productType ProductType
	if err := json.Unmarshal([]byte(request.Body), &productType) ; err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: err.Error(),
		}, nil
	}
	if created, err := productType.createProductType(db.pg); !created && err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	},nil
}

func main() {
	lambda.Start(productTypeHandler)
}


package main

import (
	"strings"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

type Reviews struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	User    int64  `json:"author"`
	Product int64  `json:"product"`
}

func (this Reviews) createReviews(db *sql.DB) (bool, error) {
	defer db.Close()
	row := db.QueryRow(`INSERT INTO reviews(title, content, "user", product, created_at) VALUES($1, $2, $3, $4, $5)`, this.Title, this.Content, this.User, this.Product, time.Now().Format(time.RFC3339))
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

type VerifyClaims struct {
	jwt.StandardClaims
	User int64
}

type pg struct {
	db *sql.DB
}

func initPg() (*pg, error) {
	var (
		host     = os.Getenv("HOST_PG")
		port     = os.Getenv("PORT_PG")
		dbname   = os.Getenv("DB_PG")
		user     = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &pg{db}, nil
}

func verifyJWT(authToken string) (int64, error) {
	onlyToken := strings.Replace(authToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(onlyToken, &VerifyClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("SECRET_KEY")
		return []byte(secret), nil
	})
	claims, ok := token.Claims.(*VerifyClaims)
	if ok && token.Valid {
		return claims.User, nil
	}
	return 0, err
}

func createReview(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var review Reviews
	db, err := initPg()
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	auth, ok := request.Headers["Authorization"]
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body: `{"message": "Don't have Token"}`,
		}, nil
	}
	userId, errToken := verifyJWT(auth)
	if errToken != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body: errToken.Error(),
		}, nil
	}
	if err := json.Unmarshal([]byte(request.Body), &review); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: err.Error(),
		}, nil
	}
	review.User = userId
	created, err := review.createReviews(db.db)
	if err != nil && !created {
		return &events.APIGatewayProxyResponse{}, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(createReview)
}

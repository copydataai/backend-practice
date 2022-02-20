package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserQuery struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type ResponseClaim struct {
	User int   `json:"user"`
	jwt.StandardClaims
}

type ResponseToken struct {
Token string `json:"token"`
}

// Verify user return JWT
func (user UserLogin) verifyUser(db *sql.DB) (string, error) {
	var userVerify UserQuery
	var SecretKey = os.Getenv("SECRET_KEY")
	defer db.Close()
	userQuery := db.QueryRow("SELECT id, password from users WHERE email = $1;", user.Email)
	if userQuery.Err() != nil {
		return "", userQuery.Err()
	}
	if err := userQuery.Scan(&userVerify.Id, &userVerify.Password); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userVerify.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	responseClaim := ResponseClaim{userVerify.Id, jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, responseClaim)
	ss, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

type pg struct {
	pg *sql.DB
}

func initPg() (*pg, error) {
	var (
		user     = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
		dbname   = os.Getenv("DB_PG")
		host     = os.Getenv("HOST_PG")
		port     = os.Getenv("PORT_PG")
	)
	configDB := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	db, err := sql.Open("postgres", configDB)
	if err != nil {
		return nil, err
	}
	return &pg{db}, nil
}

// Handler Login return string
func login(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	var user UserLogin
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: err.Error(),
		}, err
	}
	verify, errQuery := user.verifyUser(db.pg)
	if errQuery != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: `{"error": "Email or Password Incorrect"}`,
		}, nil
	}
	structByte, _ := json.Marshal(&ResponseToken{Token: verify})
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(structByte),
	}, nil
}

func main() {
	lambda.Start(login)
}

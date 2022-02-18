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
	Exp  int64 `json:"exp"`
	User int   `json:"user"`
	jwt.StandardClaims
}

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
	fmt.Println(userVerify)
	fmt.Println(user)
	if err := bcrypt.CompareHashAndPassword([]byte(userVerify.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	responseClaim := ResponseClaim{Exp: time.Now().Add(time.Hour * 72).Unix(), User: userVerify.Id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, responseClaim)
	ss, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	fmt.Println("Error to SignedString")
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
func login(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	var user UserLogin
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	verify, errorr := user.verifyUser(db.pg)
	if errorr != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       verify,
	}, nil
}

func main() {
	lambda.Start(login)
}

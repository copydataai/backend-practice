package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type UserCreate struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Speciality string `json:"speciality"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (user UserCreate) insertUser(db *sql.DB) string {
	db.QueryRow("INSERT INTO users(name, country, speciality, role, created_at, email, password) VALUES($1, $2, $3, $4, $5, $6)", user.Name, user.Country, user.Speciality, user.Role, time.Now().String(), user.Email, user.Password)
	return "User Created"
}

type pg struct {
	pg *sql.DB
}

func initPg() *pg {
	var (
		user     = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
		dbname   = os.Getenv("DB_PG")
		host     = os.Getenv("HOST_PG")
		port     = os.Getenv("PORT_PG")
	)
	configDB := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, dbname, host, port)
	db, err := sql.Open("postgres", configDB)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return &pg{db}
}

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func signup(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//db := initPg()
	var user UserCreate
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		log.Fatal(err)
	}
	newUser, _ := json.Marshal(user)
	if string(newUser) == request.Body {
		fmt.Println("This is Equal")
	}
	//body := user.insertUser(db.pg, request.Body)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(newUser),
	}, nil
}

func main() {
	lambda.Start(signup)
}

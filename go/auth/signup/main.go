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
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Speciality string `json:"speciality"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type verifyUser struct{
	Email string `json:"email"`
}

func (user UserCreate) insertUser(db *sql.DB) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	insert := db.QueryRow("INSERT INTO users(name, country, speciality, role, created_at, email, password) VALUES($1, $2, $3, $4, $5, $6, $7)", user.Name, user.Country, user.Speciality, user.Role, time.Now().Format(time.RFC3339), user.Email, string(passwordHash))
	if insert.Err() != nil {
		return false, insert.Err()
	}
	return true, nil
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

func signup(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	db, err := initPg()
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	defer db.pg.Close()
	var user UserCreate
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: err.Error(),
		}, nil
	}
	if created, err := user.insertUser(db.pg) ; !created && err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func main() {
	lambda.Start(signup)
}

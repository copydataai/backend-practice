package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user UserLogin) verifyUser(db *sql.DB) string {
	var userVerify UserLogin
	userQuery := db.QueryRow("SELECT password from users WHERE email = $1", user.Email)
	userQuery.Scan(&userVerify.Email, &userVerify.Password)
	db.Close()
	return "is user"
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
func login(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := initPg()
	var user UserLogin
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		log.Fatal(err)
	}
	body := user.verifyUser(db.pg)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(login)
}

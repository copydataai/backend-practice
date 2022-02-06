package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Reviews struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	User       int32    `json:"author"`
	Product    int32 `json:"product"`
	Created_at time.Time
}

func productType(db *sql.DB) string {
	db.QueryRow("INSERT INTO product_type(name, detail) VALUES($1, $2);")
	return "ProducType created"
}

func product(db *sql.DB) string {
	db.QueryRow("INSERT INTO products(detail, trademark, manufacturing, sku, created_at) VALUES(%d, %s, %s %s, $5)", time.Now().String())
	return "product created"
}

func (this Reviews) reviews(db *sql.DB) string {
	db.QueryRow("INSERT INTO reviews(title, content, user, product created_at) VALUES($1, $2, $3, $4, $5)")
	return "review created"
}

type pg struct {
	db *sql.DB
}

func init() (*pg, error){
	var (
		host = os.Getenv("HOST_PG")
		port = os.Getenv("PORT_PG")
		dbname = os.Getenv("DB_PG")
		user = os.Getenv("USER_PG")
		password = os.Getenv("PASSWORD_PG")
		sslmode= os.Getenv("SSLMODE_PG")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, dbname, sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
	return &pg{db}, nil
}

func createReview(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	},nil
}

func hello() (string, error){
	return "Hello Lambda Test", nil
}

func main() {
	lambda.Start(hello)
}

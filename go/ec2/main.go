package main

import (
	"log"
	"net/http"

	"github.com/copydataai/backend-practice/go/ec2/libs"
	"github.com/copydataai/backend-practice/go/ec2/services"
	_ "github.com/lib/pq"
)

func main() {
	posts, err := libs.PostgresInit()
	if err != nil {
		log.Fatal(err)
	}
	deps := services.Dependencies{Posts: posts}
	router := services.InitRouter(deps)
	log.Fatal(http.ListenAndServe(":3000", router))
}

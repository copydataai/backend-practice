package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PingResponse struct {
	Message string `json:"message"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{Message: "pong"}

	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal("err %v Error ping marshal", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(respBytes)
}

func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods("GET")

	// Version 1 API management

	router.HandleFunc("/reviews", ListReviewsHandler(deps)).Methods("GET") //.Headers(versionHeader, v1)
	router.HandleFunc("/products", ListProductsHandler(deps)).Methods("GET")
	router.HandleFunc("/types", ListProductTypeHandler(deps)).Methods("GET")
	return
}

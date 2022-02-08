package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListProductsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := deps.Posts.ListProducts()
		if err != nil {
			log.Fatal("error fetching data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(products)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(respBytes)
	})
}


func GetProductHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal("error bad request {id} = int")
			w.WriteHeader(http.StatusBadRequest)
		}
		product, err := deps.Posts.GetProductById(int64(idInt))
		if err != nil {
			log.Fatal("error fetching data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		respBytes, err := json.Marshal(product)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(respBytes)
	})
}

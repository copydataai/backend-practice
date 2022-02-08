package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListReviewsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reviews, err := deps.Posts.ListReviews()
		if err != nil {
			log.Fatal("error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(reviews)
		if err != nil {
			log.Fatal(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}

func GetReviewHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal("error bad request {id} = int")
			w.WriteHeader(http.StatusBadRequest)
		}
		review, err := deps.Posts.GetReviewById(int64(idInt))
		if err != nil {
			log.Fatal("error fetching data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		respBytes, err := json.Marshal(review)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(respBytes)
	})
}

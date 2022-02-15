package services

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListReviewsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reviews, count, err := deps.Posts.ListReviews()
		if err != nil {
			if count == 0{
				ErrNotFound.Send(w)
				return
			}
			ErrBadFetch.Send(w)
			return
		}
		Success(reviews, http.StatusOK, count).Send(w)
	})
}

func GetReviewHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			ErrBadRequest.Send(w)
			return
		}
		review, count, err := deps.Posts.GetReviewById(int64(idInt))
		if err != nil {
			if count == 0{
				ErrNotFound.Send(w)
				return
			}
			ErrBadFetch.Send(w)
			return
		}
		Success(review, http.StatusOK, count).Send(w)
	})
}

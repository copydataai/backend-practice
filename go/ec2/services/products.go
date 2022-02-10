package services

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListProductsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, count, err := deps.Posts.ListProducts()
		if err != nil {
			ErrBadFetch.Send(w)
			return

		}
		Success(products, http.StatusOK, count).Send(w)
	})
}


func GetProductHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			ErrBadRequest.Send(w)
			return
		}
		product, count, err := deps.Posts.GetProductById(int64(idInt))
		if err != nil {
			ErrBadFetch.Send(w)
			return
		}
		Success(product, http.StatusOK, count).Send(w)
	})
}

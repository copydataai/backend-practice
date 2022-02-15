package services

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListProductTypeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productTypes, count, err := deps.Posts.ListProductTypes()
		if err != nil {
			if count == 0{
				ErrNotFound.Send(w)
				return
			}
			ErrBadFetch.Send(w)
			return
		}
		Success(productTypes, http.StatusOK, count).Send(w)
	})
}

func GetProductTypeHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			ErrBadRequest.Send(w)
			return
		}
		productType, count, err := deps.Posts.GetProductTypeById(int64(idInt))
		if err != nil {
			if count == 0{
				ErrNotFound.Send(w)
				return
			}
			ErrBadFetch.Send(w)
			return
		}
		Success(productType, http.StatusOK, count).Send(w)
	})
}

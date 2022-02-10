package services

import (
	"encoding/json"
	"net/http"
)
type ResponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  Result `json:"result,omitempty"`
}

type Result struct {
	Count int  `json:"count"`
	Result interface{} `json:"result,omitempty"`
}

func (r ResponseAPI) Send(w http.ResponseWriter)  {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Result)
}

func Success(result interface{}, status int, count int) *ResponseAPI{
	return &ResponseAPI{
		Success: true,
		Status: status,
		Result: Result{count, result},
	}
}

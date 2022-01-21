package handlers

import (
	"fmt"
	"github.com/stretchr/testify/asssert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestReviewsFind(t *testing.T) {
	t.Run("Check return strings", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:3000/", nil)
		writer := httptest.NewRecoder()
		starter.CheckHealth(writer, req)
		response := writer.Result()
		body, _ := io.ReadAll(response.Body)
		assert
	})
}

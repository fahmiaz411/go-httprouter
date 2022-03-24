package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPatternNamed(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Product: "+params.ByName("id") + "id: " + params.ByName("itemId"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/4/items/az", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello 4", string(body))
}
func TestPatternCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Product: " + params.ByName("image"))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/4/items/az", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product: /4/items/az", string(body))
}
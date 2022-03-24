package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		fmt.Fprint(w, "Hello")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
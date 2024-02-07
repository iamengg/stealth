package main

import (
	"log"
	"net/http"
	 "user-service/handler"
)

func main() {
	http.HandleFunc("users", UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

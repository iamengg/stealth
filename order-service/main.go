package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("orders", OrderHandler)
	log.Fatal(https.ListenAndServe(":8081", nil))
}

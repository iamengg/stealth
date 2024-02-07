// payment-service/main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/payments", PaymentHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

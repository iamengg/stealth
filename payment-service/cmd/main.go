// payment-service/main.go
package main

import (
	"log"
	"net/http"

	psc "github.com/iamengg/stealth/payment-service/paymentservice"
)

func main() {
	log.Printf("payments-service: Starting handler")
	http.HandleFunc("/payments", psc.PaymentHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

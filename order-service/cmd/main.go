package main

import (
	"log"
	"net/http"

	osc "github.com/iamengg/stealth/order-service/orderservice"
)

func main() {
	log.Printf("orders-service: Starting handler")
	http.HandleFunc("orders", osc.OrderHandler)
	http.HandleFunc("/", osc.OrderHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// payment-service/handler.go
package main

import (
	"encoding/json"
	"net/http"
)

type Payment struct {
	ID      string `json:"id"`
	OrderID string `json:"orderId"`
	Amount  int    `json:"amount"`
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	payments := []Payment{
		{ID: "1", OrderID: "1", Amount: 100},
		{ID: "2", OrderID: "2", Amount: 200},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

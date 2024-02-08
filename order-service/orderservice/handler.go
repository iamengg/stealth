package orderservice

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Amount int    `json:"amount"`
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	orders := []Order{
		{ID: "1", UserID: "1", Amount: 100},
		{ID: "2", UserID: "2", Amount: 200},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

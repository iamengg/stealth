package userservice

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler at user-service")
	usr := []User{
		{ID: "123", Name: "pratik"},
	}
	w.Header().Set("Content-Type", "application/json")

	orderSvcUrl := os.Getenv("ORDER_SVC_URL")
	if orderSvcUrl == "" {
		log.Printf("Error in getting orderSvcUrl %v", orderSvcUrl)
		return
	}
	resp, err := http.Get(orderSvcUrl + "/orders")
	if err != nil {
		log.Printf(err.Error())
		return
	}
	log.Printf("Response from orderSvc is %v", resp.Body)
	defer resp.Body.Close()
	json.NewEncoder(w).Encode(usr)
}

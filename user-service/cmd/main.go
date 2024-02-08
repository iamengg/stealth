package main

import (
	"log"
	"net/http"
	"os"

	usc "github.com/iamengg/stealth/user-service/userservice"
)

func main() {
	setEnv()

	log.Printf("user-service: Starting handler")
	http.HandleFunc("users", usc.UserHandler)
	http.HandleFunc("/", usc.UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("user-service: Shutting down.")
}

func setEnv() {
	os.Setenv("USER_SVC_URL", "http://localhost:8080")
	os.Setenv("ORDER_SVC_URL", "http://localhost:8081")
	os.Setenv("PAYMENT_SVC_URL", "http://localhost:8082")
}

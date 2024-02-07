package userservice

import (
	"log"
	"net/http"
	"github.com/iamengg/stealth/userservice"
)

func main() {
	http.HandleFunc("users", UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

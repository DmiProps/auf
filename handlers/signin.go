package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DmiProps/auf/types"
)

// Signin is handler for signin page
func Signin(w http.ResponseWriter, r *http.Request) {

	// Get existing account data
	data := types.SignInData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Error Signin():", err)
		return
	}
	r.Body.Close()

	json.NewEncoder(w).Encode(struct{ Ok bool }{true})

}

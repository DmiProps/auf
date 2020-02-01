package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DmiProps/auf/communications"
	"github.com/DmiProps/auf/database"
	"github.com/DmiProps/auf/types"
)

type responseData struct {
	Ok       bool   `json:"ok`
	UserMsg  string `json:"userMsg`
	EmailMsg string `json:"emailMsg`
	PhoneMsg string `json:"phoneMsg`
}

// Signup is handler for signup page
func Signup(w http.ResponseWriter, r *http.Request) {

	// Get new account data
	data := types.SignUpData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}
	r.Body.Close()

	// Validate and create account
	msg, err := database.AddAccount(&data)
	if err != nil {
		log.Fatalln("Error AddAccount(): ", err)
	} else if msg != nil && len(msg) > 0 {
		response := responseData{Ok: false, UserMsg: msg["user"], EmailMsg: msg["email"], PhoneMsg: msg["phone"]}
		json.NewEncoder(w).Encode(response)
	} else {
		communications.SendActivationMail(data.User, data.Email)
		response := responseData{Ok: true}
		json.NewEncoder(w).Encode(response)
	}

}

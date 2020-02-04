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
	Ok       bool
	UserMsg  string
	EmailMsg string
	PhoneMsg string
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
		response := responseData{Ok: true}
		if err = communications.SendActivationMail(&data); err != nil {
			response.Ok = false
			response.EmailMsg = "Failed to send activation e-mail."
			log.Println(err)
		}
		json.NewEncoder(w).Encode(response)
	}

}

// ActivateViaEmail try activate account via e-mail
func ActivateViaEmail(w http.ResponseWriter, r *http.Request) {

}

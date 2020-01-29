package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmiProps/auf/communications"
)

type signUpData struct {
	User  string
	Pass  string
	Email string
	Phone string
}

// Signup is handler for signup page
func Signup(w http.ResponseWriter, r *http.Request) {

	data := signUpData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return
	}
	r.Body.Close()

	communications.SendActivationMail(data.User, data.Email)

}

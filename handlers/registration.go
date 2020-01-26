package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmiProps/auf/modules"
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

	modules.SendActivationMail(data.User, data.Email)

}

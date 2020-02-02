package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

func getDigits(in string) string {

	var digits string = "0123456789"
	var result string
	for _, ch := range in {
		if strings.ContainsRune(digits, ch) {
			result += string(ch)
		}
	}

	return result

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

	// Get phone digits
	if data.Phone != "" {
		data.PhoneDigits = getDigits(data.Phone)
	}

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

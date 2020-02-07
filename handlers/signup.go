package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/DmiProps/auf/communications"
	"github.com/DmiProps/auf/database"
	"github.com/DmiProps/auf/types"
)

// Signup is handler for signup page
func Signup(w http.ResponseWriter, r *http.Request) {

	// Get new account data
	data := types.SignUpData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}
	r.Body.Close()

	// Validate and create account
	msg, err := database.AddAccount(&data)
	if err != nil {
		log.Fatalln("Error AddAccount(): ", err)
	} else if msg != nil && len(msg) > 0 {
		response := types.SignUpResult{Ok: false, UserMsg: msg["user"], EmailMsg: msg["email"], PhoneMsg: msg["phone"]}
		json.NewEncoder(w).Encode(response)
	} else {
		response := types.SignUpResult{Ok: true}
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

	accountID := r.FormValue("link")

	response := types.ActivateEmailResult{SignInHidden: true, SignUpHidden: true, ResendLinkHidden: true, Message: ""}

	if accountID == "" {
		response.Message = "To activate your account, follow the link sent to the e-mail address specified when creating your account."
		response.SignUpHidden = false
	} else {

		msg, usr, err := database.ActivateAccountViaEmail(accountID)
		if err != nil {
			log.Println("Error ActivateAccountViaEmail(): ", err)
			response.Message = "An error occurred while activating your account. Please try again later."
		} else if msg != "" && usr == "" {
			response.Message = msg
			response.SignUpHidden = false
		} else if msg != "" && usr != "" {
			response.Message = fmt.Sprintf("Dear %s, the activation link has expired. You can resend the link.", usr)
			response.ResendLinkHidden = false
		} else {
			response.Message = fmt.Sprintf("Dear %s, your account has been successfully activated!", usr)
			response.SignInHidden = false
		}

	}

	t, _ := template.ParseFiles("www/activate-link.html")
	t.Execute(w, response)

}

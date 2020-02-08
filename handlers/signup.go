package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/DmiProps/auf/communications"
	"github.com/DmiProps/auf/database"
	"github.com/DmiProps/auf/templates"
	"github.com/DmiProps/auf/types"
)

// Signup is handler for signup page
func Signup(w http.ResponseWriter, r *http.Request) {

	// Get new account data
	data := types.SignUpData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Error Signup():", err)
		return
	}
	r.Body.Close()

	// Validate and create account
	msg, err := database.AddAccount(&data)
	if err != nil {
		log.Println("Error AddAccount():", err)
	} else if msg != nil && len(msg) > 0 {
		response := types.SignUpResult{Ok: false, UserMsg: msg["user"], EmailMsg: msg["email"], PhoneMsg: msg["phone"]}
		json.NewEncoder(w).Encode(response)
	} else {
		response := types.SignUpResult{Ok: true}
		if err = communications.SendActivationMail(&data); err != nil {
			response.Ok = false
			response.EmailMsg = templates.GetMessage(6)
			log.Println(err)
		}
		json.NewEncoder(w).Encode(response)
	}

}

// ResendLink resend e-mail link for activation account
func ResendLink(w http.ResponseWriter, r *http.Request) {

	// Get new account data
	data := types.SignUpData{ActivationLink: r.FormValue("link")}

	// Send activation e-mail
	msg, err := database.UpdateActivationLink(&data)
	if err != nil {
		log.Println("Error UpdateActivationLink():", err)
	}
	if msg != "" {
		json.NewEncoder(w).Encode(struct{ message string }{msg})
	} else if err = communications.SendActivationMail(&data); err != nil {
		log.Println("Error SendActivationMail():", err)
		json.NewEncoder(w).Encode(struct{ message string }{templates.GetMessage(6)})
	} else {
		json.NewEncoder(w).Encode(struct{ message string }{templates.GetMessage(1)})
	}

}

// ActivateViaEmail try activate account via e-mail
func ActivateViaEmail(w http.ResponseWriter, r *http.Request) {

	link := r.FormValue("link")

	response := types.ActivateEmailResult{SignInHidden: true, SignUpHidden: true, ResendLinkHidden: true, Message: ""}

	if link == "" {
		response.Message = templates.GetMessage(4)
		response.SignUpHidden = false
	} else {
		err := database.ActivateAccountViaEmail(link, &response)
		if err != nil {
			log.Println("Error ActivateAccountViaEmail():", err)
		}
	}

	t, _ := template.ParseFiles("www/activate-link.html")
	t.Execute(w, response)

}

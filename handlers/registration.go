package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	fmt.Println(data.User)
	fmt.Println(data.Pass)
	fmt.Println(data.Email)
	fmt.Println(data.Phone)

}

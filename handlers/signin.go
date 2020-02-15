package handlers

import (
	"encoding/json"
	"net/http"
)

// Signin is handler for signin page
func Signin(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(struct{ Ok bool }{true})

}

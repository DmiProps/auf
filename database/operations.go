package database

import (
	"context"
	"fmt"

	"github.com/DmiProps/auf/settings"
	"github.com/DmiProps/auf/types"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword return hash password
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

// CheckPasswordHash check password and hash
func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func validateAccount(data *types.SignUpData) (map[string]string, error) {

	msg := make(map[string]string)

	rows, err := settings.DbConnect.Query(
		context.Background(),
		`select 1 as check_type from accounts where lower(username) = lower($1)
		union all
		select 2 as check_type from accounts where lower(email) = lower($2)
		union all
		select 3 as check_type from accounts where phone <> '' and lower(phone) = lower($3)`,
		data.User,
		data.Email,
		data.Phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checkType int
	for rows.Next() {
		rows.Scan(&checkType)
		switch checkType {
		case 1:
			msg["user"] = "Username cannot be used. Please choose another username."
		case 2:
			msg["email"] = "A user is already registered with this e-mail address."
		case 3:
			msg["phone"] = "A user is already registered with this phone number."
		}
	}

	return msg, nil

}

// AddAccount validate account data and add new account
func AddAccount(data *types.SignUpData) (map[string]string, error) {

	// Validate account
	msg, err := validateAccount(data)
	if err != nil {
		return nil, err
	}
	if len(msg) > 0 {
		return msg, nil
	}

	// Get hash password
	hashPass, err := HashPassword(data.Pass)
	if err != nil {
		return nil, err
	}

	// Add account
	rows, err := settings.DbConnect.Query(
		context.Background(),
		`insert into accounts(username, email, password_hash, phone, creation_date) values ($1, $2, $3, $4, now())
		returning id`,
		data.User,
		data.Email,
		hashPass,
		data.Phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//TO-DO: create ref and code

	id := 0
	if rows.Next() {
		rows.Scan(&id)
		fmt.Println("Account Id: ", id)
	}

	return nil, nil

}

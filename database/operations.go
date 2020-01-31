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
		`select 1 as check_type, email_confirmed, phone_confirmed from account where lower(username) = lower($1)
		union
		select 2, email_confirmed, phone_confirmed from account where lower(email) = lower($2)
		union
		select 3, email_confirmed, phone_confirmed from account where lower(phone) = lower($3) and phone <> ''`,
		data.User)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checkType int
	var emailConfirmed, phoneConfirmed bool
	for rows.Next() {
		rows.Scan(&checkType, &emailConfirmed, &phoneConfirmed)
		switch checkType {
		case 1:
			if emailConfirmed || phoneConfirmed {
				msg["user"] = "A user with this name already exists"
			} else {
				msg["user"] = "*Click* to resend the activation email"
			}
		case 2:
			if emailConfirmed {
				msg["email"] = "A user with this email already exists"
			} else {
				msg["email"] = "*Click* to resend the activation email"
			}
		case 3:
			if phoneConfirmed {
				msg["phone"] = "A user with this phone number already exists"
			} else {
				msg["phone"] = "*Click* to resend the activation code"
			}
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

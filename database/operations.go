package database

import (
	"context"
	"strings"
	"time"

	"github.com/DmiProps/auf/settings"
	"github.com/DmiProps/auf/types"

	"github.com/beevik/guid"
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
		select 3 as check_type from accounts where phone_digits <> '' and phone_digits = $3`,
		data.User,
		data.Email,
		data.PhoneDigits)
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

func getActivationRef() string {

	var guid *guid.Guid = guid.New()

	return guid.String()

}

// AddAccount validate account data and add new account
func AddAccount(data *types.SignUpData) (map[string]string, error) {

	// Get phone digits
	data.PhoneDigits = getDigits(data.Phone)

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

	// Get activation ref
	data.ActivationRef = getActivationRef()

	// Add account
	rows, err := settings.DbConnect.Query(
		context.Background(),
		`insert into accounts(username, email, password_hash, phone, phone_digits, creation_date) values ($1, $2, $3, $4, $5, now())
		returning id`,
		data.User,
		data.Email,
		hashPass,
		data.Phone,
		data.PhoneDigits)
	if err != nil {
		return nil, err
	}

	// Get account id
	id := 0
	if rows.Next() {
		rows.Scan(&id)
	}
	rows.Close()

	// Add confirmation email ref
	if settings.AppSettings.Signup.ActualRefHours == 0 {
		rows, err = settings.DbConnect.Query(
			context.Background(),
			`insert into email_confirmations(account_id, ref) values ($1, $2)`,
			id,
			data.ActivationRef)
	} else {
		actualDate := time.Now().Add(time.Hour * time.Duration(settings.AppSettings.Signup.ActualRefHours))
		rows, err = settings.DbConnect.Query(
			context.Background(),
			`insert into email_confirmations(account_id, ref, actual_date) values ($1, $2, $3)`,
			id,
			data.ActivationRef,
			actualDate)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return nil, nil

}

// ActivateAccountViaEmail activate account via e-mail
func ActivateAccountViaEmail(ref string) error {

	//TO-DO: check activation ref, confirmation e-mail
	return nil

}

package database

import (
	"context"
	"strings"
	"time"

	"github.com/DmiProps/auf/settings"
	"github.com/DmiProps/auf/templates"
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

	// Get activation link
	data.ActivationLink = getActivationLink()

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

	// Add confirmation email link
	if settings.AppSettings.Signup.ActualLinkHours == 0 {
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`insert into email_confirmations(account_id, link) values ($1, $2)`,
			id,
			data.ActivationLink)
	} else {
		actualDate := time.Now().Add(time.Hour * time.Duration(settings.AppSettings.Signup.ActualLinkHours))
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`insert into email_confirmations(account_id, link, actual_date) values ($1, $2, $3)`,
			id,
			data.ActivationLink,
			actualDate)
	}
	if err != nil {
		return nil, err
	}

	return nil, nil

}

// UpdateActivationLink get user data and update activation link
func UpdateActivationLink(data *types.SignUpData) (string, error) {

	// Get account data
	rows, err := settings.DbConnect.Query(
		context.Background(),
		`select id, username, email from accounts inner join email_confirmations on account_id = id where lower(link) = lower($1)`,
		data.ActivationLink)
	if err != nil {
		return templates.GetMessage(5), err
	}

	if !rows.Next() {
		rows.Close()
		return templates.GetMessage(5), nil
	}

	var accountID int
	rows.Scan(&accountID, &data.User, &data.Email)
	rows.Close()

	// Get activation link
	data.ActivationLink = getActivationLink()

	// Update activation link
	if settings.AppSettings.Signup.ActualLinkHours == 0 {
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`update email_confirmations set link = $1, actual_date = null where account_id = $2`,
			data.ActivationLink,
			accountID)
	} else {
		actualDate := time.Now().Add(time.Hour * time.Duration(settings.AppSettings.Signup.ActualLinkHours))
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`update email_confirmations set link = $1, actual_date = $2 where account_id = $3`,
			data.ActivationLink,
			actualDate,
			accountID)
	}
	if err != nil {
		return templates.GetMessage(5), err
	}

	return "", nil

}

// ActivateAccountViaEmail activate account via e-mail
func ActivateAccountViaEmail(link string, result *types.ActivateEmailResult) error {

	rows, err := settings.DbConnect.Query(
		context.Background(),
		`select account_id, actual_date, username from email_confirmations inner join accounts on account_id = id where lower(link) = lower($1)`,
		link)
	if err != nil {
		result.Message = templates.GetMessage(2)
		return err // Try again
	}

	if !rows.Next() {
		rows.Close()
		result.Message = templates.GetMessage(0)
		result.SignUpHidden = false
		return nil // Sign Up
	}

	var accountID int
	var actualDate *time.Time
	var userName string

	err = rows.Scan(&accountID, &actualDate, &userName)
	rows.Close()
	if err != nil {
		result.Message = templates.GetMessage(2)
		return err
	}

	if actualDate == nil || actualDate.IsZero() || actualDate.After(time.Now()) {
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`update accounts set email_confirmed = true where id = $1`,
			accountID)
		if err != nil {
			result.Message = templates.GetMessage(2)
			return err // Try again
		}
		_, err = settings.DbConnect.Exec(
			context.Background(),
			`delete from email_confirmations where account_id = $1`,
			accountID)
		if err != nil {
			result.Message = templates.GetMessage(2)
			return err // Try again
		}

		result.Message = templates.GetMessage(3, userName)
		result.SignInHidden = false
		return nil // Sign In
	}

	// If the activation link has expired, must enter account information again
	result.Message = templates.GetMessage(1, userName)
	result.ResendLinkHidden = false
	return nil // Resend

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

func getActivationLink() string {

	var guid *guid.Guid = guid.New()

	return guid.String()

}

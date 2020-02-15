package types

// SignUpData contains the data of the created account and validation messages
type SignUpData struct {
	User           string
	Pass           string
	Email          string
	Phone          string
	PhoneDigits    string
	ActivationLink string
}

// SignUpResult contains result sending sign up data
type SignUpResult struct {
	Ok          bool
	UserMsg     string
	PassMsg     string
	EmailMsg    string
	PhoneMsg    string
	ActivateMsg string
}

// ActivateEmailResult contains result activation account via e-mail link
type ActivateEmailResult struct {
	SignInHidden     bool
	SignUpHidden     bool
	ResendLinkHidden bool
	Message          string
}

// SignInData contains the data of the sign in to account
type SignInData struct {
	User string
	Pass string
}

// SignInResult contains result sending sign in data
type SignInResult struct {
	Ok      bool
	UserMsg string
	PassMsg string
}

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

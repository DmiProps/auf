package templates

var (
	// EN contains messages in English
	EN = [...]string{
		"There is no such activation link. To get the activation link, go to the Sign Up page and enter the details for creating a new account.",
		"The activation link is no longer valid. You can resend the activation link.",
	}
)

// GetMessage returns a message by code in the user's language (only English is currently available)
func GetMessage(num int) string {

	return EN[num]

}

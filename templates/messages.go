package templates

import "fmt"

var (
	// Contains messages in English
	en = [...]string{
		/* 000 */ "There is no such activation link. To get the activation link, go to the Sign Up page and enter the details for creating a new account.",
		/* 001 */ "Dear %s, the activation link has expired. You can resend the activation link.",
		/* 002 */ "An error occurred while activating your account. Please try again later.",
		/* 003 */ "Dear %s, your account has been successfully activated!",
		/* 004 */ "To activate your account, follow the link sent to the e-mail address specified when creating your account.",
		/* 005 */ "An error occurred while resend activation link to your e-mail. Please try again later.",
		/* 006 */ "Failed to send activation link to e-mail. Please try again later.",
	}
)

// GetMessage returns a message by code in the user's language (only English is currently available)
func GetMessage(num int, a ...interface{}) string {

	if len(a) == 0 {
		return en[num]
	}
	return fmt.Sprintf(en[num], a...)

}

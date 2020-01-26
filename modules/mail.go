package modules

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/DmiProps/auf/settings"
)

// SendActivationMail sends activation e-mail with ref
func SendActivationMail(to string) {

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", settings.AppSettings.NoreplyEmail, settings.AppSettings.NoreplyPassword, settings.AppSettings.MailHost)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	message := `To: "Some User" <someuser@example.com>
From: "Other User" <otheruser@example.com>
Subject: Testing Email From Go!!

This is the message we are sending. That's it!
`

	if err := smtp.SendMail(settings.AppSettings.MailHost+":25", auth, settings.AppSettings.NoreplyEmail, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")

}

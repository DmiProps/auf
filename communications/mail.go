package communications

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"

	"github.com/DmiProps/auf/settings"
	"github.com/DmiProps/auf/types"
)

// SendActivationMail sends activation e-mail with ref
func SendActivationMail(data *types.SignUpData) {

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", settings.AppSettings.Email.NoreplyEmail, settings.AppSettings.Email.NoreplyPassword, settings.AppSettings.Email.MailHost)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	addr := settings.AppSettings.Email.MailHost + ":" + settings.AppSettings.Email.SMTPPort
	from := settings.AppSettings.Email.NoreplyEmail
	msg, err := makeMessage("activation-mail", data.User, data.ActivationRef, data.Email, from)
	if err != nil {
		log.Fatalln("Error makeMessage: ", err)
		return
	}

	if err := smtp.SendMail(addr, auth, from, []string{data.Email}, []byte(msg)); err != nil {
		log.Fatalln("Error SendActivationMail: ", err)
	}

}

func makeMessage(tmpl, toName string, activationRef string, a ...interface{}) (string, error) {

	wrap, err := ioutil.ReadFile("./templates/" + tmpl + ".wrap")
	if err != nil {
		return "", err
	}
	html, err := ioutil.ReadFile("./templates/" + tmpl + ".html")
	if err != nil {
		return "", err
	}

	htmlString := strings.ReplaceAll(string(html), "{User}", toName)
	htmlString = strings.ReplaceAll(htmlString, "{ActivationRef}", activationRef)
	htmlString = strings.ReplaceAll(htmlString, "{Host}", settings.AppSettings.Host)

	msg := fmt.Sprintf(string(wrap), a...)

	return msg + "\n\n" + htmlString, nil

}

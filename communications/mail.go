package communications

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"

	"github.com/DmiProps/auf/settings"
	"github.com/DmiProps/auf/types"
)

// SendActivationMail sends activation e-mail with link
func SendActivationMail(data *types.SignUpData) error {

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", settings.AppSettings.Email.NoreplyEmail, settings.AppSettings.Email.NoreplyPassword, settings.AppSettings.Email.MailHost)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	addr := settings.AppSettings.Email.MailHost + ":" + settings.AppSettings.Email.SMTPPort
	from := settings.AppSettings.Email.NoreplyEmail
	msg, err := makeMessage("activation-mail", data.User, data.ActivationLink, data.Email, from)
	if err != nil {
		return fmt.Errorf("Error makeMessage: %s", err)
	}

	if err := smtp.SendMail(addr, auth, from, []string{data.Email}, []byte(msg)); err != nil {
		return fmt.Errorf("Error SendActivationMail: %s", err)
	}

	return nil

}

func makeMessage(tmpl, toName string, activationLink string, a ...interface{}) (string, error) {

	wrap, err := ioutil.ReadFile("./templates/" + tmpl + ".wrap")
	if err != nil {
		return "", err
	}
	html, err := ioutil.ReadFile("./templates/" + tmpl + ".html")
	if err != nil {
		return "", err
	}

	htmlString := strings.ReplaceAll(string(html), "{User}", toName)
	htmlString = strings.ReplaceAll(htmlString, "{ActivationLink}", activationLink)
	htmlString = strings.ReplaceAll(htmlString, "{Host}", settings.AppSettings.Host)

	//TO-DO: Delete test message
	fmt.Println("http://localhost:8080/www/activate-link.html?link=" + activationLink)

	msg := fmt.Sprintf(string(wrap), a...)

	return msg + "\n\n" + htmlString, nil

}

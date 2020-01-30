package communications

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"

	"github.com/DmiProps/auf/settings"
)

// SendActivationMail sends activation e-mail with ref
func SendActivationMail(toName, toAddr string) {

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", settings.AppSettings.Email.NoreplyEmail, settings.AppSettings.Email.NoreplyPassword, settings.AppSettings.Email.MailHost)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	addr := settings.AppSettings.Email.MailHost + ":" + settings.AppSettings.Email.SMTPPort
	from := settings.AppSettings.Email.NoreplyEmail
	msg, err := makeMessage("activation-mail", toName, toAddr, from)
	if err != nil {
		log.Fatalln("Error makeMessage: ", err)
		return
	}

	if err := smtp.SendMail(addr, auth, from, []string{toAddr}, []byte(msg)); err != nil {
		log.Fatalln("Error SendActivationMail: ", err)
	} else {
		fmt.Println("Email Sent!")
	}

}

func makeMessage(tmpl, toName string, a ...interface{}) (string, error) {

	wrap, err := ioutil.ReadFile("./templates/" + tmpl + ".wrap")
	if err != nil {
		return "", err
	}
	html, err := ioutil.ReadFile("./templates/" + tmpl + ".html")
	if err != nil {
		return "", err
	}

	htmlString := strings.Replace(string(html), "{User}", toName, -1)
	htmlString = strings.Replace(htmlString, "{Host}", settings.AppSettings.Host, -1)

	msg := fmt.Sprintf(string(wrap), a...)

	return msg + "\n\n" + htmlString, nil

}

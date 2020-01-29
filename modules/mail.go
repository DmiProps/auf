package modules

import (
	"encoding/base64"
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
	auth := smtp.PlainAuth("", settings.AppSettings.NoreplyEmail, settings.AppSettings.NoreplyPassword, settings.AppSettings.MailHost)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	addr := settings.AppSettings.MailHost + ":" + settings.AppSettings.SMTPPort
	from := settings.AppSettings.NoreplyEmail
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

	var imgBase64 string

	img, err := ioutil.ReadFile("./www/images/auth-32x32.png")
	if err == nil {
		imgBase64 = base64.StdEncoding.EncodeToString([]byte(img))
	}

	wrap, err := ioutil.ReadFile("./templates/" + tmpl + ".wrap")
	if err != nil {
		return "", err
	}
	html, err := ioutil.ReadFile("./templates/" + tmpl + ".html")
	if err != nil {
		return "", err
	}

	htmlString := strings.Replace(string(html), "{base64}", imgBase64, -1)
	htmlString = strings.Replace(htmlString, "{user}", toName, -1)

	msg := fmt.Sprintf(string(wrap), a...)

	return msg + "\n\n" + htmlString, nil

}

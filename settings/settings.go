package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Settings with params fo rauth framework
type Settings struct {
	// Main settings
	Host string

	// E-mail settings
	MailHost        string
	SMTPPort        string
	NoreplyEmail    string
	NoreplyPassword string

	// Database settings
	DbConnection string
}

var (
	// AppSettings consists application setting from json file
	AppSettings Settings
)

// ReadSettings read settings from json file
func ReadSettings() {

	// Open our jsonFile
	jsonFile, err := os.Open("./settings/settings.json")

	// If we os.Open returns an error then handle it
	if err == nil {

		// Defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		// Unmarshal our byteArray which contains our
		// jsonFile's content into 'AppSettings' which we defined above
		json.Unmarshal(byteValue, &AppSettings)

	}

	// If the environment variables are set, then take the settings from them
	// Main settings
	a := os.Getenv("Host")
	if a != "" {
		AppSettings.Host = a
	}

	// E-mail settings
	a = os.Getenv("MailHost")
	if a != "" {
		AppSettings.MailHost = a
	}
	a = os.Getenv("SMTPPort")
	if a != "" {
		AppSettings.SMTPPort = a
	}
	a = os.Getenv("NoreplyEmail")
	if a != "" {
		AppSettings.NoreplyEmail = a
	}
	a = os.Getenv("NoreplyPassword")
	if a != "" {
		AppSettings.NoreplyPassword = a
	}

	// Database settings
	a = os.Getenv("DbConnection")
	if a != "" {
		AppSettings.DbConnection = a
	}

}

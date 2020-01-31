package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// EmailSettings with params for e-mail communications
type EmailSettings struct {
	MailHost        string
	SMTPPort        string
	NoreplyEmail    string
	NoreplyPassword string
}

// DatabaseSettings with params for database connection
type DatabaseSettings struct {
	DbConnection string
}

// Settings with params for auth framework
type Settings struct {
	Host     string
	Port     string
	Email    EmailSettings
	Database DatabaseSettings
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
	a = os.Getenv("Port")
	if a != "" {
		AppSettings.Port = a
	}

	// E-mail settings
	a = os.Getenv("MailHost")
	if a != "" {
		AppSettings.Email.MailHost = a
	}
	a = os.Getenv("SMTPPort")
	if a != "" {
		AppSettings.Email.SMTPPort = a
	}
	a = os.Getenv("NoreplyEmail")
	if a != "" {
		AppSettings.Email.NoreplyEmail = a
	}
	a = os.Getenv("NoreplyPassword")
	if a != "" {
		AppSettings.Email.NoreplyPassword = a
	}

	// Database settings
	a = os.Getenv("DATABASE_URL")
	if a != "" {
		AppSettings.Database.DbConnection = a
	}

}

package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Settings with params fo rauth framework
type Settings struct {
	MailHost        string
	NoreplyEmail    string
	NoreplyPassword string
}

var (
	// AppSettings consists application setting from json file
	AppSettings Settings
)

// ReadSettings read settings from json file
func ReadSettings() {

	// Open our jsonFile
	jsonFile, err := os.Open("/settings/settings.json")

	// If we os.Open returns an error then handle it
	if err != nil {
		return
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Unmarshal our byteArray which contains our
	// jsonFile's content into 'AppSettings' which we defined above
	json.Unmarshal(byteValue, &AppSettings)

}

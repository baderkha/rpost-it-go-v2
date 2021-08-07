package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// parseJsonFromFile : takes the path for the json string and the config , you want to cast your object to
// returns back an error if file not found or casting issues
func ParseJsonFromFile(filePath string, model interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, model)
	return nil
}

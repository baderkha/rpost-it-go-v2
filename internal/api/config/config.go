package config

import "rpost-it-go/pkg/util/json"

const (
	DefaultJSONPath = "env.json"
)

// pointer to save memory
var config *JSONEnvConfig

// JSONEnvConfig : json environment configuration for the api
type JSONEnvConfig struct {
	Port string `json:"port"`
	Auth struct {
		CookieName         string `json:"cookieName"`
		CookieDomainPrefix string `json:"cookieDomainPrefix"`
	} `json:"auth"`
	DBConfig DBConfig `json:"dbConfig"`
}

func GetInstance() *JSONEnvConfig {
	if config == nil {
		config = &JSONEnvConfig{}
		err := json.ParseJsonFromFile(DefaultJSONPath, config)
		if err != nil {
			panic(err.Error()) // should not happen , bring down the api
		}
		config.DBConfig.init() // get secrets if needed and replace the sensitive info
	}
	return config
}

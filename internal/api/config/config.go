package config

import "rpost-it-go/pkg/util/json"

const (
	DefaultJSONPath = "env.json"
)

// pointer to save memory
var config *JSONEnvConfig

// JSONEnvConfig : json environment configuration for the api
type JSONEnvConfig struct {
	Port       string   `json:"port"`
	Auth       Auth     `json:"auth"`
	DBConfig   DBConfig `json:"dbConfig"`
	MailConfig struct {
		From   string `json:"from"`
		Region string `json:"region"`
	} `json:"mailConfig"`
}

type Auth struct {
	CookieName         string       `json:"cookieName"`
	CookieDomainPrefix string       `json:"cookieDomainPrefix"`
	Verification       Verification `json:"verification"` // options about verification
}

func Get() *JSONEnvConfig {
	if config == nil {
		config = &JSONEnvConfig{}
		err := json.ParseJsonFromFile(DefaultJSONPath, config)
		if err != nil {
			panic(err.Error()) // should not happen , bring down the api
		}
		config.DBConfig.init()          // get secrets if needed and replace the sensitive info
		config.Auth.Verification.init() // get secrets if needed and replace the sensitive info
	}
	return config
}

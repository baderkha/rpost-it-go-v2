package config

import "rpost-it-go/pkg/util/secretmanager"

type DBConfig struct {
	SecretName  string   `json:"secretName"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	EndPoint    string   `json:"endpoint"`
	Port        string   `json:"port"`
	Dialect     string   `json:"dialect"`
	DBName      string   `json:"dbName"`
	QueryParams []string `json:"queryParams"`
}

// init : initalize the auth , incase it needs to fetch the secret somewhere safe
func (d *DBConfig) init() {
	if d.SecretName != "" {
		var secretConfigDb DBConfig
		err := secretmanager.New().GetSecret(d.SecretName, &secretConfigDb)
		if err != nil {
			panic(err.Error())
		}
		d.Username = secretConfigDb.Username
		d.Password = secretConfigDb.Password
	}
}

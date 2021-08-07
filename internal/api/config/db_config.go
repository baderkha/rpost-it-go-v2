package config

import (
	"fmt"
	"rpost-it-go/pkg/util/secretmanager"

	"github.com/davecgh/go-spew/spew"
)

type DBConfig struct {
	SecretName  string            `json:"secretName"`
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	EndPoint    string            `json:"endpoint"`
	Port        string            `json:"port"`
	Dialect     string            `json:"dialect"`
	DBName      string            `json:"dbName"`
	QueryParams map[string]string `json:"queryParams"`
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

func (d *DBConfig) GetDSN() string {
	queryParams := "?"
	for key, param := range d.QueryParams {
		queryParams += fmt.Sprintf("%s=%s", key, param) + "&"
	}
	spew.Dump(fmt.Sprintf("mysql://%s:%s@(%s:%s)/%s%s", d.Username, d.Password, d.EndPoint, d.Port, d.DBName, queryParams))
	return fmt.Sprintf("%s:%s@(%s:%s)/%s%s", d.Username, d.Password, d.EndPoint, d.Port, d.DBName, queryParams)
}

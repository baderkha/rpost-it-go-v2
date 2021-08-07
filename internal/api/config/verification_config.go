package config

import "rpost-it-go/pkg/util/secretmanager"

// Verification : verification config for generating tokens
type Verification struct {
	SecretName string `json:"secretName"`
	HashSecret string `json:"hashedSecret"`
	Issuer     string `json:"issuer"`
	Link       string `json:"verificationLink"`
}

func (v *Verification) init() {
	if v.SecretName != "" {
		var secretConfigDb Verification
		err := secretmanager.New().GetSecret(v.SecretName, &secretConfigDb)
		if err != nil {
			panic(err.Error())
		}
		v.HashSecret = secretConfigDb.HashSecret
	}
}

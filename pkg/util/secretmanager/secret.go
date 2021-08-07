package secretmanager

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// interface that implements a secret manager
var secretManager ISecretManager

// ISecretManager : a safe place your server has access
// to to bring you secrets like a keychain
type ISecretManager interface {
	// GetSecret : fetch the secert and cast it to the model you provide
	GetSecret(secretId string, model interface{}) error
	// GetSecretString : fetch the secret without casting to the model
	GetSecretString(secretId string) (string, error)
}

// NewSecretManager : factory that returns default secrets manager
func New() ISecretManager {
	if secretManager == nil {
		s := session.Must(session.NewSession())
		sm := secretsmanager.New(s)
		secretManager = &AwsSecretManager{
			driver: sm,
		}
	}
	return secretManager
}

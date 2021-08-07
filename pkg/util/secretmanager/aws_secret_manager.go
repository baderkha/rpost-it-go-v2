package secretmanager

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

// AwsSecretManager : Implementation of the secret manager using aws api
type AwsSecretManager struct {
	driver secretsmanageriface.SecretsManagerAPI
}

// GetSecret : fetch secret as struct
func (a *AwsSecretManager) GetSecret(secretId string, model interface{}) error {
	secret, err := a.GetSecretString(secretId)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(secret), model)
	return err
}

// GetSecretString : get the secret as a string
func (a *AwsSecretManager) GetSecretString(secretId string) (string, error) {
	out, err := a.driver.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	})
	if err != nil {
		return "", err
	}
	secret := *out.SecretString
	return secret, nil
}

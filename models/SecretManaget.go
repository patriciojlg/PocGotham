package models

import (
	settings "PocGotham/settings"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretsPayload struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type SecretManager struct {
	SecretManagerClient *secretsmanager.SecretsManager
	SecretManagerId     string
}

func getSecretManagerClient() *secretsmanager.SecretsManager {
	// Crear una sesión de AWS con la región especificada
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(settings.AWS_REGION),
	}))

	secretsManagerClient := secretsmanager.New(sess)
	return secretsManagerClient
}
func NewSecretManager() *SecretManager {

	return &SecretManager{

		SecretManagerId:     settings.SECRET_NAME,
		SecretManagerClient: getSecretManagerClient(),
	}
}
func (sm *SecretManager) GetSecret() SecretsPayload {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &sm.SecretManagerId,
	}
	result, err := sm.SecretManagerClient.GetSecretValue(input)
	if err != nil {
		return SecretsPayload{}
	}
	var secretsPayload SecretsPayload
	json.Unmarshal([]byte(*result.SecretString), &secretsPayload)
	return secretsPayload
}

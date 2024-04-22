package models

import (
	settings "PocGotham/settings"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func getClientAWSCognito() *cognitoidentityprovider.CognitoIdentityProvider {

	sm := NewSecretManager()
	secretPayload := sm.GetSecret()

	creds := aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(secretPayload.ClientId, secretPayload.ClientSecret, ""))

	// Configurar la sesión de AWS con las credenciales y la región
	sess, err := session.NewSession(creds.WithRegion(settings.AWS_REGION))
	if err != nil {
		fmt.Println("Error al crear la sesión de AWS:", err)
		os.Exit(1)
	}

	// Crear un cliente de AWS Cognito Identity Provider
	cognitoClient := cognitoidentityprovider.New(sess)

	return cognitoClient

}

type Cognito struct {
	UserPoolID    string
	ClientCognito *cognitoidentityprovider.CognitoIdentityProvider
}

func NewCognito() *Cognito {
	return &Cognito{
		UserPoolID:    settings.USER_POOL_ID,
		ClientCognito: getClientAWSCognito(),
	}
}
func createCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	// Configurar las credenciales de AWS
	sm := NewSecretManager()
	secretPayload := sm.GetSecret()

	creds := aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(secretPayload.ClientId, secretPayload.ClientSecret, ""))

	// Configurar la sesión de AWS con las credenciales y la región
	sess, err := session.NewSession(creds.WithRegion(settings.AWS_REGION))
	if err != nil {
		fmt.Println("Error al crear la sesión de AWS:", err)
		os.Exit(1)
	}

	// Crear un cliente de AWS Cognito Identity Provider
	cognitoClient := cognitoidentityprovider.New(sess)

	return cognitoClient
}
func (c *Cognito) CreateUser(username, password string) error {
	// Parametros para crear un nuevo usuario
	createUserInput := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        aws.String(c.UserPoolID),
		Username:          aws.String(username),
		TemporaryPassword: aws.String(password),
		MessageAction:     aws.String("SUPPRESS"), // Acción para el mensaje de envío de contraseña
	}
	clientCognito := createCognitoClient()
	// Crear el nuevo usuario en AWS Cognito
	_, err := clientCognito.AdminCreateUser(createUserInput)
	if err != nil {
		fmt.Println("Error al crear el nuevo usuario:", err)
		return err
	}

	fmt.Println("Nuevo usuario creado exitosamente.")
	return nil
}

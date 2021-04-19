package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

//GetPublicKey gets the public key of an account
func GetPublicKey() {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/notifications", //Specify url or use default
	})

	getPublicKeyOptions := service.NewGetPublicKeyOptions(accountID)
	result, resp, operationErr := service.GetPublicKey(getPublicKeyOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		fmt.Println("Failed to get public key: ", operationErr)
		return
	}

	fmt.Println(*result.PublicKey)
}

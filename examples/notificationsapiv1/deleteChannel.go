package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

//DeleteChannel deletes a channel
func DeleteChannel() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/notifications", //Specify url or use default
	})

	channelID := "af9affd0-8ebf-11ea-b6bb-77a55a4ff9b0"

	var deleteOptions = service.NewDeleteNotificationChannelOptions(accountID, channelID)
	result, response, err := service.DeleteNotificationChannel(deleteOptions)
	if err != nil && response.StatusCode != 200 {
		fmt.Println(response.Result)
		fmt.Println("Failed to delete channel: ", err)
		return
	}

	fmt.Println(*result.ChannelID)
	fmt.Println(*result.Message)
}

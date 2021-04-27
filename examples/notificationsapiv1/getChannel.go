package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsv1"
)

//GetChannel gets a channel
func GetChannel() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := notificationsv1.NewNotificationsV1(&notificationsv1.NotificationsV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/notifications", //Specify url or use default
	})

	channelID := "3b95c100-8ebb-11ea-a009-0fa4c7a4acaf"

	getChannelOptions := service.NewGetNotificationChannelOptions(accountID, channelID)
	result, resp, operationErr := service.GetNotificationChannel(getChannelOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		fmt.Println("Failed to get channel: ", operationErr)
		return
	}

	fmt.Println(*result.Channel.ChannelID)
	fmt.Println(*result.Channel.Name)
}

package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsv1"
)

//UpdateChannel updates a channel
func UpdateChannel() {
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

	channelID := "af9affd0-8ebf-11ea-b6bb-77a55a4ff9b0"
	channelName := "sdkTest_channel_all_exmpl"

	updateOptions := service.NewUpdateNotificationChannelOptions(accountID, channelID, channelName, "Webhook", "https://update.ss")
	updateOptions.SetHeaders(headers)
	updateOptions.SetDescription("updated from go")
	updateOptions.SetType("Webhook")
	updateOptions.SetSeverity([]string{"low"})
	updateOptions.SetEnabled(false)

	result, response, err := service.UpdateNotificationChannel(updateOptions)
	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to update channel: ", err)
		return
	}
	fmt.Println(*result.ChannelID)
	fmt.Println(*result.StatusCode)

}

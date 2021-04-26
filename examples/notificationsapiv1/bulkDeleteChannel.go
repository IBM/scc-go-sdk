package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsv1"
)

//DeleteChannels bulk deletes channels
func DeleteChannels() {
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

	body := []string{"6b64f0a0-8eba-11ea-888a-9d749f970504", "87eb0890-8eba-11ea-888a-9d749f970504"}

	var deleteOptions = service.NewDeleteNotificationChannelsOptions(accountID, body)
	result, response, err := service.DeleteNotificationChannels(deleteOptions)
	if err != nil && response.StatusCode != 200 {
		fmt.Println(response.Result)
		fmt.Println("Failed to delete channel: ", err)
		return
	}

	fmt.Println(*result.Message)
}

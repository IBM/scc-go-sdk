package main

import (
	"encoding/json"
	"fmt"
	"github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"io/ioutil"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/notificationsv1"
)

func CreateCollector() {
	apiKey := "DjsEbdqjIwuP9bfTyGATAuJ9u55dsMbVNvJ8cVWdzoxz"
	url := "https://iam.test.cloud.ibm.com/oidc/token"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
		Authenticator: authenticator,
		URL: "https://asap-dev.compliance.test.cloud.ibm.com/posture/v1/collectors?account_id={{account_id}}", //Specify url or use default
	})

	collector_name := "jason-test-collector-01"
	collector_description := ""
	channelType := "Webhook"
	severity := []string{notificationsv1.CreateNotificationChannelOptionsSeverityCriticalConst, notificationsv1.CreateNotificationChannelOptionsSeverityHighConst, notificationsv1.CreateNotificationChannelOptionsSeverityLowConst}

	var alertSource []notificationsv1.NotificationChannelAlertSourceItem
	source, _ := service.NewNotificationChannelAlertSourceItem("ATA")
	source.FindingTypes = []string{"appid", "cos", "iks"}
	alertSource = append(alertSource, *source)

	createOptions := service.NewCreateNotificationChannelOptions(accountID, channelName, channelType, endpoint)

	//Below set of calls are not required. A channel can be created with just channelName, channelType, endpoint. Rest will saaume default value.
	createOptions.SetHeaders(headers)
	createOptions.SetSeverity(severity)
	createOptions.SetEnabled(true)
	createOptions.SetDescription("this is a test")
	createOptions.SetAlertSource(alertSource)

	result, response, err := service.CreateNotificationChannel(createOptions)
	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create channel: ", err)
		return
	}
	fmt.Println(*result.ChannelID)
	fmt.Println(*result.StatusCode)

}

// +build integration

package notificationsapiv1_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

var apiKey = os.Getenv("apiKey")
var accountID = os.Getenv("accountID")
var URL = os.Getenv("URL")
var inputEnvPath = "../testInput/env"
var inputFilePath = "../testInput/json"
var notificationsServiceURL = os.Getenv("notificationsServiceURL")

var (
	service     *notificationsapiv1.NotificationsApiV1
	shouldSkip  bool = false
	err         error
	errConflict error = errors.New("Conflict")
)

func shouldSkipTest(t *testing.T) {
	if !shouldSkip {
		t.Skip("External configuration is not available, skipping...")
	}
}

func createChannelHelper(t *testing.T, path string) (*notificationsapiv1.CreateChannelsResponse, *notificationsapiv1.CreateNotificationChannelOptions, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	query, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("Failed to open file: ", err)
	}
	var createNotificationChannelOptions *notificationsapiv1.CreateNotificationChannelOptions
	json.Unmarshal([]byte(query), &createNotificationChannelOptions)
	createNotificationChannelOptions.SetHeaders(headers)
	now := time.Now()
	timeStamp := now.UnixNano()
	createNotificationChannelOptions.SetName(strconv.FormatInt(timeStamp, 10))
	createNotificationChannelOptions.AccountID = &accountID
	result, resp, operationErr := service.CreateNotificationChannel(createNotificationChannelOptions)
	if operationErr != nil && resp.StatusCode != 200 {
		t.Log("Failed to create channel: ", operationErr)
	}
	return result, createNotificationChannelOptions, operationErr
}

func deleteChannelHelper(t *testing.T, acountId string, channelId string) {
	var deleteOptions = service.NewDeleteNotificationChannelOptions(acountId, channelId)
	_, response, err := service.DeleteNotificationChannel(deleteOptions)
	if err != nil || response.StatusCode != 200 {
		t.Fatal("Failed to delete channel: ", err)
	}
}

func TestServiceSetupWithExternalConfig(t *testing.T) {
	externalConfigFile := inputEnvPath + "/notificationapiv1.env"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    URL,
	}
	err = godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Failed to load env vars: ", err)
	}

	testService, err := notificationsapiv1.NewNotificationsApiV1UsingExternalConfig(&notificationsapiv1.NotificationsApiV1Options{
		Authenticator: authenticator,
	})
	if err != nil {
		fmt.Println("Failed to load external config: ", err)
	}

	assert.NotNil(t, testService)
}

func TestServiceSetupAndSetUrl(t *testing.T) {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    URL,
	}
	testService, err := notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
		Authenticator: authenticator,
	})
	if testService == nil {
		t.Fatal("Expected service to not be nil, but got: ", service)
	}

	testService.SetServiceURL("https://dev-dallas.secadvisor.test.cloud.ibm.com/notifications")

	assert.Equal(t, testService.Service.Options.URL, "https://dev-dallas.secadvisor.test.cloud.ibm.com/notifications")
	if err != nil {
		t.Fatal("expected testServiceErr to be nil, but got: ", err)
	}

}

func TestServiceSetup(t *testing.T) { //This is required for tests that follow
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    URL,
	}
	service, err = notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
		Authenticator: authenticator,
		URL:           notificationsServiceURL,
	})
	if service == nil {
		t.Fatal("Expected service to not be nil, but got: ", service)
	}

	if err != nil {
		t.Fatal("expected testServiceErr to be nil, but got: ", err)
	}

}

func TestCreateNotificationChannelWithRequired(t *testing.T) {
	fmt.Println("creating new channel with required only")
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateChannelWithRequiredUsingStruct(t *testing.T) {
	fmt.Println("creating new channel with required only using struct instantiation")
	createChannelOptions := service.NewCreateNotificationChannelOptions(accountID, "sdkTest_channel_struct", "Webhook", "https://example.com")
	result, _, _ := service.CreateNotificationChannel(createChannelOptions)
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateChannelWithAllUsingFunctions(t *testing.T) {
	fmt.Println("creating new channel with required only using struct instantiation")
	var createChannelOptions notificationsapiv1.CreateNotificationChannelOptions

	createChannelOptions.SetName("sdkTest_channel_func")
	createChannelOptions.SetType(notificationsapiv1.CreateNotificationChannelOptions_Type_Webhook)
	createChannelOptions.SetSeverity([]string{notificationsapiv1.CreateNotificationChannelOptions_Severity_High, notificationsapiv1.CreateNotificationChannelOptions_Severity_Low})
	createChannelOptions.SetEnabled(false)
	createChannelOptions.SetDescription("description from go sdk")
	createChannelOptions.SetAccountID(accountID)
	createChannelOptions.SetEndpoint("https://example.com")
	alertSource, _ := service.NewNotificationChannelAlertSourceItem("ALL")
	createChannelOptions.SetAlertSource([]notificationsapiv1.NotificationChannelAlertSourceItem{*alertSource})

	result, _, _ := service.CreateNotificationChannel(&createChannelOptions)
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateNotificationChannelWithSeverity(t *testing.T) {
	fmt.Println("creating new channel with severity")
	result, createChannelOptions, _ := createChannelHelper(t, inputFilePath+"/channel_with_severity.json")
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateNotificationChannelWithAlertSource(t *testing.T) {
	fmt.Println("creating new channel with alert source")
	result, createChannelOptions, _ := createChannelHelper(t, inputFilePath+"/channel_with_alert_source.json")
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateNotificationChannelWithAllKeys(t *testing.T) {
	fmt.Println("creating new channel with all")
	result, createChannelOptions, _ := createChannelHelper(t, inputFilePath+"/channel_with_all.json")
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))
	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
}

func TestCreateChannelFailure(t *testing.T) {
	fmt.Println("Fails to create channel unknown alert source")
	result, _, operationErr := createChannelHelper(t, inputFilePath+"/channel_with_invalid_data.json")
	fmt.Println("Failed to create channel: ", operationErr)
	assert.NotNil(t, operationErr)
	assert.Nil(t, result)
	assert.NotNil(t, operationErr.Error(), "Alert source providers or finding types not found.")
}

func TestCreateChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to create channel nil options")

	result, resp, operationErr := service.CreateNotificationChannel(nil)

	t.Log("Failed to create channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "createNotificationChannelOptions cannot be nil")
}

func TestCreateChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to create channel with invalid options")
	channelName := "name"
	channelType := "type"
	var createNotificationChannelOptions = notificationsapiv1.CreateNotificationChannelOptions{Name: &channelName, Type: &channelType}
	result, resp, operationErr := service.CreateNotificationChannel(&createNotificationChannelOptions)

	t.Log("Failed to create channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestDeleteNotificationChannel(t *testing.T) {
	fmt.Println("creating new channel for delete test...")
	result, createChannelOptions, _ := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("Deleting channel....")
		var deleteOptions = service.NewDeleteNotificationChannelOptions(*createChannelOptions.AccountID, *result.ChannelID)
		delResult, response, err := service.DeleteNotificationChannel(deleteOptions)
		if err != nil {
			t.Fatal("Failed to delete channel: ", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, *delResult.Message, "Success")
		assert.Equal(t, *delResult.ChannelID, *result.ChannelID)
	}
}

func TestDeleteNotificationChannelUsingFunctions(t *testing.T) {
	fmt.Println("creating new channel for delete test...")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	result, _, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("Deleting channel....")
		var deleteOptions notificationsapiv1.DeleteNotificationChannelOptions
		deleteOptions.SetAccountID(accountID)
		deleteOptions.SetChannelID(*result.ChannelID)
		deleteOptions.SetHeaders(headers)

		delResult, response, err := service.DeleteNotificationChannel(&deleteOptions)

		if err != nil {
			t.Fatal("Failed to delete channel: ", err)
		}

		assert.Nil(t, err)
		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t, *delResult.Message, "Success")
		assert.Equal(t, *delResult.ChannelID, *result.ChannelID)
	}
}

func TestDeleteChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to delete channel nil options")

	result, resp, operationErr := service.DeleteNotificationChannel(nil)

	t.Log("Failed to create channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "deleteNotificationChannelOptions cannot be nil")
}

func TestDeleteChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to delete channel with invalid options")
	channelId := "id"
	var deleteNotificationChannelOptions = notificationsapiv1.DeleteNotificationChannelOptions{ChannelID: &channelId}
	result, resp, operationErr := service.DeleteNotificationChannel(&deleteNotificationChannelOptions)

	t.Log("Failed to create channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestGetNotificationChannelUsingStruct(t *testing.T) {
	fmt.Println("creating new channel for get test...")
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("Created getting channel....")
		getChannelOptions := service.NewGetNotificationChannelOptions(accountID, *result.ChannelID)
		fetchResult, resp, operationErr := service.GetNotificationChannel(getChannelOptions)
		if operationErr != nil {
			fmt.Println(resp.Result)
			t.Fatal("Failed to get channel: ", operationErr)
		}
		assert.NotNil(t, fetchResult)
		assert.Equal(t, *fetchResult.Channel.ChannelID, *result.ChannelID)
		assert.Equal(t, *fetchResult.Channel.Name, *createChannelOptions.Name)
		assert.Equal(t, *fetchResult.Channel.Type, *createChannelOptions.Type)
		assert.Equal(t, *fetchResult.Channel.AlertSource[0].ProviderName, "ALL")
		assert.Equal(t, *fetchResult.Channel.Severity.High, true)

		fmt.Println("Deleting channel....")
		deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
	}
}

func TestGetNotificationChannelUsingFunctions(t *testing.T) {
	fmt.Println("creating new channel for get test...")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println(" getting channel....")
		var getChannelOptions notificationsapiv1.GetNotificationChannelOptions
		getChannelOptions.SetAccountID(accountID)
		getChannelOptions.SetChannelID(*result.ChannelID)
		getChannelOptions.SetHeaders(headers)
		fetchResult, _, operationErr := service.GetNotificationChannel(&getChannelOptions)
		if operationErr != nil {
			t.Fatal("Failed to get channel: ", operationErr)
		}
		assert.NotNil(t, fetchResult)
		assert.Equal(t, *fetchResult.Channel.ChannelID, *result.ChannelID)
		assert.Equal(t, *fetchResult.Channel.Name, *createChannelOptions.Name)
		assert.Equal(t, *fetchResult.Channel.Type, *createChannelOptions.Type)
		assert.Equal(t, *fetchResult.Channel.AlertSource[0].ProviderName, "ALL")
		assert.Equal(t, *fetchResult.Channel.Severity.High, true)

		fmt.Println("Deleting channel....")
		deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
	}
}

func TestGetChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to get channel nil options")

	result, resp, operationErr := service.GetNotificationChannel(nil)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "getNotificationChannelOptions cannot be nil")
}

func TestGetChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to get channel with invalid options")
	channelId := "id"
	var getChannelOptions = notificationsapiv1.GetNotificationChannelOptions{ChannelID: &channelId}
	result, resp, operationErr := service.GetNotificationChannel(&getChannelOptions)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestGetChannelFailureWithInvalidId(t *testing.T) {
	fmt.Println("Fails to get channel with invalid options")
	channelId := "id"
	var getChannelOptions = notificationsapiv1.GetNotificationChannelOptions{ChannelID: &channelId, AccountID: &accountID}
	result, resp, operationErr := service.GetNotificationChannel(&getChannelOptions)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Equal(t, resp.StatusCode, 404)
	assert.NotNil(t, operationErr)
	assert.Equal(t, operationErr.Error(), "Channels not found.")

}

func TestTestNotificationChannelValid(t *testing.T) {
	fmt.Println("creating new channel for test...")
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("testing webhook....")
		testChannelOptions := service.NewTestNotificationChannelOptions(accountID, *result.ChannelID)
		testResult, _, operationErr := service.TestNotificationChannel(testChannelOptions)
		if operationErr != nil {
			t.Fatal("Failed to test channel: ", operationErr)
		}
		assert.NotNil(t, testResult)
		assert.NotNil(t, *testResult.Test)

		fmt.Println("Deleting channel....")
		deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
	}
}

func TestTestNotificationChannelInvalid(t *testing.T) {
	fmt.Println("creating new channel for test...")
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_all.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("testing webhook....")
		testChannelOptions := service.NewTestNotificationChannelOptions(accountID, *result.ChannelID)
		_, resp, operationErr := service.TestNotificationChannel(testChannelOptions)

		assert.NotNil(t, operationErr)
		assert.Equal(t, resp.StatusCode, 503)
		assert.Equal(t, operationErr.Error(), "The webhook URL specified for the channel is unavailable")

		fmt.Println("Deleting channel....")
		deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
	}
}

func TestTestNotificationChannelUsingFunctions(t *testing.T) {
	fmt.Println("creating new channel for test...")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, *result.StatusCode, int64(200))

	if result != nil {
		fmt.Println("Created new channel....")
		fmt.Println("Channel id: ", *result.ChannelID)
		fmt.Println("testing webhook....")
		var testChannelOptions notificationsapiv1.TestNotificationChannelOptions
		testChannelOptions.SetAccountID(accountID)
		testChannelOptions.SetChannelID(*result.ChannelID)
		testChannelOptions.SetHeaders(headers)
		testResult, _, operationErr := service.TestNotificationChannel(&testChannelOptions)
		if operationErr != nil {
			t.Fatal("Failed to test channel: ", operationErr)
		}
		assert.NotNil(t, testResult)
		assert.NotNil(t, *testResult.Test)

		fmt.Println("Deleting channel....")
		deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)
	}
}

func TestTestChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to test channel with nil options")

	result, resp, operationErr := service.TestNotificationChannel(nil)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "testNotificationChannelOptions cannot be nil")
}

func TestTestChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to test channel with invalid options")
	channelId := "id"
	var getChannelOptions = notificationsapiv1.TestNotificationChannelOptions{ChannelID: &channelId}
	result, resp, operationErr := service.TestNotificationChannel(&getChannelOptions)

	t.Log("Failed to test channel: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestTestChannelFailureWithInvalidId(t *testing.T) {
	fmt.Println("Fails to test channel with invalid options")
	channelId := "id"
	var getChannelOptions = notificationsapiv1.TestNotificationChannelOptions{ChannelID: &channelId, AccountID: &accountID}
	result, resp, operationErr := service.TestNotificationChannel(&getChannelOptions)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Equal(t, resp.StatusCode, 404)
	assert.NotNil(t, operationErr)
	assert.Equal(t, operationErr.Error(), "Channels not found.")

}

func TestGetPublicKeySuccess(t *testing.T) {
	fmt.Println("Getting public key....")
	getPublicKeyOptions := service.NewGetPublicKeyOptions(accountID)
	result, _, operationErr := service.GetPublicKey(getPublicKeyOptions)
	if operationErr != nil {
		t.Fatal("Failed to get public key: ", operationErr)
	}
	assert.NotNil(t, result)
	assert.NotNil(t, *result.PublicKey)
}

func TestGetPublicKeySuccessUsingFunctions(t *testing.T) {
	fmt.Println("Getting public key....")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	var getPublicKeyOptions notificationsapiv1.GetPublicKeyOptions
	getPublicKeyOptions.SetAccountID(accountID)
	getPublicKeyOptions.SetHeaders(headers)
	result, _, operationErr := service.GetPublicKey(&getPublicKeyOptions)
	if operationErr != nil {
		t.Fatal("Failed to get public key: ", operationErr)
	}
	assert.NotNil(t, result)
	assert.NotNil(t, *result.PublicKey)
}

func TestGetPublicKeyFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to get key with nil options")

	result, resp, operationErr := service.GetPublicKey(nil)

	t.Log("Failed to get key: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "getPublicKeyOptions cannot be nil")
}

func TestGetPublicKeyFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to get key with invalid options")
	var getKeyOptions = notificationsapiv1.GetPublicKeyOptions{}
	result, resp, operationErr := service.GetPublicKey(&getKeyOptions)

	t.Log("Failed to get key: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestGetPublicKeyFailureWithInvalidAccId(t *testing.T) {
	fmt.Println("Fails to test channel with invalid options")
	accId := "id"
	var getKeyOptions = notificationsapiv1.GetPublicKeyOptions{AccountID: &accId}
	result, resp, operationErr := service.GetPublicKey(&getKeyOptions)

	t.Log("Failed to get channel: ", operationErr)
	assert.Nil(t, result)
	assert.Equal(t, resp.StatusCode, 403)
	assert.NotNil(t, operationErr)
	assert.Equal(t, operationErr.Error(), "Forbidden")

}

func TestGetNotificationChannelsUsingStruct(t *testing.T) {
	fmt.Println("creating new channel for get test...")
	result1, createChannelOptions1, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	result2, createChannelOptions2, err := createChannelHelper(t, inputFilePath+"/channel_with_severity.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.Equal(t, *result1.StatusCode, int64(200))
	assert.Equal(t, *result2.StatusCode, int64(200))

	fmt.Println("Created new channels....")
	fmt.Println("getting channels....")
	listChannelOptions := service.NewListAllChannelsOptions(accountID)
	fetchResult, resp, operationErr := service.ListAllChannels(listChannelOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		t.Fatal("Failed to get channels: ", operationErr)
	}

	count := 0

	for i := 0; i < len(fetchResult.Channels); i++ {
		if *fetchResult.Channels[i].Name == *createChannelOptions2.Name || *fetchResult.Channels[i].Name == *createChannelOptions1.Name {
			count++
		}
	}

	assert.Equal(t, count, 2)

	fmt.Println("Deleting channels....")
	deleteChannelHelper(t, *createChannelOptions1.AccountID, *result1.ChannelID)
	deleteChannelHelper(t, *createChannelOptions2.AccountID, *result2.ChannelID)
}

func TestGetNotificationChannelsUsingFunctions(t *testing.T) {
	fmt.Println("creating new channels for getlist test...")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	result1, createChannelOptions1, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	result2, createChannelOptions2, err := createChannelHelper(t, inputFilePath+"/channel_with_severity.json")
	if err != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.Equal(t, *result1.StatusCode, int64(200))
	assert.Equal(t, *result2.StatusCode, int64(200))

	fmt.Println("Created new channels....")
	fmt.Println("getting channels....")
	var listChannelOptions notificationsapiv1.ListAllChannelsOptions
	listChannelOptions.SetAccountID(accountID)
	listChannelOptions.SetHeaders(headers)
	listChannelOptions.SetLimit(2)
	listChannelOptions.SetSkip(0)
	fetchResult, resp, operationErr := service.ListAllChannels(&listChannelOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		t.Fatal("Failed to get channels: ", operationErr)
	}

	count := 0

	for i := 0; i < len(fetchResult.Channels); i++ {
		if *fetchResult.Channels[i].Name == *createChannelOptions2.Name || *fetchResult.Channels[i].Name == *createChannelOptions1.Name {
			count++
		}
	}

	assert.LessOrEqual(t, count, 2)
	// assert.Equal(t, len(fetchResult.Channels), 2)

	fmt.Println("Deleting channels....")
	deleteChannelHelper(t, *createChannelOptions1.AccountID, *result1.ChannelID)
	deleteChannelHelper(t, *createChannelOptions2.AccountID, *result2.ChannelID)
}

func TestGetNotificationChannelsFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to list channels with nil options")

	result, resp, operationErr := service.ListAllChannels(nil)

	t.Log("Failed to get channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "listAllChannelsOptions cannot be nil")
}

func TestGetNotificationChannelsFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to list channels with invalid options")
	var listChannelOptions = notificationsapiv1.ListAllChannelsOptions{}
	result, resp, operationErr := service.ListAllChannels(&listChannelOptions)

	t.Log("Failed to get channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestGetNotificationChannelsFailureWithInvalidAccId(t *testing.T) {
	fmt.Println("Fails to test channel with invalid acc id")
	accId := "id"
	var listChannelOptions = notificationsapiv1.ListAllChannelsOptions{AccountID: &accId}
	result, resp, operationErr := service.ListAllChannels(&listChannelOptions)

	assert.Nil(t, result)
	assert.Equal(t, resp.StatusCode, 403)
	assert.NotNil(t, operationErr)
	assert.Equal(t, operationErr.Error(), "Forbidden")

}

func TestBulkDeleteNotificationChannel(t *testing.T) {
	fmt.Println("creating new channels for bulk delete test...")
	result1, _, err1 := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if err1 != nil {
		t.Fatal("Failed to create channel: ", err)
	}
	result2, _, err2 := createChannelHelper(t, inputFilePath+"/channel_with_severity.json")
	if err2 != nil {
		t.Fatal("Failed to create channel: ", err)
	}

	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.Equal(t, *result1.StatusCode, int64(200))
	assert.Equal(t, *result2.StatusCode, int64(200))

	fmt.Println("Created new channels....")
	fmt.Println("Deleting channels....")

	var deleteOptions = service.NewDeleteNotificationChannelsOptions(accountID, []string{*result1.ChannelID, *result2.ChannelID})
	delResult, response, err := service.DeleteNotificationChannels(deleteOptions)
	if err != nil {
		t.Fatal("Failed to delete channels: ", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, *delResult.Message, "Success")
}

func TestBulkDeleteNotificationChannelUsingFunctions(t *testing.T) {
	fmt.Println("creating new channels for bulk delete test...")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	fmt.Println("creating new channels for bulk delete test...")
	result1, _, _ := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	result2, _, _ := createChannelHelper(t, inputFilePath+"/channel_with_severity.json")
	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.Equal(t, *result1.StatusCode, int64(200))
	assert.Equal(t, *result2.StatusCode, int64(200))

	fmt.Println("Created new channels....")
	fmt.Println("Deleting channels....")
	var deleteOptions notificationsapiv1.DeleteNotificationChannelsOptions
	body := []string{*result1.ChannelID, *result2.ChannelID}
	deleteOptions.SetAccountID(accountID)
	deleteOptions.SetRequestBody(body)
	deleteOptions.SetHeaders(headers)

	delResult, response, err := service.DeleteNotificationChannels(&deleteOptions)
	if err != nil {
		t.Fatal("Failed to delete channels: ", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, *delResult.Message, "Success")
}

func TestBulkDeleteChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to bulk delete channels nil options")

	result, resp, operationErr := service.DeleteNotificationChannels(nil)

	t.Log("Failed to delete channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "deleteNotificationChannelsOptions cannot be nil")
}

func TestBulkDeleteChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to delete channels with invalid options")
	body := []string{accountID + "pridey19"}
	var deleteNotificationChannelOptions = notificationsapiv1.DeleteNotificationChannelsOptions{RequestBody: body}
	result, resp, operationErr := service.DeleteNotificationChannels(&deleteNotificationChannelOptions)

	t.Log("Failed to delete channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

func TestUpdateNotificationChannel(t *testing.T) {
	fmt.Println("creating new channel with required only")
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if result == nil {
		t.Fatal("failed to create Channel: ", err)
	}
	fmt.Println("updating channel")
	updateOptions := service.NewUpdateNotificationChannelOptions(*createChannelOptions.AccountID, *result.ChannelID, "sdkTest_channel_updated", "Webhook", "https://ss.ss")

	updateResult, response, err := service.UpdateNotificationChannel(updateOptions)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, *updateResult.ChannelID, *result.ChannelID)

	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)

}

func TestUpdateNotificationChannelWithFunction(t *testing.T) {
	fmt.Println("creating new channel with required only")
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	result, createChannelOptions, err := createChannelHelper(t, inputFilePath+"/channel_with_required_only.json")
	if result == nil {
		t.Fatal("failed to create Channel: ", err)
	}
	fmt.Println("updating channel with all keys")

	var updateOptions notificationsapiv1.UpdateNotificationChannelOptions
	updateOptions.SetAccountID(*createChannelOptions.AccountID)
	updateOptions.SetChannelID(*result.ChannelID)
	updateOptions.SetName(*createChannelOptions.Name)
	updateOptions.SetDescription("updated from go")
	updateOptions.SetType("Webhook")
	updateOptions.SetEndpoint("https://ss.ss")
	updateOptions.SetSeverity([]string{"low"})
	updateOptions.SetEnabled(false)
	alertSource, _ := service.NewNotificationChannelAlertSourceItem("ALL")
	updateOptions.SetAlertSource([]notificationsapiv1.NotificationChannelAlertSourceItem{*alertSource})
	updateOptions.SetHeaders(headers)

	updateResult, response, err := service.UpdateNotificationChannel(&updateOptions)

	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, *updateResult.ChannelID, *result.ChannelID)

	fmt.Println("cleaning up channel")
	deleteChannelHelper(t, *createChannelOptions.AccountID, *result.ChannelID)

}

func TestUpdateChannelFailureWithNilOptions(t *testing.T) {
	fmt.Println("Fails to bulk update channels nil options")

	result, resp, operationErr := service.UpdateNotificationChannel(nil)

	t.Log("Failed to delete channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Equal(t, operationErr.Error(), "updateNotificationChannelOptions cannot be nil")
}

func TestBulkUpdateChannelFailureWithInvalidOptions(t *testing.T) {
	fmt.Println("Fails to update channels with invalid options")
	var updateNotificationChannelOptions = notificationsapiv1.UpdateNotificationChannelOptions{AccountID: &accountID}
	result, resp, operationErr := service.UpdateNotificationChannel(&updateNotificationChannelOptions)

	t.Log("Failed to update channels: ", operationErr)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.NotNil(t, operationErr)
}

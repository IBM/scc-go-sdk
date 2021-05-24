package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"os"
)

func CreateCollector() {
	apiKey := os.Getenv("IAM_API_KEY")
	url := os.Getenv("IAM_APIKEY_URL")
	accountId := os.Getenv("ACCOUNT_ID")
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = os.Getenv("OAUTH_TOKEN")
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env

	}
	service, _ := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
		Authenticator: authenticator,
		URL:           "https://asap-dev.compliance.test.cloud.ibm.com/posture/v1/collectors?account_id=" + accountId, //Specify url or use default
	})

	source := service.NewCreateCollectorOptions(accountId)
	source.SetCollectorName("jason-test-collector-01")
	source.SetCollectorDescription("jason scope")
	source.SetInstallationType("installed")
	source.SetIsPublic(true)
	source.SetPassphrase("secret")

	result, response, err := service.CreateCollector(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create collector: ", err)
		return
	}
	fmt.Println(*result.CollectorID)

}

func main() {
	CreateCollector()
}

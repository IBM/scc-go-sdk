package examples

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"os"
)

func InitiateValidationScan() {
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
	service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
		Authenticator: authenticator,
		URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
	})

	source := service.NewCreateValidationOptions(accountId)
	source.SetScopeID("1380")
	source.SetProfileID("123")

	_, response, err := service.CreateValidation(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		return
	}
	fmt.Println(response.StatusCode)

}

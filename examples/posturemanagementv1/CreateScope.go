package examples

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"os"
)

func CreateScope() {
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

	source := service.NewCreateScopeOptions(accountId)
	source.SetScopeName("sample scope 01")
	source.SetScopeDescription("sample scope description")
	source.SetEnvironmentType("ibm")
	source.SetCollectorIds([]string{"1380"})
	source.SetCredentialID("cred")

	result, response, err := service.CreateScope(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		return
	}
	fmt.Println("Scope ID: ", *result.ScopeID)
	fmt.Println("Scope Description: ", *result.ScopeDescription)
	fmt.Println("Environment Type: ", *result.EnvironmentType)
	fmt.Println("Created Time: ", *result.CreatedTime)
	fmt.Println("Modified Time: ", *result.ModifiedTime)

}
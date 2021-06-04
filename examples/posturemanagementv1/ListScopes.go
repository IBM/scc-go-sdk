package examples

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"os"
)

func ListScopes(options scc.PostureManagementV1Options, scopeName string, scopeId string) bool {
	apiKey := os.Getenv("IAM_API_KEY")
	url := os.Getenv("IAM_APIKEY_URL")
	accountId := os.Getenv("ACCOUNT_ID_POSTURE")
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env

	}
	service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
		Authenticator: authenticator,
		URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
	})

	source := service.NewListScopesOptions(accountId, scopeName)
	source.SetName(scopeName)

	reply, response, err := service.ListScopes(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		panic(err)
	}

	for _, scope := range reply.Scopes {
		if *scope.ScopeID == scopeId {
			fmt.Println("scope id " + *scope.ScopeID)
			fmt.Println("scope name " + *scope.Name)
			if scope.Scans != nil {
				for _, scans := range scope.Scans {
					fmt.Println("scan id " + *scans.ScanID)
					fmt.Println("scan status " + *scans.Status)
					//if *scans.Status == "validation_completed" {
					if *scans.Status == "discovery_completed" {
						fmt.Println("discovery completed. test pass")
						return true
					}
				}
			} else {
				fmt.Println("no scans yet")
			}
		}
	}

	fmt.Println("in progress, re-checking...")
	return false

}

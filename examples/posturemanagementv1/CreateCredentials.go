package examples

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"os"
)

func CreateCredentials() {
	apiKey := os.Getenv("IAM_API_KEY")
	url := os.Getenv("IAM_APIKEY_URL")
	accountId := os.Getenv("ACCOUNT_ID")
	credentialPath := os.Getenv("CREDENTIAL_PATH")
	pemPath := os.Getenv("PEM_PATH")
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env

	}
	service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
		Authenticator: authenticator,
		URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
	})

	credentialFile, _ := os.Open(credentialPath)
	pemFile, _ := os.Open(pemPath)

	source := service.NewCreateCredentialOptions(accountId, credentialFile)
	source.SetPemFile(pemFile)

	_, response, err := service.CreateCredential(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		return
	}
	fmt.Println("Status Code: ", response.GetStatusCode())
}

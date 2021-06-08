package examples

import (
	"fmt"
	scc "github.com/ibm/scc-go-sdk/posturemanagementv1"
	"os"
)

func CreateCredentials(options scc.PostureManagementV1Options, accountId string, credentialPath string, pemPath string) (string, int) {
	service, _ := scc.NewPostureManagementV1(&options)

	credentialFile, _ := os.Open(credentialPath)
	pemFile, _ := os.Open(pemPath)

	source := service.NewCreateCredentialOptions(accountId, credentialFile)
	source.SetPemFile(pemFile)

	reply, response, err := service.CreateCredential(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		panic(err)
	}
	return *reply.CredentialID, response.GetStatusCode()
}

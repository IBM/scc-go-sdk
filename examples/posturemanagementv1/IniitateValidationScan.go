package examples

import (
	"fmt"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
)

func InitiateValidationScan(options scc.PostureManagementV1Options, accountId string, scopeId string, profileId string) (int, *string) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateValidationOptions(accountId)
	source.SetScopeID(scopeId)
	source.SetProfileID(profileId)
	source.SetGroupProfileID("0") //this is static value for now

	reply, response, err := service.CreateValidation(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		panic(err)
	}

	return response.StatusCode, reply.Message

}

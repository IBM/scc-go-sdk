package examples

import (
	"fmt"
	"github.com/google/uuid"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
)

func CreateCollector(options scc.PostureManagementV1Options, accountId string) (int, *string) {

	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateCollectorOptions(accountId)
	source.SetCollectorName("test-" + uuid.NewString())
	source.SetCollectorDescription("test collector")
	source.SetManagedBy("customer")
	source.SetIsPublic(true)
	source.SetPassPhrase("secret")

	reply, response, err := service.CreateCollector(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create collector: ", err)
		panic(err)
	}

	return response.StatusCode, reply.CollectorID
}

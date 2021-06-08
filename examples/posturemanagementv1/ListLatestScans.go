package examples

import (
	"fmt"
	scc "github.com/ibm/scc-go-sdk/posturemanagementv1"
)

func ListLatestScans(options scc.PostureManagementV1Options, accountId string) (int, []scc.ScanItem) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewListLatestScansOptions(accountId)

	reply, response, err := service.ListLatestScans(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		panic(err)
	}
	return response.StatusCode, reply.LatestScans

}

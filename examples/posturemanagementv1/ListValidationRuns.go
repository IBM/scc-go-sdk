package examples

import (
	"fmt"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
)

func ListValiadationRuns(options scc.PostureManagementV1Options, accountId string, scopeId string, profileId string) (int, scc.SummariesList) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewScanSummariesOptions(scopeId, accountId, profileId)

	reply, response, err := service.ScanSummaries(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to list validation runs: ", err)
		return 500, scc.SummariesList{}
	}

	return response.StatusCode, *reply

}

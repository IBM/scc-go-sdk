package examples

import (
	"fmt"
	scc "github.com/ibm/scc-go-sdk/posturemanagementv1"
)

func RetrieveScanSummary(options scc.PostureManagementV1Options, accountId string, scanId string, profileId string) (int, scc.Summary) {

	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewScansSummaryOptions(accountId, scanId, profileId)

	reply, response, err := service.ScansSummary(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to retrieve scan summary: ", err)
		return 500, *reply
	}

	return response.StatusCode, *reply

}

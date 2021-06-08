package examples

import (
	"fmt"
	scc "github.com/ibm/scc-go-sdk/posturemanagementv1"
)

func ListScopes(options scc.PostureManagementV1Options, accountId string, scopeName string, scopeId string, matchString string) (bool, string) {

	var scanId string

	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewListScopesOptions(accountId)
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
					scanId = *scans.ScanID
					fmt.Println("scan id " + scanId)
					fmt.Println("scan status " + *scans.Status)
					//if *scans.Status == "discovery_completed" {
					if *scans.Status == matchString {
						fmt.Println("test pass")
						return true, scanId
					}
				}
			} else {
				fmt.Println("no scans yet")
			}
		}
	}

	fmt.Println("in progress, re-checking...")
	return false, scanId

}

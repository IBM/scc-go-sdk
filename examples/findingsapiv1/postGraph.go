package examples

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/findingsv1"
)

//PostGraph posts a grapgql query
func PostGraph() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/graphql"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/findings", //Specify url or use default
	})

	//Query using ioutils
	//Graph query can also be read from a file. It is of type io.ReadClosure
	newQuery := ioutil.NopCloser(strings.NewReader(`query {findingCount: occurrenceCount(kind: "FINDING")}`))
	postGraphOptions := service.NewPostGraphOptions(accountID)
	postGraphOptions.SetBody(newQuery)
	postGraphOptions.SetHeaders(headers)
	res, operationErr := service.PostGraph(postGraphOptions)
	if operationErr != nil {
		fmt.Println("Err", operationErr)
	}
	fmt.Println(res.Result)

}

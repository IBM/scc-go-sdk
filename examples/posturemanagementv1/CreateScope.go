package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CreateScope(options scc.PostureManagementV1Options, accountId string, credentialId string, collectorIds []string) (int, *string, *string) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateScopeOptions(accountId)
	source.SetScopeName("SDKDEMO" + uuid.NewString())
	source.SetScopeDescription("sample scope description")
	source.SetEnvironmentType("ibm")
	source.SetCollectorIds(collectorIds)
	source.SetCredentialID(credentialId)

	result, response, err := service.CreateScope(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		panic(err)
	}

	return response.StatusCode, result.ScopeID, result.ScopeName

}

type access struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	Expiration   string `json:"expiration"`
	Scope        string `json:"scope"`
}

type tlDiscover struct {
	DiscoveryLevel int    `json:"discoveryLevel"`
	GatewayIds     []int  `json:"gatewayIds"`
	RequestType    string `json:"requestType"`
	SchemaId       int    `json:"schemaId"`
}

func getAuthToken() string {
	url := os.Getenv("IAM_APIKEY_URL")
	method := "POST"

	payload := strings.NewReader("apikey=" + os.Getenv("ACCOUNT_ID_POSTURE") + "&response_type=cloud_iam&grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey")

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	accessData := access{}
	strBody := string(body)
	json.Unmarshal([]byte(strBody), &accessData)
	return accessData.AccessToken
}

func demoDiscovery(gatewayIds []string, scopeId string) int {
	accountId := os.Getenv("ACCOUNT_ID_POSTURE")
	authToken := getAuthToken()
	url := os.Getenv("API_URL") + "/alpha/v1.0/schemas/tldiscover"
	method := "POST"

	tld := tlDiscover{}
	tld.DiscoveryLevel = 1
	tld.RequestType = "TLDISCOVER"
	tld.SchemaId, _ = strconv.Atoi(scopeId)

	for _, i := range gatewayIds {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		tld.GatewayIds = append(tld.GatewayIds, j)
	}

	requestByte, _ := json.Marshal(tld)
	requestReader := bytes.NewReader(requestByte)

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, requestReader)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authToken)
	req.Header.Add("REALM", accountId)
	req.Header.Add("transaction-id", uuid.New().String())

	res, _ := client.Do(req)
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)

	return res.StatusCode

}

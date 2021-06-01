/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package posturemanagementv1_test

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/google/uuid"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/**
 * This file contains an integration test for the Posture Management v1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SCC test`, func() {

	var (
		collectorId *string
		scopeId     *string
	)
	uuidWithHyphen := uuid.New().String()
	apiKey := os.Getenv("IAM_API_KEY")
	authUrl := os.Getenv("IAM_APIKEY_URL")
	accountId := os.Getenv("ACCOUNT_ID")
	apiUrl := os.Getenv("API_URL")

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    authUrl, //use for dev/preprod env
	}

	Describe(`Integration test`, func() {
		Describe(`Create collector suite`, func() {
			It(`Create collector`, func() {
				service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           apiUrl, //Specify url or use default
				})

				source := service.NewCreateCollectorOptions(accountId)
				source.SetCollectorName("test-" + uuidWithHyphen)
				source.SetCollectorDescription("test collector")
				source.SetManagedBy("customer")
				source.SetIsPublic(true)
				source.SetPassPhrase("secret")

				reply, response, err := service.CreateCollector(source)
				collectorId = reply.CollectorID

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create collector: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(201))
			})
			It(`Delete collector for cleanup`, func() {
				responseCode := hardDeleteCollector(collectorId)

				Expect(responseCode).To(Equal(200))
			})
		})
		Describe(`Create scope suite`, func() {
			It(`Create scope`, func() {
				service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           apiUrl, //Specify url or use default
				})

				source := service.NewCreateScopeOptions(accountId)
				source.SetScopeName("scope-" + uuidWithHyphen)
				source.SetScopeDescription("test scope")
				source.SetCredentialID("5645")
				source.SetCollectorIds([]string{"1417"})
				source.SetEnvironmentType("ibm")

				reply, response, err := service.CreateScope(source)
				scopeId = reply.ScopeID

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create scope: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(200))
				Expect(reply).ToNot(BeNil())
			})
			It(`Delete scope for cleanup`, func() {
				response := hardDeleteScope(scopeId)
				Expect(response).To(Equal(200))

			})
		})
		Describe(`Create credential suite`, func() {
			It(`Create credential`, func() {
				credentialPath := os.Getenv("CREDENTIAL_PATH")
				pemPath := os.Getenv("PEM_PATH")

				service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
				})

				credentialFile, _ := os.Open(credentialPath)
				pemFile, _ := os.Open(pemPath)

				source := service.NewCreateCredentialOptions(accountId, credentialFile)
				source.SetPemFile(pemFile)

				reply, response, err := service.CreateCredential(source)

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create credential: ", err)
					return
				}
				Expect(response.GetStatusCode()).To(Equal(201))
				Expect(reply).ToNot(BeNil())
			})
		})
		FDescribe(`Create scan`, func() {
			It(`Create scan`, func() {
				service, _ := scc.NewPostureManagementV1(&scc.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
				})

				source := service.NewScanSummariesOptions("1188", accountId)
				source.SetProfileID("48")
				source.SetGroupProfileID("1")

				reply, response, err := service.ScanSummaries(source)

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create scan: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(200))
				Expect(reply).ToNot(BeNil())
			})
		})

	})
})

//
// Utility functions are declared in the unit test file
//

type access struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	Expiration   string `json:"expiration"`
	Scope        string `json:"scope"`
}

func getAuthToken() string {
	url := "https://iam.test.cloud.ibm.com/identity/token"
	method := "POST"

	payload := strings.NewReader("apikey=DjsEbdqjIwuP9bfTyGATAuJ9u55dsMbVNvJ8cVWdzoxz&response_type=cloud_iam&grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey")

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

func hardDeleteCollector(collectorId *string) int {
	accountId := os.Getenv("ACCOUNT_ID")
	authToken := getAuthToken()
	collectorIdValue := *collectorId
	url := "https://asap-dev.compliance.test.cloud.ibm.com/alpha/v1.0/collectors/" + collectorIdValue
	method := "DELETE"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Authorization", authToken)
	req.Header.Add("REALM", accountId)
	req.Header.Add("transaction-id", uuid.New().String())

	res, _ := client.Do(req)
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)

	return res.StatusCode

}

func hardDeleteScope(scopeId *string) int {
	accountId := os.Getenv("ACCOUNT_ID")
	authToken := getAuthToken()
	scopeIdValue := *scopeId
	url := "https://asap-dev.compliance.test.cloud.ibm.com/alpha/v1.0/schemas/" + scopeIdValue
	method := "DELETE"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Authorization", authToken)
	req.Header.Add("REALM", accountId)
	req.Header.Add("transaction-id", uuid.New().String())

	res, _ := client.Do(req)
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)

	return res.StatusCode
}

func hardDeleteCredential() {

}

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
	"github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
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

	var collectorId *string
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
				service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           apiUrl, //Specify url or use default
				})

				source := service.NewCreateCollectorOptions(accountId)
				source.SetCollectorName("jason-" + uuidWithHyphen)
				source.SetCollectorDescription("jason collector")
				source.SetManagedBy("CUSTOMER")
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
				service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           apiUrl, //Specify url or use default
				})

				source := service.NewCreateScopeOptions(accountId)
				source.SetScopeName("scope-" + uuidWithHyphen)
				source.SetScopeDescription("jason scope")
				source.SetCredentialID("5645")
				source.SetCollectorIds([]string{"1417"})
				source.SetEnvironmentType("ibm")

				_, response, err := service.CreateScope(source)

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create scope: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(200))
			})
			It(`Delete scope for cleanup`, func() {

				//Expect(response.StatusCode).To(Equal(200))
			})
		})

		Describe(`Create credential suite`, func() {
			It(`Create credential`, func() {
				apiKey := os.Getenv("IAM_API_KEY")
				url := os.Getenv("IAM_APIKEY_URL")
				accountId := os.Getenv("ACCOUNT_ID")
				authenticator := &core.IamAuthenticator{
					ApiKey: apiKey,
					URL:    url, //use for dev/preprod env
				}
				service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
				})

				source := service.NewCreateCollectorOptions(accountId)
				source.SetCollectorName("jason-test-collector-05")
				source.SetCollectorDescription("jason scope")
				source.SetManagedBy("CUSTOMER")
				source.SetIsPublic(true)
				source.SetPassPhrase("secret")

				_, response, err := service.CreateCollector(source)

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create collector: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(200))
			})
			It(`Delete credential for cleanup`, func() {

				//Expect(response.StatusCode).To(Equal(200))
			})
		})

		Describe(`Create scan`, func() {
			It(`Create scan`, func() {
				apiKey := os.Getenv("IAM_API_KEY")
				url := os.Getenv("IAM_APIKEY_URL")
				accountId := os.Getenv("ACCOUNT_ID")
				authenticator := &core.IamAuthenticator{
					ApiKey: apiKey,
					URL:    url, //use for dev/preprod env
				}
				service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					Authenticator: authenticator,
					URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
				})

				source := service.NewCreateCollectorOptions(accountId)
				source.SetCollectorName("jason-test-collector-05")
				source.SetCollectorDescription("jason scope")
				source.SetManagedBy("CUSTOMER")
				source.SetIsPublic(true)
				source.SetPassPhrase("secret")

				_, response, err := service.CreateCollector(source)

				if err != nil {
					fmt.Println(response.Result)
					fmt.Println("Failed to create collector: ", err)
					return
				}
				Expect(response.StatusCode).To(Equal(200))
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
	collectorIdValue := *scopeId
	url := "https://asap-dev.compliance.test.cloud.ibm.com/alpha/v1.0/schemas/" + collectorIdValue
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

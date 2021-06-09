// +build integration
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
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	examples "github.com/ibm/scc-go-sdk/examples/posturemanagementv1"
	scc "github.com/ibm/scc-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the Posture Management v1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var (
	err           error
	config        map[string]string
	collectorIds  []string
	collectorId   string
	credentialId  string
	accountId     string
	apiKey        string
	authUrl       string
	apiUrl        string
	scopeId       *string
	scanId        string
	scopeName     *string
	authenticator core.IamAuthenticator
	options       scc.PostureManagementV1Options
)

var _ = XDescribe(`SCC test`, func() {
	const externalConfigFile = "../posture_management_v1.env"

	Describe(`Demo`, func() {

		BeforeEach(func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err != nil {
				return
			}
			config, err = core.GetServiceProperties(scc.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountId = config["ACCOUNT_ID"]
			if accountId == "" {
				Skip("Unable to load account configuration property, skipping tests")
			}

			apiKey = config["IAM_API_KEY"]
			if apiKey == "" {
				Skip("Unable to load API key configuration property, skipping tests")
			}

			authUrl = config["IAM_APIKEY_URL"]
			if authUrl == "" {
				Skip("Unable to load auth URL configuration property, skipping tests")
			}

			apiUrl = config["API_URL"]
			if apiUrl == "" {
				Skip("Unable to load API URL configuration property, skipping tests")
			}

			authenticator = core.IamAuthenticator{
				ApiKey: apiKey,
				URL:    authUrl,
			}

			options = scc.PostureManagementV1Options{
				Authenticator: &authenticator,
				URL:           apiUrl,
			}

		})

		//TODO fix collector issue -- skip for now
		XIt(`Create Collector`, func() {
			fmt.Println(`Create Collector`)
			statusCode, id := examples.CreateCollector(options, accountId)
			collectorId = *id
			Expect(statusCode).To(Equal(201))
			Expect(collectorId).NotTo(BeNil())

			fmt.Println("Collector Id: " + collectorId)
			fmt.Println(`Create Collector Successful`)
		})

		It(`Create Credential`, func() {
			fmt.Println(`Create Credential`)
			var statusCode int
			credentialPath := config["CREDENTIAL_PATH"]
			pemPath := config["PEM_PATH"]
			credentialId, statusCode = examples.CreateCredentials(options, accountId, credentialPath, pemPath)
			Expect(statusCode).To(Equal(201))
			Expect(credentialId).NotTo(BeNil())
			fmt.Println("Credential Id: " + credentialId)
			fmt.Println(`Create Credential Successful`)
		})

		//TODO override collector id for now until create collector is resolved
		It(`Create Scope`, func() {
			fmt.Println(`Create Scope Started`)
			var statusCode int
			collectorId = "1499"
			//credentialId = "1587"
			collectorIds = append(collectorIds, collectorId)
			statusCode, scopeId, scopeName = examples.CreateScope(options, accountId, credentialId, collectorIds)

			Expect(statusCode).To(Equal(201))
			Expect(scopeId).ToNot(BeNil())
			Expect(scopeName).ToNot(BeNil())
			fmt.Println("created scope id: " + *scopeId)
			fmt.Println("created scope name: " + *scopeName)
			fmt.Println(`Create Scope Successful`)
		})

		It(`List Scopes`, func() {
			fmt.Println(`List Scopes`)
			demoListScope(options, scopeId)

			Eventually(func() bool {
				var isCompleted bool
				isCompleted, scanId = examples.ListScopes(options, accountId, *scopeName, *scopeId, "discovery_completed")
				return isCompleted
			}, "12000s", "20s").Should(BeTrue())
			fmt.Println(`List Scope Successful`)
		})

		It(`List Profiles`, func() {
			fmt.Println(`List Profiles`)
			statusCode, profileList := examples.ListProfiles(options, accountId)
			Expect(statusCode).To(Equal(200))
			Expect(profileList).ToNot(BeNil())
			fmt.Println(`List Profiles Successful`)
		})

		It(`Initiate Scan Validation`, func() {
			fmt.Println(`Create Scan`)
			statusCode, message := examples.InitiateValidationScan(options, accountId, *scopeId, "48")
			fmt.Println("message: " + *message)
			Expect(statusCode).To(Equal(202))
			Expect(message).ToNot(BeNil())
			fmt.Println(`Create Scan Successful`)
		})
		It(`Check Scan Status`, func() {
			fmt.Println(`Check Scan status`)
			Eventually(func() bool {
				var isCompleted bool
				isCompleted, scanId = examples.ListScopes(options, accountId, *scopeName, *scopeId, "validation_completed")
				return isCompleted
			}, "12000s", "20s").Should(BeTrue())
		})
		It(`List latest scans`, func() {
			fmt.Println(`List latest scans`)
			statusCode, list := examples.ListLatestScans(options, accountId)
			Expect(statusCode).To(Equal(200))
			scanId = *list[0].ScanID
			Expect(list).ToNot(BeNil())
			fmt.Println(`List latest scans successful`)

		})

		It(`Read scan`, func() {
			fmt.Println(`Read scan summary details`)
			statusCode, list := examples.RetrieveScanSummary(options, accountId, scanId, "48")
			Expect(statusCode).To(Equal(200))
			Expect(list).ToNot(BeNil())
			fmt.Println(`Read scan summary details successful`)
		})

		It(`List Validation Runs`, func() {
			fmt.Println(`List Validation Runs`)
			statusCode, list := examples.ListValiadationRuns(options, accountId, "17885", "48")
			Expect(statusCode).To(Equal(200))
			Expect(list).ToNot(BeNil())
			fmt.Println(`List Validation Runs Successful`)
		})

	})
})

//
// Utility functions are declared in the unit test file
//

//TODO: add query parameter name to verify discovery
func demoListScope(options scc.PostureManagementV1Options, scopeId *string) {

	service, _ := scc.NewPostureManagementV1(&options)

	//var scopeIdMatch int
	source := service.NewListScopesOptions(accountId)

	reply, response, err := service.ListScopes(source)

	//
	//for _, i := range reply.Scopes {
	//	scopeIdMatch = strings.Compare(*i.ScopeID, *scopeId)
	//}

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		return
	}
	Expect(response.StatusCode).To(Equal(200))
	//Expect(scopeIdMatch).To(Equal(0))
	Expect(reply.Scopes).ToNot(BeNil())
}

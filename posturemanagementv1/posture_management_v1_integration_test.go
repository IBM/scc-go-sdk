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
	"github.com/google/uuid"
	examples "github.com/ibm-cloud-security/scc-go-sdk/examples/posturemanagementv1"
	scc "github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
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
	uuidWithHyphen string
	collectorIds   []string
	credentialId   string
	accountId      string
	scopeId        *string
	scanId         string
	scopeName      *string
)

var _ = Describe(`SCC test`, func() {

	uuidWithHyphen = uuid.New().String()
	accountId = os.Getenv("ACCOUNT_ID_POSTURE")
	apiKey := os.Getenv("IAM_API_KEY")
	authUrl := os.Getenv("IAM_APIKEY_URL")
	apiUrl := os.Getenv("API_URL")

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    authUrl, //use for dev/preprod env
	}

	options := scc.PostureManagementV1Options{
		Authenticator: authenticator,
		URL:           apiUrl,
	}

	XDescribe(`polling`, func() {
		It(`should poll`, func() {
		})
	})

	FDescribe(`Demo`, func() {

		//TODO fix collector issue -- skip for now
		XIt(`Create Collector`, func() {
			fmt.Println(`Create Collector`)
			statusCode, id := examples.CreateCollector(options, accountId)
			Expect(statusCode).To(Equal(201))
			Expect(id).NotTo(BeNil())
		})

		//new credential to be created
		FIt(`Create Credential`, func() {
			fmt.Println(`Create Credential`)
			var statusCode int
			credentialPath := os.Getenv("CREDENTIAL_PATH")
			pemPath := os.Getenv("PEM_PATH")
			credentialId, statusCode = examples.CreateCredentials(options, accountId, credentialPath, pemPath)
			Expect(statusCode).To(Equal(201))
			Expect(credentialId).NotTo(BeNil())
		})

		//TODO override collector id for now until create collector is resolved
		It(`Create Scope`, func() {
			fmt.Println(`Create Scope`)
			collectorId := "822"
			collectorIds = append(collectorIds, collectorId)
			var statusCode int
			credentialId = "1587"
			statusCode, scopeId, scopeName = examples.CreateScope(options, credentialId, collectorIds)
			fmt.Println("created scope id: " + *scopeId)
			fmt.Println("created scope name: " + *scopeName)

			Expect(statusCode).To(Equal(201))
			Expect(scopeId).ToNot(BeNil())
			Expect(scopeName).ToNot(BeNil())
		})

		It(`List Scopes`, func() {
			fmt.Println(`List Scopes`)
			demoListScope(options, scopeId)

			Eventually(func() bool {
				var isCompleted bool
				isCompleted, scanId = examples.ListScopes(options, accountId, *scopeName, *scopeId, "discovery_completed")
				return isCompleted
			}, "12000s", "20s").Should(BeTrue())
		})

		It(`List Profiles`, func() {
			fmt.Println(`List Profiles`)
			statusCode, profileList := examples.ListProfiles(options, accountId)
			Expect(statusCode).To(Equal(200))
			Expect(profileList).ToNot(BeNil())
		})

		It(`Initiate Scan Validation`, func() {
			fmt.Println(`Create Scan`)
			statusCode, message := examples.InitiateValidationScan(options, accountId, *scopeId, "48")
			fmt.Println("message: " + *message)
			Expect(statusCode).To(Equal(202))
			Expect(message).ToNot(BeNil())
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
			statusCode, list := examples.ListLatestScans(options, accountId, scanId)
			Expect(statusCode).To(Equal(200))
			Expect(list).ToNot(BeNil())

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

func demoListProfiles(options scc.PostureManagementV1Options) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewListProfilesOptions(accountId)

	reply, response, err := service.ListProfiles(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to list profiles: ", err)
		return
	}

	Expect(reply.Profiles).ToNot(BeNil())
	Expect(response.StatusCode).To(Equal(200))
}
func demoCreateScanValidation(options scc.PostureManagementV1Options, scopeId string, profileId string) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateValidationOptions(accountId)
	source.SetScopeID(scopeId)
	source.SetProfileID(profileId)

	reply, response, err := service.CreateValidation(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		return
	}
	//fmt.Println("Status code: " + string(response.StatusCode))

	Expect(response.StatusCode).To(Equal(200))
	Expect(reply.Message).ToNot(BeNil())

}
func demoListScans(options scc.PostureManagementV1Options) {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewListLatestScansOptions(accountId)
	reply, response, err := service.ListLatestScans(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		return
	}

	Expect(response.StatusCode).To(Equal(200))
	Expect(reply.LatestScans).ToNot(BeNil())

}

func demoReadScan(options scc.PostureManagementV1Options, scanId string, profileId string) *scc.Summary {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewScansSummaryOptions(accountId, scanId, profileId)
	reply, response, err := service.ScansSummary(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		panic(err)
	}

	return reply
}

func demoScanSummary(options scc.PostureManagementV1Options, scopeId string) {
	var profileId string
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewScanSummariesOptions(scopeId, accountId, profileId)
	reply, response, err := service.ScanSummaries(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scan: ", err)
		return
	}

	Expect(response.StatusCode).To(Equal(200))
	Expect(reply.Summaries).ToNot(BeNil())

}

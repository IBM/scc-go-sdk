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
	collectorId    *string
	collectorIds   []string
	credentialId   string
	accountId      string
	scopeId        *string
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

	Describe(`Demo`, func() {

		It(`Create Collector`, func() {
			fmt.Println(`Create Collector`)
			collectorId = demoCreateCollector(options)
		})
		It(`Create Credential`, func() {
			fmt.Println(`Create Credential`)
			credentialId = demoCreateCredential(options)
		})
		It(`Create Scope`, func() {
			fmt.Println(`Create Scope`)
			collectorIds = append(collectorIds, *collectorId)
			scopeId = demoCreateScope(options, credentialId, collectorIds)
		})
		It(`List Scopes`, func() {
			fmt.Println(`List Scopes`)
			demoListScope(options, scopeId)
		})
		It(`List Profiles`, func() {
			fmt.Println(`List Profiles`)
			demoListProfiles(options)
		})
		It(`Initiate Scan Validation`, func() {
			fmt.Println(`Create Scan`)
			demoCreateScanValidation(options, *scopeId, "44")
		})

	})
})

//
// Utility functions are declared in the unit test file
//
func demoCreateCollector(options scc.PostureManagementV1Options) *string {

	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateCollectorOptions(accountId)
	source.SetCollectorName("test-" + uuidWithHyphen)
	source.SetCollectorDescription("test collector")
	source.SetManagedBy("customer")
	source.SetIsPublic(true)
	source.SetPassPhrase("secret")

	reply, response, err := service.CreateCollector(source)
	collectorId := reply.CollectorID

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create collector: ", err)
		return nil
	}

	Expect(response.StatusCode).To(Equal(201))

	return collectorId
}
func demoCreateCredential(options scc.PostureManagementV1Options) string {
	credentialPath := os.Getenv("CREDENTIAL_PATH")
	pemPath := os.Getenv("PEM_PATH")

	service, _ := scc.NewPostureManagementV1(&options)

	credentialFile, _ := os.Open(credentialPath)
	pemFile, _ := os.Open(pemPath)

	source := service.NewCreateCredentialOptions(accountId, credentialFile)
	source.SetPemFile(pemFile)

	reply, response, err := service.CreateCredential(source)

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create credential: ", err)
		return err.Error()
	}

	Expect(response.StatusCode).To(Equal(201))
	Expect(reply.CredentialID).ToNot(BeNil())
	Expect(reply.CreatedTime).ToNot(BeNil())

	return *reply.CredentialID
}
func demoCreateScope(options scc.PostureManagementV1Options, credentialId string, collectorIds []string) *string {
	service, _ := scc.NewPostureManagementV1(&options)

	source := service.NewCreateScopeOptions(accountId)
	source.SetScopeName("scope-" + uuidWithHyphen)
	source.SetScopeDescription("test scope")
	source.SetCredentialID(credentialId)
	source.SetCollectorIds(collectorIds)
	source.SetEnvironmentType("ibm")

	reply, response, err := service.CreateScope(source)
	scopeId := reply.ScopeID

	if err != nil {
		fmt.Println(response.Result)
		fmt.Println("Failed to create scope: ", err)
		return nil
	}
	Expect(response.StatusCode).To(Equal(201))
	Expect(reply.ScopeID).ToNot(BeNil())
	Expect(reply.CreatedTime).ToNot(BeNil())
	Expect(reply.ModifiedTime).ToNot(BeNil())
	return scopeId
}

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

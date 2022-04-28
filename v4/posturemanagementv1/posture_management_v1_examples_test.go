//go:build examples
// +build examples

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
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv1"
)

//
// This file provides an example of how to use the Posture Management service.
//
// The following configuration properties are assumed to be defined:
// POSTURE_MANAGEMENT_URL=<service base url>
// POSTURE_MANAGEMENT_AUTH_TYPE=iam
// POSTURE_MANAGEMENT_APIKEY=<IAM apikey>
// POSTURE_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../posture_management_v1.env"

var (
	postureManagementService *posturemanagementv1.PostureManagementV1
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`PostureManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			postureManagementServiceOptions := &posturemanagementv1.PostureManagementV1Options{
				AccountID: core.StringPtr("testString"),
			}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(postureManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(postureManagementService).ToNot(BeNil())
		})
	})

	Describe(`PostureManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLatestScans request example`, func() {
			fmt.Println("\nListLatestScans() result:")
			// begin-list_latest_scans

			listLatestScansOptions := postureManagementService.NewListLatestScansOptions()

			scansList, response, err := postureManagementService.ListLatestScans(listLatestScansOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scansList, "", "  ")
			fmt.Println(string(b))

			// end-list_latest_scans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scansList).ToNot(BeNil())

		})
		It(`CreateValidation request example`, func() {
			fmt.Println("\nCreateValidation() result:")
			// begin-create_validation

			createValidationOptions := postureManagementService.NewCreateValidationOptions(
				"1",
				"6",
			)

			result, response, err := postureManagementService.CreateValidation(createValidationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-create_validation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())

		})
		It(`ScansSummary request example`, func() {
			fmt.Println("\nScansSummary() result:")
			// begin-scans_summary

			scansSummaryOptions := postureManagementService.NewScansSummaryOptions(
				"testString",
				"testString",
			)

			summary, response, err := postureManagementService.ScansSummary(scansSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(summary, "", "  ")
			fmt.Println(string(b))

			// end-scans_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summary).ToNot(BeNil())

		})
		It(`ScanSummaries request example`, func() {
			fmt.Println("\nScanSummaries() result:")
			// begin-scan_summaries

			scanSummariesOptions := postureManagementService.NewScanSummariesOptions(
				"testString",
				"testString",
			)

			summariesList, response, err := postureManagementService.ScanSummaries(scanSummariesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(summariesList, "", "  ")
			fmt.Println(string(b))

			// end-scan_summaries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summariesList).ToNot(BeNil())

		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := postureManagementService.NewListProfilesOptions()

			profilesList, response, err := postureManagementService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profilesList, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

		})
		It(`CreateScope request example`, func() {
			fmt.Println("\nCreateScope() result:")
			// begin-create_scope

			createScopeOptions := postureManagementService.NewCreateScopeOptions()
			createScopeOptions.SetScopeName("IBMSchema-new-048-test")
			createScopeOptions.SetScopeDescription("IBMSchema")
			createScopeOptions.SetCollectorIds([]string{"20"})
			createScopeOptions.SetCredentialID("5")
			createScopeOptions.SetEnvironmentType("ibm")

			scope, response, err := postureManagementService.CreateScope(createScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-create_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

		})
		It(`ListScopes request example`, func() {
			fmt.Println("\nListScopes() result:")
			// begin-list_scopes

			listScopesOptions := postureManagementService.NewListScopesOptions()

			scopesList, response, err := postureManagementService.ListScopes(listScopesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopesList, "", "  ")
			fmt.Println(string(b))

			// end-list_scopes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

		})
		It(`CreateCollector request example`, func() {
			fmt.Println("\nCreateCollector() result:")
			// begin-create_collector

			createCollectorOptions := postureManagementService.NewCreateCollectorOptions()
			createCollectorOptions.SetCollectorName("IBM-collector-sample")
			createCollectorOptions.SetCollectorDescription("sample collector")
			createCollectorOptions.SetIsPublic(true)
			createCollectorOptions.SetManagedBy("ibm")
			createCollectorOptions.SetPassPhrase("secret")

			collector, response, err := postureManagementService.CreateCollector(createCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collector, "", "  ")
			fmt.Println(string(b))

			// end-create_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collector).ToNot(BeNil())

		})
		It(`CreateCredential request example`, func() {
			fmt.Println("\nCreateCredential() result:")
			// begin-create_credential

			createCredentialOptions := postureManagementService.NewCreateCredentialOptions(
				CreateMockReader("This is a mock file."),
			)

			credential, response, err := postureManagementService.CreateCredential(createCredentialOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credential, "", "  ")
			fmt.Println(string(b))

			// end-create_credential

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())

		})
	})
})

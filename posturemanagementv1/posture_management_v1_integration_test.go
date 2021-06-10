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
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the posturemanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PostureManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../posture_management_v1.env"

	var (
		err                      error
		postureManagementService *posturemanagementv1.PostureManagementV1
		serviceURL               string
		config                   map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			postureManagementServiceOptions := &posturemanagementv1.PostureManagementV1Options{
				AccountID: core.StringPtr("testString"),
			}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(postureManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(postureManagementService).ToNot(BeNil())
			Expect(postureManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListLatestScans - List latest scans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLatestScans(listLatestScansOptions *ListLatestScansOptions)`, func() {

			listLatestScansOptions := &posturemanagementv1.ListLatestScansOptions{
				TransactionID: core.StringPtr("testString"),
				Name:          core.StringPtr("testString"),
				Offset:        core.Int64Ptr(int64(38)),
				Limit:         core.Int64Ptr(int64(100)),
			}

			scansList, response, err := postureManagementService.ListLatestScans(listLatestScansOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scansList).ToNot(BeNil())

		})
	})

	Describe(`CreateValidation - Initiate a validation scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidation(createValidationOptions *CreateValidationOptions)`, func() {

			createValidationOptions := &posturemanagementv1.CreateValidationOptions{
				ScopeID:        core.StringPtr("1"),
				ProfileID:      core.StringPtr("6"),
				GroupProfileID: core.StringPtr("13"),
				TransactionID:  core.StringPtr("testString"),
			}

			result, response, err := postureManagementService.CreateValidation(createValidationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())

		})
	})

	Describe(`ScansSummary - Retrieve the summary of a specific scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScansSummary(scansSummaryOptions *ScansSummaryOptions)`, func() {

			scansSummaryOptions := &posturemanagementv1.ScansSummaryOptions{
				ScanID:        core.StringPtr("testString"),
				ProfileID:     core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Name:          core.StringPtr("testString"),
			}

			summary, response, err := postureManagementService.ScansSummary(scansSummaryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summary).ToNot(BeNil())

		})
	})

	Describe(`ScanSummaries - List the validation summaries for a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScanSummaries(scanSummariesOptions *ScanSummariesOptions)`, func() {

			scanSummariesOptions := &posturemanagementv1.ScanSummariesOptions{
				ScopeID:        core.StringPtr("testString"),
				TransactionID:  core.StringPtr("testString"),
				ProfileID:      core.StringPtr("testString"),
				GroupProfileID: core.StringPtr("testString"),
				Offset:         core.Int64Ptr(int64(38)),
				Limit:          core.Int64Ptr(int64(100)),
			}

			summariesList, response, err := postureManagementService.ScanSummaries(scanSummariesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summariesList).ToNot(BeNil())

		})
	})

	Describe(`ListProfiles - List profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {

			listProfilesOptions := &posturemanagementv1.ListProfilesOptions{
				TransactionID: core.StringPtr("testString"),
				Name:          core.StringPtr("testString"),
				Offset:        core.Int64Ptr(int64(38)),
				Limit:         core.Int64Ptr(int64(100)),
			}

			profilesList, response, err := postureManagementService.ListProfiles(listProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profilesList).ToNot(BeNil())

		})
	})

	Describe(`CreateScope - Create a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {

			createScopeOptions := &posturemanagementv1.CreateScopeOptions{
				ScopeName:        core.StringPtr("IBM-Scope-new-048-test"),
				ScopeDescription: core.StringPtr("IBM Scope Example"),
				CollectorIds:     []string{"20"},
				CredentialID:     core.StringPtr("5"),
				EnvironmentType:  core.StringPtr("ibm"),
				TransactionID:    core.StringPtr("testString"),
			}

			scope, response, err := postureManagementService.CreateScope(createScopeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

		})
	})

	Describe(`ListScopes - List scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {

			listScopesOptions := &posturemanagementv1.ListScopesOptions{
				TransactionID: core.StringPtr("testString"),
				Name:          core.StringPtr("testString"),
			}

			scopesList, response, err := postureManagementService.ListScopes(listScopesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopesList).ToNot(BeNil())

		})
	})

	Describe(`CreateCollector - Create a collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCollector(createCollectorOptions *CreateCollectorOptions)`, func() {

			createCollectorOptions := &posturemanagementv1.CreateCollectorOptions{
				CollectorName:        core.StringPtr("IBM-collector-sample"),
				CollectorDescription: core.StringPtr("sample collector"),
				IsPublic:             core.BoolPtr(true),
				ManagedBy:            core.StringPtr("ibm"),
				PassPhrase:           core.StringPtr("secret"),
				TransactionID:        core.StringPtr("testString"),
			}

			collector, response, err := postureManagementService.CreateCollector(createCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collector).ToNot(BeNil())

		})
	})

	Describe(`CreateCredential - Create a credential`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCredential(createCredentialOptions *CreateCredentialOptions)`, func() {

			createCredentialOptions := &posturemanagementv1.CreateCredentialOptions{
				CredentialDataFile: CreateMockReader("This is a mock file."),
				PemFile:            CreateMockReader("This is a mock file."),
				TransactionID:      core.StringPtr("testString"),
			}

			credential, response, err := postureManagementService.CreateCredential(createCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//

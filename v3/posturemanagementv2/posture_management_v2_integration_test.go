//go:build integration
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

package posturemanagementv2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/posturemanagementv2"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the posturemanagementv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PostureManagementV2 Integration Tests`, func() {

	const externalConfigFile = "posture_management_v2.env"

	var (
		err                      error
		postureManagementService *posturemanagementv2.PostureManagementV2
		serviceURL               string
		config                   map[string]string
		accountID                string
		credentialID             string
		collectorID              string
		scopeID                  string
		profileID                string
		credentialIDScope        string
		collectorIDScope         string
		credentialIDScopeUpdate  string
		collectorIDScopeUpdate   string
		scanID                   string
		scopeIDScan              string
		profileIDScan            string
		groupProfileID           string
		correlationID            string
		reportSettingID          string
		transactionID            string
		profileType              string
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
			config, err = core.GetServiceProperties(posturemanagementv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}
			accountID = config["POSTURE_ACCOUNT_ID"]
			//profileID = config["PROFILE_ID"]
			scanID = config["SCAN_ID"]
			groupProfileID = config["GROUP_PROFILE_ID"]
			scopeIDScan = config["SCOPE_ID_SCAN"]
			profileIDScan = config["PROFILE_ID_SCAN"]
			credentialIDScope = config["CREDENTIAL_ID_SCOPE"]
			collectorIDScope = config["COLLECTOR_ID_SCOPE"]
			credentialIDScopeUpdate = config["CREDENTIAL_ID_SCOPE_UPDATE"]
			collectorIDScopeUpdate = config["COLLECTOR_ID_SCOPE_UPDATE"]
			correlationID = config["CORRELATION_ID"]
			reportSettingID = config["REPORT_SETTING_ID"]
			profileType = config["PROFILE_TYPE"]
			transactionID = uuid.NewString()
			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			postureManagementServiceOptions := &posturemanagementv2.PostureManagementV2Options{}

			postureManagementService, err = posturemanagementv2.NewPostureManagementV2UsingExternalConfig(postureManagementServiceOptions)

			Expect(err).To(BeNil())
			Expect(postureManagementService).ToNot(BeNil())
			Expect(postureManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			postureManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateCredential - Add a credential`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCredential(createCredentialOptions *CreateCredentialOptions)`, func() {

			newCredentialDisplayFieldsModel := &posturemanagementv2.NewCredentialDisplayFields{
				IBMAPIKey: core.StringPtr("sample_api_key"),
			}

			credentialGroupModel := &posturemanagementv2.CredentialGroup{
				ID:         core.StringPtr("1"),
				Passphrase: core.StringPtr("passphrase"),
			}

			createCredentialOptions := &posturemanagementv2.CreateCredentialOptions{
				Enabled:       core.BoolPtr(true),
				Type:          core.StringPtr("ibm_cloud"),
				Name:          core.StringPtr("test_create"),
				Description:   core.StringPtr("This credential is used for testing."),
				DisplayFields: newCredentialDisplayFieldsModel,
				Group:         credentialGroupModel,
				Purpose:       core.StringPtr("discovery_fact_collection_remediation"),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			credential, response, err := postureManagementService.CreateCredential(createCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())
			credentialID = *(credential.ID)
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ListCredentials - List credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCredentials(listCredentialsOptions *ListCredentialsOptions) with pagination`, func() {
			var result posturemanagementv2.CredentialList

			listCredentialsOptions := &posturemanagementv2.ListCredentialsOptions{
				AccountID:     core.StringPtr(accountID),
				Offset:        core.Int64Ptr(int64(2)),
				Limit:         core.Int64Ptr(int64(3)),
				TransactionID: core.StringPtr(transactionID),
			}

			listCredentialsOptions.Offset = nil
			listCredentialsOptions.Limit = core.Int64Ptr(1)

			credentialList, response, err := postureManagementService.ListCredentials(listCredentialsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialList).ToNot(BeNil())
			result.Credentials = append(result.Credentials, credentialList.Credentials...)

			listCredentialsOptions.Offset, err = credentialList.GetNextOffset()
			Expect(err).To(BeNil())
		})
	})

	Describe(`GetCredential - View credential details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCredential(getCredentialOptions *GetCredentialOptions)`, func() {

			getCredentialOptions := &posturemanagementv2.GetCredentialOptions{
				ID:            core.StringPtr(credentialID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			credential, response, err := postureManagementService.GetCredential(getCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credential).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`UpdateCredential - Update a credential`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCredential(updateCredentialOptions *UpdateCredentialOptions)`, func() {

			updateCredentialDisplayFieldsModel := &posturemanagementv2.UpdateCredentialDisplayFields{
				IBMAPIKey: core.StringPtr("sample_api_key"),
			}

			updateCredentialOptions := &posturemanagementv2.UpdateCredentialOptions{
				ID:            core.StringPtr(credentialID),
				Enabled:       core.BoolPtr(true),
				Type:          core.StringPtr("ibm_cloud"),
				Name:          core.StringPtr("test_create"),
				Description:   core.StringPtr("This credential is used for testing."),
				DisplayFields: updateCredentialDisplayFieldsModel,
				Purpose:       core.StringPtr("discovery_fact_collection_remediation"),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			credential, response, err := postureManagementService.UpdateCredential(updateCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`CreateCollector - Create a collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCollector(createCollectorOptions *CreateCollectorOptions)`, func() {
			collectorNameGen := fmt.Sprintf("%s%d", "IBM-collector-sample", time.Now().UnixNano())
			createCollectorOptions := &posturemanagementv2.CreateCollectorOptions{
				Name:          core.StringPtr(collectorNameGen),
				IsPublic:      core.BoolPtr(true),
				ManagedBy:     core.StringPtr("customer"),
				Description:   core.StringPtr("sample collector"),
				Passphrase:    core.StringPtr("secret"),
				IsUbiImage:    core.BoolPtr(true),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			collector, response, err := postureManagementService.CreateCollector(createCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collector).ToNot(BeNil())
			collectorID = *(collector.ID)
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ListCollectors - List collectors`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCollectors(listCollectorsOptions *ListCollectorsOptions)`, func() {

			listCollectorsOptions := &posturemanagementv2.ListCollectorsOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			collectorList, response, err := postureManagementService.ListCollectors(listCollectorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectorList).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetCollector - View collector details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCollector(getCollectorOptions *GetCollectorOptions)`, func() {

			getCollectorOptions := &posturemanagementv2.GetCollectorOptions{
				ID:            core.StringPtr(collectorID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			collector, response, err := postureManagementService.GetCollector(getCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collector).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`UpdateCollector - Update a collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCollector(updateCollectorOptions *UpdateCollectorOptions)`, func() {

			collectorUpdateModel := &posturemanagementv2.CollectorUpdate{
				DisplayName: core.StringPtr("test-0112-collector_jj"),
				Description: core.StringPtr("This collector is used for testing."),
			}
			collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCollectorOptions := &posturemanagementv2.UpdateCollectorOptions{
				ID:            core.StringPtr(collectorID),
				Collector:     collectorUpdateModelAsPatch,
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			collector, response, err := postureManagementService.UpdateCollector(updateCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collector).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ImportProfiles - Import profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportProfiles(importProfilesOptions *ImportProfilesOptions)`, func() {

			importProfilesOptions := &posturemanagementv2.ImportProfilesOptions{
				File: CreateMockReader("\"profilename\",\"CUSTOM PROFILE SDK\"\n" +
					"\"profilemnemonic\",\n" +
					"\"profiledescription\",\"CUSTOM PROFILE SDK\"\n" +
					"\"##METAINFO ENDS##\"\n" +
					"\"ExternalControlId\",\"Description\",\"Parent\",\"ControlId\",\"Tags\""),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			basicResult, response, err := postureManagementService.ImportProfiles(importProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(basicResult).ToNot(BeNil())
			profileID = *(basicResult.ProfileID)
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ListProfiles - List profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions) with pagination`, func() {
			var result posturemanagementv2.ProfileList

			listProfilesOptions := &posturemanagementv2.ListProfilesOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
				Offset:        core.Int64Ptr(int64(2)),
				Limit:         core.Int64Ptr(int64(3)),
			}

			listProfilesOptions.Offset = nil
			listProfilesOptions.Limit = core.Int64Ptr(1)

			profileList, response, err := postureManagementService.ListProfiles(listProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileList).ToNot(BeNil())
			result.Profiles = append(result.Profiles, profileList.Profiles...)

			listProfilesOptions.Offset, err = profileList.GetNextOffset()
			Expect(err).To(BeNil())

		})
	})

	Describe(`GetProfile - View profile details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {

			getProfileOptions := &posturemanagementv2.GetProfileOptions{
				ID:            core.StringPtr(profileID),
				ProfileType:   core.StringPtr(profileType),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			profile, response, err := postureManagementService.GetProfile(getProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`UpdateProfiles - Update a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProfiles(updateProfilesOptions *UpdateProfilesOptions)`, func() {

			updateProfilesOptions := &posturemanagementv2.UpdateProfilesOptions{
				ID:            core.StringPtr(profileID),
				Name:          core.StringPtr("CUSTOM PROFILE SDK UPDATE"),
				Description:   core.StringPtr("CUSTOM PROFILE SDK UPDATE"),
				BaseProfile:   core.StringPtr(""),
				Type:          core.StringPtr("custom"),
				IsEnabled:     core.BoolPtr(true),
				ControlIds:    []string{"1000101"},
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			profile, response, err := postureManagementService.UpdateProfiles(updateProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetProfileControls - View profile controls`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileControls(getProfileControlsOptions *GetProfileControlsOptions) with pagination`, func() {
			var result posturemanagementv2.ControlList

			getProfileControlsOptions := &posturemanagementv2.GetProfileControlsOptions{
				ProfileID:     core.StringPtr(profileID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
				Offset:        core.Int64Ptr(int64(2)),
				Limit:         core.Int64Ptr(int64(3)),
			}

			getProfileControlsOptions.Offset = nil
			getProfileControlsOptions.Limit = core.Int64Ptr(1)

			controlList, response, err := postureManagementService.GetProfileControls(getProfileControlsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlList).ToNot(BeNil())
			result.Controls = append(result.Controls, controlList.Controls...)

			getProfileControlsOptions.Offset, err = controlList.GetNextOffset()
			Expect(err).To(BeNil())

		})
	})

	Describe(`GetGroupProfileControls - View group profile controls`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetGroupProfileControls(getGroupProfileControlsOptions *GetGroupProfileControlsOptions) with pagination`, func() {
			var result posturemanagementv2.ControlList

			getGroupProfileControlsOptions := &posturemanagementv2.GetGroupProfileControlsOptions{
				GroupID:       core.StringPtr(groupProfileID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
				Offset:        core.Int64Ptr(int64(2)),
				Limit:         core.Int64Ptr(int64(3)),
			}

			getGroupProfileControlsOptions.Offset = nil
			getGroupProfileControlsOptions.Limit = core.Int64Ptr(1)

			controlList, response, err := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlList).ToNot(BeNil())
			result.Controls = append(result.Controls, controlList.Controls...)

			getGroupProfileControlsOptions.Offset, err = controlList.GetNextOffset()
			Expect(err).To(BeNil())

		})
	})

	Describe(`CreateScope - Create a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {

			scopeNameStr := fmt.Sprintf("%s%d", "IBMScope2", time.Now().UnixNano())
			createScopeOptions := &posturemanagementv2.CreateScopeOptions{
				Name:           core.StringPtr(scopeNameStr),
				Description:    core.StringPtr("IBMSchema"),
				CollectorIds:   []string{collectorIDScope},
				CredentialID:   core.StringPtr(credentialIDScope),
				CredentialType: core.StringPtr("ibm"),
				//Interval: core.Int64Ptr(int64(10)),
				//IsDiscoveryScheduled: core.BoolPtr(true),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scope, response, err := postureManagementService.CreateScope(createScopeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())
			scopeID = *(scope.ID)
			correlationID = *(scope.CorrelationID)
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 500
			//
		})
	})

	Describe(`ListScopes - List scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {

			listScopesOptions := &posturemanagementv2.ListScopesOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scopeList, response, err := postureManagementService.ListScopes(listScopesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeList).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetScopeDetails - View scope details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScopeDetails(getScopeDetailsOptions *GetScopeDetailsOptions)`, func() {

			getScopeDetailsOptions := &posturemanagementv2.GetScopeDetailsOptions{
				ID:            core.StringPtr(scopeID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scope, response, err := postureManagementService.GetScopeDetails(getScopeDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`UpdateScopeDetails - Update Scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateScopeDetails(updateScopeDetailsOptions *UpdateScopeDetailsOptions)`, func() {

			updateScopeDetailsOptions := &posturemanagementv2.UpdateScopeDetailsOptions{
				ID:            core.StringPtr(scopeID),
				Name:          core.StringPtr("Scope Test1"),
				Description:   core.StringPtr("Scope Description"),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scope, response, err := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetScopeTimeline - Get scope timelines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScopeTimeline(getScopeTimelineOptions *GetScopeTimelineOptions)`, func() {

			getScopeTimelineOptions := &posturemanagementv2.GetScopeTimelineOptions{
				ScopeID:       core.StringPtr(scopeIDScan),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			eventList, response, err := postureManagementService.GetScopeTimeline(getScopeTimelineOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(eventList).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetScopeDetailsCredentials - Get a scope's credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScopeDetailsCredentials(getScopeDetailsCredentialsOptions *GetScopeDetailsCredentialsOptions)`, func() {

			getScopeDetailsCredentialsOptions := &posturemanagementv2.GetScopeDetailsCredentialsOptions{
				ScopeID:       core.StringPtr(scopeID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scopeCredential, response, err := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCredential).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ReplaceScopeDetailsCredentials - Update a scope's credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptions *ReplaceScopeDetailsCredentialsOptions)`, func() {

			replaceScopeDetailsCredentialsOptions := &posturemanagementv2.ReplaceScopeDetailsCredentialsOptions{
				ScopeID:             core.StringPtr(scopeID),
				CredentialID:        core.StringPtr(credentialIDScopeUpdate),
				CredentialAttribute: core.StringPtr("Credentials attribute"),
				AccountID:           core.StringPtr(accountID),
				TransactionID:       core.StringPtr(transactionID),
			}

			scopeCredential, response, err := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCredential).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetScopeDetailsCollector - Get a scope's collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScopeDetailsCollector(getScopeDetailsCollectorOptions *GetScopeDetailsCollectorOptions)`, func() {

			getScopeDetailsCollectorOptions := &posturemanagementv2.GetScopeDetailsCollectorOptions{
				ScopeID:       core.StringPtr(scopeID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scopeCollector, response, err := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCollector).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ReplaceScopeDetailsCollector - Update a scope's collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptions *ReplaceScopeDetailsCollectorOptions)`, func() {

			replaceScopeDetailsCollectorOptions := &posturemanagementv2.ReplaceScopeDetailsCollectorOptions{
				ScopeID:       core.StringPtr(scopeID),
				CollectorIds:  []string{collectorIDScopeUpdate},
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scopeCollector, response, err := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCollector).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetCorrelationID - Get status of a scope by giving correlation ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCorrelationID(getCorrelationIDOptions *GetCorrelationIDOptions)`, func() {

			getCorrelationIDOptions := &posturemanagementv2.GetCorrelationIDOptions{
				CorrelationID: core.StringPtr(correlationID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			scopeTaskStatus, response, err := postureManagementService.GetCorrelationID(getCorrelationIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeTaskStatus).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ListLatestScans - List latest scans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLatestScans(listLatestScansOptions *ListLatestScansOptions) with pagination`, func() {
			var result posturemanagementv2.ScanList

			listLatestScansOptions := &posturemanagementv2.ListLatestScansOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
				Offset:        core.Int64Ptr(int64(2)),
				Limit:         core.Int64Ptr(int64(3)),
			}

			listLatestScansOptions.Offset = nil
			listLatestScansOptions.Limit = core.Int64Ptr(1)

			scanList, response, err := postureManagementService.ListLatestScans(listLatestScansOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanList).ToNot(BeNil())
			result.LatestScans = append(result.LatestScans, scanList.LatestScans...)

			listLatestScansOptions.Offset, err = scanList.GetNextOffset()
			Expect(err).To(BeNil())

		})
	})

	Describe(`CreateValidation - Initiate a validation scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateValidation(createValidationOptions *CreateValidationOptions)`, func() {

			createValidationOptions := &posturemanagementv2.CreateValidationOptions{
				ScopeID:   core.StringPtr(scopeIDScan),
				ProfileID: core.StringPtr(profileID),
				//GroupProfileID: core.StringPtr("13"),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			result, response, err := postureManagementService.CreateValidation(createValidationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ScansSummary - View a specified scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScansSummary(scansSummaryOptions *ScansSummaryOptions)`, func() {

			scansSummaryOptions := &posturemanagementv2.ScansSummaryOptions{
				ScanID:        core.StringPtr(scanID),
				ProfileID:     core.StringPtr(profileIDScan),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			summary, response, err := postureManagementService.ScansSummary(scansSummaryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summary).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`ScanSummaries - View scan summaries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScanSummaries(scanSummariesOptions *ScanSummariesOptions) with pagination`, func() {
			var result posturemanagementv2.SummaryList

			scanSummariesOptions := &posturemanagementv2.ScanSummariesOptions{
				ReportSettingID: core.StringPtr(reportSettingID),
				AccountID:       core.StringPtr(accountID),
				TransactionID:   core.StringPtr(transactionID),
				Offset:          core.Int64Ptr(int64(2)),
				Limit:           core.Int64Ptr(int64(3)),
			}

			scanSummariesOptions.Offset = nil
			scanSummariesOptions.Limit = core.Int64Ptr(1)

			summaryList, response, err := postureManagementService.ScanSummaries(scanSummariesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summaryList).ToNot(BeNil())
			result.Summaries = append(result.Summaries, summaryList.Summaries...)

			scanSummariesOptions.Offset, err = summaryList.GetNextOffset()
			Expect(err).To(BeNil())

		})
	})

	Describe(`DeleteScope - Delete a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteScope(deleteScopeOptions *DeleteScopeOptions)`, func() {

			deleteScopeOptions := &posturemanagementv2.DeleteScopeOptions{
				ID:            core.StringPtr(scopeID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			response, err := postureManagementService.DeleteScope(deleteScopeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`DeleteProfile - Delete a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {

			deleteProfileOptions := &posturemanagementv2.DeleteProfileOptions{
				ID:            core.StringPtr(profileID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			response, err := postureManagementService.DeleteProfile(deleteProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`DeleteCredential - Delete a credential`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCredential(deleteCredentialOptions *DeleteCredentialOptions)`, func() {

			deleteCredentialOptions := &posturemanagementv2.DeleteCredentialOptions{
				ID:            core.StringPtr(credentialID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			response, err := postureManagementService.DeleteCredential(deleteCredentialOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`DeleteCollector - Delete a collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCollector(deleteCollectorOptions *DeleteCollectorOptions)`, func() {

			deleteCollectorOptions := &posturemanagementv2.DeleteCollectorOptions{
				ID:            core.StringPtr(collectorID),
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr(transactionID),
			}

			response, err := postureManagementService.DeleteCollector(deleteCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//

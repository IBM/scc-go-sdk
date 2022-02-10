//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/posturemanagementv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
const externalConfigFile = "../posture_management_v2.env"

var (
	postureManagementService *posturemanagementv2.PostureManagementV2
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`PostureManagementV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(posturemanagementv2.DefaultServiceName)
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

			postureManagementServiceOptions := &posturemanagementv2.PostureManagementV2Options{}

			postureManagementService, err = posturemanagementv2.NewPostureManagementV2UsingExternalConfig(postureManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(postureManagementService).ToNot(BeNil())
		})
	})

	Describe(`PostureManagementV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCredential request example`, func() {
			fmt.Println("\nCreateCredential() result:")
			// begin-create_credential

			newCredentialDisplayFieldsModel := &posturemanagementv2.NewCredentialDisplayFields{
				Username: core.StringPtr("test"),
				Password: core.StringPtr("**********"),
			}

			credentialGroupModel := &posturemanagementv2.CredentialGroup{
				ID:         core.StringPtr("1"),
				Passphrase: core.StringPtr("passphrase"),
			}

			createCredentialOptions := postureManagementService.NewCreateCredentialOptions(
				true,
				"username_password",
				"test_create",
				"This credential is used for testing.",
				newCredentialDisplayFieldsModel,
				credentialGroupModel,
				"discovery_fact_collection_remediation",
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
		It(`ListCredentials request example`, func() {
			fmt.Println("\nListCredentials() result:")
			// begin-list_credentials

			listCredentialsOptions := postureManagementService.NewListCredentialsOptions()

			credentialList, response, err := postureManagementService.ListCredentials(listCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentialList, "", "  ")
			fmt.Println(string(b))

			// end-list_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialList).ToNot(BeNil())

		})
		It(`GetCredential request example`, func() {
			fmt.Println("\nGetCredential() result:")
			// begin-get_credential

			getCredentialOptions := postureManagementService.NewGetCredentialOptions(
				"testString",
			)

			credential, response, err := postureManagementService.GetCredential(getCredentialOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credential, "", "  ")
			fmt.Println(string(b))

			// end-get_credential

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credential).ToNot(BeNil())

		})
		It(`UpdateCredential request example`, func() {
			fmt.Println("\nUpdateCredential() result:")
			// begin-update_credential

			updateCredentialDisplayFieldsModel := &posturemanagementv2.UpdateCredentialDisplayFields{
				Username: core.StringPtr("test"),
				Password: core.StringPtr("**********"),
			}

			updateCredentialOptions := postureManagementService.NewUpdateCredentialOptions(
				"testString",
			)
			updateCredentialOptions.SetEnabled(true)
			updateCredentialOptions.SetType("username_password")
			updateCredentialOptions.SetName("test_create")
			updateCredentialOptions.SetDescription("This credential is used for testing.")
			updateCredentialOptions.SetDisplayFields(updateCredentialDisplayFieldsModel)
			updateCredentialOptions.SetPurpose("discovery_fact_collection_remediation")

			credential, response, err := postureManagementService.UpdateCredential(updateCredentialOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credential, "", "  ")
			fmt.Println(string(b))

			// end-update_credential

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(credential).ToNot(BeNil())

		})
		It(`CreateCollector request example`, func() {
			fmt.Println("\nCreateCollector() result:")
			// begin-create_collector

			createCollectorOptions := postureManagementService.NewCreateCollectorOptions(
				"IBM-collector-sample",
				true,
				"ibm",
			)
			createCollectorOptions.SetDescription("sample collector")
			createCollectorOptions.SetPassphrase("secret")

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
		It(`ListCollectors request example`, func() {
			fmt.Println("\nListCollectors() result:")
			// begin-list_collectors

			listCollectorsOptions := postureManagementService.NewListCollectorsOptions()

			collectorList, response, err := postureManagementService.ListCollectors(listCollectorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectorList, "", "  ")
			fmt.Println(string(b))

			// end-list_collectors

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectorList).ToNot(BeNil())

		})
		It(`GetCollector request example`, func() {
			fmt.Println("\nGetCollector() result:")
			// begin-get_collector

			getCollectorOptions := postureManagementService.NewGetCollectorOptions(
				"testString",
			)

			collector, response, err := postureManagementService.GetCollector(getCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collector, "", "  ")
			fmt.Println(string(b))

			// end-get_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collector).ToNot(BeNil())

		})
		It(`UpdateCollector request example`, func() {
			fmt.Println("\nUpdateCollector() result:")
			// begin-update_collector

			collectorUpdateModel := &posturemanagementv2.CollectorUpdate{
				DisplayName: core.StringPtr("test-0112-collector_jj"),
				Description: core.StringPtr("This collector is used for testing."),
			}
			collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCollectorOptions := postureManagementService.NewUpdateCollectorOptions(
				"testString",
				collectorUpdateModelAsPatch,
			)

			collector, response, err := postureManagementService.UpdateCollector(updateCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collector, "", "  ")
			fmt.Println(string(b))

			// end-update_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collector).ToNot(BeNil())

		})
		It(`ImportProfiles request example`, func() {
			fmt.Println("\nImportProfiles() result:")
			// begin-import_profiles

			importProfilesOptions := postureManagementService.NewImportProfilesOptions(
				CreateMockReader("This is a mock file."),
			)

			basicResult, response, err := postureManagementService.ImportProfiles(importProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(basicResult, "", "  ")
			fmt.Println(string(b))

			// end-import_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(basicResult).ToNot(BeNil())

		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := postureManagementService.NewListProfilesOptions()

			profileList, response, err := postureManagementService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileList, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileList).ToNot(BeNil())

		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := postureManagementService.NewGetProfileOptions(
				"testString",
				"predefined",
			)

			profile, response, err := postureManagementService.GetProfile(getProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-get_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`UpdateProfiles request example`, func() {
			fmt.Println("\nUpdateProfiles() result:")
			// begin-update_profiles

			updateProfilesOptions := postureManagementService.NewUpdateProfilesOptions(
				"testString",
			)

			profile, response, err := postureManagementService.UpdateProfiles(updateProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-update_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`GetProfileControls request example`, func() {
			fmt.Println("\nGetProfileControls() result:")
			// begin-get_profile_controls

			getProfileControlsOptions := postureManagementService.NewGetProfileControlsOptions(
				"testString",
			)

			controlList, response, err := postureManagementService.GetProfileControls(getProfileControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlList, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlList).ToNot(BeNil())

		})
		It(`GetGroupProfileControls request example`, func() {
			fmt.Println("\nGetGroupProfileControls() result:")
			// begin-get_group_profile_controls

			getGroupProfileControlsOptions := postureManagementService.NewGetGroupProfileControlsOptions(
				"testString",
			)

			controlList, response, err := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlList, "", "  ")
			fmt.Println(string(b))

			// end-get_group_profile_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlList).ToNot(BeNil())

		})
		It(`CreateScope request example`, func() {
			fmt.Println("\nCreateScope() result:")
			// begin-create_scope

			createScopeOptions := postureManagementService.NewCreateScopeOptions(
				"IBMSchema-new-048-test",
				"IBMSchema",
				[]string{"20"},
				"5",
				"on_premise",
			)
			createScopeOptions.SetInterval(int64(10))
			createScopeOptions.SetIsDiscoveryScheduled(true)

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

			scopeList, response, err := postureManagementService.ListScopes(listScopesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeList, "", "  ")
			fmt.Println(string(b))

			// end-list_scopes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeList).ToNot(BeNil())

		})
		It(`GetScopeDetails request example`, func() {
			fmt.Println("\nGetScopeDetails() result:")
			// begin-get_scope_details

			getScopeDetailsOptions := postureManagementService.NewGetScopeDetailsOptions(
				"testString",
			)

			scope, response, err := postureManagementService.GetScopeDetails(getScopeDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-get_scope_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())

		})
		It(`UpdateScopeDetails request example`, func() {
			fmt.Println("\nUpdateScopeDetails() result:")
			// begin-update_scope_details

			updateScopeDetailsOptions := postureManagementService.NewUpdateScopeDetailsOptions(
				"testString",
			)
			updateScopeDetailsOptions.SetName("Scope Test1")
			updateScopeDetailsOptions.SetDescription("Scope Description")

			scope, response, err := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-update_scope_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())

		})
		It(`GetScopeTimeline request example`, func() {
			fmt.Println("\nGetScopeTimeline() result:")
			// begin-get_scope_timeline

			getScopeTimelineOptions := postureManagementService.NewGetScopeTimelineOptions(
				"testString",
			)

			eventList, response, err := postureManagementService.GetScopeTimeline(getScopeTimelineOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(eventList, "", "  ")
			fmt.Println(string(b))

			// end-get_scope_timeline

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(eventList).ToNot(BeNil())

		})
		It(`GetScopeDetailsCredentials request example`, func() {
			fmt.Println("\nGetScopeDetailsCredentials() result:")
			// begin-get_scope_details_credentials

			getScopeDetailsCredentialsOptions := postureManagementService.NewGetScopeDetailsCredentialsOptions(
				"testString",
			)

			scopeCredential, response, err := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeCredential, "", "  ")
			fmt.Println(string(b))

			// end-get_scope_details_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCredential).ToNot(BeNil())

		})
		It(`ReplaceScopeDetailsCredentials request example`, func() {
			fmt.Println("\nReplaceScopeDetailsCredentials() result:")
			// begin-replace_scope_details_credentials

			replaceScopeDetailsCredentialsOptions := postureManagementService.NewReplaceScopeDetailsCredentialsOptions(
				"testString",
				"1",
			)
			replaceScopeDetailsCredentialsOptions.SetCredentialAttribute("Credentials attribute")

			scopeCredential, response, err := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeCredential, "", "  ")
			fmt.Println(string(b))

			// end-replace_scope_details_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCredential).ToNot(BeNil())

		})
		It(`GetScopeDetailsCollector request example`, func() {
			fmt.Println("\nGetScopeDetailsCollector() result:")
			// begin-get_scope_details_collector

			getScopeDetailsCollectorOptions := postureManagementService.NewGetScopeDetailsCollectorOptions(
				"testString",
			)

			scopeCollector, response, err := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeCollector, "", "  ")
			fmt.Println(string(b))

			// end-get_scope_details_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCollector).ToNot(BeNil())

		})
		It(`ReplaceScopeDetailsCollector request example`, func() {
			fmt.Println("\nReplaceScopeDetailsCollector() result:")
			// begin-replace_scope_details_collector

			replaceScopeDetailsCollectorOptions := postureManagementService.NewReplaceScopeDetailsCollectorOptions(
				"testString",
				[]string{"7"},
			)

			scopeCollector, response, err := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeCollector, "", "  ")
			fmt.Println(string(b))

			// end-replace_scope_details_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeCollector).ToNot(BeNil())

		})
		It(`GetCorrelationID request example`, func() {
			fmt.Println("\nGetCorrelationID() result:")
			// begin-get_correlation_id

			getCorrelationIDOptions := postureManagementService.NewGetCorrelationIDOptions(
				"testString",
			)

			scopeTaskStatus, response, err := postureManagementService.GetCorrelationID(getCorrelationIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scopeTaskStatus, "", "  ")
			fmt.Println(string(b))

			// end-get_correlation_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scopeTaskStatus).ToNot(BeNil())

		})
		It(`ListLatestScans request example`, func() {
			fmt.Println("\nListLatestScans() result:")
			// begin-list_latest_scans

			listLatestScansOptions := postureManagementService.NewListLatestScansOptions()

			scanList, response, err := postureManagementService.ListLatestScans(listLatestScansOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanList, "", "  ")
			fmt.Println(string(b))

			// end-list_latest_scans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanList).ToNot(BeNil())

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
			)

			summaryList, response, err := postureManagementService.ScanSummaries(scanSummariesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(summaryList, "", "  ")
			fmt.Println(string(b))

			// end-scan_summaries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(summaryList).ToNot(BeNil())

		})
		It(`DeleteScope request example`, func() {
			// begin-delete_scope

			deleteScopeOptions := postureManagementService.NewDeleteScopeOptions(
				"testString",
			)

			response, err := postureManagementService.DeleteScope(deleteScopeOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_scope
			fmt.Printf("\nDeleteScope() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteProfile request example`, func() {
			// begin-delete_profile

			deleteProfileOptions := postureManagementService.NewDeleteProfileOptions(
				"testString",
			)

			response, err := postureManagementService.DeleteProfile(deleteProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_profile
			fmt.Printf("\nDeleteProfile() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCredential request example`, func() {
			// begin-delete_credential

			deleteCredentialOptions := postureManagementService.NewDeleteCredentialOptions(
				"testString",
			)

			response, err := postureManagementService.DeleteCredential(deleteCredentialOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_credential
			fmt.Printf("\nDeleteCredential() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCollector request example`, func() {
			// begin-delete_collector

			deleteCollectorOptions := postureManagementService.NewDeleteCollectorOptions(
				"testString",
			)

			response, err := postureManagementService.DeleteCollector(deleteCollectorOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_collector
			fmt.Printf("\nDeleteCollector() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

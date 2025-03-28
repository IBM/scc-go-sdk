//go:build examples

/**
 * (C) Copyright IBM Corp. 2025.
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

package securityandcompliancecenterapiv3_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Security and Compliance Center API service.
//
// The following configuration properties are assumed to be defined:
// SECURITY_AND_COMPLIANCE_CENTER_API_URL=<service base url>
// SECURITY_AND_COMPLIANCE_CENTER_API_AUTH_TYPE=iam
// SECURITY_AND_COMPLIANCE_CENTER_API_APIKEY=<IAM apikey>
// SECURITY_AND_COMPLIANCE_CENTER_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`SecurityAndComplianceCenterAPIV3 Examples Tests`, func() {

	const externalConfigFile = "../security_and_compliance_center_api_v3.env"

	var (
		securityAndComplianceCenterAPIService *securityandcompliancecenterapiv3.SecurityAndComplianceCenterAPIV3
		config                                map[string]string

		// Variables to hold link values
		accountIDForReportLink                     string
		attachmentIDForReportLink                  string
		attachmentIDLink                           string
		controlLibraryIDLink                       string
		eTagLink                                   string
		eventNotificationsCRNForUpdateSettingsLink string
		groupIDForReportLink                       string
		objectStorageBucketForUpdateSettingsLink   string
		objectStorageCRNForUpdateSettingsLink      string
		objectStorageLocationForUpdateSettingsLink string
		profileIDForReportLink                     string
		profileIDLink                              string
		providerTypeInstanceIDLink                 string
		reportIDForReportLink                      string
		ruleIDLink                                 string
		scanIDForScanReportLink                    string
		scopeIDLink                                string
		subScopeIDLink                             string
		targetIDLink                               string
		typeForReportLink                          string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securityandcompliancecenterapiv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			securityAndComplianceCenterAPIServiceOptions := &securityandcompliancecenterapiv3.SecurityAndComplianceCenterAPIV3Options{}

			securityAndComplianceCenterAPIService, err = securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterAPIV3UsingExternalConfig(securityAndComplianceCenterAPIServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(securityAndComplianceCenterAPIService).ToNot(BeNil())
		})
	})

	Describe(`SecurityAndComplianceCenterAPIV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := securityAndComplianceCenterAPIService.NewGetSettingsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			settings, response, err := securityAndComplianceCenterAPIService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())

			eventNotificationsCRNForUpdateSettingsLink = *settings.EventNotifications.InstanceCRN
			fmt.Fprintf(GinkgoWriter, "Saved eventNotificationsCRNForUpdateSettingsLink value: %v\n", eventNotificationsCRNForUpdateSettingsLink)
			objectStorageCRNForUpdateSettingsLink = *settings.ObjectStorage.InstanceCRN
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageCRNForUpdateSettingsLink value: %v\n", objectStorageCRNForUpdateSettingsLink)
			objectStorageBucketForUpdateSettingsLink = *settings.ObjectStorage.Bucket
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageBucketForUpdateSettingsLink value: %v\n", objectStorageBucketForUpdateSettingsLink)
			objectStorageLocationForUpdateSettingsLink = *settings.ObjectStorage.BucketLocation
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageLocationForUpdateSettingsLink value: %v\n", objectStorageLocationForUpdateSettingsLink)
		})
		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			objectStoragePrototypeModel := &securityandcompliancecenterapiv3.ObjectStoragePrototype{
				Bucket:      core.StringPtr("px-scan-results"),
				InstanceCRN: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"),
			}

			eventNotificationsPrototypeModel := &securityandcompliancecenterapiv3.EventNotificationsPrototype{
				InstanceCRN:       core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"),
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("scc-sdk-integration"),
			}

			updateSettingsOptions := securityAndComplianceCenterAPIService.NewUpdateSettingsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			updateSettingsOptions.SetObjectStorage(objectStoragePrototypeModel)
			updateSettingsOptions.SetEventNotifications(eventNotificationsPrototypeModel)

			settings, response, err := securityAndComplianceCenterAPIService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(200), Equal(204)))
			if response.StatusCode == 200 {
				Expect(settings).ToNot(BeNil())
			} else {
				Expect(settings).To(BeNil())
			}
		})
		It(`PostTestEvent request example`, func() {
			fmt.Println("\nPostTestEvent() result:")
			// begin-post_test_event

			postTestEventOptions := securityAndComplianceCenterAPIService.NewPostTestEventOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			testEvent, response, err := securityAndComplianceCenterAPIService.PostTestEvent(postTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
		It(`ListInstanceAttachments request example`, func() {
			fmt.Println("\nListInstanceAttachments() result:")
			// begin-list_instance_attachments
			listInstanceAttachmentsOptions := &securityandcompliancecenterapiv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("6702d85a-6437-4d6f-8701-c0146648787b"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.ProfileAttachment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_instance_attachments
		})
		It(`CreateProfileAttachment request example`, func() {
			fmt.Println("\nCreateProfileAttachment() result:")
			// begin-create_profile_attachment
			parameterList := []securityandcompliancecenterapiv3.Parameter{
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
					ParameterName:        core.StringPtr("tls_version"),
					ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['1.2', '1.3']"),
				},
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Disallowed IPs for ingress to RDP port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("3389"),
				},
				securityandcompliancecenterapiv3.Parameter{
					AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
					ParameterName:        core.StringPtr("exclude_default_security_groups"),
					ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['Default']"),
				},
			}

			attachmentNotificationsControlsModel := &securityandcompliancecenterapiv3.AttachmentNotificationsControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentNotificationsModel := &securityandcompliancecenterapiv3.AttachmentNotifications{
				Enabled:  core.BoolPtr(true),
				Controls: attachmentNotificationsControlsModel,
			}

			multiCloudScopePayloadModel := &securityandcompliancecenterapiv3.MultiCloudScopePayloadByID{
				ID: core.StringPtr("8baad3b5-2e69-4027-9967-efac19508e1c"),
			}

			profileAttachmentBaseModel := &securityandcompliancecenterapiv3.ProfileAttachmentBase{
				AttachmentParameters: parameterList,
				Description:          core.StringPtr("This is a profile attachment targeting IBM CIS Foundation using a SDK"),
				Name:                 core.StringPtr("Profile Attachment for IBM CIS Foundation SDK test"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScopePayloadIntf{multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
			}

			createProfileAttachmentOptions := securityAndComplianceCenterAPIService.NewCreateProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
				[]securityandcompliancecenterapiv3.ProfileAttachmentBase{*profileAttachmentBaseModel},
			)
			createProfileAttachmentOptions.SetNewProfileID("9c265b4a-4cdf-47f1-acd3-17b5808f7f3")

			profileAttachmentResponse, response, err := securityAndComplianceCenterAPIService.CreateProfileAttachment(createProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profileAttachmentResponse).ToNot(BeNil())

			attachmentIDLink = *profileAttachmentResponse.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDLink value: %v\n", attachmentIDLink)
		})
		It(`GetProfileAttachment request example`, func() {
			fmt.Println("\nGetProfileAttachment() result:")
			// begin-get_profile_attachment

			getProfileAttachmentOptions := securityAndComplianceCenterAPIService.NewGetProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
				attachmentIDLink,
			)

			profileAttachment, response, err := securityAndComplianceCenterAPIService.GetProfileAttachment(getProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`ReplaceProfileAttachment request example`, func() {
			fmt.Println("\nReplaceProfileAttachment() result:")
			// begin-replace_profile_attachment

			parameterList := []securityandcompliancecenterapiv3.Parameter{
				{
					AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
					ParameterName:        core.StringPtr("tls_version"),
					ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['1.2', '1.3']"),
				},
				{
					AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Disallowed IPs for ingress to RDP port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("3389"),
				},
				{
					AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
					ParameterName:        core.StringPtr("exclude_default_security_groups"),
					ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['Default']"),
				},
			}

			attachmentNotificationsControlsModel := &securityandcompliancecenterapiv3.AttachmentNotificationsControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentNotificationsModel := &securityandcompliancecenterapiv3.AttachmentNotifications{
				Enabled:  core.BoolPtr(true),
				Controls: attachmentNotificationsControlsModel,
			}

			multiCloudScopePayloadModel := &securityandcompliancecenterapiv3.MultiCloudScopePayloadByID{
				ID: core.StringPtr("8baad3b5-2e69-4027-9967-efac19508e1c"),
			}

			replaceProfileAttachmentOptions := securityAndComplianceCenterAPIService.NewReplaceProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
				attachmentIDLink,
				parameterList,
				"testString",
				"testString",
				attachmentNotificationsModel,
				"daily",
				[]securityandcompliancecenterapiv3.MultiCloudScopePayloadIntf{multiCloudScopePayloadModel},
				"enabled",
			)

			profileAttachment, response, err := securityAndComplianceCenterAPIService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`UpgradeAttachment request example`, func() {
			fmt.Println("\nUpgradeAttachment() result:")
			// begin-upgrade_attachment
			parameterList := []securityandcompliancecenterapiv3.Parameter{
				{
					AssessmentID:         core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
					ParameterName:        core.StringPtr("tls_version"),
					ParameterDisplayName: core.StringPtr("IBM Cloud Internet Services TLS version"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['1.2', '1.3']"),
				},
				{
					AssessmentID:         core.StringPtr("rule-f9137be8-2490-4afb-8cd5-a201cb167eb2"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Network ACL rule for allowed IPs to SSH port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-9653d2c7-6290-4128-a5a3-65487ba40370"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Security group rule RDP allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-7c5f6385-67e4-4edf-bec8-c722558b2dec"),
					ParameterName:        core.StringPtr("ssh_port"),
					ParameterDisplayName: core.StringPtr("Security group rule SSH allow port number"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("22"),
				},
				{
					AssessmentID:         core.StringPtr("rule-f1e80ee7-88d5-4bf2-b42f-c863bb24601c"),
					ParameterName:        core.StringPtr("rdp_port"),
					ParameterDisplayName: core.StringPtr("Disallowed IPs for ingress to RDP port"),
					ParameterType:        core.StringPtr("numeric"),
					ParameterValue:       core.StringPtr("3389"),
				},
				{
					AssessmentID:         core.StringPtr("rule-96527f89-1867-4581-b923-1400e04661e0"),
					ParameterName:        core.StringPtr("exclude_default_security_groups"),
					ParameterDisplayName: core.StringPtr("Exclude the default security groups"),
					ParameterType:        core.StringPtr("string_list"),
					ParameterValue:       core.StringPtr("['Default']"),
				},
			}

			upgradeAttachmentOptions := securityAndComplianceCenterAPIService.NewUpgradeAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
				attachmentIDLink,
				parameterList,
			)

			profileAttachment, response, err := securityAndComplianceCenterAPIService.UpgradeAttachment(upgradeAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-upgrade_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`CreateScan request example`, func() {
			fmt.Println("\nCreateScan() result:")
			// begin-create_scan

			createScanOptions := securityAndComplianceCenterAPIService.NewCreateScanOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			createScanOptions.SetAttachmentID("4deb572c-9f37-4126-9cc0-d550672533cb")

			createScanResponse, response, err := securityAndComplianceCenterAPIService.CreateScan(createScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createScanResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createScanResponse).ToNot(BeNil())
		})
		It(`CreateControlLibrary request example`, func() {
			fmt.Println("\nCreateControlLibrary() result:")
			// begin-create_control_library

			assessmentPrototypeModel := &securityandcompliancecenterapiv3.AssessmentPrototype{
				AssessmentID:          core.StringPtr("rule-d1bd9f3f-bee1-46c5-9533-da8bba9eed4e"),
				AssessmentDescription: core.StringPtr("This rule will check on regulation"),
			}

			controlSpecificationPrototypeModel := &securityandcompliancecenterapiv3.ControlSpecificationPrototype{
				ComponentID:                     core.StringPtr("apprapp"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("This field is used to describe a control specification"),
				Assessments:                     []securityandcompliancecenterapiv3.AssessmentPrototype{*assessmentPrototypeModel},
			}

			controlDocModel := &securityandcompliancecenterapiv3.ControlDoc{}

			controlPrototypeModel := &securityandcompliancecenterapiv3.ControlPrototype{
				ControlName:           core.StringPtr("security"),
				ControlDescription:    core.StringPtr("This is a description of a control"),
				ControlCategory:       core.StringPtr("test-control"),
				ControlRequirement:    core.BoolPtr(true),
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("disabled"),
			}

			createControlLibraryOptions := securityAndComplianceCenterAPIService.NewCreateControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"custom control library from SDK",
				"This is a custom control library made from the SDK test framework",
				"custom",
				"0.0.1",
				[]securityandcompliancecenterapiv3.ControlPrototype{*controlPrototypeModel},
			)

			controlLibrary, response, err := securityAndComplianceCenterAPIService.CreateControlLibrary(createControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-create_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIDLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIDLink value: %v\n", controlLibraryIDLink)
		})
		It(`ListControlLibraries request example`, func() {
			fmt.Println("\nListControlLibraries() result:")
			// begin-list_control_libraries
			listControlLibrariesOptions := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterAPIService.NewControlLibrariesPager(listControlLibrariesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.ControlLibrary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_control_libraries
		})
		It(`ReplaceCustomControlLibrary request example`, func() {
			fmt.Println("\nReplaceCustomControlLibrary() result:")
			// begin-replace_custom_control_library

			assessmentPrototypeModel := &securityandcompliancecenterapiv3.AssessmentPrototype{
				AssessmentID:          core.StringPtr("rule-d1bd9f3f-bee1-46c5-9533-da8bba9eed4e"),
				AssessmentDescription: core.StringPtr("This rule will check on regulation"),
			}

			controlSpecificationPrototypeModel := &securityandcompliancecenterapiv3.ControlSpecificationPrototype{
				ComponentID:                     core.StringPtr("apprapp"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("This field is used to describe a control specification"),
				Assessments:                     []securityandcompliancecenterapiv3.AssessmentPrototype{*assessmentPrototypeModel},
			}

			controlDocModel := &securityandcompliancecenterapiv3.ControlDoc{}

			controlPrototypeModel := &securityandcompliancecenterapiv3.ControlPrototype{
				ControlName:           core.StringPtr("security"),
				ControlDescription:    core.StringPtr("This is a description of a control"),
				ControlCategory:       core.StringPtr("test-control"),
				ControlRequirement:    core.BoolPtr(true),
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("disabled"),
			}

			replaceCustomControlLibraryOptions := securityAndComplianceCenterAPIService.NewReplaceCustomControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
				"custom control library from SDK",
				"This is a custom control library made from the SDK test framework",
				"custom",
				"0.0.2",
				[]securityandcompliancecenterapiv3.ControlPrototype{*controlPrototypeModel},
			)

			controlLibrary, response, err := securityAndComplianceCenterAPIService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-replace_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`GetControlLibrary request example`, func() {
			fmt.Println("\nGetControlLibrary() result:")
			// begin-get_control_library

			getControlLibraryOptions := securityAndComplianceCenterAPIService.NewGetControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
			)

			controlLibrary, response, err := securityAndComplianceCenterAPIService.GetControlLibrary(getControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-get_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile
			profileControlsPrototypeModels := []securityandcompliancecenterapiv3.ProfileControlsPrototype{
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("2ce21ba3-0548-49a3-88e2-1122632218f4"),
				},
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("bdc5fdab-6934-461c-8bb1-9af7ed8e8d33"),
				},
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("60dae3b5-6104-4b3e-bac7-26cc7b741aca"),
				},
			}

			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr("[\"1.2\",\"1.3\"]"),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			createProfileOptions := securityAndComplianceCenterAPIService.NewCreateProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"Example Profile",
				"0.0.1",
				profileControlsPrototypeModels,
				[]securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
			)
			createProfileOptions.SetProfileDescription("This profile is created as an example of the SDK gen")
			createProfileOptions.SetLatest(true)

			profile, response, err := securityAndComplianceCenterAPIService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())

			profileIDLink = *profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIDLink value: %v\n", profileIDLink)
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles
			listProfilesOptions := &securityandcompliancecenterapiv3.ListProfilesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterAPIService.NewProfilesPager(listProfilesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Profile
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_profiles
		})
		It(`ReplaceProfile request example`, func() {
			fmt.Println("\nReplaceProfile() result:")
			// begin-replace_profile
			profileControlsModels := []securityandcompliancecenterapiv3.ProfileControls{
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("2ce21ba3-0548-49a3-88e2-1122632218f4"),
				},
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("bdc5fdab-6934-461c-8bb1-9af7ed8e8d33"),
				},
				{
					ControlLibraryID: core.StringPtr("a046fb6b-aba5-4646-b190-a2c76241e7af"),
					ControlID:        core.StringPtr("60dae3b5-6104-4b3e-bac7-26cc7b741aca"),
				},
			}

			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr("[\"1.2\",\"1.3\"]"),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			replaceProfileOptions := securityAndComplianceCenterAPIService.NewReplaceProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
				"custom",
				profileControlsModels,
				[]securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
			)
			replaceProfileOptions.SetNewProfileName("Example Profile Updated")
			replaceProfileOptions.SetNewProfileDescription("This profile has been updated")
			replaceProfileOptions.SetNewProfileVersion("0.0.2")
			replaceProfileOptions.SetNewLatest(true)

			profile, response, err := securityAndComplianceCenterAPIService.ReplaceProfile(replaceProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := securityAndComplianceCenterAPIService.NewGetProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profile, response, err := securityAndComplianceCenterAPIService.GetProfile(getProfileOptions)
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
		It(`ReplaceProfileParameters request example`, func() {
			fmt.Println("\nReplaceProfileParameters() result:")
			// begin-replace_profile_parameters

			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{}

			replaceProfileParametersOptions := securityAndComplianceCenterAPIService.NewReplaceProfileParametersOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
				[]securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
			)

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterAPIService.ReplaceProfileParameters(replaceProfileParametersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileDefaultParametersResponse, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_parameters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
		It(`ListProfileParameters request example`, func() {
			fmt.Println("\nListProfileParameters() result:")
			// begin-list_profile_parameters

			listProfileParametersOptions := securityAndComplianceCenterAPIService.NewListProfileParametersOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterAPIService.ListProfileParameters(listProfileParametersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileDefaultParametersResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_profile_parameters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
		It(`CompareProfiles request example`, func() {
			fmt.Println("\nCompareProfiles() result:")
			// begin-compare_profiles

			compareProfilesOptions := securityAndComplianceCenterAPIService.NewCompareProfilesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"2f598907-970d-4d52-9071-5cc95912f55e",
			)

			comparePredefinedProfilesResponse, response, err := securityAndComplianceCenterAPIService.CompareProfiles(compareProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(comparePredefinedProfilesResponse, "", "  ")
			fmt.Println(string(b))

			// end-compare_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comparePredefinedProfilesResponse).ToNot(BeNil())
		})
		It(`ListProfileAttachments request example`, func() {
			fmt.Println("\nListProfileAttachments() result:")
			// begin-list_profile_attachments

			listProfileAttachmentsOptions := securityAndComplianceCenterAPIService.NewListProfileAttachmentsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
			)

			profileAttachmentCollection, response, err := securityAndComplianceCenterAPIService.ListProfileAttachments(listProfileAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_profile_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachmentCollection).ToNot(BeNil())
		})
		It(`CreateScope request example`, func() {
			fmt.Println("\nCreateScope() result:")
			// begin-create_scope
			scopePropertyModel0 := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("ff88f007f9ff4622aac4fbc0eda36255"),
			}

			scopePropertyModel1 := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			createScopeOptions := securityAndComplianceCenterAPIService.NewCreateScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)
			createScopeOptions.SetName("ibm scope")
			createScopeOptions.SetDescription("The scope that is defined for IBM resources.")
			createScopeOptions.SetEnvironment("ibm-cloud")
			createScopeOptions.SetProperties([]securityandcompliancecenterapiv3.ScopePropertyIntf{scopePropertyModel0, scopePropertyModel1})

			scope, response, err := securityAndComplianceCenterAPIService.CreateScope(createScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-create_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

			scopeIDLink = *scope.ID
			fmt.Fprintf(GinkgoWriter, "Saved scopeIDLink value: %v\n", scopeIDLink)
		})
		It(`ListScopes request example`, func() {
			fmt.Println("\nListScopes() result:")
			// begin-list_scopes
			listScopesOptions := &securityandcompliancecenterapiv3.ListScopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewScopesPager(listScopesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Scope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_scopes
		})
		It(`UpdateScope request example`, func() {
			fmt.Println("\nUpdateScope() result:")
			// begin-update_scope

			updateScopeOptions := securityAndComplianceCenterAPIService.NewUpdateScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)
			updateScopeOptions.SetName("updated name of scope")
			updateScopeOptions.SetDescription("updated scope description")

			scope, response, err := securityAndComplianceCenterAPIService.UpdateScope(updateScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-update_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
		It(`GetScope request example`, func() {
			fmt.Println("\nGetScope() result:")
			// begin-get_scope

			getScopeOptions := securityAndComplianceCenterAPIService.NewGetScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)

			scope, response, err := securityAndComplianceCenterAPIService.GetScope(getScopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scope, "", "  ")
			fmt.Println(string(b))

			// end-get_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
		It(`CreateSubscope request example`, func() {
			fmt.Println("\nCreateSubscope() result:")
			// begin-create_subscope
			scopePropertyModel := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("1f689f08ec9b47b885c2659c17029581"),
			}
			scopeType := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account.resource_group"),
			}

			scopePrototypeModel := &securityandcompliancecenterapiv3.ScopePrototype{
				Name:        core.StringPtr("ibm subscope update"),
				Description: core.StringPtr("The subscope that is defined for IBM resources."),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterapiv3.ScopePropertyIntf{scopePropertyModel, scopeType},
			}

			createSubscopeOptions := securityAndComplianceCenterAPIService.NewCreateSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				[]securityandcompliancecenterapiv3.ScopePrototype{*scopePrototypeModel},
			)

			subScopeResponse, response, err := securityAndComplianceCenterAPIService.CreateSubscope(createSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScopeResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subScopeResponse).ToNot(BeNil())

			subScopeIDLink = *subScopeResponse.Subscopes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved subScopeIDLink value: %v\n", subScopeIDLink)
		})
		It(`ListSubscopes request example`, func() {
			fmt.Println("\nListSubscopes() result:")
			// begin-list_subscopes
			listSubscopesOptions := &securityandcompliancecenterapiv3.ListSubscopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				Limit:       core.Int64Ptr(int64(10)),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewSubscopesPager(listSubscopesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.SubScope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_subscopes
		})
		It(`GetSubscope request example`, func() {
			fmt.Println("\nGetSubscope() result:")
			// begin-get_subscope

			getSubscopeOptions := securityAndComplianceCenterAPIService.NewGetSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)

			subScope, response, err := securityAndComplianceCenterAPIService.GetSubscope(getSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScope, "", "  ")
			fmt.Println(string(b))

			// end-get_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
		It(`UpdateSubscope request example`, func() {
			fmt.Println("\nUpdateSubscope() result:")
			// begin-update_subscope

			updateSubscopeOptions := securityAndComplianceCenterAPIService.NewUpdateSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)
			updateSubscopeOptions.SetName("updated name of scope")
			updateSubscopeOptions.SetDescription("updated scope description")

			subScope, response, err := securityAndComplianceCenterAPIService.UpdateSubscope(updateSubscopeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subScope, "", "  ")
			fmt.Println(string(b))

			// end-update_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			createTargetOptions := securityAndComplianceCenterAPIService.NewCreateTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"62ecf99b240144dea9125666249edfcb",
				"Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54",
				"Target for IBM account",
			)

			target, response, err := securityAndComplianceCenterAPIService.CreateTarget(createTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)
		})
		It(`ListTargets request example`, func() {
			fmt.Println("\nListTargets() result:")
			// begin-list_targets

			listTargetsOptions := securityAndComplianceCenterAPIService.NewListTargetsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			targetCollection, response, err := securityAndComplianceCenterAPIService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
		It(`GetTarget request example`, func() {
			fmt.Println("\nGetTarget() result:")
			// begin-get_target

			getTargetOptions := securityAndComplianceCenterAPIService.NewGetTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
			)

			target, response, err := securityAndComplianceCenterAPIService.GetTarget(getTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-get_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
		It(`ReplaceTarget request example`, func() {
			fmt.Println("\nReplaceTarget() result:")
			// begin-replace_target

			replaceTargetOptions := securityAndComplianceCenterAPIService.NewReplaceTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
				"62ecf99b240144dea9125666249edfcb",
				"Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54",
				"Updated target for IBM account",
			)

			target, response, err := securityAndComplianceCenterAPIService.ReplaceTarget(replaceTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-replace_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
		It(`CreateProviderTypeInstance request example`, func() {
			fmt.Println("\nCreateProviderTypeInstance() result:")
			// begin-create_provider_type_instance

			createProviderTypeInstanceOptions := securityAndComplianceCenterAPIService.NewCreateProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
			)

			createProviderTypeInstanceOptions.SetName("Caveonix-instance-1")
			createProviderTypeInstanceOptions.SetAttributes(map[string]interface{}{})

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-create_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstance).ToNot(BeNil())

			providerTypeInstanceIDLink = *providerTypeInstance.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIDLink value: %v\n", providerTypeInstanceIDLink)
		})
		It(`ListProviderTypeInstances request example`, func() {
			fmt.Println("\nListProviderTypeInstances() result:")
			// begin-list_provider_type_instances

			listProviderTypeInstancesOptions := securityAndComplianceCenterAPIService.NewListProviderTypeInstancesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
			)

			providerTypeInstanceCollection, response, err := securityAndComplianceCenterAPIService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstanceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_type_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceCollection).ToNot(BeNil())
		})
		It(`GetProviderTypeInstance request example`, func() {
			fmt.Println("\nGetProviderTypeInstance() result:")
			// begin-get_provider_type_instance

			getProviderTypeInstanceOptions := securityAndComplianceCenterAPIService.NewGetProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
				providerTypeInstanceIDLink,
			)

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
		It(`UpdateProviderTypeInstance request example`, func() {
			fmt.Println("\nUpdateProviderTypeInstance() result:")
			// begin-update_provider_type_instance

			updateProviderTypeInstanceOptions := securityAndComplianceCenterAPIService.NewUpdateProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
				providerTypeInstanceIDLink,
			)

			updateProviderTypeInstanceOptions.SetName("caveonix-instance-1")
			updateProviderTypeInstanceOptions.SetAttributes(map[string]interface{}{})

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstance, "", "  ")
			fmt.Println(string(b))

			// end-update_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
		It(`ListProviderTypes request example`, func() {
			fmt.Println("\nListProviderTypes() result:")
			// begin-list_provider_types

			listProviderTypesOptions := securityAndComplianceCenterAPIService.NewListProviderTypesOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			providerTypeCollection, response, err := securityAndComplianceCenterAPIService.ListProviderTypes(listProviderTypesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_types

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeCollection).ToNot(BeNil())
		})
		It(`GetProviderTypeByID request example`, func() {
			fmt.Println("\nGetProviderTypeByID() result:")
			// begin-get_provider_type_by_id

			getProviderTypeByIDOptions := securityAndComplianceCenterAPIService.NewGetProviderTypeByIDOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
			)

			providerType, response, err := securityAndComplianceCenterAPIService.GetProviderTypeByID(getProviderTypeByIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerType, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerType).ToNot(BeNil())
		})
		It(`GetLatestReports request example`, func() {
			fmt.Println("\nGetLatestReports() result:")
			// begin-get_latest_reports

			getLatestReportsOptions := securityAndComplianceCenterAPIService.NewGetLatestReportsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			reportLatest, response, err := securityAndComplianceCenterAPIService.GetLatestReports(getLatestReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportLatest, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_reports

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportLatest).ToNot(BeNil())

			accountIDForReportLink = *reportLatest.Reports[0].Account.ID
			fmt.Fprintf(GinkgoWriter, "Saved accountIDForReportLink value: %v\n", accountIDForReportLink)
			reportIDForReportLink = *reportLatest.Reports[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved reportIDForReportLink value: %v\n", reportIDForReportLink)
			attachmentIDForReportLink = *reportLatest.Reports[0].Attachment.ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDForReportLink value: %v\n", attachmentIDForReportLink)
			groupIDForReportLink = *reportLatest.Reports[0].GroupID
			fmt.Fprintf(GinkgoWriter, "Saved groupIDForReportLink value: %v\n", groupIDForReportLink)
			profileIDForReportLink = *reportLatest.Reports[0].Profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIDForReportLink value: %v\n", profileIDForReportLink)
			typeForReportLink = *reportLatest.Reports[0].Type
			fmt.Fprintf(GinkgoWriter, "Saved typeForReportLink value: %v\n", typeForReportLink)
		})
		It(`ListReports request example`, func() {
			fmt.Println("\nListReports() result:")
			// begin-list_reports
			listReportsOptions := &securityandcompliancecenterapiv3.ListReportsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportAttachmentID: &attachmentIDForReportLink,
				GroupID:            &groupIDForReportLink,
				ReportProfileID:    &profileIDForReportLink,
				Type:               &typeForReportLink,
				Limit:              core.Int64Ptr(int64(10)),
				Sort:               core.StringPtr("profile_name"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewReportsPager(listReportsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Report
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_reports
		})
		It(`GetReport request example`, func() {
			fmt.Println("\nGetReport() result:")
			// begin-get_report

			getReportOptions := securityAndComplianceCenterAPIService.NewGetReportOptions(
				reportIDForReportLink,
				"acd7032c-15a3-484f-bf5b-67d41534d940",
			)

			report, response, err := securityAndComplianceCenterAPIService.GetReport(getReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
		It(`GetReportSummary request example`, func() {
			fmt.Println("\nGetReportSummary() result:")
			// begin-get_report_summary

			getReportSummaryOptions := securityAndComplianceCenterAPIService.NewGetReportSummaryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportSummary, response, err := securityAndComplianceCenterAPIService.GetReportSummary(getReportSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_report_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
		It(`GetReportDownloadFile request example`, func() {
			fmt.Println("\nGetReportDownloadFile() result:")
			// begin-get_report_download_file

			getReportDownloadFileOptions := securityAndComplianceCenterAPIService.NewGetReportDownloadFileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			result, response, err := securityAndComplianceCenterAPIService.GetReportDownloadFile(getReportDownloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-get_report_download_file

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`GetReportControls request example`, func() {
			fmt.Println("\nGetReportControls() result:")
			// begin-get_report_controls

			getReportControlsOptions := securityAndComplianceCenterAPIService.NewGetReportControlsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)
			getReportControlsOptions.SetStatus("compliant")

			reportControls, response, err := securityAndComplianceCenterAPIService.GetReportControls(getReportControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportControls, "", "  ")
			fmt.Println(string(b))

			// end-get_report_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportControls).ToNot(BeNil())
		})
		It(`GetReportRule request example`, func() {
			fmt.Println("\nGetReportRule() result:")
			// begin-get_report_rule

			getReportRuleOptions := securityAndComplianceCenterAPIService.NewGetReportRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				"rule-61fa114a-2bb9-43fd-8068-b873b48bdf79",
			)

			ruleInfo, response, err := securityAndComplianceCenterAPIService.GetReportRule(getReportRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleInfo, "", "  ")
			fmt.Println(string(b))

			// end-get_report_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleInfo).ToNot(BeNil())
		})
		It(`ListReportEvaluations request example`, func() {
			fmt.Println("\nListReportEvaluations() result:")
			// begin-list_report_evaluations
			listReportEvaluationsOptions := &securityandcompliancecenterapiv3.ListReportEvaluationsOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:         &reportIDForReportLink,
				AssessmentID:     core.StringPtr("testString"),
				AssessmentMethod: core.StringPtr("testString"),
				ComponentID:      core.StringPtr("testString"),
				TargetID:         core.StringPtr("testString"),
				TargetEnv:        core.StringPtr("testString"),
				TargetName:       core.StringPtr("testString"),
				Status:           core.StringPtr("failure"),
				Limit:            core.Int64Ptr(int64(10)),
				Sort:             core.StringPtr("assessment_id"),
				ScopeID:          core.StringPtr("testString"),
				SubscopeID:       core.StringPtr("testString"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_evaluations
		})
		It(`ListReportResources request example`, func() {
			fmt.Println("\nListReportResources() result:")
			// begin-list_report_resources
			listReportResourcesOptions := &securityandcompliancecenterapiv3.ListReportResourcesOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:     &reportIDForReportLink,
				ID:           core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				AccountID:    &accountIDForReportLink,
				ComponentID:  core.StringPtr("testString"),
				Status:       core.StringPtr("compliant"),
				Sort:         core.StringPtr("account_id"),
				Limit:        core.Int64Ptr(int64(10)),
				ScopeID:      core.StringPtr("testString"),
				SubscopeID:   core.StringPtr("testString"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewReportResourcesPager(listReportResourcesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_report_resources
		})
		It(`GetReportTags request example`, func() {
			fmt.Println("\nGetReportTags() result:")
			// begin-get_report_tags

			getReportTagsOptions := securityAndComplianceCenterAPIService.NewGetReportTagsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportTags, response, err := securityAndComplianceCenterAPIService.GetReportTags(getReportTagsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportTags, "", "  ")
			fmt.Println(string(b))

			// end-get_report_tags

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportTags).ToNot(BeNil())
		})
		It(`GetReportViolationsDrift request example`, func() {
			fmt.Println("\nGetReportViolationsDrift() result:")
			// begin-get_report_violations_drift

			getReportViolationsDriftOptions := securityAndComplianceCenterAPIService.NewGetReportViolationsDriftOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			reportViolationsDrift, response, err := securityAndComplianceCenterAPIService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportViolationsDrift, "", "  ")
			fmt.Println(string(b))

			// end-get_report_violations_drift

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportViolationsDrift).ToNot(BeNil())
		})
		It(`ListScanReports request example`, func() {
			fmt.Println("\nListScanReports() result:")
			// begin-list_scan_reports

			listScanReportsOptions := securityAndComplianceCenterAPIService.NewListScanReportsOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
			)

			scanReportCollection, response, err := securityAndComplianceCenterAPIService.ListScanReports(listScanReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanReportCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_scan_reports

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReportCollection).ToNot(BeNil())

			scanIDForScanReportLink = *scanReportCollection.ScanReports[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved scanIDForScanReportLink value: %v\n", scanIDForScanReportLink)
		})
		It(`CreateScanReport request example`, func() {
			fmt.Println("\nCreateScanReport() result:")
			// begin-create_scan_report

			createScanReportOptions := securityAndComplianceCenterAPIService.NewCreateScanReportOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				"csv",
			)

			createScanReport, response, err := securityAndComplianceCenterAPIService.CreateScanReport(createScanReportOptions)
			b, _ := json.MarshalIndent(createScanReport, "", "  ")
			fmt.Println(string(b))

			// end-create_scan_report

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(409))
			Expect(createScanReport).To(BeNil())
		})
		It(`GetScanReport request example`, func() {
			fmt.Println("\nGetScanReport() result:")
			// begin-get_scan_report

			getScanReportOptions := securityAndComplianceCenterAPIService.NewGetScanReportOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				scanIDForScanReportLink,
			)

			scanReport, response, err := securityAndComplianceCenterAPIService.GetScanReport(getScanReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanReport, "", "  ")
			fmt.Println(string(b))

			// end-get_scan_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReport).ToNot(BeNil())
		})
		It(`GetScanReportDownloadFile request example`, func() {
			fmt.Println("\nGetScanReportDownloadFile() result:")
			// begin-get_scan_report_download_file

			getScanReportDownloadFileOptions := securityAndComplianceCenterAPIService.NewGetScanReportDownloadFileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				reportIDForReportLink,
				scanIDForScanReportLink,
			)

			result, response, err := securityAndComplianceCenterAPIService.GetScanReportDownloadFile(getScanReportDownloadFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-get_scan_report_download_file

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules
			listRulesOptions := &securityandcompliancecenterapiv3.ListRulesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Type:        core.StringPtr("system_defined"),
				Search:      core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				Sort:        core.StringPtr("updated_on"),
			}

			pager, err := securityAndComplianceCenterAPIService.NewRulesPager(listRulesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.Rule
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_rules
		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			additionalTargetAttributeModel := &securityandcompliancecenterapiv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-east"),
			}

			ruleTargetPrototypeModel := &securityandcompliancecenterapiv3.RuleTargetPrototype{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			conditionItemModel := &securityandcompliancecenterapiv3.ConditionItemConditionBase{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterapiv3.RequiredConfigConditionListConditionListConditionAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterapiv3.ConditionItemIntf{conditionItemModel},
			}

			ruleParameterModel := &securityandcompliancecenterapiv3.RuleParameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterapiv3.Import{
				Parameters: []securityandcompliancecenterapiv3.RuleParameter{*ruleParameterModel},
			}

			createRuleOptions := securityAndComplianceCenterAPIService.NewCreateRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"Example rule",
				ruleTargetPrototypeModel,
				requiredConfigModel,
			)
			createRuleOptions.SetVersion("1.0.0")
			createRuleOptions.SetImport(importModel)
			createRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterAPIService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIDLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIDLink value: %v\n", ruleIDLink)
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := securityAndComplianceCenterAPIService.NewGetRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
			)

			rule, response, err := securityAndComplianceCenterAPIService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

			additionalTargetAttributeModel := &securityandcompliancecenterapiv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-south"),
			}

			ruleTargetPrototypeModel := &securityandcompliancecenterapiv3.RuleTargetPrototype{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			conditionItemModel := &securityandcompliancecenterapiv3.ConditionItemConditionBase{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterapiv3.RequiredConfigConditionListConditionListConditionAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterapiv3.ConditionItemIntf{conditionItemModel},
			}

			ruleParameterModel := &securityandcompliancecenterapiv3.RuleParameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterapiv3.Import{
				Parameters: []securityandcompliancecenterapiv3.RuleParameter{*ruleParameterModel},
			}

			replaceRuleOptions := securityAndComplianceCenterAPIService.NewReplaceRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
				eTagLink,
				"Example rule",
				ruleTargetPrototypeModel,
				requiredConfigModel,
			)
			replaceRuleOptions.SetVersion("1.0.1")
			replaceRuleOptions.SetImport(importModel)
			replaceRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterAPIService.ReplaceRule(replaceRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
		It(`ListServices request example`, func() {
			fmt.Println("\nListServices() result:")
			// begin-list_services

			listServicesOptions := securityAndComplianceCenterAPIService.NewListServicesOptions()

			serviceCollection, response, err := securityAndComplianceCenterAPIService.ListServices(listServicesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_services

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceCollection).ToNot(BeNil())
		})
		It(`GetService request example`, func() {
			fmt.Println("\nGetService() result:")
			// begin-get_service

			getServiceOptions := securityAndComplianceCenterAPIService.NewGetServiceOptions(
				"cloud-object-storage",
			)

			service, response, err := securityAndComplianceCenterAPIService.GetService(getServiceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(service, "", "  ")
			fmt.Println(string(b))

			// end-get_service

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(service).ToNot(BeNil())
		})
		It(`DeleteProfileAttachment request example`, func() {
			fmt.Println("\nDeleteProfileAttachment() result:")
			// begin-delete_profile_attachment

			deleteProfileAttachmentOptions := securityAndComplianceCenterAPIService.NewDeleteProfileAttachmentOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"9c265b4a-4cdf-47f1-acd3-17b5808f7f3f",
				attachmentIDLink,
			)

			profileAttachment, response, err := securityAndComplianceCenterAPIService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileAttachment, "", "  ")
			fmt.Println(string(b))

			// end-delete_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
		It(`DeleteCustomControlLibrary request example`, func() {
			fmt.Println("\nDeleteCustomControlLibrary() result:")
			// begin-delete_custom_control_library

			deleteCustomControlLibraryOptions := securityAndComplianceCenterAPIService.NewDeleteCustomControlLibraryOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				controlLibraryIDLink,
			)

			controlLibrary, response, err := securityAndComplianceCenterAPIService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`DeleteCustomProfile request example`, func() {
			fmt.Println("\nDeleteCustomProfile() result:")
			// begin-delete_custom_profile

			deleteCustomProfileOptions := securityAndComplianceCenterAPIService.NewDeleteCustomProfileOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				profileIDLink,
			)

			profile, response, err := securityAndComplianceCenterAPIService.DeleteCustomProfile(deleteCustomProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`DeleteSubscope request example`, func() {
			// begin-delete_subscope

			deleteSubscopeOptions := securityAndComplianceCenterAPIService.NewDeleteSubscopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
				subScopeIDLink,
			)

			response, err := securityAndComplianceCenterAPIService.DeleteSubscope(deleteSubscopeOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSubscope(): %d\n", response.StatusCode)
			}

			// end-delete_subscope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteScope request example`, func() {
			// begin-delete_scope

			deleteScopeOptions := securityAndComplianceCenterAPIService.NewDeleteScopeOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				scopeIDLink,
			)

			response, err := securityAndComplianceCenterAPIService.DeleteScope(deleteScopeOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteScope(): %d\n", response.StatusCode)
			}

			// end-delete_scope

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTarget request example`, func() {
			// begin-delete_target

			deleteTargetOptions := securityAndComplianceCenterAPIService.NewDeleteTargetOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				targetIDLink,
			)

			response, err := securityAndComplianceCenterAPIService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTarget(): %d\n", response.StatusCode)
			}

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProviderTypeInstance request example`, func() {
			// begin-delete_provider_type_instance

			deleteProviderTypeInstanceOptions := securityAndComplianceCenterAPIService.NewDeleteProviderTypeInstanceOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				"3e25966275dccfa2c3a34786919c5af7",
				providerTypeInstanceIDLink,
			)

			response, err := securityAndComplianceCenterAPIService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteProviderTypeInstance(): %d\n", response.StatusCode)
			}

			// end-delete_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := securityAndComplianceCenterAPIService.NewDeleteRuleOptions(
				"acd7032c-15a3-484f-bf5b-67d41534d940",
				ruleIDLink,
			)

			response, err := securityAndComplianceCenterAPIService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRule(): %d\n", response.StatusCode)
			}

			// end-delete_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

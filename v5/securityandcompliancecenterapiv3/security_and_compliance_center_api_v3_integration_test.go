//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the securityandcompliancecenterapiv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SecurityAndComplianceCenterAPIV3 Integration Tests`, func() {
	const externalConfigFile = "../security_and_compliance_center_api_v3.env"

	var (
		err                                   error
		securityAndComplianceCenterAPIService *securityandcompliancecenterapiv3.SecurityAndComplianceCenterAPIV3
		serviceURL                            string
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
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(securityandcompliancecenterapiv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			securityAndComplianceCenterAPIServiceOptions := &securityandcompliancecenterapiv3.SecurityAndComplianceCenterAPIV3Options{}

			securityAndComplianceCenterAPIService, err = securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterAPIV3UsingExternalConfig(securityAndComplianceCenterAPIServiceOptions)
			Expect(err).To(BeNil())
			Expect(securityAndComplianceCenterAPIService).ToNot(BeNil())
			Expect(securityAndComplianceCenterAPIService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			securityAndComplianceCenterAPIService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSettings - List settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
			}

			settings, response, err := securityAndComplianceCenterAPIService.GetSettings(getSettingsOptions)
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
	})

	Describe(`UpdateSettings - Update settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			objectStoragePrototypeModel := &securityandcompliancecenterapiv3.ObjectStoragePrototype{
				Bucket:      core.StringPtr("px-scan-results"),
				InstanceCRN: core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"),
			}

			eventNotificationsPrototypeModel := &securityandcompliancecenterapiv3.EventNotificationsPrototype{
				InstanceCRN:       core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"),
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("scc-sdk-integration"),
			}

			updateSettingsOptions := &securityandcompliancecenterapiv3.UpdateSettingsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ObjectStorage:      objectStoragePrototypeModel,
				EventNotifications: eventNotificationsPrototypeModel,
			}

			settings, response, err := securityAndComplianceCenterAPIService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(200), Equal(204)))
			if response.StatusCode == 200 {
				Expect(settings).ToNot(BeNil())
			} else {
				Expect(settings).To(BeNil())
			}
		})
	})

	Describe(`PostTestEvent - Create a test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
			postTestEventOptions := &securityandcompliancecenterapiv3.PostTestEventOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
			}

			testEvent, response, err := securityAndComplianceCenterAPIService.PostTestEvent(postTestEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceAttachments - Get all instance attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceAttachments(listInstanceAttachmentsOptions *ListInstanceAttachmentsOptions) with pagination`, func() {
			listInstanceAttachmentsOptions := &securityandcompliancecenterapiv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("6702d85a-6437-4d6f-8701-c0146648787b"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
				Start:             core.StringPtr("testString"),
			}

			listInstanceAttachmentsOptions.Start = nil
			listInstanceAttachmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.ProfileAttachment
			for {
				profileAttachmentCollection, response, err := securityAndComplianceCenterAPIService.ListInstanceAttachments(listInstanceAttachmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(profileAttachmentCollection).ToNot(BeNil())
				allResults = append(allResults, profileAttachmentCollection.Attachments...)

				listInstanceAttachmentsOptions.Start, err = profileAttachmentCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listInstanceAttachmentsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListInstanceAttachments(listInstanceAttachmentsOptions *ListInstanceAttachmentsOptions) using InstanceAttachmentsPager`, func() {
			listInstanceAttachmentsOptions := &securityandcompliancecenterapiv3.ListInstanceAttachmentsOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:         &accountIDForReportLink,
				VersionGroupLabel: core.StringPtr("6702d85a-6437-4d6f-8701-c0146648787b"),
				Limit:             core.Int64Ptr(int64(10)),
				Sort:              core.StringPtr("created_on"),
				Direction:         core.StringPtr("desc"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.ProfileAttachment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewInstanceAttachmentsPager(listInstanceAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListInstanceAttachments() returned a total of %d item(s) using InstanceAttachmentsPager.\n", len(allResults))
		})
	})

	Describe(`CreateProfileAttachment - Create an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfileAttachment(createProfileAttachmentOptions *CreateProfileAttachmentOptions)`, func() {
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

			endDate := strfmt.DateTime(time.Now())
			dateRangeModel := &securityandcompliancecenterapiv3.DateRange{
				StartDate: CreateMockDateTime("2025-02-28T05:42:58.000Z"),
				EndDate:   &endDate,
			}

			profileAttachmentBaseModel := &securityandcompliancecenterapiv3.ProfileAttachmentBase{
				AttachmentParameters: parameterList,
				Description:          core.StringPtr("This is a profile attachment targeting IBM CIS Foundation using a SDK"),
				Name:                 core.StringPtr("Profile Attachment for IBM CIS Foundation SDK test"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScopePayloadIntf{multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
				DataSelectionRange:   dateRangeModel,
			}

			createProfileAttachmentOptions := &securityandcompliancecenterapiv3.CreateProfileAttachmentOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:      core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				NewAttachments: []securityandcompliancecenterapiv3.ProfileAttachmentBase{*profileAttachmentBaseModel},
				NewProfileID:   core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3"),
				AccountID:      &accountIDForReportLink,
			}

			profileAttachmentResponse, response, err := securityAndComplianceCenterAPIService.CreateProfileAttachment(createProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profileAttachmentResponse).ToNot(BeNil())

			attachmentIDLink = *profileAttachmentResponse.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIDLink value: %v\n", attachmentIDLink)
		})
	})

	Describe(`GetProfileAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions)`, func() {
			getProfileAttachmentOptions := &securityandcompliancecenterapiv3.GetProfileAttachmentOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:    core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				AttachmentID: &attachmentIDLink,
				AccountID:    &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterAPIService.GetProfileAttachment(getProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
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

			endDate := strfmt.DateTime(time.Now())
			dateRangeModel := &securityandcompliancecenterapiv3.DateRange{
				StartDate: CreateMockDateTime("2025-02-28T05:42:58.000Z"),
				EndDate:   &endDate,
			}

			replaceProfileAttachmentOptions := &securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions{
				InstanceID:           core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:            core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				AttachmentID:         &attachmentIDLink,
				AttachmentParameters: parameterList,
				Description:          core.StringPtr("New Profile Attachment Update"),
				Name:                 core.StringPtr("SDK Updated Test"),
				Notifications:        attachmentNotificationsModel,
				Schedule:             core.StringPtr("daily"),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScopePayloadIntf{multiCloudScopePayloadModel},
				Status:               core.StringPtr("disabled"),
				DataSelectionRange:   dateRangeModel,
				AccountID:            &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterAPIService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`UpgradeAttachment - Upgrade an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpgradeAttachment(upgradeAttachmentOptions *UpgradeAttachmentOptions)`, func() {
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
			upgradeAttachmentOptions := &securityandcompliancecenterapiv3.UpgradeAttachmentOptions{
				InstanceID:           core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:            core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				AttachmentID:         &attachmentIDLink,
				AttachmentParameters: parameterList,
				AccountID:            &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterAPIService.UpgradeAttachment(upgradeAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`CreateScan - Create a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
			createScanOptions := &securityandcompliancecenterapiv3.CreateScanOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AttachmentID: core.StringPtr("4deb572c-9f37-4126-9cc0-d550672533cb"),
				AccountID:    &accountIDForReportLink,
			}

			createScanResponse, response, err := securityAndComplianceCenterAPIService.CreateScan(createScanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createScanResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateControlLibrary - Create a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateControlLibrary(createControlLibraryOptions *CreateControlLibraryOptions)`, func() {
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

			controlPrototypeModel := &securityandcompliancecenterapiv3.ControlPrototype{
				ControlName:           core.StringPtr("security"),
				ControlDescription:    core.StringPtr("This is a description of a control"),
				ControlCategory:       core.StringPtr("test-control"),
				ControlRequirement:    core.BoolPtr(true),
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				Status:                core.StringPtr("disabled"),
			}

			createControlLibraryOptions := &securityandcompliancecenterapiv3.CreateControlLibraryOptions{
				InstanceID:                core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryName:        core.StringPtr("custom control library from SDK"),
				ControlLibraryDescription: core.StringPtr("This is a custom control library made from the SDK test framework"),
				ControlLibraryType:        core.StringPtr("custom"),
				ControlLibraryVersion:     core.StringPtr("0.0.1"),
				Controls:                  []securityandcompliancecenterapiv3.ControlPrototype{*controlPrototypeModel},
				AccountID:                 &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterAPIService.CreateControlLibrary(createControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIDLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIDLink value: %v\n", controlLibraryIDLink)
		})
	})

	Describe(`ListControlLibraries - Get all control libraries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) with pagination`, func() {
			listControlLibrariesOptions := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
				Start:      core.StringPtr("testString"),
			}

			listControlLibrariesOptions.Start = nil
			listControlLibrariesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.ControlLibrary
			for {
				controlLibraryCollection, response, err := securityAndComplianceCenterAPIService.ListControlLibraries(listControlLibrariesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(controlLibraryCollection).ToNot(BeNil())
				allResults = append(allResults, controlLibraryCollection.ControlLibraries...)

				listControlLibrariesOptions.Start, err = controlLibraryCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listControlLibrariesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) using ControlLibrariesPager`, func() {
			listControlLibrariesOptions := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.ControlLibrary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListControlLibraries() returned a total of %d item(s) using ControlLibrariesPager.\n", len(allResults))
		})
	})

	Describe(`ReplaceCustomControlLibrary - Update a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions)`, func() {
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

			controlDocModel := &securityandcompliancecenterapiv3.ControlDoc{
				ControlDocsID:   core.StringPtr("testString"),
				ControlDocsType: core.StringPtr("testString"),
			}

			controlPrototypeModel := &securityandcompliancecenterapiv3.ControlPrototype{
				ControlName:           core.StringPtr("testString"),
				ControlDescription:    core.StringPtr("testString"),
				ControlCategory:       core.StringPtr("testString"),
				ControlRequirement:    core.BoolPtr(true),
				ControlParent:         core.StringPtr(""),
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecificationPrototype{*controlSpecificationPrototypeModel},
				ControlDocs:           controlDocModel,
				Status:                core.StringPtr("enabled"),
			}

			replaceCustomControlLibraryOptions := &securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions{
				InstanceID:                core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID:          &controlLibraryIDLink,
				ControlLibraryName:        core.StringPtr("testControlLibUpdate"),
				ControlLibraryDescription: core.StringPtr("Updates the control library"),
				ControlLibraryType:        core.StringPtr("custom"),
				ControlLibraryVersion:     core.StringPtr("0.0.2"),
				Controls:                  []securityandcompliancecenterapiv3.ControlPrototype{*controlPrototypeModel},
			}

			controlLibrary, response, err := securityAndComplianceCenterAPIService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`GetControlLibrary - Get a control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions)`, func() {
			getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID: &controlLibraryIDLink,
				AccountID:        &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterAPIService.GetControlLibrary(getControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`CreateProfile - Create a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
			profileControlsPrototypeModel := &securityandcompliancecenterapiv3.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr("51ca566e-c559-412b-8d64-f05b57044c32"),
				ControlID:        core.StringPtr("60dae3b5-6104-4b3e-bac7-26cc7b741aca"),
			}

			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr("[\"1.2\",\"1.3\"]"),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			createProfileOptions := &securityandcompliancecenterapiv3.CreateProfileOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileName:        core.StringPtr("Example Profile"),
				ProfileVersion:     core.StringPtr("0.0.1"),
				Controls:           []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
				ProfileDescription: core.StringPtr("This profile is created as an example of the SDK gen"),
				Latest:             core.BoolPtr(true),
				VersionGroupLabel:  core.StringPtr("testString"),
				AccountID:          &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterAPIService.CreateProfile(createProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())

			profileIDLink = *profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIDLink value: %v\n", profileIDLink)
		})
	})

	Describe(`ListProfiles - Get all profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions) with pagination`, func() {
			listProfilesOptions := &securityandcompliancecenterapiv3.ListProfilesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
				Start:      core.StringPtr("testString"),
			}

			listProfilesOptions.Start = nil
			listProfilesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Profile
			for {
				profileCollection, response, err := securityAndComplianceCenterAPIService.ListProfiles(listProfilesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(profileCollection).ToNot(BeNil())
				allResults = append(allResults, profileCollection.Profiles...)

				listProfilesOptions.Start, err = profileCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listProfilesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions) using ProfilesPager`, func() {
			listProfilesOptions := &securityandcompliancecenterapiv3.ListProfilesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:  &accountIDForReportLink,
				Limit:      core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Profile
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProfiles() returned a total of %d item(s) using ProfilesPager.\n", len(allResults))
		})
	})

	Describe(`ReplaceProfile - Update a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions)`, func() {
			controlDocModel := &securityandcompliancecenterapiv3.ControlDoc{
				ControlDocsID:   core.StringPtr("testString"),
				ControlDocsType: core.StringPtr("testString"),
			}

			parameterModel := &securityandcompliancecenterapiv3.Parameter{
				AssessmentType:       core.StringPtr("testString"),
				AssessmentID:         core.StringPtr("testString"),
				ParameterName:        core.StringPtr("location"),
				ParameterDisplayName: core.StringPtr("Location"),
				ParameterType:        core.StringPtr("string"),
				ParameterValue:       "testString",
			}

			assessmentModel := &securityandcompliancecenterapiv3.Assessment{
				AssessmentID:          core.StringPtr("382c2b06-e6b2-43ee-b189-c1c7743b67ee"),
				AssessmentType:        core.StringPtr("ibm-cloud-rule"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentDescription: core.StringPtr("Check whether Cloud Object Storage is accessible only by using private endpoints"),
				ParameterCount:        core.Int64Ptr(int64(1)),
				Parameters:            []securityandcompliancecenterapiv3.Parameter{*parameterModel},
			}

			controlSpecificationModel := &securityandcompliancecenterapiv3.ControlSpecification{
				ID:               core.StringPtr("testString"),
				Responsibility:   core.StringPtr("testString"),
				ComponentID:      core.StringPtr("testString"),
				ComponentName:    core.StringPtr("testString"),
				ComponentType:    core.StringPtr("testString"),
				Environment:      core.StringPtr("testString"),
				Description:      core.StringPtr("testString"),
				AssessmentsCount: core.Int64Ptr(int64(38)),
				Assessments:      []securityandcompliancecenterapiv3.Assessment{*assessmentModel},
			}

			profileControlsModel := &securityandcompliancecenterapiv3.ProfileControls{
				ControlRequirement:    core.BoolPtr(true),
				ControlLibraryID:      core.StringPtr("51ca566e-c559-412b-8d64-f05b57044c32"),
				ControlID:             core.StringPtr("60dae3b5-6104-4b3e-bac7-26cc7b741aca"),
				ControlLibraryVersion: core.StringPtr("testString"),
				ControlName:           core.StringPtr("testString"),
				ControlDescription:    core.StringPtr("testString"),
				ControlSeverity:       core.StringPtr("testString"),
				ControlCategory:       core.StringPtr("testString"),
				ControlParent:         core.StringPtr("testString"),
				ControlDocs:           controlDocModel,
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecification{*controlSpecificationModel},
			}

			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{
				AssessmentType:        core.StringPtr("automated"),
				AssessmentID:          core.StringPtr("rule-e16fcfea-fe21-4d30-a721-423611481fea"),
				ParameterName:         core.StringPtr("tls_version"),
				ParameterDefaultValue: core.StringPtr("[\"1.2\",\"1.3\"]"),
				ParameterDisplayName:  core.StringPtr("IBM Cloud Internet Services TLS version"),
				ParameterType:         core.StringPtr("string_list"),
			}

			replaceProfileOptions := &securityandcompliancecenterapiv3.ReplaceProfileOptions{
				InstanceID:            core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:             &profileIDLink,
				NewProfileType:        core.StringPtr("custom"),
				NewControls:           []securityandcompliancecenterapiv3.ProfileControls{*profileControlsModel},
				NewDefaultParameters:  []securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
				NewID:                 core.StringPtr("testString"),
				NewProfileName:        core.StringPtr("Example Profile Updated"),
				NewInstanceID:         core.StringPtr("testString"),
				NewHierarchyEnabled:   core.BoolPtr(true),
				NewProfileDescription: core.StringPtr("This profile has been updated"),
				NewProfileVersion:     core.StringPtr("0.0.2"),
				NewVersionGroupLabel:  core.StringPtr("testString"),
				NewLatest:             core.BoolPtr(true),
				NewCreatedBy:          core.StringPtr("testString"),
				NewCreatedOn:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				NewUpdatedBy:          core.StringPtr("testString"),
				NewUpdatedOn:          CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				NewControlsCount:      core.Int64Ptr(int64(38)),
				NewAttachmentsCount:   core.Int64Ptr(int64(38)),
				AccountID:             &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterAPIService.ReplaceProfile(replaceProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`GetProfile - Get a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
			getProfileOptions := &securityandcompliancecenterapiv3.GetProfileOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:  &profileIDLink,
				AccountID:  &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterAPIService.GetProfile(getProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileParameters - Update custom profile parameters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions)`, func() {
			defaultParametersModel := &securityandcompliancecenterapiv3.DefaultParameters{
				AssessmentType:        core.StringPtr("testString"),
				AssessmentID:          core.StringPtr("testString"),
				ParameterName:         core.StringPtr("testString"),
				ParameterDefaultValue: core.StringPtr("testString"),
				ParameterDisplayName:  core.StringPtr("testString"),
				ParameterType:         core.StringPtr("testString"),
			}

			replaceProfileParametersOptions := &securityandcompliancecenterapiv3.ReplaceProfileParametersOptions{
				InstanceID:        core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:         &profileIDLink,
				DefaultParameters: []securityandcompliancecenterapiv3.DefaultParameters{*defaultParametersModel},
				ID:                core.StringPtr("testString"),
				AccountID:         &accountIDForReportLink,
			}

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterAPIService.ReplaceProfileParameters(replaceProfileParametersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
	})

	Describe(`ListProfileParameters - List profile parameters for a given profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileParameters(listProfileParametersOptions *ListProfileParametersOptions)`, func() {
			listProfileParametersOptions := &securityandcompliancecenterapiv3.ListProfileParametersOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:  &profileIDLink,
			}

			profileDefaultParametersResponse, response, err := securityAndComplianceCenterAPIService.ListProfileParameters(listProfileParametersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
	})

	Describe(`CompareProfiles - Compare profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompareProfiles(compareProfilesOptions *CompareProfilesOptions)`, func() {
			compareProfilesOptions := &securityandcompliancecenterapiv3.CompareProfilesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:  core.StringPtr("2f598907-970d-4d52-9071-5cc95912f55e"),
				AccountID:  &accountIDForReportLink,
			}

			comparePredefinedProfilesResponse, response, err := securityAndComplianceCenterAPIService.CompareProfiles(compareProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(comparePredefinedProfilesResponse).ToNot(BeNil())
		})
	})

	Describe(`ListProfileAttachments - Get all attachments linked to a specific profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileAttachments(listProfileAttachmentsOptions *ListProfileAttachmentsOptions)`, func() {
			listProfileAttachmentsOptions := &securityandcompliancecenterapiv3.ListProfileAttachmentsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:  core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				AccountID:  &accountIDForReportLink,
			}

			profileAttachmentCollection, response, err := securityAndComplianceCenterAPIService.ListProfileAttachments(listProfileAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachmentCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateScope - Create a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {
			scopePropertyModel := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("ff88f007f9ff4622aac4fbc0eda36255"),
			}

			createScopeOptions := &securityandcompliancecenterapiv3.CreateScopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Name:        core.StringPtr("ibm scope"),
				Description: core.StringPtr("The scope that is defined for IBM resources."),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterapiv3.ScopePropertyIntf{scopePropertyModel},
			}

			scope, response, err := securityAndComplianceCenterAPIService.CreateScope(createScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scope).ToNot(BeNil())

			scopeIDLink = *scope.ID
			fmt.Fprintf(GinkgoWriter, "Saved scopeIDLink value: %v\n", scopeIDLink)
		})
	})

	Describe(`ListScopes - Get all scopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions) with pagination`, func() {
			listScopesOptions := &securityandcompliancecenterapiv3.ListScopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Start:       core.StringPtr("testString"),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			listScopesOptions.Start = nil
			listScopesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Scope
			for {
				scopeCollection, response, err := securityAndComplianceCenterAPIService.ListScopes(listScopesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(scopeCollection).ToNot(BeNil())
				allResults = append(allResults, scopeCollection.Scopes...)

				listScopesOptions.Start, err = scopeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listScopesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListScopes(listScopesOptions *ListScopesOptions) using ScopesPager`, func() {
			listScopesOptions := &securityandcompliancecenterapiv3.ListScopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewScopesPager(listScopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Scope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewScopesPager(listScopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListScopes() returned a total of %d item(s) using ScopesPager.\n", len(allResults))
		})
	})

	Describe(`UpdateScope - Update a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateScope(updateScopeOptions *UpdateScopeOptions)`, func() {
			updateScopeOptions := &securityandcompliancecenterapiv3.UpdateScopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				Name:        core.StringPtr("updated name of scope"),
				Description: core.StringPtr("updated scope description"),
			}

			scope, response, err := securityAndComplianceCenterAPIService.UpdateScope(updateScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
	})

	Describe(`GetScope - Get a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScope(getScopeOptions *GetScopeOptions)`, func() {
			getScopeOptions := &securityandcompliancecenterapiv3.GetScopeOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
			}

			scope, response, err := securityAndComplianceCenterAPIService.GetScope(getScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scope).ToNot(BeNil())
		})
	})

	Describe(`CreateSubscope - Create a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSubscope(createSubscopeOptions *CreateSubscopeOptions)`, func() {
			scopePropertyModel := &securityandcompliancecenterapiv3.ScopePropertyScopeAny{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("1f689f08ec9b47b885c2659c17029581"),
			}

			scopePrototypeModel := &securityandcompliancecenterapiv3.ScopePrototype{
				Name:        core.StringPtr("ibm subscope update"),
				Description: core.StringPtr("The subscope that is defined for IBM resources."),
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterapiv3.ScopePropertyIntf{scopePropertyModel},
			}

			createSubscopeOptions := &securityandcompliancecenterapiv3.CreateSubscopeOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
				Subscopes:  []securityandcompliancecenterapiv3.ScopePrototype{*scopePrototypeModel},
			}

			subScopeResponse, response, err := securityAndComplianceCenterAPIService.CreateSubscope(createSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subScopeResponse).ToNot(BeNil())

			subScopeIDLink = *subScopeResponse.Subscopes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved subScopeIDLink value: %v\n", subScopeIDLink)
		})
	})

	Describe(`ListSubscopes - Get all subscopes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSubscopes(listSubscopesOptions *ListSubscopesOptions) with pagination`, func() {
			listSubscopesOptions := &securityandcompliancecenterapiv3.ListSubscopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				Limit:       core.Int64Ptr(int64(10)),
				Start:       core.StringPtr("testString"),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			listSubscopesOptions.Start = nil
			listSubscopesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.SubScope
			for {
				subScopeCollection, response, err := securityAndComplianceCenterAPIService.ListSubscopes(listSubscopesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(subScopeCollection).ToNot(BeNil())
				allResults = append(allResults, subScopeCollection.Subscopes...)

				listSubscopesOptions.Start, err = subScopeCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listSubscopesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSubscopes(listSubscopesOptions *ListSubscopesOptions) using SubscopesPager`, func() {
			listSubscopesOptions := &securityandcompliancecenterapiv3.ListSubscopesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				Limit:       core.Int64Ptr(int64(10)),
				Name:        core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Environment: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewSubscopesPager(listSubscopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.SubScope
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewSubscopesPager(listSubscopesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSubscopes() returned a total of %d item(s) using SubscopesPager.\n", len(allResults))
		})
	})

	Describe(`GetSubscope - Get a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubscope(getSubscopeOptions *GetSubscopeOptions)`, func() {
			getSubscopeOptions := &securityandcompliancecenterapiv3.GetSubscopeOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
				SubscopeID: &subScopeIDLink,
			}

			subScope, response, err := securityAndComplianceCenterAPIService.GetSubscope(getSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
	})

	Describe(`UpdateSubscope - Update a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSubscope(updateSubscopeOptions *UpdateSubscopeOptions)`, func() {
			updateSubscopeOptions := &securityandcompliancecenterapiv3.UpdateSubscopeOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:     &scopeIDLink,
				SubscopeID:  &subScopeIDLink,
				Name:        core.StringPtr("updated name of scope"),
				Description: core.StringPtr("updated scope description"),
			}

			subScope, response, err := securityAndComplianceCenterAPIService.UpdateSubscope(updateSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subScope).ToNot(BeNil())
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
			accountModel := &securityandcompliancecenterapiv3.Account{
				ID:   core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				Name: core.StringPtr("NIST"),
				Type: core.StringPtr("account_type"),
			}

			tagsModel := &securityandcompliancecenterapiv3.Tags{
				User:    []string{"testString"},
				Access:  []string{"testString"},
				Service: []string{"testString"},
			}

			resourceModel := &securityandcompliancecenterapiv3.Resource{
				ReportID:       core.StringPtr("30b434b3-cb08-4845-af10-7a8fc682b6a8"),
				HomeAccountID:  core.StringPtr("2411ffdc16844b07b42521c3443f456d"),
				ID:             core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::"),
				ResourceName:   core.StringPtr("jeff's key"),
				Account:        accountModel,
				ComponentID:    core.StringPtr("cloud-object_storage"),
				ComponentName:  core.StringPtr("cloud-object_storage"),
				Environment:    core.StringPtr("ibm cloud"),
				Tags:           tagsModel,
				Status:         core.StringPtr("compliant"),
				TotalCount:     core.Int64Ptr(int64(140)),
				PassCount:      core.Int64Ptr(int64(123)),
				FailureCount:   core.Int64Ptr(int64(12)),
				ErrorCount:     core.Int64Ptr(int64(5)),
				SkippedCount:   core.Int64Ptr(int64(7)),
				CompletedCount: core.Int64Ptr(int64(135)),
				ServiceName:    core.StringPtr("pm-20"),
				InstanceCRN:    core.StringPtr("testString"),
			}

			credentialModel := &securityandcompliancecenterapiv3.Credential{
				SecretCRN: core.StringPtr("testString"),
				Resources: []securityandcompliancecenterapiv3.Resource{*resourceModel},
			}

			createTargetOptions := &securityandcompliancecenterapiv3.CreateTargetOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				AccountID:        core.StringPtr("62ecf99b240144dea9125666249edfcb"),
				TrustedProfileID: core.StringPtr("Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54"),
				Name:             core.StringPtr("Target for IBM account"),
				Credentials:      []securityandcompliancecenterapiv3.Credential{*credentialModel},
			}

			target, response, err := securityAndComplianceCenterAPIService.CreateTarget(createTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)
		})
	})

	Describe(`ListTargets - Get a list of targets with pagination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
			listTargetsOptions := &securityandcompliancecenterapiv3.ListTargetsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
			}

			targetCollection, response, err := securityAndComplianceCenterAPIService.ListTargets(listTargetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})
	})

	Describe(`GetTarget - Get a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
			getTargetOptions := &securityandcompliancecenterapiv3.GetTargetOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				TargetID:   &targetIDLink,
			}

			target, response, err := securityAndComplianceCenterAPIService.GetTarget(getTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`ReplaceTarget - replace a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {
			accountModel := &securityandcompliancecenterapiv3.Account{
				ID:   core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				Name: core.StringPtr("NIST"),
				Type: core.StringPtr("account_type"),
			}

			tagsModel := &securityandcompliancecenterapiv3.Tags{
				User:    []string{"testString"},
				Access:  []string{"testString"},
				Service: []string{"testString"},
			}

			resourceModel := &securityandcompliancecenterapiv3.Resource{
				ReportID:       core.StringPtr("30b434b3-cb08-4845-af10-7a8fc682b6a8"),
				HomeAccountID:  core.StringPtr("2411ffdc16844b07b42521c3443f456d"),
				ID:             core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::"),
				ResourceName:   core.StringPtr("jeff's key"),
				Account:        accountModel,
				ComponentID:    core.StringPtr("cloud-object_storage"),
				ComponentName:  core.StringPtr("cloud-object_storage"),
				Environment:    core.StringPtr("ibm cloud"),
				Tags:           tagsModel,
				Status:         core.StringPtr("compliant"),
				TotalCount:     core.Int64Ptr(int64(140)),
				PassCount:      core.Int64Ptr(int64(123)),
				FailureCount:   core.Int64Ptr(int64(12)),
				ErrorCount:     core.Int64Ptr(int64(5)),
				SkippedCount:   core.Int64Ptr(int64(7)),
				CompletedCount: core.Int64Ptr(int64(135)),
				ServiceName:    core.StringPtr("pm-20"),
				InstanceCRN:    core.StringPtr("testString"),
			}

			credentialModel := &securityandcompliancecenterapiv3.Credential{
				SecretCRN: core.StringPtr("testString"),
				Resources: []securityandcompliancecenterapiv3.Resource{*resourceModel},
			}

			replaceTargetOptions := &securityandcompliancecenterapiv3.ReplaceTargetOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				TargetID:         &targetIDLink,
				AccountID:        core.StringPtr("62ecf99b240144dea9125666249edfcb"),
				TrustedProfileID: core.StringPtr("Profile-cb2c1829-9a8d-4218-b9cd-9f83fc814e54"),
				Name:             core.StringPtr("Updated target for IBM account"),
				Credentials:      []securityandcompliancecenterapiv3.Credential{*credentialModel},
			}

			target, response, err := securityAndComplianceCenterAPIService.ReplaceTarget(replaceTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})
	})

	Describe(`CreateProviderTypeInstance - Create a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProviderTypeInstance(createProviderTypeInstanceOptions *CreateProviderTypeInstanceOptions)`, func() {
			createProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
				Name:           core.StringPtr("workload-protection-instance-1"),
				Attributes:     map[string]interface{}{"anyKey": "anyValue"},
			}

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstance).ToNot(BeNil())

			providerTypeInstanceIDLink = *providerTypeInstance.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIDLink value: %v\n", providerTypeInstanceIDLink)
		})
	})

	Describe(`ListProviderTypeInstances - List instances of a specific provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypeInstances(listProviderTypeInstancesOptions *ListProviderTypeInstancesOptions)`, func() {
			listProviderTypeInstancesOptions := &securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
			}

			providerTypeInstanceCollection, response, err := securityAndComplianceCenterAPIService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceCollection).ToNot(BeNil())
		})
	})

	Describe(`GetProviderTypeInstance - Get a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeInstance(getProviderTypeInstanceOptions *GetProviderTypeInstanceOptions)`, func() {
			getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
			}

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
	})

	Describe(`UpdateProviderTypeInstance - Update a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProviderTypeInstance(updateProviderTypeInstanceOptions *UpdateProviderTypeInstanceOptions)`, func() {
			updateProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
				Name:                   core.StringPtr("workload-protection-instance-1"),
				Attributes:             map[string]interface{}{"anyKey": "anyValue"},
			}

			providerTypeInstance, response, err := securityAndComplianceCenterAPIService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstance).ToNot(BeNil())
		})
	})

	Describe(`ListProviderTypes - List provider types`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypes(listProviderTypesOptions *ListProviderTypesOptions)`, func() {
			listProviderTypesOptions := &securityandcompliancecenterapiv3.ListProviderTypesOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
			}

			providerTypeCollection, response, err := securityAndComplianceCenterAPIService.ListProviderTypes(listProviderTypesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeCollection).ToNot(BeNil())
		})
	})

	Describe(`GetProviderTypeByID - Get a provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeByID(getProviderTypeByIDOptions *GetProviderTypeByIDOptions)`, func() {
			getProviderTypeByIDOptions := &securityandcompliancecenterapiv3.GetProviderTypeByIDOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID: core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
			}

			providerType, response, err := securityAndComplianceCenterAPIService.GetProviderTypeByID(getProviderTypeByIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerType).ToNot(BeNil())
		})
	})

	Describe(`GetLatestReports - List latest reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
			getLatestReportsOptions := &securityandcompliancecenterapiv3.GetLatestReportsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Sort:       core.StringPtr("profile_name"),
			}

			reportLatest, response, err := securityAndComplianceCenterAPIService.GetLatestReports(getLatestReportsOptions)
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
	})

	Describe(`ListReports - List reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) with pagination`, func() {
			listReportsOptions := &securityandcompliancecenterapiv3.ListReportsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportAttachmentID: &attachmentIDForReportLink,
				GroupID:            &groupIDForReportLink,
				ReportProfileID:    &profileIDForReportLink,
				Type:               &typeForReportLink,
				Start:              core.StringPtr("testString"),
				Limit:              core.Int64Ptr(int64(10)),
				Sort:               core.StringPtr("profile_name"),
			}

			listReportsOptions.Start = nil
			listReportsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Report
			for {
				reportCollection, response, err := securityAndComplianceCenterAPIService.ListReports(listReportsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reportCollection).ToNot(BeNil())
				allResults = append(allResults, reportCollection.Reports...)

				listReportsOptions.Start, err = reportCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listReportsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) using ReportsPager`, func() {
			listReportsOptions := &securityandcompliancecenterapiv3.ListReportsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportAttachmentID: &attachmentIDForReportLink,
				GroupID:            &groupIDForReportLink,
				ReportProfileID:    &profileIDForReportLink,
				Type:               &typeForReportLink,
				Limit:              core.Int64Ptr(int64(10)),
				Sort:               core.StringPtr("profile_name"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Report
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewReportsPager(listReportsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReports() returned a total of %d item(s) using ReportsPager.\n", len(allResults))
		})
	})

	Describe(`GetReport - Get a report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReport(getReportOptions *GetReportOptions)`, func() {
			getReportOptions := &securityandcompliancecenterapiv3.GetReportOptions{
				ReportID:   &reportIDForReportLink,
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    core.StringPtr("testString"),
				SubscopeID: core.StringPtr("testString"),
			}

			report, response, err := securityAndComplianceCenterAPIService.GetReport(getReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
	})

	Describe(`GetReportSummary - Get a report summary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions)`, func() {
			getReportSummaryOptions := &securityandcompliancecenterapiv3.GetReportSummaryOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
			}

			reportSummary, response, err := securityAndComplianceCenterAPIService.GetReportSummary(getReportSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetReportDownloadFile - Get report evaluation details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportDownloadFile(getReportDownloadFileOptions *GetReportDownloadFileOptions)`, func() {
			getReportDownloadFileOptions := &securityandcompliancecenterapiv3.GetReportDownloadFileOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:       &reportIDForReportLink,
				Accept:         core.StringPtr("application/csv"),
				ExcludeSummary: core.BoolPtr(true),
			}

			result, response, err := securityAndComplianceCenterAPIService.GetReportDownloadFile(getReportDownloadFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetReportControls - Get report controls`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportControls(getReportControlsOptions *GetReportControlsOptions)`, func() {
			getReportControlsOptions := &securityandcompliancecenterapiv3.GetReportControlsOptions{
				InstanceID:         core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:           &reportIDForReportLink,
				ControlID:          core.StringPtr("testString"),
				ControlName:        core.StringPtr("testString"),
				ControlDescription: core.StringPtr("testString"),
				ControlCategory:    core.StringPtr("testString"),
				Status:             core.StringPtr("compliant"),
				Sort:               core.StringPtr("control_name"),
				ScopeID:            core.StringPtr("testString"),
				SubscopeID:         core.StringPtr("testString"),
			}

			reportControls, response, err := securityAndComplianceCenterAPIService.GetReportControls(getReportControlsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportControls).ToNot(BeNil())
		})
	})

	Describe(`GetReportRule - Get a report rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportRule(getReportRuleOptions *GetReportRuleOptions)`, func() {
			getReportRuleOptions := &securityandcompliancecenterapiv3.GetReportRuleOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				RuleID:     core.StringPtr("rule-61fa114a-2bb9-43fd-8068-b873b48bdf79"),
			}

			ruleInfo, response, err := securityAndComplianceCenterAPIService.GetReportRule(getReportRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleInfo).ToNot(BeNil())
		})
	})

	Describe(`ListReportEvaluations - List report evaluations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) with pagination`, func() {
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
				Start:            core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(10)),
				Sort:             core.StringPtr("assessment_id"),
				ScopeID:          core.StringPtr("testString"),
				SubscopeID:       core.StringPtr("testString"),
			}

			listReportEvaluationsOptions.Start = nil
			listReportEvaluationsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Evaluation
			for {
				evaluationPage, response, err := securityAndComplianceCenterAPIService.ListReportEvaluations(listReportEvaluationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(evaluationPage).ToNot(BeNil())
				allResults = append(allResults, evaluationPage.Evaluations...)

				listReportEvaluationsOptions.Start, err = evaluationPage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportEvaluationsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) using ReportEvaluationsPager`, func() {
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

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Evaluation
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportEvaluations() returned a total of %d item(s) using ReportEvaluationsPager.\n", len(allResults))
		})
	})

	Describe(`ListReportResources - List report resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) with pagination`, func() {
			listReportResourcesOptions := &securityandcompliancecenterapiv3.ListReportResourcesOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:     &reportIDForReportLink,
				ID:           core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				AccountID:    &accountIDForReportLink,
				ComponentID:  core.StringPtr("testString"),
				Status:       core.StringPtr("compliant"),
				Sort:         core.StringPtr("account_id"),
				Start:        core.StringPtr("testString"),
				Limit:        core.Int64Ptr(int64(10)),
				ScopeID:      core.StringPtr("testString"),
				SubscopeID:   core.StringPtr("testString"),
			}

			listReportResourcesOptions.Start = nil
			listReportResourcesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Resource
			for {
				resourcePage, response, err := securityAndComplianceCenterAPIService.ListReportResources(listReportResourcesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(resourcePage).ToNot(BeNil())
				allResults = append(allResults, resourcePage.Resources...)

				listReportResourcesOptions.Start, err = resourcePage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportResourcesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) using ReportResourcesPager`, func() {
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

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Resource
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportResources() returned a total of %d item(s) using ReportResourcesPager.\n", len(allResults))
		})
	})

	Describe(`GetReportTags - List report tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
			getReportTagsOptions := &securityandcompliancecenterapiv3.GetReportTagsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
			}

			reportTags, response, err := securityAndComplianceCenterAPIService.GetReportTags(getReportTagsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportTags).ToNot(BeNil())
		})
	})

	Describe(`GetReportViolationsDrift - Get report violations drift`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions)`, func() {
			getReportViolationsDriftOptions := &securityandcompliancecenterapiv3.GetReportViolationsDriftOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:         &reportIDForReportLink,
				ScanTimeDuration: core.Int64Ptr(int64(0)),
				ScopeID:          core.StringPtr("testString"),
				SubscopeID:       core.StringPtr("testString"),
			}

			reportViolationsDrift, response, err := securityAndComplianceCenterAPIService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportViolationsDrift).ToNot(BeNil())
		})
	})

	Describe(`ListScanReports - List scan reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScanReports(listScanReportsOptions *ListScanReportsOptions)`, func() {
			listScanReportsOptions := &securityandcompliancecenterapiv3.ListScanReportsOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				ScopeID:    core.StringPtr("testString"),
				SubscopeID: core.StringPtr("testString"),
				Sort:       core.StringPtr("status"),
			}

			scanReportCollection, response, err := securityAndComplianceCenterAPIService.ListScanReports(listScanReportsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReportCollection).ToNot(BeNil())

			scanIDForScanReportLink = *scanReportCollection.ScanReports[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved scanIDForScanReportLink value: %v\n", scanIDForScanReportLink)
		})
	})

	Describe(`CreateScanReport - Create a scan report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScanReport(createScanReportOptions *CreateScanReportOptions)`, func() {
			createScanReportOptions := &securityandcompliancecenterapiv3.CreateScanReportOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				Format:     core.StringPtr("csv"),
				ScopeID:    core.StringPtr("132009ff-b982-412e-a110-ad8797e10f84"),
				SubscopeID: core.StringPtr("c7ddcbcc-6a43-4ab3-b6a7-b2d8f65cd54a"),
			}

			createScanReport, response, err := securityAndComplianceCenterAPIService.CreateScanReport(createScanReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createScanReport).ToNot(BeNil())
		})
	})

	Describe(`GetScanReport - Get a scan report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScanReport(getScanReportOptions *GetScanReportOptions)`, func() {
			getScanReportOptions := &securityandcompliancecenterapiv3.GetScanReportOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				JobID:      &scanIDForScanReportLink,
			}

			scanReport, response, err := securityAndComplianceCenterAPIService.GetScanReport(getScanReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanReport).ToNot(BeNil())
		})
	})

	Describe(`GetScanReportDownloadFile - Get a scan report details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetScanReportDownloadFile(getScanReportDownloadFileOptions *GetScanReportDownloadFileOptions)`, func() {
			getScanReportDownloadFileOptions := &securityandcompliancecenterapiv3.GetScanReportDownloadFileOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ReportID:   &reportIDForReportLink,
				JobID:      &scanIDForScanReportLink,
				Accept:     core.StringPtr("application/csv"),
			}

			result, response, err := securityAndComplianceCenterAPIService.GetScanReportDownloadFile(getScanReportDownloadFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`ListRules - Get all rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) with pagination`, func() {
			listRulesOptions := &securityandcompliancecenterapiv3.ListRulesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Start:       core.StringPtr("testString"),
				Type:        core.StringPtr("system_defined"),
				Search:      core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				Sort:        core.StringPtr("updated_on"),
			}

			listRulesOptions.Start = nil
			listRulesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Rule
			for {
				ruleCollection, response, err := securityAndComplianceCenterAPIService.ListRules(listRulesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(ruleCollection).ToNot(BeNil())
				allResults = append(allResults, ruleCollection.Rules...)

				listRulesOptions.Start, err = ruleCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listRulesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) using RulesPager`, func() {
			listRulesOptions := &securityandcompliancecenterapiv3.ListRulesOptions{
				InstanceID:  core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Limit:       core.Int64Ptr(int64(10)),
				Type:        core.StringPtr("system_defined"),
				Search:      core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				Sort:        core.StringPtr("updated_on"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterAPIService.NewRulesPager(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.Rule
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterAPIService.NewRulesPager(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListRules() returned a total of %d item(s) using RulesPager.\n", len(allResults))
		})
	})

	Describe(`CreateRule - Create a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
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
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
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

			createRuleOptions := &securityandcompliancecenterapiv3.CreateRuleOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				Description:    core.StringPtr("Example rule"),
				Target:         ruleTargetPrototypeModel,
				RequiredConfig: requiredConfigModel,
				Version:        core.StringPtr("1.0.0"),
				Import:         importModel,
				Labels:         []string{},
			}

			rule, response, err := securityAndComplianceCenterAPIService.CreateRule(createRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIDLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIDLink value: %v\n", ruleIDLink)
		})
	})

	Describe(`GetRule - Get a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &securityandcompliancecenterapiv3.GetRuleOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				RuleID:     &ruleIDLink,
			}

			rule, response, err := securityAndComplianceCenterAPIService.GetRule(getRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
	})

	Describe(`ReplaceRule - Update a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
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
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
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

			replaceRuleOptions := &securityandcompliancecenterapiv3.ReplaceRuleOptions{
				InstanceID:     core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				RuleID:         &ruleIDLink,
				IfMatch:        &eTagLink,
				Description:    core.StringPtr("Example rule"),
				Target:         ruleTargetPrototypeModel,
				RequiredConfig: requiredConfigModel,
				Version:        core.StringPtr("1.0.1"),
				Import:         importModel,
				Labels:         []string{},
			}

			rule, response, err := securityAndComplianceCenterAPIService.ReplaceRule(replaceRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`ListServices - List services`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServices(listServicesOptions *ListServicesOptions)`, func() {
			listServicesOptions := &securityandcompliancecenterapiv3.ListServicesOptions{}

			serviceCollection, response, err := securityAndComplianceCenterAPIService.ListServices(listServicesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceCollection).ToNot(BeNil())
		})
	})

	Describe(`GetService - Get a service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetService(getServiceOptions *GetServiceOptions)`, func() {
			getServiceOptions := &securityandcompliancecenterapiv3.GetServiceOptions{
				ServicesName: core.StringPtr("cloud-object-storage"),
			}

			service, response, err := securityAndComplianceCenterAPIService.GetService(getServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(service).ToNot(BeNil())
		})
	})

	Describe(`DeleteProfileAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions)`, func() {
			deleteProfileAttachmentOptions := &securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions{
				InstanceID:   core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:    core.StringPtr("9c265b4a-4cdf-47f1-acd3-17b5808f7f3f"),
				AttachmentID: &attachmentIDLink,
				AccountID:    &accountIDForReportLink,
			}

			profileAttachment, response, err := securityAndComplianceCenterAPIService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileAttachment).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomControlLibrary - Delete a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions)`, func() {
			deleteCustomControlLibraryOptions := &securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions{
				InstanceID:       core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ControlLibraryID: &controlLibraryIDLink,
				AccountID:        &accountIDForReportLink,
			}

			controlLibrary, response, err := securityAndComplianceCenterAPIService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomProfile - Delete a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
			deleteCustomProfileOptions := &securityandcompliancecenterapiv3.DeleteCustomProfileOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProfileID:  &profileIDLink,
				AccountID:  &accountIDForReportLink,
			}

			profile, response, err := securityAndComplianceCenterAPIService.DeleteCustomProfile(deleteCustomProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`DeleteSubscope - Delete a subscope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSubscope(deleteSubscopeOptions *DeleteSubscopeOptions)`, func() {
			deleteSubscopeOptions := &securityandcompliancecenterapiv3.DeleteSubscopeOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
				SubscopeID: &subScopeIDLink,
			}

			response, err := securityAndComplianceCenterAPIService.DeleteSubscope(deleteSubscopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteScope - Delete a scope`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteScope(deleteScopeOptions *DeleteScopeOptions)`, func() {
			deleteScopeOptions := &securityandcompliancecenterapiv3.DeleteScopeOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ScopeID:    &scopeIDLink,
			}

			response, err := securityAndComplianceCenterAPIService.DeleteScope(deleteScopeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTarget - Delete a target by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
			deleteTargetOptions := &securityandcompliancecenterapiv3.DeleteTargetOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				TargetID:   &targetIDLink,
			}

			response, err := securityAndComplianceCenterAPIService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteProviderTypeInstance - Delete a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions *DeleteProviderTypeInstanceOptions)`, func() {
			deleteProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.DeleteProviderTypeInstanceOptions{
				InstanceID:             core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				ProviderTypeID:         core.StringPtr("3e25966275dccfa2c3a34786919c5af7"),
				ProviderTypeInstanceID: &providerTypeInstanceIDLink,
			}

			response, err := securityAndComplianceCenterAPIService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteRule - Delete a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &securityandcompliancecenterapiv3.DeleteRuleOptions{
				InstanceID: core.StringPtr("acd7032c-15a3-484f-bf5b-67d41534d940"),
				RuleID:     &ruleIDLink,
			}

			response, err := securityAndComplianceCenterAPIService.DeleteRule(deleteRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//

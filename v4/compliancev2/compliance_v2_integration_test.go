//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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

package compliancev2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/compliancev2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the compliancev2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */
var _ = Describe(`ComplianceV2 Integration Tests`, func() {
	const externalConfigFile = "compliance_v2.env"

	var (
		err                 error
		complianceService   *compliancev2.ComplianceV2
		config              map[string]string
		serviceURL          string
		authenticator       core.IamAuthenticator
		accountID           string
		authUrl             string
		apiKey              string
		instanceID          string
		serviceName         string
		controlLibraryIdNew string
		profileIdNew        string
		attachmentIdNew     string
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
			config, err = core.GetServiceProperties(compliancev2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			authUrl = config["IAM_APIKEY_URL"]
			if authUrl == "" {
				Skip("Unable to load auth service URL configuration property, skipping tests")
			}

			apiKey = config["IAM"]
			if apiKey == "" {
				Skip("Unable to load IAM configuration property, skipping tests")
			}

			serviceName = config["serviceName"]
			if serviceName == "" {
				Skip("Unable to load serviceName configuration property, skipping tests")
			}

			accountID = config["accountID"]

			instanceID = config["instanceID"]
			if instanceID == "" {
				Skip("Unable to load instanceID configuration property, skipping tests")
			}

			authenticator = core.IamAuthenticator{
				ApiKey: apiKey,
				URL:    authUrl,
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
			complianceServiceOptions := &compliancev2.ComplianceV2Options{
				URL:           serviceURL,
				Authenticator: &authenticator,
				ServiceName:   serviceName,
			}

			complianceService, err = compliancev2.NewComplianceV2UsingExternalConfig(complianceServiceOptions)
			Expect(err).To(BeNil())
			Expect(complianceService).ToNot(BeNil())
			Expect(complianceService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			complianceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListControlLibraries - Get control libraries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions)`, func() {
			listControlLibrariesOptions := &compliancev2.ListControlLibrariesOptions{
				XCorrelationID:     core.StringPtr("SDK-automation-ListControlLibraries"),
				XRequestID:         core.StringPtr("SDK-automation-ListControlLibraries"),
				Limit:              core.Int64Ptr(int64(50)),
				ControlLibraryType: core.StringPtr("custom"),
			}

			controlLibraryCollection, response, err := complianceService.ListControlLibraries(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateCustomControlLibrary - Create a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions)`, func() {

			controlDocsModel := &compliancev2.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &compliancev2.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlParent:         core.StringPtr(""),
				ControlRequirement:    core.BoolPtr(false),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []compliancev2.ControlSpecifications{},
				ControlDocs:           controlDocsModel,
				Status:                core.StringPtr("enabled"),
			}

			createCustomControlLibraryOptions := &compliancev2.CreateCustomControlLibraryOptions{
				ControlLibraryName:        core.StringPtr("SDK-automation-CL"),
				ControlLibraryDescription: core.StringPtr("SDK-automation-CL"),
				ControlLibraryType:        core.StringPtr("custom"),
				Controls:                  []compliancev2.ControlsInControlLib{*controlsInControlLibModel},
				ControlLibraryVersion:     core.StringPtr("1.0.0"),
				Latest:                    core.BoolPtr(true),
				ControlsCount:             core.Int64Ptr(int64(38)),
				XCorrelationID:            core.StringPtr("SDK-automation-CreateCustomControlLibrary"),
				XRequestID:                core.StringPtr("SDK-automation-CreateCustomControlLibrary"),
			}

			controlLibrary, response, err := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptions)

			controlLibraryIdNew = *controlLibrary.ID

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
			getControlLibraryOptions := &compliancev2.GetControlLibraryOptions{
				ControlLibrariesID: core.StringPtr(controlLibraryIdNew),
				XCorrelationID:     core.StringPtr("SDK-automation-GetControlLibrary"),
				XRequestID:         core.StringPtr("SDK-automation-GetControlLibrary"),
			}

			controlLibrary, response, err := complianceService.GetControlLibrary(getControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`ReplaceCustomControlLibrary - Update a control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions)`, func() {

			controlDocsModel := &compliancev2.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &compliancev2.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlParent:         core.StringPtr(""),
				ControlRequirement:    core.BoolPtr(false),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []compliancev2.ControlSpecifications{},
				ControlDocs:           controlDocsModel,
				Status:                core.StringPtr("enabled"),
			}

			replaceCustomControlLibraryOptions := &compliancev2.ReplaceCustomControlLibraryOptions{
				ControlLibrariesID:        core.StringPtr(controlLibraryIdNew),
				ID:                        core.StringPtr(controlLibraryIdNew),
				AccountID:                 core.StringPtr(accountID),
				ControlLibraryName:        core.StringPtr("SDK-automation-CL-Edit"),
				ControlLibraryDescription: core.StringPtr("SDK-automation-CL-Edit"),
				ControlLibraryType:        core.StringPtr("custom"),
				VersionGroupLabel:         core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d"),
				ControlLibraryVersion:     core.StringPtr("1.1.0"),
				CreatedOn:                 CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				CreatedBy:                 core.StringPtr("SDK-automation"),
				UpdatedOn:                 CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				UpdatedBy:                 core.StringPtr("SDK-automation"),
				Latest:                    core.BoolPtr(true),
				HierarchyEnabled:          core.BoolPtr(true),
				ControlsCount:             core.Int64Ptr(int64(38)),
				ControlParentsCount:       core.Int64Ptr(int64(38)),
				Controls:                  []compliancev2.ControlsInControlLib{*controlsInControlLibModel},
				XCorrelationID:            core.StringPtr("SDK-automation-ReplaceCustomControlLibrary"),
				XRequestID:                core.StringPtr("SDK-automation-ReplaceCustomControlLibrary"),
			}

			controlLibrary, response, err := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())
		})
	})

	Describe(`ListProfiles - Get all profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {
			listProfilesOptions := &compliancev2.ListProfilesOptions{
				XCorrelationID: core.StringPtr("SDK-automation-ListProfiles"),
				XRequestID:     core.StringPtr("SDK-automation-ListProfiles"),
				Limit:          core.Int64Ptr(int64(50)),
				ProfileType:    core.StringPtr("custom"),
			}

			profileCollection, response, err := complianceService.ListProfiles(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateProfile - Create a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
			profileControlsPrototypeModel := &compliancev2.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr(controlLibraryIdNew),
				ControlID:        core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
			}

			defaultParametersPrototypeModel := &compliancev2.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:         core.StringPtr("session_invalidation_in_seconds"),
				ParameterDefaultValue: core.StringPtr("120"),
				ParameterDisplayName:  core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:         core.StringPtr("numeric"),
			}

			createProfileOptions := &compliancev2.CreateProfileOptions{
				ProfileName:        core.StringPtr("SDK-automation-CustomProfile"),
				ProfileDescription: core.StringPtr("SDK-automation-CustomProfile"),
				ProfileType:        core.StringPtr("custom"),
				Controls:           []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				XCorrelationID:     core.StringPtr("SDK-automation-CreateProfile"),
				XRequestID:         core.StringPtr("SDK-automation-CreateProfile"),
			}

			profile, response, err := complianceService.CreateProfile(createProfileOptions)

			profileIdNew = *profile.ID

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
			getProfileOptions := &compliancev2.GetProfileOptions{
				ProfilesID:     core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-GetProfile"),
				XRequestID:     core.StringPtr("SDK-automation-GetProfile"),
			}

			profile, response, err := complianceService.GetProfile(getProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfile - Update a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions)`, func() {
			profileControlsPrototypeModel := &compliancev2.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr(controlLibraryIdNew),
				ControlID:        core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
			}

			defaultParametersPrototypeModel := &compliancev2.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:         core.StringPtr("session_invalidation_in_seconds"),
				ParameterDefaultValue: core.StringPtr("120"),
				ParameterDisplayName:  core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:         core.StringPtr("numeric"),
			}

			replaceProfileOptions := &compliancev2.ReplaceProfileOptions{
				ProfilesID:         core.StringPtr(profileIdNew),
				ProfileName:        core.StringPtr("SDK-automation-CustomProfile-Edit"),
				ProfileDescription: core.StringPtr("SDK-automation-CustomProfile-Edit"),
				ProfileType:        core.StringPtr("custom"),
				Controls:           []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				XCorrelationID:     core.StringPtr("SDK-automation-ReplaceProfile"),
				XRequestID:         core.StringPtr("SDK-automation-ReplaceProfile"),
			}

			profile, response, err := complianceService.ReplaceProfile(replaceProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`CreateAttachment - Create an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAttachment(createAttachmentOptions *CreateAttachmentOptions)`, func() {
			propertyScopeID := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr(accountID),
			}
			propertyScopeType := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			multiCloudScopeModel := &compliancev2.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []compliancev2.PropertyItem{*propertyScopeID, *propertyScopeType},
			}

			failedControlsModel := &compliancev2.FailedControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentsNotificationsPrototypeModel := &compliancev2.AttachmentsNotificationsPrototype{
				Enabled:  core.BoolPtr(false),
				Controls: failedControlsModel,
			}

			attachmentParametersPrototypeModel := &compliancev2.AttachmentParametersPrototype{
				AssessmentType:       core.StringPtr("Automated"),
				AssessmentID:         core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterValue:       core.StringPtr("120"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			attachmentsPrototypeModel := &compliancev2.AttachmentsPrototype{
				ID:                   core.StringPtr("130003ea8bfa43c5aacea07a86da3000"),
				Name:                 core.StringPtr("SDK-Automation-Attachment"),
				Description:          core.StringPtr("SDK-Automation-Attachment"),
				Scope:                []compliancev2.MultiCloudScope{*multiCloudScopeModel},
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("daily"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel},
			}

			createAttachmentOptions := &compliancev2.CreateAttachmentOptions{
				ProfilesID:     core.StringPtr(profileIdNew),
				Attachments:    []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel},
				ProfileID:      core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-CreateAttachment"),
				XRequestID:     core.StringPtr("SDK-automation-CreateAttachment"),
			}

			attachmentPrototype, response, err := complianceService.CreateAttachment(createAttachmentOptions)

			attachmentIdNew = *attachmentPrototype.Attachments[0].ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPrototype).ToNot(BeNil())
		})
	})

	Describe(`ListAttachments - Get all attachments linked to a specific profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions)`, func() {
			listAttachmentsOptions := &compliancev2.ListAttachmentsOptions{
				ProfilesID:     core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-ListAttachments"),
				XRequestID:     core.StringPtr("SDK-automation-ListAttachments"),
				Limit:          core.Int64Ptr(int64(50)),
			}

			attachmentCollection, response, err := complianceService.ListAttachments(listAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentCollection).ToNot(BeNil())
		})
	})

	Describe(`GetProfileAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions)`, func() {
			getProfileAttachmentOptions := &compliancev2.GetProfileAttachmentOptions{
				AttachmentID:   core.StringPtr(attachmentIdNew),
				ProfilesID:     core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-GetProfileAttachment"),
				XRequestID:     core.StringPtr("SDK-automation-GetProfileAttachment"),
			}

			attachment, response, err := complianceService.GetProfileAttachment(getProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
			propertyScopeID := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr(accountID),
			}
			propertyScopeType := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			multiCloudScopeModel := &compliancev2.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []compliancev2.PropertyItem{*propertyScopeID, *propertyScopeType},
			}

			failedControlsModel := &compliancev2.FailedControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentsNotificationsPrototypeModel := &compliancev2.AttachmentsNotificationsPrototype{
				Enabled:  core.BoolPtr(false),
				Controls: failedControlsModel,
			}

			attachmentParametersPrototypeModel := &compliancev2.AttachmentParametersPrototype{
				AssessmentType:       core.StringPtr("Automated"),
				AssessmentID:         core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterValue:       core.StringPtr("120"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds."),
				ParameterType:        core.StringPtr("numeric"),
			}

			lastScanModel := &compliancev2.LastScan{
				ID:     core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a"),
				Status: core.StringPtr("in_progress"),
				Time:   core.StringPtr("testString"),
			}

			replaceProfileAttachmentOptions := &compliancev2.ReplaceProfileAttachmentOptions{
				AttachmentID:         core.StringPtr(attachmentIdNew),
				ProfilesID:           core.StringPtr(profileIdNew),
				ID:                   core.StringPtr(attachmentIdNew),
				ProfileID:            core.StringPtr(profileIdNew),
				AccountID:            core.StringPtr(accountID),
				InstanceID:           core.StringPtr(instanceID),
				Scope:                []compliancev2.MultiCloudScope{*multiCloudScopeModel},
				CreatedOn:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				CreatedBy:            core.StringPtr("SDK-Automation"),
				UpdatedOn:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				UpdatedBy:            core.StringPtr("SDK-Automation"),
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("daily"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel},
				LastScan:             lastScanModel,
				NextScanTime:         core.StringPtr(""),
				Name:                 core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b"),
				Description:          core.StringPtr("Test description"),
				XCorrelationID:       core.StringPtr("SDK-automation-ReplaceProfileAttachment"),
				XRequestID:           core.StringPtr("SDK-automation-ReplaceProfileAttachment"),
			}

			attachment, response, err := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
	})

	Describe(`CreateScan - Create a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
			createScanOptions := &compliancev2.CreateScanOptions{
				AttachmentID:   core.StringPtr(attachmentIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-CreateScan"),
				XRequestID:     core.StringPtr("SDK-automation-CreateScan"),
			}

			scan, response, err := complianceService.CreateScan(createScanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scan).ToNot(BeNil())
		})
	})

	Describe(`ListAttachmentsAccount - Get all attachments in an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions)`, func() {
			listAttachmentsAccountOptions := &compliancev2.ListAttachmentsAccountOptions{
				XCorrelationID: core.StringPtr("SDK-automation-ListAttachmentsAccount"),
				XRequestID:     core.StringPtr("SDK-automation-ListAttachmentsAccount"),
				Limit:          core.Int64Ptr(int64(50)),
			}

			attachmentCollection, response, err := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentCollection).ToNot(BeNil())
		})
	})

	Describe(`DeleteProfileAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions)`, func() {
			deleteProfileAttachmentOptions := &compliancev2.DeleteProfileAttachmentOptions{
				AttachmentID:   core.StringPtr(attachmentIdNew),
				ProfilesID:     core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-DeleteProfileAttachment"),
				XRequestID:     core.StringPtr("SDK-automation-DeleteProfileAttachment"),
			}

			_, response, _ := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteCustomProfile - Delete a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
			deleteCustomProfileOptions := &compliancev2.DeleteCustomProfileOptions{
				ProfilesID:     core.StringPtr(profileIdNew),
				XCorrelationID: core.StringPtr("SDK-automation-DeleteCustomProfile"),
				XRequestID:     core.StringPtr("SDK-automation-DeleteCustomProfile"),
			}

			profile, response, err := complianceService.DeleteCustomProfile(deleteCustomProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomControlLibrary - Delete a control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions)`, func() {
			deleteCustomControlLibraryOptions := &compliancev2.DeleteCustomControlLibraryOptions{
				ControlLibrariesID: core.StringPtr(controlLibraryIdNew),
				XCorrelationID:     core.StringPtr("SDK-automation-DeleteCustomControlLibrary"),
				XRequestID:         core.StringPtr("SDK-automation-DeleteCustomControlLibrary"),
			}

			controlLibraryDelete, response, err := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryDelete).ToNot(BeNil())
		})
	})
})

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

package securityandcompliancecenterapiv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
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

var _ = Describe(`SecurityAndComplianceCenterApiV3 Integration Tests`, func() {
	const externalConfigFile = "../security_and_compliance_center_api_v3.env"

	var (
		err                                   error
		securityAndComplianceCenterApiService *securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3
		config                                map[string]string
		serviceURL                            string
		authenticator                         core.IamAuthenticator
		authUrl                               string
		apiKey                                string
		serviceName                           string

		// Variables to hold link values
		accountIdForReportLink                     string
		attachmentIdForReportLink                  string
		attachmentIdLink                           string
		controlLibraryIdLink                       string
		eTagLink                                   string
		eventNotificationsCrnForUpdateSettingsLink string
		groupIdForReportLink                       string
		objectStorageBucketForUpdateSettingsLink   string
		objectStorageCrnForUpdateSettingsLink      string
		objectStorageLocationForUpdateSettingsLink string
		profileIdForReportLink                     string
		profileIdLink                              string
		providerTypeIdLink                         string
		providerTypeInstanceIdLink                 string
		reportIdForReportLink                      string
		ruleIdLink                                 string
		typeForReportLink                          string
		accountID                                  string
		instanceID                                 string
		createScanAttachmentID                     string
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
			authUrl = config["IAM_APIKEY_URL"]
			if authUrl == "" {
				Skip("Unable to load auth service URL configuration property, skipping tests")
			}
			apiKey = config["IAM"]
			if apiKey == "" {
				Skip("Unable to load IAM configuration property, skipping tests")
			}
			serviceName = config["SERVICENAME"]
			if serviceName == "" {
				Skip("Unable to load SERVICENAME configuration property, skipping tests")
			}
			authenticator = core.IamAuthenticator{
				ApiKey: apiKey,
				URL:    authUrl,
			}
			accountID = config["ACCOUNTID"]
			instanceID = config["INSTANCEID"]
			if instanceID == "" {
				Skip("Unable to load instanceID configuration property, skipping tests")
			}
			createScanAttachmentID = config["ATTACHMENTID"]

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			securityAndComplianceCenterApiServiceOptions := &securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL:           serviceURL,
				Authenticator: &authenticator,
				ServiceName:   serviceName,
			}

			securityAndComplianceCenterApiService, err = securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(securityAndComplianceCenterApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
			Expect(securityAndComplianceCenterApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			securityAndComplianceCenterApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSettings - Get settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{
				XCorrelationID: core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5"),
				XRequestID:     core.StringPtr("testString"),
			}

			settings, response, err := securityAndComplianceCenterApiService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())

			eventNotificationsCrnForUpdateSettingsLink = *settings.EventNotifications.InstanceCrn
			fmt.Fprintf(GinkgoWriter, "Saved eventNotificationsCrnForUpdateSettingsLink value: %v\n", eventNotificationsCrnForUpdateSettingsLink)
			objectStorageCrnForUpdateSettingsLink = *settings.ObjectStorage.InstanceCrn
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageCrnForUpdateSettingsLink value: %v\n", objectStorageCrnForUpdateSettingsLink)
			objectStorageBucketForUpdateSettingsLink = *settings.ObjectStorage.Bucket
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageBucketForUpdateSettingsLink value: %v\n", objectStorageBucketForUpdateSettingsLink)
			objectStorageLocationForUpdateSettingsLink = *settings.ObjectStorage.BucketLocation
			fmt.Fprintf(GinkgoWriter, "Saved objectStorageLocationForUpdateSettingsLink value: %v\n", objectStorageLocationForUpdateSettingsLink)
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

			targetModel := &securityandcompliancecenterapiv3.Target{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ServiceDisplayName:         core.StringPtr("testString"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			requiredConfigItemsModel := &securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel},
			}

			parameterModel := &securityandcompliancecenterapiv3.Parameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterapiv3.Import{
				Parameters: []securityandcompliancecenterapiv3.Parameter{*parameterModel},
			}

			createRuleOptions := &securityandcompliancecenterapiv3.CreateRuleOptions{
				Description:    core.StringPtr("Example rule"),
				Target:         targetModel,
				RequiredConfig: requiredConfigModel,
				Type:           core.StringPtr("user_defined"),
				Version:        core.StringPtr("1.0.0"),
				Import:         importModel,
				Labels:         []string{},
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			rule, response, err := securityAndComplianceCenterApiService.CreateRule(createRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIdLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIdLink value: %v\n", ruleIdLink)
		})
	})

	Describe(`GetRule - Get a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &securityandcompliancecenterapiv3.GetRuleOptions{
				RuleID:         &ruleIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			rule, response, err := securityAndComplianceCenterApiService.GetRule(getRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			eTagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved eTagLink value: %v\n", eTagLink)
		})
	})

	Describe(`GetLatestReports - Get the latest reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
			getLatestReportsOptions := &securityandcompliancecenterapiv3.GetLatestReportsOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Sort:           core.StringPtr("profile_name"),
			}

			reportLatest, response, err := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportLatest).ToNot(BeNil())

			accountIdForReportLink = *reportLatest.Reports[0].Account.ID
			fmt.Fprintf(GinkgoWriter, "Saved accountIdForReportLink value: %v\n", accountIdForReportLink)
			reportIdForReportLink = *reportLatest.Reports[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved reportIdForReportLink value: %v\n", reportIdForReportLink)
			attachmentIdForReportLink = *reportLatest.Reports[0].Attachment.ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIdForReportLink value: %v\n", attachmentIdForReportLink)
			groupIdForReportLink = *reportLatest.Reports[0].GroupID
			fmt.Fprintf(GinkgoWriter, "Saved groupIdForReportLink value: %v\n", groupIdForReportLink)
			profileIdForReportLink = *reportLatest.Reports[0].Profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIdForReportLink value: %v\n", profileIdForReportLink)
			typeForReportLink = *reportLatest.Reports[0].Type
			fmt.Fprintf(GinkgoWriter, "Saved typeForReportLink value: %v\n", typeForReportLink)
		})
	})

	Describe(`UpdateSettings - Update settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			eventNotificationsModel := &securityandcompliancecenterapiv3.EventNotifications{
				InstanceCrn:       &eventNotificationsCrnForUpdateSettingsLink,
				UpdatedOn:         CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				SourceID:          core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"),
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("compliance"),
			}

			objectStorageModel := &securityandcompliancecenterapiv3.ObjectStorage{
				InstanceCrn:    &objectStorageCrnForUpdateSettingsLink,
				Bucket:         &objectStorageBucketForUpdateSettingsLink,
				BucketLocation: &objectStorageLocationForUpdateSettingsLink,
				BucketEndpoint: core.StringPtr("testString"),
				UpdatedOn:      CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			updateSettingsOptions := &securityandcompliancecenterapiv3.UpdateSettingsOptions{
				EventNotifications: eventNotificationsModel,
				ObjectStorage:      objectStorageModel,
				XCorrelationID:     core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5"),
				XRequestID:         core.StringPtr("testString"),
			}

			settings, response, err := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			Expect(settings).To(BeNil())
		})
	})

	/*Describe(`PostTestEvent - Create a test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
			postTestEventOptions := &securityandcompliancecenterapiv3.PostTestEventOptions{
				XCorrelationID: core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5"),
				XRequestID: core.StringPtr("testString"),
			}

			testEvent, response, err := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})
	})*/

	Describe(`CreateCustomControlLibrary - Create a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions)`, func() {
			parameterInfoModel := &securityandcompliancecenterapiv3.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
				ParameterValue:       core.StringPtr("public"),
			}

			implementationModel := &securityandcompliancecenterapiv3.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				ParameterCount:        core.Int64Ptr(int64(38)),
				Parameters:            []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &securityandcompliancecenterapiv3.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				Responsibility:                  core.StringPtr("user"),
				ComponentID:                     core.StringPtr("iam-identity"),
				ComponenetName:                  core.StringPtr("testString"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
				AssessmentsCount:                core.Int64Ptr(int64(38)),
				Assessments:                     []securityandcompliancecenterapiv3.Implementation{*implementationModel},
			}

			controlDocsModel := &securityandcompliancecenterapiv3.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &securityandcompliancecenterapiv3.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlParent:         core.StringPtr(""),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
				ControlRequirement:    core.BoolPtr(true),
				Status:                core.StringPtr("enabled"),
			}

			createCustomControlLibraryOptions := &securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions{
				ControlLibraryName:        core.StringPtr("IBM Cloud for Financial Services"),
				ControlLibraryDescription: core.StringPtr("IBM Cloud for Financial Services"),
				ControlLibraryType:        core.StringPtr("custom"),
				Controls:                  []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel},
				ControlLibraryVersion:     core.StringPtr("1.0.0"),
				Latest:                    core.BoolPtr(true),
				ControlsCount:             core.Int64Ptr(int64(38)),
				XCorrelationID:            core.StringPtr("testString"),
				XRequestID:                core.StringPtr("testString"),
			}

			controlLibrary, response, err := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIdLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIdLink value: %v\n", controlLibraryIdLink)
		})
	})

	Describe(`ListControlLibraries - Get control libraries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) with pagination`, func() {
			listControlLibrariesOptions := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
				Limit:              core.Int64Ptr(int64(50)),
				ControlLibraryType: core.StringPtr("custom"),
				Start:              core.StringPtr("testString"),
			}

			listControlLibrariesOptions.Start = nil
			listControlLibrariesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.ControlLibraryItem
			for {
				controlLibraryCollection, response, err := securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptions)
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
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
				Limit:              core.Int64Ptr(int64(50)),
				ControlLibraryType: core.StringPtr("custom"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.ControlLibraryItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterApiService.NewControlLibrariesPager(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListControlLibraries() returned a total of %d item(s) using ControlLibrariesPager.\n", len(allResults))
		})
	})

	Describe(`GetControlLibrary - Get a control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions)`, func() {
			getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{
				ControlLibrariesID: &controlLibraryIdLink,
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
			}

			controlLibrary, response, err := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptions)
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
			parameterInfoModel := &securityandcompliancecenterapiv3.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
				ParameterValue:       core.StringPtr("public"),
			}

			implementationModel := &securityandcompliancecenterapiv3.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				ParameterCount:        core.Int64Ptr(int64(38)),
				Parameters:            []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &securityandcompliancecenterapiv3.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				Responsibility:                  core.StringPtr("user"),
				ComponentID:                     core.StringPtr("iam-identity"),
				ComponenetName:                  core.StringPtr("testString"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
				AssessmentsCount:                core.Int64Ptr(int64(38)),
				Assessments:                     []securityandcompliancecenterapiv3.Implementation{*implementationModel},
			}

			controlDocsModel := &securityandcompliancecenterapiv3.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &securityandcompliancecenterapiv3.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlParent:         core.StringPtr(""),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
				ControlRequirement:    core.BoolPtr(true),
				Status:                core.StringPtr("enabled"),
			}

			replaceCustomControlLibraryOptions := &securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions{
				ControlLibrariesID:        &controlLibraryIdLink,
				ID:                        core.StringPtr("testString"),
				AccountID:                 core.StringPtr(accountID),
				ControlLibraryName:        core.StringPtr("IBM Cloud for Financial Services"),
				ControlLibraryDescription: core.StringPtr("IBM Cloud for Financial Services"),
				ControlLibraryType:        core.StringPtr("custom"),
				ControlLibraryVersion:     core.StringPtr("1.1.0"),
				CreatedOn:                 CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				CreatedBy:                 core.StringPtr("testString"),
				UpdatedOn:                 CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				UpdatedBy:                 core.StringPtr("testString"),
				Latest:                    core.BoolPtr(true),
				HierarchyEnabled:          core.BoolPtr(true),
				ControlsCount:             core.Int64Ptr(int64(38)),
				ControlParentsCount:       core.Int64Ptr(int64(38)),
				Controls:                  []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel},
				XCorrelationID:            core.StringPtr("testString"),
				XRequestID:                core.StringPtr("testString"),
			}

			controlLibrary, response, err := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
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
				ControlLibraryID: &controlLibraryIdLink,
				ControlID:        core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
			}

			defaultParametersPrototypeModel := &securityandcompliancecenterapiv3.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:         core.StringPtr("session_invalidation_in_seconds"),
				ParameterDefaultValue: core.StringPtr("120"),
				ParameterDisplayName:  core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:         core.StringPtr("numeric"),
			}

			createProfileOptions := &securityandcompliancecenterapiv3.CreateProfileOptions{
				ProfileName:        core.StringPtr("test_profile1"),
				ProfileDescription: core.StringPtr("test_description1"),
				ProfileType:        core.StringPtr("custom"),
				Controls:           []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
			}

			profile, response, err := securityAndComplianceCenterApiService.CreateProfile(createProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())

			profileIdLink = *profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIdLink value: %v\n", profileIdLink)
		})
	})

	Describe(`ListProfiles - Get all profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions) with pagination`, func() {
			listProfilesOptions := &securityandcompliancecenterapiv3.ListProfilesOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				ProfileType:    core.StringPtr("custom"),
				Start:          core.StringPtr("testString"),
			}

			listProfilesOptions.Start = nil
			listProfilesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.ProfileItem
			for {
				profileCollection, response, err := securityAndComplianceCenterApiService.ListProfiles(listProfilesOptions)
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
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				ProfileType:    core.StringPtr("custom"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.ProfileItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterApiService.NewProfilesPager(listProfilesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProfiles() returned a total of %d item(s) using ProfilesPager.\n", len(allResults))
		})
	})

	Describe(`GetProfile - Get a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
			getProfileOptions := &securityandcompliancecenterapiv3.GetProfileOptions{
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			profile, response, err := securityAndComplianceCenterApiService.GetProfile(getProfileOptions)
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
			profileControlsPrototypeModel := &securityandcompliancecenterapiv3.ProfileControlsPrototype{
				ControlLibraryID: &controlLibraryIdLink,
				ControlID:        core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
			}

			defaultParametersPrototypeModel := &securityandcompliancecenterapiv3.DefaultParametersPrototype{
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:         core.StringPtr("session_invalidation_in_seconds"),
				ParameterDefaultValue: core.StringPtr("120"),
				ParameterDisplayName:  core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:         core.StringPtr("numeric"),
			}

			replaceProfileOptions := &securityandcompliancecenterapiv3.ReplaceProfileOptions{
				ProfilesID:         &profileIdLink,
				ProfileName:        core.StringPtr("test_profile1"),
				ProfileDescription: core.StringPtr("test_description1"),
				ProfileType:        core.StringPtr("custom"),
				Controls:           []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				DefaultParameters:  []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
			}

			profile, response, err := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
	})

	Describe(`ListRules - List all rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
			listRulesOptions := &securityandcompliancecenterapiv3.ListRulesOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Type:           core.StringPtr("system_defined"),
				Search:         core.StringPtr("testString"),
				ServiceName:    core.StringPtr("testString"),
			}

			rulesPageBase, response, err := securityAndComplianceCenterApiService.ListRules(listRulesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesPageBase).ToNot(BeNil())
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

			targetModel := &securityandcompliancecenterapiv3.Target{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ServiceDisplayName:         core.StringPtr("Cloud Object Storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			requiredConfigItemsModel := &securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("hard_quota"),
				Operator:    core.StringPtr("num_equals"),
				Value:       core.StringPtr("${hard_quota}"),
			}

			requiredConfigModel := &securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd{
				Description: core.StringPtr("The Cloud Object Storage rule."),
				And:         []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel},
			}

			parameterModel := &securityandcompliancecenterapiv3.Parameter{
				Name:        core.StringPtr("hard_quota"),
				DisplayName: core.StringPtr("The Cloud Object Storage bucket quota."),
				Description: core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket."),
				Type:        core.StringPtr("numeric"),
			}

			importModel := &securityandcompliancecenterapiv3.Import{
				Parameters: []securityandcompliancecenterapiv3.Parameter{*parameterModel},
			}

			replaceRuleOptions := &securityandcompliancecenterapiv3.ReplaceRuleOptions{
				RuleID:         &ruleIdLink,
				IfMatch:        &eTagLink,
				Description:    core.StringPtr("Example rule"),
				Target:         targetModel,
				RequiredConfig: requiredConfigModel,
				Type:           core.StringPtr("user_defined"),
				Version:        core.StringPtr("1.0.1"),
				Import:         importModel,
				Labels:         []string{},
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			rule, response, err := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`CreateAttachment - Create an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAttachment(createAttachmentOptions *CreateAttachmentOptions)`, func() {
			propertyScopeID := &securityandcompliancecenterapiv3.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr(accountID),
			}
			propertyScopeType := &securityandcompliancecenterapiv3.PropertyItem{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			multiCloudScopeModel := &securityandcompliancecenterapiv3.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterapiv3.PropertyItem{*propertyScopeID, *propertyScopeType},
			}

			failedControlsModel := &securityandcompliancecenterapiv3.FailedControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentsNotificationsPrototypeModel := &securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype{
				Enabled:  core.BoolPtr(false),
				Controls: failedControlsModel,
			}

			attachmentParameterPrototypeModel := &securityandcompliancecenterapiv3.AttachmentParameterPrototype{
				AssessmentType:       core.StringPtr("Automated"),
				AssessmentID:         core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterValue:       core.StringPtr("120"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			attachmentsPrototypeModel := &securityandcompliancecenterapiv3.AttachmentsPrototype{
				ID:                   core.StringPtr("130003ea8bfa43c5aacea07a86da3000"),
				Name:                 core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b"),
				Description:          core.StringPtr("Test description"),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel},
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("every_30_days"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel},
			}

			createAttachmentOptions := &securityandcompliancecenterapiv3.CreateAttachmentOptions{
				ProfilesID:     &profileIdLink,
				Attachments:    []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel},
				ProfileID:      &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			attachmentPrototype, response, err := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(attachmentPrototype).ToNot(BeNil())

			attachmentIdLink = *attachmentPrototype.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIdLink value: %v\n", attachmentIdLink)
		})
	})

	Describe(`ListAttachments - Get all attachments linked to a specific profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) with pagination`, func() {
			listAttachmentsOptions := &securityandcompliancecenterapiv3.ListAttachmentsOptions{
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				Start:          core.StringPtr("testString"),
			}

			listAttachmentsOptions.Start = nil
			listAttachmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for {
				attachmentCollection, response, err := securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(attachmentCollection).ToNot(BeNil())
				allResults = append(allResults, attachmentCollection.Attachments...)

				listAttachmentsOptions.Start, err = attachmentCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listAttachmentsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) using AttachmentsPager`, func() {
			listAttachmentsOptions := &securityandcompliancecenterapiv3.ListAttachmentsOptions{
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewAttachmentsPager(listAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterApiService.NewAttachmentsPager(listAttachmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAttachments() returned a total of %d item(s) using AttachmentsPager.\n", len(allResults))
		})
	})

	Describe(`GetProfileAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions)`, func() {
			getProfileAttachmentOptions := &securityandcompliancecenterapiv3.GetProfileAttachmentOptions{
				AttachmentID:   &attachmentIdLink,
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			attachmentItem, response, err := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
			propertyScopeID := &securityandcompliancecenterapiv3.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr(accountID),
			}
			propertyScopeType := &securityandcompliancecenterapiv3.PropertyItem{
				Name:  core.StringPtr("scope_type"),
				Value: core.StringPtr("account"),
			}

			multiCloudScopeModel := &securityandcompliancecenterapiv3.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []securityandcompliancecenterapiv3.PropertyItem{*propertyScopeID, *propertyScopeType},
			}

			failedControlsModel := &securityandcompliancecenterapiv3.FailedControls{
				ThresholdLimit:   core.Int64Ptr(int64(15)),
				FailedControlIds: []string{},
			}

			attachmentsNotificationsPrototypeModel := &securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype{
				Enabled:  core.BoolPtr(false),
				Controls: failedControlsModel,
			}

			attachmentParameterPrototypeModel := &securityandcompliancecenterapiv3.AttachmentParameterPrototype{
				AssessmentType:       core.StringPtr("Automated"),
				AssessmentID:         core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterValue:       core.StringPtr("120"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			lastScanModel := &securityandcompliancecenterapiv3.LastScan{
				ID:     core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a"),
				Status: core.StringPtr("in_progress"),
				Time:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			replaceProfileAttachmentOptions := &securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions{
				AttachmentID:         &attachmentIdLink,
				ProfilesID:           &profileIdLink,
				ID:                   core.StringPtr("testString"),
				ProfileID:            &profileIdLink,
				AccountID:            core.StringPtr(accountID),
				InstanceID:           core.StringPtr(instanceID),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel},
				CreatedOn:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				CreatedBy:            core.StringPtr("testString"),
				UpdatedOn:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				UpdatedBy:            core.StringPtr("testString"),
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("every_30_days"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel},
				LastScan:             lastScanModel,
				NextScanTime:         CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Name:                 core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b"),
				Description:          core.StringPtr("Test description"),
				XCorrelationID:       core.StringPtr("testString"),
				XRequestID:           core.StringPtr("testString"),
			}

			attachmentItem, response, err := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
	})

	Describe(`CreateScan - Create a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
			createScanOptions := &securityandcompliancecenterapiv3.CreateScanOptions{
				AttachmentID:   &createScanAttachmentID,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			scan, response, err := securityAndComplianceCenterApiService.CreateScan(createScanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scan).ToNot(BeNil())
		})
	})

	Describe(`ListAttachmentsAccount - Get all attachments in an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions) with pagination`, func() {
			listAttachmentsAccountOptions := &securityandcompliancecenterapiv3.ListAttachmentsAccountOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				Start:          core.StringPtr("testString"),
			}

			listAttachmentsAccountOptions.Start = nil
			listAttachmentsAccountOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for {
				attachmentCollection, response, err := securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(attachmentCollection).ToNot(BeNil())
				allResults = append(allResults, attachmentCollection.Attachments...)

				listAttachmentsAccountOptions.Start, err = attachmentCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listAttachmentsAccountOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions) using AttachmentsAccountPager`, func() {
			listAttachmentsAccountOptions := &securityandcompliancecenterapiv3.ListAttachmentsAccountOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewAttachmentsAccountPager(listAttachmentsAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = securityAndComplianceCenterApiService.NewAttachmentsAccountPager(listAttachmentsAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAttachmentsAccount() returned a total of %d item(s) using AttachmentsAccountPager.\n", len(allResults))
		})
	})

	Describe(`ListReports - List reports`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) with pagination`, func() {
			listReportsOptions := &securityandcompliancecenterapiv3.ListReportsOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				AttachmentID:   &attachmentIdForReportLink,
				GroupID:        &groupIdForReportLink,
				ProfileID:      &profileIdForReportLink,
				Type:           &typeForReportLink,
				Start:          core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				Sort:           core.StringPtr("profile_name"),
			}

			listReportsOptions.Start = nil
			listReportsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Report
			for {
				reportPage, response, err := securityAndComplianceCenterApiService.ListReports(listReportsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reportPage).ToNot(BeNil())
				allResults = append(allResults, reportPage.Reports...)

				listReportsOptions.Start, err = reportPage.GetNextStart()
				Expect(err).To(BeNil())

				if listReportsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListReports(listReportsOptions *ListReportsOptions) using ReportsPager`, func() {
			listReportsOptions := &securityandcompliancecenterapiv3.ListReportsOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				AttachmentID:   &attachmentIdForReportLink,
				GroupID:        &groupIdForReportLink,
				ProfileID:      &profileIdForReportLink,
				Type:           &typeForReportLink,
				Limit:          core.Int64Ptr(int64(10)),
				Sort:           core.StringPtr("profile_name"),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewReportsPager(listReportsOptions)
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
			pager, err = securityAndComplianceCenterApiService.NewReportsPager(listReportsOptions)
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
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			report, response, err := securityAndComplianceCenterApiService.GetReport(getReportOptions)
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
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			reportSummary, response, err := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetReportEvaluation - Get report evaluation details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportEvaluation(getReportEvaluationOptions *GetReportEvaluationOptions)`, func() {
			getReportEvaluationOptions := &securityandcompliancecenterapiv3.GetReportEvaluationOptions{
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				ExcludeSummary: core.BoolPtr(true),
			}

			result, response, err := securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptions)
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
				ReportID:           &reportIdForReportLink,
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
				ControlID:          core.StringPtr("testString"),
				ControlName:        core.StringPtr("testString"),
				ControlDescription: core.StringPtr("testString"),
				ControlCategory:    core.StringPtr("testString"),
				Status:             core.StringPtr("compliant"),
				Sort:               core.StringPtr("control_name"),
			}

			reportControls, response, err := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportControls).ToNot(BeNil())
		})
	})

	Describe(`ListReportEvaluations - List report evaluations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) with pagination`, func() {
			listReportEvaluationsOptions := &securityandcompliancecenterapiv3.ListReportEvaluationsOptions{
				ReportID: &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID: core.StringPtr("testString"),
				AssessmentID: core.StringPtr("testString"),
				ComponentID: core.StringPtr("testString"),
				TargetID: core.StringPtr("testString"),
				TargetName: core.StringPtr("testString"),
				Status: core.StringPtr("failure"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listReportEvaluationsOptions.Start = nil
			listReportEvaluationsOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Evaluation
			for {
				evaluationPage, response, err := securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptions)
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
				ReportID: &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID: core.StringPtr("testString"),
				Status: core.StringPtr("failure"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
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
			pager, err = securityAndComplianceCenterApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())

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
				ReportID: &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				AccountID: &accountIdForReportLink,
				ComponentID: core.StringPtr("testString"),
				Status: core.StringPtr("compliant"),
				Sort: core.StringPtr("account_id"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listReportResourcesOptions.Start = nil
			listReportResourcesOptions.Limit = core.Int64Ptr(1)

			var allResults []securityandcompliancecenterapiv3.Resource
			for {
				resourcePage, response, err := securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptions)
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
				ReportID: &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID: core.StringPtr("testString"),
				AccountID: &accountIdForReportLink,
				Status: core.StringPtr("compliant"),
				Sort: core.StringPtr("account_id"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := securityAndComplianceCenterApiService.NewReportResourcesPager(listReportResourcesOptions)
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
			pager, err = securityAndComplianceCenterApiService.NewReportResourcesPager(listReportResourcesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListReportResources() returned a total of %d item(s) using ReportResourcesPager.\n", len(allResults))
		})
	})

	Describe(`GetReportTags - Get report tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
			getReportTagsOptions := &securityandcompliancecenterapiv3.GetReportTagsOptions{
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			reportTags, response, err := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptions)
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
				ReportID:         &reportIdForReportLink,
				XCorrelationID:   core.StringPtr("testString"),
				XRequestID:       core.StringPtr("testString"),
				ScanTimeDuration: core.Int64Ptr(int64(0)),
			}

			reportViolationsDrift, response, err := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reportViolationsDrift).ToNot(BeNil())
		})
	})

	Describe(`ListProviderTypes - List all provider types`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypes(listProviderTypesOptions *ListProviderTypesOptions)`, func() {
			listProviderTypesOptions := &securityandcompliancecenterapiv3.ListProviderTypesOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypesCollection, response, err := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypesCollection).ToNot(BeNil())

			providerTypeIdLink = *providerTypesCollection.ProviderTypes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeIdLink value: %v\n", providerTypeIdLink)
		})
	})

	Describe(`GetProviderTypeByID - Get a provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeByID(getProviderTypeByIdOptions *GetProviderTypeByIdOptions)`, func() {
			getProviderTypeByIdOptions := &securityandcompliancecenterapiv3.GetProviderTypeByIdOptions{
				ProviderTypeID: &providerTypeIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypeItem, response, err := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeItem).ToNot(BeNil())
		})
	})

	Describe(`ListProviderTypeInstances - List all provider type instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviderTypeInstances(listProviderTypeInstancesOptions *ListProviderTypeInstancesOptions)`, func() {
			listProviderTypeInstancesOptions := &securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions{
				ProviderTypeID: &providerTypeIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypeInstancesResponse, response, err := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstancesResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateProviderTypeInstance - Create a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProviderTypeInstance(createProviderTypeInstanceOptions *CreateProviderTypeInstanceOptions)`, func() {
			createProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions{
				ProviderTypeID: &providerTypeIdLink,
				Name:           core.StringPtr("workload-protection-instance-1"),
				Attributes:     map[string]interface{}{"wp_crn": "crn:v1:staging:public:sysdig-secure:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:0df4004c-fb74-483b-97be-dd9bd35af4d8::"},
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstanceItem).ToNot(BeNil())

			providerTypeInstanceIdLink = *providerTypeInstanceItem.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIdLink value: %v\n", providerTypeInstanceIdLink)
		})
	})

	Describe(`GetProviderTypeInstance - List a provider type instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypeInstance(getProviderTypeInstanceOptions *GetProviderTypeInstanceOptions)`, func() {
			getProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions{
				ProviderTypeID:         &providerTypeIdLink,
				ProviderTypeInstanceID: &providerTypeInstanceIdLink,
				XCorrelationID:         core.StringPtr("testString"),
				XRequestID:             core.StringPtr("testString"),
			}

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceItem).ToNot(BeNil())
		})
	})

	Describe(`UpdateProviderTypeInstance - Patch a specific instance of a provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProviderTypeInstance(updateProviderTypeInstanceOptions *UpdateProviderTypeInstanceOptions)`, func() {
			updateProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions{
				ProviderTypeID:         &providerTypeIdLink,
				ProviderTypeInstanceID: &providerTypeInstanceIdLink,
				Name:                   core.StringPtr("workload-protection-instance-1"),
				Attributes:             map[string]interface{}{"wp_crn": "crn:v1:staging:public:sysdig-secure:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:0df4004c-fb74-483b-97be-dd9bd35af4d8::"},
				XCorrelationID:         core.StringPtr("testString"),
				XRequestID:             core.StringPtr("testString"),
			}

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceItem).ToNot(BeNil())
		})
	})

	Describe(`GetProviderTypesInstances - Get a list of instances for all provider types`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProviderTypesInstances(getProviderTypesInstancesOptions *GetProviderTypesInstancesOptions)`, func() {
			getProviderTypesInstancesOptions := &securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypesInstancesResponse, response, err := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypesInstancesResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteProfileAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions)`, func() {
			deleteProfileAttachmentOptions := &securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions{
				AttachmentID:   &attachmentIdLink,
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			attachmentItem, response, err := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomProfile - Delete a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
			deleteCustomProfileOptions := &securityandcompliancecenterapiv3.DeleteCustomProfileOptions{
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			profile, response, err := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptions)
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
			deleteCustomControlLibraryOptions := &securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions{
				ControlLibrariesID: &controlLibraryIdLink,
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
			}

			controlLibraryDelete, response, err := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteRule - Delete a custom rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &securityandcompliancecenterapiv3.DeleteRuleOptions{
				RuleID:         &ruleIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			response, err := securityAndComplianceCenterApiService.DeleteRule(deleteRuleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteProviderTypeInstance - Remove a specific instance of a provider type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions *DeleteProviderTypeInstanceOptions)`, func() {
			deleteProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.DeleteProviderTypeInstanceOptions{
				ProviderTypeID:         &providerTypeIdLink,
				ProviderTypeInstanceID: &providerTypeInstanceIdLink,
				XCorrelationID:         core.StringPtr("testString"),
				XRequestID:             core.StringPtr("testString"),
			}

			response, err := securityAndComplianceCenterApiService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//

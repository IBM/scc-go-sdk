//go:build examples
// +build examples

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
var _ = Describe(`SecurityAndComplianceCenterApiV3 Examples Tests`, func() {

	const externalConfigFile = "../security_and_compliance_center_api_v3.env"

	var (
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

			securityAndComplianceCenterApiServiceOptions := &securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL:           serviceURL,
				Authenticator: &authenticator,
				ServiceName:   serviceName,
			}

			securityAndComplianceCenterApiService, err = securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(securityAndComplianceCenterApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
		})
	})

	Describe(`SecurityAndComplianceCenterApiV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := securityAndComplianceCenterApiService.NewGetSettingsOptions()
			getSettingsOptions.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")

			settings, response, err := securityAndComplianceCenterApiService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

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
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

			additionalTargetAttributeModel := &securityandcompliancecenterapiv3.AdditionalTargetAttribute{
				Name:     core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-east"),
			}

			targetModel := &securityandcompliancecenterapiv3.Target{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel},
			}

			requiredConfigItemsModel := &securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase{
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
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

			createRuleOptions := securityAndComplianceCenterApiService.NewCreateRuleOptions(
				"Example rule",
				targetModel,
				requiredConfigModel,
			)
			createRuleOptions.SetVersion("1.0.0")
			createRuleOptions.SetImport(importModel)
			createRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterApiService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			ruleIdLink = *rule.ID
			fmt.Fprintf(GinkgoWriter, "Saved ruleIdLink value: %v\n", ruleIdLink)
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := securityAndComplianceCenterApiService.NewGetRuleOptions(
				ruleIdLink,
			)

			rule, response, err := securityAndComplianceCenterApiService.GetRule(getRuleOptions)
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
		It(`GetLatestReports request example`, func() {
			fmt.Println("\nGetLatestReports() result:")
			// begin-get_latest_reports

			getLatestReportsOptions := securityAndComplianceCenterApiService.NewGetLatestReportsOptions()
			getLatestReportsOptions.SetSort("profile_name")

			reportLatest, response, err := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reportLatest, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_reports

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
		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			eventNotificationsModel := &securityandcompliancecenterapiv3.EventNotifications{
				InstanceCrn:       &eventNotificationsCrnForUpdateSettingsLink,
				SourceDescription: core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center."),
				SourceName:        core.StringPtr("compliance"),
			}

			objectStorageModel := &securityandcompliancecenterapiv3.ObjectStorage{
				InstanceCrn:    &objectStorageCrnForUpdateSettingsLink,
				Bucket:         &objectStorageBucketForUpdateSettingsLink,
				BucketLocation: &objectStorageLocationForUpdateSettingsLink,
			}

			updateSettingsOptions := securityAndComplianceCenterApiService.NewUpdateSettingsOptions()
			updateSettingsOptions.SetEventNotifications(eventNotificationsModel)
			updateSettingsOptions.SetObjectStorage(objectStorageModel)
			updateSettingsOptions.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")

			settings, response, err := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			Expect(settings).To(BeNil())
		})
		/*It(`PostTestEvent request example`, func() {
			fmt.Println("\nPostTestEvent() result:")
			// begin-post_test_event

			postTestEventOptions := securityAndComplianceCenterApiService.NewPostTestEventOptions()
			postTestEventOptions.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")

			testEvent, response, err := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())
		})*/
		It(`CreateCustomControlLibrary request example`, func() {
			fmt.Println("\nCreateCustomControlLibrary() result:")
			// begin-create_custom_control_library

			parameterInfoModel := &securityandcompliancecenterapiv3.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			implementationModel := &securityandcompliancecenterapiv3.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				Parameters:            []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &securityandcompliancecenterapiv3.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				ComponentID:                     core.StringPtr("iam-identity"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
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
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
				ControlRequirement:    core.BoolPtr(true),
			}

			createCustomControlLibraryOptions := securityAndComplianceCenterApiService.NewCreateCustomControlLibraryOptions(
				"IBM Cloud for Financial Services",
				"IBM Cloud for Financial Services",
				"custom",
				[]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel},
			)
			createCustomControlLibraryOptions.SetVersionGroupLabel("")
			createCustomControlLibraryOptions.SetControlLibraryVersion("1.0.0")

			controlLibrary, response, err := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-create_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibrary).ToNot(BeNil())

			controlLibraryIdLink = *controlLibrary.ID
			fmt.Fprintf(GinkgoWriter, "Saved controlLibraryIdLink value: %v\n", controlLibraryIdLink)
		})
		It(`ListControlLibraries request example`, func() {
			fmt.Println("\nListControlLibraries() result:")
			// begin-list_control_libraries
			listControlLibrariesOptions := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
				XCorrelationID:     core.StringPtr("testString"),
				XRequestID:         core.StringPtr("testString"),
				Limit:              core.Int64Ptr(int64(50)),
				ControlLibraryType: core.StringPtr("custom"),
			}

			pager, err := securityAndComplianceCenterApiService.NewControlLibrariesPager(listControlLibrariesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.ControlLibraryItem
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
		It(`GetControlLibrary request example`, func() {
			fmt.Println("\nGetControlLibrary() result:")
			// begin-get_control_library

			getControlLibraryOptions := securityAndComplianceCenterApiService.NewGetControlLibraryOptions(
				controlLibraryIdLink,
			)

			controlLibrary, response, err := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptions)
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
		It(`ReplaceCustomControlLibrary request example`, func() {
			fmt.Println("\nReplaceCustomControlLibrary() result:")
			// begin-replace_custom_control_library

			parameterInfoModel := &securityandcompliancecenterapiv3.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			implementationModel := &securityandcompliancecenterapiv3.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				Parameters:            []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &securityandcompliancecenterapiv3.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				Responsibility:                  core.StringPtr("user"),
				ComponentID:                     core.StringPtr("iam-identity"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
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
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
				ControlRequirement:    core.BoolPtr(true),
			}

			replaceCustomControlLibraryOptions := securityAndComplianceCenterApiService.NewReplaceCustomControlLibraryOptions(
				controlLibraryIdLink,
			)
			replaceCustomControlLibraryOptions.SetControlLibraryName("IBM Cloud for Financial Services")
			replaceCustomControlLibraryOptions.SetControlLibraryDescription("IBM Cloud for Financial Services")
			replaceCustomControlLibraryOptions.SetControlLibraryType("custom")
			replaceCustomControlLibraryOptions.SetControlLibraryVersion("1.1.0")
			replaceCustomControlLibraryOptions.SetControls([]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel})

			controlLibrary, response, err := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
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
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

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

			createProfileOptions := securityAndComplianceCenterApiService.NewCreateProfileOptions(
				"test_profile1",
				"test_description1",
				"custom",
				[]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				[]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
			)

			profile, response, err := securityAndComplianceCenterApiService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

			profileIdLink = *profile.ID
			fmt.Fprintf(GinkgoWriter, "Saved profileIdLink value: %v\n", profileIdLink)
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles
			listProfilesOptions := &securityandcompliancecenterapiv3.ListProfilesOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
				ProfileType:    core.StringPtr("custom"),
			}

			pager, err := securityAndComplianceCenterApiService.NewProfilesPager(listProfilesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.ProfileItem
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
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := securityAndComplianceCenterApiService.NewGetProfileOptions(
				profileIdLink,
			)

			profile, response, err := securityAndComplianceCenterApiService.GetProfile(getProfileOptions)
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
		It(`ReplaceProfile request example`, func() {
			fmt.Println("\nReplaceProfile() result:")
			// begin-replace_profile

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

			replaceProfileOptions := securityAndComplianceCenterApiService.NewReplaceProfileOptions(
				profileIdLink,
				"test_profile1",
				"test_description1",
				"custom",
				[]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel},
				[]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel},
			)

			profile, response, err := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptions)
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
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := securityAndComplianceCenterApiService.NewListRulesOptions()
			listRulesOptions.SetType("system_defined")

			rulesPageBase, response, err := securityAndComplianceCenterApiService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rulesPageBase, "", "  ")
			fmt.Println(string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesPageBase).ToNot(BeNil())
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

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
				Property: core.StringPtr("hard_quota"),
				Operator: core.StringPtr("num_equals"),
				Value:    core.StringPtr("${hard_quota}"),
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

			replaceRuleOptions := securityAndComplianceCenterApiService.NewReplaceRuleOptions(
				ruleIdLink,
				eTagLink,
				"Example rule",
				targetModel,
				requiredConfigModel,
			)
			replaceRuleOptions.SetType("user_defined")
			replaceRuleOptions.SetVersion("1.0.1")
			replaceRuleOptions.SetImport(importModel)
			replaceRuleOptions.SetLabels([]string{})

			rule, response, err := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptions)
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
		It(`CreateAttachment request example`, func() {
			fmt.Println("\nCreateAttachment() result:")
			// begin-create_attachment

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
				Name:                 core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b"),
				Description:          core.StringPtr("Test description"),
				Scope:                []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel},
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("every_30_days"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel},
			}

			createAttachmentOptions := securityAndComplianceCenterApiService.NewCreateAttachmentOptions(
				profileIdLink,
				[]securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel},
			)

			attachmentPrototype, response, err := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentPrototype, "", "  ")
			fmt.Println(string(b))

			// end-create_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPrototype).ToNot(BeNil())

			attachmentIdLink = *attachmentPrototype.Attachments[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved attachmentIdLink value: %v\n", attachmentIdLink)
		})
		It(`ListAttachments request example`, func() {
			fmt.Println("\nListAttachments() result:")
			// begin-list_attachments
			listAttachmentsOptions := &securityandcompliancecenterapiv3.ListAttachmentsOptions{
				ProfilesID:     &profileIdLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterApiService.NewAttachmentsPager(listAttachmentsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_attachments
		})
		It(`GetProfileAttachment request example`, func() {
			fmt.Println("\nGetProfileAttachment() result:")
			// begin-get_profile_attachment

			getProfileAttachmentOptions := securityAndComplianceCenterApiService.NewGetProfileAttachmentOptions(
				attachmentIdLink,
				profileIdLink,
			)

			attachmentItem, response, err := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentItem, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
		It(`ReplaceProfileAttachment request example`, func() {
			fmt.Println("\nReplaceProfileAttachment() result:")
			// begin-replace_profile_attachment

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

			replaceProfileAttachmentOptions := securityAndComplianceCenterApiService.NewReplaceProfileAttachmentOptions(
				attachmentIdLink,
				profileIdLink,
			)
			replaceProfileAttachmentOptions.SetScope([]securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel})
			replaceProfileAttachmentOptions.SetStatus("enabled")
			replaceProfileAttachmentOptions.SetSchedule("every_30_days")
			replaceProfileAttachmentOptions.SetNotifications(attachmentsNotificationsPrototypeModel)
			replaceProfileAttachmentOptions.SetAttachmentParameters([]securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel})
			replaceProfileAttachmentOptions.SetName("account-0d8c3805dfea40aa8ad02265a18eb12b")
			replaceProfileAttachmentOptions.SetDescription("Test description")

			attachmentItem, response, err := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentItem, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
		It(`CreateScan request example`, func() {
			fmt.Println("\nCreateScan() result:")
			// begin-create_scan

			createScanOptions := securityAndComplianceCenterApiService.NewCreateScanOptions(
				createScanAttachmentID,
			)

			scan, response, err := securityAndComplianceCenterApiService.CreateScan(createScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scan, "", "  ")
			fmt.Println(string(b))

			// end-create_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scan).ToNot(BeNil())
		})
		It(`ListAttachmentsAccount request example`, func() {
			fmt.Println("\nListAttachmentsAccount() result:")
			// begin-list_attachments_account
			listAttachmentsAccountOptions := &securityandcompliancecenterapiv3.ListAttachmentsAccountOptions{
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterApiService.NewAttachmentsAccountPager(listAttachmentsAccountOptions)
			if err != nil {
				panic(err)
			}

			var allResults []securityandcompliancecenterapiv3.AttachmentItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_attachments_account
		})
		It(`ListReports request example`, func() {
			fmt.Println("\nListReports() result:")
			// begin-list_reports
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

			pager, err := securityAndComplianceCenterApiService.NewReportsPager(listReportsOptions)
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

			getReportOptions := securityAndComplianceCenterApiService.NewGetReportOptions(
				reportIdForReportLink,
			)

			report, response, err := securityAndComplianceCenterApiService.GetReport(getReportOptions)
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

			getReportSummaryOptions := securityAndComplianceCenterApiService.NewGetReportSummaryOptions(
				reportIdForReportLink,
			)

			reportSummary, response, err := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptions)
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
		It(`GetReportEvaluation request example`, func() {
			fmt.Println("\nGetReportEvaluation() result:")
			// begin-get_report_evaluation

			getReportEvaluationOptions := securityAndComplianceCenterApiService.NewGetReportEvaluationOptions(
				reportIdForReportLink,
			)

			result, response, err := securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptions)
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

			// end-get_report_evaluation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`GetReportControls request example`, func() {
			fmt.Println("\nGetReportControls() result:")
			// begin-get_report_controls

			getReportControlsOptions := securityAndComplianceCenterApiService.NewGetReportControlsOptions(
				reportIdForReportLink,
			)
			getReportControlsOptions.SetStatus("compliant")

			reportControls, response, err := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptions)
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
		It(`ListReportEvaluations request example`, func() {
			fmt.Println("\nListReportEvaluations() result:")
			// begin-list_report_evaluations
			listReportEvaluationsOptions := &securityandcompliancecenterapiv3.ListReportEvaluationsOptions{
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				AssessmentID:   core.StringPtr("testString"),
				ComponentID:    core.StringPtr("testString"),
				TargetID:       core.StringPtr("testString"),
				TargetName:     core.StringPtr("testString"),
				Status:         core.StringPtr("failure"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterApiService.NewReportEvaluationsPager(listReportEvaluationsOptions)
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
				ReportID:       &reportIdForReportLink,
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
				ResourceName:   core.StringPtr("testString"),
				AccountID:      &accountIdForReportLink,
				ComponentID:    core.StringPtr("testString"),
				Status:         core.StringPtr("compliant"),
				Sort:           core.StringPtr("account_id"),
				Limit:          core.Int64Ptr(int64(10)),
			}

			pager, err := securityAndComplianceCenterApiService.NewReportResourcesPager(listReportResourcesOptions)
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

			getReportTagsOptions := securityAndComplianceCenterApiService.NewGetReportTagsOptions(
				reportIdForReportLink,
			)

			reportTags, response, err := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptions)
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

			getReportViolationsDriftOptions := securityAndComplianceCenterApiService.NewGetReportViolationsDriftOptions(
				reportIdForReportLink,
			)

			reportViolationsDrift, response, err := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptions)
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
		It(`ListProviderTypes request example`, func() {
			fmt.Println("\nListProviderTypes() result:")
			// begin-list_provider_types

			listProviderTypesOptions := securityAndComplianceCenterApiService.NewListProviderTypesOptions()

			providerTypesCollection, response, err := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypesCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_types

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypesCollection).ToNot(BeNil())

			providerTypeIdLink = *providerTypesCollection.ProviderTypes[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeIdLink value: %v\n", providerTypeIdLink)
		})
		It(`GetProviderTypeByID request example`, func() {
			fmt.Println("\nGetProviderTypeByID() result:")
			// begin-get_provider_type_by_id

			getProviderTypeByIdOptions := securityAndComplianceCenterApiService.NewGetProviderTypeByIdOptions(
				providerTypeIdLink,
			)

			providerTypeItem, response, err := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeItem, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_by_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeItem).ToNot(BeNil())
		})
		It(`ListProviderTypeInstances request example`, func() {
			fmt.Println("\nListProviderTypeInstances() result:")
			// begin-list_provider_type_instances

			listProviderTypeInstancesOptions := securityAndComplianceCenterApiService.NewListProviderTypeInstancesOptions(
				providerTypeIdLink,
			)

			providerTypeInstancesResponse, response, err := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstancesResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_provider_type_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstancesResponse).ToNot(BeNil())
		})
		It(`CreateProviderTypeInstance request example`, func() {
			fmt.Println("\nCreateProviderTypeInstance() result:")
			// begin-create_provider_type_instance

			createProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions{
				ProviderTypeID: &providerTypeIdLink,
				Name:           core.StringPtr("workload-protection-instance-1"),
				Attributes:     map[string]interface{}{"wp_crn": "crn:v1:staging:public:sysdig-secure:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:0df4004c-fb74-483b-97be-dd9bd35af4d8::"},
				XCorrelationID: core.StringPtr("testString"),
				XRequestID:     core.StringPtr("testString"),
			}

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstanceItem, "", "  ")
			fmt.Println(string(b))

			// end-create_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(providerTypeInstanceItem).ToNot(BeNil())

			providerTypeInstanceIdLink = *providerTypeInstanceItem.ID
			fmt.Fprintf(GinkgoWriter, "Saved providerTypeInstanceIdLink value: %v\n", providerTypeInstanceIdLink)
		})
		It(`GetProviderTypeInstance request example`, func() {
			fmt.Println("\nGetProviderTypeInstance() result:")
			// begin-get_provider_type_instance

			getProviderTypeInstanceOptions := securityAndComplianceCenterApiService.NewGetProviderTypeInstanceOptions(
				providerTypeIdLink,
				providerTypeInstanceIdLink,
			)

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstanceItem, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceItem).ToNot(BeNil())
		})
		It(`UpdateProviderTypeInstance request example`, func() {
			fmt.Println("\nUpdateProviderTypeInstance() result:")
			// begin-update_provider_type_instance

			updateProviderTypeInstanceOptions := &securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions{
				ProviderTypeID:         &providerTypeIdLink,
				ProviderTypeInstanceID: &providerTypeInstanceIdLink,
				Name:                   core.StringPtr("workload-protection-instance-1"),
				Attributes:             map[string]interface{}{"wp_crn": "crn:v1:staging:public:sysdig-secure:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:0df4004c-fb74-483b-97be-dd9bd35af4d8::"},
				XCorrelationID:         core.StringPtr("testString"),
				XRequestID:             core.StringPtr("testString"),
			}

			providerTypeInstanceItem, response, err := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypeInstanceItem, "", "  ")
			fmt.Println(string(b))

			// end-update_provider_type_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypeInstanceItem).ToNot(BeNil())
		})
		It(`GetProviderTypesInstances request example`, func() {
			fmt.Println("\nGetProviderTypesInstances() result:")
			// begin-get_provider_types_instances

			getProviderTypesInstancesOptions := securityAndComplianceCenterApiService.NewGetProviderTypesInstancesOptions()

			providerTypesInstancesResponse, response, err := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(providerTypesInstancesResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_provider_types_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(providerTypesInstancesResponse).ToNot(BeNil())
		})
		It(`DeleteProfileAttachment request example`, func() {
			fmt.Println("\nDeleteProfileAttachment() result:")
			// begin-delete_profile_attachment

			deleteProfileAttachmentOptions := securityAndComplianceCenterApiService.NewDeleteProfileAttachmentOptions(
				attachmentIdLink,
				profileIdLink,
			)

			attachmentItem, response, err := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentItem, "", "  ")
			fmt.Println(string(b))

			// end-delete_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentItem).ToNot(BeNil())
		})
		It(`DeleteCustomProfile request example`, func() {
			fmt.Println("\nDeleteCustomProfile() result:")
			// begin-delete_custom_profile

			deleteCustomProfileOptions := securityAndComplianceCenterApiService.NewDeleteCustomProfileOptions(
				profileIdLink,
			)

			profile, response, err := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptions)
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
		It(`DeleteCustomControlLibrary request example`, func() {
			fmt.Println("\nDeleteCustomControlLibrary() result:")
			// begin-delete_custom_control_library

			deleteCustomControlLibraryOptions := securityAndComplianceCenterApiService.NewDeleteCustomControlLibraryOptions(
				controlLibraryIdLink,
			)

			controlLibraryDelete, response, err := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryDelete, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryDelete).ToNot(BeNil())
		})
		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := securityAndComplianceCenterApiService.NewDeleteRuleOptions(
				ruleIdLink,
			)

			response, err := securityAndComplianceCenterApiService.DeleteRule(deleteRuleOptions)
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
		It(`DeleteProviderTypeInstance request example`, func() {
			// begin-delete_provider_type_instance

			deleteProviderTypeInstanceOptions := securityAndComplianceCenterApiService.NewDeleteProviderTypeInstanceOptions(
				providerTypeIdLink,
				providerTypeInstanceIdLink,
			)

			response, err := securityAndComplianceCenterApiService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions)
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
	})
})

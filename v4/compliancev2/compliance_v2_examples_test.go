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

package compliancev2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/compliancev2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the compliance service.
//
// The following configuration properties are assumed to be defined:
// COMPLIANCE_URL=<service base url>
// COMPLIANCE_AUTH_TYPE=iam
// COMPLIANCE_APIKEY=<IAM apikey>
// COMPLIANCE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ComplianceV2 Examples Tests`, func() {

	const externalConfigFile = "compliance_v2.env"

	var (
		complianceService *compliancev2.ComplianceV2
		config            map[string]string
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
			config, err = core.GetServiceProperties(compliancev2.DefaultServiceName)
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

			complianceServiceOptions := &compliancev2.ComplianceV2Options{}

			complianceService, err = compliancev2.NewComplianceV2UsingExternalConfig(complianceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(complianceService).ToNot(BeNil())
		})
	})

	Describe(`ComplianceV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries request example`, func() {
			fmt.Println("\nListControlLibraries() result:")
			// begin-list_control_libraries

			listControlLibrariesOptions := complianceService.NewListControlLibrariesOptions()

			controlLibraryCollection, response, err := complianceService.ListControlLibraries(listControlLibrariesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_control_libraries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryCollection).ToNot(BeNil())
		})
		It(`CreateCustomControlLibrary request example`, func() {
			fmt.Println("\nCreateCustomControlLibrary() result:")
			// begin-create_custom_control_library

			parameterInfoModel := &compliancev2.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			implementationModel := &compliancev2.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				Parameters:            []compliancev2.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &compliancev2.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				ComponentID:                     core.StringPtr("iam-identity"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
				Assessments:                     []compliancev2.Implementation{*implementationModel},
			}

			controlDocsModel := &compliancev2.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &compliancev2.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlRequirement:    core.BoolPtr(false),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []compliancev2.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
			}

			createCustomControlLibraryOptions := complianceService.NewCreateCustomControlLibraryOptions(
				"IBM Cloud for Financial Services",
				"IBM Cloud for Financial Services",
				"custom",
				[]compliancev2.ControlsInControlLib{*controlsInControlLibModel},
			)
			createCustomControlLibraryOptions.SetVersionGroupLabel("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
			createCustomControlLibraryOptions.SetControlLibraryVersion("1.1.0")

			controlLibrary, response, err := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibrary, "", "  ")
			fmt.Println(string(b))

			// end-create_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(controlLibrary).ToNot(BeNil())
		})
		It(`GetControlLibrary request example`, func() {
			fmt.Println("\nGetControlLibrary() result:")
			// begin-get_control_library

			getControlLibraryOptions := complianceService.NewGetControlLibraryOptions(
				"testString",
			)

			controlLibrary, response, err := complianceService.GetControlLibrary(getControlLibraryOptions)
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

			parameterInfoModel := &compliancev2.ParameterInfo{
				ParameterName:        core.StringPtr("session_invalidation_in_seconds"),
				ParameterDisplayName: core.StringPtr("Sign out due to inactivity in seconds"),
				ParameterType:        core.StringPtr("numeric"),
			}

			implementationModel := &compliancev2.Implementation{
				AssessmentID:          core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1"),
				AssessmentMethod:      core.StringPtr("ibm-cloud-rule"),
				AssessmentType:        core.StringPtr("Automated"),
				AssessmentDescription: core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services"),
				Parameters:            []compliancev2.ParameterInfo{*parameterInfoModel},
			}

			controlSpecificationsModel := &compliancev2.ControlSpecifications{
				ControlSpecificationID:          core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184"),
				Responsibility:                  core.StringPtr("user"),
				ComponentID:                     core.StringPtr("iam-identity"),
				Environment:                     core.StringPtr("ibm-cloud"),
				ControlSpecificationDescription: core.StringPtr("IBM cloud"),
				Assessments:                     []compliancev2.Implementation{*implementationModel},
			}

			controlDocsModel := &compliancev2.ControlDocs{
				ControlDocsID:   core.StringPtr("sc-7"),
				ControlDocsType: core.StringPtr("ibm-cloud"),
			}

			controlsInControlLibModel := &compliancev2.ControlsInControlLib{
				ControlName:           core.StringPtr("SC-7"),
				ControlID:             core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790"),
				ControlDescription:    core.StringPtr("Boundary Protection"),
				ControlCategory:       core.StringPtr("System and Communications Protection"),
				ControlRequirement:    core.BoolPtr(false),
				ControlTags:           []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"},
				ControlSpecifications: []compliancev2.ControlSpecifications{*controlSpecificationsModel},
				ControlDocs:           controlDocsModel,
			}

			replaceCustomControlLibraryOptions := complianceService.NewReplaceCustomControlLibraryOptions(
				"testString",
			)
			replaceCustomControlLibraryOptions.SetControlLibraryName("IBM Cloud for Financial Services")
			replaceCustomControlLibraryOptions.SetControlLibraryDescription("IBM Cloud for Financial Services")
			replaceCustomControlLibraryOptions.SetControlLibraryType("custom")
			replaceCustomControlLibraryOptions.SetControls([]compliancev2.ControlsInControlLib{*controlsInControlLibModel})

			controlLibrary, response, err := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
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
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := complianceService.NewListProfilesOptions()

			profileCollection, response, err := complianceService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileCollection).ToNot(BeNil())
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

			profileControlsPrototypeModel := &compliancev2.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd"),
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

			createProfileOptions := complianceService.NewCreateProfileOptions(
				"test_profile1",
				"test_description1",
				"custom",
				[]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel},
				[]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel},
			)

			profile, response, err := complianceService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := complianceService.NewGetProfileOptions(
				"testString",
			)

			profile, response, err := complianceService.GetProfile(getProfileOptions)
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

			profileControlsPrototypeModel := &compliancev2.ProfileControlsPrototype{
				ControlLibraryID: core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd"),
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

			replaceProfileOptions := complianceService.NewReplaceProfileOptions(
				"testString",
				"test_profile1",
				"test_description1",
				"custom",
				[]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel},
				[]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel},
			)

			profile, response, err := complianceService.ReplaceProfile(replaceProfileOptions)
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
		It(`ListAttachments request example`, func() {
			fmt.Println("\nListAttachments() result:")
			// begin-list_attachments

			listAttachmentsOptions := complianceService.NewListAttachmentsOptions(
				"testString",
			)

			attachmentCollection, response, err := complianceService.ListAttachments(listAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentCollection).ToNot(BeNil())
		})
		It(`CreateAttachment request example`, func() {
			fmt.Println("\nCreateAttachment() result:")
			// begin-create_attachment

			propertyModel := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("cg3335893hh1428692d6747cf300yeb5"),
			}

			multiCloudScopeModel := &compliancev2.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []compliancev2.PropertyItem{*propertyModel},
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
				Name:                 core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b"),
				Description:          core.StringPtr("Test description"),
				Scope:                []compliancev2.MultiCloudScope{*multiCloudScopeModel},
				Status:               core.StringPtr("enabled"),
				Schedule:             core.StringPtr("daily"),
				Notifications:        attachmentsNotificationsPrototypeModel,
				AttachmentParameters: []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel},
			}

			createAttachmentOptions := complianceService.NewCreateAttachmentOptions(
				"testString",
				[]compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel},
			)

			attachmentPrototype, response, err := complianceService.CreateAttachment(createAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentPrototype, "", "  ")
			fmt.Println(string(b))

			// end-create_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(attachmentPrototype).ToNot(BeNil())
		})
		It(`GetProfileAttachment request example`, func() {
			fmt.Println("\nGetProfileAttachment() result:")
			// begin-get_profile_attachment

			getProfileAttachmentOptions := complianceService.NewGetProfileAttachmentOptions(
				"testString",
				"testString",
			)

			attachment, response, err := complianceService.GetProfileAttachment(getProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
		It(`ReplaceProfileAttachment request example`, func() {
			fmt.Println("\nReplaceProfileAttachment() result:")
			// begin-replace_profile_attachment

			propertyModel := &compliancev2.PropertyItem{
				Name:  core.StringPtr("scope_id"),
				Value: core.StringPtr("cg3335893hh1428692d6747cf300yeb5"),
			}

			multiCloudScopeModel := &compliancev2.MultiCloudScope{
				Environment: core.StringPtr("ibm-cloud"),
				Properties:  []compliancev2.PropertyItem{*propertyModel},
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

			replaceProfileAttachmentOptions := complianceService.NewReplaceProfileAttachmentOptions(
				"testString",
				"testString",
			)
			replaceProfileAttachmentOptions.SetScope([]compliancev2.MultiCloudScope{*multiCloudScopeModel})
			replaceProfileAttachmentOptions.SetStatus("enabled")
			replaceProfileAttachmentOptions.SetSchedule("daily")
			replaceProfileAttachmentOptions.SetNotifications(attachmentsNotificationsPrototypeModel)
			replaceProfileAttachmentOptions.SetAttachmentParameters([]compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel})
			replaceProfileAttachmentOptions.SetName("account-0d8c3805dfea40aa8ad02265a18eb12b")
			replaceProfileAttachmentOptions.SetDescription("Test description")

			attachment, response, err := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
		It(`CreateScan request example`, func() {
			fmt.Println("\nCreateScan() result:")
			// begin-create_scan

			createScanOptions := complianceService.NewCreateScanOptions(
				"testString",
			)

			scan, response, err := complianceService.CreateScan(createScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scan, "", "  ")
			fmt.Println(string(b))

			// end-create_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(scan).ToNot(BeNil())
		})
		It(`ListAttachmentsAccount request example`, func() {
			fmt.Println("\nListAttachmentsAccount() result:")
			// begin-list_attachments_account

			listAttachmentsAccountOptions := complianceService.NewListAttachmentsAccountOptions()

			attachmentCollection, response, err := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_attachments_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentCollection).ToNot(BeNil())
		})
		It(`DeleteCustomControlLibrary request example`, func() {
			fmt.Println("\nDeleteCustomControlLibrary() result:")
			// begin-delete_custom_control_library

			deleteCustomControlLibraryOptions := complianceService.NewDeleteCustomControlLibraryOptions(
				"testString",
			)

			controlLibraryDelete, response, err := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptions)
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
		It(`DeleteCustomProfile request example`, func() {
			fmt.Println("\nDeleteCustomProfile() result:")
			// begin-delete_custom_profile

			deleteCustomProfileOptions := complianceService.NewDeleteCustomProfileOptions(
				"testString",
			)

			profile, response, err := complianceService.DeleteCustomProfile(deleteCustomProfileOptions)
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
		It(`DeleteProfileAttachment request example`, func() {
			fmt.Println("\nDeleteProfileAttachment() result:")
			// begin-delete_profile_attachment

			deleteProfileAttachmentOptions := complianceService.NewDeleteProfileAttachmentOptions(
				"testString",
				"testString",
			)

			attachment, response, err := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachment, "", "  ")
			fmt.Println(string(b))

			// end-delete_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachment).ToNot(BeNil())
		})
	})
})

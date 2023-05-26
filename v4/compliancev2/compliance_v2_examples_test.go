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

// This file provides an example of how to use the SCC Phoenix Compliance Apis service.
//
// The following configuration properties are assumed to be defined:
// SCC_PHOENIX_COMPLIANCE_APIS_URL=<service base url>
// SCC_PHOENIX_COMPLIANCE_APIS_AUTH_TYPE=iam
// SCC_PHOENIX_COMPLIANCE_APIS_APIKEY=<IAM apikey>
// SCC_PHOENIX_COMPLIANCE_APIS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`SccPhoenixComplianceApisV1 Examples Tests`, func() {

	const externalConfigFile = "compliancev2.env"

	var (
		sccPhoenixComplianceApisService *compliancev2.SccPhoenixComplianceApisV1
		config                          map[string]string
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

			sccPhoenixComplianceApisServiceOptions := &compliancev2.SccPhoenixComplianceApisV1Options{}

			sccPhoenixComplianceApisService, err = compliancev2.NewSccPhoenixComplianceApisV1UsingExternalConfig(sccPhoenixComplianceApisServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
		})
	})

	Describe(`SccPhoenixComplianceApisV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

			createProfileOptions := sccPhoenixComplianceApisService.NewCreateProfileOptions(
				"testString",
			)

			profileResponse, response, err := sccPhoenixComplianceApisService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := sccPhoenixComplianceApisService.NewListProfilesOptions(
				"testString",
			)

			getAllProfilesRespBody, response, err := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getAllProfilesRespBody, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllProfilesRespBody).ToNot(BeNil())
		})
		It(`AddProfile request example`, func() {
			fmt.Println("\nAddProfile() result:")
			// begin-add_profile

			addProfileOptions := sccPhoenixComplianceApisService.NewAddProfileOptions(
				"testString",
				"testString",
			)

			profileResponse, response, err := sccPhoenixComplianceApisService.AddProfile(addProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := sccPhoenixComplianceApisService.NewGetProfileOptions(
				"testString",
				"testString",
			)

			profileResponse, response, err := sccPhoenixComplianceApisService.GetProfile(getProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
		It(`ReplaceProfileParameters request example`, func() {
			fmt.Println("\nReplaceProfileParameters() result:")
			// begin-replace_profile_parameters

			replaceProfileParametersOptions := sccPhoenixComplianceApisService.NewReplaceProfileParametersOptions(
				"testString",
				"testString",
			)

			profileDefaultParametersResponse, response, err := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptions)
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
		It(`CreateAttachment request example`, func() {
			fmt.Println("\nCreateAttachment() result:")
			// begin-create_attachment

			createAttachmentOptions := sccPhoenixComplianceApisService.NewCreateAttachmentOptions(
				"testString",
				"testString",
			)

			attachmentProfileResponse, response, err := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentProfileResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentProfileResponse).ToNot(BeNil())
		})
		It(`CheckProfileAttachmnets request example`, func() {
			fmt.Println("\nCheckProfileAttachmnets() result:")
			// begin-check_profile_attachmnets

			checkProfileAttachmnetsOptions := sccPhoenixComplianceApisService.NewCheckProfileAttachmnetsOptions(
				"testString",
				"testString",
			)

			getAllAttachmnetsForProfileRespBody, response, err := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getAllAttachmnetsForProfileRespBody, "", "  ")
			fmt.Println(string(b))

			// end-check_profile_attachmnets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllAttachmnetsForProfileRespBody).ToNot(BeNil())
		})
		It(`GetProfileAttachmnet request example`, func() {
			fmt.Println("\nGetProfileAttachmnet() result:")
			// begin-get_profile_attachmnet

			getProfileAttachmnetOptions := sccPhoenixComplianceApisService.NewGetProfileAttachmnetOptions(
				"testString",
				"testString",
				"testString",
			)

			attachmentPayload, response, err := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentPayload, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_attachmnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPayload).ToNot(BeNil())
		})
		It(`ReplaceProfileAttachment request example`, func() {
			fmt.Println("\nReplaceProfileAttachment() result:")
			// begin-replace_profile_attachment

			replaceProfileAttachmentOptions := sccPhoenixComplianceApisService.NewReplaceProfileAttachmentOptions(
				"testString",
				"testString",
				"testString",
			)

			attachmentPayload, response, err := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentPayload, "", "  ")
			fmt.Println(string(b))

			// end-replace_profile_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPayload).ToNot(BeNil())
		})
		It(`ListAttachmentParameters request example`, func() {
			fmt.Println("\nListAttachmentParameters() result:")
			// begin-list_attachment_parameters

			listAttachmentParametersOptions := sccPhoenixComplianceApisService.NewListAttachmentParametersOptions(
				"testString",
				"testString",
				"testString",
			)

			parameterDetails, response, err := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(parameterDetails, "", "  ")
			fmt.Println(string(b))

			// end-list_attachment_parameters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
		It(`ReplaceAttachment request example`, func() {
			fmt.Println("\nReplaceAttachment() result:")
			// begin-replace_attachment

			replaceAttachmentOptions := sccPhoenixComplianceApisService.NewReplaceAttachmentOptions(
				"testString",
				"testString",
				"testString",
			)

			parameterDetails, response, err := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(parameterDetails, "", "  ")
			fmt.Println(string(b))

			// end-replace_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
		It(`GetParametersByName request example`, func() {
			fmt.Println("\nGetParametersByName() result:")
			// begin-get_parameters_by_name

			getParametersByNameOptions := sccPhoenixComplianceApisService.NewGetParametersByNameOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			parameterDetails, response, err := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(parameterDetails, "", "  ")
			fmt.Println(string(b))

			// end-get_parameters_by_name

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
		It(`ReplaceAttachmnetParametersByName request example`, func() {
			fmt.Println("\nReplaceAttachmnetParametersByName() result:")
			// begin-replace_attachmnet_parameters_by_name

			replaceAttachmnetParametersByNameOptions := sccPhoenixComplianceApisService.NewReplaceAttachmnetParametersByNameOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			parameterDetails, response, err := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(parameterDetails, "", "  ")
			fmt.Println(string(b))

			// end-replace_attachmnet_parameters_by_name

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
		It(`CreateCustomControlLibrary request example`, func() {
			fmt.Println("\nCreateCustomControlLibrary() result:")
			// begin-create_custom_control_library

			createCustomControlLibraryOptions := sccPhoenixComplianceApisService.NewCreateCustomControlLibraryOptions(
				"testString",
			)

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryRequest, "", "  ")
			fmt.Println(string(b))

			// end-create_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
		It(`ListControlLibraries request example`, func() {
			fmt.Println("\nListControlLibraries() result:")
			// begin-list_control_libraries

			listControlLibrariesOptions := sccPhoenixComplianceApisService.NewListControlLibrariesOptions(
				"testString",
			)

			getAllControlLibrariesRespBody, response, err := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getAllControlLibrariesRespBody, "", "  ")
			fmt.Println(string(b))

			// end-list_control_libraries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllControlLibrariesRespBody).ToNot(BeNil())
		})
		It(`ReplaceCustomControlLibrary request example`, func() {
			fmt.Println("\nReplaceCustomControlLibrary() result:")
			// begin-replace_custom_control_library

			replaceCustomControlLibraryOptions := sccPhoenixComplianceApisService.NewReplaceCustomControlLibraryOptions(
				"testString",
				"testString",
			)

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryRequest, "", "  ")
			fmt.Println(string(b))

			// end-replace_custom_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
		It(`GetControlLibrary request example`, func() {
			fmt.Println("\nGetControlLibrary() result:")
			// begin-get_control_library

			getControlLibraryOptions := sccPhoenixComplianceApisService.NewGetControlLibraryOptions(
				"testString",
				"testString",
			)

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryRequest, "", "  ")
			fmt.Println(string(b))

			// end-get_control_library

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
		It(`CreateScan request example`, func() {
			fmt.Println("\nCreateScan() result:")
			// begin-create_scan

			createScanOptions := sccPhoenixComplianceApisService.NewCreateScanOptions(
				"testString",
			)

			createScanResponse, response, err := sccPhoenixComplianceApisService.CreateScan(createScanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createScanResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_scan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createScanResponse).ToNot(BeNil())
		})
		It(`DeleteCustomProfile request example`, func() {
			fmt.Println("\nDeleteCustomProfile() result:")
			// begin-delete_custom_profile

			deleteCustomProfileOptions := sccPhoenixComplianceApisService.NewDeleteCustomProfileOptions(
				"testString",
				"testString",
			)

			profileResponse, response, err := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
		It(`DeleteProfileAttachmnet request example`, func() {
			fmt.Println("\nDeleteProfileAttachmnet() result:")
			// begin-delete_profile_attachmnet

			deleteProfileAttachmnetOptions := sccPhoenixComplianceApisService.NewDeleteProfileAttachmnetOptions(
				"testString",
				"testString",
				"testString",
			)

			attachmentPayload, response, err := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(attachmentPayload, "", "  ")
			fmt.Println(string(b))

			// end-delete_profile_attachmnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPayload).ToNot(BeNil())
		})
		It(`DeleteCustomControllibrary request example`, func() {
			fmt.Println("\nDeleteCustomControllibrary() result:")
			// begin-delete_custom_controllibrary

			deleteCustomControllibraryOptions := sccPhoenixComplianceApisService.NewDeleteCustomControllibraryOptions(
				"testString",
				"testString",
			)

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlLibraryRequest, "", "  ")
			fmt.Println(string(b))

			// end-delete_custom_controllibrary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
	})
})

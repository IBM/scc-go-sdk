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

package sccphoenixcomplianceapisv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/sccphoenixcomplianceapisv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the sccphoenixcomplianceapisv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SccPhoenixComplianceApisV1 Integration Tests`, func() {
	const externalConfigFile = "scc_phoenix_compliance_apis_v1.env"

	var (
		err                             error
		sccPhoenixComplianceApisService *sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1
		serviceURL                      string
		config                          map[string]string
		authenticator                   core.IamAuthenticator
		accountID                       string
		authUrl                         string
		apiKey                          string
		instanceID                      string
		transactionID                   string
		serviceName                     string
		parameterName                   string
		parameterDefaultValue           string
		parameterDisplayName            string
		parameterType                   string
		parameterValue                  string
		profileName                     string
		profileDescription              string
		profileType                     string
		profileVersion                  string
		profilesID                      string
		profilesIdDelete                string
		scopeID                         string
		scopeType                       string
		createdBy                       string
		createdOn                       string
		updatedBy                       string
		updatedOn                       string
		attachmentID                    string
		attachmentIdDelete              string
		newParameterName                string
		newParameterDisplayName         string
		newParameterType                string
		newParameterValue               string
		newAssessmentType               string
		newAssessmentID                 string
		assessmentMethod                string
		assessmentDescription           string
		assessmentType                  string
		assessmentID                    string
		responsibility                  string
		componentID                     string
		environment                     string
		description                     string
		controlName                     string
		controlID                       string
		controlDescription              string
		controlCategory                 string
		controlLibraryID                string
		controlLibraryName              string
		controlLibraryDescription       string
		controlLibraryType              string
		controlLibraryVersion           string
		controlLibrariesID              string
		controlLibrariesIdDelete        string
		assessmentTypeParameter         string
		assessmentIdParameter           string
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
			config, err = core.GetServiceProperties(sccphoenixcomplianceapisv1.DefaultServiceName)
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

			accountID = config["ACCOUNT_ID"]

			instanceID = config["instanceID"]
			if instanceID == "" {
				Skip("Unable to load instanceID configuration property, skipping tests")
			}

			transactionID = config["transactionID"]
			if transactionID == "" {
				Skip("Unable to load transactionID configuration property, skipping tests")
			}

			controlLibraryID = config["controlLibraryID"]
			controlID = config["controlID"]
			parameterName = config["parameterName"]
			parameterDefaultValue = config["parameterDefaultValue"]
			parameterDisplayName = config["parameterDisplayName"]
			parameterType = config["parameterType"]
			parameterValue = config["parameterValue"]
			profileName = config["profileName"]
			profileDescription = config["profileDescription"]
			profileType = config["profileType"]
			profileVersion = config["profileVersion"]
			profilesID = config["profilesID"]
			profilesIdDelete = config["profilesIdDelete"]
			scopeID = config["scopeID"]
			scopeType = config["scopeType"]
			createdBy = config["createdBy"]
			createdOn = config["createdOn"]
			updatedBy = config["updatedBy"]
			updatedOn = config["updatedOn"]
			attachmentID = config["attachmentID"]
			attachmentIdDelete = config["attachmentIdDelete"]
			newParameterName = config["newParameterName"]
			newParameterDisplayName = config["newParameterDisplayName"]
			newParameterType = config["newParameterType"]
			newParameterValue = config["newParameterValue"]
			newAssessmentType = config["newAssessmentType"]
			newAssessmentID = config["newAssessmentID"]
			assessmentMethod = config["assessmentMethod"]
			assessmentDescription = config["assessmentDescription"]
			assessmentType = config["assessmentType"]
			assessmentID = config["assessmentID"]
			responsibility = config["responsibility"]
			componentID = config["componentID"]
			environment = config["environment"]
			description = config["description"]
			controlName = config["controlName"]
			controlDescription = config["controlDescription"]
			controlCategory = config["controlCategory"]
			controlLibraryName = config["controlLibraryName"]
			controlLibraryDescription = config["controlLibraryDescription"]
			controlLibraryType = config["controlLibraryType"]
			controlLibraryVersion = config["controlLibraryVersion"]
			controlLibrariesID = config["controlLibrariesID"]
			controlLibrariesIdDelete = config["controlLibrariesIdDelete"]
			assessmentTypeParameter = config["assessmentTypeParameter"]
			assessmentIdParameter = config["assessmentIdParameter"]

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
			sccPhoenixComplianceApisServiceOptions := &sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				URL:           serviceURL,
				Authenticator: &authenticator,
				ServiceName:   serviceName,
			}

			sccPhoenixComplianceApisService, err = sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(sccPhoenixComplianceApisServiceOptions)
			Expect(err).To(BeNil())
			Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
			Expect(sccPhoenixComplianceApisService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			sccPhoenixComplianceApisService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateProfile - Create a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
			profileControlsInRequestModel := &sccphoenixcomplianceapisv1.ProfileControlsInRequest{
				ControlLibraryID: core.StringPtr(controlLibraryID),
				ControlID:        core.StringPtr(controlID),
			}

			// defaultParametersModel := &sccphoenixcomplianceapisv1.DefaultParameters{
			// 	AssessmentType:        core.StringPtr("testString"),
			// 	AssessmentID:          core.StringPtr("testString"),
			// 	ParameterName:         core.StringPtr("testString"),
			// 	ParameterDefaultValue: core.StringPtr("testString"),
			// 	ParameterDisplayName:  core.StringPtr("testString"),
			// 	ParameterType:         core.StringPtr("numeric"),
			// }

			createProfileOptions := &sccphoenixcomplianceapisv1.CreateProfileOptions{
				InstanceID:         core.StringPtr(instanceID),
				ProfileName:        core.StringPtr(profileName),
				ProfileDescription: core.StringPtr(profileDescription),
				ProfileType:        core.StringPtr(profileType),
				ProfileVersion:     core.StringPtr(profileVersion),
				Latest:             core.BoolPtr(true),
				Controls:           []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel},
				DefaultParameters:  []sccphoenixcomplianceapisv1.DefaultParameters{},
				TransactionID:      core.StringPtr(transactionID),
			}

			_, _, err := sccPhoenixComplianceApisService.CreateProfile(createProfileOptions)
			Expect(err).To(BeNil())
			//Expect(response.StatusCode).To(Equal(500))
			// Expect(profileResponse).ToNot(BeNil())
		})
	})

	Describe(`ListProfiles - Get all predefined and user's custom profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {
			listProfilesOptions := &sccphoenixcomplianceapisv1.ListProfilesOptions{
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			getAllProfilesRespBody, response, err := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptions)
			fmt.Println(response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllProfilesRespBody).ToNot(BeNil())
		})
	})

	Describe(`AddProfile - Update a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddProfile(addProfileOptions *AddProfileOptions)`, func() {
			profileControlsInRequestModel := &sccphoenixcomplianceapisv1.ProfileControlsInRequest{
				ControlLibraryID: core.StringPtr(controlLibraryID),
				ControlID:        core.StringPtr(controlID),
			}

			// defaultParametersModel := &sccphoenixcomplianceapisv1.DefaultParameters{
			// 	AssessmentType:        core.StringPtr("testString"),
			// 	AssessmentID:          core.StringPtr("testString"),
			// 	ParameterName:         core.StringPtr("testString"),
			// 	ParameterDefaultValue: core.StringPtr("testString"),
			// 	ParameterDisplayName:  core.StringPtr("testString"),
			// 	ParameterType:         core.StringPtr("numeric"),
			// }

			addProfileOptions := &sccphoenixcomplianceapisv1.AddProfileOptions{
				ProfilesID:         core.StringPtr(profilesID),
				InstanceID:         core.StringPtr(instanceID),
				ProfileName:        core.StringPtr(profileName),
				ProfileDescription: core.StringPtr(profileDescription),
				ProfileType:        core.StringPtr(profileType),
				ProfileVersion:     core.StringPtr(profileVersion),
				Latest:             core.BoolPtr(true),
				Controls:           []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel},
				DefaultParameters:  []sccphoenixcomplianceapisv1.DefaultParameters{},
				TransactionID:      core.StringPtr(transactionID),
			}

			profileResponse, response, err := sccPhoenixComplianceApisService.AddProfile(addProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
	})

	Describe(`GetProfile - Get a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
			getProfileOptions := &sccphoenixcomplianceapisv1.GetProfileOptions{
				ProfilesID:    core.StringPtr(profilesID),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			profileResponse, response, err := sccPhoenixComplianceApisService.GetProfile(getProfileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileResponse).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileParameters - Update custom profile parameters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions)`, func() {
			defaultParametersModel := &sccphoenixcomplianceapisv1.DefaultParameters{
				AssessmentType:        core.StringPtr(assessmentType),
				AssessmentID:          core.StringPtr(assessmentID),
				ParameterName:         core.StringPtr(parameterName),
				ParameterDefaultValue: core.StringPtr(parameterDefaultValue),
				ParameterDisplayName:  core.StringPtr(parameterDisplayName),
				ParameterType:         core.StringPtr(parameterType),
			}

			replaceProfileParametersOptions := &sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions{
				ProfilesID:        core.StringPtr(profilesID),
				InstanceID:        core.StringPtr(instanceID),
				DefaultParameters: []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel},
				TransactionID:     core.StringPtr("testString"),
			}

			profileDefaultParametersResponse, response, err := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileDefaultParametersResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateAttachment - Create an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAttachment(createAttachmentOptions *CreateAttachmentOptions)`, func() {
			scopePayloadModel := &sccphoenixcomplianceapisv1.ScopePayload{
				ScopeID:   core.StringPtr(scopeID),
				ScopeType: core.StringPtr(scopeType),
			}

			// parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
			// 	ParameterName:        core.StringPtr(parameterName),
			// 	ParameterDisplayName: core.StringPtr(parameterDisplayName),
			// 	ParameterType:        core.StringPtr(parameterType),
			// }

			parameterDetailsModel := &sccphoenixcomplianceapisv1.ParameterDetails{
				ParameterName:        core.StringPtr(parameterName),
				ParameterDisplayName: core.StringPtr(parameterDisplayName),
				ParameterType:        core.StringPtr(parameterType),
				ParameterValue:       core.StringPtr(parameterValue),
				AssessmentType:       core.StringPtr(assessmentTypeParameter),
				AssessmentID:         core.StringPtr(assessmentIdParameter),
				// Parameters:           []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel},
			}

			// failedControlsModel := &sccphoenixcomplianceapisv1.FailedControls{
			// 	ThresholdLimit:   core.Int64Ptr(int64(38)),
			// 	FailedControlIds: []string{"testString"},
			// }

			// attachmentsNotificationsPayloadModel := &sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload{
			// 	Enabled:  core.BoolPtr(true),
			// 	Controls: failedControlsModel,
			// }

			attachmentPayloadModel := &sccphoenixcomplianceapisv1.AttachmentPayload{
				AccountID:            core.StringPtr(accountID),
				IncludedScope:        scopePayloadModel,
				Exclusions:           []sccphoenixcomplianceapisv1.ScopePayload{},
				CreatedBy:            core.StringPtr(createdBy),
				CreatedOn:            core.StringPtr(createdOn),
				UpdatedBy:            core.StringPtr(updatedBy),
				UpdatedOn:            core.StringPtr(updatedOn),
				Status:               core.StringPtr("enabled"),
				AttachmentParameters: []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel},
			}

			createAttachmentOptions := &sccphoenixcomplianceapisv1.CreateAttachmentOptions{
				ProfilesID:    core.StringPtr(profilesID),
				InstanceID:    core.StringPtr(instanceID),
				Attachments:   []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel},
				TransactionID: core.StringPtr(transactionID),
			}

			attachmentProfileRequest, response, err := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentProfileRequest).ToNot(BeNil())
		})
	})

	Describe(`CheckProfileAttachmnets - Get all attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckProfileAttachmnets(checkProfileAttachmnetsOptions *CheckProfileAttachmnetsOptions)`, func() {
			checkProfileAttachmnetsOptions := &sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions{
				ProfilesID:    core.StringPtr(profilesID),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			getAllAttachmnetsForProfileRespBody, response, err := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllAttachmnetsForProfileRespBody).ToNot(BeNil())
		})
	})

	Describe(`GetProfileAttachmnet - Get an attachment for a profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileAttachmnet(getProfileAttachmnetOptions *GetProfileAttachmnetOptions)`, func() {
			getProfileAttachmnetOptions := &sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions{
				ProfilesID:    core.StringPtr(profilesID),
				AttachmentID:  core.StringPtr(attachmentID),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			attachmentProfileRequest, response, err := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentProfileRequest).ToNot(BeNil())
		})
	})

	Describe(`ReplaceProfileAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
			scopePayloadModel := &sccphoenixcomplianceapisv1.ScopePayload{
				ScopeID:   core.StringPtr(scopeID),
				ScopeType: core.StringPtr(scopeType),
			}

			// parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
			// 	ParameterName:        core.StringPtr("testString"),
			// 	ParameterDisplayName: core.StringPtr("testString"),
			// 	ParameterType:        core.StringPtr("numeric"),
			// }

			parameterDetailsModel := &sccphoenixcomplianceapisv1.ParameterDetails{
				ParameterName:        core.StringPtr(parameterName),
				ParameterDisplayName: core.StringPtr(parameterDisplayName),
				ParameterType:        core.StringPtr(parameterType),
				ParameterValue:       core.StringPtr(parameterValue),
				AssessmentType:       core.StringPtr(assessmentTypeParameter),
				AssessmentID:         core.StringPtr(assessmentIdParameter),
				// Parameters:           []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel},
			}

			// failedControlsModel := &sccphoenixcomplianceapisv1.FailedControls{
			// 	ThresholdLimit:   core.Int64Ptr(int64(38)),
			// 	FailedControlIds: []string{"testString"},
			// }

			// attachmentsNotificationsPayloadModel := &sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload{
			// 	Enabled:  core.BoolPtr(true),
			// 	Controls: failedControlsModel,
			// }

			replaceProfileAttachmentOptions := &sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions{
				ProfilesID:           core.StringPtr(profilesID),
				AttachmentID:         core.StringPtr(attachmentID),
				InstanceID:           core.StringPtr(instanceID),
				ID:                   core.StringPtr(attachmentID),
				AccountID:            core.StringPtr(attachmentID),
				IncludedScope:        scopePayloadModel,
				Exclusions:           []sccphoenixcomplianceapisv1.ScopePayload{},
				CreatedBy:            core.StringPtr(createdBy),
				CreatedOn:            core.StringPtr(createdOn),
				UpdatedBy:            core.StringPtr(updatedBy),
				UpdatedOn:            core.StringPtr(updatedOn),
				Status:               core.StringPtr("enabled"),
				AttachmentParameters: []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel},
				// AttachmentNotifications: attachmentsNotificationsPayloadModel,
				TransactionID: core.StringPtr(transactionID),
			}

			attachmentPayload, response, err := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(attachmentPayload).ToNot(BeNil())
		})
	})

	Describe(`ListAttachmentParameters - Get attachment's parameters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAttachmentParameters(listAttachmentParametersOptions *ListAttachmentParametersOptions)`, func() {
			listAttachmentParametersOptions := &sccphoenixcomplianceapisv1.ListAttachmentParametersOptions{
				ProfilesID:    core.StringPtr(profilesID),
				AttachmentID:  core.StringPtr(attachmentID),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			_, response, _ := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptions)
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ReplaceAttachment - Update parameters for an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceAttachment(replaceAttachmentOptions *ReplaceAttachmentOptions)`, func() {
			parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
				ParameterName:        core.StringPtr(parameterName),
				ParameterDisplayName: core.StringPtr(parameterDisplayName),
				ParameterType:        core.StringPtr(parameterType),
			}

			replaceAttachmentOptions := &sccphoenixcomplianceapisv1.ReplaceAttachmentOptions{
				ProfilesID:           core.StringPtr(profilesID),
				AttachmentID:         core.StringPtr(attachmentID),
				InstanceID:           core.StringPtr(instanceID),
				ParameterName:        core.StringPtr(parameterName),
				ParameterDisplayName: core.StringPtr("dylannn display name"),
				ParameterType:        core.StringPtr(parameterType),
				ParameterValue:       core.StringPtr(parameterValue),
				AssessmentType:       core.StringPtr(assessmentType),
				AssessmentID:         core.StringPtr(assessmentID),
				Parameters:           []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel},
				TransactionID:        core.StringPtr(transactionID),
			}

			_, response, _ := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptions)
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetParametersByName - Get parameters by name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetParametersByName(getParametersByNameOptions *GetParametersByNameOptions)`, func() {
			getParametersByNameOptions := &sccphoenixcomplianceapisv1.GetParametersByNameOptions{
				ProfilesID:    core.StringPtr(profilesID),
				AttachmentID:  core.StringPtr(attachmentID),
				ParameterName: core.StringPtr(parameterName),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			parameterDetails, response, err := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
	})

	Describe(`ReplaceAttachmnetParametersByName - Update parameter by name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions *ReplaceAttachmnetParametersByNameOptions)`, func() {
			parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
				ParameterName:        core.StringPtr(parameterName),
				ParameterDisplayName: core.StringPtr(parameterDisplayName),
				ParameterType:        core.StringPtr(parameterType),
			}

			replaceAttachmnetParametersByNameOptions := &sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions{
				ProfilesID:              core.StringPtr(profilesID),
				AttachmentID:            core.StringPtr(attachmentID),
				ParameterName:           core.StringPtr(parameterName),
				InstanceID:              core.StringPtr(instanceID),
				NewParameterName:        core.StringPtr(newParameterName),
				NewParameterDisplayName: core.StringPtr(newParameterDisplayName),
				NewParameterType:        core.StringPtr(newParameterType),
				NewParameterValue:       core.StringPtr(newParameterValue),
				NewAssessmentType:       core.StringPtr(newAssessmentType),
				NewAssessmentID:         core.StringPtr(newAssessmentID),
				NewParameters:           []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel},
				TransactionID:           core.StringPtr(transactionID),
			}

			parameterDetails, response, err := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(parameterDetails).ToNot(BeNil())
		})
	})

	Describe(`CreateCustomControlLibrary - Create a custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions)`, func() {
			// parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
			// 	ParameterName:        core.StringPtr("testString"),
			// 	ParameterDisplayName: core.StringPtr("testString"),
			// 	ParameterType:        core.StringPtr("numeric"),
			// }

			implementationPayloadModel := &sccphoenixcomplianceapisv1.ImplementationPayload{
				AssessmentID:          core.StringPtr(assessmentID),
				AssessmentMethod:      core.StringPtr(assessmentMethod),
				AssessmentType:        core.StringPtr(assessmentType),
				AssessmentDescription: core.StringPtr(assessmentDescription),
				Parameters:            []sccphoenixcomplianceapisv1.ParameterInfo{},
			}

			controlSpecificationsModel := &sccphoenixcomplianceapisv1.ControlSpecifications{
				ID:             core.StringPtr(controlID),
				Responsibility: core.StringPtr(responsibility),
				ComponentID:    core.StringPtr(componentID),
				Environment:    core.StringPtr(environment),
				Description:    core.StringPtr(description),
				Assessments:    []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel},
			}

			controlsInControlLibRequestPayloadModel := &sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{
				ControlName:           core.StringPtr(controlName),
				ControlDescription:    core.StringPtr(controlDescription),
				ControlCategory:       core.StringPtr(controlCategory),
				ControlSpecifications: []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel},
			}

			createCustomControlLibraryOptions := &sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions{
				InstanceID:                core.StringPtr(instanceID),
				AccountID:                 core.StringPtr(accountID),
				ControlLibraryName:        core.StringPtr(controlLibraryName),
				ControlLibraryDescription: core.StringPtr(controlLibraryDescription),
				ControlLibraryType:        core.StringPtr(controlLibraryType),
				ControlLibraryVersion:     core.StringPtr(controlLibraryVersion),
				Latest:                    core.BoolPtr(true),
				Controls:                  []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel},
				TransactionID:             core.StringPtr(transactionID),
			}

			_, response, _ := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptions)
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ListControlLibraries - Get all control libraries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions)`, func() {
			listControlLibrariesOptions := &sccphoenixcomplianceapisv1.ListControlLibrariesOptions{
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			_, response, err := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ReplaceCustomControlLibrary - Update custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions)`, func() {
			// parameterInfoModel := &sccphoenixcomplianceapisv1.ParameterInfo{
			// 	ParameterName:        core.StringPtr(parameter),
			// 	ParameterDisplayName: core.StringPtr("testString"),
			// 	ParameterType:        core.StringPtr("numeric"),
			// }

			implementationPayloadModel := &sccphoenixcomplianceapisv1.ImplementationPayload{
				AssessmentID:          core.StringPtr(assessmentID),
				AssessmentMethod:      core.StringPtr(assessmentMethod),
				AssessmentType:        core.StringPtr(assessmentType),
				AssessmentDescription: core.StringPtr(assessmentDescription),
				Parameters:            []sccphoenixcomplianceapisv1.ParameterInfo{},
			}

			controlSpecificationsModel := &sccphoenixcomplianceapisv1.ControlSpecifications{
				ID:             core.StringPtr(controlID),
				Responsibility: core.StringPtr(responsibility),
				ComponentID:    core.StringPtr(componentID),
				Environment:    core.StringPtr(environment),
				Description:    core.StringPtr(description),
				Assessments:    []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel},
			}

			// controlDocsModel := &sccphoenixcomplianceapisv1.ControlDocs{
			// 	ControlDocsID:   core.StringPtr("testString"),
			// 	ControlDocsType: core.StringPtr("testString"),
			// }

			controlsInControlLibRequestPayloadModel := &sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{
				ControlName:           core.StringPtr(controlName),
				ControlDescription:    core.StringPtr(controlDescription),
				ControlCategory:       core.StringPtr(controlCategory),
				ControlSpecifications: []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel},
			}

			replaceCustomControlLibraryOptions := &sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions{
				ControlLibrariesID:        core.StringPtr(controlLibrariesID),
				InstanceID:                core.StringPtr(instanceID),
				AccountID:                 core.StringPtr(accountID),
				ControlLibraryName:        core.StringPtr(controlLibraryName),
				ControlLibraryDescription: core.StringPtr(controlLibraryDescription),
				ControlLibraryType:        core.StringPtr(controlLibraryType),
				ControlLibraryVersion:     core.StringPtr(controlLibraryVersion),
				Latest:                    core.BoolPtr(true),
				Controls:                  []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel},
				TransactionID:             core.StringPtr(transactionID),
			}

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
	})

	Describe(`GetControlLibrary - Get control library by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions)`, func() {
			getControlLibraryOptions := &sccphoenixcomplianceapisv1.GetControlLibraryOptions{
				ControlLibrariesID: core.StringPtr(controlLibrariesID),
				InstanceID:         core.StringPtr(instanceID),
				TransactionID:      core.StringPtr(transactionID),
			}

			controlLibraryRequest, response, err := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlLibraryRequest).ToNot(BeNil())
		})
	})

	Describe(`CreateScan - Create a scan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
			createScanOptions := &sccphoenixcomplianceapisv1.CreateScanOptions{
				InstanceID:    core.StringPtr(instanceID),
				AttachmentID:  core.StringPtr(attachmentID),
				TransactionID: core.StringPtr(transactionID),
			}

			createScanResponse, response, err := sccPhoenixComplianceApisService.CreateScan(createScanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createScanResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteCustomProfile - Delete a custom profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
			deleteCustomProfileOptions := &sccphoenixcomplianceapisv1.DeleteCustomProfileOptions{
				ProfilesID:    core.StringPtr(profilesIdDelete),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			_, _, err := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptions)
			Expect(err).To(BeNil())
		})
	})

	Describe(`DeleteProfileAttachmnet - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileAttachmnet(deleteProfileAttachmnetOptions *DeleteProfileAttachmnetOptions)`, func() {
			deleteProfileAttachmnetOptions := &sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions{
				ProfilesID:    core.StringPtr(profilesID),
				AttachmentID:  core.StringPtr(attachmentIdDelete),
				InstanceID:    core.StringPtr(instanceID),
				TransactionID: core.StringPtr(transactionID),
			}

			_, _, err := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptions)
			Expect(err).To(BeNil())
		})
	})

	Describe(`DeleteCustomControllibrary - Delete custom control library`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCustomControllibrary(deleteCustomControllibraryOptions *DeleteCustomControllibraryOptions)`, func() {
			deleteCustomControllibraryOptions := &sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions{
				ControlLibrariesID: core.StringPtr(controlLibrariesIdDelete),
				InstanceID:         core.StringPtr(instanceID),
				TransactionID:      core.StringPtr(transactionID),
			}

			_, response, _ := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptions)

			Expect(response.StatusCode).To(Equal(200))
		})
	})
})

//
// Utility functions are declared in the unit test file
//

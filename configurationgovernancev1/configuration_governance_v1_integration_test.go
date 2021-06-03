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

package configurationgovernancev1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm/scc-go-sdk/configurationgovernancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the configurationgovernancev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ConfigurationGovernanceV1 Integration Tests`, func() {

	const externalConfigFile = "../configuration_governance_v1.env"

	var (
		err                            error
		configurationGovernanceService *configurationgovernancev1.ConfigurationGovernanceV1
		serviceURL                     string
		config                         map[string]string
	)

	// Globlal variables to hold link values
	var (
		ruleAttachmentIDLink string
		ruleIDLink           string
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
			config, err = core.GetServiceProperties(configurationgovernancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			configurationGovernanceServiceOptions := &configurationgovernancev1.ConfigurationGovernanceV1Options{}

			configurationGovernanceService, err = configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

			Expect(err).To(BeNil())
			Expect(configurationGovernanceService).ToNot(BeNil())
			Expect(configurationGovernanceService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateRules - Create rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRules(createRulesOptions *CreateRulesOptions)`, func() {

			targetResourceAdditionalTargetAttributesItemModel := &configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{
				Name:     core.StringPtr("resource_id"),
				Value:    core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf"),
				Operator: core.StringPtr("string_equals"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName:                core.StringPtr("iam-groups"),
				ResourceKind:               core.StringPtr("service"),
				AdditionalTargetAttributes: []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel},
			}

			ruleConditionModel := &configurationgovernancev1.RuleConditionSingleProperty{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("public_access_enabled"),
				Operator:    core.StringPtr("is_false"),
				Value:       core.StringPtr("testString"),
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("Public access check"),
				And:         []configurationgovernancev1.RuleConditionIntf{ruleConditionModel},
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			ruleRequestModel := &configurationgovernancev1.RuleRequest{
				AccountID:          core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				Name:               core.StringPtr("Disable public access"),
				Description:        core.StringPtr("Ensure that public access to account resources is disabled."),
				RuleType:           core.StringPtr("user_defined"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				Labels:             []string{"Access", "IAM"},
			}

			createRuleRequestModel := &configurationgovernancev1.CreateRuleRequest{
				RequestID: core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558"),
				Rule:      ruleRequestModel,
			}

			createRulesOptions := &configurationgovernancev1.CreateRulesOptions{
				Rules:         []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createRulesResponse, response, err := configurationGovernanceService.CreateRules(createRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())

			ruleIDLink = *createRulesResponse.Rules[0].Rule.RuleID

		})
	})

	Describe(`CreateRuleAttachments - Create attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRuleAttachments(createRuleAttachmentsOptions *CreateRuleAttachmentsOptions)`, func() {

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My enterprise"),
				ScopeID:   core.StringPtr("282cf433ac91493ba860480d92519990"),
				ScopeType: core.StringPtr("enterprise"),
			}

			ruleAttachmentRequestModel := &configurationgovernancev1.RuleAttachmentRequest{
				AccountID:      core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				IncludedScope:  ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*ruleScopeModel},
			}

			createRuleAttachmentsOptions := &configurationgovernancev1.CreateRuleAttachmentsOptions{
				RuleID:        &ruleIDLink,
				Attachments:   []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createRuleAttachmentsResponse, response, err := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRuleAttachmentsResponse).ToNot(BeNil())

			ruleAttachmentIDLink = *createRuleAttachmentsResponse.Attachments[0].AttachmentID

		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {

			listRulesOptions := &configurationgovernancev1.ListRulesOptions{
				AccountID:     core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				TransactionID: core.StringPtr("testString"),
				Attached:      core.BoolPtr(true),
				Labels:        core.StringPtr("SOC2,ITCS300"),
				Scopes:        core.StringPtr("scope_id"),
				Limit:         core.Int64Ptr(int64(1000)),
				Offset:        core.Int64Ptr(int64(38)),
			}

			ruleList, response, err := configurationGovernanceService.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())

		})
	})

	Describe(`GetRule - Get a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {

			getRuleOptions := &configurationgovernancev1.GetRuleOptions{
				RuleID:        &ruleIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			rule, response, err := configurationGovernanceService.GetRule(getRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
	})

	Describe(`UpdateRule - Update a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRule(updateRuleOptions *UpdateRuleOptions)`, func() {

			targetResourceAdditionalTargetAttributesItemModel := &configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{
				Name:     core.StringPtr("testString"),
				Value:    core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName:                core.StringPtr("iam-groups"),
				ResourceKind:               core.StringPtr("service"),
				AdditionalTargetAttributes: []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel},
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigSingleProperty{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("public_access_enabled"),
				Operator:    core.StringPtr("is_false"),
				Value:       core.StringPtr("testString"),
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("audit_log"),
			}

			updateRuleOptions := &configurationgovernancev1.UpdateRuleOptions{
				RuleID:             &ruleIDLink,
				IfMatch:            core.StringPtr("testString"),
				Name:               core.StringPtr("Disable public access"),
				Description:        core.StringPtr("Ensure that public access to account resources is disabled."),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				AccountID:          core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				RuleType:           core.StringPtr("user_defined"),
				Labels:             []string{"SOC2", "ITCS300"},
				TransactionID:      core.StringPtr("testString"),
			}

			rule, response, err := configurationGovernanceService.UpdateRule(updateRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
	})

	Describe(`ListRuleAttachments - List attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRuleAttachments(listRuleAttachmentsOptions *ListRuleAttachmentsOptions)`, func() {

			listRuleAttachmentsOptions := &configurationgovernancev1.ListRuleAttachmentsOptions{
				RuleID:        &ruleIDLink,
				TransactionID: core.StringPtr("testString"),
				Limit:         core.Int64Ptr(int64(1000)),
				Offset:        core.Int64Ptr(int64(38)),
			}

			ruleAttachmentList, response, err := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleAttachmentList).ToNot(BeNil())

		})
	})

	Describe(`GetRuleAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRuleAttachment(getRuleAttachmentOptions *GetRuleAttachmentOptions)`, func() {

			getRuleAttachmentOptions := &configurationgovernancev1.GetRuleAttachmentOptions{
				RuleID:        &ruleIDLink,
				AttachmentID:  &ruleAttachmentIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			ruleAttachment, response, err := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleAttachment).ToNot(BeNil())

		})
	})

	Describe(`UpdateRuleAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRuleAttachment(updateRuleAttachmentOptions *UpdateRuleAttachmentOptions)`, func() {

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My enterprise"),
				ScopeID:   core.StringPtr("282cf433ac91493ba860480d92519990"),
				ScopeType: core.StringPtr("enterprise"),
			}

			updateRuleAttachmentOptions := &configurationgovernancev1.UpdateRuleAttachmentOptions{
				RuleID:         &ruleIDLink,
				AttachmentID:   &ruleAttachmentIDLink,
				IfMatch:        core.StringPtr("testString"),
				AccountID:      core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				IncludedScope:  ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*ruleScopeModel},
				TransactionID:  core.StringPtr("testString"),
			}

			templateAttachment, response, err := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

		})
	})

	Describe(`DeleteRuleAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRuleAttachment(deleteRuleAttachmentOptions *DeleteRuleAttachmentOptions)`, func() {

			deleteRuleAttachmentOptions := &configurationgovernancev1.DeleteRuleAttachmentOptions{
				RuleID:        &ruleIDLink,
				AttachmentID:  &ruleAttachmentIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteRuleAttachment(deleteRuleAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteRule - Delete a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {

			deleteRuleOptions := &configurationgovernancev1.DeleteRuleOptions{
				RuleID:        &ruleIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//

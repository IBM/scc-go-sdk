//go:build integration
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
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
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

var accountID = os.Getenv("ACCOUNT_ID")
var ruleLabel = os.Getenv("RULE_LABEL")
var resourceGroupID = os.Getenv("RESOURCE_GROUP_ID")
var identifier = fmt.Sprintf("go-%d", time.Now().Unix())

var _ = Describe(`ConfigurationGovernanceV1 Integration Tests`, func() {

	const externalConfigFile = "../configuration_governance_v1.env"

	var (
		err                            error
		configurationGovernanceService *configurationgovernancev1.ConfigurationGovernanceV1
		serviceURL                     string
		config                         map[string]string
	)

	// Global variables to hold link values
	var (
		ruleAttachmentIDLink     string
		ruleIDLink               string
		ruleEtag                 string
		ruleAttachmentEtag       string
		templateAttachmentIDLink string
		templateAttachmentEtag   string
		templateIDLink           string
		templateEtag             string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {

		if ruleLabel == "" {
			ruleLabel = "sdk-it"
		}

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
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel},
			}

			ruleConditionModel := &configurationgovernancev1.RuleConditionSingleProperty{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("location"),
				Operator:    core.StringPtr("string_equals"),
				Value:       core.StringPtr("us-south"),
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("Cloud Object Storage bucket"),
				And:         []configurationgovernancev1.RuleConditionIntf{ruleConditionModel},
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			ruleRequestModel := &configurationgovernancev1.RuleRequest{
				AccountID:          core.StringPtr(accountID),
				Name:               core.StringPtr("Disable public access in Dallas"),
				Description:        core.StringPtr("Ensure that public access to buckets in us-south is disabled."),
				RuleType:           core.StringPtr("user_defined"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				Labels:             []string{ruleLabel},
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
				Note:      core.StringPtr("My account"),
				ScopeID:   core.StringPtr(accountID),
				ScopeType: core.StringPtr("account"),
			}

			excludedScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My account resource group"),
				ScopeID:   core.StringPtr(resourceGroupID),
				ScopeType: core.StringPtr("account.resource_group"),
			}

			ruleAttachmentRequestModel := &configurationgovernancev1.RuleAttachmentRequest{
				AccountID:      core.StringPtr(accountID),
				IncludedScope:  ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*excludedScopeModel},
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

	Describe(`CreateTemplates - Create templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTemplates(createTemplatesOptions *CreateTemplatesOptions)`, func() {

			baseTargetAttributeModel := &configurationgovernancev1.BaseTargetAttribute{
				Name:  core.StringPtr("location"),
				Value: core.StringPtr("us-south"),
			}

			simpleTargetResourceModel := &configurationgovernancev1.SimpleTargetResource{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configurationgovernancev1.BaseTargetAttribute{*baseTargetAttributeModel},
			}

			templateCustomizedDefaultPropertyModel := &configurationgovernancev1.TemplateCustomizedDefaultProperty{
				Property: core.StringPtr("location"),
				Value:    core.StringPtr("level"),
			}

			templateModel := &configurationgovernancev1.Template{
				AccountID:          core.StringPtr(accountID),
				Name:               core.StringPtr("testString"),
				Description:        core.StringPtr("testString"),
				Target:             simpleTargetResourceModel,
				CustomizedDefaults: []configurationgovernancev1.TemplateCustomizedDefaultProperty{*templateCustomizedDefaultPropertyModel},
			}

			createTemplateRequestModel := &configurationgovernancev1.CreateTemplateRequest{
				RequestID: core.StringPtr("testString"),
				Template:  templateModel,
			}

			createTemplatesOptions := &configurationgovernancev1.CreateTemplatesOptions{
				Templates:     []configurationgovernancev1.CreateTemplateRequest{*createTemplateRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createTemplatesResponse, response, err := configurationGovernanceService.CreateTemplates(createTemplatesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplatesResponse).ToNot(BeNil())

			templateIDLink = *createTemplatesResponse.Templates[0].Template.TemplateID
		})
	})

	Describe(`CreateTemplateAttachments - Create attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTemplateAttachments(createTemplateAttachmentsOptions *CreateTemplateAttachmentsOptions)`, func() {

			templateScopeModel := &configurationgovernancev1.TemplateScope{
				Note:      core.StringPtr("My account"),
				ScopeID:   core.StringPtr(accountID),
				ScopeType: core.StringPtr("account"),
			}

			excludedScopeModel := &configurationgovernancev1.TemplateScope{
				Note:      core.StringPtr("My account resource group"),
				ScopeID:   core.StringPtr(resourceGroupID),
				ScopeType: core.StringPtr("account.resource_group"),
			}

			templateAttachmentRequestModel := &configurationgovernancev1.TemplateAttachmentRequest{
				AccountID:      core.StringPtr(accountID),
				IncludedScope:  templateScopeModel,
				ExcludedScopes: []configurationgovernancev1.TemplateScope{*excludedScopeModel},
			}

			createTemplateAttachmentsOptions := &configurationgovernancev1.CreateTemplateAttachmentsOptions{
				TemplateID:    &templateIDLink,
				Attachments:   []configurationgovernancev1.TemplateAttachmentRequest{*templateAttachmentRequestModel},
				TransactionID: core.StringPtr("testString"),
			}

			createTemplateAttachmentsResponse, response, err := configurationGovernanceService.CreateTemplateAttachments(createTemplateAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateAttachmentsResponse).ToNot(BeNil())

			templateAttachmentIDLink = *createTemplateAttachmentsResponse.Attachments[0].AttachmentID

		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {

			listRulesOptions := &configurationgovernancev1.ListRulesOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr("testString"),
				Attached:      core.BoolPtr(true),
				Labels:        core.StringPtr(ruleLabel),
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

			ruleEtag = response.Headers.Get("Etag")
		})
	})

	Describe(`UpdateRule - Update a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRule(updateRuleOptions *UpdateRuleOptions)`, func() {

			targetResourceAdditionalTargetAttributesItemModel := &configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{
				Name:     core.StringPtr("resource_id"),
				Value:    core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel},
			}

			ruleConditionModel := &configurationgovernancev1.RuleConditionSingleProperty{
				Description: core.StringPtr("testString"),
				Property:    core.StringPtr("location"),
				Operator:    core.StringPtr("string_equals"),
				Value:       core.StringPtr("us-south"),
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("Cloud Object Storage bucket"),
				And:         []configurationgovernancev1.RuleConditionIntf{ruleConditionModel},
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			updateRuleOptions := &configurationgovernancev1.UpdateRuleOptions{
				RuleID:             &ruleIDLink,
				IfMatch:            core.StringPtr(ruleEtag),
				Name:               core.StringPtr("Disable public access"),
				Description:        core.StringPtr("Disable public access in Dallas"),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				AccountID:          core.StringPtr(accountID),
				RuleType:           core.StringPtr("user_defined"),
				Labels:             []string{ruleLabel},
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

			ruleAttachmentEtag = response.Headers.Get("Etag")
		})
	})

	Describe(`UpdateRuleAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRuleAttachment(updateRuleAttachmentOptions *UpdateRuleAttachmentOptions)`, func() {

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My account"),
				ScopeID:   core.StringPtr(accountID),
				ScopeType: core.StringPtr("account"),
			}

			excludedScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My account resource group"),
				ScopeID:   core.StringPtr(resourceGroupID),
				ScopeType: core.StringPtr("account.resource_group"),
			}

			updateRuleAttachmentOptions := &configurationgovernancev1.UpdateRuleAttachmentOptions{
				RuleID:         &ruleIDLink,
				AttachmentID:   &ruleAttachmentIDLink,
				IfMatch:        core.StringPtr(ruleAttachmentEtag),
				AccountID:      core.StringPtr(accountID),
				IncludedScope:  ruleScopeModel,
				ExcludedScopes: []configurationgovernancev1.RuleScope{*excludedScopeModel},
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

	Describe(`GetTemplate - Get a template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTemplate(getTemplateOptions *GetTemplateOptions)`, func() {

			getTemplateOptions := &configurationgovernancev1.GetTemplateOptions{
				TemplateID:    &templateIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			templateResponse, response, err := configurationGovernanceService.GetTemplate(getTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())

			templateEtag = response.Headers.Get("Etag")
		})
	})

	Describe(`ListTemplates - List templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTemplates(listTemplatesOptions *ListTemplatesOptions)`, func() {

			listTemplatesOptions := &configurationgovernancev1.ListTemplatesOptions{
				AccountID:     core.StringPtr(accountID),
				TransactionID: core.StringPtr("testString"),
			}

			templateList, response, err := configurationGovernanceService.ListTemplates(listTemplatesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateList).ToNot(BeNil())

		})
	})

	Describe(`UpdateTemplate - Update a template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions)`, func() {

			baseTargetAttributeModel := &configurationgovernancev1.BaseTargetAttribute{
				Name:  core.StringPtr("location"),
				Value: core.StringPtr("us-south"),
			}

			simpleTargetResourceModel := &configurationgovernancev1.SimpleTargetResource{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configurationgovernancev1.BaseTargetAttribute{*baseTargetAttributeModel},
			}

			templateCustomizedDefaultPropertyModel := &configurationgovernancev1.TemplateCustomizedDefaultProperty{
				Property: core.StringPtr("location"),
				Value:    core.StringPtr("testString"),
			}

			updateTemplateOptions := &configurationgovernancev1.UpdateTemplateOptions{
				TemplateID:         &templateIDLink,
				IfMatch:            core.StringPtr(templateEtag),
				AccountID:          core.StringPtr(accountID),
				Name:               core.StringPtr("testString"),
				Description:        core.StringPtr("testString"),
				Target:             simpleTargetResourceModel,
				CustomizedDefaults: []configurationgovernancev1.TemplateCustomizedDefaultProperty{*templateCustomizedDefaultPropertyModel},
				TransactionID:      core.StringPtr("testString"),
			}

			templateResponse, response, err := configurationGovernanceService.UpdateTemplate(updateTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())

		})
	})

	Describe(`ListTemplateAttachments - List attachments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTemplateAttachments(listTemplateAttachmentsOptions *ListTemplateAttachmentsOptions)`, func() {

			listTemplateAttachmentsOptions := &configurationgovernancev1.ListTemplateAttachmentsOptions{
				TemplateID:    &templateIDLink,
				TransactionID: core.StringPtr("testString"),
				Limit:         core.Int64Ptr(int64(1000)),
				Offset:        core.Int64Ptr(int64(38)),
			}

			templateAttachmentList, response, err := configurationGovernanceService.ListTemplateAttachments(listTemplateAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachmentList).ToNot(BeNil())

		})
	})

	Describe(`GetTemplateAttachment - Get an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTemplateAttachment(getTemplateAttachmentOptions *GetTemplateAttachmentOptions)`, func() {

			getTemplateAttachmentOptions := &configurationgovernancev1.GetTemplateAttachmentOptions{
				TemplateID:    &templateIDLink,
				AttachmentID:  &templateAttachmentIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			templateAttachment, response, err := configurationGovernanceService.GetTemplateAttachment(getTemplateAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

			templateAttachmentEtag = response.Headers.Get("Etag")
		})
	})

	Describe(`UpdateTemplateAttachment - Update an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTemplateAttachment(updateTemplateAttachmentOptions *UpdateTemplateAttachmentOptions)`, func() {

			templateScopeModel := &configurationgovernancev1.TemplateScope{
				Note:      core.StringPtr("My account resource group new"),
				ScopeID:   core.StringPtr(resourceGroupID),
				ScopeType: core.StringPtr("account.resource_group"),
			}

			updateTemplateAttachmentOptions := &configurationgovernancev1.UpdateTemplateAttachmentOptions{
				TemplateID:    &templateIDLink,
				AttachmentID:  &templateAttachmentIDLink,
				AccountID:     core.StringPtr(accountID),
				IncludedScope: templateScopeModel,
				IfMatch:       core.StringPtr(templateAttachmentEtag),
				TransactionID: core.StringPtr("testString"),
			}

			templateAttachment, response, err := configurationGovernanceService.UpdateTemplateAttachment(updateTemplateAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

		})
	})

	Describe(`DeleteTemplateAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTemplateAttachment(deleteTemplateAttachmentOptions *DeleteTemplateAttachmentOptions)`, func() {

			deleteTemplateAttachmentOptions := &configurationgovernancev1.DeleteTemplateAttachmentOptions{
				TemplateID:    &templateIDLink,
				AttachmentID:  &templateAttachmentIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteTemplateAttachment(deleteTemplateAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteTemplate - Delete a template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {

			deleteTemplateOptions := &configurationgovernancev1.DeleteTemplateOptions{
				TemplateID:    &templateIDLink,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := configurationGovernanceService.DeleteTemplate(deleteTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
var _ = AfterSuite(func() {
	const externalConfigFile = "../configuration_governance_v1.env"
	_, err := os.Stat(externalConfigFile)
	if err != nil {
		Skip("External configuration file not found, skipping tests: " + err.Error())
	}
	os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
	config, err := core.GetServiceProperties(configurationgovernancev1.DefaultServiceName)
	if err != nil {
		Skip("Error loading service properties, skipping tests: " + err.Error())
	}
	serviceURL := config["URL"]
	if serviceURL == "" {
		Skip("Unable to load service URL configuration property, skipping tests")
	}

	fmt.Printf("cleaning up account: %s with rules labelled: %s\n", accountID, ruleLabel)

	configurationGovernanceServiceOptions := &configurationgovernancev1.ConfigurationGovernanceV1Options{}

	configurationGovernanceService, err := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

	listRulesOptions := &configurationgovernancev1.ListRulesOptions{
		AccountID: core.StringPtr(accountID),
	}

	rules, _, err := configurationGovernanceService.ListRules(listRulesOptions)
	if err != nil {
		Skip("Error occurred while listing rules for cleanup: " + err.Error())
	}

	for _, rule := range rules.Rules {
		if len(rule.Labels) > 0 && rule.Labels[0] == fmt.Sprintf("%s-%s", ruleLabel, identifier) {
			deleteRuleOptions := &configurationgovernancev1.DeleteRuleOptions{
				RuleID: rule.RuleID,
			}

			_, err := configurationGovernanceService.DeleteRule(deleteRuleOptions)
			if err != nil {
				Skip("Error occurred while deleting rule for cleanup: " + err.Error())
			}
		}
	}

	fmt.Printf("cleanup was successful\n")
})

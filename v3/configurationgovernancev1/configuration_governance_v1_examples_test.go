//go:build examples
// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Configuration Governance service.
//
// The following configuration properties are assumed to be defined:
// CONFIGURATION_GOVERNANCE_URL=<service base url>
// CONFIGURATION_GOVERNANCE_AUTH_TYPE=iam
// CONFIGURATION_GOVERNANCE_APIKEY=<IAM apikey>
// CONFIGURATION_GOVERNANCE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../configuration_governance_v1.env"

var (
	configurationGovernanceService *configurationgovernancev1.ConfigurationGovernanceV1
	config                         map[string]string
	configLoaded                   bool = false
)

// Global variables to hold link values
var (
	ruleAttachmentIDLink     string
	ruleIDLink               string
	templateAttachmentIDLink string
	templateIDLink           string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ConfigurationGovernanceV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(configurationgovernancev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			configurationGovernanceServiceOptions := &configurationgovernancev1.ConfigurationGovernanceV1Options{}

			configurationGovernanceService, err = configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(configurationGovernanceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(configurationGovernanceService).ToNot(BeNil())
		})
	})

	Describe(`ConfigurationGovernanceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRules request example`, func() {
			fmt.Println("\nCreateRules() result:")
			// begin-create_rules

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName:  core.StringPtr("cloud-object-storage"),
				ResourceKind: core.StringPtr("bucket"),
			}

			ruleConditionModel := &configurationgovernancev1.RuleConditionSingleProperty{
				Property: core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-south"),
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("Cloud Object Storage bucket"),
				And:         []configurationgovernancev1.RuleConditionIntf{ruleConditionModel},
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			ruleRequestModel := &configurationgovernancev1.RuleRequest{
				AccountID:          core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c"),
				Name:               core.StringPtr("Disable public access in Dallas"),
				Description:        core.StringPtr("Ensure that public access to buckets in us-south is disabled."),
				Target:             targetResourceModel,
				RequiredConfig:     ruleRequiredConfigModel,
				EnforcementActions: []configurationgovernancev1.EnforcementAction{*enforcementActionModel},
				Labels:             []string{"SOC2", "ITCS300"},
			}

			createRuleRequestModel := &configurationgovernancev1.CreateRuleRequest{
				RequestID: core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558"),
				Rule:      ruleRequestModel,
			}

			createRulesOptions := configurationGovernanceService.NewCreateRulesOptions(
				[]configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel},
			)

			createRulesResponse, response, err := configurationGovernanceService.CreateRules(createRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createRulesResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRulesResponse).ToNot(BeNil())

			ruleIDLink = *createRulesResponse.Rules[0].Rule.RuleID

		})
		It(`CreateRuleAttachments request example`, func() {
			fmt.Println("\nCreateRuleAttachments() result:")
			// begin-create_rule_attachments

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

			createRuleAttachmentsOptions := configurationGovernanceService.NewCreateRuleAttachmentsOptions(
				ruleIDLink,
				[]configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel},
			)

			createRuleAttachmentsResponse, response, err := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createRuleAttachmentsResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_rule_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createRuleAttachmentsResponse).ToNot(BeNil())

			ruleAttachmentIDLink = *createRuleAttachmentsResponse.Attachments[0].AttachmentID

		})
		It(`CreateTemplates request example`, func() {
			fmt.Println("\nCreateTemplates() result:")
			// begin-create_templates

			simpleTargetResourceModel := &configurationgovernancev1.SimpleTargetResource{
				ServiceName:  core.StringPtr("cloud-object-storage"),
				ResourceKind: core.StringPtr("bucket"),
			}

			templateCustomizedDefaultPropertyModel := &configurationgovernancev1.TemplateCustomizedDefaultProperty{
				Property: core.StringPtr("testString"),
				Value:    core.StringPtr("testString"),
			}

			templateModel := &configurationgovernancev1.Template{
				AccountID:          core.StringPtr("testString"),
				Name:               core.StringPtr("testString"),
				Description:        core.StringPtr("testString"),
				Target:             simpleTargetResourceModel,
				CustomizedDefaults: []configurationgovernancev1.TemplateCustomizedDefaultProperty{*templateCustomizedDefaultPropertyModel},
			}

			createTemplateRequestModel := &configurationgovernancev1.CreateTemplateRequest{
				Template: templateModel,
			}

			createTemplatesOptions := configurationGovernanceService.NewCreateTemplatesOptions(
				[]configurationgovernancev1.CreateTemplateRequest{*createTemplateRequestModel},
			)

			createTemplatesResponse, response, err := configurationGovernanceService.CreateTemplates(createTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplatesResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplatesResponse).ToNot(BeNil())

			templateIDLink = *createTemplatesResponse.Templates[0].Template.TemplateID

		})
		It(`CreateTemplateAttachments request example`, func() {
			fmt.Println("\nCreateTemplateAttachments() result:")
			// begin-create_template_attachments

			templateScopeModel := &configurationgovernancev1.TemplateScope{
				ScopeID:   core.StringPtr("testString"),
				ScopeType: core.StringPtr("enterprise"),
			}

			templateAttachmentRequestModel := &configurationgovernancev1.TemplateAttachmentRequest{
				AccountID:     core.StringPtr("testString"),
				IncludedScope: templateScopeModel,
			}

			createTemplateAttachmentsOptions := configurationGovernanceService.NewCreateTemplateAttachmentsOptions(
				templateIDLink,
				[]configurationgovernancev1.TemplateAttachmentRequest{*templateAttachmentRequestModel},
			)

			createTemplateAttachmentsResponse, response, err := configurationGovernanceService.CreateTemplateAttachments(createTemplateAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateAttachmentsResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_template_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateAttachmentsResponse).ToNot(BeNil())

			templateAttachmentIDLink = *createTemplateAttachmentsResponse.Attachments[0].AttachmentID

		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := configurationGovernanceService.NewListRulesOptions(
				"531fc3e28bfc43c5a2cea07786d93f5c",
			)
			listRulesOptions.SetAttached(true)
			listRulesOptions.SetLabels("SOC2,ITCS300")
			listRulesOptions.SetScopes("scope_id")

			ruleList, response, err := configurationGovernanceService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleList, "", "  ")
			fmt.Println(string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())

		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := configurationGovernanceService.NewGetRuleOptions(
				ruleIDLink,
			)

			rule, response, err := configurationGovernanceService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`UpdateRule request example`, func() {
			fmt.Println("\nUpdateRule() result:")
			// begin-update_rule

			targetResourceAdditionalTargetAttributesItemModel := &configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{
				Name:     core.StringPtr("testString"),
				Value:    core.StringPtr("testString"),
				Operator: core.StringPtr("string_equals"),
			}

			targetResourceModel := &configurationgovernancev1.TargetResource{
				ServiceName:                core.StringPtr("cloud-object-storage"),
				ResourceKind:               core.StringPtr("bucket"),
				AdditionalTargetAttributes: []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel},
			}

			ruleConditionModel := &configurationgovernancev1.RuleConditionSingleProperty{
				Property: core.StringPtr("location"),
				Operator: core.StringPtr("string_equals"),
				Value:    core.StringPtr("us-south"),
			}

			ruleRequiredConfigModel := &configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd{
				Description: core.StringPtr("Cloud Object Storage bucket"),
				And:         []configurationgovernancev1.RuleConditionIntf{ruleConditionModel},
			}

			enforcementActionModel := &configurationgovernancev1.EnforcementAction{
				Action: core.StringPtr("disallow"),
			}

			updateRuleOptions := configurationGovernanceService.NewUpdateRuleOptions(
				ruleIDLink,
				"testString",
				"Disable public access",
				"Disable public access in Dallas",
				targetResourceModel,
				ruleRequiredConfigModel,
				[]configurationgovernancev1.EnforcementAction{*enforcementActionModel},
			)
			updateRuleOptions.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
			updateRuleOptions.SetRuleType("user_defined")
			updateRuleOptions.SetLabels([]string{"SOC2", "ITCS300"})

			rule, response, err := configurationGovernanceService.UpdateRule(updateRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-update_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`ListRuleAttachments request example`, func() {
			fmt.Println("\nListRuleAttachments() result:")
			// begin-list_rule_attachments

			listRuleAttachmentsOptions := configurationGovernanceService.NewListRuleAttachmentsOptions(
				ruleIDLink,
			)

			ruleAttachmentList, response, err := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleAttachmentList, "", "  ")
			fmt.Println(string(b))

			// end-list_rule_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleAttachmentList).ToNot(BeNil())

		})
		It(`GetRuleAttachment request example`, func() {
			fmt.Println("\nGetRuleAttachment() result:")
			// begin-get_rule_attachment

			getRuleAttachmentOptions := configurationGovernanceService.NewGetRuleAttachmentOptions(
				ruleIDLink,
				ruleAttachmentIDLink,
			)

			ruleAttachment, response, err := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleAttachment, "", "  ")
			fmt.Println(string(b))

			// end-get_rule_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleAttachment).ToNot(BeNil())

		})
		It(`UpdateRuleAttachment request example`, func() {
			fmt.Println("\nUpdateRuleAttachment() result:")
			// begin-update_rule_attachment

			ruleScopeModel := &configurationgovernancev1.RuleScope{
				Note:      core.StringPtr("My enterprise"),
				ScopeID:   core.StringPtr("282cf433ac91493ba860480d92519990"),
				ScopeType: core.StringPtr("enterprise"),
			}

			updateRuleAttachmentOptions := configurationGovernanceService.NewUpdateRuleAttachmentOptions(
				ruleIDLink,
				ruleAttachmentIDLink,
				"testString",
				"531fc3e28bfc43c5a2cea07786d93f5c",
				ruleScopeModel,
			)
			updateRuleAttachmentOptions.SetExcludedScopes([]configurationgovernancev1.RuleScope{*ruleScopeModel})

			templateAttachment, response, err := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateAttachment, "", "  ")
			fmt.Println(string(b))

			// end-update_rule_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

		})
		It(`ListTemplates request example`, func() {
			fmt.Println("\nListTemplates() result:")
			// begin-list_templates

			listTemplatesOptions := configurationGovernanceService.NewListTemplatesOptions(
				"531fc3e28bfc43c5a2cea07786d93f5c",
			)

			templateList, response, err := configurationGovernanceService.ListTemplates(listTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateList, "", "  ")
			fmt.Println(string(b))

			// end-list_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateList).ToNot(BeNil())

		})
		It(`GetTemplate request example`, func() {
			fmt.Println("\nGetTemplate() result:")
			// begin-get_template

			getTemplateOptions := configurationGovernanceService.NewGetTemplateOptions(
				templateIDLink,
			)

			templateResponse, response, err := configurationGovernanceService.GetTemplate(getTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())

		})
		It(`UpdateTemplate request example`, func() {
			fmt.Println("\nUpdateTemplate() result:")
			// begin-update_template

			simpleTargetResourceModel := &configurationgovernancev1.SimpleTargetResource{
				ServiceName:  core.StringPtr("cloud-object-storage"),
				ResourceKind: core.StringPtr("bucket"),
			}

			templateCustomizedDefaultPropertyModel := &configurationgovernancev1.TemplateCustomizedDefaultProperty{
				Property: core.StringPtr("testString"),
				Value:    core.StringPtr("testString"),
			}

			updateTemplateOptions := configurationGovernanceService.NewUpdateTemplateOptions(
				templateIDLink,
				"testString",
				"testString",
				"testString",
				"testString",
				simpleTargetResourceModel,
				[]configurationgovernancev1.TemplateCustomizedDefaultProperty{*templateCustomizedDefaultPropertyModel},
			)

			templateResponse, response, err := configurationGovernanceService.UpdateTemplate(updateTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())

		})
		It(`ListTemplateAttachments request example`, func() {
			fmt.Println("\nListTemplateAttachments() result:")
			// begin-list_template_attachments

			listTemplateAttachmentsOptions := configurationGovernanceService.NewListTemplateAttachmentsOptions(
				templateIDLink,
			)

			templateAttachmentList, response, err := configurationGovernanceService.ListTemplateAttachments(listTemplateAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateAttachmentList, "", "  ")
			fmt.Println(string(b))

			// end-list_template_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachmentList).ToNot(BeNil())

		})
		It(`GetTemplateAttachment request example`, func() {
			fmt.Println("\nGetTemplateAttachment() result:")
			// begin-get_template_attachment

			getTemplateAttachmentOptions := configurationGovernanceService.NewGetTemplateAttachmentOptions(
				templateIDLink,
				templateAttachmentIDLink,
			)

			templateAttachment, response, err := configurationGovernanceService.GetTemplateAttachment(getTemplateAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateAttachment, "", "  ")
			fmt.Println(string(b))

			// end-get_template_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

		})
		It(`UpdateTemplateAttachment request example`, func() {
			fmt.Println("\nUpdateTemplateAttachment() result:")
			// begin-update_template_attachment

			templateScopeModel := &configurationgovernancev1.TemplateScope{
				ScopeID:   core.StringPtr("testString"),
				ScopeType: core.StringPtr("enterprise"),
			}

			updateTemplateAttachmentOptions := configurationGovernanceService.NewUpdateTemplateAttachmentOptions(
				templateIDLink,
				templateAttachmentIDLink,
				"testString",
				"testString",
				templateScopeModel,
			)

			templateAttachment, response, err := configurationGovernanceService.UpdateTemplateAttachment(updateTemplateAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateAttachment, "", "  ")
			fmt.Println(string(b))

			// end-update_template_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAttachment).ToNot(BeNil())

		})
		It(`DeleteTemplateAttachment request example`, func() {
			// begin-delete_template_attachment

			deleteTemplateAttachmentOptions := configurationGovernanceService.NewDeleteTemplateAttachmentOptions(
				templateIDLink,
				templateAttachmentIDLink,
			)

			response, err := configurationGovernanceService.DeleteTemplateAttachment(deleteTemplateAttachmentOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_template_attachment
			fmt.Printf("\nDeleteTemplateAttachment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTemplate request example`, func() {
			// begin-delete_template

			deleteTemplateOptions := configurationGovernanceService.NewDeleteTemplateOptions(
				templateIDLink,
			)

			response, err := configurationGovernanceService.DeleteTemplate(deleteTemplateOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_template
			fmt.Printf("\nDeleteTemplate() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteRuleAttachment request example`, func() {
			// begin-delete_rule_attachment

			deleteRuleAttachmentOptions := configurationGovernanceService.NewDeleteRuleAttachmentOptions(
				ruleIDLink,
				ruleAttachmentIDLink,
			)

			response, err := configurationGovernanceService.DeleteRuleAttachment(deleteRuleAttachmentOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_rule_attachment
			fmt.Printf("\nDeleteRuleAttachment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := configurationGovernanceService.NewDeleteRuleOptions(
				ruleIDLink,
			)

			response, err := configurationGovernanceService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_rule
			fmt.Printf("\nDeleteRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

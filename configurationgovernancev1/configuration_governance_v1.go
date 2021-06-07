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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.32.0-4c6a3129-20210514-210323
 */

// Package configurationgovernancev1 : Operations and models for the ConfigurationGovernanceV1 service
package configurationgovernancev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/ibm/scc-go-sdk/common"
)

// ConfigurationGovernanceV1 : API specification for the Configuration Governance service.
//
// Version: 1.0.0
type ConfigurationGovernanceV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us.compliance.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "configuration_governance"

// ConfigurationGovernanceV1Options : Service options
type ConfigurationGovernanceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewConfigurationGovernanceV1UsingExternalConfig : constructs an instance of ConfigurationGovernanceV1 with passed in options and external configuration.
func NewConfigurationGovernanceV1UsingExternalConfig(options *ConfigurationGovernanceV1Options) (configurationGovernance *ConfigurationGovernanceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	configurationGovernance, err = NewConfigurationGovernanceV1(options)
	if err != nil {
		return
	}

	err = configurationGovernance.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = configurationGovernance.Service.SetServiceURL(options.URL)
	}
	return
}

// NewConfigurationGovernanceV1 : constructs an instance of ConfigurationGovernanceV1 with passed in options.
func NewConfigurationGovernanceV1(options *ConfigurationGovernanceV1Options) (service *ConfigurationGovernanceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &ConfigurationGovernanceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://us.compliance.cloud.ibm.com",
		"us-east":  "https://us.compliance.cloud.ibm.com",
		"eu-de":    "https://eu.compliance.cloud.ibm.com",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "configurationGovernance" suitable for processing requests.
func (configurationGovernance *ConfigurationGovernanceV1) Clone() *ConfigurationGovernanceV1 {
	if core.IsNil(configurationGovernance) {
		return nil
	}
	clone := *configurationGovernance
	clone.Service = configurationGovernance.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (configurationGovernance *ConfigurationGovernanceV1) SetServiceURL(url string) error {
	return configurationGovernance.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (configurationGovernance *ConfigurationGovernanceV1) GetServiceURL() string {
	return configurationGovernance.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (configurationGovernance *ConfigurationGovernanceV1) SetDefaultHeaders(headers http.Header) {
	configurationGovernance.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (configurationGovernance *ConfigurationGovernanceV1) SetEnableGzipCompression(enableGzip bool) {
	configurationGovernance.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (configurationGovernance *ConfigurationGovernanceV1) GetEnableGzipCompression() bool {
	return configurationGovernance.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (configurationGovernance *ConfigurationGovernanceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	configurationGovernance.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (configurationGovernance *ConfigurationGovernanceV1) DisableRetries() {
	configurationGovernance.Service.DisableRetries()
}

// CreateRules : Create rules
// Creates one or more rules that you can use to govern the way that IBM Cloud resources can be provisioned and
// configured.
//
// A successful `POST /config/rules` request defines a rule based on the target, conditions, and enforcement actions
// that you specify. The response returns the ID value for your rule, along with other metadata.
//
// To learn more about rules, check out the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule).
func (configurationGovernance *ConfigurationGovernanceV1) CreateRules(createRulesOptions *CreateRulesOptions) (result *CreateRulesResponse, response *core.DetailedResponse, err error) {
	return configurationGovernance.CreateRulesWithContext(context.Background(), createRulesOptions)
}

// CreateRulesWithContext is an alternate form of the CreateRules method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) CreateRulesWithContext(ctx context.Context, createRulesOptions *CreateRulesOptions) (result *CreateRulesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRulesOptions, "createRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRulesOptions, "createRulesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "CreateRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createRulesOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createRulesOptions.Rules != nil {
		body["rules"] = createRulesOptions.Rules
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateRulesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRules : List rules
// Retrieves a list of the rules that are available in your account.
func (configurationGovernance *ConfigurationGovernanceV1) ListRules(listRulesOptions *ListRulesOptions) (result *RuleList, response *core.DetailedResponse, err error) {
	return configurationGovernance.ListRulesWithContext(context.Background(), listRulesOptions)
}

// ListRulesWithContext is an alternate form of the ListRules method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) ListRulesWithContext(ctx context.Context, listRulesOptions *ListRulesOptions) (result *RuleList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listRulesOptions, "listRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listRulesOptions, "listRulesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "ListRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listRulesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listRulesOptions.AccountID))
	if listRulesOptions.Attached != nil {
		builder.AddQuery("attached", fmt.Sprint(*listRulesOptions.Attached))
	}
	if listRulesOptions.Labels != nil {
		builder.AddQuery("labels", fmt.Sprint(*listRulesOptions.Labels))
	}
	if listRulesOptions.Scopes != nil {
		builder.AddQuery("scopes", fmt.Sprint(*listRulesOptions.Scopes))
	}
	if listRulesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listRulesOptions.Limit))
	}
	if listRulesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listRulesOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRule : Get a rule
// Retrieves an existing rule and its details.
func (configurationGovernance *ConfigurationGovernanceV1) GetRule(getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configurationGovernance.GetRuleWithContext(context.Background(), getRuleOptions)
}

// GetRuleWithContext is an alternate form of the GetRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) GetRuleWithContext(ctx context.Context, getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRuleOptions, "getRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRuleOptions, "getRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *getRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "GetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateRule : Update a rule
// Updates an existing rule based on the properties that you specify.
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRule(updateRuleOptions *UpdateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configurationGovernance.UpdateRuleWithContext(context.Background(), updateRuleOptions)
}

// UpdateRuleWithContext is an alternate form of the UpdateRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRuleWithContext(ctx context.Context, updateRuleOptions *UpdateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRuleOptions, "updateRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRuleOptions, "updateRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *updateRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "UpdateRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateRuleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateRuleOptions.IfMatch))
	}
	if updateRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateRuleOptions.Name != nil {
		body["name"] = updateRuleOptions.Name
	}
	if updateRuleOptions.Description != nil {
		body["description"] = updateRuleOptions.Description
	}
	if updateRuleOptions.Target != nil {
		body["target"] = updateRuleOptions.Target
	}
	if updateRuleOptions.RequiredConfig != nil {
		body["required_config"] = updateRuleOptions.RequiredConfig
	}
	if updateRuleOptions.EnforcementActions != nil {
		body["enforcement_actions"] = updateRuleOptions.EnforcementActions
	}
	if updateRuleOptions.AccountID != nil {
		body["account_id"] = updateRuleOptions.AccountID
	}
	if updateRuleOptions.RuleType != nil {
		body["rule_type"] = updateRuleOptions.RuleType
	}
	if updateRuleOptions.Labels != nil {
		body["labels"] = updateRuleOptions.Labels
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRule : Delete a rule
// Deletes an existing rule.
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRule(deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	return configurationGovernance.DeleteRuleWithContext(context.Background(), deleteRuleOptions)
}

// DeleteRuleWithContext is an alternate form of the DeleteRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRuleWithContext(ctx context.Context, deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRuleOptions, "deleteRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRuleOptions, "deleteRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *deleteRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "DeleteRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = configurationGovernance.Service.Request(request, nil)

	return
}

// CreateRuleAttachments : Create attachments
// Creates one or more scope attachments for an existing rule.
//
// You can attach an existing rule to a scope, such as a specific IBM Cloud account, to start evaluating the rule for
// compliance. A successful
// `POST /config/v1/rules/{rule_id}/attachments` returns the ID value for the attachment, along with other metadata.
func (configurationGovernance *ConfigurationGovernanceV1) CreateRuleAttachments(createRuleAttachmentsOptions *CreateRuleAttachmentsOptions) (result *CreateRuleAttachmentsResponse, response *core.DetailedResponse, err error) {
	return configurationGovernance.CreateRuleAttachmentsWithContext(context.Background(), createRuleAttachmentsOptions)
}

// CreateRuleAttachmentsWithContext is an alternate form of the CreateRuleAttachments method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) CreateRuleAttachmentsWithContext(ctx context.Context, createRuleAttachmentsOptions *CreateRuleAttachmentsOptions) (result *CreateRuleAttachmentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRuleAttachmentsOptions, "createRuleAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRuleAttachmentsOptions, "createRuleAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *createRuleAttachmentsOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRuleAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "CreateRuleAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRuleAttachmentsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createRuleAttachmentsOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createRuleAttachmentsOptions.Attachments != nil {
		body["attachments"] = createRuleAttachmentsOptions.Attachments
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateRuleAttachmentsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRuleAttachments : List attachments
// Retrieves a list of scope attachments that are associated with the specified rule.
func (configurationGovernance *ConfigurationGovernanceV1) ListRuleAttachments(listRuleAttachmentsOptions *ListRuleAttachmentsOptions) (result *RuleAttachmentList, response *core.DetailedResponse, err error) {
	return configurationGovernance.ListRuleAttachmentsWithContext(context.Background(), listRuleAttachmentsOptions)
}

// ListRuleAttachmentsWithContext is an alternate form of the ListRuleAttachments method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) ListRuleAttachmentsWithContext(ctx context.Context, listRuleAttachmentsOptions *ListRuleAttachmentsOptions) (result *RuleAttachmentList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listRuleAttachmentsOptions, "listRuleAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listRuleAttachmentsOptions, "listRuleAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *listRuleAttachmentsOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRuleAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "ListRuleAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRuleAttachmentsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listRuleAttachmentsOptions.TransactionID))
	}

	if listRuleAttachmentsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listRuleAttachmentsOptions.Limit))
	}
	if listRuleAttachmentsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listRuleAttachmentsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleAttachmentList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRuleAttachment : Get an attachment
// Retrieves an existing scope attachment for a rule.
func (configurationGovernance *ConfigurationGovernanceV1) GetRuleAttachment(getRuleAttachmentOptions *GetRuleAttachmentOptions) (result *RuleAttachment, response *core.DetailedResponse, err error) {
	return configurationGovernance.GetRuleAttachmentWithContext(context.Background(), getRuleAttachmentOptions)
}

// GetRuleAttachmentWithContext is an alternate form of the GetRuleAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) GetRuleAttachmentWithContext(ctx context.Context, getRuleAttachmentOptions *GetRuleAttachmentOptions) (result *RuleAttachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRuleAttachmentOptions, "getRuleAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRuleAttachmentOptions, "getRuleAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *getRuleAttachmentOptions.RuleID,
		"attachment_id": *getRuleAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "GetRuleAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getRuleAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getRuleAttachmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateRuleAttachment : Update an attachment
// Updates an existing scope attachment based on the properties that you specify.
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRuleAttachment(updateRuleAttachmentOptions *UpdateRuleAttachmentOptions) (result *TemplateAttachment, response *core.DetailedResponse, err error) {
	return configurationGovernance.UpdateRuleAttachmentWithContext(context.Background(), updateRuleAttachmentOptions)
}

// UpdateRuleAttachmentWithContext is an alternate form of the UpdateRuleAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRuleAttachmentWithContext(ctx context.Context, updateRuleAttachmentOptions *UpdateRuleAttachmentOptions) (result *TemplateAttachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRuleAttachmentOptions, "updateRuleAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRuleAttachmentOptions, "updateRuleAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *updateRuleAttachmentOptions.RuleID,
		"attachment_id": *updateRuleAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRuleAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "UpdateRuleAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateRuleAttachmentOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateRuleAttachmentOptions.IfMatch))
	}
	if updateRuleAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateRuleAttachmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateRuleAttachmentOptions.AccountID != nil {
		body["account_id"] = updateRuleAttachmentOptions.AccountID
	}
	if updateRuleAttachmentOptions.IncludedScope != nil {
		body["included_scope"] = updateRuleAttachmentOptions.IncludedScope
	}
	if updateRuleAttachmentOptions.ExcludedScopes != nil {
		body["excluded_scopes"] = updateRuleAttachmentOptions.ExcludedScopes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplateAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRuleAttachment : Delete an attachment
// Deletes an existing scope attachment.
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRuleAttachment(deleteRuleAttachmentOptions *DeleteRuleAttachmentOptions) (response *core.DetailedResponse, err error) {
	return configurationGovernance.DeleteRuleAttachmentWithContext(context.Background(), deleteRuleAttachmentOptions)
}

// DeleteRuleAttachmentWithContext is an alternate form of the DeleteRuleAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRuleAttachmentWithContext(ctx context.Context, deleteRuleAttachmentOptions *DeleteRuleAttachmentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRuleAttachmentOptions, "deleteRuleAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRuleAttachmentOptions, "deleteRuleAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *deleteRuleAttachmentOptions.RuleID,
		"attachment_id": *deleteRuleAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "DeleteRuleAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteRuleAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteRuleAttachmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = configurationGovernance.Service.Request(request, nil)

	return
}

// CreateRuleAttachmentsOptions : The CreateRuleAttachments options.
type CreateRuleAttachmentsOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	Attachments []RuleAttachmentRequest `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRuleAttachmentsOptions : Instantiate CreateRuleAttachmentsOptions
func (*ConfigurationGovernanceV1) NewCreateRuleAttachmentsOptions(ruleID string, attachments []RuleAttachmentRequest) *CreateRuleAttachmentsOptions {
	return &CreateRuleAttachmentsOptions{
		RuleID:      core.StringPtr(ruleID),
		Attachments: attachments,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *CreateRuleAttachmentsOptions) SetRuleID(ruleID string) *CreateRuleAttachmentsOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachments : Allow user to set Attachments
func (options *CreateRuleAttachmentsOptions) SetAttachments(attachments []RuleAttachmentRequest) *CreateRuleAttachmentsOptions {
	options.Attachments = attachments
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateRuleAttachmentsOptions) SetTransactionID(transactionID string) *CreateRuleAttachmentsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRuleAttachmentsOptions) SetHeaders(param map[string]string) *CreateRuleAttachmentsOptions {
	options.Headers = param
	return options
}

// CreateRuleAttachmentsResponse : CreateRuleAttachmentsResponse struct
type CreateRuleAttachmentsResponse struct {
	Attachments []RuleAttachment `json:"attachments" validate:"required"`
}

// UnmarshalCreateRuleAttachmentsResponse unmarshals an instance of CreateRuleAttachmentsResponse from the specified map of raw messages.
func UnmarshalCreateRuleAttachmentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRuleAttachmentsResponse)
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalRuleAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleRequest : A rule to be created.
type CreateRuleRequest struct {
	// A field that you can use in bulk operations to store a custom identifier for an individual request. If you omit this
	// field, the service generates and sends a `request_id` string for each new rule. The generated string corresponds
	// with the numerical order of the rules request array. For example, `"request_id": "1"`, `"request_id": "2"`.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `request_id` with
	// each request.
	RequestID *string `json:"request_id,omitempty"`

	// Properties that you can associate with a rule.
	Rule *RuleRequest `json:"rule" validate:"required"`
}

// NewCreateRuleRequest : Instantiate CreateRuleRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewCreateRuleRequest(rule *RuleRequest) (model *CreateRuleRequest, err error) {
	model = &CreateRuleRequest{
		Rule: rule,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateRuleRequest unmarshals an instance of CreateRuleRequest from the specified map of raw messages.
func UnmarshalCreateRuleRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRuleRequest)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rule", &obj.Rule, UnmarshalRuleRequest)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleResponse : Response information for a rule request.
//
// If the 'status_code' property indicates success, the 'request_id' and 'rule' properties are returned in the response.
// If the 'status_code' property indicates an error, the 'request_id', 'errors', and 'trace' fields are returned.
type CreateRuleResponse struct {
	// The identifier that is used to correlate an individual request.
	//
	// To assist with debugging, you can use this ID to identify and inspect only one request that was made as part of a
	// bulk operation.
	RequestID *string `json:"request_id,omitempty"`

	// The HTTP response status code.
	StatusCode *int64 `json:"status_code,omitempty"`

	// Information about a newly-created rule.
	//
	// This field is present for successful requests.
	Rule *Rule `json:"rule,omitempty"`

	// The error contents of the multi-status response.
	//
	// This field is present for unsuccessful requests.
	Errors []RuleResponseError `json:"errors,omitempty"`

	// The UUID that uniquely identifies the request.
	//
	// This field is present for unsuccessful requests.
	Trace *string `json:"trace,omitempty"`
}

// UnmarshalCreateRuleResponse unmarshals an instance of CreateRuleResponse from the specified map of raw messages.
func UnmarshalCreateRuleResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRuleResponse)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rule", &obj.Rule, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalRuleResponseError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRulesOptions : The CreateRules options.
type CreateRulesOptions struct {
	// A list of rules to be created.
	Rules []CreateRuleRequest `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRulesOptions : Instantiate CreateRulesOptions
func (*ConfigurationGovernanceV1) NewCreateRulesOptions(rules []CreateRuleRequest) *CreateRulesOptions {
	return &CreateRulesOptions{
		Rules: rules,
	}
}

// SetRules : Allow user to set Rules
func (options *CreateRulesOptions) SetRules(rules []CreateRuleRequest) *CreateRulesOptions {
	options.Rules = rules
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateRulesOptions) SetTransactionID(transactionID string) *CreateRulesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRulesOptions) SetHeaders(param map[string]string) *CreateRulesOptions {
	options.Headers = param
	return options
}

// CreateRulesResponse : The response associated with a request to create one or more rules.
type CreateRulesResponse struct {
	// An array of rule responses.
	Rules []CreateRuleResponse `json:"rules" validate:"required"`
}

// UnmarshalCreateRulesResponse unmarshals an instance of CreateRulesResponse from the specified map of raw messages.
func UnmarshalCreateRulesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRulesResponse)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalCreateRuleResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteRuleAttachmentOptions : The DeleteRuleAttachment options.
type DeleteRuleAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleAttachmentOptions : Instantiate DeleteRuleAttachmentOptions
func (*ConfigurationGovernanceV1) NewDeleteRuleAttachmentOptions(ruleID string, attachmentID string) *DeleteRuleAttachmentOptions {
	return &DeleteRuleAttachmentOptions{
		RuleID:       core.StringPtr(ruleID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *DeleteRuleAttachmentOptions) SetRuleID(ruleID string) *DeleteRuleAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *DeleteRuleAttachmentOptions) SetAttachmentID(attachmentID string) *DeleteRuleAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteRuleAttachmentOptions) SetTransactionID(transactionID string) *DeleteRuleAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRuleAttachmentOptions) SetHeaders(param map[string]string) *DeleteRuleAttachmentOptions {
	options.Headers = param
	return options
}

// DeleteRuleOptions : The DeleteRule options.
type DeleteRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleOptions : Instantiate DeleteRuleOptions
func (*ConfigurationGovernanceV1) NewDeleteRuleOptions(ruleID string) *DeleteRuleOptions {
	return &DeleteRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *DeleteRuleOptions) SetRuleID(ruleID string) *DeleteRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteRuleOptions) SetTransactionID(transactionID string) *DeleteRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRuleOptions) SetHeaders(param map[string]string) *DeleteRuleOptions {
	options.Headers = param
	return options
}

// EnforcementAction : EnforcementAction struct
type EnforcementAction struct {
	// To block a request from completing, use `disallow`. To log the request to Activity Tracker with LogDNA, use
	// `audit_log`.
	Action *string `json:"action" validate:"required"`
}

// Constants associated with the EnforcementAction.Action property.
// To block a request from completing, use `disallow`. To log the request to Activity Tracker with LogDNA, use
// `audit_log`.
const (
	EnforcementActionActionAuditLogConst = "audit_log"
	EnforcementActionActionDisallowConst = "disallow"
)

// NewEnforcementAction : Instantiate EnforcementAction (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewEnforcementAction(action string) (model *EnforcementAction, err error) {
	model = &EnforcementAction{
		Action: core.StringPtr(action),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEnforcementAction unmarshals an instance of EnforcementAction from the specified map of raw messages.
func UnmarshalEnforcementAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnforcementAction)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRuleAttachmentOptions : The GetRuleAttachment options.
type GetRuleAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleAttachmentOptions : Instantiate GetRuleAttachmentOptions
func (*ConfigurationGovernanceV1) NewGetRuleAttachmentOptions(ruleID string, attachmentID string) *GetRuleAttachmentOptions {
	return &GetRuleAttachmentOptions{
		RuleID:       core.StringPtr(ruleID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *GetRuleAttachmentOptions) SetRuleID(ruleID string) *GetRuleAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *GetRuleAttachmentOptions) SetAttachmentID(attachmentID string) *GetRuleAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetRuleAttachmentOptions) SetTransactionID(transactionID string) *GetRuleAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRuleAttachmentOptions) SetHeaders(param map[string]string) *GetRuleAttachmentOptions {
	options.Headers = param
	return options
}

// GetRuleOptions : The GetRule options.
type GetRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleOptions : Instantiate GetRuleOptions
func (*ConfigurationGovernanceV1) NewGetRuleOptions(ruleID string) *GetRuleOptions {
	return &GetRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *GetRuleOptions) SetRuleID(ruleID string) *GetRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetRuleOptions) SetTransactionID(transactionID string) *GetRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRuleOptions) SetHeaders(param map[string]string) *GetRuleOptions {
	options.Headers = param
	return options
}

// Link : A link that is used to paginate through available resources.
type Link struct {
	// The URL for the first, previous, next, or last page of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalLink unmarshals an instance of Link from the specified map of raw messages.
func UnmarshalLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Link)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListRuleAttachmentsOptions : The ListRuleAttachments options.
type ListRuleAttachmentsOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// The number of resources to retrieve. By default, list operations return the first 100 items. To retrieve a different
	// set of items, use `limit` with `offset` to page through your available resources.
	//
	// **Usage:** If you have 20 rules, and you want to retrieve only the first 5 rules, use
	// `../rules?account_id={account_id}&limit=5`.
	Limit *int64

	// The number of resources to skip. By specifying `offset`, you retrieve a subset of resources that starts with the
	// `offset` value. Use `offset` with `limit` to page through your available resources.
	//
	// **Usage:** If you have 100 rules, and you want to retrieve rules 26 through 50, use
	// `../rules?account_id={account_id}&offset=25&limit=5`.
	Offset *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRuleAttachmentsOptions : Instantiate ListRuleAttachmentsOptions
func (*ConfigurationGovernanceV1) NewListRuleAttachmentsOptions(ruleID string) *ListRuleAttachmentsOptions {
	return &ListRuleAttachmentsOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *ListRuleAttachmentsOptions) SetRuleID(ruleID string) *ListRuleAttachmentsOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListRuleAttachmentsOptions) SetTransactionID(transactionID string) *ListRuleAttachmentsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListRuleAttachmentsOptions) SetLimit(limit int64) *ListRuleAttachmentsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListRuleAttachmentsOptions) SetOffset(offset int64) *ListRuleAttachmentsOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListRuleAttachmentsOptions) SetHeaders(param map[string]string) *ListRuleAttachmentsOptions {
	options.Headers = param
	return options
}

// ListRulesOptions : The ListRules options.
type ListRulesOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Retrieves a list of rules that have scope attachments.
	Attached *bool

	// Retrieves a list of rules that match the labels that you specify.
	Labels *string

	// Retrieves a list of rules that match the scope ID that you specify.
	Scopes *string

	// The number of resources to retrieve. By default, list operations return the first 100 items. To retrieve a different
	// set of items, use `limit` with `offset` to page through your available resources.
	//
	// **Usage:** If you have 20 rules, and you want to retrieve only the first 5 rules, use
	// `../rules?account_id={account_id}&limit=5`.
	Limit *int64

	// The number of resources to skip. By specifying `offset`, you retrieve a subset of resources that starts with the
	// `offset` value. Use `offset` with `limit` to page through your available resources.
	//
	// **Usage:** If you have 100 rules, and you want to retrieve rules 26 through 50, use
	// `../rules?account_id={account_id}&offset=25&limit=5`.
	Offset *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRulesOptions : Instantiate ListRulesOptions
func (*ConfigurationGovernanceV1) NewListRulesOptions(accountID string) *ListRulesOptions {
	return &ListRulesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListRulesOptions) SetAccountID(accountID string) *ListRulesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListRulesOptions) SetTransactionID(transactionID string) *ListRulesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetAttached : Allow user to set Attached
func (options *ListRulesOptions) SetAttached(attached bool) *ListRulesOptions {
	options.Attached = core.BoolPtr(attached)
	return options
}

// SetLabels : Allow user to set Labels
func (options *ListRulesOptions) SetLabels(labels string) *ListRulesOptions {
	options.Labels = core.StringPtr(labels)
	return options
}

// SetScopes : Allow user to set Scopes
func (options *ListRulesOptions) SetScopes(scopes string) *ListRulesOptions {
	options.Scopes = core.StringPtr(scopes)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListRulesOptions) SetLimit(limit int64) *ListRulesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListRulesOptions) SetOffset(offset int64) *ListRulesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListRulesOptions) SetHeaders(param map[string]string) *ListRulesOptions {
	options.Headers = param
	return options
}

// Rule : Properties associated with a rule, including both user-defined and server-populated properties.
type Rule struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id,omitempty"`

	// A human-readable alias to assign to your rule.
	Name *string `json:"name" validate:"required"`

	// An extended description of your rule.
	Description *string `json:"description" validate:"required"`

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string `json:"rule_type,omitempty"`

	// The properties that describe the resource that you want to target
	// with the rule or template.
	Target *TargetResource `json:"target" validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `json:"required_config" validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `json:"enforcement_actions" validate:"required"`

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string `json:"labels,omitempty"`

	// The UUID that uniquely identifies the rule.
	RuleID *string `json:"rule_id,omitempty"`

	// The date the resource was created.
	CreationDate *strfmt.DateTime `json:"creation_date,omitempty"`

	// The unique identifier for the user or application that created the resource.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date the resource was last modified.
	ModificationDate *strfmt.DateTime `json:"modification_date,omitempty"`

	// The unique identifier for the user or application that last modified the resource.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// The number of scope attachments that are associated with the rule.
	NumberOfAttachments *int64 `json:"number_of_attachments,omitempty"`
}

// Constants associated with the Rule.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	RuleRuleTypeUserDefinedConst = "user_defined"
)

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_type", &obj.RuleType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTargetResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_config", &obj.RequiredConfig, UnmarshalRuleRequiredConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enforcement_actions", &obj.EnforcementActions, UnmarshalEnforcementAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creation_date", &obj.CreationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modification_date", &obj.ModificationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "number_of_attachments", &obj.NumberOfAttachments)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleAttachment : The scopes to attach to a rule.
type RuleAttachment struct {
	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `json:"attachment_id" validate:"required"`

	// The UUID that uniquely identifies the rule.
	RuleID *string `json:"rule_id" validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `json:"included_scope" validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope `json:"excluded_scopes,omitempty"`
}

// UnmarshalRuleAttachment unmarshals an instance of RuleAttachment from the specified map of raw messages.
func UnmarshalRuleAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleAttachment)
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalRuleScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded_scopes", &obj.ExcludedScopes, UnmarshalRuleScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleAttachmentList : A list of attachments.
type RuleAttachmentList struct {
	// The requested offset for the returned items.
	Offset *int64 `json:"offset" validate:"required"`

	// The requested limit for the returned items.
	Limit *int64 `json:"limit" validate:"required"`

	// The total number of available items.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The first page of available items.
	First *Link `json:"first" validate:"required"`

	// The last page of available items.
	Last *Link `json:"last" validate:"required"`

	Attachments []RuleAttachment `json:"attachments" validate:"required"`
}

// UnmarshalRuleAttachmentList unmarshals an instance of RuleAttachmentList from the specified map of raw messages.
func UnmarshalRuleAttachmentList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleAttachmentList)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalRuleAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleAttachmentRequest : The scopes to attach to a rule.
type RuleAttachmentRequest struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `json:"included_scope" validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope `json:"excluded_scopes,omitempty"`
}

// NewRuleAttachmentRequest : Instantiate RuleAttachmentRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleAttachmentRequest(accountID string, includedScope *RuleScope) (model *RuleAttachmentRequest, err error) {
	model = &RuleAttachmentRequest{
		AccountID:     core.StringPtr(accountID),
		IncludedScope: includedScope,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleAttachmentRequest unmarshals an instance of RuleAttachmentRequest from the specified map of raw messages.
func UnmarshalRuleAttachmentRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleAttachmentRequest)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalRuleScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded_scopes", &obj.ExcludedScopes, UnmarshalRuleScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleCondition : RuleCondition struct
// Models which "extend" this model:
// - RuleConditionSingleProperty
// - RuleConditionOrLvl2
// - RuleConditionAndLvl2
type RuleCondition struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property,omitempty"`

	// The way in which the `property` field is compared to its value.
	//
	// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
	Operator *string `json:"operator,omitempty"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`

	Or []RuleSingleProperty `json:"or,omitempty"`

	And []RuleSingleProperty `json:"and,omitempty"`
}

// Constants associated with the RuleCondition.Operator property.
// The way in which the `property` field is compared to its value.
//
// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
const (
	RuleConditionOperatorIpsInRangeConst           = "ips_in_range"
	RuleConditionOperatorIsEmptyConst              = "is_empty"
	RuleConditionOperatorIsFalseConst              = "is_false"
	RuleConditionOperatorIsNotEmptyConst           = "is_not_empty"
	RuleConditionOperatorIsTrueConst               = "is_true"
	RuleConditionOperatorNumEqualsConst            = "num_equals"
	RuleConditionOperatorNumGreaterThanConst       = "num_greater_than"
	RuleConditionOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleConditionOperatorNumLessThanConst          = "num_less_than"
	RuleConditionOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleConditionOperatorNumNotEqualsConst         = "num_not_equals"
	RuleConditionOperatorStringEqualsConst         = "string_equals"
	RuleConditionOperatorStringMatchConst          = "string_match"
	RuleConditionOperatorStringNotEqualsConst      = "string_not_equals"
	RuleConditionOperatorStringNotMatchConst       = "string_not_match"
	RuleConditionOperatorStringsInListConst        = "strings_in_list"
)

func (*RuleCondition) isaRuleCondition() bool {
	return true
}

type RuleConditionIntf interface {
	isaRuleCondition() bool
}

// UnmarshalRuleCondition unmarshals an instance of RuleCondition from the specified map of raw messages.
func UnmarshalRuleCondition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleCondition)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleList : A list of rules.
type RuleList struct {
	// The requested offset for the returned items.
	Offset *int64 `json:"offset" validate:"required"`

	// The requested limit for the returned items.
	Limit *int64 `json:"limit" validate:"required"`

	// The total number of available items.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The first page of available items.
	First *Link `json:"first" validate:"required"`

	// The last page of available items.
	Last *Link `json:"last" validate:"required"`

	// An array of rules.
	Rules []Rule `json:"rules" validate:"required"`
}

// UnmarshalRuleList unmarshals an instance of RuleList from the specified map of raw messages.
func UnmarshalRuleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleList)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequest : Properties that you can associate with a rule.
type RuleRequest struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id,omitempty"`

	// A human-readable alias to assign to your rule.
	Name *string `json:"name" validate:"required"`

	// An extended description of your rule.
	Description *string `json:"description" validate:"required"`

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string `json:"rule_type,omitempty"`

	// The properties that describe the resource that you want to target
	// with the rule or template.
	Target *TargetResource `json:"target" validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `json:"required_config" validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `json:"enforcement_actions" validate:"required"`

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string `json:"labels,omitempty"`
}

// Constants associated with the RuleRequest.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	RuleRequestRuleTypeUserDefinedConst = "user_defined"
)

// NewRuleRequest : Instantiate RuleRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequest(name string, description string, target *TargetResource, requiredConfig RuleRequiredConfigIntf, enforcementActions []EnforcementAction) (model *RuleRequest, err error) {
	model = &RuleRequest{
		Name:               core.StringPtr(name),
		Description:        core.StringPtr(description),
		Target:             target,
		RequiredConfig:     requiredConfig,
		EnforcementActions: enforcementActions,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleRequest unmarshals an instance of RuleRequest from the specified map of raw messages.
func UnmarshalRuleRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequest)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_type", &obj.RuleType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTargetResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_config", &obj.RequiredConfig, UnmarshalRuleRequiredConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enforcement_actions", &obj.EnforcementActions, UnmarshalEnforcementAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfig : RuleRequiredConfig struct
// Models which "extend" this model:
// - RuleRequiredConfigSingleProperty
// - RuleRequiredConfigMultipleProperties
type RuleRequiredConfig struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property,omitempty"`

	// The way in which the `property` field is compared to its value.
	//
	// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
	Operator *string `json:"operator,omitempty"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`

	Or []RuleConditionIntf `json:"or,omitempty"`

	And []RuleConditionIntf `json:"and,omitempty"`
}

// Constants associated with the RuleRequiredConfig.Operator property.
// The way in which the `property` field is compared to its value.
//
// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
const (
	RuleRequiredConfigOperatorIpsInRangeConst           = "ips_in_range"
	RuleRequiredConfigOperatorIsEmptyConst              = "is_empty"
	RuleRequiredConfigOperatorIsFalseConst              = "is_false"
	RuleRequiredConfigOperatorIsNotEmptyConst           = "is_not_empty"
	RuleRequiredConfigOperatorIsTrueConst               = "is_true"
	RuleRequiredConfigOperatorNumEqualsConst            = "num_equals"
	RuleRequiredConfigOperatorNumGreaterThanConst       = "num_greater_than"
	RuleRequiredConfigOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleRequiredConfigOperatorNumLessThanConst          = "num_less_than"
	RuleRequiredConfigOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleRequiredConfigOperatorNumNotEqualsConst         = "num_not_equals"
	RuleRequiredConfigOperatorStringEqualsConst         = "string_equals"
	RuleRequiredConfigOperatorStringMatchConst          = "string_match"
	RuleRequiredConfigOperatorStringNotEqualsConst      = "string_not_equals"
	RuleRequiredConfigOperatorStringNotMatchConst       = "string_not_match"
	RuleRequiredConfigOperatorStringsInListConst        = "strings_in_list"
)

func (*RuleRequiredConfig) isaRuleRequiredConfig() bool {
	return true
}

type RuleRequiredConfigIntf interface {
	isaRuleRequiredConfig() bool
}

// UnmarshalRuleRequiredConfig unmarshals an instance of RuleRequiredConfig from the specified map of raw messages.
func UnmarshalRuleRequiredConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfig)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleResponseError : RuleResponseError struct
type RuleResponseError struct {
	// Specifies the problem that caused the error.
	Code *string `json:"code" validate:"required"`

	// Describes the problem.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalRuleResponseError unmarshals an instance of RuleResponseError from the specified map of raw messages.
func UnmarshalRuleResponseError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleResponseError)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleScope : The extent at which the rule can be attached across your accounts.
type RuleScope struct {
	// A short description or alias to assign to the scope.
	Note *string `json:"note,omitempty"`

	// The ID of the scope, such as an enterprise, account, or account group, that you want to evaluate.
	ScopeID *string `json:"scope_id" validate:"required"`

	// The type of scope that you want to evaluate.
	ScopeType *string `json:"scope_type" validate:"required"`
}

// Constants associated with the RuleScope.ScopeType property.
// The type of scope that you want to evaluate.
const (
	RuleScopeScopeTypeAccountConst                = "account"
	RuleScopeScopeTypeAccountResourceGroupConst   = "account.resource_group"
	RuleScopeScopeTypeEnterpriseConst             = "enterprise"
	RuleScopeScopeTypeEnterpriseAccountConst      = "enterprise.account"
	RuleScopeScopeTypeEnterpriseAccountGroupConst = "enterprise.account_group"
)

// NewRuleScope : Instantiate RuleScope (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleScope(scopeID string, scopeType string) (model *RuleScope, err error) {
	model = &RuleScope{
		ScopeID:   core.StringPtr(scopeID),
		ScopeType: core.StringPtr(scopeType),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleScope unmarshals an instance of RuleScope from the specified map of raw messages.
func UnmarshalRuleScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleScope)
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_type", &obj.ScopeType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
type RuleSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
const (
	RuleSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleSingleProperty : Instantiate RuleSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleSingleProperty(property string, operator string) (model *RuleSingleProperty, err error) {
	model = &RuleSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleSingleProperty unmarshals an instance of RuleSingleProperty from the specified map of raw messages.
func UnmarshalRuleSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetResource : The properties that describe the resource that you want to target with the rule or template.
type TargetResource struct {
	// The programmatic name of the IBM Cloud service that you want to target with the rule or template.
	ServiceName *string `json:"service_name" validate:"required"`

	// The type of resource that you want to target.
	ResourceKind *string `json:"resource_kind" validate:"required"`

	// An extra qualifier for the resource kind. When you include additional attributes, only the resources that match the
	// definition are included in the rule or template.
	AdditionalTargetAttributes []TargetResourceAdditionalTargetAttributesItem `json:"additional_target_attributes,omitempty"`
}

// NewTargetResource : Instantiate TargetResource (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewTargetResource(serviceName string, resourceKind string) (model *TargetResource, err error) {
	model = &TargetResource{
		ServiceName:  core.StringPtr(serviceName),
		ResourceKind: core.StringPtr(resourceKind),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTargetResource unmarshals an instance of TargetResource from the specified map of raw messages.
func UnmarshalTargetResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetResource)
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_kind", &obj.ResourceKind)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "additional_target_attributes", &obj.AdditionalTargetAttributes, UnmarshalTargetResourceAdditionalTargetAttributesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetResourceAdditionalTargetAttributesItem : The attributes that are associated with a rule or template target.
type TargetResourceAdditionalTargetAttributesItem struct {
	// The name of the additional attribute that you want to use to further qualify the target.
	//
	// Options differ depending on the service or resource that you are targeting with a rule or template. For more
	// information, refer to the service documentation.
	Name *string `json:"name" validate:"required"`

	// The value that you want to apply to `name` field.
	//
	// Options differ depending on the rule or template that you configure. For more information, refer to the service
	// documentation.
	Value *string `json:"value" validate:"required"`

	// The way in which the `name` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator" validate:"required"`
}

// Constants associated with the TargetResourceAdditionalTargetAttributesItem.Operator property.
// The way in which the `name` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	TargetResourceAdditionalTargetAttributesItemOperatorIpsInRangeConst           = "ips_in_range"
	TargetResourceAdditionalTargetAttributesItemOperatorIsEmptyConst              = "is_empty"
	TargetResourceAdditionalTargetAttributesItemOperatorIsFalseConst              = "is_false"
	TargetResourceAdditionalTargetAttributesItemOperatorIsNotEmptyConst           = "is_not_empty"
	TargetResourceAdditionalTargetAttributesItemOperatorIsTrueConst               = "is_true"
	TargetResourceAdditionalTargetAttributesItemOperatorNumEqualsConst            = "num_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorNumGreaterThanConst       = "num_greater_than"
	TargetResourceAdditionalTargetAttributesItemOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorNumLessThanConst          = "num_less_than"
	TargetResourceAdditionalTargetAttributesItemOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorNumNotEqualsConst         = "num_not_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorStringEqualsConst         = "string_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorStringMatchConst          = "string_match"
	TargetResourceAdditionalTargetAttributesItemOperatorStringNotEqualsConst      = "string_not_equals"
	TargetResourceAdditionalTargetAttributesItemOperatorStringNotMatchConst       = "string_not_match"
	TargetResourceAdditionalTargetAttributesItemOperatorStringsInListConst        = "strings_in_list"
)

// NewTargetResourceAdditionalTargetAttributesItem : Instantiate TargetResourceAdditionalTargetAttributesItem (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewTargetResourceAdditionalTargetAttributesItem(name string, value string, operator string) (model *TargetResourceAdditionalTargetAttributesItem, err error) {
	model = &TargetResourceAdditionalTargetAttributesItem{
		Name:     core.StringPtr(name),
		Value:    core.StringPtr(value),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTargetResourceAdditionalTargetAttributesItem unmarshals an instance of TargetResourceAdditionalTargetAttributesItem from the specified map of raw messages.
func UnmarshalTargetResourceAdditionalTargetAttributesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetResourceAdditionalTargetAttributesItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TemplateAttachment : The scopes to attach to a template.
type TemplateAttachment struct {
	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `json:"attachment_id" validate:"required"`

	// The UUID that uniquely identifies the template.
	TemplateID *string `json:"template_id" validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The extent at which the template can be attached across your accounts.
	IncludedScope *TemplateScope `json:"included_scope" validate:"required"`

	ExcludedScopes []TemplateScope `json:"excluded_scopes,omitempty"`
}

// UnmarshalTemplateAttachment unmarshals an instance of TemplateAttachment from the specified map of raw messages.
func UnmarshalTemplateAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateAttachment)
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_id", &obj.TemplateID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalTemplateScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded_scopes", &obj.ExcludedScopes, UnmarshalTemplateScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TemplateScope : The extent at which the template can be attached across your accounts.
type TemplateScope struct {
	// A short description or alias to assign to the scope.
	Note *string `json:"note,omitempty"`

	// The ID of the scope, such as an enterprise, account, or account group, where you want to apply the customized
	// defaults that are associated with a template.
	ScopeID *string `json:"scope_id" validate:"required"`

	// The type of scope.
	ScopeType *string `json:"scope_type" validate:"required"`
}

// Constants associated with the TemplateScope.ScopeType property.
// The type of scope.
const (
	TemplateScopeScopeTypeAccountConst                = "account"
	TemplateScopeScopeTypeAccountResourceGroupConst   = "account.resource_group"
	TemplateScopeScopeTypeEnterpriseConst             = "enterprise"
	TemplateScopeScopeTypeEnterpriseAccountConst      = "enterprise.account"
	TemplateScopeScopeTypeEnterpriseAccountGroupConst = "enterprise.account_group"
)

// UnmarshalTemplateScope unmarshals an instance of TemplateScope from the specified map of raw messages.
func UnmarshalTemplateScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateScope)
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_type", &obj.ScopeType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateRuleAttachmentOptions : The UpdateRuleAttachment options.
type UpdateRuleAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// Compares a supplied `Etag` value with the version that is stored for the requested resource. If the values match,
	// the server allows the request method to continue.
	//
	// To find the `Etag` value, run a GET request on the resource that you want to modify, and check the response headers.
	IfMatch *string `validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRuleAttachmentOptions : Instantiate UpdateRuleAttachmentOptions
func (*ConfigurationGovernanceV1) NewUpdateRuleAttachmentOptions(ruleID string, attachmentID string, ifMatch string, accountID string, includedScope *RuleScope) *UpdateRuleAttachmentOptions {
	return &UpdateRuleAttachmentOptions{
		RuleID:        core.StringPtr(ruleID),
		AttachmentID:  core.StringPtr(attachmentID),
		IfMatch:       core.StringPtr(ifMatch),
		AccountID:     core.StringPtr(accountID),
		IncludedScope: includedScope,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *UpdateRuleAttachmentOptions) SetRuleID(ruleID string) *UpdateRuleAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *UpdateRuleAttachmentOptions) SetAttachmentID(attachmentID string) *UpdateRuleAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateRuleAttachmentOptions) SetIfMatch(ifMatch string) *UpdateRuleAttachmentOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateRuleAttachmentOptions) SetAccountID(accountID string) *UpdateRuleAttachmentOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIncludedScope : Allow user to set IncludedScope
func (options *UpdateRuleAttachmentOptions) SetIncludedScope(includedScope *RuleScope) *UpdateRuleAttachmentOptions {
	options.IncludedScope = includedScope
	return options
}

// SetExcludedScopes : Allow user to set ExcludedScopes
func (options *UpdateRuleAttachmentOptions) SetExcludedScopes(excludedScopes []RuleScope) *UpdateRuleAttachmentOptions {
	options.ExcludedScopes = excludedScopes
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateRuleAttachmentOptions) SetTransactionID(transactionID string) *UpdateRuleAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRuleAttachmentOptions) SetHeaders(param map[string]string) *UpdateRuleAttachmentOptions {
	options.Headers = param
	return options
}

// UpdateRuleOptions : The UpdateRule options.
type UpdateRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// Compares a supplied `Etag` value with the version that is stored for the requested resource. If the values match,
	// the server allows the request method to continue.
	//
	// To find the `Etag` value, run a GET request on the resource that you want to modify, and check the response headers.
	IfMatch *string `validate:"required"`

	// A human-readable alias to assign to your rule.
	Name *string `validate:"required"`

	// An extended description of your rule.
	Description *string `validate:"required"`

	// The properties that describe the resource that you want to target
	// with the rule or template.
	Target *TargetResource `validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateRuleOptions.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	UpdateRuleOptionsRuleTypeUserDefinedConst = "user_defined"
)

// NewUpdateRuleOptions : Instantiate UpdateRuleOptions
func (*ConfigurationGovernanceV1) NewUpdateRuleOptions(ruleID string, ifMatch string, name string, description string, target *TargetResource, requiredConfig RuleRequiredConfigIntf, enforcementActions []EnforcementAction) *UpdateRuleOptions {
	return &UpdateRuleOptions{
		RuleID:             core.StringPtr(ruleID),
		IfMatch:            core.StringPtr(ifMatch),
		Name:               core.StringPtr(name),
		Description:        core.StringPtr(description),
		Target:             target,
		RequiredConfig:     requiredConfig,
		EnforcementActions: enforcementActions,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *UpdateRuleOptions) SetRuleID(ruleID string) *UpdateRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateRuleOptions) SetIfMatch(ifMatch string) *UpdateRuleOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateRuleOptions) SetName(name string) *UpdateRuleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateRuleOptions) SetDescription(description string) *UpdateRuleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTarget : Allow user to set Target
func (options *UpdateRuleOptions) SetTarget(target *TargetResource) *UpdateRuleOptions {
	options.Target = target
	return options
}

// SetRequiredConfig : Allow user to set RequiredConfig
func (options *UpdateRuleOptions) SetRequiredConfig(requiredConfig RuleRequiredConfigIntf) *UpdateRuleOptions {
	options.RequiredConfig = requiredConfig
	return options
}

// SetEnforcementActions : Allow user to set EnforcementActions
func (options *UpdateRuleOptions) SetEnforcementActions(enforcementActions []EnforcementAction) *UpdateRuleOptions {
	options.EnforcementActions = enforcementActions
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateRuleOptions) SetAccountID(accountID string) *UpdateRuleOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetRuleType : Allow user to set RuleType
func (options *UpdateRuleOptions) SetRuleType(ruleType string) *UpdateRuleOptions {
	options.RuleType = core.StringPtr(ruleType)
	return options
}

// SetLabels : Allow user to set Labels
func (options *UpdateRuleOptions) SetLabels(labels []string) *UpdateRuleOptions {
	options.Labels = labels
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateRuleOptions) SetTransactionID(transactionID string) *UpdateRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRuleOptions) SetHeaders(param map[string]string) *UpdateRuleOptions {
	options.Headers = param
	return options
}

// RuleConditionAndLvl2 : A condition with the `and` logical operator.
// This model "extends" RuleCondition
type RuleConditionAndLvl2 struct {
	Description *string `json:"description,omitempty"`

	And []RuleSingleProperty `json:"and" validate:"required"`
}

// NewRuleConditionAndLvl2 : Instantiate RuleConditionAndLvl2 (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionAndLvl2(and []RuleSingleProperty) (model *RuleConditionAndLvl2, err error) {
	model = &RuleConditionAndLvl2{
		And: and,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionAndLvl2) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionAndLvl2 unmarshals an instance of RuleConditionAndLvl2 from the specified map of raw messages.
func UnmarshalRuleConditionAndLvl2(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionAndLvl2)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleConditionOrLvl2 : A condition with the `or` logical operator.
// This model "extends" RuleCondition
type RuleConditionOrLvl2 struct {
	Description *string `json:"description,omitempty"`

	Or []RuleSingleProperty `json:"or" validate:"required"`
}

// NewRuleConditionOrLvl2 : Instantiate RuleConditionOrLvl2 (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionOrLvl2(or []RuleSingleProperty) (model *RuleConditionOrLvl2, err error) {
	model = &RuleConditionOrLvl2{
		Or: or,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionOrLvl2) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionOrLvl2 unmarshals an instance of RuleConditionOrLvl2 from the specified map of raw messages.
func UnmarshalRuleConditionOrLvl2(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionOrLvl2)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleConditionSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
// This model "extends" RuleCondition
type RuleConditionSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleConditionSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
const (
	RuleConditionSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleConditionSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleConditionSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleConditionSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleConditionSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleConditionSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleConditionSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleConditionSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleConditionSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleConditionSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleConditionSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleConditionSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleConditionSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleConditionSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleConditionSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleConditionSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleConditionSingleProperty : Instantiate RuleConditionSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionSingleProperty(property string, operator string) (model *RuleConditionSingleProperty, err error) {
	model = &RuleConditionSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionSingleProperty) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionSingleProperty unmarshals an instance of RuleConditionSingleProperty from the specified map of raw messages.
func UnmarshalRuleConditionSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigMultipleProperties : The requirements that must be met to determine the resource's level of compliance in accordance with the rule.
//
// Use logical operators (`and`/`or`) to define multiple property checks and conditions. To define requirements for a
// rule, list one or more property check objects in the `and` array. To add conditions to a property check, use `or`.
// Models which "extend" this model:
// - RuleRequiredConfigMultiplePropertiesConditionOr
// - RuleRequiredConfigMultiplePropertiesConditionAnd
// This model "extends" RuleRequiredConfig
type RuleRequiredConfigMultipleProperties struct {
	Description *string `json:"description,omitempty"`

	Or []RuleConditionIntf `json:"or,omitempty"`

	And []RuleConditionIntf `json:"and,omitempty"`
}

func (*RuleRequiredConfigMultipleProperties) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

type RuleRequiredConfigMultiplePropertiesIntf interface {
	RuleRequiredConfigIntf
	isaRuleRequiredConfigMultipleProperties() bool
}

func (*RuleRequiredConfigMultipleProperties) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultipleProperties unmarshals an instance of RuleRequiredConfigMultipleProperties from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultipleProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultipleProperties)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
// This model "extends" RuleRequiredConfig
type RuleRequiredConfigSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleRequiredConfigSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// To learn more, see the [docs](/docs/security-compliance?topic=security-compliance-what-is-rule#rule-operators).
const (
	RuleRequiredConfigSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleRequiredConfigSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleRequiredConfigSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleRequiredConfigSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleRequiredConfigSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleRequiredConfigSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleRequiredConfigSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleRequiredConfigSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleRequiredConfigSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleRequiredConfigSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleRequiredConfigSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleRequiredConfigSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleRequiredConfigSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleRequiredConfigSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleRequiredConfigSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleRequiredConfigSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleRequiredConfigSingleProperty : Instantiate RuleRequiredConfigSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigSingleProperty(property string, operator string) (model *RuleRequiredConfigSingleProperty, err error) {
	model = &RuleRequiredConfigSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigSingleProperty) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigSingleProperty unmarshals an instance of RuleRequiredConfigSingleProperty from the specified map of raw messages.
func UnmarshalRuleRequiredConfigSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigMultiplePropertiesConditionAnd : A condition with the `and` logical operator.
// This model "extends" RuleRequiredConfigMultipleProperties
type RuleRequiredConfigMultiplePropertiesConditionAnd struct {
	Description *string `json:"description,omitempty"`

	And []RuleConditionIntf `json:"and" validate:"required"`
}

// NewRuleRequiredConfigMultiplePropertiesConditionAnd : Instantiate RuleRequiredConfigMultiplePropertiesConditionAnd (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigMultiplePropertiesConditionAnd(and []RuleConditionIntf) (model *RuleRequiredConfigMultiplePropertiesConditionAnd, err error) {
	model = &RuleRequiredConfigMultiplePropertiesConditionAnd{
		And: and,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigMultiplePropertiesConditionAnd) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

func (*RuleRequiredConfigMultiplePropertiesConditionAnd) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultiplePropertiesConditionAnd unmarshals an instance of RuleRequiredConfigMultiplePropertiesConditionAnd from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultiplePropertiesConditionAnd(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultiplePropertiesConditionAnd)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigMultiplePropertiesConditionOr : A condition with the `or` logical operator.
// This model "extends" RuleRequiredConfigMultipleProperties
type RuleRequiredConfigMultiplePropertiesConditionOr struct {
	Description *string `json:"description,omitempty"`

	Or []RuleConditionIntf `json:"or" validate:"required"`
}

// NewRuleRequiredConfigMultiplePropertiesConditionOr : Instantiate RuleRequiredConfigMultiplePropertiesConditionOr (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigMultiplePropertiesConditionOr(or []RuleConditionIntf) (model *RuleRequiredConfigMultiplePropertiesConditionOr, err error) {
	model = &RuleRequiredConfigMultiplePropertiesConditionOr{
		Or: or,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigMultiplePropertiesConditionOr) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

func (*RuleRequiredConfigMultiplePropertiesConditionOr) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultiplePropertiesConditionOr unmarshals an instance of RuleRequiredConfigMultiplePropertiesConditionOr from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultiplePropertiesConditionOr(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultiplePropertiesConditionOr)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

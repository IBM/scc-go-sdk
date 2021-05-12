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
 * IBM OpenAPI SDK Code Generator Version: 3.31.0-902c9336-20210504-161156
 */

// Package posturemanagementv1 : Operations and models for the PostureManagementV1 service
package posturemanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	common "github.com/ibm-cloud-security/scc-go-sdk/common"
)

// PostureManagementV1 : With IBM Cloud® Security and Compliance Center, you can embed checks into your every day
// workflows to help manage your current security and compliance posture. By monitoring for risks, you can identify
// security vulnerabilities and quickly work to mitigate the impact.
//
// Version: 1.0.0
type PostureManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "posture_management"

// PostureManagementV1Options : Service options
type PostureManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPostureManagementV1UsingExternalConfig : constructs an instance of PostureManagementV1 with passed in options and external configuration.
func NewPostureManagementV1UsingExternalConfig(options *PostureManagementV1Options) (postureManagement *PostureManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	postureManagement, err = NewPostureManagementV1(options)
	if err != nil {
		return
	}

	err = postureManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = postureManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewPostureManagementV1 : constructs an instance of PostureManagementV1 with passed in options.
func NewPostureManagementV1(options *PostureManagementV1Options) (service *PostureManagementV1, err error) {
	serviceOptions := &core.ServiceOptions{
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

	service = &PostureManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "postureManagement" suitable for processing requests.
func (postureManagement *PostureManagementV1) Clone() *PostureManagementV1 {
	if core.IsNil(postureManagement) {
		return nil
	}
	clone := *postureManagement
	clone.Service = postureManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (postureManagement *PostureManagementV1) SetServiceURL(url string) error {
	return postureManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (postureManagement *PostureManagementV1) GetServiceURL() string {
	return postureManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (postureManagement *PostureManagementV1) SetDefaultHeaders(headers http.Header) {
	postureManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (postureManagement *PostureManagementV1) SetEnableGzipCompression(enableGzip bool) {
	postureManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (postureManagement *PostureManagementV1) GetEnableGzipCompression() bool {
	return postureManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (postureManagement *PostureManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	postureManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (postureManagement *PostureManagementV1) DisableRetries() {
	postureManagement.Service.DisableRetries()
}

// ListScans : List scans
// List all of the scans that are available in your account. To view a specific scans, you can filter by scope or by
// profile.
func (postureManagement *PostureManagementV1) ListScans(listScansOptions *ListScansOptions) (result *ScansList, response *core.DetailedResponse, err error) {
	return postureManagement.ListScansWithContext(context.Background(), listScansOptions)
}

// ListScansWithContext is an alternate form of the ListScans method which supports a Context parameter
func (postureManagement *PostureManagementV1) ListScansWithContext(ctx context.Context, listScansOptions *ListScansOptions) (result *ScansList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listScansOptions, "listScansOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listScansOptions, "listScansOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scans/validations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listScansOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "ListScans")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listScansOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listScansOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listScansOptions.AccountID))
	if listScansOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listScansOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScansList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateValidation : Initiate a validation scan
// Validation scans determine a specified scope's adherence to regulatory controls by validating the configuration of
// the resources in your scope to the attached profile. To initiate a scan, you must have configured a collector,
// provided credentials, and completed both a fact collection and discovery scan. [Learn
// more](/docs/security-compliance?topic=security-compliance-schedule-scan).
func (postureManagement *PostureManagementV1) CreateValidation(createValidationOptions *CreateValidationOptions) (result *Result, response *core.DetailedResponse, err error) {
	return postureManagement.CreateValidationWithContext(context.Background(), createValidationOptions)
}

// CreateValidationWithContext is an alternate form of the CreateValidation method which supports a Context parameter
func (postureManagement *PostureManagementV1) CreateValidationWithContext(ctx context.Context, createValidationOptions *CreateValidationOptions) (result *Result, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createValidationOptions, "createValidationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createValidationOptions, "createValidationOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scans/validations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createValidationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "CreateValidation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createValidationOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createValidationOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*createValidationOptions.AccountID))

	body := make(map[string]interface{})
	if createValidationOptions.ScopeID != nil {
		body["scope_id"] = createValidationOptions.ScopeID
	}
	if createValidationOptions.ProfileID != nil {
		body["profile_id"] = createValidationOptions.ProfileID
	}
	if createValidationOptions.GroupProfileID != nil {
		body["group_profile_id"] = createValidationOptions.GroupProfileID
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
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProfiles : List profiles
// List all of the profiles that are available in your account. To view a specific profile, you can filter by name.
func (postureManagement *PostureManagementV1) ListProfiles(listProfilesOptions *ListProfilesOptions) (result *ProfilesList, response *core.DetailedResponse, err error) {
	return postureManagement.ListProfilesWithContext(context.Background(), listProfilesOptions)
}

// ListProfilesWithContext is an alternate form of the ListProfiles method which supports a Context parameter
func (postureManagement *PostureManagementV1) ListProfilesWithContext(ctx context.Context, listProfilesOptions *ListProfilesOptions) (result *ProfilesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listProfilesOptions, "listProfilesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listProfilesOptions, "listProfilesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "ListProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listProfilesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listProfilesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listProfilesOptions.AccountID))
	if listProfilesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listProfilesOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfilesList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateScope : Create a scope
// Creating a scope lets you determine your security and compliance score across a specific area of your business.
func (postureManagement *PostureManagementV1) CreateScope(createScopeOptions *CreateScopeOptions) (result *Scope, response *core.DetailedResponse, err error) {
	return postureManagement.CreateScopeWithContext(context.Background(), createScopeOptions)
}

// CreateScopeWithContext is an alternate form of the CreateScope method which supports a Context parameter
func (postureManagement *PostureManagementV1) CreateScopeWithContext(ctx context.Context, createScopeOptions *CreateScopeOptions) (result *Scope, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createScopeOptions, "createScopeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createScopeOptions, "createScopeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scopes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createScopeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "CreateScope")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createScopeOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createScopeOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*createScopeOptions.AccountID))

	body := make(map[string]interface{})
	if createScopeOptions.ScopeName != nil {
		body["scope_name"] = createScopeOptions.ScopeName
	}
	if createScopeOptions.ScopeDescription != nil {
		body["scope_description"] = createScopeOptions.ScopeDescription
	}
	if createScopeOptions.CollectorIds != nil {
		body["collector_ids"] = createScopeOptions.CollectorIds
	}
	if createScopeOptions.CredentialID != nil {
		body["credential_id"] = createScopeOptions.CredentialID
	}
	if createScopeOptions.EnvironmentType != nil {
		body["environment_type"] = createScopeOptions.EnvironmentType
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
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScope)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListScopes : List scopes
// List all of the scopes that are available in your account. To view a specific scope, you can filter by name.
func (postureManagement *PostureManagementV1) ListScopes(listScopesOptions *ListScopesOptions) (result *ScopesList, response *core.DetailedResponse, err error) {
	return postureManagement.ListScopesWithContext(context.Background(), listScopesOptions)
}

// ListScopesWithContext is an alternate form of the ListScopes method which supports a Context parameter
func (postureManagement *PostureManagementV1) ListScopesWithContext(ctx context.Context, listScopesOptions *ListScopesOptions) (result *ScopesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listScopesOptions, "listScopesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listScopesOptions, "listScopesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scopes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listScopesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "ListScopes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listScopesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listScopesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listScopesOptions.AccountID))
	if listScopesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listScopesOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScopesList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCollector : Create a collector
// Create a collector to fetch the configuration of your cloud environment and perform a validation afterwards.
func (postureManagement *PostureManagementV1) CreateCollector(createCollectorOptions *CreateCollectorOptions) (result *Collector, response *core.DetailedResponse, err error) {
	return postureManagement.CreateCollectorWithContext(context.Background(), createCollectorOptions)
}

// CreateCollectorWithContext is an alternate form of the CreateCollector method which supports a Context parameter
func (postureManagement *PostureManagementV1) CreateCollectorWithContext(ctx context.Context, createCollectorOptions *CreateCollectorOptions) (result *Collector, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCollectorOptions, "createCollectorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectorOptions, "createCollectorOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/collectors`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCollectorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "CreateCollector")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createCollectorOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createCollectorOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*createCollectorOptions.AccountID))

	body := make(map[string]interface{})
	if createCollectorOptions.CollectorName != nil {
		body["collector_name"] = createCollectorOptions.CollectorName
	}
	if createCollectorOptions.CollectorDescription != nil {
		body["collector_description"] = createCollectorOptions.CollectorDescription
	}
	if createCollectorOptions.IsPublic != nil {
		body["is_public"] = createCollectorOptions.IsPublic
	}
	if createCollectorOptions.ManagedBy != nil {
		body["managed_by"] = createCollectorOptions.ManagedBy
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
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCollector)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCredential : Create a credential
// Credential description.
func (postureManagement *PostureManagementV1) CreateCredential(createCredentialOptions *CreateCredentialOptions) (result *Credential, response *core.DetailedResponse, err error) {
	return postureManagement.CreateCredentialWithContext(context.Background(), createCredentialOptions)
}

// CreateCredentialWithContext is an alternate form of the CreateCredential method which supports a Context parameter
func (postureManagement *PostureManagementV1) CreateCredentialWithContext(ctx context.Context, createCredentialOptions *CreateCredentialOptions) (result *Credential, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCredentialOptions, "createCredentialOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCredentialOptions, "createCredentialOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/credentials`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCredentialOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "CreateCredential")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createCredentialOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createCredentialOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*createCredentialOptions.AccountID))

	builder.AddFormData("credential_data_file", "filename",
		"application/json", createCredentialOptions.CredentialDataFile)
	if createCredentialOptions.PemFile != nil {
		builder.AddFormData("pem_file", "filename",
			"text/plain", createCredentialOptions.PemFile)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCredential)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ApplicabilityCriteria : The criteria that defines how a profile applies.
type ApplicabilityCriteria struct {
	// A list of environments that a profile can be applied to.
	Environment []string `json:"environment,omitempty"`

	// A list of resources that a profile can be used with.
	Resource []string `json:"resource,omitempty"`

	// The type of environment that a profile is able to be applied to.
	EnvironmentCategory []string `json:"environment_category,omitempty"`

	// The type of resource that a profile is able to be applied to.
	ResourceCategory []string `json:"resource_category,omitempty"`

	// The resource type that the profile applies to.
	ResourceType []string `json:"resource_type,omitempty"`

	// The software that the profile applies to.
	SoftwareDetails interface{} `json:"software_details,omitempty"`

	// The operatoring system that the profile applies to.
	OsDetails interface{} `json:"os_details,omitempty"`

	// Any additional details about the profile.
	AdditionalDetails interface{} `json:"additional_details,omitempty"`

	// The type of environment that your scope is targeted to.
	EnvironmentCategoryDescription map[string]string `json:"environment_category_description,omitempty"`

	// The environment that your scope is targeted to.
	EnvironmentDescription map[string]string `json:"environment_description,omitempty"`

	// The type of resource that your scope is targeted to.
	ResourceCategoryDescription map[string]string `json:"resource_category_description,omitempty"`

	// A further classification of the type of resource that your scope is targeted to.
	ResourceTypeDescription map[string]string `json:"resource_type_description,omitempty"`

	// The resource that is scanned as part of your scope.
	ResourceDescription map[string]string `json:"resource_description,omitempty"`
}

// UnmarshalApplicabilityCriteria unmarshals an instance of ApplicabilityCriteria from the specified map of raw messages.
func UnmarshalApplicabilityCriteria(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicabilityCriteria)
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource", &obj.Resource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_category", &obj.EnvironmentCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_category", &obj.ResourceCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "software_details", &obj.SoftwareDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "os_details", &obj.OsDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "additional_details", &obj.AdditionalDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_category_description", &obj.EnvironmentCategoryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_description", &obj.EnvironmentDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_category_description", &obj.ResourceCategoryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type_description", &obj.ResourceTypeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_description", &obj.ResourceDescription)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Collector : Collector struct
type Collector struct {
	// An identifier of the collector.
	CollectorID *string `json:"collector_id,omitempty"`
}

// UnmarshalCollector unmarshals an instance of Collector from the specified map of raw messages.
func UnmarshalCollector(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Collector)
	err = core.UnmarshalPrimitive(m, "collector_id", &obj.CollectorID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollectorOptions : The CreateCollector options.
type CreateCollectorOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// A unique name for your collector.
	CollectorName *string

	// A detailed description of the collector.
	CollectorDescription *string

	// Determines whether the collector endpoint is accessible on a public network.
	//
	// If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the
	// collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network.
	IsPublic *bool

	// Determines whether the collector is IBM or customer-managed.
	//
	// Use `IBM` to allow Security and Compliance Center to create, install, and manage the collector on your behalf. The
	// collector is installed in an OpenShift cluster and approved automatically for use. Use `CUSTOMER` if you would like
	// to install the collector by using your own virtual machine. For more information, check out the
	// [docs](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-collector).
	ManagedBy *string

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCollectorOptions.ManagedBy property.
// Determines whether the collector is IBM or customer-managed.
//
// Use `IBM` to allow Security and Compliance Center to create, install, and manage the collector on your behalf. The
// collector is installed in an OpenShift cluster and approved automatically for use. Use `CUSTOMER` if you would like
// to install the collector by using your own virtual machine. For more information, check out the
// [docs](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-collector).
const (
	CreateCollectorOptionsManagedByCustomerConst = "CUSTOMER"
	CreateCollectorOptionsManagedByIBMConst = "IBM"
)

// NewCreateCollectorOptions : Instantiate CreateCollectorOptions
func (*PostureManagementV1) NewCreateCollectorOptions(accountID string) *CreateCollectorOptions {
	return &CreateCollectorOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateCollectorOptions) SetAccountID(accountID string) *CreateCollectorOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetCollectorName : Allow user to set CollectorName
func (options *CreateCollectorOptions) SetCollectorName(collectorName string) *CreateCollectorOptions {
	options.CollectorName = core.StringPtr(collectorName)
	return options
}

// SetCollectorDescription : Allow user to set CollectorDescription
func (options *CreateCollectorOptions) SetCollectorDescription(collectorDescription string) *CreateCollectorOptions {
	options.CollectorDescription = core.StringPtr(collectorDescription)
	return options
}

// SetIsPublic : Allow user to set IsPublic
func (options *CreateCollectorOptions) SetIsPublic(isPublic bool) *CreateCollectorOptions {
	options.IsPublic = core.BoolPtr(isPublic)
	return options
}

// SetManagedBy : Allow user to set ManagedBy
func (options *CreateCollectorOptions) SetManagedBy(managedBy string) *CreateCollectorOptions {
	options.ManagedBy = core.StringPtr(managedBy)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateCollectorOptions) SetTransactionID(transactionID string) *CreateCollectorOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectorOptions) SetHeaders(param map[string]string) *CreateCollectorOptions {
	options.Headers = param
	return options
}

// CreateCredentialOptions : The CreateCredential options.
type CreateCredentialOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	CredentialDataFile io.ReadCloser `validate:"required"`

	PemFile io.ReadCloser

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCredentialOptions : Instantiate CreateCredentialOptions
func (*PostureManagementV1) NewCreateCredentialOptions(accountID string, credentialDataFile io.ReadCloser) *CreateCredentialOptions {
	return &CreateCredentialOptions{
		AccountID: core.StringPtr(accountID),
		CredentialDataFile: credentialDataFile,
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateCredentialOptions) SetAccountID(accountID string) *CreateCredentialOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetCredentialDataFile : Allow user to set CredentialDataFile
func (options *CreateCredentialOptions) SetCredentialDataFile(credentialDataFile io.ReadCloser) *CreateCredentialOptions {
	options.CredentialDataFile = credentialDataFile
	return options
}

// SetPemFile : Allow user to set PemFile
func (options *CreateCredentialOptions) SetPemFile(pemFile io.ReadCloser) *CreateCredentialOptions {
	options.PemFile = pemFile
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateCredentialOptions) SetTransactionID(transactionID string) *CreateCredentialOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCredentialOptions) SetHeaders(param map[string]string) *CreateCredentialOptions {
	options.Headers = param
	return options
}

// CreateScopeOptions : The CreateScope options.
type CreateScopeOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// A unique name for your scope.
	ScopeName *string

	// A detailed description of the scope.
	ScopeDescription *string

	// The unique IDs of the collectors that are attached to the scope.
	CollectorIds []string

	// The unique identifier of the credential.
	CredentialID *string

	// The environment that the scope is targeted to.
	EnvironmentType *string

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateScopeOptions.EnvironmentType property.
// The environment that the scope is targeted to.
const (
	CreateScopeOptionsEnvironmentTypeAwsConst = "aws"
	CreateScopeOptionsEnvironmentTypeAzureConst = "azure"
	CreateScopeOptionsEnvironmentTypeGcpConst = "gcp"
	CreateScopeOptionsEnvironmentTypeHostedConst = "hosted"
	CreateScopeOptionsEnvironmentTypeIBMConst = "ibm"
	CreateScopeOptionsEnvironmentTypeOnPremiseConst = "on_premise"
	CreateScopeOptionsEnvironmentTypeOpenstackConst = "openstack"
	CreateScopeOptionsEnvironmentTypeServicesConst = "services"
)

// NewCreateScopeOptions : Instantiate CreateScopeOptions
func (*PostureManagementV1) NewCreateScopeOptions(accountID string) *CreateScopeOptions {
	return &CreateScopeOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateScopeOptions) SetAccountID(accountID string) *CreateScopeOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetScopeName : Allow user to set ScopeName
func (options *CreateScopeOptions) SetScopeName(scopeName string) *CreateScopeOptions {
	options.ScopeName = core.StringPtr(scopeName)
	return options
}

// SetScopeDescription : Allow user to set ScopeDescription
func (options *CreateScopeOptions) SetScopeDescription(scopeDescription string) *CreateScopeOptions {
	options.ScopeDescription = core.StringPtr(scopeDescription)
	return options
}

// SetCollectorIds : Allow user to set CollectorIds
func (options *CreateScopeOptions) SetCollectorIds(collectorIds []string) *CreateScopeOptions {
	options.CollectorIds = collectorIds
	return options
}

// SetCredentialID : Allow user to set CredentialID
func (options *CreateScopeOptions) SetCredentialID(credentialID string) *CreateScopeOptions {
	options.CredentialID = core.StringPtr(credentialID)
	return options
}

// SetEnvironmentType : Allow user to set EnvironmentType
func (options *CreateScopeOptions) SetEnvironmentType(environmentType string) *CreateScopeOptions {
	options.EnvironmentType = core.StringPtr(environmentType)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateScopeOptions) SetTransactionID(transactionID string) *CreateScopeOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateScopeOptions) SetHeaders(param map[string]string) *CreateScopeOptions {
	options.Headers = param
	return options
}

// CreateValidationOptions : The CreateValidation options.
type CreateValidationOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique ID of the scope.
	ScopeID *string

	// The unique ID of the profile.
	ProfileID *string

	// The ID of the profile group.
	GroupProfileID *string

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateValidationOptions : Instantiate CreateValidationOptions
func (*PostureManagementV1) NewCreateValidationOptions(accountID string) *CreateValidationOptions {
	return &CreateValidationOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateValidationOptions) SetAccountID(accountID string) *CreateValidationOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetScopeID : Allow user to set ScopeID
func (options *CreateValidationOptions) SetScopeID(scopeID string) *CreateValidationOptions {
	options.ScopeID = core.StringPtr(scopeID)
	return options
}

// SetProfileID : Allow user to set ProfileID
func (options *CreateValidationOptions) SetProfileID(profileID string) *CreateValidationOptions {
	options.ProfileID = core.StringPtr(profileID)
	return options
}

// SetGroupProfileID : Allow user to set GroupProfileID
func (options *CreateValidationOptions) SetGroupProfileID(groupProfileID string) *CreateValidationOptions {
	options.GroupProfileID = core.StringPtr(groupProfileID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateValidationOptions) SetTransactionID(transactionID string) *CreateValidationOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateValidationOptions) SetHeaders(param map[string]string) *CreateValidationOptions {
	options.Headers = param
	return options
}

// Credential : The details of a credential.
type Credential struct {
	// The unique ID of the credential.
	CredentialID *string `json:"credential_id,omitempty"`

	// The unique name of the credential.
	CredentialName *string `json:"credential_name,omitempty"`

	// The creation time of the credential.
	CreatedTime *strfmt.DateTime `json:"created_time,omitempty"`
}

// UnmarshalCredential unmarshals an instance of Credential from the specified map of raw messages.
func UnmarshalCredential(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Credential)
	err = core.UnmarshalPrimitive(m, "credential_id", &obj.CredentialID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credential_name", &obj.CredentialName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListProfilesOptions : The ListProfiles options.
type ListProfilesOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// The name of the profile.
	Name *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProfilesOptions : Instantiate ListProfilesOptions
func (*PostureManagementV1) NewListProfilesOptions(accountID string) *ListProfilesOptions {
	return &ListProfilesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListProfilesOptions) SetAccountID(accountID string) *ListProfilesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListProfilesOptions) SetTransactionID(transactionID string) *ListProfilesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetName : Allow user to set Name
func (options *ListProfilesOptions) SetName(name string) *ListProfilesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListProfilesOptions) SetHeaders(param map[string]string) *ListProfilesOptions {
	options.Headers = param
	return options
}

// ListScansOptions : The ListScans options.
type ListScansOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// The name of the scan.
	Name *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListScansOptions : Instantiate ListScansOptions
func (*PostureManagementV1) NewListScansOptions(accountID string) *ListScansOptions {
	return &ListScansOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListScansOptions) SetAccountID(accountID string) *ListScansOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListScansOptions) SetTransactionID(transactionID string) *ListScansOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetName : Allow user to set Name
func (options *ListScansOptions) SetName(name string) *ListScansOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListScansOptions) SetHeaders(param map[string]string) *ListScansOptions {
	options.Headers = param
	return options
}

// ListScopesOptions : The ListScopes options.
type ListScopesOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request.
	TransactionID *string

	// The name of the scope.
	Name *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListScopesOptions : Instantiate ListScopesOptions
func (*PostureManagementV1) NewListScopesOptions(accountID string) *ListScopesOptions {
	return &ListScopesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListScopesOptions) SetAccountID(accountID string) *ListScopesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListScopesOptions) SetTransactionID(transactionID string) *ListScopesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetName : Allow user to set Name
func (options *ListScopesOptions) SetName(name string) *ListScopesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListScopesOptions) SetHeaders(param map[string]string) *ListScopesOptions {
	options.Headers = param
	return options
}

// Profile : Profile.
type Profile struct {
	// The name of the profile.
	Name *string `json:"name,omitempty"`

	// A description of the profile.
	Description *string `json:"description,omitempty"`

	// The version of the profile.
	Version *int64 `json:"version,omitempty"`

	// The user who created the profile.
	CreatedBy *string `json:"created_by,omitempty"`

	// The user who last modified the profile.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// A reason that you want to delete a profile.
	ReasonForDelete *string `json:"reason_for_delete,omitempty"`

	// The criteria that defines how a profile applies.
	ApplicabilityCriteria *ApplicabilityCriteria `json:"applicability_criteria,omitempty"`

	// An auto-generated unique identifying number of the profile.
	ProfileID *string `json:"profile_id,omitempty"`

	// The base profile that the controls are pulled from.
	BaseProfile *string `json:"base_profile,omitempty"`

	// The type of profile.
	ProfileType *string `json:"profile_type,omitempty"`

	// The time that the profile was created in UTC.
	CreatedTime *strfmt.DateTime `json:"created_time,omitempty"`

	// The time that the profile was most recently modified in UTC.
	ModifiedTime *strfmt.DateTime `json:"modified_time,omitempty"`

	// The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// Constants associated with the Profile.ProfileType property.
// The type of profile.
const (
	ProfileProfileTypeCustomConst = "custom"
	ProfileProfileTypePredefinedConst = "predefined"
	ProfileProfileTypeTemplateGroupConst = "template_group"
)

// UnmarshalProfile unmarshals an instance of Profile from the specified map of raw messages.
func UnmarshalProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Profile)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason_for_delete", &obj.ReasonForDelete)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "applicability_criteria", &obj.ApplicabilityCriteria, UnmarshalApplicabilityCriteria)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_profile", &obj.BaseProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_type", &obj.ProfileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_time", &obj.ModifiedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfilesList : A list of profiles.
type ProfilesList struct {
	// Profiles.
	Profiles []Profile `json:"profiles,omitempty"`
}

// UnmarshalProfilesList unmarshals an instance of ProfilesList from the specified map of raw messages.
func UnmarshalProfilesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfilesList)
	err = core.UnmarshalModel(m, "profiles", &obj.Profiles, UnmarshalProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Result : Result.
type Result struct {
	// Result.
	Result *bool `json:"result,omitempty"`

	// A message is returned.
	Message *string `json:"message,omitempty"`
}

// UnmarshalResult unmarshals an instance of Result from the specified map of raw messages.
func UnmarshalResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Result)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
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

// Scan : Scan.
type Scan struct {
	// An auto-generated unique identifier for the scan.
	ScanID *string `json:"scan_id,omitempty"`

	// An auto-generated unique identifier for discovery.
	DiscoverID *string `json:"discover_id,omitempty"`

	// The status of the collector as it completes a scan.
	Status *string `json:"status,omitempty"`

	// The current status of the collector.
	StatusMessage *string `json:"status_message,omitempty"`
}

// Constants associated with the Scan.Status property.
// The status of the collector as it completes a scan.
const (
	ScanStatusAbortTaskRequestCompletedConst = "abort_task_request_completed"
	ScanStatusAbortTaskRequestFailedConst = "abort_task_request_failed"
	ScanStatusAbortTaskRequestReceivedConst = "abort_task_request_received"
	ScanStatusControllerAbortedConst = "controller_aborted"
	ScanStatusDiscoveryCompletedConst = "discovery_completed"
	ScanStatusDiscoveryInProgressConst = "discovery_in_progress"
	ScanStatusDiscoveryResultPostedNoErrorConst = "discovery_result_posted_no_error"
	ScanStatusDiscoveryResultPostedWithErrorConst = "discovery_result_posted_with_error"
	ScanStatusDiscoveryStartedConst = "discovery_started"
	ScanStatusErrorInAbortTaskRequestConst = "error_in_abort_task_request"
	ScanStatusErrorInDiscoveryConst = "error_in_discovery"
	ScanStatusErrorInFactCollectionConst = "error_in_fact_collection"
	ScanStatusErrorInFactValidationConst = "error_in_fact_validation"
	ScanStatusErrorInInventoryConst = "error_in_inventory"
	ScanStatusErrorInRemediationConst = "error_in_remediation"
	ScanStatusErrorInValidationConst = "error_in_validation"
	ScanStatusFactCollectionCompletedConst = "fact_collection_completed"
	ScanStatusFactCollectionInProgressConst = "fact_collection_in_progress"
	ScanStatusFactCollectionStartedConst = "fact_collection_started"
	ScanStatusFactValidationCompletedConst = "fact_validation_completed"
	ScanStatusFactValidationInProgressConst = "fact_validation_in_progress"
	ScanStatusFactValidationStartedConst = "fact_validation_started"
	ScanStatusGatewayAbortedConst = "gateway_aborted"
	ScanStatusInventoryCompletedConst = "inventory_completed"
	ScanStatusInventoryCompletedWithErrorConst = "inventory_completed_with_error"
	ScanStatusInventoryInProgressConst = "inventory_in_progress"
	ScanStatusInventoryStartedConst = "inventory_started"
	ScanStatusNotAcceptedConst = "not_accepted"
	ScanStatusPendingConst = "pending"
	ScanStatusRemediationCompletedConst = "remediation_completed"
	ScanStatusRemediationInProgressConst = "remediation_in_progress"
	ScanStatusRemediationStartedConst = "remediation_started"
	ScanStatusSentToCollectorConst = "sent_to_collector"
	ScanStatusUserAbortedConst = "user_aborted"
	ScanStatusValidationCompletedConst = "validation_completed"
	ScanStatusValidationInProgressConst = "validation_in_progress"
	ScanStatusValidationResultPostedNoErrorConst = "validation_result_posted_no_error"
	ScanStatusValidationResultPostedWithErrorConst = "validation_result_posted_with_error"
	ScanStatusValidationStartedConst = "validation_started"
	ScanStatusWaitingForRefineConst = "waiting_for_refine"
)

// UnmarshalScan unmarshals an instance of Scan from the specified map of raw messages.
func UnmarshalScan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scan)
	err = core.UnmarshalPrimitive(m, "scan_id", &obj.ScanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "discover_id", &obj.DiscoverID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_message", &obj.StatusMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScanResult : The result of a scan.
type ScanResult struct {
	// The number of goals that passed the scan.
	GoalsPassCount *int64 `json:"goals_pass_count,omitempty"`

	// The number of _Unable to perform_ (u2p) goals. A goal is evaluated as 'Unable to perform' when information about its
	// associated resource can't be collected.
	GoalsU2pCount *int64 `json:"goals_u2p_count,omitempty"`

	// The number of _Not applicable_ (na) goals. A goal is evaluated as 'Not applicable' when its associated resource
	// can't be found.
	GoalsNaCount *int64 `json:"goals_na_count,omitempty"`

	// The number of goals that failed the scan.
	GoalsFailCount *int64 `json:"goals_fail_count,omitempty"`

	// The total number of goals that were included in the scan.
	GoalsTotalCount *int64 `json:"goals_total_count,omitempty"`

	// The number of controls that passed the scan.
	ControlsPassCount *int64 `json:"controls_pass_count,omitempty"`

	// The number of controls that failed the scan.
	ControlsFailCount *int64 `json:"controls_fail_count,omitempty"`

	// The number of _Not applicable_ (na) controls. A control is evaluated as 'Not applicable' when its associated
	// resource can't be found.
	ControlsNaCount *int64 `json:"controls_na_count,omitempty"`

	// The number of _Unable to perform_ (u2p) controls. A control is evaluated as 'Unable to perform' when information
	// about its associated resource can't be collected.
	ControlsU2pCount *int64 `json:"controls_u2p_count,omitempty"`

	// The total number of controls that were included in the scan.
	ControlsTotalCount *int64 `json:"controls_total_count,omitempty"`
}

// UnmarshalScanResult unmarshals an instance of ScanResult from the specified map of raw messages.
func UnmarshalScanResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScanResult)
	err = core.UnmarshalPrimitive(m, "goals_pass_count", &obj.GoalsPassCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "goals_u2p_count", &obj.GoalsU2pCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "goals_na_count", &obj.GoalsNaCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "goals_fail_count", &obj.GoalsFailCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "goals_total_count", &obj.GoalsTotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_pass_count", &obj.ControlsPassCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_fail_count", &obj.ControlsFailCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_na_count", &obj.ControlsNaCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_u2p_count", &obj.ControlsU2pCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_total_count", &obj.ControlsTotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Scans : Scans struct
type Scans struct {
	// The ID of the scan.
	ScanID *string `json:"scan_id,omitempty"`

	// The name of the scan.
	ScanName *string `json:"scan_name,omitempty"`

	// The ID of the scan.
	ScopeID *string `json:"scope_id,omitempty"`

	// The name of the scope.
	ScopeName *string `json:"scope_name,omitempty"`

	// The ID of the profile.
	ProfileID *string `json:"profile_id,omitempty"`

	// The name of the profile.
	ProfileName *string `json:"profile_name,omitempty"`

	// The type of profile. To learn more about profile types, check out the [docs]
	// (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
	ProfileType *string `json:"profile_type,omitempty"`

	// The group ID of profile.
	ProfileGroupID *string `json:"profile_group_id,omitempty"`

	// The group name of the profile.
	ProfileGroupName *string `json:"profile_group_name,omitempty"`

	// The entity that ran the report.
	ReportRunBy *string `json:"report_run_by,omitempty"`

	// The amount of time that the scan took to complete.
	ScanTime *strfmt.DateTime `json:"scan_time,omitempty"`

	// The date and time the scan was run.
	StartedTime *strfmt.DateTime `json:"started_time,omitempty"`

	// The date and time the scan completed.
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// The result of a scan.
	Result *ScanResult `json:"result,omitempty"`
}

// Constants associated with the Scans.ProfileType property.
// The type of profile. To learn more about profile types, check out the [docs]
// (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
const (
	ScansProfileTypeAuthoredConst = "authored"
	ScansProfileTypeCustomConst = "custom"
	ScansProfileTypeStandardConst = "standard"
	ScansProfileTypeStandardCertificateConst = "standard_certificate"
	ScansProfileTypeStandardCvConst = "standard_cv"
	ScansProfileTypeTemmplategroupConst = "temmplategroup"
)

// UnmarshalScans unmarshals an instance of Scans from the specified map of raw messages.
func UnmarshalScans(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scans)
	err = core.UnmarshalPrimitive(m, "scan_id", &obj.ScanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scan_name", &obj.ScanName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_name", &obj.ScopeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_name", &obj.ProfileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_type", &obj.ProfileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_group_id", &obj.ProfileGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_group_name", &obj.ProfileGroupName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_run_by", &obj.ReportRunBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scan_time", &obj.ScanTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_time", &obj.StartedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalScanResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScansList : A list of scans.
type ScansList struct {
	// The details of a scan.
	Scans []Scans `json:"scans,omitempty"`
}

// UnmarshalScansList unmarshals an instance of ScansList from the specified map of raw messages.
func UnmarshalScansList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScansList)
	err = core.UnmarshalModel(m, "scans", &obj.Scans, UnmarshalScans)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Scope : Create Scope Response.
type Scope struct {
	// A unique identifier for your scope.
	ScopeID *string `json:"scope_id,omitempty"`

	// A unique name for your scope.
	ScopeName *string `json:"scope_name,omitempty"`

	// A detailed description of the scope.
	ScopeDescription *string `json:"scope_description,omitempty"`

	// The unique IDs of the collectors that are attached to the scope.
	CollectorIds []string `json:"collector_ids,omitempty"`

	// The unique identifier of the credential.
	CredentialID *string `json:"credential_id,omitempty"`

	// The environment that the scope is targeted to.
	EnvironmentType *string `json:"environment_type,omitempty"`

	// The time that the scope was created in UTC.
	CreatedTime *strfmt.DateTime `json:"created_time,omitempty"`

	// The time that the scope was last modified in UTC.
	ModifiedTime *strfmt.DateTime `json:"modified_time,omitempty"`
}

// Constants associated with the Scope.EnvironmentType property.
// The environment that the scope is targeted to.
const (
	ScopeEnvironmentTypeAwsConst = "aws"
	ScopeEnvironmentTypeAzureConst = "azure"
	ScopeEnvironmentTypeGcpConst = "gcp"
	ScopeEnvironmentTypeHostedConst = "hosted"
	ScopeEnvironmentTypeIBMConst = "ibm"
	ScopeEnvironmentTypeOnPremiseConst = "on_premise"
	ScopeEnvironmentTypeOpenstackConst = "openstack"
	ScopeEnvironmentTypeServicesConst = "services"
)

// UnmarshalScope unmarshals an instance of Scope from the specified map of raw messages.
func UnmarshalScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scope)
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_name", &obj.ScopeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_description", &obj.ScopeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collector_ids", &obj.CollectorIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credential_id", &obj.CredentialID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_type", &obj.EnvironmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_time", &obj.ModifiedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScopeItem : Scope.
type ScopeItem struct {
	// A detailed description of the scope.
	Description *string `json:"description,omitempty"`

	// The user who created the scope.
	CreatedBy *string `json:"created_by,omitempty"`

	// The user who most recently modified the scope.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// An auto-generated unique identifier for the scope.
	ScopeID *string `json:"scope_id,omitempty"`

	// A unique name for your scope.
	Name *string `json:"name,omitempty"`

	// Indicates whether scope is enabled/disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The environment that the scope is targeted to.
	EnvironmentType *string `json:"environment_type,omitempty"`

	// The time that the scope was created in UTC.
	CreatedTime *strfmt.DateTime `json:"created_time,omitempty"`

	// The time that the scope was last modified in UTC.
	ModifiedTime *strfmt.DateTime `json:"modified_time,omitempty"`

	// The last type of scan that was run on the scope.
	LastScanType *string `json:"last_scan_type,omitempty"`

	// A description of the last scan type.
	LastScanTypeDescription *string `json:"last_scan_type_description,omitempty"`

	// The last time that a scan status for a scope was updated in UTC.
	LastScanStatusUpdatedTime *strfmt.DateTime `json:"last_scan_status_updated_time,omitempty"`

	// The unique IDs of the collectors that are attached to the scope.
	CollectorsID []string `json:"collectors_id,omitempty"`

	// A list of the scans that have been run on the scope.
	Scans []Scan `json:"scans,omitempty"`
}

// Constants associated with the ScopeItem.EnvironmentType property.
// The environment that the scope is targeted to.
const (
	ScopeItemEnvironmentTypeAwsConst = "aws"
	ScopeItemEnvironmentTypeAzureConst = "azure"
	ScopeItemEnvironmentTypeGcpConst = "gcp"
	ScopeItemEnvironmentTypeHostedConst = "hosted"
	ScopeItemEnvironmentTypeIBMConst = "ibm"
	ScopeItemEnvironmentTypeOnPremiseConst = "on_premise"
	ScopeItemEnvironmentTypeOpenstackConst = "openstack"
	ScopeItemEnvironmentTypeServicesConst = "services"
)

// Constants associated with the ScopeItem.LastScanType property.
// The last type of scan that was run on the scope.
const (
	ScopeItemLastScanTypeAbortTasksConst = "abort_tasks"
	ScopeItemLastScanTypeDiscoveryConst = "discovery"
	ScopeItemLastScanTypeEvidenceConst = "evidence"
	ScopeItemLastScanTypeFactCollectionConst = "fact_collection"
	ScopeItemLastScanTypeFactValidationConst = "fact_validation"
	ScopeItemLastScanTypeInventoryConst = "inventory"
	ScopeItemLastScanTypeRemediationConst = "remediation"
	ScopeItemLastScanTypeScriptConst = "script"
	ScopeItemLastScanTypeValidationConst = "validation"
)

// UnmarshalScopeItem unmarshals an instance of ScopeItem from the specified map of raw messages.
func UnmarshalScopeItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScopeItem)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_type", &obj.EnvironmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_time", &obj.ModifiedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_type", &obj.LastScanType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_type_description", &obj.LastScanTypeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_status_updated_time", &obj.LastScanStatusUpdatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collectors_id", &obj.CollectorsID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scans", &obj.Scans, UnmarshalScan)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScopesList : Scopes list.
type ScopesList struct {
	// Scopes.
	Scopes []ScopeItem `json:"scopes,omitempty"`
}

// UnmarshalScopesList unmarshals an instance of ScopesList from the specified map of raw messages.
func UnmarshalScopesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScopesList)
	err = core.UnmarshalModel(m, "scopes", &obj.Scopes, UnmarshalScopeItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

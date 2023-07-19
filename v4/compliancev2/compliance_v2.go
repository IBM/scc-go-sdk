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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.73.0-eeee85a9-20230607-165104
 */

// Package compliancev2 : Operations and models for the ComplianceV2 service
package compliancev2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/scc-go-sdk/v4/common"
	"github.com/go-openapi/strfmt"
)

// ComplianceV2 : Security and Compliance Center API reference.
//
// API Version: 2.0.0
type ComplianceV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.compliance.cloud.ibm.com/instances/edf9524f-406c-412c-acbb-ee371a5cabda/v3"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "compliance"

const ParameterizedServiceURL = "https://{region}.cloud.ibm.com/instances/{instance_id}/v3"

var defaultUrlVariables = map[string]string{
	"region":      "us-south.compliance",
	"instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda",
}

// ComplianceV2Options : Service options
type ComplianceV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewComplianceV2UsingExternalConfig : constructs an instance of ComplianceV2 with passed in options and external configuration.
func NewComplianceV2UsingExternalConfig(options *ComplianceV2Options) (compliance *ComplianceV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	compliance, err = NewComplianceV2(options)
	if err != nil {
		return
	}

	err = compliance.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = compliance.Service.SetServiceURL(options.URL)
	}
	return
}

// NewComplianceV2 : constructs an instance of ComplianceV2 with passed in options.
func NewComplianceV2(options *ComplianceV2Options) (service *ComplianceV2, err error) {
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

	service = &ComplianceV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "compliance" suitable for processing requests.
func (compliance *ComplianceV2) Clone() *ComplianceV2 {
	if core.IsNil(compliance) {
		return nil
	}
	clone := *compliance
	clone.Service = compliance.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (compliance *ComplianceV2) SetServiceURL(url string) error {
	return compliance.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (compliance *ComplianceV2) GetServiceURL() string {
	return compliance.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (compliance *ComplianceV2) SetDefaultHeaders(headers http.Header) {
	compliance.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (compliance *ComplianceV2) SetEnableGzipCompression(enableGzip bool) {
	compliance.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (compliance *ComplianceV2) GetEnableGzipCompression() bool {
	return compliance.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (compliance *ComplianceV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	compliance.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (compliance *ComplianceV2) DisableRetries() {
	compliance.Service.DisableRetries()
}

// ListControlLibraries : Get control libraries
// Retrieve all of the control libraries that are available in your account, including predefined, and custom libraries.
//
// With Security and Compliance Center, you can create a custom control library that is specific to your organization's
// needs.  You define the controls and specifications before you map previously created assessments. Each control has
// several specifications  and assessments that are mapped to it. A specification is a defined requirement that is
// specific to a component. An assessment, or several,  are mapped to each specification with a detailed evaluation that
// is done to check whether the specification is compliant. For more information, see [Creating custom
// libraries](/docs/security-compliance?topic=security-compliance-custom-library).
func (compliance *ComplianceV2) ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) (result *ControlLibraryCollection, response *core.DetailedResponse, err error) {
	return compliance.ListControlLibrariesWithContext(context.Background(), listControlLibrariesOptions)
}

// ListControlLibrariesWithContext is an alternate form of the ListControlLibraries method which supports a Context parameter
func (compliance *ComplianceV2) ListControlLibrariesWithContext(ctx context.Context, listControlLibrariesOptions *ListControlLibrariesOptions) (result *ControlLibraryCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listControlLibrariesOptions, "listControlLibrariesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/control_libraries`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listControlLibrariesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ListControlLibraries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listControlLibrariesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listControlLibrariesOptions.XCorrelationID))
	}
	if listControlLibrariesOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*listControlLibrariesOptions.XRequestID))
	}

	if listControlLibrariesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listControlLibrariesOptions.Limit))
	}
	if listControlLibrariesOptions.ControlLibraryType != nil {
		builder.AddQuery("control_library_type", fmt.Sprint(*listControlLibrariesOptions.ControlLibraryType))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomControlLibrary : Create a custom control library
// Create a custom control library that is specific to your organization's needs.
//
// With Security and Compliance Center, you can create a custom control library that is specific to your organization's
// needs.  You define the controls and specifications before you map previously created assessments. Each control has
// several specifications  and assessments that are mapped to it. A specification is a defined requirement that is
// specific to a component. An assessment, or several,  are mapped to each specification with a detailed evaluation that
// is done to check whether the specification is compliant. For more information, see [Creating custom
// libraries](/docs/security-compliance?topic=security-compliance-custom-library).
func (compliance *ComplianceV2) CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	return compliance.CreateCustomControlLibraryWithContext(context.Background(), createCustomControlLibraryOptions)
}

// CreateCustomControlLibraryWithContext is an alternate form of the CreateCustomControlLibrary method which supports a Context parameter
func (compliance *ComplianceV2) CreateCustomControlLibraryWithContext(ctx context.Context, createCustomControlLibraryOptions *CreateCustomControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomControlLibraryOptions, "createCustomControlLibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomControlLibraryOptions, "createCustomControlLibraryOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/control_libraries`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "CreateCustomControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createCustomControlLibraryOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createCustomControlLibraryOptions.XCorrelationID))
	}
	if createCustomControlLibraryOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*createCustomControlLibraryOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if createCustomControlLibraryOptions.ControlLibraryName != nil {
		body["control_library_name"] = createCustomControlLibraryOptions.ControlLibraryName
	}
	if createCustomControlLibraryOptions.ControlLibraryDescription != nil {
		body["control_library_description"] = createCustomControlLibraryOptions.ControlLibraryDescription
	}
	if createCustomControlLibraryOptions.ControlLibraryType != nil {
		body["control_library_type"] = createCustomControlLibraryOptions.ControlLibraryType
	}
	if createCustomControlLibraryOptions.Controls != nil {
		body["controls"] = createCustomControlLibraryOptions.Controls
	}
	if createCustomControlLibraryOptions.VersionGroupLabel != nil {
		body["version_group_label"] = createCustomControlLibraryOptions.VersionGroupLabel
	}
	if createCustomControlLibraryOptions.ControlLibraryVersion != nil {
		body["control_library_version"] = createCustomControlLibraryOptions.ControlLibraryVersion
	}
	if createCustomControlLibraryOptions.Latest != nil {
		body["latest"] = createCustomControlLibraryOptions.Latest
	}
	if createCustomControlLibraryOptions.ControlsCount != nil {
		body["controls_count"] = createCustomControlLibraryOptions.ControlsCount
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibrary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetControlLibrary : Get a control library
// View the details of a control library by specifying its ID.
//
// With Security and Compliance Center, you can create a custom control library that is specific to your organization's
// needs.  You define the controls and specifications before you map previously created assessments. Each control has
// several specifications  and assessments that are mapped to it. A specification is a defined requirement that is
// specific to a component. An assessment, or several,  are mapped to each specification with a detailed evaluation that
// is done to check whether the specification is compliant. For more information, see [Creating custom
// libraries](/docs/security-compliance?topic=security-compliance-custom-library).
func (compliance *ComplianceV2) GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	return compliance.GetControlLibraryWithContext(context.Background(), getControlLibraryOptions)
}

// GetControlLibraryWithContext is an alternate form of the GetControlLibrary method which supports a Context parameter
func (compliance *ComplianceV2) GetControlLibraryWithContext(ctx context.Context, getControlLibraryOptions *GetControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getControlLibraryOptions, "getControlLibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getControlLibraryOptions, "getControlLibraryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"control_libraries_id": *getControlLibraryOptions.ControlLibrariesID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "GetControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getControlLibraryOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getControlLibraryOptions.XCorrelationID))
	}
	if getControlLibraryOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*getControlLibraryOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibrary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceCustomControlLibrary : Update a control library
// Update a custom control library by providing the control library ID. You can find this ID in the Security and
// Compliance Center UI.
//
// With Security and Compliance Center, you can create and update a custom control library that is specific to your
// organization's needs.  You define the controls and specifications before you map previously created assessments. Each
// control has several specifications  and assessments that are mapped to it. For more information, see [Creating custom
// libraries](/docs/security-compliance?topic=security-compliance-custom-library).
func (compliance *ComplianceV2) ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	return compliance.ReplaceCustomControlLibraryWithContext(context.Background(), replaceCustomControlLibraryOptions)
}

// ReplaceCustomControlLibraryWithContext is an alternate form of the ReplaceCustomControlLibrary method which supports a Context parameter
func (compliance *ComplianceV2) ReplaceCustomControlLibraryWithContext(ctx context.Context, replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions) (result *ControlLibrary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceCustomControlLibraryOptions, "replaceCustomControlLibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceCustomControlLibraryOptions, "replaceCustomControlLibraryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"control_libraries_id": *replaceCustomControlLibraryOptions.ControlLibrariesID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceCustomControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ReplaceCustomControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceCustomControlLibraryOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*replaceCustomControlLibraryOptions.XCorrelationID))
	}
	if replaceCustomControlLibraryOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*replaceCustomControlLibraryOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if replaceCustomControlLibraryOptions.ID != nil {
		body["id"] = replaceCustomControlLibraryOptions.ID
	}
	if replaceCustomControlLibraryOptions.AccountID != nil {
		body["account_id"] = replaceCustomControlLibraryOptions.AccountID
	}
	if replaceCustomControlLibraryOptions.ControlLibraryName != nil {
		body["control_library_name"] = replaceCustomControlLibraryOptions.ControlLibraryName
	}
	if replaceCustomControlLibraryOptions.ControlLibraryDescription != nil {
		body["control_library_description"] = replaceCustomControlLibraryOptions.ControlLibraryDescription
	}
	if replaceCustomControlLibraryOptions.ControlLibraryType != nil {
		body["control_library_type"] = replaceCustomControlLibraryOptions.ControlLibraryType
	}
	if replaceCustomControlLibraryOptions.VersionGroupLabel != nil {
		body["version_group_label"] = replaceCustomControlLibraryOptions.VersionGroupLabel
	}
	if replaceCustomControlLibraryOptions.ControlLibraryVersion != nil {
		body["control_library_version"] = replaceCustomControlLibraryOptions.ControlLibraryVersion
	}
	if replaceCustomControlLibraryOptions.CreatedOn != nil {
		body["created_on"] = replaceCustomControlLibraryOptions.CreatedOn
	}
	if replaceCustomControlLibraryOptions.CreatedBy != nil {
		body["created_by"] = replaceCustomControlLibraryOptions.CreatedBy
	}
	if replaceCustomControlLibraryOptions.UpdatedOn != nil {
		body["updated_on"] = replaceCustomControlLibraryOptions.UpdatedOn
	}
	if replaceCustomControlLibraryOptions.UpdatedBy != nil {
		body["updated_by"] = replaceCustomControlLibraryOptions.UpdatedBy
	}
	if replaceCustomControlLibraryOptions.Latest != nil {
		body["latest"] = replaceCustomControlLibraryOptions.Latest
	}
	if replaceCustomControlLibraryOptions.HierarchyEnabled != nil {
		body["hierarchy_enabled"] = replaceCustomControlLibraryOptions.HierarchyEnabled
	}
	if replaceCustomControlLibraryOptions.ControlsCount != nil {
		body["controls_count"] = replaceCustomControlLibraryOptions.ControlsCount
	}
	if replaceCustomControlLibraryOptions.ControlParentsCount != nil {
		body["control_parents_count"] = replaceCustomControlLibraryOptions.ControlParentsCount
	}
	if replaceCustomControlLibraryOptions.Controls != nil {
		body["controls"] = replaceCustomControlLibraryOptions.Controls
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibrary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomControlLibrary : Delete a control library
// Delete a custom control library by providing the control library ID.  You can find this ID by looking in the Security
// and Compliance Center UI.
//
// With Security and Compliance Center, you can manage a custom control library  that is specific to your organization's
// needs. Each control has several specifications  and assessments that are mapped to it.  For more information, see
// [Creating custom libraries](/docs/security-compliance?topic=security-compliance-custom-library).
func (compliance *ComplianceV2) DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions) (result *ControlLibraryDelete, response *core.DetailedResponse, err error) {
	return compliance.DeleteCustomControlLibraryWithContext(context.Background(), deleteCustomControlLibraryOptions)
}

// DeleteCustomControlLibraryWithContext is an alternate form of the DeleteCustomControlLibrary method which supports a Context parameter
func (compliance *ComplianceV2) DeleteCustomControlLibraryWithContext(ctx context.Context, deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions) (result *ControlLibraryDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomControlLibraryOptions, "deleteCustomControlLibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomControlLibraryOptions, "deleteCustomControlLibraryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"control_libraries_id": *deleteCustomControlLibraryOptions.ControlLibrariesID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "DeleteCustomControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteCustomControlLibraryOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteCustomControlLibraryOptions.XCorrelationID))
	}
	if deleteCustomControlLibraryOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*deleteCustomControlLibraryOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProfiles : Get all profiles
// View all of the predefined and custom profiles that are available in your account.
func (compliance *ComplianceV2) ListProfiles(listProfilesOptions *ListProfilesOptions) (result *ProfileCollection, response *core.DetailedResponse, err error) {
	return compliance.ListProfilesWithContext(context.Background(), listProfilesOptions)
}

// ListProfilesWithContext is an alternate form of the ListProfiles method which supports a Context parameter
func (compliance *ComplianceV2) ListProfilesWithContext(ctx context.Context, listProfilesOptions *ListProfilesOptions) (result *ProfileCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProfilesOptions, "listProfilesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ListProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listProfilesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listProfilesOptions.XCorrelationID))
	}
	if listProfilesOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*listProfilesOptions.XRequestID))
	}

	if listProfilesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listProfilesOptions.Limit))
	}
	if listProfilesOptions.ProfileType != nil {
		builder.AddQuery("profile_type", fmt.Sprint(*listProfilesOptions.ProfileType))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateProfile : Create a custom profile
// Create a custom profile that is specific to your usecase, by using an existing library as a starting point.  For more
// information, see [Building custom
// profiles](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-build-custom-profiles&interface=api).
func (compliance *ComplianceV2) CreateProfile(createProfileOptions *CreateProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	return compliance.CreateProfileWithContext(context.Background(), createProfileOptions)
}

// CreateProfileWithContext is an alternate form of the CreateProfile method which supports a Context parameter
func (compliance *ComplianceV2) CreateProfileWithContext(ctx context.Context, createProfileOptions *CreateProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProfileOptions, "createProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProfileOptions, "createProfileOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "CreateProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createProfileOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createProfileOptions.XCorrelationID))
	}
	if createProfileOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*createProfileOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if createProfileOptions.ProfileName != nil {
		body["profile_name"] = createProfileOptions.ProfileName
	}
	if createProfileOptions.ProfileDescription != nil {
		body["profile_description"] = createProfileOptions.ProfileDescription
	}
	if createProfileOptions.ProfileType != nil {
		body["profile_type"] = createProfileOptions.ProfileType
	}
	if createProfileOptions.Controls != nil {
		body["controls"] = createProfileOptions.Controls
	}
	if createProfileOptions.DefaultParameters != nil {
		body["default_parameters"] = createProfileOptions.DefaultParameters
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProfile : Get a profile
// View the details of a profile by providing the profile ID.  You can find the profile ID in the Security and
// Compliance Center UI. For more information, see [Building custom
// profiles](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-build-custom-profiles&interface=api).
func (compliance *ComplianceV2) GetProfile(getProfileOptions *GetProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	return compliance.GetProfileWithContext(context.Background(), getProfileOptions)
}

// GetProfileWithContext is an alternate form of the GetProfile method which supports a Context parameter
func (compliance *ComplianceV2) GetProfileWithContext(ctx context.Context, getProfileOptions *GetProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProfileOptions, "getProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProfileOptions, "getProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *getProfileOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "GetProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getProfileOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getProfileOptions.XCorrelationID))
	}
	if getProfileOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*getProfileOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceProfile : Update a profile
// Update the details of a custom profile. With Security and Compliance Center, you can manage  a profile that is
// specific to your usecase, by using an existing library as a starting point.  For more information, see [Building
// custom
// profiles](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-build-custom-profiles&interface=api).
func (compliance *ComplianceV2) ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	return compliance.ReplaceProfileWithContext(context.Background(), replaceProfileOptions)
}

// ReplaceProfileWithContext is an alternate form of the ReplaceProfile method which supports a Context parameter
func (compliance *ComplianceV2) ReplaceProfileWithContext(ctx context.Context, replaceProfileOptions *ReplaceProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceProfileOptions, "replaceProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceProfileOptions, "replaceProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *replaceProfileOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ReplaceProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceProfileOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*replaceProfileOptions.XCorrelationID))
	}
	if replaceProfileOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*replaceProfileOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if replaceProfileOptions.ProfileName != nil {
		body["profile_name"] = replaceProfileOptions.ProfileName
	}
	if replaceProfileOptions.ProfileDescription != nil {
		body["profile_description"] = replaceProfileOptions.ProfileDescription
	}
	if replaceProfileOptions.ProfileType != nil {
		body["profile_type"] = replaceProfileOptions.ProfileType
	}
	if replaceProfileOptions.Controls != nil {
		body["controls"] = replaceProfileOptions.Controls
	}
	if replaceProfileOptions.DefaultParameters != nil {
		body["default_parameters"] = replaceProfileOptions.DefaultParameters
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomProfile : Delete a custom profile
// Delete a custom profile by providing the profile ID.  You can find the ID in the Security and Compliance Center UI.
// For more information about managing your custom profiles, see [Building custom
// profiles](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-build-custom-profiles&interface=api).
func (compliance *ComplianceV2) DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	return compliance.DeleteCustomProfileWithContext(context.Background(), deleteCustomProfileOptions)
}

// DeleteCustomProfileWithContext is an alternate form of the DeleteCustomProfile method which supports a Context parameter
func (compliance *ComplianceV2) DeleteCustomProfileWithContext(ctx context.Context, deleteCustomProfileOptions *DeleteCustomProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomProfileOptions, "deleteCustomProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomProfileOptions, "deleteCustomProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *deleteCustomProfileOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "DeleteCustomProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteCustomProfileOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteCustomProfileOptions.XCorrelationID))
	}
	if deleteCustomProfileOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*deleteCustomProfileOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAttachments : Get all attachments linked to a specific profile
// View all of the attachments that are linked to a specific profile.  An attachment is the association between the set
// of resources that you want to evaluate  and a profile that contains the specific controls that you want to use. For
// more information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) (result *AttachmentCollection, response *core.DetailedResponse, err error) {
	return compliance.ListAttachmentsWithContext(context.Background(), listAttachmentsOptions)
}

// ListAttachmentsWithContext is an alternate form of the ListAttachments method which supports a Context parameter
func (compliance *ComplianceV2) ListAttachmentsWithContext(ctx context.Context, listAttachmentsOptions *ListAttachmentsOptions) (result *AttachmentCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAttachmentsOptions, "listAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAttachmentsOptions, "listAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *listAttachmentsOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ListAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAttachmentsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listAttachmentsOptions.XCorrelationID))
	}
	if listAttachmentsOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*listAttachmentsOptions.XRequestID))
	}

	if listAttachmentsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAttachmentsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateAttachment : Create an attachment
// Create an attachment to link to a profile to schedule evaluations  of your resources on a recurring schedule, or
// on-demand. For more information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) CreateAttachment(createAttachmentOptions *CreateAttachmentOptions) (result *AttachmentPrototype, response *core.DetailedResponse, err error) {
	return compliance.CreateAttachmentWithContext(context.Background(), createAttachmentOptions)
}

// CreateAttachmentWithContext is an alternate form of the CreateAttachment method which supports a Context parameter
func (compliance *ComplianceV2) CreateAttachmentWithContext(ctx context.Context, createAttachmentOptions *CreateAttachmentOptions) (result *AttachmentPrototype, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAttachmentOptions, "createAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAttachmentOptions, "createAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *createAttachmentOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "CreateAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createAttachmentOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createAttachmentOptions.XCorrelationID))
	}
	if createAttachmentOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*createAttachmentOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if createAttachmentOptions.Attachments != nil {
		body["attachments"] = createAttachmentOptions.Attachments
	}
	if createAttachmentOptions.ProfileID != nil {
		body["profile_id"] = createAttachmentOptions.ProfileID
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentPrototype)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProfileAttachment : Get an attachment
// View the details of an attachment a profile by providing the attachment ID.  You can find this value in the Security
// and Compliance Center UI. For more information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	return compliance.GetProfileAttachmentWithContext(context.Background(), getProfileAttachmentOptions)
}

// GetProfileAttachmentWithContext is an alternate form of the GetProfileAttachment method which supports a Context parameter
func (compliance *ComplianceV2) GetProfileAttachmentWithContext(ctx context.Context, getProfileAttachmentOptions *GetProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProfileAttachmentOptions, "getProfileAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProfileAttachmentOptions, "getProfileAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"attachment_id": *getProfileAttachmentOptions.AttachmentID,
		"profiles_id":   *getProfileAttachmentOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfileAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "GetProfileAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getProfileAttachmentOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getProfileAttachmentOptions.XCorrelationID))
	}
	if getProfileAttachmentOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*getProfileAttachmentOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceProfileAttachment : Update an attachment
// Update an attachment that is linked to a profile to evaluate your resources  on a recurring schedule, or on-demand.
// For more information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	return compliance.ReplaceProfileAttachmentWithContext(context.Background(), replaceProfileAttachmentOptions)
}

// ReplaceProfileAttachmentWithContext is an alternate form of the ReplaceProfileAttachment method which supports a Context parameter
func (compliance *ComplianceV2) ReplaceProfileAttachmentWithContext(ctx context.Context, replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceProfileAttachmentOptions, "replaceProfileAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceProfileAttachmentOptions, "replaceProfileAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"attachment_id": *replaceProfileAttachmentOptions.AttachmentID,
		"profiles_id":   *replaceProfileAttachmentOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceProfileAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ReplaceProfileAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceProfileAttachmentOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*replaceProfileAttachmentOptions.XCorrelationID))
	}
	if replaceProfileAttachmentOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*replaceProfileAttachmentOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if replaceProfileAttachmentOptions.ID != nil {
		body["id"] = replaceProfileAttachmentOptions.ID
	}
	if replaceProfileAttachmentOptions.ProfileID != nil {
		body["profile_id"] = replaceProfileAttachmentOptions.ProfileID
	}
	if replaceProfileAttachmentOptions.AccountID != nil {
		body["account_id"] = replaceProfileAttachmentOptions.AccountID
	}
	if replaceProfileAttachmentOptions.InstanceID != nil {
		body["instance_id"] = replaceProfileAttachmentOptions.InstanceID
	}
	if replaceProfileAttachmentOptions.Scope != nil {
		body["scope"] = replaceProfileAttachmentOptions.Scope
	}
	if replaceProfileAttachmentOptions.CreatedOn != nil {
		body["created_on"] = replaceProfileAttachmentOptions.CreatedOn
	}
	if replaceProfileAttachmentOptions.CreatedBy != nil {
		body["created_by"] = replaceProfileAttachmentOptions.CreatedBy
	}
	if replaceProfileAttachmentOptions.UpdatedOn != nil {
		body["updated_on"] = replaceProfileAttachmentOptions.UpdatedOn
	}
	if replaceProfileAttachmentOptions.UpdatedBy != nil {
		body["updated_by"] = replaceProfileAttachmentOptions.UpdatedBy
	}
	if replaceProfileAttachmentOptions.Status != nil {
		body["status"] = replaceProfileAttachmentOptions.Status
	}
	if replaceProfileAttachmentOptions.Schedule != nil {
		body["schedule"] = replaceProfileAttachmentOptions.Schedule
	}
	if replaceProfileAttachmentOptions.Notifications != nil {
		body["notifications"] = replaceProfileAttachmentOptions.Notifications
	}
	if replaceProfileAttachmentOptions.AttachmentParameters != nil {
		body["attachment_parameters"] = replaceProfileAttachmentOptions.AttachmentParameters
	}
	if replaceProfileAttachmentOptions.LastScan != nil {
		body["last_scan"] = replaceProfileAttachmentOptions.LastScan
	}
	if replaceProfileAttachmentOptions.NextScanTime != nil {
		body["next_scan_time"] = replaceProfileAttachmentOptions.NextScanTime
	}
	if replaceProfileAttachmentOptions.Name != nil {
		body["name"] = replaceProfileAttachmentOptions.Name
	}
	if replaceProfileAttachmentOptions.Description != nil {
		body["description"] = replaceProfileAttachmentOptions.Description
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteProfileAttachment : Delete an attachment
// Delete an attachment. Alternatively, if you think that you might need  this configuration in the future, you can
// pause an attachment to stop being charged. For more information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	return compliance.DeleteProfileAttachmentWithContext(context.Background(), deleteProfileAttachmentOptions)
}

// DeleteProfileAttachmentWithContext is an alternate form of the DeleteProfileAttachment method which supports a Context parameter
func (compliance *ComplianceV2) DeleteProfileAttachmentWithContext(ctx context.Context, deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions) (result *AttachmentItem, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProfileAttachmentOptions, "deleteProfileAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProfileAttachmentOptions, "deleteProfileAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"attachment_id": *deleteProfileAttachmentOptions.AttachmentID,
		"profiles_id":   *deleteProfileAttachmentOptions.ProfilesID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProfileAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "DeleteProfileAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteProfileAttachmentOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteProfileAttachmentOptions.XCorrelationID))
	}
	if deleteProfileAttachmentOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*deleteProfileAttachmentOptions.XRequestID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateScan : Create a scan
// Create a scan to evaluate your resources on a recurring basis or on demand.
func (compliance *ComplianceV2) CreateScan(createScanOptions *CreateScanOptions) (result *Scan, response *core.DetailedResponse, err error) {
	return compliance.CreateScanWithContext(context.Background(), createScanOptions)
}

// CreateScanWithContext is an alternate form of the CreateScan method which supports a Context parameter
func (compliance *ComplianceV2) CreateScanWithContext(ctx context.Context, createScanOptions *CreateScanOptions) (result *Scan, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createScanOptions, "createScanOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createScanOptions, "createScanOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/scans`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createScanOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "CreateScan")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createScanOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createScanOptions.XCorrelationID))
	}
	if createScanOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*createScanOptions.XRequestID))
	}

	body := make(map[string]interface{})
	if createScanOptions.AttachmentID != nil {
		body["attachment_id"] = createScanOptions.AttachmentID
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
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScan)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAttachmentsAccount : Get all attachments in an instance
// View all of the attachments that are linked to an account. An attachment is the association between the set of
// resources that you want to evaluate  and a profile that contains the specific controls that you want to use. For more
// information, see [Running an evaluation for IBM
// Cloud](https://test.cloud.ibm.com/docs/security-compliance?topic=security-compliance-scan-resources).
func (compliance *ComplianceV2) ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions) (result *AttachmentCollection, response *core.DetailedResponse, err error) {
	return compliance.ListAttachmentsAccountWithContext(context.Background(), listAttachmentsAccountOptions)
}

// ListAttachmentsAccountWithContext is an alternate form of the ListAttachmentsAccount method which supports a Context parameter
func (compliance *ComplianceV2) ListAttachmentsAccountWithContext(ctx context.Context, listAttachmentsAccountOptions *ListAttachmentsAccountOptions) (result *AttachmentCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAttachmentsAccountOptions, "listAttachmentsAccountOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = compliance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(compliance.Service.Options.URL, `/attachments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAttachmentsAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("compliance", "V2", "ListAttachmentsAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAttachmentsAccountOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listAttachmentsAccountOptions.XCorrelationID))
	}
	if listAttachmentsAccountOptions.XRequestID != nil {
		builder.AddHeader("X-Request-ID", fmt.Sprint(*listAttachmentsAccountOptions.XRequestID))
	}

	if listAttachmentsAccountOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAttachmentsAccountOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = compliance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Attachment : The request payload of the attachments parameter.
type AttachmentItem struct {
	// The ID of the attachment.
	ID *string `json:"id,omitempty"`

	// The ID of the profile that is specified in the attachment.
	ProfileID *string `json:"profile_id,omitempty"`

	// The account ID that is associated to the attachment.
	AccountID *string `json:"account_id,omitempty"`

	// The instance ID of the account that is associated to the attachment.
	InstanceID *string `json:"instance_id,omitempty"`

	// The scope payload for the multi cloud feature.
	Scope []MultiCloudScope `json:"scope,omitempty"`

	// The date when the attachment was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who created the attachment.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the attachment was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The user who updated the attachment.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The status of an attachment evaluation.
	Status *string `json:"status,omitempty"`

	// The schedule of an attachment evaluation.
	Schedule *string `json:"schedule,omitempty"`

	// The request payload of the attachment notifications.
	Notifications *AttachmentsNotificationsPrototype `json:"notifications,omitempty"`

	// The profile parameters for the attachment.
	AttachmentParameters []AttachmentParametersPrototype `json:"attachment_parameters,omitempty"`

	// The details of the last scan of an attachment.
	LastScan *LastScan `json:"last_scan,omitempty"`

	// The start time of the next scan.
	NextScanTime *string `json:"next_scan_time,omitempty"`

	// The name that is generated from the scope type and ID.
	Name *string `json:"name,omitempty"`

	// The description for the attachment.
	Description *string `json:"description,omitempty"`
}

// Constants associated with the Attachment.Status property.
// The status of an attachment evaluation.
const (
	Attachment_Status_Disabled = "disabled"
	Attachment_Status_Enabled  = "enabled"
)

// Constants associated with the Attachment.Schedule property.
// The schedule of an attachment evaluation.
const (
	Attachment_Schedule_Daily       = "daily"
	Attachment_Schedule_Every30Days = "every_30_days"
	Attachment_Schedule_Every7Days  = "every_7_days"
)

// UnmarshalAttachment unmarshals an instance of Attachment from the specified map of raw messages.
func UnmarshalAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scope", &obj.Scope, UnmarshalMultiCloudScope)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schedule", &obj.Schedule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalAttachmentsNotificationsPrototype)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment_parameters", &obj.AttachmentParameters, UnmarshalAttachmentParametersPrototype)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_scan", &obj.LastScan, UnmarshalLastScan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_scan_time", &obj.NextScanTime)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentCollection : The response body of an attachment.
type AttachmentCollection struct {
	// The number of attachments.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The limit of attachments per request.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first page of entries.
	First *PaginatedCollectionFirst `json:"first,omitempty"`

	// The reference URL for the next few entries.
	Next *PaginatedCollectionNext `json:"next,omitempty"`

	// The list of attachments.
	Attachments []AttachmentItem `json:"attachments,omitempty"`
}

// UnmarshalAttachmentCollection unmarshals an instance of AttachmentCollection from the specified map of raw messages.
func UnmarshalAttachmentCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentCollection)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedCollectionFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedCollectionNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentParametersPrototype : The control details of a profile.
type AttachmentParametersPrototype struct {
	// The type of the implementation.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The implementation ID of the parameter.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// The parameter name.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The value of the parameter.
	ParameterValue *string `json:"parameter_value,omitempty"`

	// The parameter display name.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// The parameter type.
	ParameterType *string `json:"parameter_type,omitempty"`
}

// Constants associated with the AttachmentParametersPrototype.ParameterType property.
// The parameter type.
const (
	AttachmentParametersPrototype_ParameterType_Boolean    = "boolean"
	AttachmentParametersPrototype_ParameterType_General    = "general"
	AttachmentParametersPrototype_ParameterType_IpList     = "ip_list"
	AttachmentParametersPrototype_ParameterType_Numeric    = "numeric"
	AttachmentParametersPrototype_ParameterType_String     = "string"
	AttachmentParametersPrototype_ParameterType_StringList = "string_list"
	AttachmentParametersPrototype_ParameterType_Timestamp  = "timestamp"
)

// UnmarshalAttachmentParametersPrototype unmarshals an instance of AttachmentParametersPrototype from the specified map of raw messages.
func UnmarshalAttachmentParametersPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentParametersPrototype)
	err = core.UnmarshalPrimitive(m, "assessment_type", &obj.AssessmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_id", &obj.AssessmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_name", &obj.ParameterName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_value", &obj.ParameterValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_display_name", &obj.ParameterDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_type", &obj.ParameterType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentPrototype : The request body of getting an attachment that is associated with your account.
type AttachmentPrototype struct {
	// The ID of the profile that is specified in the attachment.
	ProfileID *string `json:"profile_id,omitempty"`

	// The array that displays all of the available attachments.
	Attachments []AttachmentsPrototype `json:"attachments" validate:"required"`
}

// NewAttachmentPrototype : Instantiate AttachmentPrototype (Generic Model Constructor)
func (*ComplianceV2) NewAttachmentPrototype(attachments []AttachmentsPrototype) (_model *AttachmentPrototype, err error) {
	_model = &AttachmentPrototype{
		Attachments: attachments,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAttachmentPrototype unmarshals an instance of AttachmentPrototype from the specified map of raw messages.
func UnmarshalAttachmentPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentPrototype)
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachmentsPrototype)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentsNotificationsPrototype : The request payload of the attachment notifications.
type AttachmentsNotificationsPrototype struct {
	// enabled notifications.
	Enabled *bool `json:"enabled" validate:"required"`

	// The failed controls.
	Controls *FailedControls `json:"controls" validate:"required"`
}

// NewAttachmentsNotificationsPrototype : Instantiate AttachmentsNotificationsPrototype (Generic Model Constructor)
func (*ComplianceV2) NewAttachmentsNotificationsPrototype(enabled bool, controls *FailedControls) (_model *AttachmentsNotificationsPrototype, err error) {
	_model = &AttachmentsNotificationsPrototype{
		Enabled:  core.BoolPtr(enabled),
		Controls: controls,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAttachmentsNotificationsPrototype unmarshals an instance of AttachmentsNotificationsPrototype from the specified map of raw messages.
func UnmarshalAttachmentsNotificationsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentsNotificationsPrototype)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalFailedControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentsPrototype : The request payload of getting all of the attachments that are associated with the account.
type AttachmentsPrototype struct {
	// The ID of the attachment.
	ID *string `json:"id,omitempty"`

	// The name that is generated from the scope type and ID.
	Name *string `json:"name" validate:"required"`

	// The description for the attachment.
	Description *string `json:"description,omitempty"`

	// The scope payload for the multi cloud feature.
	Scope []MultiCloudScope `json:"scope" validate:"required"`

	// The status of the scan of an attachment.
	Status *string `json:"status" validate:"required"`

	// The schedule of an attachment evaluation.
	Schedule *string `json:"schedule" validate:"required"`

	// The request payload of the attachment notifications.
	Notifications *AttachmentsNotificationsPrototype `json:"notifications,omitempty"`

	// The profile parameters for the attachment.
	AttachmentParameters []AttachmentParametersPrototype `json:"attachment_parameters" validate:"required"`
}

// Constants associated with the AttachmentsPrototype.Status property.
// The status of the scan of an attachment.
const (
	AttachmentsPrototype_Status_Disabled = "disabled"
	AttachmentsPrototype_Status_Enabled  = "enabled"
)

// Constants associated with the AttachmentsPrototype.Schedule property.
// The schedule of an attachment evaluation.
const (
	AttachmentsPrototype_Schedule_Daily       = "daily"
	AttachmentsPrototype_Schedule_Every30Days = "every_30_days"
	AttachmentsPrototype_Schedule_Every7Days  = "every_7_days"
)

// NewAttachmentsPrototype : Instantiate AttachmentsPrototype (Generic Model Constructor)
func (*ComplianceV2) NewAttachmentsPrototype(name string, scope []MultiCloudScope, status string, schedule string, attachmentParameters []AttachmentParametersPrototype) (_model *AttachmentsPrototype, err error) {
	_model = &AttachmentsPrototype{
		Name:                 core.StringPtr(name),
		Scope:                scope,
		Status:               core.StringPtr(status),
		Schedule:             core.StringPtr(schedule),
		AttachmentParameters: attachmentParameters,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAttachmentsPrototype unmarshals an instance of AttachmentsPrototype from the specified map of raw messages.
func UnmarshalAttachmentsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentsPrototype)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalModel(m, "scope", &obj.Scope, UnmarshalMultiCloudScope)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schedule", &obj.Schedule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalAttachmentsNotificationsPrototype)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment_parameters", &obj.AttachmentParameters, UnmarshalAttachmentParametersPrototype)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlDocs : The control documentation.
type ControlDocs struct {
	// The ID of the control documentation.
	ControlDocsID *string `json:"control_docs_id,omitempty"`

	// The type of control documentation.
	ControlDocsType *string `json:"control_docs_type,omitempty"`
}

// UnmarshalControlDocs unmarshals an instance of ControlDocs from the specified map of raw messages.
func UnmarshalControlDocs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlDocs)
	err = core.UnmarshalPrimitive(m, "control_docs_id", &obj.ControlDocsID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_docs_type", &obj.ControlDocsType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlLibrary : The request payload of the control library.
type ControlLibrary struct {
	// The control library ID.
	ID *string `json:"id,omitempty"`

	// The account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The control library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// The control library description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// The control library type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// The version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The date when the control library was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who created the control library.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the control library was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The user who updated the control library.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The latest version of the control library.
	Latest *bool `json:"latest,omitempty"`

	// The indication of whether hierarchy is enabled for the profile.
	HierarchyEnabled *bool `json:"hierarchy_enabled,omitempty"`

	// The number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// The number of parent controls in the profile.
	ControlParentsCount *int64 `json:"control_parents_count,omitempty"`

	// The controls.
	Controls []ControlsInControlLib `json:"controls,omitempty"`
}

// Constants associated with the ControlLibrary.ControlLibraryType property.
// The control library type.
const (
	ControlLibrary_ControlLibraryType_Custom     = "custom"
	ControlLibrary_ControlLibraryType_Predefined = "predefined"
)

// UnmarshalControlLibrary unmarshals an instance of ControlLibrary from the specified map of raw messages.
func UnmarshalControlLibrary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibrary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_name", &obj.ControlLibraryName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_description", &obj.ControlLibraryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_type", &obj.ControlLibraryType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_group_label", &obj.VersionGroupLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_version", &obj.ControlLibraryVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest", &obj.Latest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hierarchy_enabled", &obj.HierarchyEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_count", &obj.ControlsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_parents_count", &obj.ControlParentsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalControlsInControlLib)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlLibraryCollection : The response body of control libraries.
type ControlLibraryCollection struct {
	// The number of control libraries.
	TotalCount *int64 `json:"total_count,omitempty"`

	// limit.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first page of entries.
	First *PaginatedCollectionFirst `json:"first,omitempty"`

	// The reference URL for the next few entries.
	Next *PaginatedCollectionNext `json:"next,omitempty"`

	// The control libraries.
	ControlLibraries []ControlLibraryItem `json:"control_libraries,omitempty"`
}

// UnmarshalControlLibraryCollection unmarshals an instance of ControlLibraryCollection from the specified map of raw messages.
func UnmarshalControlLibraryCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibraryCollection)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedCollectionFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedCollectionNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_libraries", &obj.ControlLibraries, UnmarshalControlLibraryItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlLibraryDelete : The response body of deleting of a control library.
type ControlLibraryDelete struct {
	// The delete message of a control library.
	Deleted *string `json:"deleted,omitempty"`
}

// UnmarshalControlLibraryDelete unmarshals an instance of ControlLibraryDelete from the specified map of raw messages.
func UnmarshalControlLibraryDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibraryDelete)
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlLibraryItem : ControlLibraryItem struct
type ControlLibraryItem struct {
	// The ID of the control library.
	ID *string `json:"id,omitempty"`

	// The Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The control library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// The control library description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// The control library type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// The date when the control library was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who created the control library.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the control library was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The use who updated the control library.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The latest control library version.
	Latest *bool `json:"latest,omitempty"`

	// The number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`
}

// UnmarshalControlLibraryItem unmarshals an instance of ControlLibraryItem from the specified map of raw messages.
func UnmarshalControlLibraryItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibraryItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_name", &obj.ControlLibraryName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_description", &obj.ControlLibraryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_type", &obj.ControlLibraryType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_group_label", &obj.VersionGroupLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_version", &obj.ControlLibraryVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest", &obj.Latest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_count", &obj.ControlsCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlSpecifications : The control specifications of a control library.
type ControlSpecifications struct {
	// The control specification ID.
	ControlSpecificationID *string `json:"control_specification_id,omitempty"`

	// The responsibility for managing the control.
	Responsibility *string `json:"responsibility,omitempty"`

	// The component ID.
	ComponentID *string `json:"component_id,omitempty"`

	// The component name.
	ComponenetName *string `json:"componenet_name,omitempty"`

	// The control specifications environment.
	Environment *string `json:"environment,omitempty"`

	// The control specifications description.
	ControlSpecificationDescription *string `json:"control_specification_description,omitempty"`

	// The number of assessments.
	AssessmentsCount *int64 `json:"assessments_count,omitempty"`

	// The assessments.
	Assessments []Implementation `json:"assessments,omitempty"`
}

// Constants associated with the ControlSpecifications.Responsibility property.
// The responsibility for managing the control.
const (
	ControlSpecifications_Responsibility_User = "user"
)

// UnmarshalControlSpecifications unmarshals an instance of ControlSpecifications from the specified map of raw messages.
func UnmarshalControlSpecifications(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlSpecifications)
	err = core.UnmarshalPrimitive(m, "control_specification_id", &obj.ControlSpecificationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "responsibility", &obj.Responsibility)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "component_id", &obj.ComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "componenet_name", &obj.ComponenetName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_specification_description", &obj.ControlSpecificationDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessments_count", &obj.AssessmentsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "assessments", &obj.Assessments, UnmarshalImplementation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlsInControlLib : The control details of a control library.
type ControlsInControlLib struct {
	// The ID of the control library that contains the profile.
	ControlName *string `json:"control_name,omitempty"`

	// The control name.
	ControlID *string `json:"control_id,omitempty"`

	// The control description.
	ControlDescription *string `json:"control_description,omitempty"`

	// The control category.
	ControlCategory *string `json:"control_category,omitempty"`

	// The parent control.
	ControlParent *string `json:"control_parent,omitempty"`

	// The indication of whether the control is required.
	ControlRequirement *bool `json:"control_requirement,omitempty"`

	// The control tags.
	ControlTags []string `json:"control_tags,omitempty"`

	// The control specifications.
	ControlSpecifications []ControlSpecifications `json:"control_specifications,omitempty"`

	// The control documentation.
	ControlDocs *ControlDocs `json:"control_docs,omitempty"`

	// The control status.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the ControlsInControlLib.Status property.
// The control status.
const (
	ControlsInControlLib_Status_Disabled = "disabled"
	ControlsInControlLib_Status_Enabled  = "enabled"
)

// UnmarshalControlsInControlLib unmarshals an instance of ControlsInControlLib from the specified map of raw messages.
func UnmarshalControlsInControlLib(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlsInControlLib)
	err = core.UnmarshalPrimitive(m, "control_name", &obj.ControlName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_id", &obj.ControlID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_description", &obj.ControlDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_category", &obj.ControlCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_parent", &obj.ControlParent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_requirement", &obj.ControlRequirement)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_tags", &obj.ControlTags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_specifications", &obj.ControlSpecifications, UnmarshalControlSpecifications)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_docs", &obj.ControlDocs, UnmarshalControlDocs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateAttachmentOptions : The CreateAttachment options.
type CreateAttachmentOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The array that displays all of the available attachments.
	Attachments []AttachmentsPrototype `json:"attachments" validate:"required"`

	// The ID of the profile that is specified in the attachment.
	ProfileID *string `json:"profile_id,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAttachmentOptions : Instantiate CreateAttachmentOptions
func (*ComplianceV2) NewCreateAttachmentOptions(profilesID string, attachments []AttachmentsPrototype) *CreateAttachmentOptions {
	return &CreateAttachmentOptions{
		ProfilesID:  core.StringPtr(profilesID),
		Attachments: attachments,
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *CreateAttachmentOptions) SetProfilesID(profilesID string) *CreateAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachments : Allow user to set Attachments
func (_options *CreateAttachmentOptions) SetAttachments(attachments []AttachmentsPrototype) *CreateAttachmentOptions {
	_options.Attachments = attachments
	return _options
}

// SetProfileID : Allow user to set ProfileID
func (_options *CreateAttachmentOptions) SetProfileID(profileID string) *CreateAttachmentOptions {
	_options.ProfileID = core.StringPtr(profileID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateAttachmentOptions) SetXCorrelationID(xCorrelationID string) *CreateAttachmentOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *CreateAttachmentOptions) SetXRequestID(xRequestID string) *CreateAttachmentOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAttachmentOptions) SetHeaders(param map[string]string) *CreateAttachmentOptions {
	options.Headers = param
	return options
}

// CreateCustomControlLibraryOptions : The CreateCustomControlLibrary options.
type CreateCustomControlLibraryOptions struct {
	// The control library name.
	ControlLibraryName *string `json:"control_library_name" validate:"required"`

	// The control library description.
	ControlLibraryDescription *string `json:"control_library_description" validate:"required"`

	// The control library type.
	ControlLibraryType *string `json:"control_library_type" validate:"required"`

	// The controls.
	Controls []ControlsInControlLib `json:"controls" validate:"required"`

	// The version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The latest control library version.
	Latest *bool `json:"latest,omitempty"`

	// The number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCustomControlLibraryOptions.ControlLibraryType property.
// The control library type.
const (
	CreateCustomControlLibraryOptions_ControlLibraryType_Custom     = "custom"
	CreateCustomControlLibraryOptions_ControlLibraryType_Predefined = "predefined"
)

// NewCreateCustomControlLibraryOptions : Instantiate CreateCustomControlLibraryOptions
func (*ComplianceV2) NewCreateCustomControlLibraryOptions(controlLibraryName string, controlLibraryDescription string, controlLibraryType string, controls []ControlsInControlLib) *CreateCustomControlLibraryOptions {
	return &CreateCustomControlLibraryOptions{
		ControlLibraryName:        core.StringPtr(controlLibraryName),
		ControlLibraryDescription: core.StringPtr(controlLibraryDescription),
		ControlLibraryType:        core.StringPtr(controlLibraryType),
		Controls:                  controls,
	}
}

// SetControlLibraryName : Allow user to set ControlLibraryName
func (_options *CreateCustomControlLibraryOptions) SetControlLibraryName(controlLibraryName string) *CreateCustomControlLibraryOptions {
	_options.ControlLibraryName = core.StringPtr(controlLibraryName)
	return _options
}

// SetControlLibraryDescription : Allow user to set ControlLibraryDescription
func (_options *CreateCustomControlLibraryOptions) SetControlLibraryDescription(controlLibraryDescription string) *CreateCustomControlLibraryOptions {
	_options.ControlLibraryDescription = core.StringPtr(controlLibraryDescription)
	return _options
}

// SetControlLibraryType : Allow user to set ControlLibraryType
func (_options *CreateCustomControlLibraryOptions) SetControlLibraryType(controlLibraryType string) *CreateCustomControlLibraryOptions {
	_options.ControlLibraryType = core.StringPtr(controlLibraryType)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *CreateCustomControlLibraryOptions) SetControls(controls []ControlsInControlLib) *CreateCustomControlLibraryOptions {
	_options.Controls = controls
	return _options
}

// SetVersionGroupLabel : Allow user to set VersionGroupLabel
func (_options *CreateCustomControlLibraryOptions) SetVersionGroupLabel(versionGroupLabel string) *CreateCustomControlLibraryOptions {
	_options.VersionGroupLabel = core.StringPtr(versionGroupLabel)
	return _options
}

// SetControlLibraryVersion : Allow user to set ControlLibraryVersion
func (_options *CreateCustomControlLibraryOptions) SetControlLibraryVersion(controlLibraryVersion string) *CreateCustomControlLibraryOptions {
	_options.ControlLibraryVersion = core.StringPtr(controlLibraryVersion)
	return _options
}

// SetLatest : Allow user to set Latest
func (_options *CreateCustomControlLibraryOptions) SetLatest(latest bool) *CreateCustomControlLibraryOptions {
	_options.Latest = core.BoolPtr(latest)
	return _options
}

// SetControlsCount : Allow user to set ControlsCount
func (_options *CreateCustomControlLibraryOptions) SetControlsCount(controlsCount int64) *CreateCustomControlLibraryOptions {
	_options.ControlsCount = core.Int64Ptr(controlsCount)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateCustomControlLibraryOptions) SetXCorrelationID(xCorrelationID string) *CreateCustomControlLibraryOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *CreateCustomControlLibraryOptions) SetXRequestID(xRequestID string) *CreateCustomControlLibraryOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomControlLibraryOptions) SetHeaders(param map[string]string) *CreateCustomControlLibraryOptions {
	options.Headers = param
	return options
}

// CreateProfileOptions : The CreateProfile options.
type CreateProfileOptions struct {
	// The name of the profile.
	ProfileName *string `json:"profile_name" validate:"required"`

	// The description of the profile.
	ProfileDescription *string `json:"profile_description" validate:"required"`

	// The profile type.
	ProfileType *string `json:"profile_type" validate:"required"`

	// The controls that are in the profile.
	Controls []ProfileControlsPrototype `json:"controls" validate:"required"`

	// The default parameters of the profile.
	DefaultParameters []DefaultParametersPrototype `json:"default_parameters" validate:"required"`

	// The supplied or generated value of this header is logged for a request, and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests, and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateProfileOptions.ProfileType property.
// The profile type.
const (
	CreateProfileOptions_ProfileType_Custom     = "custom"
	CreateProfileOptions_ProfileType_Predefined = "predefined"
)

// NewCreateProfileOptions : Instantiate CreateProfileOptions
func (*ComplianceV2) NewCreateProfileOptions(profileName string, profileDescription string, profileType string, controls []ProfileControlsPrototype, defaultParameters []DefaultParametersPrototype) *CreateProfileOptions {
	return &CreateProfileOptions{
		ProfileName:        core.StringPtr(profileName),
		ProfileDescription: core.StringPtr(profileDescription),
		ProfileType:        core.StringPtr(profileType),
		Controls:           controls,
		DefaultParameters:  defaultParameters,
	}
}

// SetProfileName : Allow user to set ProfileName
func (_options *CreateProfileOptions) SetProfileName(profileName string) *CreateProfileOptions {
	_options.ProfileName = core.StringPtr(profileName)
	return _options
}

// SetProfileDescription : Allow user to set ProfileDescription
func (_options *CreateProfileOptions) SetProfileDescription(profileDescription string) *CreateProfileOptions {
	_options.ProfileDescription = core.StringPtr(profileDescription)
	return _options
}

// SetProfileType : Allow user to set ProfileType
func (_options *CreateProfileOptions) SetProfileType(profileType string) *CreateProfileOptions {
	_options.ProfileType = core.StringPtr(profileType)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *CreateProfileOptions) SetControls(controls []ProfileControlsPrototype) *CreateProfileOptions {
	_options.Controls = controls
	return _options
}

// SetDefaultParameters : Allow user to set DefaultParameters
func (_options *CreateProfileOptions) SetDefaultParameters(defaultParameters []DefaultParametersPrototype) *CreateProfileOptions {
	_options.DefaultParameters = defaultParameters
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateProfileOptions) SetXCorrelationID(xCorrelationID string) *CreateProfileOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *CreateProfileOptions) SetXRequestID(xRequestID string) *CreateProfileOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProfileOptions) SetHeaders(param map[string]string) *CreateProfileOptions {
	options.Headers = param
	return options
}

// CreateScanOptions : The CreateScan options.
type CreateScanOptions struct {
	// The attachment ID of a profile.
	AttachmentID *string `json:"attachment_id" validate:"required"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateScanOptions : Instantiate CreateScanOptions
func (*ComplianceV2) NewCreateScanOptions(attachmentID string) *CreateScanOptions {
	return &CreateScanOptions{
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *CreateScanOptions) SetAttachmentID(attachmentID string) *CreateScanOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateScanOptions) SetXCorrelationID(xCorrelationID string) *CreateScanOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *CreateScanOptions) SetXRequestID(xRequestID string) *CreateScanOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateScanOptions) SetHeaders(param map[string]string) *CreateScanOptions {
	options.Headers = param
	return options
}

// DefaultParametersPrototype : The control details of a profile.
type DefaultParametersPrototype struct {
	// The type of the implementation.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The implementation ID of the parameter.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// The parameter name.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The default value of the parameter.
	ParameterDefaultValue *string `json:"parameter_default_value,omitempty"`

	// The parameter display name.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// The parameter type.
	ParameterType *string `json:"parameter_type,omitempty"`
}

// Constants associated with the DefaultParametersPrototype.ParameterType property.
// The parameter type.
const (
	DefaultParametersPrototype_ParameterType_Boolean    = "boolean"
	DefaultParametersPrototype_ParameterType_General    = "general"
	DefaultParametersPrototype_ParameterType_IpList     = "ip_list"
	DefaultParametersPrototype_ParameterType_Numeric    = "numeric"
	DefaultParametersPrototype_ParameterType_String     = "string"
	DefaultParametersPrototype_ParameterType_StringList = "string_list"
	DefaultParametersPrototype_ParameterType_Timestamp  = "timestamp"
)

// UnmarshalDefaultParametersPrototype unmarshals an instance of DefaultParametersPrototype from the specified map of raw messages.
func UnmarshalDefaultParametersPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultParametersPrototype)
	err = core.UnmarshalPrimitive(m, "assessment_type", &obj.AssessmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_id", &obj.AssessmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_name", &obj.ParameterName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_default_value", &obj.ParameterDefaultValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_display_name", &obj.ParameterDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_type", &obj.ParameterType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCustomControlLibraryOptions : The DeleteCustomControlLibrary options.
type DeleteCustomControlLibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomControlLibraryOptions : Instantiate DeleteCustomControlLibraryOptions
func (*ComplianceV2) NewDeleteCustomControlLibraryOptions(controlLibrariesID string) *DeleteCustomControlLibraryOptions {
	return &DeleteCustomControlLibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *DeleteCustomControlLibraryOptions) SetControlLibrariesID(controlLibrariesID string) *DeleteCustomControlLibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteCustomControlLibraryOptions) SetXCorrelationID(xCorrelationID string) *DeleteCustomControlLibraryOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *DeleteCustomControlLibraryOptions) SetXRequestID(xRequestID string) *DeleteCustomControlLibraryOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomControlLibraryOptions) SetHeaders(param map[string]string) *DeleteCustomControlLibraryOptions {
	options.Headers = param
	return options
}

// DeleteCustomProfileOptions : The DeleteCustomProfile options.
type DeleteCustomProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomProfileOptions : Instantiate DeleteCustomProfileOptions
func (*ComplianceV2) NewDeleteCustomProfileOptions(profilesID string) *DeleteCustomProfileOptions {
	return &DeleteCustomProfileOptions{
		ProfilesID: core.StringPtr(profilesID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *DeleteCustomProfileOptions) SetProfilesID(profilesID string) *DeleteCustomProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteCustomProfileOptions) SetXCorrelationID(xCorrelationID string) *DeleteCustomProfileOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *DeleteCustomProfileOptions) SetXRequestID(xRequestID string) *DeleteCustomProfileOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomProfileOptions) SetHeaders(param map[string]string) *DeleteCustomProfileOptions {
	options.Headers = param
	return options
}

// DeleteProfileAttachmentOptions : The DeleteProfileAttachment options.
type DeleteProfileAttachmentOptions struct {
	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProfileAttachmentOptions : Instantiate DeleteProfileAttachmentOptions
func (*ComplianceV2) NewDeleteProfileAttachmentOptions(attachmentID string, profilesID string) *DeleteProfileAttachmentOptions {
	return &DeleteProfileAttachmentOptions{
		AttachmentID: core.StringPtr(attachmentID),
		ProfilesID:   core.StringPtr(profilesID),
	}
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *DeleteProfileAttachmentOptions) SetAttachmentID(attachmentID string) *DeleteProfileAttachmentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *DeleteProfileAttachmentOptions) SetProfilesID(profilesID string) *DeleteProfileAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteProfileAttachmentOptions) SetXCorrelationID(xCorrelationID string) *DeleteProfileAttachmentOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *DeleteProfileAttachmentOptions) SetXRequestID(xRequestID string) *DeleteProfileAttachmentOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProfileAttachmentOptions) SetHeaders(param map[string]string) *DeleteProfileAttachmentOptions {
	options.Headers = param
	return options
}

// FailedControls : The failed controls.
type FailedControls struct {
	// The threshold limit.
	ThresholdLimit *int64 `json:"threshold_limit,omitempty"`

	// The failed control IDs.
	FailedControlIds []string `json:"failed_control_ids,omitempty"`
}

// UnmarshalFailedControls unmarshals an instance of FailedControls from the specified map of raw messages.
func UnmarshalFailedControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FailedControls)
	err = core.UnmarshalPrimitive(m, "threshold_limit", &obj.ThresholdLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed_control_ids", &obj.FailedControlIds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetControlLibraryOptions : The GetControlLibrary options.
type GetControlLibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetControlLibraryOptions : Instantiate GetControlLibraryOptions
func (*ComplianceV2) NewGetControlLibraryOptions(controlLibrariesID string) *GetControlLibraryOptions {
	return &GetControlLibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *GetControlLibraryOptions) SetControlLibrariesID(controlLibrariesID string) *GetControlLibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetControlLibraryOptions) SetXCorrelationID(xCorrelationID string) *GetControlLibraryOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *GetControlLibraryOptions) SetXRequestID(xRequestID string) *GetControlLibraryOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetControlLibraryOptions) SetHeaders(param map[string]string) *GetControlLibraryOptions {
	options.Headers = param
	return options
}

// GetProfileAttachmentOptions : The GetProfileAttachment options.
type GetProfileAttachmentOptions struct {
	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfileAttachmentOptions : Instantiate GetProfileAttachmentOptions
func (*ComplianceV2) NewGetProfileAttachmentOptions(attachmentID string, profilesID string) *GetProfileAttachmentOptions {
	return &GetProfileAttachmentOptions{
		AttachmentID: core.StringPtr(attachmentID),
		ProfilesID:   core.StringPtr(profilesID),
	}
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *GetProfileAttachmentOptions) SetAttachmentID(attachmentID string) *GetProfileAttachmentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *GetProfileAttachmentOptions) SetProfilesID(profilesID string) *GetProfileAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetProfileAttachmentOptions) SetXCorrelationID(xCorrelationID string) *GetProfileAttachmentOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *GetProfileAttachmentOptions) SetXRequestID(xRequestID string) *GetProfileAttachmentOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProfileAttachmentOptions) SetHeaders(param map[string]string) *GetProfileAttachmentOptions {
	options.Headers = param
	return options
}

// GetProfileOptions : The GetProfile options.
type GetProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfileOptions : Instantiate GetProfileOptions
func (*ComplianceV2) NewGetProfileOptions(profilesID string) *GetProfileOptions {
	return &GetProfileOptions{
		ProfilesID: core.StringPtr(profilesID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *GetProfileOptions) SetProfilesID(profilesID string) *GetProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetProfileOptions) SetXCorrelationID(xCorrelationID string) *GetProfileOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *GetProfileOptions) SetXRequestID(xRequestID string) *GetProfileOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProfileOptions) SetHeaders(param map[string]string) *GetProfileOptions {
	options.Headers = param
	return options
}

// Implementation : The implementation details of a control library.
type Implementation struct {
	// The assessment ID.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// The assessment method.
	AssessmentMethod *string `json:"assessment_method,omitempty"`

	// The assessment type.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The assessment description.
	AssessmentDescription *string `json:"assessment_description,omitempty"`

	// The parameter count.
	ParameterCount *int64 `json:"parameter_count,omitempty"`

	// The parameters.
	Parameters []ParameterInfo `json:"parameters,omitempty"`
}

// UnmarshalImplementation unmarshals an instance of Implementation from the specified map of raw messages.
func UnmarshalImplementation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Implementation)
	err = core.UnmarshalPrimitive(m, "assessment_id", &obj.AssessmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_method", &obj.AssessmentMethod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_type", &obj.AssessmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_description", &obj.AssessmentDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_count", &obj.ParameterCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalParameterInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LastScan : The details of the last scan of an attachment.
type LastScan struct {
	// The ID of the last scan of an attachment.
	ID *string `json:"id,omitempty"`

	// The status of the last scan of an attachment.
	Status *string `json:"status,omitempty"`

	// The time when the last scan started.
	Time *string `json:"time,omitempty"`
}

// Constants associated with the LastScan.Status property.
// The status of the last scan of an attachment.
const (
	LastScan_Status_Completed  = "completed"
	LastScan_Status_InProgress = "in_progress"
)

// UnmarshalLastScan unmarshals an instance of LastScan from the specified map of raw messages.
func UnmarshalLastScan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LastScan)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time", &obj.Time)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAttachmentsAccountOptions : The ListAttachmentsAccount options.
type ListAttachmentsAccountOptions struct {
	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// The indication of how many resources to return, unless the response is the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAttachmentsAccountOptions : Instantiate ListAttachmentsAccountOptions
func (*ComplianceV2) NewListAttachmentsAccountOptions() *ListAttachmentsAccountOptions {
	return &ListAttachmentsAccountOptions{}
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListAttachmentsAccountOptions) SetXCorrelationID(xCorrelationID string) *ListAttachmentsAccountOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ListAttachmentsAccountOptions) SetXRequestID(xRequestID string) *ListAttachmentsAccountOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAttachmentsAccountOptions) SetLimit(limit int64) *ListAttachmentsAccountOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAttachmentsAccountOptions) SetHeaders(param map[string]string) *ListAttachmentsAccountOptions {
	options.Headers = param
	return options
}

// ListAttachmentsOptions : The ListAttachments options.
type ListAttachmentsOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// The indication of how many resources to return, unless the response is the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAttachmentsOptions : Instantiate ListAttachmentsOptions
func (*ComplianceV2) NewListAttachmentsOptions(profilesID string) *ListAttachmentsOptions {
	return &ListAttachmentsOptions{
		ProfilesID: core.StringPtr(profilesID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ListAttachmentsOptions) SetProfilesID(profilesID string) *ListAttachmentsOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListAttachmentsOptions) SetXCorrelationID(xCorrelationID string) *ListAttachmentsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ListAttachmentsOptions) SetXRequestID(xRequestID string) *ListAttachmentsOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAttachmentsOptions) SetLimit(limit int64) *ListAttachmentsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAttachmentsOptions) SetHeaders(param map[string]string) *ListAttachmentsOptions {
	options.Headers = param
	return options
}

// ListControlLibrariesOptions : The ListControlLibraries options.
type ListControlLibrariesOptions struct {
	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// The field that indicates how many resources to return, unless the response is the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// The field that indicate how you want the resources to be filtered by.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListControlLibrariesOptions : Instantiate ListControlLibrariesOptions
func (*ComplianceV2) NewListControlLibrariesOptions() *ListControlLibrariesOptions {
	return &ListControlLibrariesOptions{}
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListControlLibrariesOptions) SetXCorrelationID(xCorrelationID string) *ListControlLibrariesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ListControlLibrariesOptions) SetXRequestID(xRequestID string) *ListControlLibrariesOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListControlLibrariesOptions) SetLimit(limit int64) *ListControlLibrariesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetControlLibraryType : Allow user to set ControlLibraryType
func (_options *ListControlLibrariesOptions) SetControlLibraryType(controlLibraryType string) *ListControlLibrariesOptions {
	_options.ControlLibraryType = core.StringPtr(controlLibraryType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListControlLibrariesOptions) SetHeaders(param map[string]string) *ListControlLibrariesOptions {
	options.Headers = param
	return options
}

// ListProfilesOptions : The ListProfiles options.
type ListProfilesOptions struct {
	// The supplied or generated value of this header is logged for a request, and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests, and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// The indication of how many resources to return, unless the response is the last page of resources.
	Limit *int64 `json:"limit,omitempty"`

	// The field that indicate how you want the resources to be filtered by.
	ProfileType *string `json:"profile_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProfilesOptions : Instantiate ListProfilesOptions
func (*ComplianceV2) NewListProfilesOptions() *ListProfilesOptions {
	return &ListProfilesOptions{}
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListProfilesOptions) SetXCorrelationID(xCorrelationID string) *ListProfilesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ListProfilesOptions) SetXRequestID(xRequestID string) *ListProfilesOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListProfilesOptions) SetLimit(limit int64) *ListProfilesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetProfileType : Allow user to set ProfileType
func (_options *ListProfilesOptions) SetProfileType(profileType string) *ListProfilesOptions {
	_options.ProfileType = core.StringPtr(profileType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProfilesOptions) SetHeaders(param map[string]string) *ListProfilesOptions {
	options.Headers = param
	return options
}

// MultiCloudScope : The scope payload for the multi cloud feature.
type MultiCloudScope struct {
	// The environment that relates to this scope.
	Environment *string `json:"environment" validate:"required"`

	// The properties supported for scoping by this environment.
	Properties []PropertyItem `json:"properties" validate:"required"`
}

// NewMultiCloudScope : Instantiate MultiCloudScope (Generic Model Constructor)
func (*ComplianceV2) NewMultiCloudScope(environment string, properties []PropertyItem) (_model *MultiCloudScope, err error) {
	_model = &MultiCloudScope{
		Environment: core.StringPtr(environment),
		Properties:  properties,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMultiCloudScope unmarshals an instance of MultiCloudScope from the specified map of raw messages.
func UnmarshalMultiCloudScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MultiCloudScope)
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginatedCollectionFirst : The reference to the first page of entries.
type PaginatedCollectionFirst struct {
	// The reference URL for the first few entries.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPaginatedCollectionFirst unmarshals an instance of PaginatedCollectionFirst from the specified map of raw messages.
func UnmarshalPaginatedCollectionFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginatedCollectionFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginatedCollectionNext : The reference URL for the next few entries.
type PaginatedCollectionNext struct {
	// The reference URL for the entries.
	Href *string `json:"href,omitempty"`

	// The reference to the start of the list of entries.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginatedCollectionNext unmarshals an instance of PaginatedCollectionNext from the specified map of raw messages.
func UnmarshalPaginatedCollectionNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginatedCollectionNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ParameterInfo : The parameter details.
type ParameterInfo struct {
	// The parameter name.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The parameter display name.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// The parameter type.
	ParameterType *string `json:"parameter_type,omitempty"`
}

// Constants associated with the ParameterInfo.ParameterType property.
// The parameter type.
const (
	ParameterInfo_ParameterType_Boolean    = "boolean"
	ParameterInfo_ParameterType_General    = "general"
	ParameterInfo_ParameterType_IpList     = "ip_list"
	ParameterInfo_ParameterType_Numeric    = "numeric"
	ParameterInfo_ParameterType_String     = "string"
	ParameterInfo_ParameterType_StringList = "string_list"
	ParameterInfo_ParameterType_Timestamp  = "timestamp"
)

// UnmarshalParameterInfo unmarshals an instance of ParameterInfo from the specified map of raw messages.
func UnmarshalParameterInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ParameterInfo)
	err = core.UnmarshalPrimitive(m, "parameter_name", &obj.ParameterName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_display_name", &obj.ParameterDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameter_type", &obj.ParameterType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Profile : The response body of the profile.
type Profile struct {
	// The unique ID of the profile.
	ID *string `json:"id,omitempty"`

	// The profile name.
	ProfileName *string `json:"profile_name,omitempty"`

	// The profile description.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// The profile type, such as custom or predefined.
	ProfileType *string `json:"profile_type,omitempty"`

	// The version status of the profile.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// The version group label of the profile.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The instance ID.
	InstanceID *string `json:"instance_id,omitempty"`

	// The latest version of the profile.
	Latest *bool `json:"latest,omitempty"`

	// The indication of whether hierarchy is enabled for the profile.
	HierarchyEnabled *bool `json:"hierarchy_enabled,omitempty"`

	// The user who created the profile.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the profile was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who updated the profile.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The date when the profile was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The number of controls for the profile.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// The number of parent controls for the profile.
	ControlParentsCount *int64 `json:"control_parents_count,omitempty"`

	// The number of attachments related to this profile.
	AttachmentsCount *int64 `json:"attachments_count,omitempty"`

	// The array of controls that are used to create the profile.
	Controls []ProfileControls `json:"controls,omitempty"`

	// The default parameters of the profile.
	DefaultParameters []DefaultParametersPrototype `json:"default_parameters,omitempty"`
}

// Constants associated with the Profile.ProfileType property.
// The profile type, such as custom or predefined.
const (
	Profile_ProfileType_Custom     = "custom"
	Profile_ProfileType_Predefined = "predefined"
)

// UnmarshalProfile unmarshals an instance of Profile from the specified map of raw messages.
func UnmarshalProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Profile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_name", &obj.ProfileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_description", &obj.ProfileDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_type", &obj.ProfileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_version", &obj.ProfileVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_group_label", &obj.VersionGroupLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest", &obj.Latest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hierarchy_enabled", &obj.HierarchyEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_count", &obj.ControlsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_parents_count", &obj.ControlParentsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attachments_count", &obj.AttachmentsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalProfileControls)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_parameters", &obj.DefaultParameters, UnmarshalDefaultParametersPrototype)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileCollection : The response body to get all profiles that are linked to your account.
type ProfileCollection struct {
	// The number of profiles.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The limit of profiles that can be created.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first page of entries.
	First *PaginatedCollectionFirst `json:"first,omitempty"`

	// The reference URL for the next few entries.
	Next *PaginatedCollectionNext `json:"next,omitempty"`

	// The profiles.
	Profiles []ProfileItem `json:"profiles,omitempty"`
}

// UnmarshalProfileCollection unmarshals an instance of ProfileCollection from the specified map of raw messages.
func UnmarshalProfileCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileCollection)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginatedCollectionFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginatedCollectionNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "profiles", &obj.Profiles, UnmarshalProfileItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileControls : The control details for the profile.
type ProfileControls struct {
	// The ID of the control library that contains the profile.
	ControlLibraryID *string `json:"control_library_id,omitempty"`

	// The unique ID of the control library that contains the profile.
	ControlID *string `json:"control_id,omitempty"`

	// The most recent version of the control library.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The control name.
	ControlName *string `json:"control_name,omitempty"`

	// The control description.
	ControlDescription *string `json:"control_description,omitempty"`

	// The control category.
	ControlCategory *string `json:"control_category,omitempty"`

	// The parent control.
	ControlParent *string `json:"control_parent,omitempty"`

	// The indication of whether the control is required.
	ControlRequirement *bool `json:"control_requirement,omitempty"`

	// The control documentation.
	ControlDocs *ControlDocs `json:"control_docs,omitempty"`

	// The number of control specifications.
	ControlSpecificationsCount *int64 `json:"control_specifications_count,omitempty"`

	// The control specifications.
	ControlSpecifications []ControlSpecifications `json:"control_specifications,omitempty"`
}

// UnmarshalProfileControls unmarshals an instance of ProfileControls from the specified map of raw messages.
func UnmarshalProfileControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileControls)
	err = core.UnmarshalPrimitive(m, "control_library_id", &obj.ControlLibraryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_id", &obj.ControlID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_library_version", &obj.ControlLibraryVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_name", &obj.ControlName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_description", &obj.ControlDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_category", &obj.ControlCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_parent", &obj.ControlParent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_requirement", &obj.ControlRequirement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_docs", &obj.ControlDocs, UnmarshalControlDocs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_specifications_count", &obj.ControlSpecificationsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_specifications", &obj.ControlSpecifications, UnmarshalControlSpecifications)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileControlsPrototype : The control details of a profile.
type ProfileControlsPrototype struct {
	// The ID of the control library that contains the profile.
	ControlLibraryID *string `json:"control_library_id,omitempty"`

	// The control ID.
	ControlID *string `json:"control_id,omitempty"`
}

// UnmarshalProfileControlsPrototype unmarshals an instance of ProfileControlsPrototype from the specified map of raw messages.
func UnmarshalProfileControlsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileControlsPrototype)
	err = core.UnmarshalPrimitive(m, "control_library_id", &obj.ControlLibraryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "control_id", &obj.ControlID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileItem : ProfileItem struct
type ProfileItem struct {
	// The profile ID.
	ID *string `json:"id,omitempty"`

	// The profile name.
	ProfileName *string `json:"profile_name,omitempty"`

	// The profile description.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// The profile type.
	ProfileType *string `json:"profile_type,omitempty"`

	// The profile version.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// The version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The latest profile.
	Latest *bool `json:"latest,omitempty"`

	// The user who created the profile.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the profile was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who updated the profile.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The date when the profile was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// The number of attachments.
	AttachmentsCount *int64 `json:"attachments_count,omitempty"`
}

// UnmarshalProfileItem unmarshals an instance of ProfileItem from the specified map of raw messages.
func UnmarshalProfileItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_name", &obj.ProfileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_description", &obj.ProfileDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_type", &obj.ProfileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_version", &obj.ProfileVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_group_label", &obj.VersionGroupLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest", &obj.Latest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_on", &obj.UpdatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_count", &obj.ControlsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attachments_count", &obj.AttachmentsCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Property : The properties supported for scoping by this environment.
type PropertyItem struct {
	// The name of the property.
	Name *string `json:"name,omitempty"`

	// The value of the property.
	Value interface{} `json:"value,omitempty"`
}

// UnmarshalProperty unmarshals an instance of Property from the specified map of raw messages.
func UnmarshalProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PropertyItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// ReplaceCustomControlLibraryOptions : The ReplaceCustomControlLibrary options.
type ReplaceCustomControlLibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// The control library ID.
	ID *string `json:"id,omitempty"`

	// The account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The control library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// The control library description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// The control library type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// The version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// The control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The date when the control library was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who created the control library.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the control library was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The user who updated the control library.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The latest version of the control library.
	Latest *bool `json:"latest,omitempty"`

	// The indication of whether hierarchy is enabled for the profile.
	HierarchyEnabled *bool `json:"hierarchy_enabled,omitempty"`

	// The number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// The number of parent controls in the profile.
	ControlParentsCount *int64 `json:"control_parents_count,omitempty"`

	// The controls.
	Controls []ControlsInControlLib `json:"controls,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceCustomControlLibraryOptions.ControlLibraryType property.
// The control library type.
const (
	ReplaceCustomControlLibraryOptions_ControlLibraryType_Custom     = "custom"
	ReplaceCustomControlLibraryOptions_ControlLibraryType_Predefined = "predefined"
)

// NewReplaceCustomControlLibraryOptions : Instantiate ReplaceCustomControlLibraryOptions
func (*ComplianceV2) NewReplaceCustomControlLibraryOptions(controlLibrariesID string) *ReplaceCustomControlLibraryOptions {
	return &ReplaceCustomControlLibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibrariesID(controlLibrariesID string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ReplaceCustomControlLibraryOptions) SetID(id string) *ReplaceCustomControlLibraryOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceCustomControlLibraryOptions) SetAccountID(accountID string) *ReplaceCustomControlLibraryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetControlLibraryName : Allow user to set ControlLibraryName
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibraryName(controlLibraryName string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibraryName = core.StringPtr(controlLibraryName)
	return _options
}

// SetControlLibraryDescription : Allow user to set ControlLibraryDescription
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibraryDescription(controlLibraryDescription string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibraryDescription = core.StringPtr(controlLibraryDescription)
	return _options
}

// SetControlLibraryType : Allow user to set ControlLibraryType
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibraryType(controlLibraryType string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibraryType = core.StringPtr(controlLibraryType)
	return _options
}

// SetVersionGroupLabel : Allow user to set VersionGroupLabel
func (_options *ReplaceCustomControlLibraryOptions) SetVersionGroupLabel(versionGroupLabel string) *ReplaceCustomControlLibraryOptions {
	_options.VersionGroupLabel = core.StringPtr(versionGroupLabel)
	return _options
}

// SetControlLibraryVersion : Allow user to set ControlLibraryVersion
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibraryVersion(controlLibraryVersion string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibraryVersion = core.StringPtr(controlLibraryVersion)
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *ReplaceCustomControlLibraryOptions) SetCreatedOn(createdOn *strfmt.DateTime) *ReplaceCustomControlLibraryOptions {
	_options.CreatedOn = createdOn
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *ReplaceCustomControlLibraryOptions) SetCreatedBy(createdBy string) *ReplaceCustomControlLibraryOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetUpdatedOn : Allow user to set UpdatedOn
func (_options *ReplaceCustomControlLibraryOptions) SetUpdatedOn(updatedOn *strfmt.DateTime) *ReplaceCustomControlLibraryOptions {
	_options.UpdatedOn = updatedOn
	return _options
}

// SetUpdatedBy : Allow user to set UpdatedBy
func (_options *ReplaceCustomControlLibraryOptions) SetUpdatedBy(updatedBy string) *ReplaceCustomControlLibraryOptions {
	_options.UpdatedBy = core.StringPtr(updatedBy)
	return _options
}

// SetLatest : Allow user to set Latest
func (_options *ReplaceCustomControlLibraryOptions) SetLatest(latest bool) *ReplaceCustomControlLibraryOptions {
	_options.Latest = core.BoolPtr(latest)
	return _options
}

// SetHierarchyEnabled : Allow user to set HierarchyEnabled
func (_options *ReplaceCustomControlLibraryOptions) SetHierarchyEnabled(hierarchyEnabled bool) *ReplaceCustomControlLibraryOptions {
	_options.HierarchyEnabled = core.BoolPtr(hierarchyEnabled)
	return _options
}

// SetControlsCount : Allow user to set ControlsCount
func (_options *ReplaceCustomControlLibraryOptions) SetControlsCount(controlsCount int64) *ReplaceCustomControlLibraryOptions {
	_options.ControlsCount = core.Int64Ptr(controlsCount)
	return _options
}

// SetControlParentsCount : Allow user to set ControlParentsCount
func (_options *ReplaceCustomControlLibraryOptions) SetControlParentsCount(controlParentsCount int64) *ReplaceCustomControlLibraryOptions {
	_options.ControlParentsCount = core.Int64Ptr(controlParentsCount)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *ReplaceCustomControlLibraryOptions) SetControls(controls []ControlsInControlLib) *ReplaceCustomControlLibraryOptions {
	_options.Controls = controls
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ReplaceCustomControlLibraryOptions) SetXCorrelationID(xCorrelationID string) *ReplaceCustomControlLibraryOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ReplaceCustomControlLibraryOptions) SetXRequestID(xRequestID string) *ReplaceCustomControlLibraryOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceCustomControlLibraryOptions) SetHeaders(param map[string]string) *ReplaceCustomControlLibraryOptions {
	options.Headers = param
	return options
}

// ReplaceProfileAttachmentOptions : The ReplaceProfileAttachment options.
type ReplaceProfileAttachmentOptions struct {
	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The ID of the attachment.
	ID *string `json:"id,omitempty"`

	// The ID of the profile that is specified in the attachment.
	ProfileID *string `json:"profile_id,omitempty"`

	// The account ID that is associated to the attachment.
	AccountID *string `json:"account_id,omitempty"`

	// The instance ID of the account that is associated to the attachment.
	InstanceID *string `json:"instance_id,omitempty"`

	// The scope payload for the multi cloud feature.
	Scope []MultiCloudScope `json:"scope,omitempty"`

	// The date when the attachment was created.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// The user who created the attachment.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date when the attachment was updated.
	UpdatedOn *strfmt.DateTime `json:"updated_on,omitempty"`

	// The user who updated the attachment.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The status of an attachment evaluation.
	Status *string `json:"status,omitempty"`

	// The schedule of an attachment evaluation.
	Schedule *string `json:"schedule,omitempty"`

	// The request payload of the attachment notifications.
	Notifications *AttachmentsNotificationsPrototype `json:"notifications,omitempty"`

	// The profile parameters for the attachment.
	AttachmentParameters []AttachmentParametersPrototype `json:"attachment_parameters,omitempty"`

	// The details of the last scan of an attachment.
	LastScan *LastScan `json:"last_scan,omitempty"`

	// The start time of the next scan.
	NextScanTime *string `json:"next_scan_time,omitempty"`

	// The name that is generated from the scope type and ID.
	Name *string `json:"name,omitempty"`

	// The description for the attachment.
	Description *string `json:"description,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceProfileAttachmentOptions.Status property.
// The status of an attachment evaluation.
const (
	ReplaceProfileAttachmentOptions_Status_Disabled = "disabled"
	ReplaceProfileAttachmentOptions_Status_Enabled  = "enabled"
)

// Constants associated with the ReplaceProfileAttachmentOptions.Schedule property.
// The schedule of an attachment evaluation.
const (
	ReplaceProfileAttachmentOptions_Schedule_Daily       = "daily"
	ReplaceProfileAttachmentOptions_Schedule_Every30Days = "every_30_days"
	ReplaceProfileAttachmentOptions_Schedule_Every7Days  = "every_7_days"
)

// NewReplaceProfileAttachmentOptions : Instantiate ReplaceProfileAttachmentOptions
func (*ComplianceV2) NewReplaceProfileAttachmentOptions(attachmentID string, profilesID string) *ReplaceProfileAttachmentOptions {
	return &ReplaceProfileAttachmentOptions{
		AttachmentID: core.StringPtr(attachmentID),
		ProfilesID:   core.StringPtr(profilesID),
	}
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ReplaceProfileAttachmentOptions) SetAttachmentID(attachmentID string) *ReplaceProfileAttachmentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceProfileAttachmentOptions) SetProfilesID(profilesID string) *ReplaceProfileAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ReplaceProfileAttachmentOptions) SetID(id string) *ReplaceProfileAttachmentOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetProfileID : Allow user to set ProfileID
func (_options *ReplaceProfileAttachmentOptions) SetProfileID(profileID string) *ReplaceProfileAttachmentOptions {
	_options.ProfileID = core.StringPtr(profileID)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceProfileAttachmentOptions) SetAccountID(accountID string) *ReplaceProfileAttachmentOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceProfileAttachmentOptions) SetInstanceID(instanceID string) *ReplaceProfileAttachmentOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetScope : Allow user to set Scope
func (_options *ReplaceProfileAttachmentOptions) SetScope(scope []MultiCloudScope) *ReplaceProfileAttachmentOptions {
	_options.Scope = scope
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *ReplaceProfileAttachmentOptions) SetCreatedOn(createdOn *strfmt.DateTime) *ReplaceProfileAttachmentOptions {
	_options.CreatedOn = createdOn
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *ReplaceProfileAttachmentOptions) SetCreatedBy(createdBy string) *ReplaceProfileAttachmentOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetUpdatedOn : Allow user to set UpdatedOn
func (_options *ReplaceProfileAttachmentOptions) SetUpdatedOn(updatedOn *strfmt.DateTime) *ReplaceProfileAttachmentOptions {
	_options.UpdatedOn = updatedOn
	return _options
}

// SetUpdatedBy : Allow user to set UpdatedBy
func (_options *ReplaceProfileAttachmentOptions) SetUpdatedBy(updatedBy string) *ReplaceProfileAttachmentOptions {
	_options.UpdatedBy = core.StringPtr(updatedBy)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ReplaceProfileAttachmentOptions) SetStatus(status string) *ReplaceProfileAttachmentOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetSchedule : Allow user to set Schedule
func (_options *ReplaceProfileAttachmentOptions) SetSchedule(schedule string) *ReplaceProfileAttachmentOptions {
	_options.Schedule = core.StringPtr(schedule)
	return _options
}

// SetNotifications : Allow user to set Notifications
func (_options *ReplaceProfileAttachmentOptions) SetNotifications(notifications *AttachmentsNotificationsPrototype) *ReplaceProfileAttachmentOptions {
	_options.Notifications = notifications
	return _options
}

// SetAttachmentParameters : Allow user to set AttachmentParameters
func (_options *ReplaceProfileAttachmentOptions) SetAttachmentParameters(attachmentParameters []AttachmentParametersPrototype) *ReplaceProfileAttachmentOptions {
	_options.AttachmentParameters = attachmentParameters
	return _options
}

// SetLastScan : Allow user to set LastScan
func (_options *ReplaceProfileAttachmentOptions) SetLastScan(lastScan *LastScan) *ReplaceProfileAttachmentOptions {
	_options.LastScan = lastScan
	return _options
}

// SetNextScanTime : Allow user to set NextScanTime
func (_options *ReplaceProfileAttachmentOptions) SetNextScanTime(nextScanTime string) *ReplaceProfileAttachmentOptions {
	_options.NextScanTime = core.StringPtr(nextScanTime)
	return _options
}

// SetName : Allow user to set Name
func (_options *ReplaceProfileAttachmentOptions) SetName(name string) *ReplaceProfileAttachmentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ReplaceProfileAttachmentOptions) SetDescription(description string) *ReplaceProfileAttachmentOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ReplaceProfileAttachmentOptions) SetXCorrelationID(xCorrelationID string) *ReplaceProfileAttachmentOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ReplaceProfileAttachmentOptions) SetXRequestID(xRequestID string) *ReplaceProfileAttachmentOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceProfileAttachmentOptions) SetHeaders(param map[string]string) *ReplaceProfileAttachmentOptions {
	options.Headers = param
	return options
}

// ReplaceProfileOptions : The ReplaceProfile options.
type ReplaceProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The name of the profile.
	ProfileName *string `json:"profile_name" validate:"required"`

	// The description of the profile.
	ProfileDescription *string `json:"profile_description" validate:"required"`

	// The profile type.
	ProfileType *string `json:"profile_type" validate:"required"`

	// The controls that are in the profile.
	Controls []ProfileControlsPrototype `json:"controls" validate:"required"`

	// The default parameters of the profile.
	DefaultParameters []DefaultParametersPrototype `json:"default_parameters" validate:"required"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is used for downstream requests and retries of those requests. If a value of
	// this header is not supplied in a request, the service generates a random (version 4) UUID.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// The supplied or generated value of this header is logged for a request and repeated in a response header for the
	// corresponding response. The same value is not used for downstream requests and retries of those requests. If a value
	// of this header is not supplied in a request, the service generates a random (version 4) UUID.
	XRequestID *string `json:"X-Request-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceProfileOptions.ProfileType property.
// The profile type.
const (
	ReplaceProfileOptions_ProfileType_Custom     = "custom"
	ReplaceProfileOptions_ProfileType_Predefined = "predefined"
)

// NewReplaceProfileOptions : Instantiate ReplaceProfileOptions
func (*ComplianceV2) NewReplaceProfileOptions(profilesID string, profileName string, profileDescription string, profileType string, controls []ProfileControlsPrototype, defaultParameters []DefaultParametersPrototype) *ReplaceProfileOptions {
	return &ReplaceProfileOptions{
		ProfilesID:         core.StringPtr(profilesID),
		ProfileName:        core.StringPtr(profileName),
		ProfileDescription: core.StringPtr(profileDescription),
		ProfileType:        core.StringPtr(profileType),
		Controls:           controls,
		DefaultParameters:  defaultParameters,
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceProfileOptions) SetProfilesID(profilesID string) *ReplaceProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetProfileName : Allow user to set ProfileName
func (_options *ReplaceProfileOptions) SetProfileName(profileName string) *ReplaceProfileOptions {
	_options.ProfileName = core.StringPtr(profileName)
	return _options
}

// SetProfileDescription : Allow user to set ProfileDescription
func (_options *ReplaceProfileOptions) SetProfileDescription(profileDescription string) *ReplaceProfileOptions {
	_options.ProfileDescription = core.StringPtr(profileDescription)
	return _options
}

// SetProfileType : Allow user to set ProfileType
func (_options *ReplaceProfileOptions) SetProfileType(profileType string) *ReplaceProfileOptions {
	_options.ProfileType = core.StringPtr(profileType)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *ReplaceProfileOptions) SetControls(controls []ProfileControlsPrototype) *ReplaceProfileOptions {
	_options.Controls = controls
	return _options
}

// SetDefaultParameters : Allow user to set DefaultParameters
func (_options *ReplaceProfileOptions) SetDefaultParameters(defaultParameters []DefaultParametersPrototype) *ReplaceProfileOptions {
	_options.DefaultParameters = defaultParameters
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ReplaceProfileOptions) SetXCorrelationID(xCorrelationID string) *ReplaceProfileOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetXRequestID : Allow user to set XRequestID
func (_options *ReplaceProfileOptions) SetXRequestID(xRequestID string) *ReplaceProfileOptions {
	_options.XRequestID = core.StringPtr(xRequestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceProfileOptions) SetHeaders(param map[string]string) *ReplaceProfileOptions {
	options.Headers = param
	return options
}

// Scan : The response schema for creating a scan.
type Scan struct {
	// The scan ID.
	ID *string `json:"id,omitempty"`

	// The account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The attachment ID of a profile.
	AttachmentID *string `json:"attachment_id,omitempty"`

	// The report ID.
	ReportID *string `json:"report_id,omitempty"`

	// The status of the scan.
	Status *string `json:"status,omitempty"`

	// The last scan time.
	LastScanTime *string `json:"last_scan_time,omitempty"`

	// The next scan time.
	NextScanTime *string `json:"next_scan_time,omitempty"`

	// The type of scan.
	ScanType *string `json:"scan_type,omitempty"`

	// The occurrence of the scan.
	Occurence *int64 `json:"occurence,omitempty"`
}

// Constants associated with the Scan.Status property.
// The status of the scan.
const (
	Scan_Status_Completed  = "completed"
	Scan_Status_InProgress = "in_progress"
)

// Constants associated with the Scan.ScanType property.
// The type of scan.
const (
	Scan_ScanType_Ondemand  = "ondemand"
	Scan_ScanType_Scheduled = "scheduled"
)

// UnmarshalScan unmarshals an instance of Scan from the specified map of raw messages.
func UnmarshalScan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scan)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "report_id", &obj.ReportID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_time", &obj.LastScanTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_scan_time", &obj.NextScanTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scan_type", &obj.ScanType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "occurence", &obj.Occurence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

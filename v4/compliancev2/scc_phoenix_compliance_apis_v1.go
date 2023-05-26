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
 * IBM OpenAPI SDK Code Generator Version: 3.72.0-5d70f2bb-20230511-203609
 */

// Package compliancev2 : Operations and models for the SccPhoenixComplianceApisV1 service
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
)

// SccPhoenixComplianceApisV1 : The SCC Phoenix Compliance APIs.
//
// API Version: 1.0.0
type SccPhoenixComplianceApisV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "scc_phoenix_compliance_apis"

// SccPhoenixComplianceApisV1Options : Service options
type SccPhoenixComplianceApisV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSccPhoenixComplianceApisV1UsingExternalConfig : constructs an instance of SccPhoenixComplianceApisV1 with passed in options and external configuration.
func NewSccPhoenixComplianceApisV1UsingExternalConfig(options *SccPhoenixComplianceApisV1Options) (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	sccPhoenixComplianceApis, err = NewSccPhoenixComplianceApisV1(options)
	if err != nil {
		return
	}

	err = sccPhoenixComplianceApis.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = sccPhoenixComplianceApis.Service.SetServiceURL(options.URL)
	}
	return
}

// NewSccPhoenixComplianceApisV1 : constructs an instance of SccPhoenixComplianceApisV1 with passed in options.
func NewSccPhoenixComplianceApisV1(options *SccPhoenixComplianceApisV1Options) (service *SccPhoenixComplianceApisV1, err error) {
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

	service = &SccPhoenixComplianceApisV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "sccPhoenixComplianceApis" suitable for processing requests.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) Clone() *SccPhoenixComplianceApisV1 {
	if core.IsNil(sccPhoenixComplianceApis) {
		return nil
	}
	clone := *sccPhoenixComplianceApis
	clone.Service = sccPhoenixComplianceApis.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) SetServiceURL(url string) error {
	return sccPhoenixComplianceApis.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetServiceURL() string {
	return sccPhoenixComplianceApis.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) SetDefaultHeaders(headers http.Header) {
	sccPhoenixComplianceApis.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) SetEnableGzipCompression(enableGzip bool) {
	sccPhoenixComplianceApis.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetEnableGzipCompression() bool {
	return sccPhoenixComplianceApis.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	sccPhoenixComplianceApis.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DisableRetries() {
	sccPhoenixComplianceApis.Service.DisableRetries()
}

// CreateProfile : Create a custom profile
// Create a user-defined custom profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateProfile(createProfileOptions *CreateProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.CreateProfileWithContext(context.Background(), createProfileOptions)
}

// CreateProfileWithContext is an alternate form of the CreateProfile method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateProfileWithContext(ctx context.Context, createProfileOptions *CreateProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProfileOptions, "createProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProfileOptions, "createProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createProfileOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "CreateProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createProfileOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createProfileOptions.TransactionID))
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
	if createProfileOptions.ProfileVersion != nil {
		body["profile_version"] = createProfileOptions.ProfileVersion
	}
	if createProfileOptions.Latest != nil {
		body["latest"] = createProfileOptions.Latest
	}
	if createProfileOptions.VersionGroupLabel != nil {
		body["version_group_label"] = createProfileOptions.VersionGroupLabel
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProfiles : Get all predefined and user's custom profiles
// Get all predefined and user's custom profiles.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListProfiles(listProfilesOptions *ListProfilesOptions) (result *GetAllProfilesRespBody, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ListProfilesWithContext(context.Background(), listProfilesOptions)
}

// ListProfilesWithContext is an alternate form of the ListProfiles method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListProfilesWithContext(ctx context.Context, listProfilesOptions *ListProfilesOptions) (result *GetAllProfilesRespBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listProfilesOptions, "listProfilesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listProfilesOptions, "listProfilesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listProfilesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ListProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listProfilesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listProfilesOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAllProfilesRespBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddProfile : Update a custom profile
// Update a user-defined custom profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) AddProfile(addProfileOptions *AddProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.AddProfileWithContext(context.Background(), addProfileOptions)
}

// AddProfileWithContext is an alternate form of the AddProfile method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) AddProfileWithContext(ctx context.Context, addProfileOptions *AddProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addProfileOptions, "addProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addProfileOptions, "addProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *addProfileOptions.ProfilesID,
		"instance_id": *addProfileOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "AddProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addProfileOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addProfileOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if addProfileOptions.ProfileName != nil {
		body["profile_name"] = addProfileOptions.ProfileName
	}
	if addProfileOptions.ProfileDescription != nil {
		body["profile_description"] = addProfileOptions.ProfileDescription
	}
	if addProfileOptions.ProfileType != nil {
		body["profile_type"] = addProfileOptions.ProfileType
	}
	if addProfileOptions.ProfileVersion != nil {
		body["profile_version"] = addProfileOptions.ProfileVersion
	}
	if addProfileOptions.Latest != nil {
		body["latest"] = addProfileOptions.Latest
	}
	if addProfileOptions.VersionGroupLabel != nil {
		body["version_group_label"] = addProfileOptions.VersionGroupLabel
	}
	if addProfileOptions.Controls != nil {
		body["controls"] = addProfileOptions.Controls
	}
	if addProfileOptions.DefaultParameters != nil {
		body["default_parameters"] = addProfileOptions.DefaultParameters
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProfile : Get a profile
// Retrieve a profile by specifying the profile ID.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetProfile(getProfileOptions *GetProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.GetProfileWithContext(context.Background(), getProfileOptions)
}

// GetProfileWithContext is an alternate form of the GetProfile method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetProfileWithContext(ctx context.Context, getProfileOptions *GetProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
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
		"instance_id": *getProfileOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "GetProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getProfileOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getProfileOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomProfile : Delete a custom profile
// Delete a custom profile by specifying the profile ID.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.DeleteCustomProfileWithContext(context.Background(), deleteCustomProfileOptions)
}

// DeleteCustomProfileWithContext is an alternate form of the DeleteCustomProfile method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteCustomProfileWithContext(ctx context.Context, deleteCustomProfileOptions *DeleteCustomProfileOptions) (result *ProfileResponse, response *core.DetailedResponse, err error) {
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
		"instance_id": *deleteCustomProfileOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "DeleteCustomProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteCustomProfileOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteCustomProfileOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceProfileParameters : Update custom profile parameters
// Update the parameters of a custom profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions) (result *ProfileDefaultParametersResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ReplaceProfileParametersWithContext(context.Background(), replaceProfileParametersOptions)
}

// ReplaceProfileParametersWithContext is an alternate form of the ReplaceProfileParameters method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceProfileParametersWithContext(ctx context.Context, replaceProfileParametersOptions *ReplaceProfileParametersOptions) (result *ProfileDefaultParametersResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceProfileParametersOptions, "replaceProfileParametersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceProfileParametersOptions, "replaceProfileParametersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *replaceProfileParametersOptions.ProfilesID,
		"instance_id": *replaceProfileParametersOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/parameters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceProfileParametersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ReplaceProfileParameters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceProfileParametersOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceProfileParametersOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceProfileParametersOptions.ID != nil {
		body["id"] = replaceProfileParametersOptions.ID
	}
	if replaceProfileParametersOptions.DefaultParameters != nil {
		body["default_parameters"] = replaceProfileParametersOptions.DefaultParameters
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfileDefaultParametersResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateAttachment : Create an attachment
// Create an attachment to link to a profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateAttachment(createAttachmentOptions *CreateAttachmentOptions) (result *AttachmentProfileResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.CreateAttachmentWithContext(context.Background(), createAttachmentOptions)
}

// CreateAttachmentWithContext is an alternate form of the CreateAttachment method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateAttachmentWithContext(ctx context.Context, createAttachmentOptions *CreateAttachmentOptions) (result *AttachmentProfileResponse, response *core.DetailedResponse, err error) {
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
		"instance_id": *createAttachmentOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "CreateAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createAttachmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createAttachmentOptions.Attachments != nil {
		body["attachments"] = createAttachmentOptions.Attachments
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentProfileResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CheckProfileAttachmnets : Get all attachments
// Retrieve all attachments that are linked to a profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CheckProfileAttachmnets(checkProfileAttachmnetsOptions *CheckProfileAttachmnetsOptions) (result *GetAllAttachmnetsForProfileRespBody, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.CheckProfileAttachmnetsWithContext(context.Background(), checkProfileAttachmnetsOptions)
}

// CheckProfileAttachmnetsWithContext is an alternate form of the CheckProfileAttachmnets method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CheckProfileAttachmnetsWithContext(ctx context.Context, checkProfileAttachmnetsOptions *CheckProfileAttachmnetsOptions) (result *GetAllAttachmnetsForProfileRespBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(checkProfileAttachmnetsOptions, "checkProfileAttachmnetsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(checkProfileAttachmnetsOptions, "checkProfileAttachmnetsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id": *checkProfileAttachmnetsOptions.ProfilesID,
		"instance_id": *checkProfileAttachmnetsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkProfileAttachmnetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "CheckProfileAttachmnets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if checkProfileAttachmnetsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*checkProfileAttachmnetsOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAllAttachmnetsForProfileRespBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProfileAttachmnet : Get an attachment for a profile
// Retrieve an attachment that is linked to a profile by specifying the attachment ID.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetProfileAttachmnet(getProfileAttachmnetOptions *GetProfileAttachmnetOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.GetProfileAttachmnetWithContext(context.Background(), getProfileAttachmnetOptions)
}

// GetProfileAttachmnetWithContext is an alternate form of the GetProfileAttachmnet method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetProfileAttachmnetWithContext(ctx context.Context, getProfileAttachmnetOptions *GetProfileAttachmnetOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProfileAttachmnetOptions, "getProfileAttachmnetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProfileAttachmnetOptions, "getProfileAttachmnetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":   *getProfileAttachmnetOptions.ProfilesID,
		"attachment_id": *getProfileAttachmnetOptions.AttachmentID,
		"instance_id":   *getProfileAttachmnetOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfileAttachmnetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "GetProfileAttachmnet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getProfileAttachmnetOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getProfileAttachmnetOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentPayload)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceProfileAttachment : Update an attachment
// Update an attachment that is linked to a profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ReplaceProfileAttachmentWithContext(context.Background(), replaceProfileAttachmentOptions)
}

// ReplaceProfileAttachmentWithContext is an alternate form of the ReplaceProfileAttachment method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceProfileAttachmentWithContext(ctx context.Context, replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceProfileAttachmentOptions, "replaceProfileAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceProfileAttachmentOptions, "replaceProfileAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":   *replaceProfileAttachmentOptions.ProfilesID,
		"attachment_id": *replaceProfileAttachmentOptions.AttachmentID,
		"instance_id":   *replaceProfileAttachmentOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceProfileAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ReplaceProfileAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceProfileAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceProfileAttachmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceProfileAttachmentOptions.ID != nil {
		body["id"] = replaceProfileAttachmentOptions.ID
	}
	if replaceProfileAttachmentOptions.AccountID != nil {
		body["account_id"] = replaceProfileAttachmentOptions.AccountID
	}
	if replaceProfileAttachmentOptions.IncludedScope != nil {
		body["included_scope"] = replaceProfileAttachmentOptions.IncludedScope
	}
	if replaceProfileAttachmentOptions.Exclusions != nil {
		body["exclusions"] = replaceProfileAttachmentOptions.Exclusions
	}
	if replaceProfileAttachmentOptions.CreatedBy != nil {
		body["created_by"] = replaceProfileAttachmentOptions.CreatedBy
	}
	if replaceProfileAttachmentOptions.CreatedOn != nil {
		body["created_on"] = replaceProfileAttachmentOptions.CreatedOn
	}
	if replaceProfileAttachmentOptions.UpdatedBy != nil {
		body["updated_by"] = replaceProfileAttachmentOptions.UpdatedBy
	}
	if replaceProfileAttachmentOptions.UpdatedOn != nil {
		body["updated_on"] = replaceProfileAttachmentOptions.UpdatedOn
	}
	if replaceProfileAttachmentOptions.Status != nil {
		body["status"] = replaceProfileAttachmentOptions.Status
	}
	if replaceProfileAttachmentOptions.AttachmentParameters != nil {
		body["attachment_parameters"] = replaceProfileAttachmentOptions.AttachmentParameters
	}
	if replaceProfileAttachmentOptions.AttachmentNotifications != nil {
		body["attachment_notifications"] = replaceProfileAttachmentOptions.AttachmentNotifications
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentPayload)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteProfileAttachmnet : Delete an attachment
// Delete an attachment that is linked to a profile.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteProfileAttachmnet(deleteProfileAttachmnetOptions *DeleteProfileAttachmnetOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.DeleteProfileAttachmnetWithContext(context.Background(), deleteProfileAttachmnetOptions)
}

// DeleteProfileAttachmnetWithContext is an alternate form of the DeleteProfileAttachmnet method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteProfileAttachmnetWithContext(ctx context.Context, deleteProfileAttachmnetOptions *DeleteProfileAttachmnetOptions) (result *AttachmentPayload, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProfileAttachmnetOptions, "deleteProfileAttachmnetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProfileAttachmnetOptions, "deleteProfileAttachmnetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":   *deleteProfileAttachmnetOptions.ProfilesID,
		"attachment_id": *deleteProfileAttachmnetOptions.AttachmentID,
		"instance_id":   *deleteProfileAttachmnetOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProfileAttachmnetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "DeleteProfileAttachmnet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteProfileAttachmnetOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteProfileAttachmnetOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentPayload)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAttachmentParameters : Get attachment's parameters
// Get attachment's parameters.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListAttachmentParameters(listAttachmentParametersOptions *ListAttachmentParametersOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ListAttachmentParametersWithContext(context.Background(), listAttachmentParametersOptions)
}

// ListAttachmentParametersWithContext is an alternate form of the ListAttachmentParameters method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListAttachmentParametersWithContext(ctx context.Context, listAttachmentParametersOptions *ListAttachmentParametersOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAttachmentParametersOptions, "listAttachmentParametersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAttachmentParametersOptions, "listAttachmentParametersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":   *listAttachmentParametersOptions.ProfilesID,
		"attachment_id": *listAttachmentParametersOptions.AttachmentID,
		"instance_id":   *listAttachmentParametersOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}/parameters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAttachmentParametersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ListAttachmentParameters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAttachmentParametersOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAttachmentParametersOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalParameterDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceAttachment : Update parameters for an attachment
// Update parameters for an attachment.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceAttachment(replaceAttachmentOptions *ReplaceAttachmentOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ReplaceAttachmentWithContext(context.Background(), replaceAttachmentOptions)
}

// ReplaceAttachmentWithContext is an alternate form of the ReplaceAttachment method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceAttachmentWithContext(ctx context.Context, replaceAttachmentOptions *ReplaceAttachmentOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceAttachmentOptions, "replaceAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceAttachmentOptions, "replaceAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":   *replaceAttachmentOptions.ProfilesID,
		"attachment_id": *replaceAttachmentOptions.AttachmentID,
		"instance_id":   *replaceAttachmentOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}/parameters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ReplaceAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceAttachmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceAttachmentOptions.ParameterName != nil {
		body["parameter_name"] = replaceAttachmentOptions.ParameterName
	}
	if replaceAttachmentOptions.ParameterDisplayName != nil {
		body["parameter_display_name"] = replaceAttachmentOptions.ParameterDisplayName
	}
	if replaceAttachmentOptions.ParameterType != nil {
		body["parameter_type"] = replaceAttachmentOptions.ParameterType
	}
	if replaceAttachmentOptions.ParameterValue != nil {
		body["parameter_value"] = replaceAttachmentOptions.ParameterValue
	}
	if replaceAttachmentOptions.AssessmentType != nil {
		body["assessment_type"] = replaceAttachmentOptions.AssessmentType
	}
	if replaceAttachmentOptions.AssessmentID != nil {
		body["assessment_id"] = replaceAttachmentOptions.AssessmentID
	}
	if replaceAttachmentOptions.Parameters != nil {
		body["parameters"] = replaceAttachmentOptions.Parameters
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalParameterDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetParametersByName : Get parameters by name
// Get parametes by name.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetParametersByName(getParametersByNameOptions *GetParametersByNameOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.GetParametersByNameWithContext(context.Background(), getParametersByNameOptions)
}

// GetParametersByNameWithContext is an alternate form of the GetParametersByName method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetParametersByNameWithContext(ctx context.Context, getParametersByNameOptions *GetParametersByNameOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getParametersByNameOptions, "getParametersByNameOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getParametersByNameOptions, "getParametersByNameOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":    *getParametersByNameOptions.ProfilesID,
		"attachment_id":  *getParametersByNameOptions.AttachmentID,
		"parameter_name": *getParametersByNameOptions.ParameterName,
		"instance_id":    *getParametersByNameOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}/parameters/{parameter_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getParametersByNameOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "GetParametersByName")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getParametersByNameOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getParametersByNameOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalParameterDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceAttachmnetParametersByName : Update parameter by name
// Update parameter by name.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions *ReplaceAttachmnetParametersByNameOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ReplaceAttachmnetParametersByNameWithContext(context.Background(), replaceAttachmnetParametersByNameOptions)
}

// ReplaceAttachmnetParametersByNameWithContext is an alternate form of the ReplaceAttachmnetParametersByName method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceAttachmnetParametersByNameWithContext(ctx context.Context, replaceAttachmnetParametersByNameOptions *ReplaceAttachmnetParametersByNameOptions) (result *ParameterDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceAttachmnetParametersByNameOptions, "replaceAttachmnetParametersByNameOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceAttachmnetParametersByNameOptions, "replaceAttachmnetParametersByNameOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"profiles_id":    *replaceAttachmnetParametersByNameOptions.ProfilesID,
		"attachment_id":  *replaceAttachmnetParametersByNameOptions.AttachmentID,
		"parameter_name": *replaceAttachmnetParametersByNameOptions.ParameterName,
		"instance_id":    *replaceAttachmnetParametersByNameOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/profiles/{profiles_id}/attachments/{attachment_id}/parameters/{parameter_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceAttachmnetParametersByNameOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ReplaceAttachmnetParametersByName")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceAttachmnetParametersByNameOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceAttachmnetParametersByNameOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceAttachmnetParametersByNameOptions.NewParameterName != nil {
		body["parameter_name"] = replaceAttachmnetParametersByNameOptions.NewParameterName
	}
	if replaceAttachmnetParametersByNameOptions.NewParameterDisplayName != nil {
		body["parameter_display_name"] = replaceAttachmnetParametersByNameOptions.NewParameterDisplayName
	}
	if replaceAttachmnetParametersByNameOptions.NewParameterType != nil {
		body["parameter_type"] = replaceAttachmnetParametersByNameOptions.NewParameterType
	}
	if replaceAttachmnetParametersByNameOptions.NewParameterValue != nil {
		body["parameter_value"] = replaceAttachmnetParametersByNameOptions.NewParameterValue
	}
	if replaceAttachmnetParametersByNameOptions.NewAssessmentType != nil {
		body["assessment_type"] = replaceAttachmnetParametersByNameOptions.NewAssessmentType
	}
	if replaceAttachmnetParametersByNameOptions.NewAssessmentID != nil {
		body["assessment_id"] = replaceAttachmnetParametersByNameOptions.NewAssessmentID
	}
	if replaceAttachmnetParametersByNameOptions.NewParameters != nil {
		body["parameters"] = replaceAttachmnetParametersByNameOptions.NewParameters
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalParameterDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomControlLibrary : Create a custom control library
// Create a custom control library.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.CreateCustomControlLibraryWithContext(context.Background(), createCustomControlLibraryOptions)
}

// CreateCustomControlLibraryWithContext is an alternate form of the CreateCustomControlLibrary method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateCustomControlLibraryWithContext(ctx context.Context, createCustomControlLibraryOptions *CreateCustomControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomControlLibraryOptions, "createCustomControlLibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomControlLibraryOptions, "createCustomControlLibraryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createCustomControlLibraryOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/control_libraries`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "CreateCustomControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createCustomControlLibraryOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createCustomControlLibraryOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createCustomControlLibraryOptions.ID != nil {
		body["id"] = createCustomControlLibraryOptions.ID
	}
	if createCustomControlLibraryOptions.AccountID != nil {
		body["account_id"] = createCustomControlLibraryOptions.AccountID
	}
	if createCustomControlLibraryOptions.ControlLibraryName != nil {
		body["control_library_name"] = createCustomControlLibraryOptions.ControlLibraryName
	}
	if createCustomControlLibraryOptions.ControlLibraryDescription != nil {
		body["control_library_description"] = createCustomControlLibraryOptions.ControlLibraryDescription
	}
	if createCustomControlLibraryOptions.ControlLibraryType != nil {
		body["control_library_type"] = createCustomControlLibraryOptions.ControlLibraryType
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
	if createCustomControlLibraryOptions.Controls != nil {
		body["controls"] = createCustomControlLibraryOptions.Controls
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryRequest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListControlLibraries : Get all control libraries
// Retrieve all the control libraries, including predefined and custom libraries.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) (result *GetAllControlLibrariesRespBody, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ListControlLibrariesWithContext(context.Background(), listControlLibrariesOptions)
}

// ListControlLibrariesWithContext is an alternate form of the ListControlLibraries method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ListControlLibrariesWithContext(ctx context.Context, listControlLibrariesOptions *ListControlLibrariesOptions) (result *GetAllControlLibrariesRespBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listControlLibrariesOptions, "listControlLibrariesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listControlLibrariesOptions, "listControlLibrariesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listControlLibrariesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/control_libraries`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listControlLibrariesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ListControlLibraries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listControlLibrariesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listControlLibrariesOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAllControlLibrariesRespBody)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceCustomControlLibrary : Update custom control library
// Update a custom control library by specifying the control library ID.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.ReplaceCustomControlLibraryWithContext(context.Background(), replaceCustomControlLibraryOptions)
}

// ReplaceCustomControlLibraryWithContext is an alternate form of the ReplaceCustomControlLibrary method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) ReplaceCustomControlLibraryWithContext(ctx context.Context, replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
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
		"instance_id":          *replaceCustomControlLibraryOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceCustomControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "ReplaceCustomControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceCustomControlLibraryOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceCustomControlLibraryOptions.TransactionID))
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
	if replaceCustomControlLibraryOptions.Latest != nil {
		body["latest"] = replaceCustomControlLibraryOptions.Latest
	}
	if replaceCustomControlLibraryOptions.ControlsCount != nil {
		body["controls_count"] = replaceCustomControlLibraryOptions.ControlsCount
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryRequest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetControlLibrary : Get control library by id
// Get control library by id.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.GetControlLibraryWithContext(context.Background(), getControlLibraryOptions)
}

// GetControlLibraryWithContext is an alternate form of the GetControlLibrary method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) GetControlLibraryWithContext(ctx context.Context, getControlLibraryOptions *GetControlLibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
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
		"instance_id":          *getControlLibraryOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getControlLibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "GetControlLibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getControlLibraryOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getControlLibraryOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryRequest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomControllibrary : Delete custom control library
// Delete custom control library.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteCustomControllibrary(deleteCustomControllibraryOptions *DeleteCustomControllibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.DeleteCustomControllibraryWithContext(context.Background(), deleteCustomControllibraryOptions)
}

// DeleteCustomControllibraryWithContext is an alternate form of the DeleteCustomControllibrary method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) DeleteCustomControllibraryWithContext(ctx context.Context, deleteCustomControllibraryOptions *DeleteCustomControllibraryOptions) (result *ControlLibraryRequest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomControllibraryOptions, "deleteCustomControllibraryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomControllibraryOptions, "deleteCustomControllibraryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"control_libraries_id": *deleteCustomControllibraryOptions.ControlLibrariesID,
		"instance_id":          *deleteCustomControllibraryOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/control_libraries/{control_libraries_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomControllibraryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "DeleteCustomControllibrary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteCustomControllibraryOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteCustomControllibraryOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalControlLibraryRequest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateScan : Create a scan
// Create a scan.
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateScan(createScanOptions *CreateScanOptions) (result *CreateScanResponse, response *core.DetailedResponse, err error) {
	return sccPhoenixComplianceApis.CreateScanWithContext(context.Background(), createScanOptions)
}

// CreateScanWithContext is an alternate form of the CreateScan method which supports a Context parameter
func (sccPhoenixComplianceApis *SccPhoenixComplianceApisV1) CreateScanWithContext(ctx context.Context, createScanOptions *CreateScanOptions) (result *CreateScanResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createScanOptions, "createScanOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createScanOptions, "createScanOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createScanOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = sccPhoenixComplianceApis.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(sccPhoenixComplianceApis.Service.Options.URL, `/instances/{instance_id}/v3/scans`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createScanOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("scc_phoenix_compliance_apis", "V1", "CreateScan")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createScanOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createScanOptions.TransactionID))
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
	response, err = sccPhoenixComplianceApis.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateScanResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddProfileOptions : The AddProfile options.
type AddProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Name of the Profile.
	ProfileName *string `json:"profile_name,omitempty"`

	// Description of the profile.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// Type of the profile.
	ProfileType *string `json:"profile_type,omitempty"`

	// Version of the profile.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// If Latest is enabled or not.
	Latest *bool `json:"latest,omitempty"`

	// The version group label of the profile.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Controls in the profile.
	Controls []ProfileControlsInRequest `json:"controls,omitempty"`

	// default parameters of the profile.
	DefaultParameters []DefaultParameters `json:"default_parameters,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the AddProfileOptions.ProfileType property.
// Type of the profile.
const (
	AddProfileOptions_ProfileType_Custom     = "custom"
	AddProfileOptions_ProfileType_Predefined = "predefined"
)

// NewAddProfileOptions : Instantiate AddProfileOptions
func (*SccPhoenixComplianceApisV1) NewAddProfileOptions(profilesID string, instanceID string) *AddProfileOptions {
	return &AddProfileOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *AddProfileOptions) SetProfilesID(profilesID string) *AddProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *AddProfileOptions) SetInstanceID(instanceID string) *AddProfileOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetProfileName : Allow user to set ProfileName
func (_options *AddProfileOptions) SetProfileName(profileName string) *AddProfileOptions {
	_options.ProfileName = core.StringPtr(profileName)
	return _options
}

// SetProfileDescription : Allow user to set ProfileDescription
func (_options *AddProfileOptions) SetProfileDescription(profileDescription string) *AddProfileOptions {
	_options.ProfileDescription = core.StringPtr(profileDescription)
	return _options
}

// SetProfileType : Allow user to set ProfileType
func (_options *AddProfileOptions) SetProfileType(profileType string) *AddProfileOptions {
	_options.ProfileType = core.StringPtr(profileType)
	return _options
}

// SetProfileVersion : Allow user to set ProfileVersion
func (_options *AddProfileOptions) SetProfileVersion(profileVersion string) *AddProfileOptions {
	_options.ProfileVersion = core.StringPtr(profileVersion)
	return _options
}

// SetLatest : Allow user to set Latest
func (_options *AddProfileOptions) SetLatest(latest bool) *AddProfileOptions {
	_options.Latest = core.BoolPtr(latest)
	return _options
}

// SetVersionGroupLabel : Allow user to set VersionGroupLabel
func (_options *AddProfileOptions) SetVersionGroupLabel(versionGroupLabel string) *AddProfileOptions {
	_options.VersionGroupLabel = core.StringPtr(versionGroupLabel)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *AddProfileOptions) SetControls(controls []ProfileControlsInRequest) *AddProfileOptions {
	_options.Controls = controls
	return _options
}

// SetDefaultParameters : Allow user to set DefaultParameters
func (_options *AddProfileOptions) SetDefaultParameters(defaultParameters []DefaultParameters) *AddProfileOptions {
	_options.DefaultParameters = defaultParameters
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddProfileOptions) SetTransactionID(transactionID string) *AddProfileOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddProfileOptions) SetHeaders(param map[string]string) *AddProfileOptions {
	options.Headers = param
	return options
}

// AttachmentPayload : The attachment details of a profile.
type AttachmentPayload struct {
	// attachment id.
	ID *string `json:"id,omitempty"`

	// account id.
	AccountID *string `json:"account_id,omitempty"`

	// scope payload.
	IncludedScope *ScopePayload `json:"included_scope,omitempty"`

	// exclusions.
	Exclusions []ScopePayload `json:"exclusions,omitempty"`

	// created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// updated by.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// updated on.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// status.
	Status *string `json:"status,omitempty"`

	// attachment parameters.
	AttachmentParameters []ParameterDetails `json:"attachment_parameters,omitempty"`

	// payload of the attachments notifications.
	AttachmentNotifications *AttachmentsNotificationsPayload `json:"attachment_notifications,omitempty"`
}

// Constants associated with the AttachmentPayload.Status property.
// status.
const (
	AttachmentPayload_Status_Disabled = "disabled"
	AttachmentPayload_Status_Enabled  = "enabled"
)

// UnmarshalAttachmentPayload unmarshals an instance of AttachmentPayload from the specified map of raw messages.
func UnmarshalAttachmentPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentPayload)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalScopePayload)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "exclusions", &obj.Exclusions, UnmarshalScopePayload)
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
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment_parameters", &obj.AttachmentParameters, UnmarshalParameterDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment_notifications", &obj.AttachmentNotifications, UnmarshalAttachmentsNotificationsPayload)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentProfileRequest : request body of attachments of a profile.
type AttachmentProfileRequest struct {
	// the attachments of a profile.
	Attachments []AttachmentPayload `json:"attachments,omitempty"`
}

// UnmarshalAttachmentProfileRequest unmarshals an instance of AttachmentProfileRequest from the specified map of raw messages.
func UnmarshalAttachmentProfileRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentProfileRequest)
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachmentPayload)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentProfileResponse : Response body for attachment profile.
type AttachmentProfileResponse struct {
	// Profile id.
	ProfileID *string `json:"profile_id,omitempty"`

	// List of attachments.
	Attachments []AttachmentResponse `json:"attachments,omitempty"`
}

// UnmarshalAttachmentProfileResponse unmarshals an instance of AttachmentProfileResponse from the specified map of raw messages.
func UnmarshalAttachmentProfileResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentProfileResponse)
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachmentResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentResponse : attachment details for profile.
type AttachmentResponse struct {
	// Attachment id.
	ID *string `json:"id,omitempty"`

	// Account id.
	AccountID *string `json:"account_id,omitempty"`

	// scope payload.
	IncludedScope *ScopePayload `json:"included_scope,omitempty"`

	// Excluded scopes.
	Exclusions []ScopePayload `json:"exclusions,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// Updated by.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// Updated on.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// Status.
	Status *string `json:"status,omitempty"`

	// Attachment parameters.
	AttachmentParameters []ParameterDetails `json:"attachment_parameters,omitempty"`

	// Last scan id.
	LastScan *string `json:"last_scan,omitempty"`

	// Last scan status.
	LastScanStatus *string `json:"last_scan_status,omitempty"`

	// Last scan time.
	LastScanTime *string `json:"last_scan_time,omitempty"`
}

// UnmarshalAttachmentResponse unmarshals an instance of AttachmentResponse from the specified map of raw messages.
func UnmarshalAttachmentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalScopePayload)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "exclusions", &obj.Exclusions, UnmarshalScopePayload)
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
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachment_parameters", &obj.AttachmentParameters, UnmarshalParameterDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan", &obj.LastScan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_status", &obj.LastScanStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_time", &obj.LastScanTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CheckProfileAttachmnetsOptions : The CheckProfileAttachmnets options.
type CheckProfileAttachmnetsOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckProfileAttachmnetsOptions : Instantiate CheckProfileAttachmnetsOptions
func (*SccPhoenixComplianceApisV1) NewCheckProfileAttachmnetsOptions(profilesID string, instanceID string) *CheckProfileAttachmnetsOptions {
	return &CheckProfileAttachmnetsOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *CheckProfileAttachmnetsOptions) SetProfilesID(profilesID string) *CheckProfileAttachmnetsOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CheckProfileAttachmnetsOptions) SetInstanceID(instanceID string) *CheckProfileAttachmnetsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CheckProfileAttachmnetsOptions) SetTransactionID(transactionID string) *CheckProfileAttachmnetsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CheckProfileAttachmnetsOptions) SetHeaders(param map[string]string) *CheckProfileAttachmnetsOptions {
	options.Headers = param
	return options
}

// ControlDocs : Control Docs.
type ControlDocs struct {
	// ID of Control Docs.
	ControlDocsID *string `json:"control_docs_id,omitempty"`

	// Type of Control Docs.
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

// ControlLibraryListResponse : ControlLibraryListResponse struct
type ControlLibraryListResponse struct {
	// The ID of the control library.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The Control Library Name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// Control Library Description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// Control Library Type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// Created On.
	CreatedOn *string `json:"created_on,omitempty"`

	// Created By.
	CreatedBy *string `json:"created_by,omitempty"`

	// Updated ON.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// Updated By.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// Version Group Label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Control Library Version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// Latest.
	Latest *bool `json:"latest,omitempty"`

	// Number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`
}

// UnmarshalControlLibraryListResponse unmarshals an instance of ControlLibraryListResponse from the specified map of raw messages.
func UnmarshalControlLibraryListResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibraryListResponse)
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

// ControlLibraryRequest : Request payload of the Control Library.
type ControlLibraryRequest struct {
	// Control Library ID.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Control Library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// Control Library Description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// Control Library Type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// Version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Control Library Version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// Latest.
	Latest *bool `json:"latest,omitempty"`

	// Number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// Controls.
	Controls []ControlsInControlLibRequestPayload `json:"controls,omitempty"`
}

// Constants associated with the ControlLibraryRequest.ControlLibraryType property.
// Control Library Type.
const (
	ControlLibraryRequest_ControlLibraryType_Custom     = "custom"
	ControlLibraryRequest_ControlLibraryType_Predefined = "predefined"
)

// UnmarshalControlLibraryRequest unmarshals an instance of ControlLibraryRequest from the specified map of raw messages.
func UnmarshalControlLibraryRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlLibraryRequest)
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
	err = core.UnmarshalPrimitive(m, "latest", &obj.Latest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "controls_count", &obj.ControlsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalControlsInControlLibRequestPayload)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlSpecifications : The control specifications for a control library.
type ControlSpecifications struct {
	// Control Specification ID.
	ID *string `json:"id,omitempty"`

	// Responsibility.
	Responsibility *string `json:"responsibility,omitempty"`

	// Component ID.
	ComponentID *string `json:"component_id,omitempty"`

	// Environment of control specifications.
	Environment *string `json:"environment,omitempty"`

	// Description of control specifications.
	Description *string `json:"description,omitempty"`

	// Number of Assessment.
	AssessmentsCount *int64 `json:"assessments_count,omitempty"`

	// Assessments.
	Assessments []ImplementationPayload `json:"assessments,omitempty"`
}

// Constants associated with the ControlSpecifications.Responsibility property.
// Responsibility.
const (
	ControlSpecifications_Responsibility_User = "user"
)

// UnmarshalControlSpecifications unmarshals an instance of ControlSpecifications from the specified map of raw messages.
func UnmarshalControlSpecifications(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlSpecifications)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessments_count", &obj.AssessmentsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "assessments", &obj.Assessments, UnmarshalImplementationPayload)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ControlsInControlLibRequestPayload : The control details of a control library.
type ControlsInControlLibRequestPayload struct {
	// The ID of the control library that contains the profile.
	ControlName *string `json:"control_name,omitempty"`

	// The control name.
	ControlID *string `json:"control_id,omitempty"`

	// The control description.
	ControlDescription *string `json:"control_description,omitempty"`

	// control category.
	ControlCategory *string `json:"control_category,omitempty"`

	// control parent.
	ControlParent *string `json:"control_parent,omitempty"`

	// Control severity.
	ControlSeverity *string `json:"control_severity,omitempty"`

	// Control Tags.
	ControlTags []string `json:"control_tags,omitempty"`

	// control specifications.
	ControlSpecifications []ControlSpecifications `json:"control_specifications,omitempty"`

	// Control Docs.
	ControlDocs *ControlDocs `json:"control_docs,omitempty"`

	// Status.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the ControlsInControlLibRequestPayload.Status property.
// Status.
const (
	ControlsInControlLibRequestPayload_Status_Disabled = "disabled"
	ControlsInControlLibRequestPayload_Status_Enabled  = "enabled"
)

// UnmarshalControlsInControlLibRequestPayload unmarshals an instance of ControlsInControlLibRequestPayload from the specified map of raw messages.
func UnmarshalControlsInControlLibRequestPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ControlsInControlLibRequestPayload)
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
	err = core.UnmarshalPrimitive(m, "control_severity", &obj.ControlSeverity)
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

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// the attachments of a profile.
	Attachments []AttachmentPayload `json:"attachments,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAttachmentOptions : Instantiate CreateAttachmentOptions
func (*SccPhoenixComplianceApisV1) NewCreateAttachmentOptions(profilesID string, instanceID string) *CreateAttachmentOptions {
	return &CreateAttachmentOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *CreateAttachmentOptions) SetProfilesID(profilesID string) *CreateAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateAttachmentOptions) SetInstanceID(instanceID string) *CreateAttachmentOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetAttachments : Allow user to set Attachments
func (_options *CreateAttachmentOptions) SetAttachments(attachments []AttachmentPayload) *CreateAttachmentOptions {
	_options.Attachments = attachments
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateAttachmentOptions) SetTransactionID(transactionID string) *CreateAttachmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAttachmentOptions) SetHeaders(param map[string]string) *CreateAttachmentOptions {
	options.Headers = param
	return options
}

// CreateCustomControlLibraryOptions : The CreateCustomControlLibrary options.
type CreateCustomControlLibraryOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Control Library ID.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Control Library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// Control Library Description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// Control Library Type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// Version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Control Library Version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// Latest.
	Latest *bool `json:"latest,omitempty"`

	// Number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// Controls.
	Controls []ControlsInControlLibRequestPayload `json:"controls,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCustomControlLibraryOptions.ControlLibraryType property.
// Control Library Type.
const (
	CreateCustomControlLibraryOptions_ControlLibraryType_Custom     = "custom"
	CreateCustomControlLibraryOptions_ControlLibraryType_Predefined = "predefined"
)

// NewCreateCustomControlLibraryOptions : Instantiate CreateCustomControlLibraryOptions
func (*SccPhoenixComplianceApisV1) NewCreateCustomControlLibraryOptions(instanceID string) *CreateCustomControlLibraryOptions {
	return &CreateCustomControlLibraryOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateCustomControlLibraryOptions) SetInstanceID(instanceID string) *CreateCustomControlLibraryOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateCustomControlLibraryOptions) SetID(id string) *CreateCustomControlLibraryOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *CreateCustomControlLibraryOptions) SetAccountID(accountID string) *CreateCustomControlLibraryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
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

// SetControls : Allow user to set Controls
func (_options *CreateCustomControlLibraryOptions) SetControls(controls []ControlsInControlLibRequestPayload) *CreateCustomControlLibraryOptions {
	_options.Controls = controls
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateCustomControlLibraryOptions) SetTransactionID(transactionID string) *CreateCustomControlLibraryOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomControlLibraryOptions) SetHeaders(param map[string]string) *CreateCustomControlLibraryOptions {
	options.Headers = param
	return options
}

// CreateProfileOptions : The CreateProfile options.
type CreateProfileOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Name of the Profile.
	ProfileName *string `json:"profile_name,omitempty"`

	// Description of the profile.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// Type of the profile.
	ProfileType *string `json:"profile_type,omitempty"`

	// Version of the profile.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// If Latest is enabled or not.
	Latest *bool `json:"latest,omitempty"`

	// The version group label of the profile.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Controls in the profile.
	Controls []ProfileControlsInRequest `json:"controls,omitempty"`

	// default parameters of the profile.
	DefaultParameters []DefaultParameters `json:"default_parameters,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateProfileOptions.ProfileType property.
// Type of the profile.
const (
	CreateProfileOptions_ProfileType_Custom     = "custom"
	CreateProfileOptions_ProfileType_Predefined = "predefined"
)

// NewCreateProfileOptions : Instantiate CreateProfileOptions
func (*SccPhoenixComplianceApisV1) NewCreateProfileOptions(instanceID string) *CreateProfileOptions {
	return &CreateProfileOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateProfileOptions) SetInstanceID(instanceID string) *CreateProfileOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
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

// SetProfileVersion : Allow user to set ProfileVersion
func (_options *CreateProfileOptions) SetProfileVersion(profileVersion string) *CreateProfileOptions {
	_options.ProfileVersion = core.StringPtr(profileVersion)
	return _options
}

// SetLatest : Allow user to set Latest
func (_options *CreateProfileOptions) SetLatest(latest bool) *CreateProfileOptions {
	_options.Latest = core.BoolPtr(latest)
	return _options
}

// SetVersionGroupLabel : Allow user to set VersionGroupLabel
func (_options *CreateProfileOptions) SetVersionGroupLabel(versionGroupLabel string) *CreateProfileOptions {
	_options.VersionGroupLabel = core.StringPtr(versionGroupLabel)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *CreateProfileOptions) SetControls(controls []ProfileControlsInRequest) *CreateProfileOptions {
	_options.Controls = controls
	return _options
}

// SetDefaultParameters : Allow user to set DefaultParameters
func (_options *CreateProfileOptions) SetDefaultParameters(defaultParameters []DefaultParameters) *CreateProfileOptions {
	_options.DefaultParameters = defaultParameters
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateProfileOptions) SetTransactionID(transactionID string) *CreateProfileOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProfileOptions) SetHeaders(param map[string]string) *CreateProfileOptions {
	options.Headers = param
	return options
}

// CreateScanOptions : The CreateScan options.
type CreateScanOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Attachment ID.
	AttachmentID *string `json:"attachment_id,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateScanOptions : Instantiate CreateScanOptions
func (*SccPhoenixComplianceApisV1) NewCreateScanOptions(instanceID string) *CreateScanOptions {
	return &CreateScanOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateScanOptions) SetInstanceID(instanceID string) *CreateScanOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *CreateScanOptions) SetAttachmentID(attachmentID string) *CreateScanOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateScanOptions) SetTransactionID(transactionID string) *CreateScanOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateScanOptions) SetHeaders(param map[string]string) *CreateScanOptions {
	options.Headers = param
	return options
}

// CreateScanResponse : Response schema for creating a scan.
type CreateScanResponse struct {
	// Scan ID.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Attachment ID.
	AttachmentID *string `json:"attachment_id,omitempty"`

	// Report ID.
	ReportID *string `json:"report_id,omitempty"`

	// Status.
	Status *string `json:"status,omitempty"`

	// Last Scan Time.
	LastScanTime *string `json:"last_scan_time,omitempty"`

	// Next Scan Time.
	NextScanTime *string `json:"next_scan_time,omitempty"`

	// Type of Scan.
	ScanType *string `json:"scan_type,omitempty"`

	// Occurance of Scan.
	Occurence *int64 `json:"occurence,omitempty"`
}

// UnmarshalCreateScanResponse unmarshals an instance of CreateScanResponse from the specified map of raw messages.
func UnmarshalCreateScanResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateScanResponse)
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

// DefaultParameters : The control details of a profile.
type DefaultParameters struct {
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

// Constants associated with the DefaultParameters.ParameterType property.
// The parameter type.
const (
	DefaultParameters_ParameterType_Numeric    = "numeric"
	DefaultParameters_ParameterType_StringList = "string_list"
)

// UnmarshalDefaultParameters unmarshals an instance of DefaultParameters from the specified map of raw messages.
func UnmarshalDefaultParameters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DefaultParameters)
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

// DeleteCustomControllibraryOptions : The DeleteCustomControllibrary options.
type DeleteCustomControllibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomControllibraryOptions : Instantiate DeleteCustomControllibraryOptions
func (*SccPhoenixComplianceApisV1) NewDeleteCustomControllibraryOptions(controlLibrariesID string, instanceID string) *DeleteCustomControllibraryOptions {
	return &DeleteCustomControllibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
		InstanceID:         core.StringPtr(instanceID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *DeleteCustomControllibraryOptions) SetControlLibrariesID(controlLibrariesID string) *DeleteCustomControllibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteCustomControllibraryOptions) SetInstanceID(instanceID string) *DeleteCustomControllibraryOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteCustomControllibraryOptions) SetTransactionID(transactionID string) *DeleteCustomControllibraryOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomControllibraryOptions) SetHeaders(param map[string]string) *DeleteCustomControllibraryOptions {
	options.Headers = param
	return options
}

// DeleteCustomProfileOptions : The DeleteCustomProfile options.
type DeleteCustomProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomProfileOptions : Instantiate DeleteCustomProfileOptions
func (*SccPhoenixComplianceApisV1) NewDeleteCustomProfileOptions(profilesID string, instanceID string) *DeleteCustomProfileOptions {
	return &DeleteCustomProfileOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *DeleteCustomProfileOptions) SetProfilesID(profilesID string) *DeleteCustomProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteCustomProfileOptions) SetInstanceID(instanceID string) *DeleteCustomProfileOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteCustomProfileOptions) SetTransactionID(transactionID string) *DeleteCustomProfileOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomProfileOptions) SetHeaders(param map[string]string) *DeleteCustomProfileOptions {
	options.Headers = param
	return options
}

// DeleteProfileAttachmnetOptions : The DeleteProfileAttachmnet options.
type DeleteProfileAttachmnetOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProfileAttachmnetOptions : Instantiate DeleteProfileAttachmnetOptions
func (*SccPhoenixComplianceApisV1) NewDeleteProfileAttachmnetOptions(profilesID string, attachmentID string, instanceID string) *DeleteProfileAttachmnetOptions {
	return &DeleteProfileAttachmnetOptions{
		ProfilesID:   core.StringPtr(profilesID),
		AttachmentID: core.StringPtr(attachmentID),
		InstanceID:   core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *DeleteProfileAttachmnetOptions) SetProfilesID(profilesID string) *DeleteProfileAttachmnetOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *DeleteProfileAttachmnetOptions) SetAttachmentID(attachmentID string) *DeleteProfileAttachmnetOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteProfileAttachmnetOptions) SetInstanceID(instanceID string) *DeleteProfileAttachmnetOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteProfileAttachmnetOptions) SetTransactionID(transactionID string) *DeleteProfileAttachmnetOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProfileAttachmnetOptions) SetHeaders(param map[string]string) *DeleteProfileAttachmnetOptions {
	options.Headers = param
	return options
}

// GetControlLibraryOptions : The GetControlLibrary options.
type GetControlLibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetControlLibraryOptions : Instantiate GetControlLibraryOptions
func (*SccPhoenixComplianceApisV1) NewGetControlLibraryOptions(controlLibrariesID string, instanceID string) *GetControlLibraryOptions {
	return &GetControlLibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
		InstanceID:         core.StringPtr(instanceID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *GetControlLibraryOptions) SetControlLibrariesID(controlLibrariesID string) *GetControlLibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetControlLibraryOptions) SetInstanceID(instanceID string) *GetControlLibraryOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetControlLibraryOptions) SetTransactionID(transactionID string) *GetControlLibraryOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetControlLibraryOptions) SetHeaders(param map[string]string) *GetControlLibraryOptions {
	options.Headers = param
	return options
}

// GetParametersByNameOptions : The GetParametersByName options.
type GetParametersByNameOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The parameter name.
	ParameterName *string `json:"parameter_name" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetParametersByNameOptions : Instantiate GetParametersByNameOptions
func (*SccPhoenixComplianceApisV1) NewGetParametersByNameOptions(profilesID string, attachmentID string, parameterName string, instanceID string) *GetParametersByNameOptions {
	return &GetParametersByNameOptions{
		ProfilesID:    core.StringPtr(profilesID),
		AttachmentID:  core.StringPtr(attachmentID),
		ParameterName: core.StringPtr(parameterName),
		InstanceID:    core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *GetParametersByNameOptions) SetProfilesID(profilesID string) *GetParametersByNameOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *GetParametersByNameOptions) SetAttachmentID(attachmentID string) *GetParametersByNameOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetParameterName : Allow user to set ParameterName
func (_options *GetParametersByNameOptions) SetParameterName(parameterName string) *GetParametersByNameOptions {
	_options.ParameterName = core.StringPtr(parameterName)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetParametersByNameOptions) SetInstanceID(instanceID string) *GetParametersByNameOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetParametersByNameOptions) SetTransactionID(transactionID string) *GetParametersByNameOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetParametersByNameOptions) SetHeaders(param map[string]string) *GetParametersByNameOptions {
	options.Headers = param
	return options
}

// GetProfileAttachmnetOptions : The GetProfileAttachmnet options.
type GetProfileAttachmnetOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfileAttachmnetOptions : Instantiate GetProfileAttachmnetOptions
func (*SccPhoenixComplianceApisV1) NewGetProfileAttachmnetOptions(profilesID string, attachmentID string, instanceID string) *GetProfileAttachmnetOptions {
	return &GetProfileAttachmnetOptions{
		ProfilesID:   core.StringPtr(profilesID),
		AttachmentID: core.StringPtr(attachmentID),
		InstanceID:   core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *GetProfileAttachmnetOptions) SetProfilesID(profilesID string) *GetProfileAttachmnetOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *GetProfileAttachmnetOptions) SetAttachmentID(attachmentID string) *GetProfileAttachmnetOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetProfileAttachmnetOptions) SetInstanceID(instanceID string) *GetProfileAttachmnetOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetProfileAttachmnetOptions) SetTransactionID(transactionID string) *GetProfileAttachmnetOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProfileAttachmnetOptions) SetHeaders(param map[string]string) *GetProfileAttachmnetOptions {
	options.Headers = param
	return options
}

// GetProfileOptions : The GetProfile options.
type GetProfileOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfileOptions : Instantiate GetProfileOptions
func (*SccPhoenixComplianceApisV1) NewGetProfileOptions(profilesID string, instanceID string) *GetProfileOptions {
	return &GetProfileOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *GetProfileOptions) SetProfilesID(profilesID string) *GetProfileOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetProfileOptions) SetInstanceID(instanceID string) *GetProfileOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetProfileOptions) SetTransactionID(transactionID string) *GetProfileOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProfileOptions) SetHeaders(param map[string]string) *GetProfileOptions {
	options.Headers = param
	return options
}

// ImplementationPayload : The implementation details of a control library.
type ImplementationPayload struct {
	// Assessment ID.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// Method of Assessment.
	AssessmentMethod *string `json:"assessment_method,omitempty"`

	// Type of Assessment.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// Description of Assessment.
	AssessmentDescription *string `json:"assessment_description,omitempty"`

	// Parameter Count.
	ParameterCount *int64 `json:"parameter_count,omitempty"`

	// Parameters.
	Parameters []ParameterInfo `json:"parameters,omitempty"`
}

// UnmarshalImplementationPayload unmarshals an instance of ImplementationPayload from the specified map of raw messages.
func UnmarshalImplementationPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImplementationPayload)
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

// ListAttachmentParametersOptions : The ListAttachmentParameters options.
type ListAttachmentParametersOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAttachmentParametersOptions : Instantiate ListAttachmentParametersOptions
func (*SccPhoenixComplianceApisV1) NewListAttachmentParametersOptions(profilesID string, attachmentID string, instanceID string) *ListAttachmentParametersOptions {
	return &ListAttachmentParametersOptions{
		ProfilesID:   core.StringPtr(profilesID),
		AttachmentID: core.StringPtr(attachmentID),
		InstanceID:   core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ListAttachmentParametersOptions) SetProfilesID(profilesID string) *ListAttachmentParametersOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ListAttachmentParametersOptions) SetAttachmentID(attachmentID string) *ListAttachmentParametersOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListAttachmentParametersOptions) SetInstanceID(instanceID string) *ListAttachmentParametersOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListAttachmentParametersOptions) SetTransactionID(transactionID string) *ListAttachmentParametersOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAttachmentParametersOptions) SetHeaders(param map[string]string) *ListAttachmentParametersOptions {
	options.Headers = param
	return options
}

// ListControlLibrariesOptions : The ListControlLibraries options.
type ListControlLibrariesOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListControlLibrariesOptions : Instantiate ListControlLibrariesOptions
func (*SccPhoenixComplianceApisV1) NewListControlLibrariesOptions(instanceID string) *ListControlLibrariesOptions {
	return &ListControlLibrariesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListControlLibrariesOptions) SetInstanceID(instanceID string) *ListControlLibrariesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListControlLibrariesOptions) SetTransactionID(transactionID string) *ListControlLibrariesOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListControlLibrariesOptions) SetHeaders(param map[string]string) *ListControlLibrariesOptions {
	options.Headers = param
	return options
}

// ListProfilesOptions : The ListProfiles options.
type ListProfilesOptions struct {
	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProfilesOptions : Instantiate ListProfilesOptions
func (*SccPhoenixComplianceApisV1) NewListProfilesOptions(instanceID string) *ListProfilesOptions {
	return &ListProfilesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListProfilesOptions) SetInstanceID(instanceID string) *ListProfilesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListProfilesOptions) SetTransactionID(transactionID string) *ListProfilesOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProfilesOptions) SetHeaders(param map[string]string) *ListProfilesOptions {
	options.Headers = param
	return options
}

// PageRefFirst : Reference page first.
type PageRefFirst struct {
	// Reference URL.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPageRefFirst unmarshals an instance of PageRefFirst from the specified map of raw messages.
func UnmarshalPageRefFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PageRefFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PageRefNext : Reference page next.
type PageRefNext struct {
	// Reference URL.
	Href *string `json:"href,omitempty"`

	// Reference start.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPageRefNext unmarshals an instance of PageRefNext from the specified map of raw messages.
func UnmarshalPageRefNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PageRefNext)
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

// ParameterDetails : The details of the parameter.
type ParameterDetails struct {
	// The name of the parameter.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The display name of the parameter.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// the type of the parameter.
	ParameterType *string `json:"parameter_type,omitempty"`

	// The value of the parameter.
	ParameterValue *string `json:"parameter_value,omitempty"`

	// The assessment type of the parameter.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The Assessment ID of the parameter.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// Parameters.
	Parameters []ParameterInfo `json:"parameters,omitempty"`
}

// Constants associated with the ParameterDetails.ParameterType property.
// the type of the parameter.
const (
	ParameterDetails_ParameterType_Numeric    = "numeric"
	ParameterDetails_ParameterType_StringList = "string_list"
)

// UnmarshalParameterDetails unmarshals an instance of ParameterDetails from the specified map of raw messages.
func UnmarshalParameterDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ParameterDetails)
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
	err = core.UnmarshalPrimitive(m, "parameter_value", &obj.ParameterValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_type", &obj.AssessmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assessment_id", &obj.AssessmentID)
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

// ParameterInfo : The parameters details.
type ParameterInfo struct {
	// Parameter Name.
	ParameterName *string `json:"parameter_name,omitempty"`

	// Parameter Display Name.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// Parameter Type.
	ParameterType *string `json:"parameter_type,omitempty"`
}

// Constants associated with the ParameterInfo.ParameterType property.
// Parameter Type.
const (
	ParameterInfo_ParameterType_Numeric    = "numeric"
	ParameterInfo_ParameterType_StringList = "string_list"
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

// ProfileControlsInRequest : The control details of a profile.
type ProfileControlsInRequest struct {
	// The ID of the control library that contains the profile.
	ControlLibraryID *string `json:"control_library_id,omitempty"`

	// The control ID.
	ControlID *string `json:"control_id,omitempty"`
}

// UnmarshalProfileControlsInRequest unmarshals an instance of ProfileControlsInRequest from the specified map of raw messages.
func UnmarshalProfileControlsInRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileControlsInRequest)
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

// ProfileControlsInResponse : The control details for a profile.
type ProfileControlsInResponse struct {
	// The ID of the control library that contains a profile.
	ControlLibraryID *string `json:"control_library_id,omitempty"`

	// control id.
	ControlID *string `json:"control_id,omitempty"`

	// control library version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// The control name.
	ControlName *string `json:"control_name,omitempty"`

	// The control description.
	ControlDescription *string `json:"control_description,omitempty"`

	// The control severity.
	ControlSeverity *string `json:"control_severity,omitempty"`

	// The control category.
	ControlCategory *string `json:"control_category,omitempty"`

	// The control parent.
	ControlParent *string `json:"control_parent,omitempty"`

	// Control Docs.
	ControlDocs *ControlDocs `json:"control_docs,omitempty"`

	// control specifications.
	ControlSpecifications []ControlSpecifications `json:"control_specifications,omitempty"`
}

// UnmarshalProfileControlsInResponse unmarshals an instance of ProfileControlsInResponse from the specified map of raw messages.
func UnmarshalProfileControlsInResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileControlsInResponse)
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
	err = core.UnmarshalPrimitive(m, "control_severity", &obj.ControlSeverity)
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
	err = core.UnmarshalModel(m, "control_docs", &obj.ControlDocs, UnmarshalControlDocs)
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

// ProfileDefaultParametersResponse : The default parameters of a profile.
type ProfileDefaultParametersResponse struct {
	// id of parameter.
	ID *string `json:"id,omitempty"`

	// default parameters.
	DefaultParameters []DefaultParameters `json:"default_parameters,omitempty"`
}

// UnmarshalProfileDefaultParametersResponse unmarshals an instance of ProfileDefaultParametersResponse from the specified map of raw messages.
func UnmarshalProfileDefaultParametersResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileDefaultParametersResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_parameters", &obj.DefaultParameters, UnmarshalDefaultParameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileResponse : Response body of the Profile.
type ProfileResponse struct {
	// Profile ID.
	ID *string `json:"id,omitempty"`

	// Profile name.
	ProfileName *string `json:"profile_name,omitempty"`

	// Profile Description.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// Profile Type.
	ProfileType *string `json:"profile_type,omitempty"`

	// Profile Version.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// Version Group Label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Latest.
	Latest *bool `json:"latest,omitempty"`

	// Created By.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created On.
	CreatedOn *string `json:"created_on,omitempty"`

	// Updated by.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// Updated On.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// Number of Controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// Number of attachments.
	AttachmentsCount *int64 `json:"attachments_count,omitempty"`

	// Control IDs.
	Controls []ProfileControlsInResponse `json:"controls,omitempty"`

	// The default parameters of a profile.
	DefaultParameters []DefaultParameters `json:"default_parameters,omitempty"`
}

// UnmarshalProfileResponse unmarshals an instance of ProfileResponse from the specified map of raw messages.
func UnmarshalProfileResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfileResponse)
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
	err = core.UnmarshalModel(m, "controls", &obj.Controls, UnmarshalProfileControlsInResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "default_parameters", &obj.DefaultParameters, UnmarshalDefaultParameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceAttachmentOptions : The ReplaceAttachment options.
type ReplaceAttachmentOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The name of the parameter.
	ParameterName *string `json:"parameter_name,omitempty"`

	// The display name of the parameter.
	ParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// the type of the parameter.
	ParameterType *string `json:"parameter_type,omitempty"`

	// The value of the parameter.
	ParameterValue *string `json:"parameter_value,omitempty"`

	// The assessment type of the parameter.
	AssessmentType *string `json:"assessment_type,omitempty"`

	// The Assessment ID of the parameter.
	AssessmentID *string `json:"assessment_id,omitempty"`

	// Parameters.
	Parameters []ParameterInfo `json:"parameters,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceAttachmentOptions.ParameterType property.
// the type of the parameter.
const (
	ReplaceAttachmentOptions_ParameterType_Numeric    = "numeric"
	ReplaceAttachmentOptions_ParameterType_StringList = "string_list"
)

// NewReplaceAttachmentOptions : Instantiate ReplaceAttachmentOptions
func (*SccPhoenixComplianceApisV1) NewReplaceAttachmentOptions(profilesID string, attachmentID string, instanceID string) *ReplaceAttachmentOptions {
	return &ReplaceAttachmentOptions{
		ProfilesID:   core.StringPtr(profilesID),
		AttachmentID: core.StringPtr(attachmentID),
		InstanceID:   core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceAttachmentOptions) SetProfilesID(profilesID string) *ReplaceAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ReplaceAttachmentOptions) SetAttachmentID(attachmentID string) *ReplaceAttachmentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceAttachmentOptions) SetInstanceID(instanceID string) *ReplaceAttachmentOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetParameterName : Allow user to set ParameterName
func (_options *ReplaceAttachmentOptions) SetParameterName(parameterName string) *ReplaceAttachmentOptions {
	_options.ParameterName = core.StringPtr(parameterName)
	return _options
}

// SetParameterDisplayName : Allow user to set ParameterDisplayName
func (_options *ReplaceAttachmentOptions) SetParameterDisplayName(parameterDisplayName string) *ReplaceAttachmentOptions {
	_options.ParameterDisplayName = core.StringPtr(parameterDisplayName)
	return _options
}

// SetParameterType : Allow user to set ParameterType
func (_options *ReplaceAttachmentOptions) SetParameterType(parameterType string) *ReplaceAttachmentOptions {
	_options.ParameterType = core.StringPtr(parameterType)
	return _options
}

// SetParameterValue : Allow user to set ParameterValue
func (_options *ReplaceAttachmentOptions) SetParameterValue(parameterValue string) *ReplaceAttachmentOptions {
	_options.ParameterValue = core.StringPtr(parameterValue)
	return _options
}

// SetAssessmentType : Allow user to set AssessmentType
func (_options *ReplaceAttachmentOptions) SetAssessmentType(assessmentType string) *ReplaceAttachmentOptions {
	_options.AssessmentType = core.StringPtr(assessmentType)
	return _options
}

// SetAssessmentID : Allow user to set AssessmentID
func (_options *ReplaceAttachmentOptions) SetAssessmentID(assessmentID string) *ReplaceAttachmentOptions {
	_options.AssessmentID = core.StringPtr(assessmentID)
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *ReplaceAttachmentOptions) SetParameters(parameters []ParameterInfo) *ReplaceAttachmentOptions {
	_options.Parameters = parameters
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceAttachmentOptions) SetTransactionID(transactionID string) *ReplaceAttachmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceAttachmentOptions) SetHeaders(param map[string]string) *ReplaceAttachmentOptions {
	options.Headers = param
	return options
}

// ReplaceAttachmnetParametersByNameOptions : The ReplaceAttachmnetParametersByName options.
type ReplaceAttachmnetParametersByNameOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The parameter name.
	ParameterName *string `json:"-" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The name of the parameter.
	NewParameterName *string `json:"parameter_name,omitempty"`

	// The display name of the parameter.
	NewParameterDisplayName *string `json:"parameter_display_name,omitempty"`

	// the type of the parameter.
	NewParameterType *string `json:"parameter_type,omitempty"`

	// The value of the parameter.
	NewParameterValue *string `json:"parameter_value,omitempty"`

	// The assessment type of the parameter.
	NewAssessmentType *string `json:"assessment_type,omitempty"`

	// The Assessment ID of the parameter.
	NewAssessmentID *string `json:"assessment_id,omitempty"`

	// Parameters.
	NewParameters []ParameterInfo `json:"parameters,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceAttachmnetParametersByNameOptions.NewParameterType property.
// the type of the parameter.
const (
	ReplaceAttachmnetParametersByNameOptions_NewParameterType_Numeric    = "numeric"
	ReplaceAttachmnetParametersByNameOptions_NewParameterType_StringList = "string_list"
)

// NewReplaceAttachmnetParametersByNameOptions : Instantiate ReplaceAttachmnetParametersByNameOptions
func (*SccPhoenixComplianceApisV1) NewReplaceAttachmnetParametersByNameOptions(profilesID string, attachmentID string, parameterName string, instanceID string) *ReplaceAttachmnetParametersByNameOptions {
	return &ReplaceAttachmnetParametersByNameOptions{
		ProfilesID:    core.StringPtr(profilesID),
		AttachmentID:  core.StringPtr(attachmentID),
		ParameterName: core.StringPtr(parameterName),
		InstanceID:    core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceAttachmnetParametersByNameOptions) SetProfilesID(profilesID string) *ReplaceAttachmnetParametersByNameOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ReplaceAttachmnetParametersByNameOptions) SetAttachmentID(attachmentID string) *ReplaceAttachmnetParametersByNameOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetParameterName : Allow user to set ParameterName
func (_options *ReplaceAttachmnetParametersByNameOptions) SetParameterName(parameterName string) *ReplaceAttachmnetParametersByNameOptions {
	_options.ParameterName = core.StringPtr(parameterName)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceAttachmnetParametersByNameOptions) SetInstanceID(instanceID string) *ReplaceAttachmnetParametersByNameOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetNewParameterName : Allow user to set NewParameterName
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewParameterName(newParameterName string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewParameterName = core.StringPtr(newParameterName)
	return _options
}

// SetNewParameterDisplayName : Allow user to set NewParameterDisplayName
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewParameterDisplayName(newParameterDisplayName string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewParameterDisplayName = core.StringPtr(newParameterDisplayName)
	return _options
}

// SetNewParameterType : Allow user to set NewParameterType
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewParameterType(newParameterType string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewParameterType = core.StringPtr(newParameterType)
	return _options
}

// SetNewParameterValue : Allow user to set NewParameterValue
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewParameterValue(newParameterValue string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewParameterValue = core.StringPtr(newParameterValue)
	return _options
}

// SetNewAssessmentType : Allow user to set NewAssessmentType
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewAssessmentType(newAssessmentType string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewAssessmentType = core.StringPtr(newAssessmentType)
	return _options
}

// SetNewAssessmentID : Allow user to set NewAssessmentID
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewAssessmentID(newAssessmentID string) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewAssessmentID = core.StringPtr(newAssessmentID)
	return _options
}

// SetNewParameters : Allow user to set NewParameters
func (_options *ReplaceAttachmnetParametersByNameOptions) SetNewParameters(newParameters []ParameterInfo) *ReplaceAttachmnetParametersByNameOptions {
	_options.NewParameters = newParameters
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceAttachmnetParametersByNameOptions) SetTransactionID(transactionID string) *ReplaceAttachmnetParametersByNameOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceAttachmnetParametersByNameOptions) SetHeaders(param map[string]string) *ReplaceAttachmnetParametersByNameOptions {
	options.Headers = param
	return options
}

// ReplaceCustomControlLibraryOptions : The ReplaceCustomControlLibrary options.
type ReplaceCustomControlLibraryOptions struct {
	// The control library ID.
	ControlLibrariesID *string `json:"control_libraries_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Control Library ID.
	ID *string `json:"id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Control Library name.
	ControlLibraryName *string `json:"control_library_name,omitempty"`

	// Control Library Description.
	ControlLibraryDescription *string `json:"control_library_description,omitempty"`

	// Control Library Type.
	ControlLibraryType *string `json:"control_library_type,omitempty"`

	// Version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// Control Library Version.
	ControlLibraryVersion *string `json:"control_library_version,omitempty"`

	// Latest.
	Latest *bool `json:"latest,omitempty"`

	// Number of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// Controls.
	Controls []ControlsInControlLibRequestPayload `json:"controls,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceCustomControlLibraryOptions.ControlLibraryType property.
// Control Library Type.
const (
	ReplaceCustomControlLibraryOptions_ControlLibraryType_Custom     = "custom"
	ReplaceCustomControlLibraryOptions_ControlLibraryType_Predefined = "predefined"
)

// NewReplaceCustomControlLibraryOptions : Instantiate ReplaceCustomControlLibraryOptions
func (*SccPhoenixComplianceApisV1) NewReplaceCustomControlLibraryOptions(controlLibrariesID string, instanceID string) *ReplaceCustomControlLibraryOptions {
	return &ReplaceCustomControlLibraryOptions{
		ControlLibrariesID: core.StringPtr(controlLibrariesID),
		InstanceID:         core.StringPtr(instanceID),
	}
}

// SetControlLibrariesID : Allow user to set ControlLibrariesID
func (_options *ReplaceCustomControlLibraryOptions) SetControlLibrariesID(controlLibrariesID string) *ReplaceCustomControlLibraryOptions {
	_options.ControlLibrariesID = core.StringPtr(controlLibrariesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceCustomControlLibraryOptions) SetInstanceID(instanceID string) *ReplaceCustomControlLibraryOptions {
	_options.InstanceID = core.StringPtr(instanceID)
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

// SetLatest : Allow user to set Latest
func (_options *ReplaceCustomControlLibraryOptions) SetLatest(latest bool) *ReplaceCustomControlLibraryOptions {
	_options.Latest = core.BoolPtr(latest)
	return _options
}

// SetControlsCount : Allow user to set ControlsCount
func (_options *ReplaceCustomControlLibraryOptions) SetControlsCount(controlsCount int64) *ReplaceCustomControlLibraryOptions {
	_options.ControlsCount = core.Int64Ptr(controlsCount)
	return _options
}

// SetControls : Allow user to set Controls
func (_options *ReplaceCustomControlLibraryOptions) SetControls(controls []ControlsInControlLibRequestPayload) *ReplaceCustomControlLibraryOptions {
	_options.Controls = controls
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceCustomControlLibraryOptions) SetTransactionID(transactionID string) *ReplaceCustomControlLibraryOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceCustomControlLibraryOptions) SetHeaders(param map[string]string) *ReplaceCustomControlLibraryOptions {
	options.Headers = param
	return options
}

// ReplaceProfileAttachmentOptions : The ReplaceProfileAttachment options.
type ReplaceProfileAttachmentOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// The attachment ID.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// attachment id.
	ID *string `json:"id,omitempty"`

	// account id.
	AccountID *string `json:"account_id,omitempty"`

	// scope payload.
	IncludedScope *ScopePayload `json:"included_scope,omitempty"`

	// exclusions.
	Exclusions []ScopePayload `json:"exclusions,omitempty"`

	// created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// updated by.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// updated on.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// status.
	Status *string `json:"status,omitempty"`

	// attachment parameters.
	AttachmentParameters []ParameterDetails `json:"attachment_parameters,omitempty"`

	// payload of the attachments notifications.
	AttachmentNotifications *AttachmentsNotificationsPayload `json:"attachment_notifications,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceProfileAttachmentOptions.Status property.
// status.
const (
	ReplaceProfileAttachmentOptions_Status_Disabled = "disabled"
	ReplaceProfileAttachmentOptions_Status_Enabled  = "enabled"
)

// NewReplaceProfileAttachmentOptions : Instantiate ReplaceProfileAttachmentOptions
func (*SccPhoenixComplianceApisV1) NewReplaceProfileAttachmentOptions(profilesID string, attachmentID string, instanceID string) *ReplaceProfileAttachmentOptions {
	return &ReplaceProfileAttachmentOptions{
		ProfilesID:   core.StringPtr(profilesID),
		AttachmentID: core.StringPtr(attachmentID),
		InstanceID:   core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceProfileAttachmentOptions) SetProfilesID(profilesID string) *ReplaceProfileAttachmentOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ReplaceProfileAttachmentOptions) SetAttachmentID(attachmentID string) *ReplaceProfileAttachmentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceProfileAttachmentOptions) SetInstanceID(instanceID string) *ReplaceProfileAttachmentOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ReplaceProfileAttachmentOptions) SetID(id string) *ReplaceProfileAttachmentOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceProfileAttachmentOptions) SetAccountID(accountID string) *ReplaceProfileAttachmentOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIncludedScope : Allow user to set IncludedScope
func (_options *ReplaceProfileAttachmentOptions) SetIncludedScope(includedScope *ScopePayload) *ReplaceProfileAttachmentOptions {
	_options.IncludedScope = includedScope
	return _options
}

// SetExclusions : Allow user to set Exclusions
func (_options *ReplaceProfileAttachmentOptions) SetExclusions(exclusions []ScopePayload) *ReplaceProfileAttachmentOptions {
	_options.Exclusions = exclusions
	return _options
}

// SetCreatedBy : Allow user to set CreatedBy
func (_options *ReplaceProfileAttachmentOptions) SetCreatedBy(createdBy string) *ReplaceProfileAttachmentOptions {
	_options.CreatedBy = core.StringPtr(createdBy)
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *ReplaceProfileAttachmentOptions) SetCreatedOn(createdOn string) *ReplaceProfileAttachmentOptions {
	_options.CreatedOn = core.StringPtr(createdOn)
	return _options
}

// SetUpdatedBy : Allow user to set UpdatedBy
func (_options *ReplaceProfileAttachmentOptions) SetUpdatedBy(updatedBy string) *ReplaceProfileAttachmentOptions {
	_options.UpdatedBy = core.StringPtr(updatedBy)
	return _options
}

// SetUpdatedOn : Allow user to set UpdatedOn
func (_options *ReplaceProfileAttachmentOptions) SetUpdatedOn(updatedOn string) *ReplaceProfileAttachmentOptions {
	_options.UpdatedOn = core.StringPtr(updatedOn)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *ReplaceProfileAttachmentOptions) SetStatus(status string) *ReplaceProfileAttachmentOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetAttachmentParameters : Allow user to set AttachmentParameters
func (_options *ReplaceProfileAttachmentOptions) SetAttachmentParameters(attachmentParameters []ParameterDetails) *ReplaceProfileAttachmentOptions {
	_options.AttachmentParameters = attachmentParameters
	return _options
}

// SetAttachmentNotifications : Allow user to set AttachmentNotifications
func (_options *ReplaceProfileAttachmentOptions) SetAttachmentNotifications(attachmentNotifications *AttachmentsNotificationsPayload) *ReplaceProfileAttachmentOptions {
	_options.AttachmentNotifications = attachmentNotifications
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceProfileAttachmentOptions) SetTransactionID(transactionID string) *ReplaceProfileAttachmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceProfileAttachmentOptions) SetHeaders(param map[string]string) *ReplaceProfileAttachmentOptions {
	options.Headers = param
	return options
}

// ReplaceProfileParametersOptions : The ReplaceProfileParameters options.
type ReplaceProfileParametersOptions struct {
	// The profile ID.
	ProfilesID *string `json:"profiles_id" validate:"required,ne="`

	// Instance id.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// id of parameter.
	ID *string `json:"id,omitempty"`

	// default parameters.
	DefaultParameters []DefaultParameters `json:"default_parameters,omitempty"`

	// The transaction ID for the request in UUID v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceProfileParametersOptions : Instantiate ReplaceProfileParametersOptions
func (*SccPhoenixComplianceApisV1) NewReplaceProfileParametersOptions(profilesID string, instanceID string) *ReplaceProfileParametersOptions {
	return &ReplaceProfileParametersOptions{
		ProfilesID: core.StringPtr(profilesID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetProfilesID : Allow user to set ProfilesID
func (_options *ReplaceProfileParametersOptions) SetProfilesID(profilesID string) *ReplaceProfileParametersOptions {
	_options.ProfilesID = core.StringPtr(profilesID)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceProfileParametersOptions) SetInstanceID(instanceID string) *ReplaceProfileParametersOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ReplaceProfileParametersOptions) SetID(id string) *ReplaceProfileParametersOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDefaultParameters : Allow user to set DefaultParameters
func (_options *ReplaceProfileParametersOptions) SetDefaultParameters(defaultParameters []DefaultParameters) *ReplaceProfileParametersOptions {
	_options.DefaultParameters = defaultParameters
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceProfileParametersOptions) SetTransactionID(transactionID string) *ReplaceProfileParametersOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceProfileParametersOptions) SetHeaders(param map[string]string) *ReplaceProfileParametersOptions {
	options.Headers = param
	return options
}

// ScopePayload : scope payload.
type ScopePayload struct {
	// scope id.
	ScopeID *string `json:"scope_id,omitempty"`

	// Scope type.
	ScopeType *string `json:"scope_type,omitempty"`
}

// UnmarshalScopePayload unmarshals an instance of ScopePayload from the specified map of raw messages.
func UnmarshalScopePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScopePayload)
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

// AttachmentsNotificationsPayload : payload of the attachments notifications.
type AttachmentsNotificationsPayload struct {
	// enabled notifications.
	Enabled *bool `json:"enabled,omitempty"`

	// failed controls.
	Controls *FailedControls `json:"controls,omitempty"`
}

// UnmarshalAttachmentsNotificationsPayload unmarshals an instance of AttachmentsNotificationsPayload from the specified map of raw messages.
func UnmarshalAttachmentsNotificationsPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentsNotificationsPayload)
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

// FailedControls : failed controls.
type FailedControls struct {
	// threshold limit.
	ThresholdLimit *int64 `json:"threshold_limit,omitempty"`

	// failed control ids.
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

// GetAllAttachmnetsForProfileRespBody : All the attachments present in a profile.
type GetAllAttachmnetsForProfileRespBody struct {
	// Number of attachments.
	TotalCount *int64 `json:"total_count,omitempty"`

	// Limit on Attachments.
	Limit *int64 `json:"limit,omitempty"`

	// Reference page first.
	First *PageRefFirst `json:"first,omitempty"`

	// Reference page next.
	Next *PageRefNext `json:"next,omitempty"`

	// Profile ID.
	ProfileID *string `json:"profile_id,omitempty"`

	// Account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The Control Library ID.
	ControlLibraries []ControlLibraryListResponse `json:"control_libraries,omitempty"`

	// the attachments of a profile.
	Attachments []AttachmentProfileRequest `json:"attachments,omitempty"`
}

// UnmarshalGetAllAttachmnetsForProfileRespBody unmarshals an instance of GetAllAttachmnetsForProfileRespBody from the specified map of raw messages.
func UnmarshalGetAllAttachmnetsForProfileRespBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAllAttachmnetsForProfileRespBody)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageRefFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageRefNext)
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
	err = core.UnmarshalModel(m, "control_libraries", &obj.ControlLibraries, UnmarshalControlLibraryListResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachmentProfileRequest)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAllControlLibrariesRespBody : response body of control libraries.
type GetAllControlLibrariesRespBody struct {
	// number of control libraries.
	TotalCount *int64 `json:"total_count,omitempty"`

	// limit.
	Limit *int64 `json:"limit,omitempty"`

	// Reference page first.
	First *PageRefFirst `json:"first,omitempty"`

	// Reference page next.
	Next *PageRefNext `json:"next,omitempty"`

	// control libraries.
	ControlLibraries []ControlLibraryListResponse `json:"control_libraries,omitempty"`
}

// UnmarshalGetAllControlLibrariesRespBody unmarshals an instance of GetAllControlLibrariesRespBody from the specified map of raw messages.
func UnmarshalGetAllControlLibrariesRespBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAllControlLibrariesRespBody)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageRefFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageRefNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "control_libraries", &obj.ControlLibraries, UnmarshalControlLibraryListResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAllProfilesRespBody : Response body of get All profiles.
type GetAllProfilesRespBody struct {
	// Number of profiles.
	TotalCount *int64 `json:"total_count,omitempty"`

	// limit.
	Limit *int64 `json:"limit,omitempty"`

	// Reference page first.
	First *PageRefFirst `json:"first,omitempty"`

	// Reference page next.
	Next *PageRefNext `json:"next,omitempty"`

	// Profiles.
	Profiles []ListProfilesResponseStructure `json:"profiles,omitempty"`
}

// UnmarshalGetAllProfilesRespBody unmarshals an instance of GetAllProfilesRespBody from the specified map of raw messages.
func UnmarshalGetAllProfilesRespBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAllProfilesRespBody)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPageRefFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPageRefNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "profiles", &obj.Profiles, UnmarshalListProfilesResponseStructure)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListProfilesResponseStructure : ListProfilesResponseStructure struct
type ListProfilesResponseStructure struct {
	// id of the profile.
	ID *string `json:"id,omitempty"`

	// name of the profile.
	ProfileName *string `json:"profile_name,omitempty"`

	// description of the profile.
	ProfileDescription *string `json:"profile_description,omitempty"`

	// type of the profile.
	ProfileType *string `json:"profile_type,omitempty"`

	// version of the profile.
	ProfileVersion *string `json:"profile_version,omitempty"`

	// version group label.
	VersionGroupLabel *string `json:"version_group_label,omitempty"`

	// latest.
	Latest *bool `json:"latest,omitempty"`

	// created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// updated by.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// updated on.
	UpdatedOn *string `json:"updated_on,omitempty"`

	// No of controls.
	ControlsCount *int64 `json:"controls_count,omitempty"`

	// No of attachments.
	AttachmentsCount *int64 `json:"attachments_count,omitempty"`
}

// UnmarshalListProfilesResponseStructure unmarshals an instance of ListProfilesResponseStructure from the specified map of raw messages.
func UnmarshalListProfilesResponseStructure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListProfilesResponseStructure)
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

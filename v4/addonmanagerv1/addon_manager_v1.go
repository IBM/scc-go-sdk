/**
 * (C) Copyright IBM Corp. 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.46.0-a4e29da0-20220224-210428
 */

// Package addonmanagerv1 : Operations and models for the AddonManagerV1 service
package addonmanagerv1

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

// AddonManagerV1 : IBM Cloud Security and Compliance Center Addon Manager API
//
// API Version: 1.0.0
type AddonManagerV1 struct {
	Service *core.BaseService

	// Account ID.
	AccountID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.secadvisor.cloud.ibm.com/addonmgr"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "addon_manager"

// AddonManagerV1Options : Service options
type AddonManagerV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Account ID.
	AccountID *string `validate:"required"`
}

// NewAddonManagerV1UsingExternalConfig : constructs an instance of AddonManagerV1 with passed in options and external configuration.
func NewAddonManagerV1UsingExternalConfig(options *AddonManagerV1Options) (addonManager *AddonManagerV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	addonManager, err = NewAddonManagerV1(options)
	if err != nil {
		return
	}

	err = addonManager.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = addonManager.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAddonManagerV1 : constructs an instance of AddonManagerV1 with passed in options.
func NewAddonManagerV1(options *AddonManagerV1Options) (service *AddonManagerV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
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

	service = &AddonManagerV1{
		Service:   baseService,
		AccountID: options.AccountID,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://us-south.secadvisor.cloud.ibm.com/addonmgr",
		"us-east":  "https://us-south.secadvisor.cloud.ibm.com/addonmgr",
		"eu-gb":    "https://eu-gb.secadvisor.cloud.ibm.com/addonmgr",
		"eu-de":    "https://eu.compliance.cloud.ibm.com/si/addonmgr",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "addonManager" suitable for processing requests.
func (addonManager *AddonManagerV1) Clone() *AddonManagerV1 {
	if core.IsNil(addonManager) {
		return nil
	}
	clone := *addonManager
	clone.Service = addonManager.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (addonManager *AddonManagerV1) SetServiceURL(url string) error {
	return addonManager.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (addonManager *AddonManagerV1) GetServiceURL() string {
	return addonManager.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (addonManager *AddonManagerV1) SetDefaultHeaders(headers http.Header) {
	addonManager.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (addonManager *AddonManagerV1) SetEnableGzipCompression(enableGzip bool) {
	addonManager.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (addonManager *AddonManagerV1) GetEnableGzipCompression() bool {
	return addonManager.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (addonManager *AddonManagerV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	addonManager.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (addonManager *AddonManagerV1) DisableRetries() {
	addonManager.Service.DisableRetries()
}

// GetSupportedInsightsV2 : Get list of all supported addons
// Get list of all supported addons.
func (addonManager *AddonManagerV1) GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options) (result *AllInsights, response *core.DetailedResponse, err error) {
	return addonManager.GetSupportedInsightsV2WithContext(context.Background(), getSupportedInsightsV2Options)
}

// GetSupportedInsightsV2WithContext is an alternate form of the GetSupportedInsightsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) GetSupportedInsightsV2WithContext(ctx context.Context, getSupportedInsightsV2Options *GetSupportedInsightsV2Options) (result *AllInsights, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSupportedInsightsV2Options, "getSupportedInsightsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/insights`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSupportedInsightsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "GetSupportedInsightsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAllInsights)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddNetworkInsightsCosDetailsV2 : Add new COS buckets to Network Insights
// Add new COS buckets to Network Insights.
func (addonManager *AddonManagerV1) AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options) (result *NetworkInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	return addonManager.AddNetworkInsightsCosDetailsV2WithContext(context.Background(), addNetworkInsightsCosDetailsV2Options)
}

// AddNetworkInsightsCosDetailsV2WithContext is an alternate form of the AddNetworkInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) AddNetworkInsightsCosDetailsV2WithContext(ctx context.Context, addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options) (result *NetworkInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addNetworkInsightsCosDetailsV2Options, "addNetworkInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addNetworkInsightsCosDetailsV2Options, "addNetworkInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/network_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addNetworkInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "AddNetworkInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addNetworkInsightsCosDetailsV2Options.RegionID != nil {
		body["region_id"] = addNetworkInsightsCosDetailsV2Options.RegionID
	}
	if addNetworkInsightsCosDetailsV2Options.CosDetails != nil {
		body["cos_details"] = addNetworkInsightsCosDetailsV2Options.CosDetails
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
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetworkInsightsCosDetailsOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteNetworkInsightsCosDetailsV2 : Remove COS buckets from Network Insights
// Remove COS buckets from Network Insights.
func (addonManager *AddonManagerV1) DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonManager.DeleteNetworkInsightsCosDetailsV2WithContext(context.Background(), deleteNetworkInsightsCosDetailsV2Options)
}

// DeleteNetworkInsightsCosDetailsV2WithContext is an alternate form of the DeleteNetworkInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) DeleteNetworkInsightsCosDetailsV2WithContext(ctx context.Context, deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNetworkInsightsCosDetailsV2Options, "deleteNetworkInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNetworkInsightsCosDetailsV2Options, "deleteNetworkInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/network_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNetworkInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "DeleteNetworkInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if deleteNetworkInsightsCosDetailsV2Options.Ids != nil {
		body["ids"] = deleteNetworkInsightsCosDetailsV2Options.Ids
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonManager.Service.Request(request, nil)

	return
}

// GetNetworkInsightsCosDetailsV2 : Get all COS buckets details from Network Insights
// Get all COS buckets details from Network Insights.
func (addonManager *AddonManagerV1) GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options *GetNetworkInsightsCosDetailsV2Options) (result *NetworkInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	return addonManager.GetNetworkInsightsCosDetailsV2WithContext(context.Background(), getNetworkInsightsCosDetailsV2Options)
}

// GetNetworkInsightsCosDetailsV2WithContext is an alternate form of the GetNetworkInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) GetNetworkInsightsCosDetailsV2WithContext(ctx context.Context, getNetworkInsightsCosDetailsV2Options *GetNetworkInsightsCosDetailsV2Options) (result *NetworkInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getNetworkInsightsCosDetailsV2Options, "getNetworkInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/network_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNetworkInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "GetNetworkInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetworkInsightsCosDetailsOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetNetworkInsightStatusV2 : Get Network Insights status configuration
// Get Network Insights status configuration.
func (addonManager *AddonManagerV1) GetNetworkInsightStatusV2(getNetworkInsightStatusV2Options *GetNetworkInsightStatusV2Options) (result *NetworkInsightsStatusConfigOutput, response *core.DetailedResponse, err error) {
	return addonManager.GetNetworkInsightStatusV2WithContext(context.Background(), getNetworkInsightStatusV2Options)
}

// GetNetworkInsightStatusV2WithContext is an alternate form of the GetNetworkInsightStatusV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) GetNetworkInsightStatusV2WithContext(ctx context.Context, getNetworkInsightStatusV2Options *GetNetworkInsightStatusV2Options) (result *NetworkInsightsStatusConfigOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getNetworkInsightStatusV2Options, "getNetworkInsightStatusV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/network_insights/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNetworkInsightStatusV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "GetNetworkInsightStatusV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetworkInsightsStatusConfigOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateNetworkInsightStatusV2 : Update Network Insights status configuration
// Update Network Insights status configuration.
func (addonManager *AddonManagerV1) UpdateNetworkInsightStatusV2(updateNetworkInsightStatusV2Options *UpdateNetworkInsightStatusV2Options) (response *core.DetailedResponse, err error) {
	return addonManager.UpdateNetworkInsightStatusV2WithContext(context.Background(), updateNetworkInsightStatusV2Options)
}

// UpdateNetworkInsightStatusV2WithContext is an alternate form of the UpdateNetworkInsightStatusV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) UpdateNetworkInsightStatusV2WithContext(ctx context.Context, updateNetworkInsightStatusV2Options *UpdateNetworkInsightStatusV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNetworkInsightStatusV2Options, "updateNetworkInsightStatusV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateNetworkInsightStatusV2Options, "updateNetworkInsightStatusV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/network_insights/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateNetworkInsightStatusV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "UpdateNetworkInsightStatusV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateNetworkInsightStatusV2Options.RegionID != nil {
		body["region_id"] = updateNetworkInsightStatusV2Options.RegionID
	}
	if updateNetworkInsightStatusV2Options.Status != nil {
		body["status"] = updateNetworkInsightStatusV2Options.Status
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonManager.Service.Request(request, nil)

	return
}

// AddActivityInsightsCosDetailsV2 : Add new COS buckets to Activity Insights
// Add new COS buckets to Activity Insights.
func (addonManager *AddonManagerV1) AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options) (result *ActivityInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	return addonManager.AddActivityInsightsCosDetailsV2WithContext(context.Background(), addActivityInsightsCosDetailsV2Options)
}

// AddActivityInsightsCosDetailsV2WithContext is an alternate form of the AddActivityInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) AddActivityInsightsCosDetailsV2WithContext(ctx context.Context, addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options) (result *ActivityInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addActivityInsightsCosDetailsV2Options, "addActivityInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addActivityInsightsCosDetailsV2Options, "addActivityInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/activity_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addActivityInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "AddActivityInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addActivityInsightsCosDetailsV2Options.RegionID != nil {
		body["region_id"] = addActivityInsightsCosDetailsV2Options.RegionID
	}
	if addActivityInsightsCosDetailsV2Options.CosDetails != nil {
		body["cos_details"] = addActivityInsightsCosDetailsV2Options.CosDetails
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
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalActivityInsightsCosDetailsOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteActivityInsightsCosDetailsV2 : Remove COS buckets from Activity Insights
// Remove COS buckets from Activity Insights.
func (addonManager *AddonManagerV1) DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonManager.DeleteActivityInsightsCosDetailsV2WithContext(context.Background(), deleteActivityInsightsCosDetailsV2Options)
}

// DeleteActivityInsightsCosDetailsV2WithContext is an alternate form of the DeleteActivityInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) DeleteActivityInsightsCosDetailsV2WithContext(ctx context.Context, deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteActivityInsightsCosDetailsV2Options, "deleteActivityInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteActivityInsightsCosDetailsV2Options, "deleteActivityInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/activity_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteActivityInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "DeleteActivityInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if deleteActivityInsightsCosDetailsV2Options.Ids != nil {
		body["ids"] = deleteActivityInsightsCosDetailsV2Options.Ids
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonManager.Service.Request(request, nil)

	return
}

// GetActivityInsightsCosDetailsV2 : Get all COS buckets details from Activity Insights
// Get all COS buckets details from Activity Insights.
func (addonManager *AddonManagerV1) GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options *GetActivityInsightsCosDetailsV2Options) (result *ActivityInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	return addonManager.GetActivityInsightsCosDetailsV2WithContext(context.Background(), getActivityInsightsCosDetailsV2Options)
}

// GetActivityInsightsCosDetailsV2WithContext is an alternate form of the GetActivityInsightsCosDetailsV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) GetActivityInsightsCosDetailsV2WithContext(ctx context.Context, getActivityInsightsCosDetailsV2Options *GetActivityInsightsCosDetailsV2Options) (result *ActivityInsightsCosDetailsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getActivityInsightsCosDetailsV2Options, "getActivityInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/activity_insights/buckets`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getActivityInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "GetActivityInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalActivityInsightsCosDetailsOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetActivityInsightStatusV2 : Get Activity Insights status configuration
// Get Activity Insights status configuration.
func (addonManager *AddonManagerV1) GetActivityInsightStatusV2(getActivityInsightStatusV2Options *GetActivityInsightStatusV2Options) (result *ActivityInsightsStatusConfigOutput, response *core.DetailedResponse, err error) {
	return addonManager.GetActivityInsightStatusV2WithContext(context.Background(), getActivityInsightStatusV2Options)
}

// GetActivityInsightStatusV2WithContext is an alternate form of the GetActivityInsightStatusV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) GetActivityInsightStatusV2WithContext(ctx context.Context, getActivityInsightStatusV2Options *GetActivityInsightStatusV2Options) (result *ActivityInsightsStatusConfigOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getActivityInsightStatusV2Options, "getActivityInsightStatusV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/activity_insights/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getActivityInsightStatusV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "GetActivityInsightStatusV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonManager.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalActivityInsightsStatusConfigOutput)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateActivityInsightStatusV2 : Update Activity Insights status configuration
// Update Activity Insights status configuration.
func (addonManager *AddonManagerV1) UpdateActivityInsightStatusV2(updateActivityInsightStatusV2Options *UpdateActivityInsightStatusV2Options) (response *core.DetailedResponse, err error) {
	return addonManager.UpdateActivityInsightStatusV2WithContext(context.Background(), updateActivityInsightStatusV2Options)
}

// UpdateActivityInsightStatusV2WithContext is an alternate form of the UpdateActivityInsightStatusV2 method which supports a Context parameter
func (addonManager *AddonManagerV1) UpdateActivityInsightStatusV2WithContext(ctx context.Context, updateActivityInsightStatusV2Options *UpdateActivityInsightStatusV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateActivityInsightStatusV2Options, "updateActivityInsightStatusV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateActivityInsightStatusV2Options, "updateActivityInsightStatusV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonManager.AccountID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonManager.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonManager.Service.Options.URL, `/v2/addons/{account_id}/activity_insights/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateActivityInsightStatusV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_manager", "V1", "UpdateActivityInsightStatusV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateActivityInsightStatusV2Options.RegionID != nil {
		body["region_id"] = updateActivityInsightStatusV2Options.RegionID
	}
	if updateActivityInsightStatusV2Options.Status != nil {
		body["status"] = updateActivityInsightStatusV2Options.Status
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonManager.Service.Request(request, nil)

	return
}

// ActivityInsightsCosDetailsOutput : List of COS buckets details added to Activity Insights.
type ActivityInsightsCosDetailsOutput struct {
	CosDetails []CosDetailsWithID `json:"cos_details" validate:"required"`
}

// UnmarshalActivityInsightsCosDetailsOutput unmarshals an instance of ActivityInsightsCosDetailsOutput from the specified map of raw messages.
func UnmarshalActivityInsightsCosDetailsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActivityInsightsCosDetailsOutput)
	err = core.UnmarshalModel(m, "cos_details", &obj.CosDetails, UnmarshalCosDetailsWithID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ActivityInsightsStatusConfigOutput : Network Insights status configuration.
type ActivityInsightsStatusConfigOutput struct {
	// The type of insight whose status is to be displayed.
	Addon *string `json:"addon" validate:"required"`

	// The status of the insight, ie., enable or disable.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the ActivityInsightsStatusConfigOutput.Status property.
// The status of the insight, ie., enable or disable.
const (
	ActivityInsightsStatusConfigOutputStatusDisableConst = "disable"
	ActivityInsightsStatusConfigOutputStatusEnableConst  = "enable"
)

// UnmarshalActivityInsightsStatusConfigOutput unmarshals an instance of ActivityInsightsStatusConfigOutput from the specified map of raw messages.
func UnmarshalActivityInsightsStatusConfigOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActivityInsightsStatusConfigOutput)
	err = core.UnmarshalPrimitive(m, "addon", &obj.Addon)
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

// AddActivityInsightsCosDetailsV2Options : The AddActivityInsightsCosDetailsV2 options.
type AddActivityInsightsCosDetailsV2Options struct {
	// Region ID, e.g., us, eu, etc.
	RegionID *string `json:"region_id" validate:"required"`

	CosDetails []CosDetails `json:"cos_details" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddActivityInsightsCosDetailsV2Options : Instantiate AddActivityInsightsCosDetailsV2Options
func (*AddonManagerV1) NewAddActivityInsightsCosDetailsV2Options(regionID string, cosDetails []CosDetails) *AddActivityInsightsCosDetailsV2Options {
	return &AddActivityInsightsCosDetailsV2Options{
		RegionID:   core.StringPtr(regionID),
		CosDetails: cosDetails,
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *AddActivityInsightsCosDetailsV2Options) SetRegionID(regionID string) *AddActivityInsightsCosDetailsV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetCosDetails : Allow user to set CosDetails
func (_options *AddActivityInsightsCosDetailsV2Options) SetCosDetails(cosDetails []CosDetails) *AddActivityInsightsCosDetailsV2Options {
	_options.CosDetails = cosDetails
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddActivityInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *AddActivityInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// AddNetworkInsightsCosDetailsV2Options : The AddNetworkInsightsCosDetailsV2 options.
type AddNetworkInsightsCosDetailsV2Options struct {
	// Region ID, e.g., us, eu, etc.
	RegionID *string `json:"region_id" validate:"required"`

	CosDetails []CosDetails `json:"cos_details" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddNetworkInsightsCosDetailsV2Options : Instantiate AddNetworkInsightsCosDetailsV2Options
func (*AddonManagerV1) NewAddNetworkInsightsCosDetailsV2Options(regionID string, cosDetails []CosDetails) *AddNetworkInsightsCosDetailsV2Options {
	return &AddNetworkInsightsCosDetailsV2Options{
		RegionID:   core.StringPtr(regionID),
		CosDetails: cosDetails,
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *AddNetworkInsightsCosDetailsV2Options) SetRegionID(regionID string) *AddNetworkInsightsCosDetailsV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetCosDetails : Allow user to set CosDetails
func (_options *AddNetworkInsightsCosDetailsV2Options) SetCosDetails(cosDetails []CosDetails) *AddNetworkInsightsCosDetailsV2Options {
	_options.CosDetails = cosDetails
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddNetworkInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *AddNetworkInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// AllInsights : Response for getting the list of all supported insights.
type AllInsights struct {
	// List of all supported insights.
	Type []string `json:"type,omitempty"`
}

// Constants associated with the AllInsights.Type property.
const (
	AllInsightsTypeActivityInsightsConst = "activity_insights"
	AllInsightsTypeNetworkInsightsConst  = "network_insights"
)

// UnmarshalAllInsights unmarshals an instance of AllInsights from the specified map of raw messages.
func UnmarshalAllInsights(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AllInsights)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CosDetails : Details of a specific COS bucket.
type CosDetails struct {
	// Insight type associated with this bucket, ie, network_insights or activity_insights.
	Type *string `json:"type,omitempty"`

	CosInstance *string `json:"cos_instance,omitempty"`

	BucketName *string `json:"bucket_name,omitempty"`

	Description *string `json:"description,omitempty"`

	CosBucketURL *string `json:"cos_bucket_url,omitempty"`
}

// Constants associated with the CosDetails.Type property.
// Insight type associated with this bucket, ie, network_insights or activity_insights.
const (
	CosDetailsTypeActivityInsightsConst = "activity_insights"
	CosDetailsTypeNetworkInsightsConst  = "network_insights"
)

// UnmarshalCosDetails unmarshals an instance of CosDetails from the specified map of raw messages.
func UnmarshalCosDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosDetails)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_instance", &obj.CosInstance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_bucket_url", &obj.CosBucketURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CosDetailsWithID : Details of a specific COS bucket added.
type CosDetailsWithID struct {
	// Insight type associated with this bucket, ie, network_insights or activity_insights.
	Type *string `json:"type,omitempty"`

	CosInstance *string `json:"cos_instance,omitempty"`

	BucketName *string `json:"bucket_name,omitempty"`

	Description *string `json:"description,omitempty"`

	CosBucketURL *string `json:"cos_bucket_url,omitempty"`

	ID *string `json:"id,omitempty"`
}

// Constants associated with the CosDetailsWithID.Type property.
// Insight type associated with this bucket, ie, network_insights or activity_insights.
const (
	CosDetailsWithIDTypeActivityInsightsConst = "activity_insights"
	CosDetailsWithIDTypeNetworkInsightsConst  = "network_insights"
)

// UnmarshalCosDetailsWithID unmarshals an instance of CosDetailsWithID from the specified map of raw messages.
func UnmarshalCosDetailsWithID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosDetailsWithID)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_instance", &obj.CosInstance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_bucket_url", &obj.CosBucketURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteActivityInsightsCosDetailsV2Options : The DeleteActivityInsightsCosDetailsV2 options.
type DeleteActivityInsightsCosDetailsV2Options struct {
	// Array of Ids of COS entries.
	Ids []string `json:"ids,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteActivityInsightsCosDetailsV2Options : Instantiate DeleteActivityInsightsCosDetailsV2Options
func (*AddonManagerV1) NewDeleteActivityInsightsCosDetailsV2Options() *DeleteActivityInsightsCosDetailsV2Options {
	return &DeleteActivityInsightsCosDetailsV2Options{}
}

// SetIds : Allow user to set Ids
func (_options *DeleteActivityInsightsCosDetailsV2Options) SetIds(ids []string) *DeleteActivityInsightsCosDetailsV2Options {
	_options.Ids = ids
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteActivityInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *DeleteActivityInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// DeleteNetworkInsightsCosDetailsV2Options : The DeleteNetworkInsightsCosDetailsV2 options.
type DeleteNetworkInsightsCosDetailsV2Options struct {
	// Array of Ids of COS entries.
	Ids []string `json:"ids,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNetworkInsightsCosDetailsV2Options : Instantiate DeleteNetworkInsightsCosDetailsV2Options
func (*AddonManagerV1) NewDeleteNetworkInsightsCosDetailsV2Options() *DeleteNetworkInsightsCosDetailsV2Options {
	return &DeleteNetworkInsightsCosDetailsV2Options{}
}

// SetIds : Allow user to set Ids
func (_options *DeleteNetworkInsightsCosDetailsV2Options) SetIds(ids []string) *DeleteNetworkInsightsCosDetailsV2Options {
	_options.Ids = ids
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNetworkInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *DeleteNetworkInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// GetActivityInsightStatusV2Options : The GetActivityInsightStatusV2 options.
type GetActivityInsightStatusV2Options struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetActivityInsightStatusV2Options : Instantiate GetActivityInsightStatusV2Options
func (*AddonManagerV1) NewGetActivityInsightStatusV2Options() *GetActivityInsightStatusV2Options {
	return &GetActivityInsightStatusV2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetActivityInsightStatusV2Options) SetHeaders(param map[string]string) *GetActivityInsightStatusV2Options {
	options.Headers = param
	return options
}

// GetActivityInsightsCosDetailsV2Options : The GetActivityInsightsCosDetailsV2 options.
type GetActivityInsightsCosDetailsV2Options struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetActivityInsightsCosDetailsV2Options : Instantiate GetActivityInsightsCosDetailsV2Options
func (*AddonManagerV1) NewGetActivityInsightsCosDetailsV2Options() *GetActivityInsightsCosDetailsV2Options {
	return &GetActivityInsightsCosDetailsV2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetActivityInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *GetActivityInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// GetNetworkInsightStatusV2Options : The GetNetworkInsightStatusV2 options.
type GetNetworkInsightStatusV2Options struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNetworkInsightStatusV2Options : Instantiate GetNetworkInsightStatusV2Options
func (*AddonManagerV1) NewGetNetworkInsightStatusV2Options() *GetNetworkInsightStatusV2Options {
	return &GetNetworkInsightStatusV2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetNetworkInsightStatusV2Options) SetHeaders(param map[string]string) *GetNetworkInsightStatusV2Options {
	options.Headers = param
	return options
}

// GetNetworkInsightsCosDetailsV2Options : The GetNetworkInsightsCosDetailsV2 options.
type GetNetworkInsightsCosDetailsV2Options struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNetworkInsightsCosDetailsV2Options : Instantiate GetNetworkInsightsCosDetailsV2Options
func (*AddonManagerV1) NewGetNetworkInsightsCosDetailsV2Options() *GetNetworkInsightsCosDetailsV2Options {
	return &GetNetworkInsightsCosDetailsV2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetNetworkInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *GetNetworkInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// GetSupportedInsightsV2Options : The GetSupportedInsightsV2 options.
type GetSupportedInsightsV2Options struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSupportedInsightsV2Options : Instantiate GetSupportedInsightsV2Options
func (*AddonManagerV1) NewGetSupportedInsightsV2Options() *GetSupportedInsightsV2Options {
	return &GetSupportedInsightsV2Options{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSupportedInsightsV2Options) SetHeaders(param map[string]string) *GetSupportedInsightsV2Options {
	options.Headers = param
	return options
}

// NetworkInsightsCosDetailsOutput : List of COS buckets details added to Network Insights.
type NetworkInsightsCosDetailsOutput struct {
	CosDetails []CosDetailsWithID `json:"cos_details" validate:"required"`
}

// UnmarshalNetworkInsightsCosDetailsOutput unmarshals an instance of NetworkInsightsCosDetailsOutput from the specified map of raw messages.
func UnmarshalNetworkInsightsCosDetailsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetworkInsightsCosDetailsOutput)
	err = core.UnmarshalModel(m, "cos_details", &obj.CosDetails, UnmarshalCosDetailsWithID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetworkInsightsStatusConfigOutput : Network Insights status configuration.
type NetworkInsightsStatusConfigOutput struct {
	// The type of insight whose status is to be displayed.
	Addon *string `json:"addon" validate:"required"`

	// The status of the insight, ie., enable or disable.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the NetworkInsightsStatusConfigOutput.Status property.
// The status of the insight, ie., enable or disable.
const (
	NetworkInsightsStatusConfigOutputStatusDisableConst = "disable"
	NetworkInsightsStatusConfigOutputStatusEnableConst  = "enable"
)

// UnmarshalNetworkInsightsStatusConfigOutput unmarshals an instance of NetworkInsightsStatusConfigOutput from the specified map of raw messages.
func UnmarshalNetworkInsightsStatusConfigOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetworkInsightsStatusConfigOutput)
	err = core.UnmarshalPrimitive(m, "addon", &obj.Addon)
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

// UpdateActivityInsightStatusV2Options : The UpdateActivityInsightStatusV2 options.
type UpdateActivityInsightStatusV2Options struct {
	// Region ID, e.g., us, eu, etc.
	RegionID *string `json:"region_id" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateActivityInsightStatusV2Options.Status property.
const (
	UpdateActivityInsightStatusV2OptionsStatusDisableConst = "disable"
	UpdateActivityInsightStatusV2OptionsStatusEnableConst  = "enable"
)

// NewUpdateActivityInsightStatusV2Options : Instantiate UpdateActivityInsightStatusV2Options
func (*AddonManagerV1) NewUpdateActivityInsightStatusV2Options(regionID string, status string) *UpdateActivityInsightStatusV2Options {
	return &UpdateActivityInsightStatusV2Options{
		RegionID: core.StringPtr(regionID),
		Status:   core.StringPtr(status),
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *UpdateActivityInsightStatusV2Options) SetRegionID(regionID string) *UpdateActivityInsightStatusV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *UpdateActivityInsightStatusV2Options) SetStatus(status string) *UpdateActivityInsightStatusV2Options {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateActivityInsightStatusV2Options) SetHeaders(param map[string]string) *UpdateActivityInsightStatusV2Options {
	options.Headers = param
	return options
}

// UpdateNetworkInsightStatusV2Options : The UpdateNetworkInsightStatusV2 options.
type UpdateNetworkInsightStatusV2Options struct {
	// Region ID, e.g., us, eu, etc.
	RegionID *string `json:"region_id" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateNetworkInsightStatusV2Options.Status property.
const (
	UpdateNetworkInsightStatusV2OptionsStatusDisableConst = "disable"
	UpdateNetworkInsightStatusV2OptionsStatusEnableConst  = "enable"
)

// NewUpdateNetworkInsightStatusV2Options : Instantiate UpdateNetworkInsightStatusV2Options
func (*AddonManagerV1) NewUpdateNetworkInsightStatusV2Options(regionID string, status string) *UpdateNetworkInsightStatusV2Options {
	return &UpdateNetworkInsightStatusV2Options{
		RegionID: core.StringPtr(regionID),
		Status:   core.StringPtr(status),
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *UpdateNetworkInsightStatusV2Options) SetRegionID(regionID string) *UpdateNetworkInsightStatusV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *UpdateNetworkInsightStatusV2Options) SetStatus(status string) *UpdateNetworkInsightStatusV2Options {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNetworkInsightStatusV2Options) SetHeaders(param map[string]string) *UpdateNetworkInsightStatusV2Options {
	options.Headers = param
	return options
}

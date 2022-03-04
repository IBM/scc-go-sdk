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

// Package addonmgrv1 : Operations and models for the AddonMgrV1 service
package addonmgrv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/scc-go-sdk/v3/common"
)

// AddonMgrV1 : The Addon Manager API
//
// API Version: 1.0.0
type AddonMgrV1 struct {
	Service *core.BaseService

	// Account ID.
	AccountID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.secadvisor.cloud.ibm.com/addonmgr"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "addon_mgr"

// AddonMgrV1Options : Service options
type AddonMgrV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Account ID.
	AccountID *string `validate:"required"`
}

// NewAddonMgrV1UsingExternalConfig : constructs an instance of AddonMgrV1 with passed in options and external configuration.
func NewAddonMgrV1UsingExternalConfig(options *AddonMgrV1Options) (addonMgr *AddonMgrV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	addonMgr, err = NewAddonMgrV1(options)
	if err != nil {
		return
	}

	err = addonMgr.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = addonMgr.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAddonMgrV1 : constructs an instance of AddonMgrV1 with passed in options.
func NewAddonMgrV1(options *AddonMgrV1Options) (service *AddonMgrV1, err error) {
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

	service = &AddonMgrV1{
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

// Clone makes a copy of "addonMgr" suitable for processing requests.
func (addonMgr *AddonMgrV1) Clone() *AddonMgrV1 {
	if core.IsNil(addonMgr) {
		return nil
	}
	clone := *addonMgr
	clone.Service = addonMgr.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (addonMgr *AddonMgrV1) SetServiceURL(url string) error {
	return addonMgr.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (addonMgr *AddonMgrV1) GetServiceURL() string {
	return addonMgr.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (addonMgr *AddonMgrV1) SetDefaultHeaders(headers http.Header) {
	addonMgr.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (addonMgr *AddonMgrV1) SetEnableGzipCompression(enableGzip bool) {
	addonMgr.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (addonMgr *AddonMgrV1) GetEnableGzipCompression() bool {
	return addonMgr.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (addonMgr *AddonMgrV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	addonMgr.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (addonMgr *AddonMgrV1) DisableRetries() {
	addonMgr.Service.DisableRetries()
}

// AddNetworkInsightsCosDetailsV2 : Add cos details
// Addcos details.
func (addonMgr *AddonMgrV1) AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.AddNetworkInsightsCosDetailsV2WithContext(context.Background(), addNetworkInsightsCosDetailsV2Options)
}

// AddNetworkInsightsCosDetailsV2WithContext is an alternate form of the AddNetworkInsightsCosDetailsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) AddNetworkInsightsCosDetailsV2WithContext(ctx context.Context, addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addNetworkInsightsCosDetailsV2Options, "addNetworkInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addNetworkInsightsCosDetailsV2Options, "addNetworkInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/network-insights/cos`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addNetworkInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "AddNetworkInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addNetworkInsightsCosDetailsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addNetworkInsightsCosDetailsV2Options.TransactionID))
	}

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

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// DeleteNetworkInsightsCosDetailsV2 : Delete cos details
// Delete NA cos details.
func (addonMgr *AddonMgrV1) DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.DeleteNetworkInsightsCosDetailsV2WithContext(context.Background(), deleteNetworkInsightsCosDetailsV2Options)
}

// DeleteNetworkInsightsCosDetailsV2WithContext is an alternate form of the DeleteNetworkInsightsCosDetailsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) DeleteNetworkInsightsCosDetailsV2WithContext(ctx context.Context, deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNetworkInsightsCosDetailsV2Options, "deleteNetworkInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNetworkInsightsCosDetailsV2Options, "deleteNetworkInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/network-insights/cos`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNetworkInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "DeleteNetworkInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if deleteNetworkInsightsCosDetailsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteNetworkInsightsCosDetailsV2Options.TransactionID))
	}

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

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// AddActivityInsightsCosDetailsV2 : Add cos details
// Addcos details.
func (addonMgr *AddonMgrV1) AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.AddActivityInsightsCosDetailsV2WithContext(context.Background(), addActivityInsightsCosDetailsV2Options)
}

// AddActivityInsightsCosDetailsV2WithContext is an alternate form of the AddActivityInsightsCosDetailsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) AddActivityInsightsCosDetailsV2WithContext(ctx context.Context, addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addActivityInsightsCosDetailsV2Options, "addActivityInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addActivityInsightsCosDetailsV2Options, "addActivityInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/activity-insights/cos`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addActivityInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "AddActivityInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addActivityInsightsCosDetailsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addActivityInsightsCosDetailsV2Options.TransactionID))
	}

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

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// DeleteActivityInsightsCosDetailsV2 : Delete cos details
// Delete AT cos details.
func (addonMgr *AddonMgrV1) DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.DeleteActivityInsightsCosDetailsV2WithContext(context.Background(), deleteActivityInsightsCosDetailsV2Options)
}

// DeleteActivityInsightsCosDetailsV2WithContext is an alternate form of the DeleteActivityInsightsCosDetailsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) DeleteActivityInsightsCosDetailsV2WithContext(ctx context.Context, deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteActivityInsightsCosDetailsV2Options, "deleteActivityInsightsCosDetailsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteActivityInsightsCosDetailsV2Options, "deleteActivityInsightsCosDetailsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/activity-insights/cos`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteActivityInsightsCosDetailsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "DeleteActivityInsightsCosDetailsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if deleteActivityInsightsCosDetailsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteActivityInsightsCosDetailsV2Options.TransactionID))
	}

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

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// DisableInsightsV2 : Disable add-on
// Disable add-on.
func (addonMgr *AddonMgrV1) DisableInsightsV2(disableInsightsV2Options *DisableInsightsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.DisableInsightsV2WithContext(context.Background(), disableInsightsV2Options)
}

// DisableInsightsV2WithContext is an alternate form of the DisableInsightsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) DisableInsightsV2WithContext(ctx context.Context, disableInsightsV2Options *DisableInsightsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(disableInsightsV2Options, "disableInsightsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(disableInsightsV2Options, "disableInsightsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/disable`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range disableInsightsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "DisableInsightsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if disableInsightsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*disableInsightsV2Options.TransactionID))
	}

	body := make(map[string]interface{})
	if disableInsightsV2Options.RegionID != nil {
		body["region_id"] = disableInsightsV2Options.RegionID
	}
	if disableInsightsV2Options.NetworkInsights != nil {
		body["network-insights"] = disableInsightsV2Options.NetworkInsights
	}
	if disableInsightsV2Options.ActivityInsights != nil {
		body["activity-insights"] = disableInsightsV2Options.ActivityInsights
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// EnableInsightsV2 : Enable add-on
// Enable add-on.
func (addonMgr *AddonMgrV1) EnableInsightsV2(enableInsightsV2Options *EnableInsightsV2Options) (response *core.DetailedResponse, err error) {
	return addonMgr.EnableInsightsV2WithContext(context.Background(), enableInsightsV2Options)
}

// EnableInsightsV2WithContext is an alternate form of the EnableInsightsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) EnableInsightsV2WithContext(ctx context.Context, enableInsightsV2Options *EnableInsightsV2Options) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(enableInsightsV2Options, "enableInsightsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(enableInsightsV2Options, "enableInsightsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/enable`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range enableInsightsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "EnableInsightsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if enableInsightsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*enableInsightsV2Options.TransactionID))
	}

	body := make(map[string]interface{})
	if enableInsightsV2Options.RegionID != nil {
		body["region_id"] = enableInsightsV2Options.RegionID
	}
	if enableInsightsV2Options.NetworkInsights != nil {
		body["network-insights"] = enableInsightsV2Options.NetworkInsights
	}
	if enableInsightsV2Options.ActivityInsights != nil {
		body["activity-insights"] = enableInsightsV2Options.ActivityInsights
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = addonMgr.Service.Request(request, nil)

	return
}

// GetSupportedInsightsV2 : Fetch supported insights
// Retrieve insights details.
func (addonMgr *AddonMgrV1) GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options) (result *AllInsights, response *core.DetailedResponse, err error) {
	return addonMgr.GetSupportedInsightsV2WithContext(context.Background(), getSupportedInsightsV2Options)
}

// GetSupportedInsightsV2WithContext is an alternate form of the GetSupportedInsightsV2 method which supports a Context parameter
func (addonMgr *AddonMgrV1) GetSupportedInsightsV2WithContext(ctx context.Context, getSupportedInsightsV2Options *GetSupportedInsightsV2Options) (result *AllInsights, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSupportedInsightsV2Options, "getSupportedInsightsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addonMgr.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = addonMgr.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(addonMgr.Service.Options.URL, `/v2/addons/{account_id}/insights`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSupportedInsightsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("addon_mgr", "V1", "GetSupportedInsightsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSupportedInsightsV2Options.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getSupportedInsightsV2Options.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = addonMgr.Service.Request(request, &rawResponse)
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

// AddActivityInsightsCosDetailsV2Options : The AddActivityInsightsCosDetailsV2 options.
type AddActivityInsightsCosDetailsV2Options struct {
	// Region for example - us-south, eu-gb.
	RegionID *string `json:"region_id" validate:"required"`

	CosDetails []CosDetailsV2CosDetailsItem `json:"cos_details" validate:"required"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddActivityInsightsCosDetailsV2Options : Instantiate AddActivityInsightsCosDetailsV2Options
func (*AddonMgrV1) NewAddActivityInsightsCosDetailsV2Options(regionID string, cosDetails []CosDetailsV2CosDetailsItem) *AddActivityInsightsCosDetailsV2Options {
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
func (_options *AddActivityInsightsCosDetailsV2Options) SetCosDetails(cosDetails []CosDetailsV2CosDetailsItem) *AddActivityInsightsCosDetailsV2Options {
	_options.CosDetails = cosDetails
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddActivityInsightsCosDetailsV2Options) SetTransactionID(transactionID string) *AddActivityInsightsCosDetailsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddActivityInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *AddActivityInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// AddNetworkInsightsCosDetailsV2Options : The AddNetworkInsightsCosDetailsV2 options.
type AddNetworkInsightsCosDetailsV2Options struct {
	// Region for example - us-south, eu-gb.
	RegionID *string `json:"region_id" validate:"required"`

	CosDetails []CosDetailsV2CosDetailsItem `json:"cos_details" validate:"required"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddNetworkInsightsCosDetailsV2Options : Instantiate AddNetworkInsightsCosDetailsV2Options
func (*AddonMgrV1) NewAddNetworkInsightsCosDetailsV2Options(regionID string, cosDetails []CosDetailsV2CosDetailsItem) *AddNetworkInsightsCosDetailsV2Options {
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
func (_options *AddNetworkInsightsCosDetailsV2Options) SetCosDetails(cosDetails []CosDetailsV2CosDetailsItem) *AddNetworkInsightsCosDetailsV2Options {
	_options.CosDetails = cosDetails
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddNetworkInsightsCosDetailsV2Options) SetTransactionID(transactionID string) *AddNetworkInsightsCosDetailsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddNetworkInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *AddNetworkInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// AllInsights : AllInsights struct
type AllInsights struct {
	Type []string `json:"type,omitempty"`
}

// Constants associated with the AllInsights.Type property.
// Insights type.
const (
	AllInsightsTypeActivityInsightsConst = "activity-insights"
	AllInsightsTypeNetworkInsightsConst  = "network-insights"
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

// CosDetailsV2CosDetailsItem : CosDetailsV2CosDetailsItem struct
type CosDetailsV2CosDetailsItem struct {
	CosInstance *string `json:"cos_instance" validate:"required"`

	BucketName *string `json:"bucket_name" validate:"required"`

	Description *string `json:"description" validate:"required"`

	// Insights type.
	Type *string `json:"type" validate:"required"`

	// cos bucket url.
	CosBucketURL *string `json:"cos_bucket_url" validate:"required"`
}

// Constants associated with the CosDetailsV2CosDetailsItem.Type property.
// Insights type.
const (
	CosDetailsV2CosDetailsItemTypeActivityInsightsConst = "activity-insights"
	CosDetailsV2CosDetailsItemTypeNetworkInsightsConst  = "network-insights"
)

// NewCosDetailsV2CosDetailsItem : Instantiate CosDetailsV2CosDetailsItem (Generic Model Constructor)
func (*AddonMgrV1) NewCosDetailsV2CosDetailsItem(cosInstance string, bucketName string, description string, typeVar string, cosBucketURL string) (_model *CosDetailsV2CosDetailsItem, err error) {
	_model = &CosDetailsV2CosDetailsItem{
		CosInstance:  core.StringPtr(cosInstance),
		BucketName:   core.StringPtr(bucketName),
		Description:  core.StringPtr(description),
		Type:         core.StringPtr(typeVar),
		CosBucketURL: core.StringPtr(cosBucketURL),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCosDetailsV2CosDetailsItem unmarshals an instance of CosDetailsV2CosDetailsItem from the specified map of raw messages.
func UnmarshalCosDetailsV2CosDetailsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosDetailsV2CosDetailsItem)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// DeleteActivityInsightsCosDetailsV2Options : The DeleteActivityInsightsCosDetailsV2 options.
type DeleteActivityInsightsCosDetailsV2Options struct {
	// Array of Ids of COS entries.
	Ids []string `json:"ids,omitempty"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteActivityInsightsCosDetailsV2Options : Instantiate DeleteActivityInsightsCosDetailsV2Options
func (*AddonMgrV1) NewDeleteActivityInsightsCosDetailsV2Options() *DeleteActivityInsightsCosDetailsV2Options {
	return &DeleteActivityInsightsCosDetailsV2Options{}
}

// SetIds : Allow user to set Ids
func (_options *DeleteActivityInsightsCosDetailsV2Options) SetIds(ids []string) *DeleteActivityInsightsCosDetailsV2Options {
	_options.Ids = ids
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteActivityInsightsCosDetailsV2Options) SetTransactionID(transactionID string) *DeleteActivityInsightsCosDetailsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
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

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNetworkInsightsCosDetailsV2Options : Instantiate DeleteNetworkInsightsCosDetailsV2Options
func (*AddonMgrV1) NewDeleteNetworkInsightsCosDetailsV2Options() *DeleteNetworkInsightsCosDetailsV2Options {
	return &DeleteNetworkInsightsCosDetailsV2Options{}
}

// SetIds : Allow user to set Ids
func (_options *DeleteNetworkInsightsCosDetailsV2Options) SetIds(ids []string) *DeleteNetworkInsightsCosDetailsV2Options {
	_options.Ids = ids
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteNetworkInsightsCosDetailsV2Options) SetTransactionID(transactionID string) *DeleteNetworkInsightsCosDetailsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNetworkInsightsCosDetailsV2Options) SetHeaders(param map[string]string) *DeleteNetworkInsightsCosDetailsV2Options {
	options.Headers = param
	return options
}

// DisableInsightsV2Options : The DisableInsightsV2 options.
type DisableInsightsV2Options struct {
	// Region id for example - us.
	RegionID *string `json:"region_id" validate:"required"`

	NetworkInsights *bool `json:"network-insights,omitempty"`

	ActivityInsights *bool `json:"activity-insights,omitempty"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDisableInsightsV2Options : Instantiate DisableInsightsV2Options
func (*AddonMgrV1) NewDisableInsightsV2Options(regionID string) *DisableInsightsV2Options {
	return &DisableInsightsV2Options{
		RegionID: core.StringPtr(regionID),
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *DisableInsightsV2Options) SetRegionID(regionID string) *DisableInsightsV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetNetworkInsights : Allow user to set NetworkInsights
func (_options *DisableInsightsV2Options) SetNetworkInsights(networkInsights bool) *DisableInsightsV2Options {
	_options.NetworkInsights = core.BoolPtr(networkInsights)
	return _options
}

// SetActivityInsights : Allow user to set ActivityInsights
func (_options *DisableInsightsV2Options) SetActivityInsights(activityInsights bool) *DisableInsightsV2Options {
	_options.ActivityInsights = core.BoolPtr(activityInsights)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DisableInsightsV2Options) SetTransactionID(transactionID string) *DisableInsightsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DisableInsightsV2Options) SetHeaders(param map[string]string) *DisableInsightsV2Options {
	options.Headers = param
	return options
}

// EnableInsightsV2Options : The EnableInsightsV2 options.
type EnableInsightsV2Options struct {
	// Region id for example - us.
	RegionID *string `json:"region_id" validate:"required"`

	NetworkInsights *bool `json:"network-insights,omitempty"`

	ActivityInsights *bool `json:"activity-insights,omitempty"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEnableInsightsV2Options : Instantiate EnableInsightsV2Options
func (*AddonMgrV1) NewEnableInsightsV2Options(regionID string) *EnableInsightsV2Options {
	return &EnableInsightsV2Options{
		RegionID: core.StringPtr(regionID),
	}
}

// SetRegionID : Allow user to set RegionID
func (_options *EnableInsightsV2Options) SetRegionID(regionID string) *EnableInsightsV2Options {
	_options.RegionID = core.StringPtr(regionID)
	return _options
}

// SetNetworkInsights : Allow user to set NetworkInsights
func (_options *EnableInsightsV2Options) SetNetworkInsights(networkInsights bool) *EnableInsightsV2Options {
	_options.NetworkInsights = core.BoolPtr(networkInsights)
	return _options
}

// SetActivityInsights : Allow user to set ActivityInsights
func (_options *EnableInsightsV2Options) SetActivityInsights(activityInsights bool) *EnableInsightsV2Options {
	_options.ActivityInsights = core.BoolPtr(activityInsights)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *EnableInsightsV2Options) SetTransactionID(transactionID string) *EnableInsightsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *EnableInsightsV2Options) SetHeaders(param map[string]string) *EnableInsightsV2Options {
	options.Headers = param
	return options
}

// GetSupportedInsightsV2Options : The GetSupportedInsightsV2 options.
type GetSupportedInsightsV2Options struct {
	// The transaction id for the request in uuid v4 format.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSupportedInsightsV2Options : Instantiate GetSupportedInsightsV2Options
func (*AddonMgrV1) NewGetSupportedInsightsV2Options() *GetSupportedInsightsV2Options {
	return &GetSupportedInsightsV2Options{}
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetSupportedInsightsV2Options) SetTransactionID(transactionID string) *GetSupportedInsightsV2Options {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSupportedInsightsV2Options) SetHeaders(param map[string]string) *GetSupportedInsightsV2Options {
	options.Headers = param
	return options
}

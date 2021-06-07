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

// Package notificationsv1 : Operations and models for the NotificationsV1 service
package notificationsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/ibm/scc-go-sdk/common"
)

// NotificationsV1 : API specification for the Notifications service.
//
// Version: 1.0.0
type NotificationsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.secadvisor.cloud.ibm.com/notifications"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "notifications"

// NotificationsV1Options : Service options
type NotificationsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewNotificationsV1UsingExternalConfig : constructs an instance of NotificationsV1 with passed in options and external configuration.
func NewNotificationsV1UsingExternalConfig(options *NotificationsV1Options) (notifications *NotificationsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	notifications, err = NewNotificationsV1(options)
	if err != nil {
		return
	}

	err = notifications.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = notifications.Service.SetServiceURL(options.URL)
	}
	return
}

// NewNotificationsV1 : constructs an instance of NotificationsV1 with passed in options.
func NewNotificationsV1(options *NotificationsV1Options) (service *NotificationsV1, err error) {
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

	service = &NotificationsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-south": "https://us-south.secadvisor.cloud.ibm.com/notifications",
		"us-east":  "https://us-south.secadvisor.cloud.ibm.com/notifications",
		"eu-gb":    "https://eu-gb.secadvisor.cloud.ibm.com/notifications",
		"eu-de":    "https://eu.compliance.cloud.ibm.com/si/notifications",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "notifications" suitable for processing requests.
func (notifications *NotificationsV1) Clone() *NotificationsV1 {
	if core.IsNil(notifications) {
		return nil
	}
	clone := *notifications
	clone.Service = notifications.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (notifications *NotificationsV1) SetServiceURL(url string) error {
	return notifications.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (notifications *NotificationsV1) GetServiceURL() string {
	return notifications.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (notifications *NotificationsV1) SetDefaultHeaders(headers http.Header) {
	notifications.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (notifications *NotificationsV1) SetEnableGzipCompression(enableGzip bool) {
	notifications.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (notifications *NotificationsV1) GetEnableGzipCompression() bool {
	return notifications.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (notifications *NotificationsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	notifications.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (notifications *NotificationsV1) DisableRetries() {
	notifications.Service.DisableRetries()
}

// ListAllChannels : list all channels
// list all channels under this account.
func (notifications *NotificationsV1) ListAllChannels(listAllChannelsOptions *ListAllChannelsOptions) (result *ChannelsList, response *core.DetailedResponse, err error) {
	return notifications.ListAllChannelsWithContext(context.Background(), listAllChannelsOptions)
}

// ListAllChannelsWithContext is an alternate form of the ListAllChannels method which supports a Context parameter
func (notifications *NotificationsV1) ListAllChannelsWithContext(ctx context.Context, listAllChannelsOptions *ListAllChannelsOptions) (result *ChannelsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAllChannelsOptions, "listAllChannelsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAllChannelsOptions, "listAllChannelsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *listAllChannelsOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAllChannelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "ListAllChannels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAllChannelsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAllChannelsOptions.TransactionID))
	}

	if listAllChannelsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAllChannelsOptions.Limit))
	}
	if listAllChannelsOptions.Skip != nil {
		builder.AddQuery("skip", fmt.Sprint(*listAllChannelsOptions.Skip))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateNotificationChannel : create notification channel
// create notification channel.
func (notifications *NotificationsV1) CreateNotificationChannel(createNotificationChannelOptions *CreateNotificationChannelOptions) (result *ChannelInfo, response *core.DetailedResponse, err error) {
	return notifications.CreateNotificationChannelWithContext(context.Background(), createNotificationChannelOptions)
}

// CreateNotificationChannelWithContext is an alternate form of the CreateNotificationChannel method which supports a Context parameter
func (notifications *NotificationsV1) CreateNotificationChannelWithContext(ctx context.Context, createNotificationChannelOptions *CreateNotificationChannelOptions) (result *ChannelInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNotificationChannelOptions, "createNotificationChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createNotificationChannelOptions, "createNotificationChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *createNotificationChannelOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createNotificationChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "CreateNotificationChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createNotificationChannelOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createNotificationChannelOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createNotificationChannelOptions.Name != nil {
		body["name"] = createNotificationChannelOptions.Name
	}
	if createNotificationChannelOptions.Type != nil {
		body["type"] = createNotificationChannelOptions.Type
	}
	if createNotificationChannelOptions.Endpoint != nil {
		body["endpoint"] = createNotificationChannelOptions.Endpoint
	}
	if createNotificationChannelOptions.Description != nil {
		body["description"] = createNotificationChannelOptions.Description
	}
	if createNotificationChannelOptions.Severity != nil {
		body["severity"] = createNotificationChannelOptions.Severity
	}
	if createNotificationChannelOptions.Enabled != nil {
		body["enabled"] = createNotificationChannelOptions.Enabled
	}
	if createNotificationChannelOptions.AlertSource != nil {
		body["alert_source"] = createNotificationChannelOptions.AlertSource
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
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelInfo)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteNotificationChannels : bulk delete of channels
// bulk delete of channels.
func (notifications *NotificationsV1) DeleteNotificationChannels(deleteNotificationChannelsOptions *DeleteNotificationChannelsOptions) (result *ChannelsDelete, response *core.DetailedResponse, err error) {
	return notifications.DeleteNotificationChannelsWithContext(context.Background(), deleteNotificationChannelsOptions)
}

// DeleteNotificationChannelsWithContext is an alternate form of the DeleteNotificationChannels method which supports a Context parameter
func (notifications *NotificationsV1) DeleteNotificationChannelsWithContext(ctx context.Context, deleteNotificationChannelsOptions *DeleteNotificationChannelsOptions) (result *ChannelsDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNotificationChannelsOptions, "deleteNotificationChannelsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNotificationChannelsOptions, "deleteNotificationChannelsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *deleteNotificationChannelsOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNotificationChannelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "DeleteNotificationChannels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if deleteNotificationChannelsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteNotificationChannelsOptions.TransactionID))
	}

	_, err = builder.SetBodyContentJSON(deleteNotificationChannelsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelsDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteNotificationChannel : delete the details of a specific channel
// delete the details of a specific channel.
func (notifications *NotificationsV1) DeleteNotificationChannel(deleteNotificationChannelOptions *DeleteNotificationChannelOptions) (result *ChannelDelete, response *core.DetailedResponse, err error) {
	return notifications.DeleteNotificationChannelWithContext(context.Background(), deleteNotificationChannelOptions)
}

// DeleteNotificationChannelWithContext is an alternate form of the DeleteNotificationChannel method which supports a Context parameter
func (notifications *NotificationsV1) DeleteNotificationChannelWithContext(ctx context.Context, deleteNotificationChannelOptions *DeleteNotificationChannelOptions) (result *ChannelDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNotificationChannelOptions, "deleteNotificationChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNotificationChannelOptions, "deleteNotificationChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *deleteNotificationChannelOptions.AccountID,
		"channel_id": *deleteNotificationChannelOptions.ChannelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels/{channel_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNotificationChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "DeleteNotificationChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteNotificationChannelOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteNotificationChannelOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetNotificationChannel : get the details of a specific channel
// get the details of a specific channel.
func (notifications *NotificationsV1) GetNotificationChannel(getNotificationChannelOptions *GetNotificationChannelOptions) (result *ChannelGet, response *core.DetailedResponse, err error) {
	return notifications.GetNotificationChannelWithContext(context.Background(), getNotificationChannelOptions)
}

// GetNotificationChannelWithContext is an alternate form of the GetNotificationChannel method which supports a Context parameter
func (notifications *NotificationsV1) GetNotificationChannelWithContext(ctx context.Context, getNotificationChannelOptions *GetNotificationChannelOptions) (result *ChannelGet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getNotificationChannelOptions, "getNotificationChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getNotificationChannelOptions, "getNotificationChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getNotificationChannelOptions.AccountID,
		"channel_id": *getNotificationChannelOptions.ChannelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels/{channel_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNotificationChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "GetNotificationChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getNotificationChannelOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getNotificationChannelOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelGet)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateNotificationChannel : update notification channel
// update notification channel.
func (notifications *NotificationsV1) UpdateNotificationChannel(updateNotificationChannelOptions *UpdateNotificationChannelOptions) (result *ChannelInfo, response *core.DetailedResponse, err error) {
	return notifications.UpdateNotificationChannelWithContext(context.Background(), updateNotificationChannelOptions)
}

// UpdateNotificationChannelWithContext is an alternate form of the UpdateNotificationChannel method which supports a Context parameter
func (notifications *NotificationsV1) UpdateNotificationChannelWithContext(ctx context.Context, updateNotificationChannelOptions *UpdateNotificationChannelOptions) (result *ChannelInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNotificationChannelOptions, "updateNotificationChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateNotificationChannelOptions, "updateNotificationChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *updateNotificationChannelOptions.AccountID,
		"channel_id": *updateNotificationChannelOptions.ChannelID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels/{channel_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateNotificationChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "UpdateNotificationChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateNotificationChannelOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateNotificationChannelOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateNotificationChannelOptions.Name != nil {
		body["name"] = updateNotificationChannelOptions.Name
	}
	if updateNotificationChannelOptions.Type != nil {
		body["type"] = updateNotificationChannelOptions.Type
	}
	if updateNotificationChannelOptions.Endpoint != nil {
		body["endpoint"] = updateNotificationChannelOptions.Endpoint
	}
	if updateNotificationChannelOptions.Description != nil {
		body["description"] = updateNotificationChannelOptions.Description
	}
	if updateNotificationChannelOptions.Severity != nil {
		body["severity"] = updateNotificationChannelOptions.Severity
	}
	if updateNotificationChannelOptions.Enabled != nil {
		body["enabled"] = updateNotificationChannelOptions.Enabled
	}
	if updateNotificationChannelOptions.AlertSource != nil {
		body["alert_source"] = updateNotificationChannelOptions.AlertSource
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
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChannelInfo)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TestNotificationChannel : test notification channel
// test a nofication channel under this account.
func (notifications *NotificationsV1) TestNotificationChannel(testNotificationChannelOptions *TestNotificationChannelOptions) (result *TestChannel, response *core.DetailedResponse, err error) {
	return notifications.TestNotificationChannelWithContext(context.Background(), testNotificationChannelOptions)
}

// TestNotificationChannelWithContext is an alternate form of the TestNotificationChannel method which supports a Context parameter
func (notifications *NotificationsV1) TestNotificationChannelWithContext(ctx context.Context, testNotificationChannelOptions *TestNotificationChannelOptions) (result *TestChannel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(testNotificationChannelOptions, "testNotificationChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(testNotificationChannelOptions, "testNotificationChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *testNotificationChannelOptions.AccountID,
		"channel_id": *testNotificationChannelOptions.ChannelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/channels/{channel_id}/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range testNotificationChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "TestNotificationChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if testNotificationChannelOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*testNotificationChannelOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestChannel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPublicKey : fetch notifications public key
// fetch public key to decrypt messages in notification payload.
func (notifications *NotificationsV1) GetPublicKey(getPublicKeyOptions *GetPublicKeyOptions) (result *PublicKeyGet, response *core.DetailedResponse, err error) {
	return notifications.GetPublicKeyWithContext(context.Background(), getPublicKeyOptions)
}

// GetPublicKeyWithContext is an alternate form of the GetPublicKey method which supports a Context parameter
func (notifications *NotificationsV1) GetPublicKeyWithContext(ctx context.Context, getPublicKeyOptions *GetPublicKeyOptions) (result *PublicKeyGet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPublicKeyOptions, "getPublicKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPublicKeyOptions, "getPublicKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getPublicKeyOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = notifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(notifications.Service.Options.URL, `/v1/{account_id}/notifications/public_key`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPublicKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("notifications", "V1", "GetPublicKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPublicKeyOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getPublicKeyOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = notifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPublicKeyGet)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ChannelAlertSourceItem : The providers that act as alert sources and the potential findings that can be flagged as alerts.
type ChannelAlertSourceItem struct {
	// The providers that you can receive alerts for. To view your available providers, you can call the
	// /v1/{account_id}/providers endpoint of the Findings API.
	ProviderName *string `json:"provider_name,omitempty"`

	// The types of findings for each provider that you want to receive alerts for. Options are dependent upon the provider
	// that you select. Depending on that selection, some available options include `image_with_vulnerabilities`,
	// `anonym_server`, `server_suspected_ratio`, `appid`, `cos`, `expired_cert`, and `expiring_1day_cert`For a full list
	// of available finding types, see [the docs](/docs/).
	FindingTypes []interface{} `json:"finding_types,omitempty"`
}

// Constants associated with the ChannelAlertSourceItem.ProviderName property.
// The providers that you can receive alerts for. To view your available providers, you can call the
// /v1/{account_id}/providers endpoint of the Findings API.
const (
	ChannelAlertSourceItemProviderNameAllConst  = "ALL"
	ChannelAlertSourceItemProviderNameAtaConst  = "ATA"
	ChannelAlertSourceItemProviderNameCertConst = "CERT"
	ChannelAlertSourceItemProviderNameNaConst   = "NA"
	ChannelAlertSourceItemProviderNameVaConst   = "VA"
)

// UnmarshalChannelAlertSourceItem unmarshals an instance of ChannelAlertSourceItem from the specified map of raw messages.
func UnmarshalChannelAlertSourceItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelAlertSourceItem)
	err = core.UnmarshalPrimitive(m, "provider_name", &obj.ProviderName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finding_types", &obj.FindingTypes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelDelete : The returned response when a channel is deleted.
type ChannelDelete struct {
	// The ID of the deleted channel.
	ChannelID *string `json:"channel_id,omitempty"`

	// response message.
	Message *string `json:"message,omitempty"`
}

// UnmarshalChannelDelete unmarshals an instance of ChannelDelete from the specified map of raw messages.
func UnmarshalChannelDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelDelete)
	err = core.UnmarshalPrimitive(m, "channel_id", &obj.ChannelID)
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

// ChannelGet : The returned response when get channel is run.
type ChannelGet struct {
	// get channel.
	Channel *Channel `json:"channel,omitempty"`
}

// UnmarshalChannelGet unmarshals an instance of ChannelGet from the specified map of raw messages.
func UnmarshalChannelGet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelGet)
	err = core.UnmarshalModel(m, "channel", &obj.Channel, UnmarshalChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelInfo : The returned response when a channel is created or updated.
type ChannelInfo struct {
	// The ID of the created channel.
	ChannelID *string `json:"channel_id,omitempty"`

	// response code.
	StatusCode *int64 `json:"status_code,omitempty"`
}

// UnmarshalChannelInfo unmarshals an instance of ChannelInfo from the specified map of raw messages.
func UnmarshalChannelInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelInfo)
	err = core.UnmarshalPrimitive(m, "channel_id", &obj.ChannelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelSeverity : The severity of the notification.
type ChannelSeverity struct {
	// Critical severity.
	Critical *bool `json:"critical,omitempty"`

	// High severity.
	High *bool `json:"high,omitempty"`

	// Medium severity.
	Medium *bool `json:"medium,omitempty"`

	// Low severity.
	Low *bool `json:"low,omitempty"`
}

// UnmarshalChannelSeverity unmarshals an instance of ChannelSeverity from the specified map of raw messages.
func UnmarshalChannelSeverity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelSeverity)
	err = core.UnmarshalPrimitive(m, "critical", &obj.Critical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "high", &obj.High)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "medium", &obj.Medium)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "low", &obj.Low)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelsDelete : The returned response when more than one channel is deleted.
type ChannelsDelete struct {
	// response message.
	Message *string `json:"message,omitempty"`
}

// UnmarshalChannelsDelete unmarshals an instance of ChannelsDelete from the specified map of raw messages.
func UnmarshalChannelsDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelsDelete)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChannelsList : Available channels in your account are listed.
type ChannelsList struct {
	// List of channels.
	Channels []Channel `json:"channels,omitempty"`
}

// UnmarshalChannelsList unmarshals an instance of ChannelsList from the specified map of raw messages.
func UnmarshalChannelsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChannelsList)
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalChannel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateNotificationChannelOptions : The CreateNotificationChannel options.
type CreateNotificationChannelOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// The name of the notification channel in the form "v1/{account_id}/notifications/channelName".
	Name *string `validate:"required"`

	// Type of callback URL.
	Type *string `validate:"required"`

	// The callback URL which receives the notification.
	Endpoint *string `validate:"required"`

	// A one sentence description of this `Channel`.
	Description *string

	// Severity of the notification to be received.
	Severity []string

	// Channel is enabled or not. Default is disabled.
	Enabled *bool

	AlertSource []NotificationChannelAlertSourceItem

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateNotificationChannelOptions.Type property.
// Type of callback URL.
const (
	CreateNotificationChannelOptionsTypeWebhookConst = "Webhook"
)

// Constants associated with the CreateNotificationChannelOptions.Severity property.
const (
	CreateNotificationChannelOptionsSeverityCriticalConst = "critical"
	CreateNotificationChannelOptionsSeverityHighConst     = "high"
	CreateNotificationChannelOptionsSeverityLowConst      = "low"
	CreateNotificationChannelOptionsSeverityMediumConst   = "medium"
)

// NewCreateNotificationChannelOptions : Instantiate CreateNotificationChannelOptions
func (*NotificationsV1) NewCreateNotificationChannelOptions(accountID string, name string, typeVar string, endpoint string) *CreateNotificationChannelOptions {
	return &CreateNotificationChannelOptions{
		AccountID: core.StringPtr(accountID),
		Name:      core.StringPtr(name),
		Type:      core.StringPtr(typeVar),
		Endpoint:  core.StringPtr(endpoint),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateNotificationChannelOptions) SetAccountID(accountID string) *CreateNotificationChannelOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateNotificationChannelOptions) SetName(name string) *CreateNotificationChannelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *CreateNotificationChannelOptions) SetType(typeVar string) *CreateNotificationChannelOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetEndpoint : Allow user to set Endpoint
func (options *CreateNotificationChannelOptions) SetEndpoint(endpoint string) *CreateNotificationChannelOptions {
	options.Endpoint = core.StringPtr(endpoint)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateNotificationChannelOptions) SetDescription(description string) *CreateNotificationChannelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetSeverity : Allow user to set Severity
func (options *CreateNotificationChannelOptions) SetSeverity(severity []string) *CreateNotificationChannelOptions {
	options.Severity = severity
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *CreateNotificationChannelOptions) SetEnabled(enabled bool) *CreateNotificationChannelOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetAlertSource : Allow user to set AlertSource
func (options *CreateNotificationChannelOptions) SetAlertSource(alertSource []NotificationChannelAlertSourceItem) *CreateNotificationChannelOptions {
	options.AlertSource = alertSource
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateNotificationChannelOptions) SetTransactionID(transactionID string) *CreateNotificationChannelOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNotificationChannelOptions) SetHeaders(param map[string]string) *CreateNotificationChannelOptions {
	options.Headers = param
	return options
}

// DeleteNotificationChannelOptions : The DeleteNotificationChannel options.
type DeleteNotificationChannelOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// Channel ID.
	ChannelID *string `validate:"required,ne="`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNotificationChannelOptions : Instantiate DeleteNotificationChannelOptions
func (*NotificationsV1) NewDeleteNotificationChannelOptions(accountID string, channelID string) *DeleteNotificationChannelOptions {
	return &DeleteNotificationChannelOptions{
		AccountID: core.StringPtr(accountID),
		ChannelID: core.StringPtr(channelID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *DeleteNotificationChannelOptions) SetAccountID(accountID string) *DeleteNotificationChannelOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetChannelID : Allow user to set ChannelID
func (options *DeleteNotificationChannelOptions) SetChannelID(channelID string) *DeleteNotificationChannelOptions {
	options.ChannelID = core.StringPtr(channelID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteNotificationChannelOptions) SetTransactionID(transactionID string) *DeleteNotificationChannelOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNotificationChannelOptions) SetHeaders(param map[string]string) *DeleteNotificationChannelOptions {
	options.Headers = param
	return options
}

// DeleteNotificationChannelsOptions : The DeleteNotificationChannels options.
type DeleteNotificationChannelsOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// Body for bulk delete notification channels.
	Body []string `validate:"required"`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNotificationChannelsOptions : Instantiate DeleteNotificationChannelsOptions
func (*NotificationsV1) NewDeleteNotificationChannelsOptions(accountID string, body []string) *DeleteNotificationChannelsOptions {
	return &DeleteNotificationChannelsOptions{
		AccountID: core.StringPtr(accountID),
		Body:      body,
	}
}

// SetAccountID : Allow user to set AccountID
func (options *DeleteNotificationChannelsOptions) SetAccountID(accountID string) *DeleteNotificationChannelsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetBody : Allow user to set Body
func (options *DeleteNotificationChannelsOptions) SetBody(body []string) *DeleteNotificationChannelsOptions {
	options.Body = body
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteNotificationChannelsOptions) SetTransactionID(transactionID string) *DeleteNotificationChannelsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNotificationChannelsOptions) SetHeaders(param map[string]string) *DeleteNotificationChannelsOptions {
	options.Headers = param
	return options
}

// GetNotificationChannelOptions : The GetNotificationChannel options.
type GetNotificationChannelOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// Channel ID.
	ChannelID *string `validate:"required,ne="`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNotificationChannelOptions : Instantiate GetNotificationChannelOptions
func (*NotificationsV1) NewGetNotificationChannelOptions(accountID string, channelID string) *GetNotificationChannelOptions {
	return &GetNotificationChannelOptions{
		AccountID: core.StringPtr(accountID),
		ChannelID: core.StringPtr(channelID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetNotificationChannelOptions) SetAccountID(accountID string) *GetNotificationChannelOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetChannelID : Allow user to set ChannelID
func (options *GetNotificationChannelOptions) SetChannelID(channelID string) *GetNotificationChannelOptions {
	options.ChannelID = core.StringPtr(channelID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetNotificationChannelOptions) SetTransactionID(transactionID string) *GetNotificationChannelOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetNotificationChannelOptions) SetHeaders(param map[string]string) *GetNotificationChannelOptions {
	options.Headers = param
	return options
}

// GetPublicKeyOptions : The GetPublicKey options.
type GetPublicKeyOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPublicKeyOptions : Instantiate GetPublicKeyOptions
func (*NotificationsV1) NewGetPublicKeyOptions(accountID string) *GetPublicKeyOptions {
	return &GetPublicKeyOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetPublicKeyOptions) SetAccountID(accountID string) *GetPublicKeyOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetPublicKeyOptions) SetTransactionID(transactionID string) *GetPublicKeyOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPublicKeyOptions) SetHeaders(param map[string]string) *GetPublicKeyOptions {
	options.Headers = param
	return options
}

// ListAllChannelsOptions : The ListAllChannels options.
type ListAllChannelsOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Limit the number of the returned documents to the specified number.
	Limit *int64

	// The offset is the index of the item from which you want to start returning data from. Default is 0.
	Skip *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAllChannelsOptions : Instantiate ListAllChannelsOptions
func (*NotificationsV1) NewListAllChannelsOptions(accountID string) *ListAllChannelsOptions {
	return &ListAllChannelsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListAllChannelsOptions) SetAccountID(accountID string) *ListAllChannelsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListAllChannelsOptions) SetTransactionID(transactionID string) *ListAllChannelsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAllChannelsOptions) SetLimit(limit int64) *ListAllChannelsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetSkip : Allow user to set Skip
func (options *ListAllChannelsOptions) SetSkip(skip int64) *ListAllChannelsOptions {
	options.Skip = core.Int64Ptr(skip)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllChannelsOptions) SetHeaders(param map[string]string) *ListAllChannelsOptions {
	options.Headers = param
	return options
}

// NotificationChannelAlertSourceItem : The alert sources. They identify the providers and their finding types which makes the findings available to Security
// Advisor.
type NotificationChannelAlertSourceItem struct {
	// Below is a list of builtin providers that you can select in addition to the ones you obtain by calling Findings API
	// /v1/{account_id}/providers :
	//  | provider_name | The source they represent |
	//  |-----|-----|
	//  | VA  | Vulnerable image findings|
	//  | NA  | Network Insights findings|
	//  | ATA | Activity Insights findings|
	//  | CERT | Certificate Manager findings|
	//  | ALL | Special provider name to represent all the providers. Its mutually exclusive with other providers meaning
	// either you choose ALL or you don't|.
	ProviderName *string `json:"provider_name" validate:"required"`

	// An array of the finding types of the provider_name or "ALL" to specify all finding types under that provider Below
	// is a list of supported finding types for each built in providers
	// | provider_name | Supported finding types |
	// |-----|-----|
	// | VA  | "image_with_vulnerabilities", "image_with_config_issues"|
	// | NA  | "anonym_server", "malware_server", "bot_server", "miner_server", "server_suspected_ratio",
	// "server_response", "data_extrusion", "server_weaponized_total"|
	// | ATA | "appid", "cos", "iks", "iam", "kms", "cert", "account", "app"|
	// | CERT | "expired_cert", "expiring_1day_cert", "expiring_10day_cert", "expiring_30day_cert", "expiring_60day_cert",
	// "expiring_90day_cert"|
	// | ALL | "ALL"|.
	FindingTypes []string `json:"finding_types,omitempty"`
}

// NewNotificationChannelAlertSourceItem : Instantiate NotificationChannelAlertSourceItem (Generic Model Constructor)
func (*NotificationsV1) NewNotificationChannelAlertSourceItem(providerName string) (model *NotificationChannelAlertSourceItem, err error) {
	model = &NotificationChannelAlertSourceItem{
		ProviderName: core.StringPtr(providerName),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalNotificationChannelAlertSourceItem unmarshals an instance of NotificationChannelAlertSourceItem from the specified map of raw messages.
func UnmarshalNotificationChannelAlertSourceItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationChannelAlertSourceItem)
	err = core.UnmarshalPrimitive(m, "provider_name", &obj.ProviderName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finding_types", &obj.FindingTypes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PublicKeyGet : PublicKeyGet struct
type PublicKeyGet struct {
	PublicKey *string `json:"public_key" validate:"required"`
}

// UnmarshalPublicKeyGet unmarshals an instance of PublicKeyGet from the specified map of raw messages.
func UnmarshalPublicKeyGet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PublicKeyGet)
	err = core.UnmarshalPrimitive(m, "public_key", &obj.PublicKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestChannel : The returned response when a webhook is tested for a channel.
type TestChannel struct {
	// response status.
	Test *string `json:"test,omitempty"`
}

// UnmarshalTestChannel unmarshals an instance of TestChannel from the specified map of raw messages.
func UnmarshalTestChannel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestChannel)
	err = core.UnmarshalPrimitive(m, "test", &obj.Test)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestNotificationChannelOptions : The TestNotificationChannel options.
type TestNotificationChannelOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// Channel ID.
	ChannelID *string `validate:"required,ne="`

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTestNotificationChannelOptions : Instantiate TestNotificationChannelOptions
func (*NotificationsV1) NewTestNotificationChannelOptions(accountID string, channelID string) *TestNotificationChannelOptions {
	return &TestNotificationChannelOptions{
		AccountID: core.StringPtr(accountID),
		ChannelID: core.StringPtr(channelID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *TestNotificationChannelOptions) SetAccountID(accountID string) *TestNotificationChannelOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetChannelID : Allow user to set ChannelID
func (options *TestNotificationChannelOptions) SetChannelID(channelID string) *TestNotificationChannelOptions {
	options.ChannelID = core.StringPtr(channelID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *TestNotificationChannelOptions) SetTransactionID(transactionID string) *TestNotificationChannelOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TestNotificationChannelOptions) SetHeaders(param map[string]string) *TestNotificationChannelOptions {
	options.Headers = param
	return options
}

// UpdateNotificationChannelOptions : The UpdateNotificationChannel options.
type UpdateNotificationChannelOptions struct {
	// Account ID.
	AccountID *string `validate:"required,ne="`

	// Channel ID.
	ChannelID *string `validate:"required,ne="`

	// The name of the notification channel in the form "v1/{account_id}/notifications/channelName".
	Name *string `validate:"required"`

	// Type of callback URL.
	Type *string `validate:"required"`

	// The callback URL which receives the notification.
	Endpoint *string `validate:"required"`

	// A one sentence description of this `Channel`.
	Description *string

	// Severity of the notification to be received.
	Severity []string

	// Channel is enabled or not. Default is disabled.
	Enabled *bool

	AlertSource []NotificationChannelAlertSourceItem

	// The transaction id for the request in uuid v4 format.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateNotificationChannelOptions.Type property.
// Type of callback URL.
const (
	UpdateNotificationChannelOptionsTypeWebhookConst = "Webhook"
)

// Constants associated with the UpdateNotificationChannelOptions.Severity property.
const (
	UpdateNotificationChannelOptionsSeverityCriticalConst = "critical"
	UpdateNotificationChannelOptionsSeverityHighConst     = "high"
	UpdateNotificationChannelOptionsSeverityLowConst      = "low"
	UpdateNotificationChannelOptionsSeverityMediumConst   = "medium"
)

// NewUpdateNotificationChannelOptions : Instantiate UpdateNotificationChannelOptions
func (*NotificationsV1) NewUpdateNotificationChannelOptions(accountID string, channelID string, name string, typeVar string, endpoint string) *UpdateNotificationChannelOptions {
	return &UpdateNotificationChannelOptions{
		AccountID: core.StringPtr(accountID),
		ChannelID: core.StringPtr(channelID),
		Name:      core.StringPtr(name),
		Type:      core.StringPtr(typeVar),
		Endpoint:  core.StringPtr(endpoint),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateNotificationChannelOptions) SetAccountID(accountID string) *UpdateNotificationChannelOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetChannelID : Allow user to set ChannelID
func (options *UpdateNotificationChannelOptions) SetChannelID(channelID string) *UpdateNotificationChannelOptions {
	options.ChannelID = core.StringPtr(channelID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateNotificationChannelOptions) SetName(name string) *UpdateNotificationChannelOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *UpdateNotificationChannelOptions) SetType(typeVar string) *UpdateNotificationChannelOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetEndpoint : Allow user to set Endpoint
func (options *UpdateNotificationChannelOptions) SetEndpoint(endpoint string) *UpdateNotificationChannelOptions {
	options.Endpoint = core.StringPtr(endpoint)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateNotificationChannelOptions) SetDescription(description string) *UpdateNotificationChannelOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetSeverity : Allow user to set Severity
func (options *UpdateNotificationChannelOptions) SetSeverity(severity []string) *UpdateNotificationChannelOptions {
	options.Severity = severity
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *UpdateNotificationChannelOptions) SetEnabled(enabled bool) *UpdateNotificationChannelOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetAlertSource : Allow user to set AlertSource
func (options *UpdateNotificationChannelOptions) SetAlertSource(alertSource []NotificationChannelAlertSourceItem) *UpdateNotificationChannelOptions {
	options.AlertSource = alertSource
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateNotificationChannelOptions) SetTransactionID(transactionID string) *UpdateNotificationChannelOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNotificationChannelOptions) SetHeaders(param map[string]string) *UpdateNotificationChannelOptions {
	options.Headers = param
	return options
}

// Channel : Response including channels.
type Channel struct {
	// A unique ID for the channel.
	ChannelID *string `json:"channel_id,omitempty"`

	// The name of the notification channel in the form "v1/{account_id}/notifications/channelName".
	Name *string `json:"name,omitempty"`

	// A one sentence description of this `Channel`.
	Description *string `json:"description,omitempty"`

	// Type of callback URL.
	Type *string `json:"type,omitempty"`

	// The severity of the notification.
	Severity *ChannelSeverity `json:"severity,omitempty"`

	// The callback URL which receives the notification.
	Endpoint *string `json:"endpoint,omitempty"`

	// Whether the channel is enabled. The default is disabled.
	Enabled *bool `json:"enabled,omitempty"`

	AlertSource []ChannelAlertSourceItem `json:"alert_source,omitempty"`

	Frequency *string `json:"frequency,omitempty"`
}

// Constants associated with the Channel.Type property.
// Type of callback URL.
const (
	ChannelTypeWebhookConst = "Webhook"
)

// UnmarshalChannel unmarshals an instance of Channel from the specified map of raw messages.
func UnmarshalChannel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Channel)
	err = core.UnmarshalPrimitive(m, "channel_id", &obj.ChannelID)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "severity", &obj.Severity, UnmarshalChannelSeverity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "alert_source", &obj.AlertSource, UnmarshalChannelAlertSourceItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "frequency", &obj.Frequency)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

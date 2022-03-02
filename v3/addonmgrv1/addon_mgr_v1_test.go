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

package addonmgrv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/addonmgrv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AddonMgrV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		accountID := "testString"
		It(`Instantiate service client`, func() {
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				AccountID:     core.StringPtr(accountID),
			})
			Expect(addonMgrService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
				URL:       "{BAD_URL_STRING",
				AccountID: core.StringPtr(accountID),
			})
			Expect(addonMgrService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
				URL:       "https://addonmgrv1/api",
				AccountID: core.StringPtr(accountID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(addonMgrService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{})
			Expect(addonMgrService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		accountID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADDON_MGR_URL":       "https://addonmgrv1/api",
				"ADDON_MGR_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1UsingExternalConfig(&addonmgrv1.AddonMgrV1Options{
					AccountID: core.StringPtr(accountID),
				})
				Expect(addonMgrService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := addonMgrService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != addonMgrService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(addonMgrService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(addonMgrService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1UsingExternalConfig(&addonmgrv1.AddonMgrV1Options{
					URL:       "https://testService/api",
					AccountID: core.StringPtr(accountID),
				})
				Expect(addonMgrService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := addonMgrService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != addonMgrService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(addonMgrService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(addonMgrService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1UsingExternalConfig(&addonmgrv1.AddonMgrV1Options{
					AccountID: core.StringPtr(accountID),
				})
				err := addonMgrService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := addonMgrService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != addonMgrService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(addonMgrService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(addonMgrService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADDON_MGR_URL":       "https://addonmgrv1/api",
				"ADDON_MGR_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1UsingExternalConfig(&addonmgrv1.AddonMgrV1Options{
				AccountID: core.StringPtr(accountID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(addonMgrService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADDON_MGR_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1UsingExternalConfig(&addonmgrv1.AddonMgrV1Options{
				URL:       "{BAD_URL_STRING",
				AccountID: core.StringPtr(accountID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(addonMgrService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = addonmgrv1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us-south.secadvisor.cloud.ibm.com/addonmgr"))
			Expect(err).To(BeNil())

			url, err = addonmgrv1.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us-south.secadvisor.cloud.ibm.com/addonmgr"))
			Expect(err).To(BeNil())

			url, err = addonmgrv1.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://eu-gb.secadvisor.cloud.ibm.com/addonmgr"))
			Expect(err).To(BeNil())

			url, err = addonmgrv1.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu.compliance.cloud.ibm.com/si/addonmgr"))
			Expect(err).To(BeNil())

			url, err = addonmgrv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options)`, func() {
		accountID := "testString"
		addNetworkInsightsCosDetailsV2Path := "/v2/addons/testString/network-insights/cos"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addNetworkInsightsCosDetailsV2Path))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddNetworkInsightsCosDetailsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.AddNetworkInsightsCosDetailsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")

				// Construct an instance of the AddNetworkInsightsCosDetailsV2Options model
				addNetworkInsightsCosDetailsV2OptionsModel := new(addonmgrv1.AddNetworkInsightsCosDetailsV2Options)
				addNetworkInsightsCosDetailsV2OptionsModel.RegionID = core.StringPtr("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.CosDetails = []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}
				addNetworkInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddNetworkInsightsCosDetailsV2 with error: Operation validation and request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")

				// Construct an instance of the AddNetworkInsightsCosDetailsV2Options model
				addNetworkInsightsCosDetailsV2OptionsModel := new(addonmgrv1.AddNetworkInsightsCosDetailsV2Options)
				addNetworkInsightsCosDetailsV2OptionsModel.RegionID = core.StringPtr("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.CosDetails = []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}
				addNetworkInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddNetworkInsightsCosDetailsV2Options model with no property values
				addNetworkInsightsCosDetailsV2OptionsModelNew := new(addonmgrv1.AddNetworkInsightsCosDetailsV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = addonMgrService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options)`, func() {
		accountID := "testString"
		deleteNetworkInsightsCosDetailsV2Path := "/v2/addons/testString/network-insights/cos"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNetworkInsightsCosDetailsV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteNetworkInsightsCosDetailsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.DeleteNetworkInsightsCosDetailsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNetworkInsightsCosDetailsV2Options model
				deleteNetworkInsightsCosDetailsV2OptionsModel := new(addonmgrv1.DeleteNetworkInsightsCosDetailsV2Options)
				deleteNetworkInsightsCosDetailsV2OptionsModel.Ids = []string{"testString"}
				deleteNetworkInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				deleteNetworkInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNetworkInsightsCosDetailsV2 with error: Operation request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the DeleteNetworkInsightsCosDetailsV2Options model
				deleteNetworkInsightsCosDetailsV2OptionsModel := new(addonmgrv1.DeleteNetworkInsightsCosDetailsV2Options)
				deleteNetworkInsightsCosDetailsV2OptionsModel.Ids = []string{"testString"}
				deleteNetworkInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				deleteNetworkInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options)`, func() {
		accountID := "testString"
		addActivityInsightsCosDetailsV2Path := "/v2/addons/testString/activity-insights/cos"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addActivityInsightsCosDetailsV2Path))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddActivityInsightsCosDetailsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.AddActivityInsightsCosDetailsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")

				// Construct an instance of the AddActivityInsightsCosDetailsV2Options model
				addActivityInsightsCosDetailsV2OptionsModel := new(addonmgrv1.AddActivityInsightsCosDetailsV2Options)
				addActivityInsightsCosDetailsV2OptionsModel.RegionID = core.StringPtr("testString")
				addActivityInsightsCosDetailsV2OptionsModel.CosDetails = []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}
				addActivityInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				addActivityInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddActivityInsightsCosDetailsV2 with error: Operation validation and request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")

				// Construct an instance of the AddActivityInsightsCosDetailsV2Options model
				addActivityInsightsCosDetailsV2OptionsModel := new(addonmgrv1.AddActivityInsightsCosDetailsV2Options)
				addActivityInsightsCosDetailsV2OptionsModel.RegionID = core.StringPtr("testString")
				addActivityInsightsCosDetailsV2OptionsModel.CosDetails = []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}
				addActivityInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				addActivityInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddActivityInsightsCosDetailsV2Options model with no property values
				addActivityInsightsCosDetailsV2OptionsModelNew := new(addonmgrv1.AddActivityInsightsCosDetailsV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = addonMgrService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options)`, func() {
		accountID := "testString"
		deleteActivityInsightsCosDetailsV2Path := "/v2/addons/testString/activity-insights/cos"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteActivityInsightsCosDetailsV2Path))
					Expect(req.Method).To(Equal("DELETE"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteActivityInsightsCosDetailsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.DeleteActivityInsightsCosDetailsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteActivityInsightsCosDetailsV2Options model
				deleteActivityInsightsCosDetailsV2OptionsModel := new(addonmgrv1.DeleteActivityInsightsCosDetailsV2Options)
				deleteActivityInsightsCosDetailsV2OptionsModel.Ids = []string{"testString"}
				deleteActivityInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				deleteActivityInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteActivityInsightsCosDetailsV2 with error: Operation request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the DeleteActivityInsightsCosDetailsV2Options model
				deleteActivityInsightsCosDetailsV2OptionsModel := new(addonmgrv1.DeleteActivityInsightsCosDetailsV2Options)
				deleteActivityInsightsCosDetailsV2OptionsModel.Ids = []string{"testString"}
				deleteActivityInsightsCosDetailsV2OptionsModel.TransactionID = core.StringPtr("testString")
				deleteActivityInsightsCosDetailsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DisableInsightsV2(disableInsightsV2Options *DisableInsightsV2Options)`, func() {
		accountID := "testString"
		disableInsightsV2Path := "/v2/addons/testString/disable"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(disableInsightsV2Path))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DisableInsightsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.DisableInsightsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DisableInsightsV2Options model
				disableInsightsV2OptionsModel := new(addonmgrv1.DisableInsightsV2Options)
				disableInsightsV2OptionsModel.RegionID = core.StringPtr("testString")
				disableInsightsV2OptionsModel.NetworkInsights = core.BoolPtr(true)
				disableInsightsV2OptionsModel.ActivityInsights = core.BoolPtr(true)
				disableInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				disableInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.DisableInsightsV2(disableInsightsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DisableInsightsV2 with error: Operation validation and request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the DisableInsightsV2Options model
				disableInsightsV2OptionsModel := new(addonmgrv1.DisableInsightsV2Options)
				disableInsightsV2OptionsModel.RegionID = core.StringPtr("testString")
				disableInsightsV2OptionsModel.NetworkInsights = core.BoolPtr(true)
				disableInsightsV2OptionsModel.ActivityInsights = core.BoolPtr(true)
				disableInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				disableInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.DisableInsightsV2(disableInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DisableInsightsV2Options model with no property values
				disableInsightsV2OptionsModelNew := new(addonmgrv1.DisableInsightsV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = addonMgrService.DisableInsightsV2(disableInsightsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EnableInsightsV2(enableInsightsV2Options *EnableInsightsV2Options)`, func() {
		accountID := "testString"
		enableInsightsV2Path := "/v2/addons/testString/enable"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(enableInsightsV2Path))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke EnableInsightsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.EnableInsightsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the EnableInsightsV2Options model
				enableInsightsV2OptionsModel := new(addonmgrv1.EnableInsightsV2Options)
				enableInsightsV2OptionsModel.RegionID = core.StringPtr("testString")
				enableInsightsV2OptionsModel.NetworkInsights = core.BoolPtr(true)
				enableInsightsV2OptionsModel.ActivityInsights = core.BoolPtr(true)
				enableInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				enableInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.EnableInsightsV2(enableInsightsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke EnableInsightsV2 with error: Operation validation and request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the EnableInsightsV2Options model
				enableInsightsV2OptionsModel := new(addonmgrv1.EnableInsightsV2Options)
				enableInsightsV2OptionsModel.RegionID = core.StringPtr("testString")
				enableInsightsV2OptionsModel.NetworkInsights = core.BoolPtr(true)
				enableInsightsV2OptionsModel.ActivityInsights = core.BoolPtr(true)
				enableInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				enableInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.EnableInsightsV2(enableInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the EnableInsightsV2Options model with no property values
				enableInsightsV2OptionsModelNew := new(addonmgrv1.EnableInsightsV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = addonMgrService.EnableInsightsV2(enableInsightsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options) - Operation response error`, func() {
		accountID := "testString"
		getSupportedInsightsV2Path := "/v2/addons/testString/insights"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedInsightsV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSupportedInsightsV2 with error: Operation response processing error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := new(addonmgrv1.GetSupportedInsightsV2Options)
				getSupportedInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				getSupportedInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				addonMgrService.EnableRetries(0, 0)
				result, response, operationErr = addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options)`, func() {
		accountID := "testString"
		getSupportedInsightsV2Path := "/v2/addons/testString/insights"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedInsightsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": ["network-insights"]}`)
				}))
			})
			It(`Invoke GetSupportedInsightsV2 successfully with retries`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())
				addonMgrService.EnableRetries(0, 0)

				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := new(addonmgrv1.GetSupportedInsightsV2Options)
				getSupportedInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				getSupportedInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := addonMgrService.GetSupportedInsightsV2WithContext(ctx, getSupportedInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				addonMgrService.DisableRetries()
				result, response, operationErr := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = addonMgrService.GetSupportedInsightsV2WithContext(ctx, getSupportedInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedInsightsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": ["network-insights"]}`)
				}))
			})
			It(`Invoke GetSupportedInsightsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := addonMgrService.GetSupportedInsightsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := new(addonmgrv1.GetSupportedInsightsV2Options)
				getSupportedInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				getSupportedInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSupportedInsightsV2 with error: Operation request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := new(addonmgrv1.GetSupportedInsightsV2Options)
				getSupportedInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				getSupportedInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSupportedInsightsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := new(addonmgrv1.GetSupportedInsightsV2Options)
				getSupportedInsightsV2OptionsModel.TransactionID = core.StringPtr("testString")
				getSupportedInsightsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestAiFindingsV2(testAiFindingsV2Options *TestAiFindingsV2Options)`, func() {
		accountID := "testString"
		testAiFindingsV2Path := "/v2/addons/testString/activity-insights/test-ai-findings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testAiFindingsV2Path))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke TestAiFindingsV2 successfully`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := addonMgrService.TestAiFindingsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the TestAiFindingsV2Options model
				testAiFindingsV2OptionsModel := new(addonmgrv1.TestAiFindingsV2Options)
				testAiFindingsV2OptionsModel.RegionID = core.StringPtr("testString")
				testAiFindingsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = addonMgrService.TestAiFindingsV2(testAiFindingsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke TestAiFindingsV2 with error: Operation validation and request error`, func() {
				addonMgrService, serviceErr := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					AccountID:     core.StringPtr(accountID),
				})
				Expect(serviceErr).To(BeNil())
				Expect(addonMgrService).ToNot(BeNil())

				// Construct an instance of the TestAiFindingsV2Options model
				testAiFindingsV2OptionsModel := new(addonmgrv1.TestAiFindingsV2Options)
				testAiFindingsV2OptionsModel.RegionID = core.StringPtr("testString")
				testAiFindingsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := addonMgrService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := addonMgrService.TestAiFindingsV2(testAiFindingsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the TestAiFindingsV2Options model with no property values
				testAiFindingsV2OptionsModelNew := new(addonmgrv1.TestAiFindingsV2Options)
				// Invoke operation with invalid model (negative test)
				response, operationErr = addonMgrService.TestAiFindingsV2(testAiFindingsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			accountID := "testString"
			addonMgrService, _ := addonmgrv1.NewAddonMgrV1(&addonmgrv1.AddonMgrV1Options{
				URL:           "http://addonmgrv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				AccountID:     core.StringPtr(accountID),
			})
			It(`Invoke NewAddActivityInsightsCosDetailsV2Options successfully`, func() {
				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				Expect(cosDetailsV2CosDetailsItemModel).ToNot(BeNil())
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")
				Expect(cosDetailsV2CosDetailsItemModel.CosInstance).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.BucketName).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.Type).To(Equal(core.StringPtr("network-insights")))
				Expect(cosDetailsV2CosDetailsItemModel.CosBucketURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddActivityInsightsCosDetailsV2Options model
				addActivityInsightsCosDetailsV2OptionsRegionID := "testString"
				addActivityInsightsCosDetailsV2OptionsCosDetails := []addonmgrv1.CosDetailsV2CosDetailsItem{}
				addActivityInsightsCosDetailsV2OptionsModel := addonMgrService.NewAddActivityInsightsCosDetailsV2Options(addActivityInsightsCosDetailsV2OptionsRegionID, addActivityInsightsCosDetailsV2OptionsCosDetails)
				addActivityInsightsCosDetailsV2OptionsModel.SetRegionID("testString")
				addActivityInsightsCosDetailsV2OptionsModel.SetCosDetails([]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel})
				addActivityInsightsCosDetailsV2OptionsModel.SetTransactionID("testString")
				addActivityInsightsCosDetailsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addActivityInsightsCosDetailsV2OptionsModel).ToNot(BeNil())
				Expect(addActivityInsightsCosDetailsV2OptionsModel.RegionID).To(Equal(core.StringPtr("testString")))
				Expect(addActivityInsightsCosDetailsV2OptionsModel.CosDetails).To(Equal([]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}))
				Expect(addActivityInsightsCosDetailsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addActivityInsightsCosDetailsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddNetworkInsightsCosDetailsV2Options successfully`, func() {
				// Construct an instance of the CosDetailsV2CosDetailsItem model
				cosDetailsV2CosDetailsItemModel := new(addonmgrv1.CosDetailsV2CosDetailsItem)
				Expect(cosDetailsV2CosDetailsItemModel).ToNot(BeNil())
				cosDetailsV2CosDetailsItemModel.CosInstance = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.BucketName = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Description = core.StringPtr("testString")
				cosDetailsV2CosDetailsItemModel.Type = core.StringPtr("network-insights")
				cosDetailsV2CosDetailsItemModel.CosBucketURL = core.StringPtr("testString")
				Expect(cosDetailsV2CosDetailsItemModel.CosInstance).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.BucketName).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(cosDetailsV2CosDetailsItemModel.Type).To(Equal(core.StringPtr("network-insights")))
				Expect(cosDetailsV2CosDetailsItemModel.CosBucketURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddNetworkInsightsCosDetailsV2Options model
				addNetworkInsightsCosDetailsV2OptionsRegionID := "testString"
				addNetworkInsightsCosDetailsV2OptionsCosDetails := []addonmgrv1.CosDetailsV2CosDetailsItem{}
				addNetworkInsightsCosDetailsV2OptionsModel := addonMgrService.NewAddNetworkInsightsCosDetailsV2Options(addNetworkInsightsCosDetailsV2OptionsRegionID, addNetworkInsightsCosDetailsV2OptionsCosDetails)
				addNetworkInsightsCosDetailsV2OptionsModel.SetRegionID("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.SetCosDetails([]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel})
				addNetworkInsightsCosDetailsV2OptionsModel.SetTransactionID("testString")
				addNetworkInsightsCosDetailsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addNetworkInsightsCosDetailsV2OptionsModel).ToNot(BeNil())
				Expect(addNetworkInsightsCosDetailsV2OptionsModel.RegionID).To(Equal(core.StringPtr("testString")))
				Expect(addNetworkInsightsCosDetailsV2OptionsModel.CosDetails).To(Equal([]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel}))
				Expect(addNetworkInsightsCosDetailsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addNetworkInsightsCosDetailsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCosDetailsV2CosDetailsItem successfully`, func() {
				cosInstance := "testString"
				bucketName := "testString"
				description := "testString"
				typeVar := "network-insights"
				cosBucketURL := "testString"
				_model, err := addonMgrService.NewCosDetailsV2CosDetailsItem(cosInstance, bucketName, description, typeVar, cosBucketURL)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteActivityInsightsCosDetailsV2Options successfully`, func() {
				// Construct an instance of the DeleteActivityInsightsCosDetailsV2Options model
				deleteActivityInsightsCosDetailsV2OptionsModel := addonMgrService.NewDeleteActivityInsightsCosDetailsV2Options()
				deleteActivityInsightsCosDetailsV2OptionsModel.SetIds([]string{"testString"})
				deleteActivityInsightsCosDetailsV2OptionsModel.SetTransactionID("testString")
				deleteActivityInsightsCosDetailsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteActivityInsightsCosDetailsV2OptionsModel).ToNot(BeNil())
				Expect(deleteActivityInsightsCosDetailsV2OptionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(deleteActivityInsightsCosDetailsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteActivityInsightsCosDetailsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNetworkInsightsCosDetailsV2Options successfully`, func() {
				// Construct an instance of the DeleteNetworkInsightsCosDetailsV2Options model
				deleteNetworkInsightsCosDetailsV2OptionsModel := addonMgrService.NewDeleteNetworkInsightsCosDetailsV2Options()
				deleteNetworkInsightsCosDetailsV2OptionsModel.SetIds([]string{"testString"})
				deleteNetworkInsightsCosDetailsV2OptionsModel.SetTransactionID("testString")
				deleteNetworkInsightsCosDetailsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNetworkInsightsCosDetailsV2OptionsModel).ToNot(BeNil())
				Expect(deleteNetworkInsightsCosDetailsV2OptionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(deleteNetworkInsightsCosDetailsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNetworkInsightsCosDetailsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDisableInsightsV2Options successfully`, func() {
				// Construct an instance of the DisableInsightsV2Options model
				disableInsightsV2OptionsRegionID := "testString"
				disableInsightsV2OptionsModel := addonMgrService.NewDisableInsightsV2Options(disableInsightsV2OptionsRegionID)
				disableInsightsV2OptionsModel.SetRegionID("testString")
				disableInsightsV2OptionsModel.SetNetworkInsights(true)
				disableInsightsV2OptionsModel.SetActivityInsights(true)
				disableInsightsV2OptionsModel.SetTransactionID("testString")
				disableInsightsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(disableInsightsV2OptionsModel).ToNot(BeNil())
				Expect(disableInsightsV2OptionsModel.RegionID).To(Equal(core.StringPtr("testString")))
				Expect(disableInsightsV2OptionsModel.NetworkInsights).To(Equal(core.BoolPtr(true)))
				Expect(disableInsightsV2OptionsModel.ActivityInsights).To(Equal(core.BoolPtr(true)))
				Expect(disableInsightsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(disableInsightsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnableInsightsV2Options successfully`, func() {
				// Construct an instance of the EnableInsightsV2Options model
				enableInsightsV2OptionsRegionID := "testString"
				enableInsightsV2OptionsModel := addonMgrService.NewEnableInsightsV2Options(enableInsightsV2OptionsRegionID)
				enableInsightsV2OptionsModel.SetRegionID("testString")
				enableInsightsV2OptionsModel.SetNetworkInsights(true)
				enableInsightsV2OptionsModel.SetActivityInsights(true)
				enableInsightsV2OptionsModel.SetTransactionID("testString")
				enableInsightsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(enableInsightsV2OptionsModel).ToNot(BeNil())
				Expect(enableInsightsV2OptionsModel.RegionID).To(Equal(core.StringPtr("testString")))
				Expect(enableInsightsV2OptionsModel.NetworkInsights).To(Equal(core.BoolPtr(true)))
				Expect(enableInsightsV2OptionsModel.ActivityInsights).To(Equal(core.BoolPtr(true)))
				Expect(enableInsightsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(enableInsightsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSupportedInsightsV2Options successfully`, func() {
				// Construct an instance of the GetSupportedInsightsV2Options model
				getSupportedInsightsV2OptionsModel := addonMgrService.NewGetSupportedInsightsV2Options()
				getSupportedInsightsV2OptionsModel.SetTransactionID("testString")
				getSupportedInsightsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSupportedInsightsV2OptionsModel).ToNot(BeNil())
				Expect(getSupportedInsightsV2OptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getSupportedInsightsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTestAiFindingsV2Options successfully`, func() {
				// Construct an instance of the TestAiFindingsV2Options model
				testAiFindingsV2OptionsRegionID := "testString"
				testAiFindingsV2OptionsModel := addonMgrService.NewTestAiFindingsV2Options(testAiFindingsV2OptionsRegionID)
				testAiFindingsV2OptionsModel.SetRegionID("testString")
				testAiFindingsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testAiFindingsV2OptionsModel).ToNot(BeNil())
				Expect(testAiFindingsV2OptionsModel.RegionID).To(Equal(core.StringPtr("testString")))
				Expect(testAiFindingsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}

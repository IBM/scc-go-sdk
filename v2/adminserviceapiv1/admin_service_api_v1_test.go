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

package adminserviceapiv1_test

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
	"github.com/IBM/scc-go-sdk/v2/adminserviceapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AdminServiceApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(adminServiceApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(adminServiceApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
				URL: "https://adminserviceapiv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(adminServiceApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_API_URL":       "https://adminserviceapiv1/api",
				"ADMIN_SERVICE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(&adminserviceapiv1.AdminServiceApiV1Options{})
				Expect(adminServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(&adminserviceapiv1.AdminServiceApiV1Options{
					URL: "https://testService/api",
				})
				Expect(adminServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(&adminserviceapiv1.AdminServiceApiV1Options{})
				err := adminServiceApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminServiceApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_API_URL":       "https://adminserviceapiv1/api",
				"ADMIN_SERVICE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(&adminserviceapiv1.AdminServiceApiV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(adminServiceApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMIN_SERVICE_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(&adminserviceapiv1.AdminServiceApiV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(adminServiceApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = adminserviceapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/admin/v1/accounts/testString/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminserviceapiv1.GetSettingsOptions)
				getSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/admin/v1/accounts/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "us"}}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())
				adminServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminserviceapiv1.GetSettingsOptions)
				getSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceApiService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceApiService.DisableRetries()
				result, response, operationErr := adminServiceApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceApiService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "us"}}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceApiService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminserviceapiv1.GetSettingsOptions)
				getSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation validation and request error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminserviceapiv1.GetSettingsOptions)
				getSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSettingsOptions model with no property values
				getSettingsOptionsModelNew := new(adminserviceapiv1.GetSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceApiService.GetSettings(getSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke GetSettings successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(adminserviceapiv1.GetSettingsOptions)
				getSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceApiService.GetSettings(getSettingsOptionsModel)
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
	Describe(`PatchAccountSettings(patchAccountSettingsOptions *PatchAccountSettingsOptions) - Operation response error`, func() {
		patchAccountSettingsPath := "/admin/v1/accounts/testString/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchAccountSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PatchAccountSettings with error: Operation response processing error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				locationIdModel.ID = core.StringPtr("eu")

				// Construct an instance of the PatchAccountSettingsOptions model
				patchAccountSettingsOptionsModel := new(adminserviceapiv1.PatchAccountSettingsOptions)
				patchAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				patchAccountSettingsOptionsModel.Location = locationIdModel
				patchAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchAccountSettings(patchAccountSettingsOptions *PatchAccountSettingsOptions)`, func() {
		patchAccountSettingsPath := "/admin/v1/accounts/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchAccountSettingsPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "us"}}`)
				}))
			})
			It(`Invoke PatchAccountSettings successfully with retries`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())
				adminServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				locationIdModel.ID = core.StringPtr("eu")

				// Construct an instance of the PatchAccountSettingsOptions model
				patchAccountSettingsOptionsModel := new(adminserviceapiv1.PatchAccountSettingsOptions)
				patchAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				patchAccountSettingsOptionsModel.Location = locationIdModel
				patchAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceApiService.PatchAccountSettingsWithContext(ctx, patchAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceApiService.DisableRetries()
				result, response, operationErr := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceApiService.PatchAccountSettingsWithContext(ctx, patchAccountSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(patchAccountSettingsPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": {"id": "us"}}`)
				}))
			})
			It(`Invoke PatchAccountSettings successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceApiService.PatchAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				locationIdModel.ID = core.StringPtr("eu")

				// Construct an instance of the PatchAccountSettingsOptions model
				patchAccountSettingsOptionsModel := new(adminserviceapiv1.PatchAccountSettingsOptions)
				patchAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				patchAccountSettingsOptionsModel.Location = locationIdModel
				patchAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PatchAccountSettings with error: Operation validation and request error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				locationIdModel.ID = core.StringPtr("eu")

				// Construct an instance of the PatchAccountSettingsOptions model
				patchAccountSettingsOptionsModel := new(adminserviceapiv1.PatchAccountSettingsOptions)
				patchAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				patchAccountSettingsOptionsModel.Location = locationIdModel
				patchAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PatchAccountSettingsOptions model with no property values
				patchAccountSettingsOptionsModelNew := new(adminserviceapiv1.PatchAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke PatchAccountSettings successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				locationIdModel.ID = core.StringPtr("eu")

				// Construct an instance of the PatchAccountSettingsOptions model
				patchAccountSettingsOptionsModel := new(adminserviceapiv1.PatchAccountSettingsOptions)
				patchAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				patchAccountSettingsOptionsModel.Location = locationIdModel
				patchAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptionsModel)
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
	Describe(`ListLocations(listLocationsOptions *ListLocationsOptions) - Operation response error`, func() {
		listLocationsPath := "/admin/v1/locations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLocations with error: Operation response processing error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := new(adminserviceapiv1.ListLocationsOptions)
				listLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceApiService.ListLocations(listLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceApiService.ListLocations(listLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLocations(listLocationsOptions *ListLocationsOptions)`, func() {
		listLocationsPath := "/admin/v1/locations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"id": "us", "main_endpoint_url": "MainEndpointURL", "governance_endpoint_url": "GovernanceEndpointURL", "results_endpoint_url": "ResultsEndpointURL", "compliance_endpoint_url": "ComplianceEndpointURL", "analytics_endpoint_url": "AnalyticsEndpointURL", "si_endpoint_url": "SiEndpointURL", "regions": [{"id": "ID"}]}]}`)
				}))
			})
			It(`Invoke ListLocations successfully with retries`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())
				adminServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := new(adminserviceapiv1.ListLocationsOptions)
				listLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceApiService.ListLocationsWithContext(ctx, listLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceApiService.DisableRetries()
				result, response, operationErr := adminServiceApiService.ListLocations(listLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceApiService.ListLocationsWithContext(ctx, listLocationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"id": "us", "main_endpoint_url": "MainEndpointURL", "governance_endpoint_url": "GovernanceEndpointURL", "results_endpoint_url": "ResultsEndpointURL", "compliance_endpoint_url": "ComplianceEndpointURL", "analytics_endpoint_url": "AnalyticsEndpointURL", "si_endpoint_url": "SiEndpointURL", "regions": [{"id": "ID"}]}]}`)
				}))
			})
			It(`Invoke ListLocations successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceApiService.ListLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := new(adminserviceapiv1.ListLocationsOptions)
				listLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceApiService.ListLocations(listLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLocations with error: Operation request error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := new(adminserviceapiv1.ListLocationsOptions)
				listLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceApiService.ListLocations(listLocationsOptionsModel)
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
			It(`Invoke ListLocations successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := new(adminserviceapiv1.ListLocationsOptions)
				listLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceApiService.ListLocations(listLocationsOptionsModel)
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
	Describe(`GetLocation(getLocationOptions *GetLocationOptions) - Operation response error`, func() {
		getLocationPath := "/admin/v1/locations/us"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLocation with error: Operation response processing error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetLocationOptions model
				getLocationOptionsModel := new(adminserviceapiv1.GetLocationOptions)
				getLocationOptionsModel.LocationID = core.StringPtr("us")
				getLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminServiceApiService.GetLocation(getLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = adminServiceApiService.GetLocation(getLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLocation(getLocationOptions *GetLocationOptions)`, func() {
		getLocationPath := "/admin/v1/locations/us"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "us", "main_endpoint_url": "MainEndpointURL", "governance_endpoint_url": "GovernanceEndpointURL", "results_endpoint_url": "ResultsEndpointURL", "compliance_endpoint_url": "ComplianceEndpointURL", "analytics_endpoint_url": "AnalyticsEndpointURL", "si_endpoint_url": "SiEndpointURL", "regions": [{"id": "ID"}]}`)
				}))
			})
			It(`Invoke GetLocation successfully with retries`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())
				adminServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLocationOptions model
				getLocationOptionsModel := new(adminserviceapiv1.GetLocationOptions)
				getLocationOptionsModel.LocationID = core.StringPtr("us")
				getLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminServiceApiService.GetLocationWithContext(ctx, getLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminServiceApiService.DisableRetries()
				result, response, operationErr := adminServiceApiService.GetLocation(getLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminServiceApiService.GetLocationWithContext(ctx, getLocationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "us", "main_endpoint_url": "MainEndpointURL", "governance_endpoint_url": "GovernanceEndpointURL", "results_endpoint_url": "ResultsEndpointURL", "compliance_endpoint_url": "ComplianceEndpointURL", "analytics_endpoint_url": "AnalyticsEndpointURL", "si_endpoint_url": "SiEndpointURL", "regions": [{"id": "ID"}]}`)
				}))
			})
			It(`Invoke GetLocation successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminServiceApiService.GetLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLocationOptions model
				getLocationOptionsModel := new(adminserviceapiv1.GetLocationOptions)
				getLocationOptionsModel.LocationID = core.StringPtr("us")
				getLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminServiceApiService.GetLocation(getLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLocation with error: Operation validation and request error`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetLocationOptions model
				getLocationOptionsModel := new(adminserviceapiv1.GetLocationOptions)
				getLocationOptionsModel.LocationID = core.StringPtr("us")
				getLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminServiceApiService.GetLocation(getLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLocationOptions model with no property values
				getLocationOptionsModelNew := new(adminserviceapiv1.GetLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminServiceApiService.GetLocation(getLocationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke GetLocation successfully`, func() {
				adminServiceApiService, serviceErr := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetLocationOptions model
				getLocationOptionsModel := new(adminserviceapiv1.GetLocationOptions)
				getLocationOptionsModel.LocationID = core.StringPtr("us")
				getLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminServiceApiService.GetLocation(getLocationOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			adminServiceApiService, _ := adminserviceapiv1.NewAdminServiceApiV1(&adminserviceapiv1.AdminServiceApiV1Options{
				URL:           "http://adminserviceapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAccountSettings successfully`, func() {
				var location *adminserviceapiv1.LocationID = nil
				_, err := adminServiceApiService.NewAccountSettings(location)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGetLocationOptions successfully`, func() {
				// Construct an instance of the GetLocationOptions model
				locationID := "us"
				getLocationOptionsModel := adminServiceApiService.NewGetLocationOptions(locationID)
				getLocationOptionsModel.SetLocationID("us")
				getLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLocationOptionsModel).ToNot(BeNil())
				Expect(getLocationOptionsModel.LocationID).To(Equal(core.StringPtr("us")))
				Expect(getLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				accountID := "testString"
				getSettingsOptionsModel := adminServiceApiService.NewGetSettingsOptions(accountID)
				getSettingsOptionsModel.SetAccountID("testString")
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLocationsOptions successfully`, func() {
				// Construct an instance of the ListLocationsOptions model
				listLocationsOptionsModel := adminServiceApiService.NewListLocationsOptions()
				listLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLocationsOptionsModel).ToNot(BeNil())
				Expect(listLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLocationID successfully`, func() {
				id := "us"
				_model, err := adminServiceApiService.NewLocationID(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPatchAccountSettingsOptions successfully`, func() {
				// Construct an instance of the LocationID model
				locationIdModel := new(adminserviceapiv1.LocationID)
				Expect(locationIdModel).ToNot(BeNil())
				locationIdModel.ID = core.StringPtr("eu")
				Expect(locationIdModel.ID).To(Equal(core.StringPtr("eu")))

				// Construct an instance of the PatchAccountSettingsOptions model
				accountID := "testString"
				var patchAccountSettingsOptionsLocation *adminserviceapiv1.LocationID = nil
				patchAccountSettingsOptionsModel := adminServiceApiService.NewPatchAccountSettingsOptions(accountID, patchAccountSettingsOptionsLocation)
				patchAccountSettingsOptionsModel.SetAccountID("testString")
				patchAccountSettingsOptionsModel.SetLocation(locationIdModel)
				patchAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(patchAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(patchAccountSettingsOptionsModel.Location).To(Equal(locationIdModel))
				Expect(patchAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

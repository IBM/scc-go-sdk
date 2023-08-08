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

package securityandcompliancecenterapiv3_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SecurityAndComplianceCenterApiV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(securityAndComplianceCenterApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL: "https://securityandcompliancecenterapiv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(securityAndComplianceCenterApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_AND_COMPLIANCE_CENTER_API_URL": "https://securityandcompliancecenterapiv3/api",
				"SECURITY_AND_COMPLIANCE_CENTER_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				})
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := securityAndComplianceCenterApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityAndComplianceCenterApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityAndComplianceCenterApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityAndComplianceCenterApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL: "https://testService/api",
				})
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityAndComplianceCenterApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityAndComplianceCenterApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityAndComplianceCenterApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityAndComplianceCenterApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				})
				err := securityAndComplianceCenterApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityAndComplianceCenterApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityAndComplianceCenterApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityAndComplianceCenterApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityAndComplianceCenterApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_AND_COMPLIANCE_CENTER_API_URL": "https://securityandcompliancecenterapiv3/api",
				"SECURITY_AND_COMPLIANCE_CENTER_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityAndComplianceCenterApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_AND_COMPLIANCE_CENTER_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3UsingExternalConfig(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityAndComplianceCenterApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = securityandcompliancecenterapiv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := securityandcompliancecenterapiv3.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.compliance.cloud.ibm.com/instances/instance_id/v3"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := securityandcompliancecenterapiv3.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(securityandcompliancecenterapiv3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
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
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::", "updated_on": "2019-01-01T12:00:00.000Z", "source_id": "crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::", "source_description": "This source is used for integration with IBM Cloud Security and Compliance Center.", "source_name": "compliance"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(securityandcompliancecenterapiv3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::", "updated_on": "2019-01-01T12:00:00.000Z", "source_id": "crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::", "source_description": "This source is used for integration with IBM Cloud Security and Compliance Center.", "source_name": "compliance"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(securityandcompliancecenterapiv3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(securityandcompliancecenterapiv3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
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
			It(`Invoke GetSettings successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(securityandcompliancecenterapiv3.GetSettingsOptions)
				getSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetSettings(getSettingsOptionsModel)
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
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) - Operation response error`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSettings with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(securityandcompliancecenterapiv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.EventNotifications = eventNotificationsModel
				updateSettingsOptionsModel.ObjectStorage = objectStorageModel
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::", "updated_on": "2019-01-01T12:00:00.000Z", "source_id": "crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::", "source_description": "This source is used for integration with IBM Cloud Security and Compliance Center.", "source_name": "compliance"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateSettings successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(securityandcompliancecenterapiv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.EventNotifications = eventNotificationsModel
				updateSettingsOptionsModel.ObjectStorage = objectStorageModel
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event_notifications": {"instance_crn": "crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::", "updated_on": "2019-01-01T12:00:00.000Z", "source_id": "crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::", "source_description": "This source is used for integration with IBM Cloud Security and Compliance Center.", "source_name": "compliance"}, "object_storage": {"instance_crn": "InstanceCrn", "bucket": "Bucket", "bucket_location": "BucketLocation", "bucket_endpoint": "BucketEndpoint", "updated_on": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateSettings successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(securityandcompliancecenterapiv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.EventNotifications = eventNotificationsModel
				updateSettingsOptionsModel.ObjectStorage = objectStorageModel
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSettings with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(securityandcompliancecenterapiv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.EventNotifications = eventNotificationsModel
				updateSettingsOptionsModel.ObjectStorage = objectStorageModel
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
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
			It(`Invoke UpdateSettings successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(securityandcompliancecenterapiv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.EventNotifications = eventNotificationsModel
				updateSettingsOptionsModel.ObjectStorage = objectStorageModel
				updateSettingsOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.XRequestID = core.StringPtr("testString")
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateSettings(updateSettingsOptionsModel)
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
	Describe(`PostTestEvent(postTestEventOptions *PostTestEventOptions) - Operation response error`, func() {
		postTestEventPath := "/test_event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostTestEvent with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(securityandcompliancecenterapiv3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostTestEvent(postTestEventOptions *PostTestEventOptions)`, func() {
		postTestEventPath := "/test_event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostTestEvent successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(securityandcompliancecenterapiv3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.PostTestEventWithContext(ctx, postTestEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.PostTestEventWithContext(ctx, postTestEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postTestEventPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"success": false}`)
				}))
			})
			It(`Invoke PostTestEvent successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.PostTestEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(securityandcompliancecenterapiv3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostTestEvent with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(securityandcompliancecenterapiv3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke PostTestEvent successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := new(securityandcompliancecenterapiv3.PostTestEventOptions)
				postTestEventOptionsModel.XCorrelationID = core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.XRequestID = core.StringPtr("testString")
				postTestEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.PostTestEvent(postTestEventOptionsModel)
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
	Describe(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions) - Operation response error`, func() {
		listControlLibrariesPath := "/control_libraries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListControlLibraries with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(securityandcompliancecenterapiv3.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("custom")
				listControlLibrariesOptionsModel.Start = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListControlLibraries(listControlLibrariesOptions *ListControlLibrariesOptions)`, func() {
		listControlLibrariesPath := "/control_libraries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(securityandcompliancecenterapiv3.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("custom")
				listControlLibrariesOptionsModel.Start = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListControlLibraries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(securityandcompliancecenterapiv3.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("custom")
				listControlLibrariesOptionsModel.Start = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListControlLibraries with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(securityandcompliancecenterapiv3.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("custom")
				listControlLibrariesOptionsModel.Start = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
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
			It(`Invoke ListControlLibraries successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(securityandcompliancecenterapiv3.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("custom")
				listControlLibrariesOptionsModel.Start = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ControlLibraryCollection)
				nextObject := new(securityandcompliancecenterapiv3.PaginatedCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ControlLibraryCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"control_libraries":[{"id":"ID","account_id":"AccountID","control_library_name":"ControlLibraryName","control_library_description":"ControlLibraryDescription","control_library_type":"ControlLibraryType","created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","version_group_label":"VersionGroupLabel","control_library_version":"ControlLibraryVersion","latest":true,"controls_count":13}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"control_libraries":[{"id":"ID","account_id":"AccountID","control_library_name":"ControlLibraryName","control_library_description":"ControlLibraryDescription","control_library_type":"ControlLibraryType","created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","version_group_label":"VersionGroupLabel","control_library_version":"ControlLibraryVersion","latest":true,"controls_count":13}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ControlLibrariesPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listControlLibrariesOptionsModel := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(50)),
					ControlLibraryType: core.StringPtr("custom"),
				}

				pager, err := securityAndComplianceCenterApiService.NewControlLibrariesPager(listControlLibrariesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.ControlLibraryItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ControlLibrariesPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listControlLibrariesOptionsModel := &securityandcompliancecenterapiv3.ListControlLibrariesOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(50)),
					ControlLibraryType: core.StringPtr("custom"),
				}

				pager, err := securityAndComplianceCenterApiService.NewControlLibrariesPager(listControlLibrariesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions) - Operation response error`, func() {
		createCustomControlLibraryPath := "/control_libraries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomControlLibraryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.0.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomControlLibrary(createCustomControlLibraryOptions *CreateCustomControlLibraryOptions)`, func() {
		createCustomControlLibraryPath := "/control_libraries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomControlLibraryPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.0.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCustomControlLibraryPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.0.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomControlLibrary with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.0.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomControlLibraryOptions model with no property values
				createCustomControlLibraryOptionsModelNew := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.0.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
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
	Describe(`DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions) - Operation response error`, func() {
		deleteCustomControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControlLibraryPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCustomControlLibrary with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomControlLibrary(deleteCustomControlLibraryOptions *DeleteCustomControlLibraryOptions)`, func() {
		deleteCustomControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControlLibraryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteCustomControlLibrary successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibraryWithContext(ctx, deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.DeleteCustomControlLibraryWithContext(ctx, deleteCustomControlLibraryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControlLibraryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomControlLibrary with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomControlLibraryOptions model with no property values
				deleteCustomControlLibraryOptionsModelNew := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModelNew)
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
			It(`Invoke DeleteCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
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
	Describe(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions) - Operation response error`, func() {
		getControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getControlLibraryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetControlLibrary with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetControlLibrary(getControlLibraryOptions *GetControlLibraryOptions)`, func() {
		getControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getControlLibraryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getControlLibraryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetControlLibrary with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetControlLibraryOptions model with no property values
				getControlLibraryOptionsModelNew := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModelNew)
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
			It(`Invoke GetControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetControlLibrary(getControlLibraryOptionsModel)
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
	Describe(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions) - Operation response error`, func() {
		replaceCustomControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceCustomControlLibraryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceCustomControlLibrary(replaceCustomControlLibraryOptions *ReplaceCustomControlLibraryOptions)`, func() {
		replaceCustomControlLibraryPath := "/control_libraries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceCustomControlLibraryPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceCustomControlLibraryPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_requirement": true, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceCustomControlLibrary with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceCustomControlLibraryOptions model with no property values
				replaceCustomControlLibraryOptionsModelNew := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModelNew)
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
			It(`Invoke ReplaceCustomControlLibrary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(securityandcompliancecenterapiv3.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
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
	Describe(`ListProfiles(listProfilesOptions *ListProfilesOptions) - Operation response error`, func() {
		listProfilesPath := "/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(securityandcompliancecenterapiv3.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProfilesOptionsModel.ProfileType = core.StringPtr("custom")
				listProfilesOptionsModel.Start = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {
		listProfilesPath := "/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(securityandcompliancecenterapiv3.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProfilesOptionsModel.ProfileType = core.StringPtr("custom")
				listProfilesOptionsModel.Start = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"custom"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(securityandcompliancecenterapiv3.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProfilesOptionsModel.ProfileType = core.StringPtr("custom")
				listProfilesOptionsModel.Start = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(securityandcompliancecenterapiv3.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProfilesOptionsModel.ProfileType = core.StringPtr("custom")
				listProfilesOptionsModel.Start = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
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
			It(`Invoke ListProfiles successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(securityandcompliancecenterapiv3.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProfilesOptionsModel.ProfileType = core.StringPtr("custom")
				listProfilesOptionsModel.Start = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ProfileCollection)
				nextObject := new(securityandcompliancecenterapiv3.PaginatedCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ProfileCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"profiles":[{"id":"ID","profile_name":"ProfileName","profile_description":"ProfileDescription","profile_type":"ProfileType","profile_version":"ProfileVersion","version_group_label":"VersionGroupLabel","latest":true,"created_by":"CreatedBy","created_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","updated_on":"2019-01-01T12:00:00.000Z","controls_count":13,"attachments_count":16}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"profiles":[{"id":"ID","profile_name":"ProfileName","profile_description":"ProfileDescription","profile_type":"ProfileType","profile_version":"ProfileVersion","version_group_label":"VersionGroupLabel","latest":true,"created_by":"CreatedBy","created_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","updated_on":"2019-01-01T12:00:00.000Z","controls_count":13,"attachments_count":16}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ProfilesPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listProfilesOptionsModel := &securityandcompliancecenterapiv3.ListProfilesOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					ProfileType: core.StringPtr("custom"),
				}

				pager, err := securityAndComplianceCenterApiService.NewProfilesPager(listProfilesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.ProfileItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ProfilesPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listProfilesOptionsModel := &securityandcompliancecenterapiv3.ListProfilesOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					ProfileType: core.StringPtr("custom"),
				}

				pager, err := securityAndComplianceCenterApiService.NewProfilesPager(listProfilesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions) - Operation response error`, func() {
		createProfilePath := "/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProfile with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
		createProfilePath := "/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateProfileWithContext(ctx, createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateProfileWithContext(ctx, createProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfile with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProfileOptions model with no property values
				createProfileOptionsModelNew := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(securityandcompliancecenterapiv3.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProfile(createProfileOptionsModel)
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
	Describe(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions) - Operation response error`, func() {
		deleteCustomProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomProfilePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCustomProfile with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomProfile(deleteCustomProfileOptions *DeleteCustomProfileOptions)`, func() {
		deleteCustomProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomProfile with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomProfileOptions model with no property values
				deleteCustomProfileOptionsModelNew := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModelNew)
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
			It(`Invoke DeleteCustomProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(securityandcompliancecenterapiv3.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
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
	Describe(`GetProfile(getProfileOptions *GetProfileOptions) - Operation response error`, func() {
		getProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfile with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(securityandcompliancecenterapiv3.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {
		getProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(securityandcompliancecenterapiv3.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetProfileWithContext(ctx, getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetProfileWithContext(ctx, getProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(securityandcompliancecenterapiv3.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(securityandcompliancecenterapiv3.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfileOptionsModelNew := new(securityandcompliancecenterapiv3.GetProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModelNew)
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
			It(`Invoke GetProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(securityandcompliancecenterapiv3.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfile(getProfileOptionsModel)
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
	Describe(`ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions) - Operation response error`, func() {
		replaceProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfilePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceProfile with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceProfile(replaceProfileOptions *ReplaceProfileOptions)`, func() {
		replaceProfilePath := "/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfilePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke ReplaceProfile successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ReplaceProfileWithContext(ctx, replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ReplaceProfileWithContext(ctx, replaceProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfilePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke ReplaceProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfile with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileOptions model with no property values
				replaceProfileOptionsModelNew := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModelNew)
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
			It(`Invoke ReplaceProfile successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfile(replaceProfileOptionsModel)
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
	Describe(`ListRules(listRulesOptions *ListRulesOptions) - Operation response error`, func() {
		listRulesPath := "/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"system_defined"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRules with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(securityandcompliancecenterapiv3.ListRulesOptions)
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.XRequestID = core.StringPtr("testString")
				listRulesOptionsModel.Type = core.StringPtr("system_defined")
				listRulesOptionsModel.Search = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
		listRulesPath := "/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"system_defined"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 50, "total_count": 230, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "rules": [{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(securityandcompliancecenterapiv3.ListRulesOptions)
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.XRequestID = core.StringPtr("testString")
				listRulesOptionsModel.Type = core.StringPtr("system_defined")
				listRulesOptionsModel.Search = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListRulesWithContext(ctx, listRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"system_defined"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 50, "total_count": 230, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "rules": [{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(securityandcompliancecenterapiv3.ListRulesOptions)
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.XRequestID = core.StringPtr("testString")
				listRulesOptionsModel.Type = core.StringPtr("system_defined")
				listRulesOptionsModel.Search = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(securityandcompliancecenterapiv3.ListRulesOptions)
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.XRequestID = core.StringPtr("testString")
				listRulesOptionsModel.Type = core.StringPtr("system_defined")
				listRulesOptionsModel.Search = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
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
			It(`Invoke ListRules successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(securityandcompliancecenterapiv3.ListRulesOptions)
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.XRequestID = core.StringPtr("testString")
				listRulesOptionsModel.Type = core.StringPtr("system_defined")
				listRulesOptionsModel.Search = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListRules(listRulesOptionsModel)
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
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions) - Operation response error`, func() {
		createRulePath := "/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRule with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.XRequestID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
		createRulePath := "/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke CreateRule successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.XRequestID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateRuleWithContext(ctx, createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateRuleWithContext(ctx, createRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.XRequestID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRule with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.XRequestID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRuleOptions model with no property values
				createRuleOptionsModelNew := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(securityandcompliancecenterapiv3.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("Example rule")
				createRuleOptionsModel.Target = targetModel
				createRuleOptionsModel.RequiredConfig = requiredConfigModel
				createRuleOptionsModel.Type = core.StringPtr("user_defined")
				createRuleOptionsModel.Version = core.StringPtr("1.0.0")
				createRuleOptionsModel.Import = importModel
				createRuleOptionsModel.Labels = []string{}
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.XRequestID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateRule(createRuleOptionsModel)
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
	Describe(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
		deleteRulePath := "/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := securityAndComplianceCenterApiService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(securityandcompliancecenterapiv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteRuleOptionsModel.XRequestID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = securityAndComplianceCenterApiService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(securityandcompliancecenterapiv3.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteRuleOptionsModel.XRequestID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := securityAndComplianceCenterApiService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(securityandcompliancecenterapiv3.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = securityAndComplianceCenterApiService.DeleteRule(deleteRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRule(getRuleOptions *GetRuleOptions) - Operation response error`, func() {
		getRulePath := "/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRule with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(securityandcompliancecenterapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
		getRulePath := "/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(securityandcompliancecenterapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetRuleWithContext(ctx, getRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(securityandcompliancecenterapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(securityandcompliancecenterapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(securityandcompliancecenterapiv3.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetRule(getRuleOptionsModelNew)
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
			It(`Invoke GetRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(securityandcompliancecenterapiv3.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetRule(getRuleOptionsModel)
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
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) - Operation response error`, func() {
		replaceRulePath := "/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRule with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.XRequestID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
		replaceRulePath := "/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke ReplaceRule successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.XRequestID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "id": "ID", "account_id": "AccountID", "description": "Description", "type": "user_defined", "version": "Version", "import": {"parameters": [{"name": "Name", "display_name": "DisplayName", "description": "Description", "type": "string"}]}, "target": {"service_name": "ServiceName", "service_display_name": "ServiceDisplayName", "resource_kind": "ResourceKind", "additional_target_attributes": [{"name": "Name", "operator": "string_equals", "value": "Value"}]}, "required_config": {"description": "Description", "and": [{"description": "Description"}]}, "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke ReplaceRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.XRequestID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRule with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.XRequestID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRuleOptions model with no property values
				replaceRuleOptionsModelNew := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModelNew)
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
			It(`Invoke ReplaceRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(securityandcompliancecenterapiv3.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("Example rule")
				replaceRuleOptionsModel.Target = targetModel
				replaceRuleOptionsModel.RequiredConfig = requiredConfigModel
				replaceRuleOptionsModel.Type = core.StringPtr("user_defined")
				replaceRuleOptionsModel.Version = core.StringPtr("1.0.1")
				replaceRuleOptionsModel.Import = importModel
				replaceRuleOptionsModel.Labels = []string{}
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.XRequestID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceRule(replaceRuleOptionsModel)
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
	Describe(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) - Operation response error`, func() {
		listAttachmentsPath := "/profiles/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachments with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAttachments(listAttachmentsOptions *ListAttachmentsOptions)`, func() {
		listAttachmentsPath := "/profiles/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachments with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAttachmentsOptions model with no property values
				listAttachmentsOptionsModelNew := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModelNew)
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
			It(`Invoke ListAttachments successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.AttachmentCollection)
				nextObject := new(securityandcompliancecenterapiv3.PaginatedCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.AttachmentCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"attachments":[{"id":"130003ea8bfa43c5aacea07a86da3000","profile_id":"7ec45986-54fc-4b66-a303-d9577b078c65","account_id":"130003ea8bfa43c5aacea07a86da3000","instance_id":"edf9524f-406c-412c-acbb-ee371a5cabda","scope":[{"environment":"Environment","properties":[{"name":"Name","value":"Value"}]}],"created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","status":"enabled","schedule":"daily","notifications":{"enabled":false,"controls":{"threshold_limit":14,"failed_control_ids":["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}},"attachment_parameters":[{"assessment_type":"AssessmentType","assessment_id":"AssessmentID","parameter_name":"ParameterName","parameter_value":"ParameterValue","parameter_display_name":"ParameterDisplayName","parameter_type":"string"}],"last_scan":{"id":"e8a39d25-0051-4328-8462-988ad321f49a","status":"in_progress","time":"2019-01-01T12:00:00.000Z"},"next_scan_time":"2019-01-01T12:00:00.000Z","name":"account-130003ea8bfa43c5aacea07a86da3000","description":"Test description"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"attachments":[{"id":"130003ea8bfa43c5aacea07a86da3000","profile_id":"7ec45986-54fc-4b66-a303-d9577b078c65","account_id":"130003ea8bfa43c5aacea07a86da3000","instance_id":"edf9524f-406c-412c-acbb-ee371a5cabda","scope":[{"environment":"Environment","properties":[{"name":"Name","value":"Value"}]}],"created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","status":"enabled","schedule":"daily","notifications":{"enabled":false,"controls":{"threshold_limit":14,"failed_control_ids":["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}},"attachment_parameters":[{"assessment_type":"AssessmentType","assessment_id":"AssessmentID","parameter_name":"ParameterName","parameter_value":"ParameterValue","parameter_display_name":"ParameterDisplayName","parameter_type":"string"}],"last_scan":{"id":"e8a39d25-0051-4328-8462-988ad321f49a","status":"in_progress","time":"2019-01-01T12:00:00.000Z"},"next_scan_time":"2019-01-01T12:00:00.000Z","name":"account-130003ea8bfa43c5aacea07a86da3000","description":"Test description"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AttachmentsPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listAttachmentsOptionsModel := &securityandcompliancecenterapiv3.ListAttachmentsOptions{
					ProfilesID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewAttachmentsPager(listAttachmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.AttachmentItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AttachmentsPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listAttachmentsOptionsModel := &securityandcompliancecenterapiv3.ListAttachmentsOptions{
					ProfilesID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewAttachmentsPager(listAttachmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateAttachment(createAttachmentOptions *CreateAttachmentOptions) - Operation response error`, func() {
		createAttachmentPath := "/profiles/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAttachment with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAttachment(createAttachmentOptions *CreateAttachmentOptions)`, func() {
		createAttachmentPath := "/profiles/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"profile_id": "ProfileID", "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}`)
				}))
			})
			It(`Invoke CreateAttachment successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"profile_id": "ProfileID", "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}`)
				}))
			})
			It(`Invoke CreateAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAttachment with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAttachmentOptions model with no property values
				createAttachmentOptionsModelNew := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(securityandcompliancecenterapiv3.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateAttachment(createAttachmentOptionsModel)
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
	Describe(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions) - Operation response error`, func() {
		deleteProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteProfileAttachment with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProfileAttachment(deleteProfileAttachmentOptions *DeleteProfileAttachmentOptions)`, func() {
		deleteProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke DeleteProfileAttachment successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachmentWithContext(ctx, deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.DeleteProfileAttachmentWithContext(ctx, deleteProfileAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke DeleteProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProfileAttachment with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProfileAttachmentOptions model with no property values
				deleteProfileAttachmentOptionsModelNew := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModelNew)
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
			It(`Invoke DeleteProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
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
	Describe(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions) - Operation response error`, func() {
		getProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileAttachment with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileAttachment(getProfileAttachmentOptions *GetProfileAttachmentOptions)`, func() {
		getProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke GetProfileAttachment successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetProfileAttachmentWithContext(ctx, getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetProfileAttachmentWithContext(ctx, getProfileAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke GetProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileAttachment with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileAttachmentOptions model with no property values
				getProfileAttachmentOptionsModelNew := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModelNew)
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
			It(`Invoke GetProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetProfileAttachment(getProfileAttachmentOptionsModel)
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
	Describe(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions) - Operation response error`, func() {
		replaceProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("every_30_days")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceProfileAttachment(replaceProfileAttachmentOptions *ReplaceProfileAttachmentOptions)`, func() {
		replaceProfileAttachmentPath := "/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileAttachmentPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("every_30_days")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileAttachmentPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("every_30_days")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfileAttachment with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("every_30_days")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileAttachmentOptions model with no property values
				replaceProfileAttachmentOptionsModelNew := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModelNew)
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
			It(`Invoke ReplaceProfileAttachment successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(securityandcompliancecenterapiv3.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("every_30_days")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
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
	Describe(`CreateScan(createScanOptions *CreateScanOptions) - Operation response error`, func() {
		createScanPath := "/scans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createScanPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateScan with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(securityandcompliancecenterapiv3.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateScan(createScanOptions *CreateScanOptions)`, func() {
		createScanPath := "/scans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createScanPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "attachment_id": "AttachmentID", "report_id": "ReportID", "status": "completed", "last_scan_time": "LastScanTime", "next_scan_time": "NextScanTime", "scan_type": "ondemand", "occurence": 9}`)
				}))
			})
			It(`Invoke CreateScan successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(securityandcompliancecenterapiv3.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateScanWithContext(ctx, createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateScanWithContext(ctx, createScanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createScanPath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "attachment_id": "AttachmentID", "report_id": "ReportID", "status": "completed", "last_scan_time": "LastScanTime", "next_scan_time": "NextScanTime", "scan_type": "ondemand", "occurence": 9}`)
				}))
			})
			It(`Invoke CreateScan successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateScan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(securityandcompliancecenterapiv3.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateScan with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(securityandcompliancecenterapiv3.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateScanOptions model with no property values
				createScanOptionsModelNew := new(securityandcompliancecenterapiv3.CreateScanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateScan(createScanOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateScan successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(securityandcompliancecenterapiv3.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateScan(createScanOptionsModel)
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
	Describe(`ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions) - Operation response error`, func() {
		listAttachmentsAccountPath := "/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsAccountPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachmentsAccount with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsAccountOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAttachmentsAccount(listAttachmentsAccountOptions *ListAttachmentsAccountOptions)`, func() {
		listAttachmentsAccountPath := "/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsAccountPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachmentsAccount successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsAccountOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccountWithContext(ctx, listAttachmentsAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListAttachmentsAccountWithContext(ctx, listAttachmentsAccountOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsAccountPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "2019-01-01T12:00:00.000Z"}, "next_scan_time": "2019-01-01T12:00:00.000Z", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachmentsAccount successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsAccountOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachmentsAccount with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsAccountOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
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
			It(`Invoke ListAttachmentsAccount successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(securityandcompliancecenterapiv3.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAttachmentsAccountOptionsModel.Start = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.AttachmentCollection)
				nextObject := new(securityandcompliancecenterapiv3.PaginatedCollectionNext)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.AttachmentCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentsAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"attachments":[{"id":"130003ea8bfa43c5aacea07a86da3000","profile_id":"7ec45986-54fc-4b66-a303-d9577b078c65","account_id":"130003ea8bfa43c5aacea07a86da3000","instance_id":"edf9524f-406c-412c-acbb-ee371a5cabda","scope":[{"environment":"Environment","properties":[{"name":"Name","value":"Value"}]}],"created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","status":"enabled","schedule":"daily","notifications":{"enabled":false,"controls":{"threshold_limit":14,"failed_control_ids":["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}},"attachment_parameters":[{"assessment_type":"AssessmentType","assessment_id":"AssessmentID","parameter_name":"ParameterName","parameter_value":"ParameterValue","parameter_display_name":"ParameterDisplayName","parameter_type":"string"}],"last_scan":{"id":"e8a39d25-0051-4328-8462-988ad321f49a","status":"in_progress","time":"2019-01-01T12:00:00.000Z"},"next_scan_time":"2019-01-01T12:00:00.000Z","name":"account-130003ea8bfa43c5aacea07a86da3000","description":"Test description"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"attachments":[{"id":"130003ea8bfa43c5aacea07a86da3000","profile_id":"7ec45986-54fc-4b66-a303-d9577b078c65","account_id":"130003ea8bfa43c5aacea07a86da3000","instance_id":"edf9524f-406c-412c-acbb-ee371a5cabda","scope":[{"environment":"Environment","properties":[{"name":"Name","value":"Value"}]}],"created_on":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_on":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy","status":"enabled","schedule":"daily","notifications":{"enabled":false,"controls":{"threshold_limit":14,"failed_control_ids":["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}},"attachment_parameters":[{"assessment_type":"AssessmentType","assessment_id":"AssessmentID","parameter_name":"ParameterName","parameter_value":"ParameterValue","parameter_display_name":"ParameterDisplayName","parameter_type":"string"}],"last_scan":{"id":"e8a39d25-0051-4328-8462-988ad321f49a","status":"in_progress","time":"2019-01-01T12:00:00.000Z"},"next_scan_time":"2019-01-01T12:00:00.000Z","name":"account-130003ea8bfa43c5aacea07a86da3000","description":"Test description"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AttachmentsAccountPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listAttachmentsAccountOptionsModel := &securityandcompliancecenterapiv3.ListAttachmentsAccountOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewAttachmentsAccountPager(listAttachmentsAccountOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.AttachmentItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AttachmentsAccountPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listAttachmentsAccountOptionsModel := &securityandcompliancecenterapiv3.ListAttachmentsAccountOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewAttachmentsAccountPager(listAttachmentsAccountOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions) - Operation response error`, func() {
		getLatestReportsPath := "/reports/latest"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestReports with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(securityandcompliancecenterapiv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.XRequestID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("profile_name")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestReports(getLatestReportsOptions *GetLatestReportsOptions)`, func() {
		getLatestReportsPath := "/reports/latest"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "controls_summary": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations_summary": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}]}`)
				}))
			})
			It(`Invoke GetLatestReports successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(securityandcompliancecenterapiv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.XRequestID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("profile_name")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetLatestReportsWithContext(ctx, getLatestReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetLatestReportsWithContext(ctx, getLatestReportsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLatestReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "controls_summary": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations_summary": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}]}`)
				}))
			})
			It(`Invoke GetLatestReports successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetLatestReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(securityandcompliancecenterapiv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.XRequestID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("profile_name")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestReports with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(securityandcompliancecenterapiv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.XRequestID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("profile_name")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
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
			It(`Invoke GetLatestReports successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := new(securityandcompliancecenterapiv3.GetLatestReportsOptions)
				getLatestReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getLatestReportsOptionsModel.XRequestID = core.StringPtr("testString")
				getLatestReportsOptionsModel.Sort = core.StringPtr("profile_name")
				getLatestReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetLatestReports(getLatestReportsOptionsModel)
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
	Describe(`ListReports(listReportsOptions *ListReportsOptions) - Operation response error`, func() {
		listReportsPath := "/reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReports with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(securityandcompliancecenterapiv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("profile_name")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReports(listReportsOptions *ListReportsOptions)`, func() {
		listReportsPath := "/reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}]}`)
				}))
			})
			It(`Invoke ListReports successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(securityandcompliancecenterapiv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("profile_name")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListReportsWithContext(ctx, listReportsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListReportsWithContext(ctx, listReportsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["attachment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"scheduled"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"profile_name"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "reports": [{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}]}`)
				}))
			})
			It(`Invoke ListReports successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListReports(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(securityandcompliancecenterapiv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("profile_name")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReports with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(securityandcompliancecenterapiv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("profile_name")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
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
			It(`Invoke ListReports successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := new(securityandcompliancecenterapiv3.ListReportsOptions)
				listReportsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportsOptionsModel.AttachmentID = core.StringPtr("testString")
				listReportsOptionsModel.GroupID = core.StringPtr("testString")
				listReportsOptionsModel.ProfileID = core.StringPtr("testString")
				listReportsOptionsModel.Type = core.StringPtr("scheduled")
				listReportsOptionsModel.Start = core.StringPtr("testString")
				listReportsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportsOptionsModel.Sort = core.StringPtr("profile_name")
				listReportsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListReports(listReportsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ReportPage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ReportPage)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ReportPage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"reports":[{"id":"44a5-a292-32114fa73558","group_id":"55b6-b3A4-432250b84669","created_on":"2022-08-15T12:30:01Z","scan_time":"2022-08-15T12:30:01Z","type":"scheduled","cos_object":"crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7","instance_id":"84644a08-31b6-4988-b504-49a46ca69ccd","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"profile":{"id":"44a5-a292-32114fa73558","name":"IBM FS Cloud","version":"0.1"},"attachment":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"resource group - Default","description":"Scoped to the Default resource group","schedule":"daily","scope":[{"id":"ca0941aa-b7e2-43a3-9794-1b3d322474d9","environment":"ibm-cloud","properties":[{"name":"scope_id","value":"18d32a4430e54c81a6668952609763b2"}]}]}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"reports":[{"id":"44a5-a292-32114fa73558","group_id":"55b6-b3A4-432250b84669","created_on":"2022-08-15T12:30:01Z","scan_time":"2022-08-15T12:30:01Z","type":"scheduled","cos_object":"crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7","instance_id":"84644a08-31b6-4988-b504-49a46ca69ccd","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"profile":{"id":"44a5-a292-32114fa73558","name":"IBM FS Cloud","version":"0.1"},"attachment":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"resource group - Default","description":"Scoped to the Default resource group","schedule":"daily","scope":[{"id":"ca0941aa-b7e2-43a3-9794-1b3d322474d9","environment":"ibm-cloud","properties":[{"name":"scope_id","value":"18d32a4430e54c81a6668952609763b2"}]}]}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportsPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportsOptionsModel := &securityandcompliancecenterapiv3.ListReportsOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					AttachmentID: core.StringPtr("testString"),
					GroupID: core.StringPtr("testString"),
					ProfileID: core.StringPtr("testString"),
					Type: core.StringPtr("scheduled"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: core.StringPtr("profile_name"),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportsPager(listReportsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.Report
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportsPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportsOptionsModel := &securityandcompliancecenterapiv3.ListReportsOptions{
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					AttachmentID: core.StringPtr("testString"),
					GroupID: core.StringPtr("testString"),
					ProfileID: core.StringPtr("testString"),
					Type: core.StringPtr("scheduled"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: core.StringPtr("profile_name"),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportsPager(listReportsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetReport(getReportOptions *GetReportOptions) - Operation response error`, func() {
		getReportPath := "/reports/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReport with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(securityandcompliancecenterapiv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.XRequestID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReport(getReportOptions *GetReportOptions)`, func() {
		getReportPath := "/reports/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}`)
				}))
			})
			It(`Invoke GetReport successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(securityandcompliancecenterapiv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.XRequestID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportWithContext(ctx, getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportWithContext(ctx, getReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "44a5-a292-32114fa73558", "group_id": "55b6-b3A4-432250b84669", "created_on": "2022-08-15T12:30:01Z", "scan_time": "2022-08-15T12:30:01Z", "type": "scheduled", "cos_object": "crn:v1:bluemix:public:cloud-object-storage:global:a/531fc3e28bfc43c5a2cea07786d93f5c:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:b1a8f3da-49d2-4966-ae83-a8d02bc2aac7", "instance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "profile": {"id": "44a5-a292-32114fa73558", "name": "IBM FS Cloud", "version": "0.1"}, "attachment": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "resource group - Default", "description": "Scoped to the Default resource group", "schedule": "daily", "scope": [{"id": "ca0941aa-b7e2-43a3-9794-1b3d322474d9", "environment": "ibm-cloud", "properties": [{"name": "scope_id", "value": "18d32a4430e54c81a6668952609763b2"}]}]}}`)
				}))
			})
			It(`Invoke GetReport successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(securityandcompliancecenterapiv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.XRequestID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReport with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(securityandcompliancecenterapiv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.XRequestID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportOptions model with no property values
				getReportOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReport(getReportOptionsModelNew)
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
			It(`Invoke GetReport successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportOptions model
				getReportOptionsModel := new(securityandcompliancecenterapiv3.GetReportOptions)
				getReportOptionsModel.ReportID = core.StringPtr("testString")
				getReportOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportOptionsModel.XRequestID = core.StringPtr("testString")
				getReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReport(getReportOptionsModel)
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
	Describe(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions) - Operation response error`, func() {
		getReportSummaryPath := "/reports/testString/summary"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportSummary with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XRequestID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportSummary(getReportSummaryOptions *GetReportSummaryOptions)`, func() {
		getReportSummaryPath := "/reports/testString/summary"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "isntance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "resources": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "top_failed": [{"name": "my-bucket", "id": "531fc3e28bfc43c5a2cea07786d93f5c", "service": "cloud-object-storage", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}, "account": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}}`)
				}))
			})
			It(`Invoke GetReportSummary successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XRequestID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportSummaryWithContext(ctx, getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportSummaryWithContext(ctx, getReportSummaryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "isntance_id": "84644a08-31b6-4988-b504-49a46ca69ccd", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "score": {"passed": 1, "total_count": 4, "percent": 25}, "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}, "evaluations": {"status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}, "resources": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "top_failed": [{"name": "my-bucket", "id": "531fc3e28bfc43c5a2cea07786d93f5c", "service": "cloud-object-storage", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}, "account": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}}`)
				}))
			})
			It(`Invoke GetReportSummary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XRequestID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportSummary with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XRequestID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportSummaryOptions model with no property values
				getReportSummaryOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModelNew)
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
			It(`Invoke GetReportSummary successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportSummaryOptions model
				getReportSummaryOptionsModel := new(securityandcompliancecenterapiv3.GetReportSummaryOptions)
				getReportSummaryOptionsModel.ReportID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportSummaryOptionsModel.XRequestID = core.StringPtr("testString")
				getReportSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportSummary(getReportSummaryOptionsModel)
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
	Describe(`GetReportEvaluation(getReportEvaluationOptions *GetReportEvaluationOptions)`, func() {
		getReportEvaluationPath := "/reports/testString/download"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportEvaluationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for exclude_summary query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetReportEvaluation successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(securityandcompliancecenterapiv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XRequestID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportEvaluationWithContext(ctx, getReportEvaluationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportEvaluationWithContext(ctx, getReportEvaluationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportEvaluationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for exclude_summary query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/csv")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetReportEvaluation successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportEvaluation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(securityandcompliancecenterapiv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XRequestID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportEvaluation with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(securityandcompliancecenterapiv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XRequestID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportEvaluationOptions model with no property values
				getReportEvaluationOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportEvaluationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptionsModelNew)
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
			It(`Invoke GetReportEvaluation successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportEvaluationOptions model
				getReportEvaluationOptionsModel := new(securityandcompliancecenterapiv3.GetReportEvaluationOptions)
				getReportEvaluationOptionsModel.ReportID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.XRequestID = core.StringPtr("testString")
				getReportEvaluationOptionsModel.ExcludeSummary = core.BoolPtr(true)
				getReportEvaluationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportEvaluation(getReportEvaluationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())


				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportControls(getReportControlsOptions *GetReportControlsOptions) - Operation response error`, func() {
		getReportControlsPath := "/reports/testString/controls"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"control_name"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportControls with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("control_name")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportControls(getReportControlsOptions *GetReportControlsOptions)`, func() {
		getReportControlsPath := "/reports/testString/controls"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"control_name"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "home_account_id": "HomeAccountID", "report_id": "ReportID", "controls": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_version": "v1.2.3", "control_name": "Password Management", "control_description": "Password Management", "control_category": "Access Control", "control_path": "AC-2(a)", "control_specifications": [{"control_specification_id": "18d32a4430e54c81a6668952609763b2", "component_id": "cloud-object_storage", "control_specification_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "environment": "ibm cloud", "responsibility": "user", "assessments": [{"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}]}`)
				}))
			})
			It(`Invoke GetReportControls successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("control_name")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportControlsWithContext(ctx, getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportControlsWithContext(ctx, getReportControlsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["control_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_description"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["control_category"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"control_name"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10, "home_account_id": "HomeAccountID", "report_id": "ReportID", "controls": [{"id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_id": "531fc3e28bfc43c5a2cea07786d93f5c", "control_library_version": "v1.2.3", "control_name": "Password Management", "control_description": "Password Management", "control_category": "Access Control", "control_path": "AC-2(a)", "control_specifications": [{"control_specification_id": "18d32a4430e54c81a6668952609763b2", "component_id": "cloud-object_storage", "control_specification_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "environment": "ibm cloud", "responsibility": "user", "assessments": [{"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}], "status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}]}`)
				}))
			})
			It(`Invoke GetReportControls successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportControls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("control_name")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportControls with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("control_name")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportControlsOptions model with no property values
				getReportControlsOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModelNew)
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
			It(`Invoke GetReportControls successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportControlsOptions model
				getReportControlsOptionsModel := new(securityandcompliancecenterapiv3.GetReportControlsOptions)
				getReportControlsOptionsModel.ReportID = core.StringPtr("testString")
				getReportControlsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportControlsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlID = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlName = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlDescription = core.StringPtr("testString")
				getReportControlsOptionsModel.ControlCategory = core.StringPtr("testString")
				getReportControlsOptionsModel.Status = core.StringPtr("compliant")
				getReportControlsOptionsModel.Sort = core.StringPtr("control_name")
				getReportControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportControls(getReportControlsOptionsModel)
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
	Describe(`GetReportRule(getReportRuleOptions *GetReportRuleOptions) - Operation response error`, func() {
		getReportRulePath := "/reports/testString/rules/rule-8d444f8c-fd1d-48de-bcaa-f43732568761"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportRule with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportRule(getReportRuleOptions *GetReportRuleOptions)`, func() {
		getReportRulePath := "/reports/testString/rules/rule-8d444f8c-fd1d-48de-bcaa-f43732568761"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "rule-7b0560a4-df94-4629-bb76-680f3155ddda", "type": "user_defined/system_defined", "description": "rule", "version": "1.2.3", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "created_on": "2022-08-15T12:30:01Z", "created_by": "IBMid-12345", "updated_on": "2022-08-15T12:30:01Z", "updated_by": "IBMid-12345", "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetReportRule successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportRuleWithContext(ctx, getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportRuleWithContext(ctx, getReportRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "rule-7b0560a4-df94-4629-bb76-680f3155ddda", "type": "user_defined/system_defined", "description": "rule", "version": "1.2.3", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "created_on": "2022-08-15T12:30:01Z", "created_by": "IBMid-12345", "updated_on": "2022-08-15T12:30:01Z", "updated_by": "IBMid-12345", "labels": ["Labels"]}`)
				}))
			})
			It(`Invoke GetReportRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportRule with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportRuleOptions model with no property values
				getReportRuleOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModelNew)
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
			It(`Invoke GetReportRule successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportRuleOptions model
				getReportRuleOptionsModel := new(securityandcompliancecenterapiv3.GetReportRuleOptions)
				getReportRuleOptionsModel.ReportID = core.StringPtr("testString")
				getReportRuleOptionsModel.RuleID = core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportRuleOptionsModel.XRequestID = core.StringPtr("testString")
				getReportRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportRule(getReportRuleOptionsModel)
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
	Describe(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions) - Operation response error`, func() {
		listReportEvaluationsPath := "/reports/testString/evaluations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReportEvaluations with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReportEvaluations(listReportEvaluationsOptions *ListReportEvaluationsOptions)`, func() {
		listReportEvaluationsPath := "/reports/testString/evaluations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "evaluations": [{"home_account_id": "be200c80cabc456e91139e4152327456", "report_id": "44a5-a292-32114fa73558", "control_id": "28016c95-b389-447f-8a05-eabe1ad7fd24", "component_id": "cloud-object_storage", "assessment": {"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}, "evaluate_time": "2022-06-30T11:03:44.630150782Z", "target": {"id": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "resource_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "resource_name": "mybucket", "service_name": "cloud-object-storage"}, "status": "failure", "reason": "One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met", "details": {"properties": [{"property": "allowed_network", "property_description": "A description for this property", "operator": "string_equals", "expected_value": "anyValue", "found_value": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke ListReportEvaluations successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListReportEvaluationsWithContext(ctx, listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListReportEvaluationsWithContext(ctx, listReportEvaluationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["assessment_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"failure"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "evaluations": [{"home_account_id": "be200c80cabc456e91139e4152327456", "report_id": "44a5-a292-32114fa73558", "control_id": "28016c95-b389-447f-8a05-eabe1ad7fd24", "component_id": "cloud-object_storage", "assessment": {"assessment_id": "382c2b06-e6b2-43ee-b189-c1c7743b67ee", "assessment_type": "ibm-cloud-rule", "assessment_method": "ibm-cloud-rule", "assessment_description": "Check whether Cloud Object Storage is accessible only by using private endpoints", "parameter_count": 1, "parameters": [{"parameter_name": "location", "parameter_display_name": "Location", "parameter_type": "string", "parameter_value": "anyValue"}]}, "evaluate_time": "2022-06-30T11:03:44.630150782Z", "target": {"id": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "account_id": "59bcbfa6ea2f006b4ed7094c1a08dcdd", "resource_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket", "resource_name": "mybucket", "service_name": "cloud-object-storage"}, "status": "failure", "reason": "One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met", "details": {"properties": [{"property": "allowed_network", "property_description": "A description for this property", "operator": "string_equals", "expected_value": "anyValue", "found_value": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke ListReportEvaluations successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportEvaluations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReportEvaluations with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListReportEvaluationsOptions model with no property values
				listReportEvaluationsOptionsModelNew := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModelNew)
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
			It(`Invoke ListReportEvaluations successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportEvaluationsOptions model
				listReportEvaluationsOptionsModel := new(securityandcompliancecenterapiv3.ListReportEvaluationsOptions)
				listReportEvaluationsOptionsModel.ReportID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.XRequestID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.AssessmentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.ComponentID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetID = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.TargetName = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Status = core.StringPtr("failure")
				listReportEvaluationsOptionsModel.Start = core.StringPtr("testString")
				listReportEvaluationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportEvaluationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportEvaluations(listReportEvaluationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.EvaluationPage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.EvaluationPage)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(securityandcompliancecenterapiv3.EvaluationPage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportEvaluationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"evaluations":[{"home_account_id":"be200c80cabc456e91139e4152327456","report_id":"44a5-a292-32114fa73558","control_id":"28016c95-b389-447f-8a05-eabe1ad7fd24","component_id":"cloud-object_storage","assessment":{"assessment_id":"382c2b06-e6b2-43ee-b189-c1c7743b67ee","assessment_type":"ibm-cloud-rule","assessment_method":"ibm-cloud-rule","assessment_description":"Check whether Cloud Object Storage is accessible only by using private endpoints","parameter_count":1,"parameters":[{"parameter_name":"location","parameter_display_name":"Location","parameter_type":"string","parameter_value":"anyValue"}]},"evaluate_time":"2022-06-30T11:03:44.630150782Z","target":{"id":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","account_id":"59bcbfa6ea2f006b4ed7094c1a08dcdd","resource_crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","resource_name":"mybucket","service_name":"cloud-object-storage"},"status":"failure","reason":"One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met","details":{"properties":[{"property":"allowed_network","property_description":"A description for this property","operator":"string_equals","expected_value":"anyValue","found_value":"anyValue"}]}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"evaluations":[{"home_account_id":"be200c80cabc456e91139e4152327456","report_id":"44a5-a292-32114fa73558","control_id":"28016c95-b389-447f-8a05-eabe1ad7fd24","component_id":"cloud-object_storage","assessment":{"assessment_id":"382c2b06-e6b2-43ee-b189-c1c7743b67ee","assessment_type":"ibm-cloud-rule","assessment_method":"ibm-cloud-rule","assessment_description":"Check whether Cloud Object Storage is accessible only by using private endpoints","parameter_count":1,"parameters":[{"parameter_name":"location","parameter_display_name":"Location","parameter_type":"string","parameter_value":"anyValue"}]},"evaluate_time":"2022-06-30T11:03:44.630150782Z","target":{"id":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","account_id":"59bcbfa6ea2f006b4ed7094c1a08dcdd","resource_crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/59bcbfa6ea2f006b4ed7094c1a08dcdd:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:mybucket","resource_name":"mybucket","service_name":"cloud-object-storage"},"status":"failure","reason":"One or more conditions in rule rule-7b0560a4-df94-4629-bb76-680f3155ddda were not met","details":{"properties":[{"property":"allowed_network","property_description":"A description for this property","operator":"string_equals","expected_value":"anyValue","found_value":"anyValue"}]}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportEvaluationsPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportEvaluationsOptionsModel := &securityandcompliancecenterapiv3.ListReportEvaluationsOptions{
					ReportID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					AssessmentID: core.StringPtr("testString"),
					ComponentID: core.StringPtr("testString"),
					TargetID: core.StringPtr("testString"),
					TargetName: core.StringPtr("testString"),
					Status: core.StringPtr("failure"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportEvaluationsPager(listReportEvaluationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.Evaluation
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportEvaluationsPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportEvaluationsOptionsModel := &securityandcompliancecenterapiv3.ListReportEvaluationsOptions{
					ReportID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					AssessmentID: core.StringPtr("testString"),
					ComponentID: core.StringPtr("testString"),
					TargetID: core.StringPtr("testString"),
					TargetName: core.StringPtr("testString"),
					Status: core.StringPtr("failure"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportEvaluationsPager(listReportEvaluationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions) - Operation response error`, func() {
		listReportResourcesPath := "/reports/testString/resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"account_id"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReportResources with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XRequestID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Sort = core.StringPtr("account_id")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReportResources(listReportResourcesOptions *ListReportResourcesOptions)`, func() {
		listReportResourcesPath := "/reports/testString/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"account_id"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "resources": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "id": "crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::", "resource_name": "jeff's key", "component_id": "cloud-object_storage", "environment": "ibm cloud", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}`)
				}))
			})
			It(`Invoke ListReportResources successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XRequestID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Sort = core.StringPtr("account_id")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListReportResourcesWithContext(ctx, listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListReportResourcesWithContext(ctx, listReportResourcesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"compliant"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"account_id"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 230, "limit": 50, "start": "Start", "first": {"href": "Href"}, "next": {"href": "Href"}, "home_account_id": "HomeAccountID", "report_id": "ReportID", "resources": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "id": "crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::", "resource_name": "jeff's key", "component_id": "cloud-object_storage", "environment": "ibm cloud", "account": {"id": "531fc3e28bfc43c5a2cea07786d93f5c", "name": "NIST", "type": "account_type"}, "status": "compliant", "total_count": 140, "pass_count": 123, "failure_count": 12, "error_count": 5, "completed_count": 135}]}`)
				}))
			})
			It(`Invoke ListReportResources successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XRequestID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Sort = core.StringPtr("account_id")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReportResources with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XRequestID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Sort = core.StringPtr("account_id")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListReportResourcesOptions model with no property values
				listReportResourcesOptionsModelNew := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModelNew)
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
			It(`Invoke ListReportResources successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListReportResourcesOptions model
				listReportResourcesOptionsModel := new(securityandcompliancecenterapiv3.ListReportResourcesOptions)
				listReportResourcesOptionsModel.ReportID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listReportResourcesOptionsModel.XRequestID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ResourceName = core.StringPtr("testString")
				listReportResourcesOptionsModel.AccountID = core.StringPtr("testString")
				listReportResourcesOptionsModel.ComponentID = core.StringPtr("testString")
				listReportResourcesOptionsModel.Status = core.StringPtr("compliant")
				listReportResourcesOptionsModel.Sort = core.StringPtr("account_id")
				listReportResourcesOptionsModel.Start = core.StringPtr("testString")
				listReportResourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listReportResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListReportResources(listReportResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ResourcePage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ResourcePage)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(securityandcompliancecenterapiv3.ResourcePage)
				nextObject := new(securityandcompliancecenterapiv3.PageHRef)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReportResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"total_count":2,"limit":1,"resources":[{"report_id":"30b434b3-cb08-4845-af10-7a8fc682b6a8","id":"crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::","resource_name":"jeff's key","component_id":"cloud-object_storage","environment":"ibm cloud","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"status":"compliant","total_count":140,"pass_count":123,"failure_count":12,"error_count":5,"completed_count":135}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"report_id":"30b434b3-cb08-4845-af10-7a8fc682b6a8","id":"crn:v1:bluemix:public:kms:us-south:a/5af747ca19a8a278b1b6e4eec20df507:03502a50-4ea9-463c-80e5-e27ed838cdb6::","resource_name":"jeff's key","component_id":"cloud-object_storage","environment":"ibm cloud","account":{"id":"531fc3e28bfc43c5a2cea07786d93f5c","name":"NIST","type":"account_type"},"status":"compliant","total_count":140,"pass_count":123,"failure_count":12,"error_count":5,"completed_count":135}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ReportResourcesPager.GetNext successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportResourcesOptionsModel := &securityandcompliancecenterapiv3.ListReportResourcesOptions{
					ReportID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					ResourceName: core.StringPtr("testString"),
					AccountID: core.StringPtr("testString"),
					ComponentID: core.StringPtr("testString"),
					Status: core.StringPtr("compliant"),
					Sort: core.StringPtr("account_id"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportResourcesPager(listReportResourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []securityandcompliancecenterapiv3.Resource
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ReportResourcesPager.GetAll successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				listReportResourcesOptionsModel := &securityandcompliancecenterapiv3.ListReportResourcesOptions{
					ReportID: core.StringPtr("testString"),
					XCorrelationID: core.StringPtr("testString"),
					XRequestID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					ResourceName: core.StringPtr("testString"),
					AccountID: core.StringPtr("testString"),
					ComponentID: core.StringPtr("testString"),
					Status: core.StringPtr("compliant"),
					Sort: core.StringPtr("account_id"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := securityAndComplianceCenterApiService.NewReportResourcesPager(listReportResourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetReportTags(getReportTagsOptions *GetReportTagsOptions) - Operation response error`, func() {
		getReportTagsPath := "/reports/testString/tags"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportTags with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportTags(getReportTagsOptions *GetReportTagsOptions)`, func() {
		getReportTagsPath := "/reports/testString/tags"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "ReportID", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}}`)
				}))
			})
			It(`Invoke GetReportTags successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportTagsWithContext(ctx, getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportTagsWithContext(ctx, getReportTagsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportTagsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"report_id": "ReportID", "tags": {"user": ["User"], "access": ["Access"], "service": ["Service"]}}`)
				}))
			})
			It(`Invoke GetReportTags successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportTags with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportTagsOptions model with no property values
				getReportTagsOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModelNew)
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
			It(`Invoke GetReportTags successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportTagsOptions model
				getReportTagsOptionsModel := new(securityandcompliancecenterapiv3.GetReportTagsOptions)
				getReportTagsOptionsModel.ReportID = core.StringPtr("testString")
				getReportTagsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportTagsOptionsModel.XRequestID = core.StringPtr("testString")
				getReportTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportTags(getReportTagsOptionsModel)
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
	Describe(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions) - Operation response error`, func() {
		getReportViolationsDriftPath := "/reports/testString/violations_drift"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReportViolationsDrift with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XRequestID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReportViolationsDrift(getReportViolationsDriftOptions *GetReportViolationsDriftOptions)`, func() {
		getReportViolationsDriftPath := "/reports/testString/violations_drift"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "report_id": "ReportID", "data_points": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "report_group_id": "55b6-b3A4-432250b84669", "scan_time": "2022-08-15T12:30:01Z", "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}}]}`)
				}))
			})
			It(`Invoke GetReportViolationsDrift successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XRequestID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDriftWithContext(ctx, getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetReportViolationsDriftWithContext(ctx, getReportViolationsDriftOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getReportViolationsDriftPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["scan_time_duration"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"home_account_id": "HomeAccountID", "report_id": "ReportID", "data_points": [{"report_id": "30b434b3-cb08-4845-af10-7a8fc682b6a8", "report_group_id": "55b6-b3A4-432250b84669", "scan_time": "2022-08-15T12:30:01Z", "controls": {"status": "compliant", "total_count": 150, "compliant_count": 130, "not_compliant_count": 5, "unable_to_perform_count": 5, "user_evaluation_required_count": 10}}]}`)
				}))
			})
			It(`Invoke GetReportViolationsDrift successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDrift(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XRequestID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReportViolationsDrift with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XRequestID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReportViolationsDriftOptions model with no property values
				getReportViolationsDriftOptionsModelNew := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModelNew)
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
			It(`Invoke GetReportViolationsDrift successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetReportViolationsDriftOptions model
				getReportViolationsDriftOptionsModel := new(securityandcompliancecenterapiv3.GetReportViolationsDriftOptions)
				getReportViolationsDriftOptionsModel.ReportID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XCorrelationID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.XRequestID = core.StringPtr("testString")
				getReportViolationsDriftOptionsModel.ScanTimeDuration = core.Int64Ptr(int64(0))
				getReportViolationsDriftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetReportViolationsDrift(getReportViolationsDriftOptionsModel)
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
	Describe(`ListProviderTypes(listProviderTypesOptions *ListProviderTypesOptions) - Operation response error`, func() {
		listProviderTypesPath := "/provider_types"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderTypes with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypesOptions)
				listProviderTypesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProviderTypes(listProviderTypesOptions *ListProviderTypesOptions)`, func() {
		listProviderTypesPath := "/provider_types"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_types": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection", "description": "Security and Compliance Center Workload Protection helps you accelerate your Kubernetes and cloud adoption by addressing security and regulatory compliance. Easily identify vulnerabilities, check compliance, block threats and respond faster at every stage of the container and Kubernetes lifecycle.", "s2s_enabled": true, "instance_limit": 1, "mode": "PULL", "data_type": "com.sysdig.secure.results", "icon": "PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiBkYXRhLW5hbWU9IkJ1aWxkIGljb24gaGVyZSIgdmlld0JveD0iMCAwIDMyIDMyIj48ZGVmcz48bGluZWFyR3JhZGllbnQgaWQ9ImEiIHgxPSItMjgxMS4xOTgiIHgyPSItMjgxNC4xOTgiIHkxPSI1NTcuNTE3IiB5Mj0iNTU3LjUxNyIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSgyODMxLjE5OCAtNTQyLjAxNykiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLW9wYWNpdHk9IjAiLz48c3RvcCBvZmZzZXQ9Ii44Ii8+PC9saW5lYXJHcmFkaWVudD48bGluZWFyR3JhZGllbnQgeGxpbms6aHJlZj0iI2EiIGlkPSJiIiB4MT0iLTgwNi4xOTgiIHgyPSItNzk5LjE5OCIgeTE9Ii0yNDE0LjQ4MSIgeTI9Ii0yNDE0LjQ4MSIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSg4MjUuMTk4IDI0MjguOTgxKSIvPjxsaW5lYXJHcmFkaWVudCB4bGluazpocmVmPSIjYSIgaWQ9ImMiIHgxPSItODEwLjE5OCIgeDI9Ii03OTguMTk4IiB5MT0iLTI0MTkuOTgxIiB5Mj0iLTI0MTkuOTgxIiBncmFkaWVudFRyYW5zZm9ybT0idHJhbnNsYXRlKDgzMi4xOTggMjQzMi45ODEpIi8+PGxpbmVhckdyYWRpZW50IGlkPSJlIiB4MT0iLTI1MTQiIHgyPSItMjQ4MiIgeTE9Ii0yNDgyIiB5Mj0iLTI1MTQiIGdyYWRpZW50VHJhbnNmb3JtPSJtYXRyaXgoMSAwIDAgLTEgMjUxNCAtMjQ4MikiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLWNvbG9yPSIjMDhiZGJhIi8+PHN0b3Agb2Zmc2V0PSIuOSIgc3RvcC1jb2xvcj0iIzBmNjJmZSIvPjwvbGluZWFyR3JhZGllbnQ+PG1hc2sgaWQ9ImQiIHdpZHRoPSIyOS4wMTciIGhlaWdodD0iMjcuOTk2IiB4PSIxLjk4MyIgeT0iMiIgZGF0YS1uYW1lPSJtYXNrIiBtYXNrVW5pdHM9InVzZXJTcGFjZU9uVXNlIj48ZyBmaWxsPSIjZmZmIj48cGF0aCBkPSJNMjkuOTc2IDE2YzAtMy43MzktMS40NTYtNy4yNTUtNC4xMDEtOS44OTlTMTkuNzE1IDIgMTUuOTc2IDIgOC43MjEgMy40NTYgNi4wNzcgNi4xMDFjLTUuNDU5IDUuNDU5LTUuNDU5IDE0LjM0IDAgMTkuNzk4QTE0LjA0NCAxNC4wNDQgMCAwIDAgMTYgMjkuOTk1di0yLjAwMWExMi4wNCAxMi4wNCAwIDAgMS04LjUwOS0zLjUxYy00LjY3OS00LjY3OS00LjY3OS0xMi4yOTIgMC0xNi45NzEgMi4yNjctMi4yNjcgNS4yOC0zLjUxNSA4LjQ4NS0zLjUxNXM2LjIxOSAxLjI0OCA4LjQ4NSAzLjUxNSAzLjUxNSA1LjI4IDMuNTE1IDguNDg1YzAgMS4zMDgtLjIxOCAyLjU4LS42MTggMy43ODZsMS44OTcuNjMyYy40NjctMS40MDguNzIyLTIuODkyLjcyMi00LjQxOFoiLz48cGF0aCBkPSJNMjQuNyAxMy42NzVhOC45NCA4Ljk0IDAgMCAwLTQuMTkzLTUuNDY1IDguOTQyIDguOTQyIDAgMCAwLTYuODMtLjg5OSA4Ljk3MSA4Ljk3MSAwIDAgMC01LjQ2MSA0LjE5NSA4Ljk4IDguOTggMCAwIDAtLjkwMyA2LjgyOGMxLjA3NyA0LjAxNiA0LjcyMiA2LjY2IDguNjk1IDYuNjYxdi0xLjk5OGMtMy4wOS0uMDAxLTUuOTI2LTIuMDU4LTYuNzYzLTUuMTgxYTcuMDEgNy4wMSAwIDAgMSA0Ljk1LTguNTc0IDYuOTU4IDYuOTU4IDAgMCAxIDUuMzEyLjY5OSA2Ljk1NCA2Ljk1NCAwIDAgMSAzLjI2MSA0LjI1Yy4zNTkgMS4zNDIuMjc1IDIuNzMyLS4xNTQgNC4wMTNsMS45MDkuNjM2YTguOTU5IDguOTU5IDAgMCAwIC4xNzYtNS4xNjdaIi8+PC9nPjxwYXRoIGZpbGw9IiNmZmYiIGQ9Ik0xNCAxNmMwLTEuMTAzLjg5Ny0yIDItMnMyIC44OTcgMiAyYTIgMiAwIDAgMS0uMTExLjYzbDEuODg5LjYzYy4xMzMtLjM5OC4yMjItLjgxNy4yMjItMS4yNTlhNCA0IDAgMSAwLTQgNHYtMmMtMS4xMDMgMC0yLS44OTctMi0yWiIvPjxwYXRoIGZpbGw9InVybCgjYSkiIGQ9Ik0xNyAxNGgzdjNoLTN6IiB0cmFuc2Zvcm09InJvdGF0ZSgtOTAgMTguNSAxNS41KSIvPjxwYXRoIGZpbGw9InVybCgjYikiIGQ9Ik0xOSAxMmg3djVoLTd6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyMi41IDE0LjUpIi8+PHBhdGggZmlsbD0idXJsKCNjKSIgZD0iTTIyIDEwaDEydjZIMjJ6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyOCAxMykiLz48cGF0aCBkPSJNMjUgMTloNnY0aC02ek0yMCAxOGg1djVoLTV6TTE3IDE3aDN2NmgtM3oiLz48L21hc2s+PC9kZWZzPjxwYXRoIGZpbGw9IiMwMDFkNmMiIGQ9Im0yNSAzMS4wMDEtMi4xMzktMS4wMTNBNS4wMjIgNS4wMjIgMCAwIDEgMjAgMjUuNDY4VjE5aDEwdjYuNDY4YTUuMDIzIDUuMDIzIDAgMCAxLTIuODYxIDQuNTJMMjUgMzEuMDAxWm0tMy0xMHY0LjQ2OGMwIDEuMTUzLjY3NCAyLjIxOCAxLjcxNyAyLjcxMWwxLjI4My42MDcgMS4yODMtLjYwN0EzLjAxMiAzLjAxMiAwIDAgMCAyOCAyNS40Njl2LTQuNDY4aC02WiIgZGF0YS1uYW1lPSJ1dWlkLTU1ODMwNDRiLWZmMjQtNGUyNy05MDU0LTI0MDQzYWRkZmMwNiIvPjxnIG1hc2s9InVybCgjZCkiPjxwYXRoIGZpbGw9InVybCgjZSkiIGQ9Ik0wIDBoMzJ2MzJIMHoiIHRyYW5zZm9ybT0icm90YXRlKC05MCAxNiAxNikiLz48L2c+PC9zdmc+", "label": {"text": "1 per instance", "tip": "Only 1 per instance"}, "attributes": {"mapKey": {"type": "text", "display_name": "Workload Protection Instance CRN"}}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke ListProviderTypes successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypesOptions)
				listProviderTypesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListProviderTypesWithContext(ctx, listProviderTypesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListProviderTypesWithContext(ctx, listProviderTypesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_types": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection", "description": "Security and Compliance Center Workload Protection helps you accelerate your Kubernetes and cloud adoption by addressing security and regulatory compliance. Easily identify vulnerabilities, check compliance, block threats and respond faster at every stage of the container and Kubernetes lifecycle.", "s2s_enabled": true, "instance_limit": 1, "mode": "PULL", "data_type": "com.sysdig.secure.results", "icon": "PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiBkYXRhLW5hbWU9IkJ1aWxkIGljb24gaGVyZSIgdmlld0JveD0iMCAwIDMyIDMyIj48ZGVmcz48bGluZWFyR3JhZGllbnQgaWQ9ImEiIHgxPSItMjgxMS4xOTgiIHgyPSItMjgxNC4xOTgiIHkxPSI1NTcuNTE3IiB5Mj0iNTU3LjUxNyIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSgyODMxLjE5OCAtNTQyLjAxNykiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLW9wYWNpdHk9IjAiLz48c3RvcCBvZmZzZXQ9Ii44Ii8+PC9saW5lYXJHcmFkaWVudD48bGluZWFyR3JhZGllbnQgeGxpbms6aHJlZj0iI2EiIGlkPSJiIiB4MT0iLTgwNi4xOTgiIHgyPSItNzk5LjE5OCIgeTE9Ii0yNDE0LjQ4MSIgeTI9Ii0yNDE0LjQ4MSIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSg4MjUuMTk4IDI0MjguOTgxKSIvPjxsaW5lYXJHcmFkaWVudCB4bGluazpocmVmPSIjYSIgaWQ9ImMiIHgxPSItODEwLjE5OCIgeDI9Ii03OTguMTk4IiB5MT0iLTI0MTkuOTgxIiB5Mj0iLTI0MTkuOTgxIiBncmFkaWVudFRyYW5zZm9ybT0idHJhbnNsYXRlKDgzMi4xOTggMjQzMi45ODEpIi8+PGxpbmVhckdyYWRpZW50IGlkPSJlIiB4MT0iLTI1MTQiIHgyPSItMjQ4MiIgeTE9Ii0yNDgyIiB5Mj0iLTI1MTQiIGdyYWRpZW50VHJhbnNmb3JtPSJtYXRyaXgoMSAwIDAgLTEgMjUxNCAtMjQ4MikiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLWNvbG9yPSIjMDhiZGJhIi8+PHN0b3Agb2Zmc2V0PSIuOSIgc3RvcC1jb2xvcj0iIzBmNjJmZSIvPjwvbGluZWFyR3JhZGllbnQ+PG1hc2sgaWQ9ImQiIHdpZHRoPSIyOS4wMTciIGhlaWdodD0iMjcuOTk2IiB4PSIxLjk4MyIgeT0iMiIgZGF0YS1uYW1lPSJtYXNrIiBtYXNrVW5pdHM9InVzZXJTcGFjZU9uVXNlIj48ZyBmaWxsPSIjZmZmIj48cGF0aCBkPSJNMjkuOTc2IDE2YzAtMy43MzktMS40NTYtNy4yNTUtNC4xMDEtOS44OTlTMTkuNzE1IDIgMTUuOTc2IDIgOC43MjEgMy40NTYgNi4wNzcgNi4xMDFjLTUuNDU5IDUuNDU5LTUuNDU5IDE0LjM0IDAgMTkuNzk4QTE0LjA0NCAxNC4wNDQgMCAwIDAgMTYgMjkuOTk1di0yLjAwMWExMi4wNCAxMi4wNCAwIDAgMS04LjUwOS0zLjUxYy00LjY3OS00LjY3OS00LjY3OS0xMi4yOTIgMC0xNi45NzEgMi4yNjctMi4yNjcgNS4yOC0zLjUxNSA4LjQ4NS0zLjUxNXM2LjIxOSAxLjI0OCA4LjQ4NSAzLjUxNSAzLjUxNSA1LjI4IDMuNTE1IDguNDg1YzAgMS4zMDgtLjIxOCAyLjU4LS42MTggMy43ODZsMS44OTcuNjMyYy40NjctMS40MDguNzIyLTIuODkyLjcyMi00LjQxOFoiLz48cGF0aCBkPSJNMjQuNyAxMy42NzVhOC45NCA4Ljk0IDAgMCAwLTQuMTkzLTUuNDY1IDguOTQyIDguOTQyIDAgMCAwLTYuODMtLjg5OSA4Ljk3MSA4Ljk3MSAwIDAgMC01LjQ2MSA0LjE5NSA4Ljk4IDguOTggMCAwIDAtLjkwMyA2LjgyOGMxLjA3NyA0LjAxNiA0LjcyMiA2LjY2IDguNjk1IDYuNjYxdi0xLjk5OGMtMy4wOS0uMDAxLTUuOTI2LTIuMDU4LTYuNzYzLTUuMTgxYTcuMDEgNy4wMSAwIDAgMSA0Ljk1LTguNTc0IDYuOTU4IDYuOTU4IDAgMCAxIDUuMzEyLjY5OSA2Ljk1NCA2Ljk1NCAwIDAgMSAzLjI2MSA0LjI1Yy4zNTkgMS4zNDIuMjc1IDIuNzMyLS4xNTQgNC4wMTNsMS45MDkuNjM2YTguOTU5IDguOTU5IDAgMCAwIC4xNzYtNS4xNjdaIi8+PC9nPjxwYXRoIGZpbGw9IiNmZmYiIGQ9Ik0xNCAxNmMwLTEuMTAzLjg5Ny0yIDItMnMyIC44OTcgMiAyYTIgMiAwIDAgMS0uMTExLjYzbDEuODg5LjYzYy4xMzMtLjM5OC4yMjItLjgxNy4yMjItMS4yNTlhNCA0IDAgMSAwLTQgNHYtMmMtMS4xMDMgMC0yLS44OTctMi0yWiIvPjxwYXRoIGZpbGw9InVybCgjYSkiIGQ9Ik0xNyAxNGgzdjNoLTN6IiB0cmFuc2Zvcm09InJvdGF0ZSgtOTAgMTguNSAxNS41KSIvPjxwYXRoIGZpbGw9InVybCgjYikiIGQ9Ik0xOSAxMmg3djVoLTd6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyMi41IDE0LjUpIi8+PHBhdGggZmlsbD0idXJsKCNjKSIgZD0iTTIyIDEwaDEydjZIMjJ6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyOCAxMykiLz48cGF0aCBkPSJNMjUgMTloNnY0aC02ek0yMCAxOGg1djVoLTV6TTE3IDE3aDN2NmgtM3oiLz48L21hc2s+PC9kZWZzPjxwYXRoIGZpbGw9IiMwMDFkNmMiIGQ9Im0yNSAzMS4wMDEtMi4xMzktMS4wMTNBNS4wMjIgNS4wMjIgMCAwIDEgMjAgMjUuNDY4VjE5aDEwdjYuNDY4YTUuMDIzIDUuMDIzIDAgMCAxLTIuODYxIDQuNTJMMjUgMzEuMDAxWm0tMy0xMHY0LjQ2OGMwIDEuMTUzLjY3NCAyLjIxOCAxLjcxNyAyLjcxMWwxLjI4My42MDcgMS4yODMtLjYwN0EzLjAxMiAzLjAxMiAwIDAgMCAyOCAyNS40Njl2LTQuNDY4aC02WiIgZGF0YS1uYW1lPSJ1dWlkLTU1ODMwNDRiLWZmMjQtNGUyNy05MDU0LTI0MDQzYWRkZmMwNiIvPjxnIG1hc2s9InVybCgjZCkiPjxwYXRoIGZpbGw9InVybCgjZSkiIGQ9Ik0wIDBoMzJ2MzJIMHoiIHRyYW5zZm9ybT0icm90YXRlKC05MCAxNiAxNikiLz48L2c+PC9zdmc+", "label": {"text": "1 per instance", "tip": "Only 1 per instance"}, "attributes": {"mapKey": {"type": "text", "display_name": "Workload Protection Instance CRN"}}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke ListProviderTypes successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypesOptions)
				listProviderTypesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProviderTypes with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypesOptions)
				listProviderTypesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
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
			It(`Invoke ListProviderTypes successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypesOptions)
				listProviderTypesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypes(listProviderTypesOptionsModel)
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
	Describe(`GetProviderTypeByID(getProviderTypeByIdOptions *GetProviderTypeByIdOptions) - Operation response error`, func() {
		getProviderTypeByIDPath := "/provider_types/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeByIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderTypeByID with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeByIdOptions model
				getProviderTypeByIdOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				getProviderTypeByIdOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProviderTypeByID(getProviderTypeByIdOptions *GetProviderTypeByIdOptions)`, func() {
		getProviderTypeByIDPath := "/provider_types/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeByIDPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection", "description": "Security and Compliance Center Workload Protection helps you accelerate your Kubernetes and cloud adoption by addressing security and regulatory compliance. Easily identify vulnerabilities, check compliance, block threats and respond faster at every stage of the container and Kubernetes lifecycle.", "s2s_enabled": true, "instance_limit": 1, "mode": "PULL", "data_type": "com.sysdig.secure.results", "icon": "PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiBkYXRhLW5hbWU9IkJ1aWxkIGljb24gaGVyZSIgdmlld0JveD0iMCAwIDMyIDMyIj48ZGVmcz48bGluZWFyR3JhZGllbnQgaWQ9ImEiIHgxPSItMjgxMS4xOTgiIHgyPSItMjgxNC4xOTgiIHkxPSI1NTcuNTE3IiB5Mj0iNTU3LjUxNyIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSgyODMxLjE5OCAtNTQyLjAxNykiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLW9wYWNpdHk9IjAiLz48c3RvcCBvZmZzZXQ9Ii44Ii8+PC9saW5lYXJHcmFkaWVudD48bGluZWFyR3JhZGllbnQgeGxpbms6aHJlZj0iI2EiIGlkPSJiIiB4MT0iLTgwNi4xOTgiIHgyPSItNzk5LjE5OCIgeTE9Ii0yNDE0LjQ4MSIgeTI9Ii0yNDE0LjQ4MSIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSg4MjUuMTk4IDI0MjguOTgxKSIvPjxsaW5lYXJHcmFkaWVudCB4bGluazpocmVmPSIjYSIgaWQ9ImMiIHgxPSItODEwLjE5OCIgeDI9Ii03OTguMTk4IiB5MT0iLTI0MTkuOTgxIiB5Mj0iLTI0MTkuOTgxIiBncmFkaWVudFRyYW5zZm9ybT0idHJhbnNsYXRlKDgzMi4xOTggMjQzMi45ODEpIi8+PGxpbmVhckdyYWRpZW50IGlkPSJlIiB4MT0iLTI1MTQiIHgyPSItMjQ4MiIgeTE9Ii0yNDgyIiB5Mj0iLTI1MTQiIGdyYWRpZW50VHJhbnNmb3JtPSJtYXRyaXgoMSAwIDAgLTEgMjUxNCAtMjQ4MikiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLWNvbG9yPSIjMDhiZGJhIi8+PHN0b3Agb2Zmc2V0PSIuOSIgc3RvcC1jb2xvcj0iIzBmNjJmZSIvPjwvbGluZWFyR3JhZGllbnQ+PG1hc2sgaWQ9ImQiIHdpZHRoPSIyOS4wMTciIGhlaWdodD0iMjcuOTk2IiB4PSIxLjk4MyIgeT0iMiIgZGF0YS1uYW1lPSJtYXNrIiBtYXNrVW5pdHM9InVzZXJTcGFjZU9uVXNlIj48ZyBmaWxsPSIjZmZmIj48cGF0aCBkPSJNMjkuOTc2IDE2YzAtMy43MzktMS40NTYtNy4yNTUtNC4xMDEtOS44OTlTMTkuNzE1IDIgMTUuOTc2IDIgOC43MjEgMy40NTYgNi4wNzcgNi4xMDFjLTUuNDU5IDUuNDU5LTUuNDU5IDE0LjM0IDAgMTkuNzk4QTE0LjA0NCAxNC4wNDQgMCAwIDAgMTYgMjkuOTk1di0yLjAwMWExMi4wNCAxMi4wNCAwIDAgMS04LjUwOS0zLjUxYy00LjY3OS00LjY3OS00LjY3OS0xMi4yOTIgMC0xNi45NzEgMi4yNjctMi4yNjcgNS4yOC0zLjUxNSA4LjQ4NS0zLjUxNXM2LjIxOSAxLjI0OCA4LjQ4NSAzLjUxNSAzLjUxNSA1LjI4IDMuNTE1IDguNDg1YzAgMS4zMDgtLjIxOCAyLjU4LS42MTggMy43ODZsMS44OTcuNjMyYy40NjctMS40MDguNzIyLTIuODkyLjcyMi00LjQxOFoiLz48cGF0aCBkPSJNMjQuNyAxMy42NzVhOC45NCA4Ljk0IDAgMCAwLTQuMTkzLTUuNDY1IDguOTQyIDguOTQyIDAgMCAwLTYuODMtLjg5OSA4Ljk3MSA4Ljk3MSAwIDAgMC01LjQ2MSA0LjE5NSA4Ljk4IDguOTggMCAwIDAtLjkwMyA2LjgyOGMxLjA3NyA0LjAxNiA0LjcyMiA2LjY2IDguNjk1IDYuNjYxdi0xLjk5OGMtMy4wOS0uMDAxLTUuOTI2LTIuMDU4LTYuNzYzLTUuMTgxYTcuMDEgNy4wMSAwIDAgMSA0Ljk1LTguNTc0IDYuOTU4IDYuOTU4IDAgMCAxIDUuMzEyLjY5OSA2Ljk1NCA2Ljk1NCAwIDAgMSAzLjI2MSA0LjI1Yy4zNTkgMS4zNDIuMjc1IDIuNzMyLS4xNTQgNC4wMTNsMS45MDkuNjM2YTguOTU5IDguOTU5IDAgMCAwIC4xNzYtNS4xNjdaIi8+PC9nPjxwYXRoIGZpbGw9IiNmZmYiIGQ9Ik0xNCAxNmMwLTEuMTAzLjg5Ny0yIDItMnMyIC44OTcgMiAyYTIgMiAwIDAgMS0uMTExLjYzbDEuODg5LjYzYy4xMzMtLjM5OC4yMjItLjgxNy4yMjItMS4yNTlhNCA0IDAgMSAwLTQgNHYtMmMtMS4xMDMgMC0yLS44OTctMi0yWiIvPjxwYXRoIGZpbGw9InVybCgjYSkiIGQ9Ik0xNyAxNGgzdjNoLTN6IiB0cmFuc2Zvcm09InJvdGF0ZSgtOTAgMTguNSAxNS41KSIvPjxwYXRoIGZpbGw9InVybCgjYikiIGQ9Ik0xOSAxMmg3djVoLTd6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyMi41IDE0LjUpIi8+PHBhdGggZmlsbD0idXJsKCNjKSIgZD0iTTIyIDEwaDEydjZIMjJ6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyOCAxMykiLz48cGF0aCBkPSJNMjUgMTloNnY0aC02ek0yMCAxOGg1djVoLTV6TTE3IDE3aDN2NmgtM3oiLz48L21hc2s+PC9kZWZzPjxwYXRoIGZpbGw9IiMwMDFkNmMiIGQ9Im0yNSAzMS4wMDEtMi4xMzktMS4wMTNBNS4wMjIgNS4wMjIgMCAwIDEgMjAgMjUuNDY4VjE5aDEwdjYuNDY4YTUuMDIzIDUuMDIzIDAgMCAxLTIuODYxIDQuNTJMMjUgMzEuMDAxWm0tMy0xMHY0LjQ2OGMwIDEuMTUzLjY3NCAyLjIxOCAxLjcxNyAyLjcxMWwxLjI4My42MDcgMS4yODMtLjYwN0EzLjAxMiAzLjAxMiAwIDAgMCAyOCAyNS40Njl2LTQuNDY4aC02WiIgZGF0YS1uYW1lPSJ1dWlkLTU1ODMwNDRiLWZmMjQtNGUyNy05MDU0LTI0MDQzYWRkZmMwNiIvPjxnIG1hc2s9InVybCgjZCkiPjxwYXRoIGZpbGw9InVybCgjZSkiIGQ9Ik0wIDBoMzJ2MzJIMHoiIHRyYW5zZm9ybT0icm90YXRlKC05MCAxNiAxNikiLz48L2c+PC9zdmc+", "label": {"text": "1 per instance", "tip": "Only 1 per instance"}, "attributes": {"mapKey": {"type": "text", "display_name": "Workload Protection Instance CRN"}}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke GetProviderTypeByID successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetProviderTypeByIdOptions model
				getProviderTypeByIdOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				getProviderTypeByIdOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByIDWithContext(ctx, getProviderTypeByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetProviderTypeByIDWithContext(ctx, getProviderTypeByIdOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeByIDPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection", "description": "Security and Compliance Center Workload Protection helps you accelerate your Kubernetes and cloud adoption by addressing security and regulatory compliance. Easily identify vulnerabilities, check compliance, block threats and respond faster at every stage of the container and Kubernetes lifecycle.", "s2s_enabled": true, "instance_limit": 1, "mode": "PULL", "data_type": "com.sysdig.secure.results", "icon": "PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiBkYXRhLW5hbWU9IkJ1aWxkIGljb24gaGVyZSIgdmlld0JveD0iMCAwIDMyIDMyIj48ZGVmcz48bGluZWFyR3JhZGllbnQgaWQ9ImEiIHgxPSItMjgxMS4xOTgiIHgyPSItMjgxNC4xOTgiIHkxPSI1NTcuNTE3IiB5Mj0iNTU3LjUxNyIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSgyODMxLjE5OCAtNTQyLjAxNykiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLW9wYWNpdHk9IjAiLz48c3RvcCBvZmZzZXQ9Ii44Ii8+PC9saW5lYXJHcmFkaWVudD48bGluZWFyR3JhZGllbnQgeGxpbms6aHJlZj0iI2EiIGlkPSJiIiB4MT0iLTgwNi4xOTgiIHgyPSItNzk5LjE5OCIgeTE9Ii0yNDE0LjQ4MSIgeTI9Ii0yNDE0LjQ4MSIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSg4MjUuMTk4IDI0MjguOTgxKSIvPjxsaW5lYXJHcmFkaWVudCB4bGluazpocmVmPSIjYSIgaWQ9ImMiIHgxPSItODEwLjE5OCIgeDI9Ii03OTguMTk4IiB5MT0iLTI0MTkuOTgxIiB5Mj0iLTI0MTkuOTgxIiBncmFkaWVudFRyYW5zZm9ybT0idHJhbnNsYXRlKDgzMi4xOTggMjQzMi45ODEpIi8+PGxpbmVhckdyYWRpZW50IGlkPSJlIiB4MT0iLTI1MTQiIHgyPSItMjQ4MiIgeTE9Ii0yNDgyIiB5Mj0iLTI1MTQiIGdyYWRpZW50VHJhbnNmb3JtPSJtYXRyaXgoMSAwIDAgLTEgMjUxNCAtMjQ4MikiIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj48c3RvcCBvZmZzZXQ9Ii4xIiBzdG9wLWNvbG9yPSIjMDhiZGJhIi8+PHN0b3Agb2Zmc2V0PSIuOSIgc3RvcC1jb2xvcj0iIzBmNjJmZSIvPjwvbGluZWFyR3JhZGllbnQ+PG1hc2sgaWQ9ImQiIHdpZHRoPSIyOS4wMTciIGhlaWdodD0iMjcuOTk2IiB4PSIxLjk4MyIgeT0iMiIgZGF0YS1uYW1lPSJtYXNrIiBtYXNrVW5pdHM9InVzZXJTcGFjZU9uVXNlIj48ZyBmaWxsPSIjZmZmIj48cGF0aCBkPSJNMjkuOTc2IDE2YzAtMy43MzktMS40NTYtNy4yNTUtNC4xMDEtOS44OTlTMTkuNzE1IDIgMTUuOTc2IDIgOC43MjEgMy40NTYgNi4wNzcgNi4xMDFjLTUuNDU5IDUuNDU5LTUuNDU5IDE0LjM0IDAgMTkuNzk4QTE0LjA0NCAxNC4wNDQgMCAwIDAgMTYgMjkuOTk1di0yLjAwMWExMi4wNCAxMi4wNCAwIDAgMS04LjUwOS0zLjUxYy00LjY3OS00LjY3OS00LjY3OS0xMi4yOTIgMC0xNi45NzEgMi4yNjctMi4yNjcgNS4yOC0zLjUxNSA4LjQ4NS0zLjUxNXM2LjIxOSAxLjI0OCA4LjQ4NSAzLjUxNSAzLjUxNSA1LjI4IDMuNTE1IDguNDg1YzAgMS4zMDgtLjIxOCAyLjU4LS42MTggMy43ODZsMS44OTcuNjMyYy40NjctMS40MDguNzIyLTIuODkyLjcyMi00LjQxOFoiLz48cGF0aCBkPSJNMjQuNyAxMy42NzVhOC45NCA4Ljk0IDAgMCAwLTQuMTkzLTUuNDY1IDguOTQyIDguOTQyIDAgMCAwLTYuODMtLjg5OSA4Ljk3MSA4Ljk3MSAwIDAgMC01LjQ2MSA0LjE5NSA4Ljk4IDguOTggMCAwIDAtLjkwMyA2LjgyOGMxLjA3NyA0LjAxNiA0LjcyMiA2LjY2IDguNjk1IDYuNjYxdi0xLjk5OGMtMy4wOS0uMDAxLTUuOTI2LTIuMDU4LTYuNzYzLTUuMTgxYTcuMDEgNy4wMSAwIDAgMSA0Ljk1LTguNTc0IDYuOTU4IDYuOTU4IDAgMCAxIDUuMzEyLjY5OSA2Ljk1NCA2Ljk1NCAwIDAgMSAzLjI2MSA0LjI1Yy4zNTkgMS4zNDIuMjc1IDIuNzMyLS4xNTQgNC4wMTNsMS45MDkuNjM2YTguOTU5IDguOTU5IDAgMCAwIC4xNzYtNS4xNjdaIi8+PC9nPjxwYXRoIGZpbGw9IiNmZmYiIGQ9Ik0xNCAxNmMwLTEuMTAzLjg5Ny0yIDItMnMyIC44OTcgMiAyYTIgMiAwIDAgMS0uMTExLjYzbDEuODg5LjYzYy4xMzMtLjM5OC4yMjItLjgxNy4yMjItMS4yNTlhNCA0IDAgMSAwLTQgNHYtMmMtMS4xMDMgMC0yLS44OTctMi0yWiIvPjxwYXRoIGZpbGw9InVybCgjYSkiIGQ9Ik0xNyAxNGgzdjNoLTN6IiB0cmFuc2Zvcm09InJvdGF0ZSgtOTAgMTguNSAxNS41KSIvPjxwYXRoIGZpbGw9InVybCgjYikiIGQ9Ik0xOSAxMmg3djVoLTd6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyMi41IDE0LjUpIi8+PHBhdGggZmlsbD0idXJsKCNjKSIgZD0iTTIyIDEwaDEydjZIMjJ6IiB0cmFuc2Zvcm09InJvdGF0ZSg5MCAyOCAxMykiLz48cGF0aCBkPSJNMjUgMTloNnY0aC02ek0yMCAxOGg1djVoLTV6TTE3IDE3aDN2NmgtM3oiLz48L21hc2s+PC9kZWZzPjxwYXRoIGZpbGw9IiMwMDFkNmMiIGQ9Im0yNSAzMS4wMDEtMi4xMzktMS4wMTNBNS4wMjIgNS4wMjIgMCAwIDEgMjAgMjUuNDY4VjE5aDEwdjYuNDY4YTUuMDIzIDUuMDIzIDAgMCAxLTIuODYxIDQuNTJMMjUgMzEuMDAxWm0tMy0xMHY0LjQ2OGMwIDEuMTUzLjY3NCAyLjIxOCAxLjcxNyAyLjcxMWwxLjI4My42MDcgMS4yODMtLjYwN0EzLjAxMiAzLjAxMiAwIDAgMCAyOCAyNS40Njl2LTQuNDY4aC02WiIgZGF0YS1uYW1lPSJ1dWlkLTU1ODMwNDRiLWZmMjQtNGUyNy05MDU0LTI0MDQzYWRkZmMwNiIvPjxnIG1hc2s9InVybCgjZCkiPjxwYXRoIGZpbGw9InVybCgjZSkiIGQ9Ik0wIDBoMzJ2MzJIMHoiIHRyYW5zZm9ybT0icm90YXRlKC05MCAxNiAxNikiLz48L2c+PC9zdmc+", "label": {"text": "1 per instance", "tip": "Only 1 per instance"}, "attributes": {"mapKey": {"type": "text", "display_name": "Workload Protection Instance CRN"}}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke GetProviderTypeByID successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderTypeByIdOptions model
				getProviderTypeByIdOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				getProviderTypeByIdOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProviderTypeByID with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeByIdOptions model
				getProviderTypeByIdOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				getProviderTypeByIdOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderTypeByIdOptions model with no property values
				getProviderTypeByIdOptionsModelNew := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModelNew)
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
			It(`Invoke GetProviderTypeByID successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeByIdOptions model
				getProviderTypeByIdOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeByIdOptions)
				getProviderTypeByIdOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeByID(getProviderTypeByIdOptionsModel)
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
	Describe(`ListProviderTypeInstances(listProviderTypeInstancesOptions *ListProviderTypeInstancesOptions) - Operation response error`, func() {
		listProviderTypeInstancesPath := "/provider_types/testString/provider_type_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypeInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviderTypeInstances with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypeInstancesOptions model
				listProviderTypeInstancesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				listProviderTypeInstancesOptionsModel.ProviderTypeID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProviderTypeInstances(listProviderTypeInstancesOptions *ListProviderTypeInstancesOptions)`, func() {
		listProviderTypeInstancesPath := "/provider_types/testString/provider_type_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypeInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_type_instances": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke ListProviderTypeInstances successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the ListProviderTypeInstancesOptions model
				listProviderTypeInstancesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				listProviderTypeInstancesOptionsModel.ProviderTypeID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstancesWithContext(ctx, listProviderTypeInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.ListProviderTypeInstancesWithContext(ctx, listProviderTypeInstancesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProviderTypeInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_type_instances": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke ListProviderTypeInstances successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProviderTypeInstancesOptions model
				listProviderTypeInstancesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				listProviderTypeInstancesOptionsModel.ProviderTypeID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProviderTypeInstances with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypeInstancesOptions model
				listProviderTypeInstancesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				listProviderTypeInstancesOptionsModel.ProviderTypeID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProviderTypeInstancesOptions model with no property values
				listProviderTypeInstancesOptionsModelNew := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModelNew)
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
			It(`Invoke ListProviderTypeInstances successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the ListProviderTypeInstancesOptions model
				listProviderTypeInstancesOptionsModel := new(securityandcompliancecenterapiv3.ListProviderTypeInstancesOptions)
				listProviderTypeInstancesOptionsModel.ProviderTypeID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				listProviderTypeInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.ListProviderTypeInstances(listProviderTypeInstancesOptionsModel)
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
	Describe(`CreateProviderTypeInstance(createProviderTypeInstanceOptions *CreateProviderTypeInstanceOptions) - Operation response error`, func() {
		createProviderTypeInstancePath := "/provider_types/testString/provider_type_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProviderTypeInstancePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProviderTypeInstance with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateProviderTypeInstanceOptions model
				createProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				createProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				createProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProviderTypeInstance(createProviderTypeInstanceOptions *CreateProviderTypeInstanceOptions)`, func() {
		createProviderTypeInstancePath := "/provider_types/testString/provider_type_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProviderTypeInstancePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke CreateProviderTypeInstance successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the CreateProviderTypeInstanceOptions model
				createProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				createProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				createProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstanceWithContext(ctx, createProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.CreateProviderTypeInstanceWithContext(ctx, createProviderTypeInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProviderTypeInstancePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke CreateProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProviderTypeInstanceOptions model
				createProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				createProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				createProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProviderTypeInstance with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateProviderTypeInstanceOptions model
				createProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				createProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				createProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProviderTypeInstanceOptions model with no property values
				createProviderTypeInstanceOptionsModelNew := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the CreateProviderTypeInstanceOptions model
				createProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.CreateProviderTypeInstanceOptions)
				createProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				createProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				createProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.CreateProviderTypeInstance(createProviderTypeInstanceOptionsModel)
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
	Describe(`DeleteProviderTypeInstance(deleteProviderTypeInstanceOptions *DeleteProviderTypeInstanceOptions)`, func() {
		deleteProviderTypeInstancePath := "/provider_types/testString/provider_type_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProviderTypeInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := securityAndComplianceCenterApiService.DeleteProviderTypeInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProviderTypeInstanceOptions model
				deleteProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.DeleteProviderTypeInstanceOptions)
				deleteProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = securityAndComplianceCenterApiService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProviderTypeInstance with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the DeleteProviderTypeInstanceOptions model
				deleteProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.DeleteProviderTypeInstanceOptions)
				deleteProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := securityAndComplianceCenterApiService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProviderTypeInstanceOptions model with no property values
				deleteProviderTypeInstanceOptionsModelNew := new(securityandcompliancecenterapiv3.DeleteProviderTypeInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = securityAndComplianceCenterApiService.DeleteProviderTypeInstance(deleteProviderTypeInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProviderTypeInstance(getProviderTypeInstanceOptions *GetProviderTypeInstanceOptions) - Operation response error`, func() {
		getProviderTypeInstancePath := "/provider_types/testString/provider_type_instances/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderTypeInstance with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeInstanceOptions model
				getProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				getProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProviderTypeInstance(getProviderTypeInstanceOptions *GetProviderTypeInstanceOptions)`, func() {
		getProviderTypeInstancePath := "/provider_types/testString/provider_type_instances/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke GetProviderTypeInstance successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetProviderTypeInstanceOptions model
				getProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				getProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstanceWithContext(ctx, getProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetProviderTypeInstanceWithContext(ctx, getProviderTypeInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypeInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke GetProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderTypeInstanceOptions model
				getProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				getProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProviderTypeInstance with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeInstanceOptions model
				getProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				getProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProviderTypeInstanceOptions model with no property values
				getProviderTypeInstanceOptionsModelNew := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModelNew)
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
			It(`Invoke GetProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypeInstanceOptions model
				getProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypeInstanceOptions)
				getProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypeInstance(getProviderTypeInstanceOptionsModel)
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
	Describe(`UpdateProviderTypeInstance(updateProviderTypeInstanceOptions *UpdateProviderTypeInstanceOptions) - Operation response error`, func() {
		updateProviderTypeInstancePath := "/provider_types/testString/provider_type_instances/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderTypeInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProviderTypeInstance with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				updateProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				updateProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				updateProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProviderTypeInstance(updateProviderTypeInstanceOptions *UpdateProviderTypeInstanceOptions)`, func() {
		updateProviderTypeInstancePath := "/provider_types/testString/provider_type_instances/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderTypeInstancePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke UpdateProviderTypeInstance successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				updateProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				updateProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				updateProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstanceWithContext(ctx, updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.UpdateProviderTypeInstanceWithContext(ctx, updateProviderTypeInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProviderTypeInstancePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}`)
				}))
			})
			It(`Invoke UpdateProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				updateProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				updateProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				updateProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProviderTypeInstance with error: Operation validation and request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				updateProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				updateProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				updateProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProviderTypeInstanceOptions model with no property values
				updateProviderTypeInstanceOptionsModelNew := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModelNew)
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
			It(`Invoke UpdateProviderTypeInstance successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				updateProviderTypeInstanceOptionsModel := new(securityandcompliancecenterapiv3.UpdateProviderTypeInstanceOptions)
				updateProviderTypeInstanceOptionsModel.ProviderTypeID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Name = core.StringPtr("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.Attributes = map[string]interface{}{"anyKey": "anyValue"}
				updateProviderTypeInstanceOptionsModel.XCorrelationID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.XRequestID = core.StringPtr("testString")
				updateProviderTypeInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.UpdateProviderTypeInstance(updateProviderTypeInstanceOptionsModel)
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
	Describe(`GetProviderTypesInstances(getProviderTypesInstancesOptions *GetProviderTypesInstancesOptions) - Operation response error`, func() {
		getProviderTypesInstancesPath := "/provider_types_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypesInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProviderTypesInstances with error: Operation response processing error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions)
				getProviderTypesInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityAndComplianceCenterApiService.EnableRetries(0, 0)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProviderTypesInstances(getProviderTypesInstancesOptions *GetProviderTypesInstancesOptions)`, func() {
		getProviderTypesInstancesPath := "/provider_types_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypesInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_types_instances": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke GetProviderTypesInstances successfully with retries`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())
				securityAndComplianceCenterApiService.EnableRetries(0, 0)

				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions)
				getProviderTypesInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstancesWithContext(ctx, getProviderTypesInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityAndComplianceCenterApiService.DisableRetries()
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityAndComplianceCenterApiService.GetProviderTypesInstancesWithContext(ctx, getProviderTypesInstancesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProviderTypesInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Request-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Request-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"provider_types_instances": [{"id": "7588190cce3c05ac8f7942ea597dafce", "type": "workload-protection", "name": "workload-protection-instance-1", "attributes": {"anyKey": "anyValue"}, "created_at": "2023-07-24T13:14:18.884Z", "updated_at": "2023-07-24T13:14:18.884Z"}]}`)
				}))
			})
			It(`Invoke GetProviderTypesInstances successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions)
				getProviderTypesInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProviderTypesInstances with error: Operation request error`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions)
				getProviderTypesInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityAndComplianceCenterApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
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
			It(`Invoke GetProviderTypesInstances successfully`, func() {
				securityAndComplianceCenterApiService, serviceErr := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityAndComplianceCenterApiService).ToNot(BeNil())

				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := new(securityandcompliancecenterapiv3.GetProviderTypesInstancesOptions)
				getProviderTypesInstancesOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.XRequestID = core.StringPtr("testString")
				getProviderTypesInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := securityAndComplianceCenterApiService.GetProviderTypesInstances(getProviderTypesInstancesOptionsModel)
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
			securityAndComplianceCenterApiService, _ := securityandcompliancecenterapiv3.NewSecurityAndComplianceCenterApiV3(&securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3Options{
				URL:           "http://securityandcompliancecenterapiv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAttachmentPrototype successfully`, func() {
				attachments := []securityandcompliancecenterapiv3.AttachmentsPrototype{}
				_model, err := securityAndComplianceCenterApiService.NewAttachmentPrototype(attachments)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAttachmentsNotificationsPrototype successfully`, func() {
				enabled := true
				var controls *securityandcompliancecenterapiv3.FailedControls = nil
				_, err := securityAndComplianceCenterApiService.NewAttachmentsNotificationsPrototype(enabled, controls)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAttachmentsPrototype successfully`, func() {
				name := "account-130003ea8bfa43c5aacea07a86da3000"
				scope := []securityandcompliancecenterapiv3.MultiCloudScope{}
				status := "enabled"
				schedule := "daily"
				attachmentParameters := []securityandcompliancecenterapiv3.AttachmentParameterPrototype{}
				_model, err := securityAndComplianceCenterApiService.NewAttachmentsPrototype(name, scope, status, schedule, attachmentParameters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateAttachmentOptions successfully`, func() {
				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				Expect(propertyItemModel).ToNot(BeNil())
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")
				Expect(propertyItemModel.Name).To(Equal(core.StringPtr("scope_id")))
				Expect(propertyItemModel.Value).To(Equal(core.StringPtr("cg3335893hh1428692d6747cf300yeb5")))

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				Expect(multiCloudScopeModel).ToNot(BeNil())
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}
				Expect(multiCloudScopeModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(multiCloudScopeModel.Properties).To(Equal([]securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(15))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{}))

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				Expect(attachmentsNotificationsPrototypeModel).ToNot(BeNil())
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPrototypeModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(attachmentsNotificationsPrototypeModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				Expect(attachmentParameterPrototypeModel).ToNot(BeNil())
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(attachmentParameterPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(attachmentParameterPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(attachmentParameterPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(attachmentParameterPrototypeModel.ParameterValue).To(Equal(core.StringPtr("120")))
				Expect(attachmentParameterPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(attachmentParameterPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsPrototype)
				Expect(attachmentsPrototypeModel).ToNot(BeNil())
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("every_30_days")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}
				Expect(attachmentsPrototypeModel.ID).To(Equal(core.StringPtr("130003ea8bfa43c5aacea07a86da3000")))
				Expect(attachmentsPrototypeModel.Name).To(Equal(core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")))
				Expect(attachmentsPrototypeModel.Description).To(Equal(core.StringPtr("Test description")))
				Expect(attachmentsPrototypeModel.Scope).To(Equal([]securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}))
				Expect(attachmentsPrototypeModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(attachmentsPrototypeModel.Schedule).To(Equal(core.StringPtr("every_30_days")))
				Expect(attachmentsPrototypeModel.Notifications).To(Equal(attachmentsNotificationsPrototypeModel))
				Expect(attachmentsPrototypeModel.AttachmentParameters).To(Equal([]securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}))

				// Construct an instance of the CreateAttachmentOptions model
				profilesID := "testString"
				createAttachmentOptionsAttachments := []securityandcompliancecenterapiv3.AttachmentsPrototype{}
				createAttachmentOptionsModel := securityAndComplianceCenterApiService.NewCreateAttachmentOptions(profilesID, createAttachmentOptionsAttachments)
				createAttachmentOptionsModel.SetProfilesID("testString")
				createAttachmentOptionsModel.SetAttachments([]securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel})
				createAttachmentOptionsModel.SetProfileID("testString")
				createAttachmentOptionsModel.SetXCorrelationID("testString")
				createAttachmentOptionsModel.SetXRequestID("testString")
				createAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAttachmentOptionsModel).ToNot(BeNil())
				Expect(createAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Attachments).To(Equal([]securityandcompliancecenterapiv3.AttachmentsPrototype{*attachmentsPrototypeModel}))
				Expect(createAttachmentOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(parameterInfoModel.ParameterValue).To(Equal(core.StringPtr("public")))

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				Expect(implementationModel).ToNot(BeNil())
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}
				Expect(implementationModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(implementationModel.AssessmentMethod).To(Equal(core.StringPtr("ibm-cloud-rule")))
				Expect(implementationModel.AssessmentType).To(Equal(core.StringPtr("automated")))
				Expect(implementationModel.AssessmentDescription).To(Equal(core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")))
				Expect(implementationModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationModel.Parameters).To(Equal([]securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}
				Expect(controlSpecificationsModel.ControlSpecificationID).To(Equal(core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("iam-identity")))
				Expect(controlSpecificationsModel.ComponenetName).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(controlSpecificationsModel.ControlSpecificationDescription).To(Equal(core.StringPtr("IBM cloud")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]securityandcompliancecenterapiv3.Implementation{*implementationModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("sc-7")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("ibm-cloud")))

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				Expect(controlsInControlLibModel).ToNot(BeNil())
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibModel.ControlName).To(Equal(core.StringPtr("SC-7")))
				Expect(controlsInControlLibModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))
				Expect(controlsInControlLibModel.ControlDescription).To(Equal(core.StringPtr("Boundary Protection")))
				Expect(controlsInControlLibModel.ControlCategory).To(Equal(core.StringPtr("System and Communications Protection")))
				Expect(controlsInControlLibModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibModel.ControlTags).To(Equal([]string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}))
				Expect(controlsInControlLibModel.ControlSpecifications).To(Equal([]securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibModel.ControlRequirement).To(Equal(core.BoolPtr(true)))
				Expect(controlsInControlLibModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsControlLibraryName := "IBM Cloud for Financial Services"
				createCustomControlLibraryOptionsControlLibraryDescription := "IBM Cloud for Financial Services"
				createCustomControlLibraryOptionsControlLibraryType := "custom"
				createCustomControlLibraryOptionsControls := []securityandcompliancecenterapiv3.ControlsInControlLib{}
				createCustomControlLibraryOptionsModel := securityAndComplianceCenterApiService.NewCreateCustomControlLibraryOptions(createCustomControlLibraryOptionsControlLibraryName, createCustomControlLibraryOptionsControlLibraryDescription, createCustomControlLibraryOptionsControlLibraryType, createCustomControlLibraryOptionsControls)
				createCustomControlLibraryOptionsModel.SetControlLibraryName("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.SetControlLibraryDescription("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.SetControlLibraryType("custom")
				createCustomControlLibraryOptionsModel.SetControls([]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel})
				createCustomControlLibraryOptionsModel.SetVersionGroupLabel("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.SetControlLibraryVersion("1.0.0")
				createCustomControlLibraryOptionsModel.SetLatest(true)
				createCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				createCustomControlLibraryOptionsModel.SetXCorrelationID("testString")
				createCustomControlLibraryOptionsModel.SetXRequestID("testString")
				createCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryName).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryDescription).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("custom")))
				Expect(createCustomControlLibraryOptionsModel.Controls).To(Equal([]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}))
				Expect(createCustomControlLibraryOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("1.0.0")))
				Expect(createCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(createCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createCustomControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				Expect(profileControlsPrototypeModel).ToNot(BeNil())
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				Expect(profileControlsPrototypeModel.ControlLibraryID).To(Equal(core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")))
				Expect(profileControlsPrototypeModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				Expect(defaultParametersPrototypeModel).ToNot(BeNil())
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(defaultParametersPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(defaultParametersPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(defaultParametersPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(defaultParametersPrototypeModel.ParameterDefaultValue).To(Equal(core.StringPtr("120")))
				Expect(defaultParametersPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(defaultParametersPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsProfileName := "test_profile1"
				createProfileOptionsProfileDescription := "test_description1"
				createProfileOptionsProfileType := "custom"
				createProfileOptionsControls := []securityandcompliancecenterapiv3.ProfileControlsPrototype{}
				createProfileOptionsDefaultParameters := []securityandcompliancecenterapiv3.DefaultParametersPrototype{}
				createProfileOptionsModel := securityAndComplianceCenterApiService.NewCreateProfileOptions(createProfileOptionsProfileName, createProfileOptionsProfileDescription, createProfileOptionsProfileType, createProfileOptionsControls, createProfileOptionsDefaultParameters)
				createProfileOptionsModel.SetProfileName("test_profile1")
				createProfileOptionsModel.SetProfileDescription("test_description1")
				createProfileOptionsModel.SetProfileType("custom")
				createProfileOptionsModel.SetControls([]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel})
				createProfileOptionsModel.SetDefaultParameters([]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel})
				createProfileOptionsModel.SetXCorrelationID("testString")
				createProfileOptionsModel.SetXRequestID("testString")
				createProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileOptionsModel).ToNot(BeNil())
				Expect(createProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("test_profile1")))
				Expect(createProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("test_description1")))
				Expect(createProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("custom")))
				Expect(createProfileOptionsModel.Controls).To(Equal([]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}))
				Expect(createProfileOptionsModel.DefaultParameters).To(Equal([]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}))
				Expect(createProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProviderTypeInstanceOptions successfully`, func() {
				// Construct an instance of the CreateProviderTypeInstanceOptions model
				providerTypeID := "testString"
				createProviderTypeInstanceOptionsModel := securityAndComplianceCenterApiService.NewCreateProviderTypeInstanceOptions(providerTypeID)
				createProviderTypeInstanceOptionsModel.SetProviderTypeID("testString")
				createProviderTypeInstanceOptionsModel.SetName("workload-protection-instance-1")
				createProviderTypeInstanceOptionsModel.SetAttributes(map[string]interface{}{"anyKey": "anyValue"})
				createProviderTypeInstanceOptionsModel.SetXCorrelationID("testString")
				createProviderTypeInstanceOptionsModel.SetXRequestID("testString")
				createProviderTypeInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProviderTypeInstanceOptionsModel).ToNot(BeNil())
				Expect(createProviderTypeInstanceOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(createProviderTypeInstanceOptionsModel.Name).To(Equal(core.StringPtr("workload-protection-instance-1")))
				Expect(createProviderTypeInstanceOptionsModel.Attributes).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createProviderTypeInstanceOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createProviderTypeInstanceOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createProviderTypeInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRuleOptions successfully`, func() {
				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				Expect(additionalTargetAttributeModel).ToNot(BeNil())
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-east")
				Expect(additionalTargetAttributeModel.Name).To(Equal(core.StringPtr("location")))
				Expect(additionalTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(additionalTargetAttributeModel.Value).To(Equal(core.StringPtr("us-east")))

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("testString")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("cloud-object-storage")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("bucket")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}))

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				Expect(requiredConfigItemsModel).ToNot(BeNil())
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")
				Expect(requiredConfigItemsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(requiredConfigItemsModel.Property).To(Equal(core.StringPtr("hard_quota")))
				Expect(requiredConfigItemsModel.Operator).To(Equal(core.StringPtr("num_equals")))
				Expect(requiredConfigItemsModel.Value).To(Equal(core.StringPtr("${hard_quota}")))

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}
				Expect(requiredConfigModel.Description).To(Equal(core.StringPtr("The Cloud Object Storage rule.")))
				Expect(requiredConfigModel.And).To(Equal([]securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}))

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				Expect(parameterModel).ToNot(BeNil())
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")
				Expect(parameterModel.Name).To(Equal(core.StringPtr("hard_quota")))
				Expect(parameterModel.DisplayName).To(Equal(core.StringPtr("The Cloud Object Storage bucket quota.")))
				Expect(parameterModel.Description).To(Equal(core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")))
				Expect(parameterModel.Type).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}
				Expect(importModel.Parameters).To(Equal([]securityandcompliancecenterapiv3.Parameter{*parameterModel}))

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsDescription := "Example rule"
				var createRuleOptionsTarget *securityandcompliancecenterapiv3.Target = nil
				var createRuleOptionsRequiredConfig securityandcompliancecenterapiv3.RequiredConfigIntf = nil
				createRuleOptionsModel := securityAndComplianceCenterApiService.NewCreateRuleOptions(createRuleOptionsDescription, createRuleOptionsTarget, createRuleOptionsRequiredConfig)
				createRuleOptionsModel.SetDescription("Example rule")
				createRuleOptionsModel.SetTarget(targetModel)
				createRuleOptionsModel.SetRequiredConfig(requiredConfigModel)
				createRuleOptionsModel.SetType("user_defined")
				createRuleOptionsModel.SetVersion("1.0.0")
				createRuleOptionsModel.SetImport(importModel)
				createRuleOptionsModel.SetLabels([]string{})
				createRuleOptionsModel.SetXCorrelationID("testString")
				createRuleOptionsModel.SetXRequestID("testString")
				createRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRuleOptionsModel).ToNot(BeNil())
				Expect(createRuleOptionsModel.Description).To(Equal(core.StringPtr("Example rule")))
				Expect(createRuleOptionsModel.Target).To(Equal(targetModel))
				Expect(createRuleOptionsModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(createRuleOptionsModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(createRuleOptionsModel.Version).To(Equal(core.StringPtr("1.0.0")))
				Expect(createRuleOptionsModel.Import).To(Equal(importModel))
				Expect(createRuleOptionsModel.Labels).To(Equal([]string{}))
				Expect(createRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateScanOptions successfully`, func() {
				// Construct an instance of the CreateScanOptions model
				createScanOptionsAttachmentID := "testString"
				createScanOptionsModel := securityAndComplianceCenterApiService.NewCreateScanOptions(createScanOptionsAttachmentID)
				createScanOptionsModel.SetAttachmentID("testString")
				createScanOptionsModel.SetXCorrelationID("testString")
				createScanOptionsModel.SetXRequestID("testString")
				createScanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createScanOptionsModel).ToNot(BeNil())
				Expect(createScanOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the DeleteCustomControlLibraryOptions model
				controlLibrariesID := "testString"
				deleteCustomControlLibraryOptionsModel := securityAndComplianceCenterApiService.NewDeleteCustomControlLibraryOptions(controlLibrariesID)
				deleteCustomControlLibraryOptionsModel.SetControlLibrariesID("testString")
				deleteCustomControlLibraryOptionsModel.SetXCorrelationID("testString")
				deleteCustomControlLibraryOptionsModel.SetXRequestID("testString")
				deleteCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(deleteCustomControlLibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomProfileOptions successfully`, func() {
				// Construct an instance of the DeleteCustomProfileOptions model
				profilesID := "testString"
				deleteCustomProfileOptionsModel := securityAndComplianceCenterApiService.NewDeleteCustomProfileOptions(profilesID)
				deleteCustomProfileOptionsModel.SetProfilesID("testString")
				deleteCustomProfileOptionsModel.SetXCorrelationID("testString")
				deleteCustomProfileOptionsModel.SetXRequestID("testString")
				deleteCustomProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomProfileOptionsModel).ToNot(BeNil())
				Expect(deleteCustomProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the DeleteProfileAttachmentOptions model
				attachmentID := "testString"
				profilesID := "testString"
				deleteProfileAttachmentOptionsModel := securityAndComplianceCenterApiService.NewDeleteProfileAttachmentOptions(attachmentID, profilesID)
				deleteProfileAttachmentOptionsModel.SetAttachmentID("testString")
				deleteProfileAttachmentOptionsModel.SetProfilesID("testString")
				deleteProfileAttachmentOptionsModel.SetXCorrelationID("testString")
				deleteProfileAttachmentOptionsModel.SetXRequestID("testString")
				deleteProfileAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileAttachmentOptionsModel).ToNot(BeNil())
				Expect(deleteProfileAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProviderTypeInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteProviderTypeInstanceOptions model
				providerTypeID := "testString"
				providerTypeInstanceID := "testString"
				deleteProviderTypeInstanceOptionsModel := securityAndComplianceCenterApiService.NewDeleteProviderTypeInstanceOptions(providerTypeID, providerTypeInstanceID)
				deleteProviderTypeInstanceOptionsModel.SetProviderTypeID("testString")
				deleteProviderTypeInstanceOptionsModel.SetProviderTypeInstanceID("testString")
				deleteProviderTypeInstanceOptionsModel.SetXCorrelationID("testString")
				deleteProviderTypeInstanceOptionsModel.SetXRequestID("testString")
				deleteProviderTypeInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProviderTypeInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteProviderTypeInstanceOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderTypeInstanceOptionsModel.ProviderTypeInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderTypeInstanceOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderTypeInstanceOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProviderTypeInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				deleteRuleOptionsModel := securityAndComplianceCenterApiService.NewDeleteRuleOptions(ruleID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetXCorrelationID("testString")
				deleteRuleOptionsModel.SetXRequestID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetControlLibraryOptions successfully`, func() {
				// Construct an instance of the GetControlLibraryOptions model
				controlLibrariesID := "testString"
				getControlLibraryOptionsModel := securityAndComplianceCenterApiService.NewGetControlLibraryOptions(controlLibrariesID)
				getControlLibraryOptionsModel.SetControlLibrariesID("testString")
				getControlLibraryOptionsModel.SetXCorrelationID("testString")
				getControlLibraryOptionsModel.SetXRequestID("testString")
				getControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getControlLibraryOptionsModel).ToNot(BeNil())
				Expect(getControlLibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLatestReportsOptions successfully`, func() {
				// Construct an instance of the GetLatestReportsOptions model
				getLatestReportsOptionsModel := securityAndComplianceCenterApiService.NewGetLatestReportsOptions()
				getLatestReportsOptionsModel.SetXCorrelationID("testString")
				getLatestReportsOptionsModel.SetXRequestID("testString")
				getLatestReportsOptionsModel.SetSort("profile_name")
				getLatestReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestReportsOptionsModel).ToNot(BeNil())
				Expect(getLatestReportsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestReportsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestReportsOptionsModel.Sort).To(Equal(core.StringPtr("profile_name")))
				Expect(getLatestReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the GetProfileAttachmentOptions model
				attachmentID := "testString"
				profilesID := "testString"
				getProfileAttachmentOptionsModel := securityAndComplianceCenterApiService.NewGetProfileAttachmentOptions(attachmentID, profilesID)
				getProfileAttachmentOptionsModel.SetAttachmentID("testString")
				getProfileAttachmentOptionsModel.SetProfilesID("testString")
				getProfileAttachmentOptionsModel.SetXCorrelationID("testString")
				getProfileAttachmentOptionsModel.SetXRequestID("testString")
				getProfileAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileAttachmentOptionsModel).ToNot(BeNil())
				Expect(getProfileAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileOptions successfully`, func() {
				// Construct an instance of the GetProfileOptions model
				profilesID := "testString"
				getProfileOptionsModel := securityAndComplianceCenterApiService.NewGetProfileOptions(profilesID)
				getProfileOptionsModel.SetProfilesID("testString")
				getProfileOptionsModel.SetXCorrelationID("testString")
				getProfileOptionsModel.SetXRequestID("testString")
				getProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileOptionsModel).ToNot(BeNil())
				Expect(getProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderTypeByIdOptions successfully`, func() {
				// Construct an instance of the GetProviderTypeByIdOptions model
				providerTypeID := "testString"
				getProviderTypeByIdOptionsModel := securityAndComplianceCenterApiService.NewGetProviderTypeByIdOptions(providerTypeID)
				getProviderTypeByIdOptionsModel.SetProviderTypeID("testString")
				getProviderTypeByIdOptionsModel.SetXCorrelationID("testString")
				getProviderTypeByIdOptionsModel.SetXRequestID("testString")
				getProviderTypeByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderTypeByIdOptionsModel).ToNot(BeNil())
				Expect(getProviderTypeByIdOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeByIdOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeByIdOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderTypeInstanceOptions successfully`, func() {
				// Construct an instance of the GetProviderTypeInstanceOptions model
				providerTypeID := "testString"
				providerTypeInstanceID := "testString"
				getProviderTypeInstanceOptionsModel := securityAndComplianceCenterApiService.NewGetProviderTypeInstanceOptions(providerTypeID, providerTypeInstanceID)
				getProviderTypeInstanceOptionsModel.SetProviderTypeID("testString")
				getProviderTypeInstanceOptionsModel.SetProviderTypeInstanceID("testString")
				getProviderTypeInstanceOptionsModel.SetXCorrelationID("testString")
				getProviderTypeInstanceOptionsModel.SetXRequestID("testString")
				getProviderTypeInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderTypeInstanceOptionsModel).ToNot(BeNil())
				Expect(getProviderTypeInstanceOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeInstanceOptionsModel.ProviderTypeInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeInstanceOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeInstanceOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypeInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProviderTypesInstancesOptions successfully`, func() {
				// Construct an instance of the GetProviderTypesInstancesOptions model
				getProviderTypesInstancesOptionsModel := securityAndComplianceCenterApiService.NewGetProviderTypesInstancesOptions()
				getProviderTypesInstancesOptionsModel.SetXCorrelationID("testString")
				getProviderTypesInstancesOptionsModel.SetXRequestID("testString")
				getProviderTypesInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProviderTypesInstancesOptionsModel).ToNot(BeNil())
				Expect(getProviderTypesInstancesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypesInstancesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getProviderTypesInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportControlsOptions successfully`, func() {
				// Construct an instance of the GetReportControlsOptions model
				reportID := "testString"
				getReportControlsOptionsModel := securityAndComplianceCenterApiService.NewGetReportControlsOptions(reportID)
				getReportControlsOptionsModel.SetReportID("testString")
				getReportControlsOptionsModel.SetXCorrelationID("testString")
				getReportControlsOptionsModel.SetXRequestID("testString")
				getReportControlsOptionsModel.SetControlID("testString")
				getReportControlsOptionsModel.SetControlName("testString")
				getReportControlsOptionsModel.SetControlDescription("testString")
				getReportControlsOptionsModel.SetControlCategory("testString")
				getReportControlsOptionsModel.SetStatus("compliant")
				getReportControlsOptionsModel.SetSort("control_name")
				getReportControlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportControlsOptionsModel).ToNot(BeNil())
				Expect(getReportControlsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlID).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlName).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlDescription).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.ControlCategory).To(Equal(core.StringPtr("testString")))
				Expect(getReportControlsOptionsModel.Status).To(Equal(core.StringPtr("compliant")))
				Expect(getReportControlsOptionsModel.Sort).To(Equal(core.StringPtr("control_name")))
				Expect(getReportControlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportEvaluationOptions successfully`, func() {
				// Construct an instance of the GetReportEvaluationOptions model
				reportID := "testString"
				getReportEvaluationOptionsModel := securityAndComplianceCenterApiService.NewGetReportEvaluationOptions(reportID)
				getReportEvaluationOptionsModel.SetReportID("testString")
				getReportEvaluationOptionsModel.SetXCorrelationID("testString")
				getReportEvaluationOptionsModel.SetXRequestID("testString")
				getReportEvaluationOptionsModel.SetExcludeSummary(true)
				getReportEvaluationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportEvaluationOptionsModel).ToNot(BeNil())
				Expect(getReportEvaluationOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportEvaluationOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportEvaluationOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportEvaluationOptionsModel.ExcludeSummary).To(Equal(core.BoolPtr(true)))
				Expect(getReportEvaluationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportOptions successfully`, func() {
				// Construct an instance of the GetReportOptions model
				reportID := "testString"
				getReportOptionsModel := securityAndComplianceCenterApiService.NewGetReportOptions(reportID)
				getReportOptionsModel.SetReportID("testString")
				getReportOptionsModel.SetXCorrelationID("testString")
				getReportOptionsModel.SetXRequestID("testString")
				getReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportOptionsModel).ToNot(BeNil())
				Expect(getReportOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportRuleOptions successfully`, func() {
				// Construct an instance of the GetReportRuleOptions model
				reportID := "testString"
				ruleID := "rule-8d444f8c-fd1d-48de-bcaa-f43732568761"
				getReportRuleOptionsModel := securityAndComplianceCenterApiService.NewGetReportRuleOptions(reportID, ruleID)
				getReportRuleOptionsModel.SetReportID("testString")
				getReportRuleOptionsModel.SetRuleID("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")
				getReportRuleOptionsModel.SetXCorrelationID("testString")
				getReportRuleOptionsModel.SetXRequestID("testString")
				getReportRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportRuleOptionsModel).ToNot(BeNil())
				Expect(getReportRuleOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.RuleID).To(Equal(core.StringPtr("rule-8d444f8c-fd1d-48de-bcaa-f43732568761")))
				Expect(getReportRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportSummaryOptions successfully`, func() {
				// Construct an instance of the GetReportSummaryOptions model
				reportID := "testString"
				getReportSummaryOptionsModel := securityAndComplianceCenterApiService.NewGetReportSummaryOptions(reportID)
				getReportSummaryOptionsModel.SetReportID("testString")
				getReportSummaryOptionsModel.SetXCorrelationID("testString")
				getReportSummaryOptionsModel.SetXRequestID("testString")
				getReportSummaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportSummaryOptionsModel).ToNot(BeNil())
				Expect(getReportSummaryOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportSummaryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportSummaryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportSummaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportTagsOptions successfully`, func() {
				// Construct an instance of the GetReportTagsOptions model
				reportID := "testString"
				getReportTagsOptionsModel := securityAndComplianceCenterApiService.NewGetReportTagsOptions(reportID)
				getReportTagsOptionsModel.SetReportID("testString")
				getReportTagsOptionsModel.SetXCorrelationID("testString")
				getReportTagsOptionsModel.SetXRequestID("testString")
				getReportTagsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportTagsOptionsModel).ToNot(BeNil())
				Expect(getReportTagsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportTagsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportTagsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportTagsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReportViolationsDriftOptions successfully`, func() {
				// Construct an instance of the GetReportViolationsDriftOptions model
				reportID := "testString"
				getReportViolationsDriftOptionsModel := securityAndComplianceCenterApiService.NewGetReportViolationsDriftOptions(reportID)
				getReportViolationsDriftOptionsModel.SetReportID("testString")
				getReportViolationsDriftOptionsModel.SetXCorrelationID("testString")
				getReportViolationsDriftOptionsModel.SetXRequestID("testString")
				getReportViolationsDriftOptionsModel.SetScanTimeDuration(int64(0))
				getReportViolationsDriftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReportViolationsDriftOptionsModel).ToNot(BeNil())
				Expect(getReportViolationsDriftOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(getReportViolationsDriftOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getReportViolationsDriftOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getReportViolationsDriftOptionsModel.ScanTimeDuration).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getReportViolationsDriftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				getRuleOptionsModel := securityAndComplianceCenterApiService.NewGetRuleOptions(ruleID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetXCorrelationID("testString")
				getRuleOptionsModel.SetXRequestID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := securityAndComplianceCenterApiService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				getSettingsOptionsModel.SetXRequestID("testString")
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
				Expect(getSettingsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAttachmentsAccountOptions successfully`, func() {
				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := securityAndComplianceCenterApiService.NewListAttachmentsAccountOptions()
				listAttachmentsAccountOptionsModel.SetXCorrelationID("testString")
				listAttachmentsAccountOptionsModel.SetXRequestID("testString")
				listAttachmentsAccountOptionsModel.SetLimit(int64(10))
				listAttachmentsAccountOptionsModel.SetStart("testString")
				listAttachmentsAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentsAccountOptionsModel).ToNot(BeNil())
				Expect(listAttachmentsAccountOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsAccountOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsAccountOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAttachmentsAccountOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAttachmentsOptions successfully`, func() {
				// Construct an instance of the ListAttachmentsOptions model
				profilesID := "testString"
				listAttachmentsOptionsModel := securityAndComplianceCenterApiService.NewListAttachmentsOptions(profilesID)
				listAttachmentsOptionsModel.SetProfilesID("testString")
				listAttachmentsOptionsModel.SetXCorrelationID("testString")
				listAttachmentsOptionsModel.SetXRequestID("testString")
				listAttachmentsOptionsModel.SetLimit(int64(10))
				listAttachmentsOptionsModel.SetStart("testString")
				listAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentsOptionsModel).ToNot(BeNil())
				Expect(listAttachmentsOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAttachmentsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListControlLibrariesOptions successfully`, func() {
				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := securityAndComplianceCenterApiService.NewListControlLibrariesOptions()
				listControlLibrariesOptionsModel.SetXCorrelationID("testString")
				listControlLibrariesOptionsModel.SetXRequestID("testString")
				listControlLibrariesOptionsModel.SetLimit(int64(50))
				listControlLibrariesOptionsModel.SetControlLibraryType("custom")
				listControlLibrariesOptionsModel.SetStart("testString")
				listControlLibrariesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listControlLibrariesOptionsModel).ToNot(BeNil())
				Expect(listControlLibrariesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listControlLibrariesOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("custom")))
				Expect(listControlLibrariesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := securityAndComplianceCenterApiService.NewListProfilesOptions()
				listProfilesOptionsModel.SetXCorrelationID("testString")
				listProfilesOptionsModel.SetXRequestID("testString")
				listProfilesOptionsModel.SetLimit(int64(10))
				listProfilesOptionsModel.SetProfileType("custom")
				listProfilesOptionsModel.SetStart("testString")
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listProfilesOptionsModel.ProfileType).To(Equal(core.StringPtr("custom")))
				Expect(listProfilesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderTypeInstancesOptions successfully`, func() {
				// Construct an instance of the ListProviderTypeInstancesOptions model
				providerTypeID := "testString"
				listProviderTypeInstancesOptionsModel := securityAndComplianceCenterApiService.NewListProviderTypeInstancesOptions(providerTypeID)
				listProviderTypeInstancesOptionsModel.SetProviderTypeID("testString")
				listProviderTypeInstancesOptionsModel.SetXCorrelationID("testString")
				listProviderTypeInstancesOptionsModel.SetXRequestID("testString")
				listProviderTypeInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderTypeInstancesOptionsModel).ToNot(BeNil())
				Expect(listProviderTypeInstancesOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(listProviderTypeInstancesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listProviderTypeInstancesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listProviderTypeInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProviderTypesOptions successfully`, func() {
				// Construct an instance of the ListProviderTypesOptions model
				listProviderTypesOptionsModel := securityAndComplianceCenterApiService.NewListProviderTypesOptions()
				listProviderTypesOptionsModel.SetXCorrelationID("testString")
				listProviderTypesOptionsModel.SetXRequestID("testString")
				listProviderTypesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProviderTypesOptionsModel).ToNot(BeNil())
				Expect(listProviderTypesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listProviderTypesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listProviderTypesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportEvaluationsOptions successfully`, func() {
				// Construct an instance of the ListReportEvaluationsOptions model
				reportID := "testString"
				listReportEvaluationsOptionsModel := securityAndComplianceCenterApiService.NewListReportEvaluationsOptions(reportID)
				listReportEvaluationsOptionsModel.SetReportID("testString")
				listReportEvaluationsOptionsModel.SetXCorrelationID("testString")
				listReportEvaluationsOptionsModel.SetXRequestID("testString")
				listReportEvaluationsOptionsModel.SetAssessmentID("testString")
				listReportEvaluationsOptionsModel.SetComponentID("testString")
				listReportEvaluationsOptionsModel.SetTargetID("testString")
				listReportEvaluationsOptionsModel.SetTargetName("testString")
				listReportEvaluationsOptionsModel.SetStatus("failure")
				listReportEvaluationsOptionsModel.SetStart("testString")
				listReportEvaluationsOptionsModel.SetLimit(int64(10))
				listReportEvaluationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportEvaluationsOptionsModel).ToNot(BeNil())
				Expect(listReportEvaluationsOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.TargetID).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.TargetName).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.Status).To(Equal(core.StringPtr("failure")))
				Expect(listReportEvaluationsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportEvaluationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportEvaluationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportResourcesOptions successfully`, func() {
				// Construct an instance of the ListReportResourcesOptions model
				reportID := "testString"
				listReportResourcesOptionsModel := securityAndComplianceCenterApiService.NewListReportResourcesOptions(reportID)
				listReportResourcesOptionsModel.SetReportID("testString")
				listReportResourcesOptionsModel.SetXCorrelationID("testString")
				listReportResourcesOptionsModel.SetXRequestID("testString")
				listReportResourcesOptionsModel.SetID("testString")
				listReportResourcesOptionsModel.SetResourceName("testString")
				listReportResourcesOptionsModel.SetAccountID("testString")
				listReportResourcesOptionsModel.SetComponentID("testString")
				listReportResourcesOptionsModel.SetStatus("compliant")
				listReportResourcesOptionsModel.SetSort("account_id")
				listReportResourcesOptionsModel.SetStart("testString")
				listReportResourcesOptionsModel.SetLimit(int64(10))
				listReportResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportResourcesOptionsModel).ToNot(BeNil())
				Expect(listReportResourcesOptionsModel.ReportID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ResourceName).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.Status).To(Equal(core.StringPtr("compliant")))
				Expect(listReportResourcesOptionsModel.Sort).To(Equal(core.StringPtr("account_id")))
				Expect(listReportResourcesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportResourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReportsOptions successfully`, func() {
				// Construct an instance of the ListReportsOptions model
				listReportsOptionsModel := securityAndComplianceCenterApiService.NewListReportsOptions()
				listReportsOptionsModel.SetXCorrelationID("testString")
				listReportsOptionsModel.SetXRequestID("testString")
				listReportsOptionsModel.SetAttachmentID("testString")
				listReportsOptionsModel.SetGroupID("testString")
				listReportsOptionsModel.SetProfileID("testString")
				listReportsOptionsModel.SetType("scheduled")
				listReportsOptionsModel.SetStart("testString")
				listReportsOptionsModel.SetLimit(int64(10))
				listReportsOptionsModel.SetSort("profile_name")
				listReportsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReportsOptionsModel).ToNot(BeNil())
				Expect(listReportsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.Type).To(Equal(core.StringPtr("scheduled")))
				Expect(listReportsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listReportsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listReportsOptionsModel.Sort).To(Equal(core.StringPtr("profile_name")))
				Expect(listReportsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := securityAndComplianceCenterApiService.NewListRulesOptions()
				listRulesOptionsModel.SetXCorrelationID("testString")
				listRulesOptionsModel.SetXRequestID("testString")
				listRulesOptionsModel.SetType("system_defined")
				listRulesOptionsModel.SetSearch("testString")
				listRulesOptionsModel.SetServiceName("testString")
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Type).To(Equal(core.StringPtr("system_defined")))
				Expect(listRulesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMultiCloudScope successfully`, func() {
				environment := "testString"
				properties := []securityandcompliancecenterapiv3.PropertyItem{}
				_model, err := securityAndComplianceCenterApiService.NewMultiCloudScope(environment, properties)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostTestEventOptions successfully`, func() {
				// Construct an instance of the PostTestEventOptions model
				postTestEventOptionsModel := securityAndComplianceCenterApiService.NewPostTestEventOptions()
				postTestEventOptionsModel.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				postTestEventOptionsModel.SetXRequestID("testString")
				postTestEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postTestEventOptionsModel).ToNot(BeNil())
				Expect(postTestEventOptionsModel.XCorrelationID).To(Equal(core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
				Expect(postTestEventOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(postTestEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(securityandcompliancecenterapiv3.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				parameterInfoModel.ParameterValue = core.StringPtr("public")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(parameterInfoModel.ParameterValue).To(Equal(core.StringPtr("public")))

				// Construct an instance of the Implementation model
				implementationModel := new(securityandcompliancecenterapiv3.Implementation)
				Expect(implementationModel).ToNot(BeNil())
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}
				Expect(implementationModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(implementationModel.AssessmentMethod).To(Equal(core.StringPtr("ibm-cloud-rule")))
				Expect(implementationModel.AssessmentType).To(Equal(core.StringPtr("automated")))
				Expect(implementationModel.AssessmentDescription).To(Equal(core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")))
				Expect(implementationModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationModel.Parameters).To(Equal([]securityandcompliancecenterapiv3.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(securityandcompliancecenterapiv3.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []securityandcompliancecenterapiv3.Implementation{*implementationModel}
				Expect(controlSpecificationsModel.ControlSpecificationID).To(Equal(core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("iam-identity")))
				Expect(controlSpecificationsModel.ComponenetName).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(controlSpecificationsModel.ControlSpecificationDescription).To(Equal(core.StringPtr("IBM cloud")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]securityandcompliancecenterapiv3.Implementation{*implementationModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(securityandcompliancecenterapiv3.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("sc-7")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("ibm-cloud")))

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(securityandcompliancecenterapiv3.ControlsInControlLib)
				Expect(controlsInControlLibModel).ToNot(BeNil())
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(true)
				controlsInControlLibModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibModel.ControlName).To(Equal(core.StringPtr("SC-7")))
				Expect(controlsInControlLibModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))
				Expect(controlsInControlLibModel.ControlDescription).To(Equal(core.StringPtr("Boundary Protection")))
				Expect(controlsInControlLibModel.ControlCategory).To(Equal(core.StringPtr("System and Communications Protection")))
				Expect(controlsInControlLibModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibModel.ControlTags).To(Equal([]string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}))
				Expect(controlsInControlLibModel.ControlSpecifications).To(Equal([]securityandcompliancecenterapiv3.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibModel.ControlRequirement).To(Equal(core.BoolPtr(true)))
				Expect(controlsInControlLibModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				controlLibrariesID := "testString"
				replaceCustomControlLibraryOptionsModel := securityAndComplianceCenterApiService.NewReplaceCustomControlLibraryOptions(controlLibrariesID)
				replaceCustomControlLibraryOptionsModel.SetControlLibrariesID("testString")
				replaceCustomControlLibraryOptionsModel.SetID("testString")
				replaceCustomControlLibraryOptionsModel.SetAccountID("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryName("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryDescription("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryType("custom")
				replaceCustomControlLibraryOptionsModel.SetVersionGroupLabel("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryVersion("1.1.0")
				replaceCustomControlLibraryOptionsModel.SetCreatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceCustomControlLibraryOptionsModel.SetCreatedBy("testString")
				replaceCustomControlLibraryOptionsModel.SetUpdatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceCustomControlLibraryOptionsModel.SetUpdatedBy("testString")
				replaceCustomControlLibraryOptionsModel.SetLatest(true)
				replaceCustomControlLibraryOptionsModel.SetHierarchyEnabled(true)
				replaceCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				replaceCustomControlLibraryOptionsModel.SetControlParentsCount(int64(38))
				replaceCustomControlLibraryOptionsModel.SetControls([]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel})
				replaceCustomControlLibraryOptionsModel.SetXCorrelationID("testString")
				replaceCustomControlLibraryOptionsModel.SetXRequestID("testString")
				replaceCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryName).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryDescription).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("custom")))
				Expect(replaceCustomControlLibraryOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("1.1.0")))
				Expect(replaceCustomControlLibraryOptionsModel.CreatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceCustomControlLibraryOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceCustomControlLibraryOptionsModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(replaceCustomControlLibraryOptionsModel.HierarchyEnabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceCustomControlLibraryOptionsModel.ControlParentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceCustomControlLibraryOptionsModel.Controls).To(Equal([]securityandcompliancecenterapiv3.ControlsInControlLib{*controlsInControlLibModel}))
				Expect(replaceCustomControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the PropertyItem model
				propertyItemModel := new(securityandcompliancecenterapiv3.PropertyItem)
				Expect(propertyItemModel).ToNot(BeNil())
				propertyItemModel.Name = core.StringPtr("scope_id")
				propertyItemModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")
				Expect(propertyItemModel.Name).To(Equal(core.StringPtr("scope_id")))
				Expect(propertyItemModel.Value).To(Equal(core.StringPtr("cg3335893hh1428692d6747cf300yeb5")))

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(securityandcompliancecenterapiv3.MultiCloudScope)
				Expect(multiCloudScopeModel).ToNot(BeNil())
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}
				Expect(multiCloudScopeModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(multiCloudScopeModel.Properties).To(Equal([]securityandcompliancecenterapiv3.PropertyItem{*propertyItemModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(securityandcompliancecenterapiv3.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(15))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{}))

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentsNotificationsPrototype)
				Expect(attachmentsNotificationsPrototypeModel).ToNot(BeNil())
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPrototypeModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(attachmentsNotificationsPrototypeModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the AttachmentParameterPrototype model
				attachmentParameterPrototypeModel := new(securityandcompliancecenterapiv3.AttachmentParameterPrototype)
				Expect(attachmentParameterPrototypeModel).ToNot(BeNil())
				attachmentParameterPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParameterPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParameterPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParameterPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParameterPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParameterPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(attachmentParameterPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(attachmentParameterPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(attachmentParameterPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(attachmentParameterPrototypeModel.ParameterValue).To(Equal(core.StringPtr("120")))
				Expect(attachmentParameterPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(attachmentParameterPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the LastScan model
				lastScanModel := new(securityandcompliancecenterapiv3.LastScan)
				Expect(lastScanModel).ToNot(BeNil())
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				Expect(lastScanModel.ID).To(Equal(core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")))
				Expect(lastScanModel.Status).To(Equal(core.StringPtr("in_progress")))
				Expect(lastScanModel.Time).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				attachmentID := "testString"
				profilesID := "testString"
				replaceProfileAttachmentOptionsModel := securityAndComplianceCenterApiService.NewReplaceProfileAttachmentOptions(attachmentID, profilesID)
				replaceProfileAttachmentOptionsModel.SetAttachmentID("testString")
				replaceProfileAttachmentOptionsModel.SetProfilesID("testString")
				replaceProfileAttachmentOptionsModel.SetID("testString")
				replaceProfileAttachmentOptionsModel.SetProfileID("testString")
				replaceProfileAttachmentOptionsModel.SetAccountID("testString")
				replaceProfileAttachmentOptionsModel.SetInstanceID("testString")
				replaceProfileAttachmentOptionsModel.SetScope([]securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel})
				replaceProfileAttachmentOptionsModel.SetCreatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceProfileAttachmentOptionsModel.SetCreatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetUpdatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceProfileAttachmentOptionsModel.SetUpdatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetStatus("enabled")
				replaceProfileAttachmentOptionsModel.SetSchedule("every_30_days")
				replaceProfileAttachmentOptionsModel.SetNotifications(attachmentsNotificationsPrototypeModel)
				replaceProfileAttachmentOptionsModel.SetAttachmentParameters([]securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel})
				replaceProfileAttachmentOptionsModel.SetLastScan(lastScanModel)
				replaceProfileAttachmentOptionsModel.SetNextScanTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceProfileAttachmentOptionsModel.SetName("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.SetDescription("Test description")
				replaceProfileAttachmentOptionsModel.SetXCorrelationID("testString")
				replaceProfileAttachmentOptionsModel.SetXRequestID("testString")
				replaceProfileAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProfileAttachmentOptionsModel).ToNot(BeNil())
				Expect(replaceProfileAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Scope).To(Equal([]securityandcompliancecenterapiv3.MultiCloudScope{*multiCloudScopeModel}))
				Expect(replaceProfileAttachmentOptionsModel.CreatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceProfileAttachmentOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(replaceProfileAttachmentOptionsModel.Schedule).To(Equal(core.StringPtr("every_30_days")))
				Expect(replaceProfileAttachmentOptionsModel.Notifications).To(Equal(attachmentsNotificationsPrototypeModel))
				Expect(replaceProfileAttachmentOptionsModel.AttachmentParameters).To(Equal([]securityandcompliancecenterapiv3.AttachmentParameterPrototype{*attachmentParameterPrototypeModel}))
				Expect(replaceProfileAttachmentOptionsModel.LastScan).To(Equal(lastScanModel))
				Expect(replaceProfileAttachmentOptionsModel.NextScanTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceProfileAttachmentOptionsModel.Name).To(Equal(core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")))
				Expect(replaceProfileAttachmentOptionsModel.Description).To(Equal(core.StringPtr("Test description")))
				Expect(replaceProfileAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(securityandcompliancecenterapiv3.ProfileControlsPrototype)
				Expect(profileControlsPrototypeModel).ToNot(BeNil())
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				Expect(profileControlsPrototypeModel.ControlLibraryID).To(Equal(core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")))
				Expect(profileControlsPrototypeModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(securityandcompliancecenterapiv3.DefaultParametersPrototype)
				Expect(defaultParametersPrototypeModel).ToNot(BeNil())
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(defaultParametersPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(defaultParametersPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(defaultParametersPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(defaultParametersPrototypeModel.ParameterDefaultValue).To(Equal(core.StringPtr("120")))
				Expect(defaultParametersPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(defaultParametersPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ReplaceProfileOptions model
				profilesID := "testString"
				replaceProfileOptionsProfileName := "test_profile1"
				replaceProfileOptionsProfileDescription := "test_description1"
				replaceProfileOptionsProfileType := "custom"
				replaceProfileOptionsControls := []securityandcompliancecenterapiv3.ProfileControlsPrototype{}
				replaceProfileOptionsDefaultParameters := []securityandcompliancecenterapiv3.DefaultParametersPrototype{}
				replaceProfileOptionsModel := securityAndComplianceCenterApiService.NewReplaceProfileOptions(profilesID, replaceProfileOptionsProfileName, replaceProfileOptionsProfileDescription, replaceProfileOptionsProfileType, replaceProfileOptionsControls, replaceProfileOptionsDefaultParameters)
				replaceProfileOptionsModel.SetProfilesID("testString")
				replaceProfileOptionsModel.SetProfileName("test_profile1")
				replaceProfileOptionsModel.SetProfileDescription("test_description1")
				replaceProfileOptionsModel.SetProfileType("custom")
				replaceProfileOptionsModel.SetControls([]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel})
				replaceProfileOptionsModel.SetDefaultParameters([]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel})
				replaceProfileOptionsModel.SetXCorrelationID("testString")
				replaceProfileOptionsModel.SetXRequestID("testString")
				replaceProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProfileOptionsModel).ToNot(BeNil())
				Expect(replaceProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("test_profile1")))
				Expect(replaceProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("test_description1")))
				Expect(replaceProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("custom")))
				Expect(replaceProfileOptionsModel.Controls).To(Equal([]securityandcompliancecenterapiv3.ProfileControlsPrototype{*profileControlsPrototypeModel}))
				Expect(replaceProfileOptionsModel.DefaultParameters).To(Equal([]securityandcompliancecenterapiv3.DefaultParametersPrototype{*defaultParametersPrototypeModel}))
				Expect(replaceProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceRuleOptions successfully`, func() {
				// Construct an instance of the AdditionalTargetAttribute model
				additionalTargetAttributeModel := new(securityandcompliancecenterapiv3.AdditionalTargetAttribute)
				Expect(additionalTargetAttributeModel).ToNot(BeNil())
				additionalTargetAttributeModel.Name = core.StringPtr("location")
				additionalTargetAttributeModel.Operator = core.StringPtr("string_equals")
				additionalTargetAttributeModel.Value = core.StringPtr("us-south")
				Expect(additionalTargetAttributeModel.Name).To(Equal(core.StringPtr("location")))
				Expect(additionalTargetAttributeModel.Operator).To(Equal(core.StringPtr("string_equals")))
				Expect(additionalTargetAttributeModel.Value).To(Equal(core.StringPtr("us-south")))

				// Construct an instance of the Target model
				targetModel := new(securityandcompliancecenterapiv3.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.ServiceName = core.StringPtr("cloud-object-storage")
				targetModel.ServiceDisplayName = core.StringPtr("Cloud Object Storage")
				targetModel.ResourceKind = core.StringPtr("bucket")
				targetModel.AdditionalTargetAttributes = []securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}
				Expect(targetModel.ServiceName).To(Equal(core.StringPtr("cloud-object-storage")))
				Expect(targetModel.ServiceDisplayName).To(Equal(core.StringPtr("Cloud Object Storage")))
				Expect(targetModel.ResourceKind).To(Equal(core.StringPtr("bucket")))
				Expect(targetModel.AdditionalTargetAttributes).To(Equal([]securityandcompliancecenterapiv3.AdditionalTargetAttribute{*additionalTargetAttributeModel}))

				// Construct an instance of the RequiredConfigItemsRequiredConfigBase model
				requiredConfigItemsModel := new(securityandcompliancecenterapiv3.RequiredConfigItemsRequiredConfigBase)
				Expect(requiredConfigItemsModel).ToNot(BeNil())
				requiredConfigItemsModel.Description = core.StringPtr("testString")
				requiredConfigItemsModel.Property = core.StringPtr("hard_quota")
				requiredConfigItemsModel.Operator = core.StringPtr("num_equals")
				requiredConfigItemsModel.Value = core.StringPtr("${hard_quota}")
				Expect(requiredConfigItemsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(requiredConfigItemsModel.Property).To(Equal(core.StringPtr("hard_quota")))
				Expect(requiredConfigItemsModel.Operator).To(Equal(core.StringPtr("num_equals")))
				Expect(requiredConfigItemsModel.Value).To(Equal(core.StringPtr("${hard_quota}")))

				// Construct an instance of the RequiredConfigRequiredConfigAnd model
				requiredConfigModel := new(securityandcompliancecenterapiv3.RequiredConfigRequiredConfigAnd)
				Expect(requiredConfigModel).ToNot(BeNil())
				requiredConfigModel.Description = core.StringPtr("The Cloud Object Storage rule.")
				requiredConfigModel.And = []securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}
				Expect(requiredConfigModel.Description).To(Equal(core.StringPtr("The Cloud Object Storage rule.")))
				Expect(requiredConfigModel.And).To(Equal([]securityandcompliancecenterapiv3.RequiredConfigItemsIntf{requiredConfigItemsModel}))

				// Construct an instance of the Parameter model
				parameterModel := new(securityandcompliancecenterapiv3.Parameter)
				Expect(parameterModel).ToNot(BeNil())
				parameterModel.Name = core.StringPtr("hard_quota")
				parameterModel.DisplayName = core.StringPtr("The Cloud Object Storage bucket quota.")
				parameterModel.Description = core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")
				parameterModel.Type = core.StringPtr("numeric")
				Expect(parameterModel.Name).To(Equal(core.StringPtr("hard_quota")))
				Expect(parameterModel.DisplayName).To(Equal(core.StringPtr("The Cloud Object Storage bucket quota.")))
				Expect(parameterModel.Description).To(Equal(core.StringPtr("The maximum bytes that are allocated to the Cloud Object Storage bucket.")))
				Expect(parameterModel.Type).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Import model
				importModel := new(securityandcompliancecenterapiv3.Import)
				Expect(importModel).ToNot(BeNil())
				importModel.Parameters = []securityandcompliancecenterapiv3.Parameter{*parameterModel}
				Expect(importModel.Parameters).To(Equal([]securityandcompliancecenterapiv3.Parameter{*parameterModel}))

				// Construct an instance of the ReplaceRuleOptions model
				ruleID := "testString"
				ifMatch := "testString"
				replaceRuleOptionsDescription := "Example rule"
				var replaceRuleOptionsTarget *securityandcompliancecenterapiv3.Target = nil
				var replaceRuleOptionsRequiredConfig securityandcompliancecenterapiv3.RequiredConfigIntf = nil
				replaceRuleOptionsModel := securityAndComplianceCenterApiService.NewReplaceRuleOptions(ruleID, ifMatch, replaceRuleOptionsDescription, replaceRuleOptionsTarget, replaceRuleOptionsRequiredConfig)
				replaceRuleOptionsModel.SetRuleID("testString")
				replaceRuleOptionsModel.SetIfMatch("testString")
				replaceRuleOptionsModel.SetDescription("Example rule")
				replaceRuleOptionsModel.SetTarget(targetModel)
				replaceRuleOptionsModel.SetRequiredConfig(requiredConfigModel)
				replaceRuleOptionsModel.SetType("user_defined")
				replaceRuleOptionsModel.SetVersion("1.0.1")
				replaceRuleOptionsModel.SetImport(importModel)
				replaceRuleOptionsModel.SetLabels([]string{})
				replaceRuleOptionsModel.SetXCorrelationID("testString")
				replaceRuleOptionsModel.SetXRequestID("testString")
				replaceRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRuleOptionsModel).ToNot(BeNil())
				Expect(replaceRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.Description).To(Equal(core.StringPtr("Example rule")))
				Expect(replaceRuleOptionsModel.Target).To(Equal(targetModel))
				Expect(replaceRuleOptionsModel.RequiredConfig).To(Equal(requiredConfigModel))
				Expect(replaceRuleOptionsModel.Type).To(Equal(core.StringPtr("user_defined")))
				Expect(replaceRuleOptionsModel.Version).To(Equal(core.StringPtr("1.0.1")))
				Expect(replaceRuleOptionsModel.Import).To(Equal(importModel))
				Expect(replaceRuleOptionsModel.Labels).To(Equal([]string{}))
				Expect(replaceRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTarget successfully`, func() {
				serviceName := "testString"
				resourceKind := "testString"
				_model, err := securityAndComplianceCenterApiService.NewTarget(serviceName, resourceKind)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateProviderTypeInstanceOptions successfully`, func() {
				// Construct an instance of the UpdateProviderTypeInstanceOptions model
				providerTypeID := "testString"
				providerTypeInstanceID := "testString"
				updateProviderTypeInstanceOptionsModel := securityAndComplianceCenterApiService.NewUpdateProviderTypeInstanceOptions(providerTypeID, providerTypeInstanceID)
				updateProviderTypeInstanceOptionsModel.SetProviderTypeID("testString")
				updateProviderTypeInstanceOptionsModel.SetProviderTypeInstanceID("testString")
				updateProviderTypeInstanceOptionsModel.SetName("workload-protection-instance-1")
				updateProviderTypeInstanceOptionsModel.SetAttributes(map[string]interface{}{"anyKey": "anyValue"})
				updateProviderTypeInstanceOptionsModel.SetXCorrelationID("testString")
				updateProviderTypeInstanceOptionsModel.SetXRequestID("testString")
				updateProviderTypeInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProviderTypeInstanceOptionsModel).ToNot(BeNil())
				Expect(updateProviderTypeInstanceOptionsModel.ProviderTypeID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderTypeInstanceOptionsModel.ProviderTypeInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderTypeInstanceOptionsModel.Name).To(Equal(core.StringPtr("workload-protection-instance-1")))
				Expect(updateProviderTypeInstanceOptionsModel.Attributes).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateProviderTypeInstanceOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderTypeInstanceOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(updateProviderTypeInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSettingsOptions successfully`, func() {
				// Construct an instance of the EventNotifications model
				eventNotificationsModel := new(securityandcompliancecenterapiv3.EventNotifications)
				Expect(eventNotificationsModel).ToNot(BeNil())
				eventNotificationsModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				eventNotificationsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventNotificationsModel.SourceID = core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")
				eventNotificationsModel.SourceDescription = core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")
				eventNotificationsModel.SourceName = core.StringPtr("compliance")
				Expect(eventNotificationsModel.InstanceCrn).To(Equal(core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")))
				Expect(eventNotificationsModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(eventNotificationsModel.SourceID).To(Equal(core.StringPtr("crn:v1:staging:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::")))
				Expect(eventNotificationsModel.SourceDescription).To(Equal(core.StringPtr("This source is used for integration with IBM Cloud Security and Compliance Center.")))
				Expect(eventNotificationsModel.SourceName).To(Equal(core.StringPtr("compliance")))

				// Construct an instance of the ObjectStorage model
				objectStorageModel := new(securityandcompliancecenterapiv3.ObjectStorage)
				Expect(objectStorageModel).ToNot(BeNil())
				objectStorageModel.InstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")
				objectStorageModel.Bucket = core.StringPtr("px-scan-results")
				objectStorageModel.BucketLocation = core.StringPtr("us-south")
				objectStorageModel.BucketEndpoint = core.StringPtr("testString")
				objectStorageModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				Expect(objectStorageModel.InstanceCrn).To(Equal(core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::")))
				Expect(objectStorageModel.Bucket).To(Equal(core.StringPtr("px-scan-results")))
				Expect(objectStorageModel.BucketLocation).To(Equal(core.StringPtr("us-south")))
				Expect(objectStorageModel.BucketEndpoint).To(Equal(core.StringPtr("testString")))
				Expect(objectStorageModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := securityAndComplianceCenterApiService.NewUpdateSettingsOptions()
				updateSettingsOptionsModel.SetEventNotifications(eventNotificationsModel)
				updateSettingsOptionsModel.SetObjectStorage(objectStorageModel)
				updateSettingsOptionsModel.SetXCorrelationID("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")
				updateSettingsOptionsModel.SetXRequestID("testString")
				updateSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSettingsOptionsModel).ToNot(BeNil())
				Expect(updateSettingsOptionsModel.EventNotifications).To(Equal(eventNotificationsModel))
				Expect(updateSettingsOptionsModel.ObjectStorage).To(Equal(objectStorageModel))
				Expect(updateSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("1a2b3c4d-5e6f-4a7b-8c9d-e0f1a2b3c4d5")))
				Expect(updateSettingsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(updateSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRequiredConfigItemsRequiredConfigBase successfully`, func() {
				property := "testString"
				operator := "string_equals"
				_model, err := securityAndComplianceCenterApiService.NewRequiredConfigItemsRequiredConfigBase(property, operator)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRequiredConfigRequiredConfigBase successfully`, func() {
				property := "testString"
				operator := "string_equals"
				_model, err := securityAndComplianceCenterApiService.NewRequiredConfigRequiredConfigBase(property, operator)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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

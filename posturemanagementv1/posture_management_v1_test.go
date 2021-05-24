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

package posturemanagementv1_test

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
	"github.com/go-openapi/strfmt"
	"github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe(`PostureManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(postureManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(postureManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
				URL: "https://posturemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(postureManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POSTURE_MANAGEMENT_URL": "https://posturemanagementv1/api",
				"POSTURE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
				})
				Expect(postureManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := postureManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != postureManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(postureManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(postureManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(postureManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := postureManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != postureManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(postureManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(postureManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
				})
				err := postureManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := postureManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != postureManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(postureManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(postureManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POSTURE_MANAGEMENT_URL": "https://posturemanagementv1/api",
				"POSTURE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(postureManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POSTURE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1UsingExternalConfig(&posturemanagementv1.PostureManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(postureManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = posturemanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListScans(listScansOptions *ListScansOptions) - Operation response error`, func() {
		listScansPath := "/posture/v1/scans/validations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScansPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListScans with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScansOptions model
				listScansOptionsModel := new(posturemanagementv1.ListScansOptions)
				listScansOptionsModel.AccountID = core.StringPtr("testString")
				listScansOptionsModel.TransactionID = core.StringPtr("testString")
				listScansOptionsModel.Name = core.StringPtr("testString")
				listScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListScans(listScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListScans(listScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListScans(listScansOptions *ListScansOptions)`, func() {
		listScansPath := "/posture/v1/scans/validations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scans": [{"scan_id": "262", "scan_name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "profile_id": "48", "profile_name": "IBM Cloud Best Practices Controls 1.0", "profile_type": "standard", "profile_group_id": "1", "profile_group_name": "CIS Windows Server Benchmarks", "report_run_by": "controller", "scan_time": "2020-09-23T12:45:24.000Z", "started_time": "2020-09-23T12:45:24.000Z", "end_time": "2020-09-23T12:45:24.000Z", "result": {"goals_pass_count": 118, "goals_u2p_count": 16, "goals_na_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_na_count": 6, "controls_u2p_count": 16, "controls_total_count": 293}}]}`)
				}))
			})
			It(`Invoke ListScans successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListScansOptions model
				listScansOptionsModel := new(posturemanagementv1.ListScansOptions)
				listScansOptionsModel.AccountID = core.StringPtr("testString")
				listScansOptionsModel.TransactionID = core.StringPtr("testString")
				listScansOptionsModel.Name = core.StringPtr("testString")
				listScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListScansWithContext(ctx, listScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListScans(listScansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListScansWithContext(ctx, listScansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listScansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scans": [{"scan_id": "262", "scan_name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "profile_id": "48", "profile_name": "IBM Cloud Best Practices Controls 1.0", "profile_type": "standard", "profile_group_id": "1", "profile_group_name": "CIS Windows Server Benchmarks", "report_run_by": "controller", "scan_time": "2020-09-23T12:45:24.000Z", "started_time": "2020-09-23T12:45:24.000Z", "end_time": "2020-09-23T12:45:24.000Z", "result": {"goals_pass_count": 118, "goals_u2p_count": 16, "goals_na_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_na_count": 6, "controls_u2p_count": 16, "controls_total_count": 293}}]}`)
				}))
			})
			It(`Invoke ListScans successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListScans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListScansOptions model
				listScansOptionsModel := new(posturemanagementv1.ListScansOptions)
				listScansOptionsModel.AccountID = core.StringPtr("testString")
				listScansOptionsModel.TransactionID = core.StringPtr("testString")
				listScansOptionsModel.Name = core.StringPtr("testString")
				listScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListScans(listScansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListScans with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScansOptions model
				listScansOptionsModel := new(posturemanagementv1.ListScansOptions)
				listScansOptionsModel.AccountID = core.StringPtr("testString")
				listScansOptionsModel.TransactionID = core.StringPtr("testString")
				listScansOptionsModel.Name = core.StringPtr("testString")
				listScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListScans(listScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListScansOptions model with no property values
				listScansOptionsModelNew := new(posturemanagementv1.ListScansOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ListScans(listScansOptionsModelNew)
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
			It(`Invoke ListScans successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScansOptions model
				listScansOptionsModel := new(posturemanagementv1.ListScansOptions)
				listScansOptionsModel.AccountID = core.StringPtr("testString")
				listScansOptionsModel.TransactionID = core.StringPtr("testString")
				listScansOptionsModel.Name = core.StringPtr("testString")
				listScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListScans(listScansOptionsModel)
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
	Describe(`CreateValidation(createValidationOptions *CreateValidationOptions) - Operation response error`, func() {
		createValidationPath := "/posture/v1/scans/validations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValidationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateValidation with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv1.CreateValidationOptions)
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateValidation(createValidationOptions *CreateValidationOptions)`, func() {
		createValidationPath := "/posture/v1/scans/validations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValidationPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"result": true, "message": "Success: The validation is in progress. To see the results, go to Security & Compliance > Assess > Scans in the service dashboard and select the scan My_Example_scan."}`)
				}))
			})
			It(`Invoke CreateValidation successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv1.CreateValidationOptions)
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.CreateValidationWithContext(ctx, createValidationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.CreateValidationWithContext(ctx, createValidationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createValidationPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"result": true, "message": "Success: The validation is in progress. To see the results, go to Security & Compliance > Assess > Scans in the service dashboard and select the scan My_Example_scan."}`)
				}))
			})
			It(`Invoke CreateValidation successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.CreateValidation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv1.CreateValidationOptions)
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateValidation with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv1.CreateValidationOptions)
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateValidationOptions model with no property values
				createValidationOptionsModelNew := new(posturemanagementv1.CreateValidationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.CreateValidation(createValidationOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateValidation successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv1.CreateValidationOptions)
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.CreateValidation(createValidationOptionsModel)
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
		listProfilesPath := "/posture/v1/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListProfiles(listProfilesOptionsModel)
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
		listProfilesPath := "/posture/v1/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "CIS IBM Foundations Benchmark 1.0.0", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "profile_id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "profile_type": "predefined", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "CIS IBM Foundations Benchmark 1.0.0", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "profile_id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "profile_type": "predefined", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProfilesOptions model with no property values
				listProfilesOptionsModelNew := new(posturemanagementv1.ListProfilesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ListProfiles(listProfilesOptionsModelNew)
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
			It(`Invoke ListProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv1.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Name = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListProfiles(listProfilesOptionsModel)
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
	Describe(`CreateScope(createScopeOptions *CreateScopeOptions) - Operation response error`, func() {
		createScopePath := "/posture/v1/scopes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createScopePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateScope with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv1.CreateScopeOptions)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.ScopeName = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.ScopeDescription = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.EnvironmentType = core.StringPtr("ibm")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateScope(createScopeOptions *CreateScopeOptions)`, func() {
		createScopePath := "/posture/v1/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createScopePath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"scope_id": "My_Example_Scope", "scope_name": "My_Example_Scope", "scope_description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "collector_ids": ["CollectorIds"], "credential_id": "IBM Cloud Tn03", "environment_type": "ibm", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z"}`)
				}))
			})
			It(`Invoke CreateScope successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv1.CreateScopeOptions)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.ScopeName = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.ScopeDescription = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.EnvironmentType = core.StringPtr("ibm")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.CreateScopeWithContext(ctx, createScopeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.CreateScopeWithContext(ctx, createScopeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createScopePath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"scope_id": "My_Example_Scope", "scope_name": "My_Example_Scope", "scope_description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "collector_ids": ["CollectorIds"], "credential_id": "IBM Cloud Tn03", "environment_type": "ibm", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z"}`)
				}))
			})
			It(`Invoke CreateScope successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.CreateScope(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv1.CreateScopeOptions)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.ScopeName = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.ScopeDescription = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.EnvironmentType = core.StringPtr("ibm")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateScope with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv1.CreateScopeOptions)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.ScopeName = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.ScopeDescription = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.EnvironmentType = core.StringPtr("ibm")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateScopeOptions model with no property values
				createScopeOptionsModelNew := new(posturemanagementv1.CreateScopeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.CreateScope(createScopeOptionsModelNew)
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
			It(`Invoke CreateScope successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv1.CreateScopeOptions)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.ScopeName = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.ScopeDescription = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.EnvironmentType = core.StringPtr("ibm")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.CreateScope(createScopeOptionsModel)
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
	Describe(`ListScopes(listScopesOptions *ListScopesOptions) - Operation response error`, func() {
		listScopesPath := "/posture/v1/scopes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListScopes with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {
		listScopesPath := "/posture/v1/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "scope_id": "1", "name": "My_Example_Scope", "enabled": true, "environment_type": "ibm", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z", "last_scan_type": "fact_collection", "last_scan_type_description": "Fact collection", "last_scan_status_updated_time": "2021-02-26T04:07:25.000Z", "collectors_id": ["CollectorsID"], "scans": [{"scan_id": "235", "discover_id": "49", "status": "validation_completed", "status_message": "The collector aborted the task during upgrade."}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListScopesWithContext(ctx, listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListScopesWithContext(ctx, listScopesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "scope_id": "1", "name": "My_Example_Scope", "enabled": true, "environment_type": "ibm", "created_time": "2021-02-26T04:07:25.000Z", "modified_time": "2021-02-26T04:07:25.000Z", "last_scan_type": "fact_collection", "last_scan_type_description": "Fact collection", "last_scan_status_updated_time": "2021-02-26T04:07:25.000Z", "collectors_id": ["CollectorsID"], "scans": [{"scan_id": "235", "discover_id": "49", "status": "validation_completed", "status_message": "The collector aborted the task during upgrade."}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListScopes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListScopes with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListScopesOptions model with no property values
				listScopesOptionsModelNew := new(posturemanagementv1.ListScopesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ListScopes(listScopesOptionsModelNew)
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
			It(`Invoke ListScopes successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListScopes(listScopesOptionsModel)
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
	Describe(`CreateCollector(createCollectorOptions *CreateCollectorOptions) - Operation response error`, func() {
		createCollectorPath := "/posture/v1/collectors"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCollector with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv1.CreateCollectorOptions)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.CollectorName = core.StringPtr("IBMSchema-new-048-test")
				createCollectorOptionsModel.CollectorDescription = core.StringPtr("IBMSchema")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("IBM")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCollector(createCollectorOptions *CreateCollectorOptions)`, func() {
		createCollectorPath := "/posture/v1/collectors"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"collector_id": "1"}`)
				}))
			})
			It(`Invoke CreateCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv1.CreateCollectorOptions)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.CollectorName = core.StringPtr("IBMSchema-new-048-test")
				createCollectorOptionsModel.CollectorDescription = core.StringPtr("IBMSchema")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("IBM")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.CreateCollectorWithContext(ctx, createCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.CreateCollectorWithContext(ctx, createCollectorOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"collector_id": "1"}`)
				}))
			})
			It(`Invoke CreateCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.CreateCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv1.CreateCollectorOptions)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.CollectorName = core.StringPtr("IBMSchema-new-048-test")
				createCollectorOptionsModel.CollectorDescription = core.StringPtr("IBMSchema")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("IBM")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv1.CreateCollectorOptions)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.CollectorName = core.StringPtr("IBMSchema-new-048-test")
				createCollectorOptionsModel.CollectorDescription = core.StringPtr("IBMSchema")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("IBM")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCollectorOptions model with no property values
				createCollectorOptionsModelNew := new(posturemanagementv1.CreateCollectorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.CreateCollector(createCollectorOptionsModelNew)
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
			It(`Invoke CreateCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv1.CreateCollectorOptions)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.CollectorName = core.StringPtr("IBMSchema-new-048-test")
				createCollectorOptionsModel.CollectorDescription = core.StringPtr("IBMSchema")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("IBM")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.CreateCollector(createCollectorOptionsModel)
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
	Describe(`CreateCredential(createCredentialOptions *CreateCredentialOptions) - Operation response error`, func() {
		createCredentialPath := "/posture/v1/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCredential with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv1.CreateCredentialOptions)
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.CredentialDataFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.PemFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCredential(createCredentialOptions *CreateCredentialOptions)`, func() {
		createCredentialPath := "/posture/v1/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"credential_id": "1", "credential_name": "aws cred", "created_time": "2021-04-08T10:22:04.000Z"}`)
				}))
			})
			It(`Invoke CreateCredential successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv1.CreateCredentialOptions)
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.CredentialDataFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.PemFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.CreateCredentialWithContext(ctx, createCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.CreateCredentialWithContext(ctx, createCredentialOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"credential_id": "1", "credential_name": "aws cred", "created_time": "2021-04-08T10:22:04.000Z"}`)
				}))
			})
			It(`Invoke CreateCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.CreateCredential(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv1.CreateCredentialOptions)
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.CredentialDataFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.PemFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCredential with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv1.CreateCredentialOptions)
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.CredentialDataFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.PemFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCredentialOptions model with no property values
				createCredentialOptionsModelNew := new(posturemanagementv1.CreateCredentialOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.CreateCredential(createCredentialOptionsModelNew)
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
			It(`Invoke CreateCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv1.CreateCredentialOptions)
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.CredentialDataFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.PemFile = CreateMockReader("This is a mock file.")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.CreateCredential(createCredentialOptionsModel)
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
			postureManagementService, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
				URL:           "http://posturemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateCollectorOptions successfully`, func() {
				// Construct an instance of the CreateCollectorOptions model
				accountID := "testString"
				createCollectorOptionsModel := postureManagementService.NewCreateCollectorOptions(accountID)
				createCollectorOptionsModel.SetAccountID("testString")
				createCollectorOptionsModel.SetCollectorName("IBMSchema-new-048-test")
				createCollectorOptionsModel.SetCollectorDescription("IBMSchema")
				createCollectorOptionsModel.SetIsPublic(true)
				createCollectorOptionsModel.SetManagedBy("IBM")
				createCollectorOptionsModel.SetTransactionID("testString")
				createCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectorOptionsModel).ToNot(BeNil())
				Expect(createCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectorOptionsModel.CollectorName).To(Equal(core.StringPtr("IBMSchema-new-048-test")))
				Expect(createCollectorOptionsModel.CollectorDescription).To(Equal(core.StringPtr("IBMSchema")))
				Expect(createCollectorOptionsModel.IsPublic).To(Equal(core.BoolPtr(true)))
				Expect(createCollectorOptionsModel.ManagedBy).To(Equal(core.StringPtr("IBM")))
				Expect(createCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCredentialOptions successfully`, func() {
				// Construct an instance of the CreateCredentialOptions model
				accountID := "testString"
				credentialDataFile := CreateMockReader("This is a mock file.")
				createCredentialOptionsModel := postureManagementService.NewCreateCredentialOptions(accountID, credentialDataFile)
				createCredentialOptionsModel.SetAccountID("testString")
				createCredentialOptionsModel.SetCredentialDataFile(CreateMockReader("This is a mock file."))
				createCredentialOptionsModel.SetPemFile(CreateMockReader("This is a mock file."))
				createCredentialOptionsModel.SetTransactionID("testString")
				createCredentialOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCredentialOptionsModel).ToNot(BeNil())
				Expect(createCredentialOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createCredentialOptionsModel.CredentialDataFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createCredentialOptionsModel.PemFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createCredentialOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createCredentialOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateScopeOptions successfully`, func() {
				// Construct an instance of the CreateScopeOptions model
				accountID := "testString"
				createScopeOptionsModel := postureManagementService.NewCreateScopeOptions(accountID)
				createScopeOptionsModel.SetAccountID("testString")
				createScopeOptionsModel.SetScopeName("IBMSchema-new-048-test")
				createScopeOptionsModel.SetScopeDescription("IBMSchema")
				createScopeOptionsModel.SetCollectorIds([]string{"20"})
				createScopeOptionsModel.SetCredentialID("5")
				createScopeOptionsModel.SetEnvironmentType("ibm")
				createScopeOptionsModel.SetTransactionID("testString")
				createScopeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createScopeOptionsModel).ToNot(BeNil())
				Expect(createScopeOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createScopeOptionsModel.ScopeName).To(Equal(core.StringPtr("IBMSchema-new-048-test")))
				Expect(createScopeOptionsModel.ScopeDescription).To(Equal(core.StringPtr("IBMSchema")))
				Expect(createScopeOptionsModel.CollectorIds).To(Equal([]string{"20"}))
				Expect(createScopeOptionsModel.CredentialID).To(Equal(core.StringPtr("5")))
				Expect(createScopeOptionsModel.EnvironmentType).To(Equal(core.StringPtr("ibm")))
				Expect(createScopeOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createScopeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateValidationOptions successfully`, func() {
				// Construct an instance of the CreateValidationOptions model
				accountID := "testString"
				createValidationOptionsModel := postureManagementService.NewCreateValidationOptions(accountID)
				createValidationOptionsModel.SetAccountID("testString")
				createValidationOptionsModel.SetScopeID("1")
				createValidationOptionsModel.SetProfileID("6")
				createValidationOptionsModel.SetGroupProfileID("13")
				createValidationOptionsModel.SetTransactionID("testString")
				createValidationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createValidationOptionsModel).ToNot(BeNil())
				Expect(createValidationOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createValidationOptionsModel.ScopeID).To(Equal(core.StringPtr("1")))
				Expect(createValidationOptionsModel.ProfileID).To(Equal(core.StringPtr("6")))
				Expect(createValidationOptionsModel.GroupProfileID).To(Equal(core.StringPtr("13")))
				Expect(createValidationOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createValidationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				accountID := "testString"
				listProfilesOptionsModel := postureManagementService.NewListProfilesOptions(accountID)
				listProfilesOptionsModel.SetAccountID("testString")
				listProfilesOptionsModel.SetTransactionID("testString")
				listProfilesOptionsModel.SetName("testString")
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListScansOptions successfully`, func() {
				// Construct an instance of the ListScansOptions model
				accountID := "testString"
				listScansOptionsModel := postureManagementService.NewListScansOptions(accountID)
				listScansOptionsModel.SetAccountID("testString")
				listScansOptionsModel.SetTransactionID("testString")
				listScansOptionsModel.SetName("testString")
				listScansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listScansOptionsModel).ToNot(BeNil())
				Expect(listScansOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listScansOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listScansOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listScansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListScopesOptions successfully`, func() {
				// Construct an instance of the ListScopesOptions model
				accountID := "testString"
				listScopesOptionsModel := postureManagementService.NewListScopesOptions(accountID)
				listScopesOptionsModel.SetAccountID("testString")
				listScopesOptionsModel.SetTransactionID("testString")
				listScopesOptionsModel.SetName("testString")
				listScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listScopesOptionsModel).ToNot(BeNil())
				Expect(listScopesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

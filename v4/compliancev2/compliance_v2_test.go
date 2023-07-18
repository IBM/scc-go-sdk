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

package compliancev2_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/compliancev2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ComplianceV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(complianceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(complianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
				URL: "https://compliancev2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(complianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPLIANCE_URL":       "https://compliancev2/api",
				"COMPLIANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				complianceService, serviceErr := compliancev2.NewComplianceV2UsingExternalConfig(&compliancev2.ComplianceV2Options{})
				Expect(complianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := complianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != complianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(complianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(complianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				complianceService, serviceErr := compliancev2.NewComplianceV2UsingExternalConfig(&compliancev2.ComplianceV2Options{
					URL: "https://testService/api",
				})
				Expect(complianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(complianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := complianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != complianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(complianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(complianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				complianceService, serviceErr := compliancev2.NewComplianceV2UsingExternalConfig(&compliancev2.ComplianceV2Options{})
				err := complianceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(complianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := complianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != complianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(complianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(complianceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPLIANCE_URL":       "https://compliancev2/api",
				"COMPLIANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			complianceService, serviceErr := compliancev2.NewComplianceV2UsingExternalConfig(&compliancev2.ComplianceV2Options{})

			It(`Instantiate service client with error`, func() {
				Expect(complianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"COMPLIANCE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			complianceService, serviceErr := compliancev2.NewComplianceV2UsingExternalConfig(&compliancev2.ComplianceV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(complianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = compliancev2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := compliancev2.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.compliance.cloud.ibm.com/instances/edf9524f-406c-412c-acbb-ee371a5cabda/v3"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := compliancev2.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
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
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListControlLibraries with error: Operation response processing error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(compliancev2.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
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
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(compliancev2.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
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
					Expect(req.URL.Query()["control_library_type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ListControlLibraries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(compliancev2.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListControlLibraries with error: Operation request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(compliancev2.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(compliancev2.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.XRequestID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listControlLibrariesOptionsModel.ControlLibraryType = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ListControlLibraries(listControlLibrariesOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(compliancev2.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(compliancev2.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.CreateCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(compliancev2.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomControlLibrary with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(compliancev2.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomControlLibraryOptions model with no property values
				createCustomControlLibraryOptionsModelNew := new(compliancev2.CreateCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(compliancev2.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				createCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("1.1.0")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(compliancev2.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.GetControlLibrary(getControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(compliancev2.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.GetControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(compliancev2.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetControlLibrary with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(compliancev2.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetControlLibraryOptions model with no property values
				getControlLibraryOptionsModelNew := new(compliancev2.GetControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.GetControlLibrary(getControlLibraryOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(compliancev2.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				getControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.GetControlLibrary(getControlLibraryOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(compliancev2.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(compliancev2.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f3517159-889e-4781-819a-89d89b747c85", "account_id": "130003ea8bfa43c5aacea07a86da3000", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "control_library_version": "ControlLibraryVersion", "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "latest": true, "hierarchy_enabled": true, "controls_count": 13, "control_parents_count": 19, "controls": [{"control_name": "ControlName", "control_id": "1fa45e17-9322-4e6c-bbd6-1c51db08e790", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_tags": ["ControlTags"], "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ReplaceCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(compliancev2.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceCustomControlLibrary with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(compliancev2.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceCustomControlLibraryOptions model with no property values
				replaceCustomControlLibraryOptionsModelNew := new(compliancev2.ReplaceCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(compliancev2.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("custom")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceCustomControlLibraryOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.HierarchyEnabled = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.ControlParentsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []compliancev2.ControlsInControlLib{*controlsInControlLibModel}
				replaceCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(compliancev2.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(compliancev2.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.DeleteCustomControlLibraryWithContext(ctx, deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.DeleteCustomControlLibraryWithContext(ctx, deleteCustomControlLibraryOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.DeleteCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(compliancev2.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomControlLibrary with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(compliancev2.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomControlLibraryOptions model with no property values
				deleteCustomControlLibraryOptionsModelNew := new(compliancev2.DeleteCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControlLibraryOptions model
				deleteCustomControlLibraryOptionsModel := new(compliancev2.DeleteCustomControlLibraryOptions)
				deleteCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.DeleteCustomControlLibrary(deleteCustomControlLibraryOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(compliancev2.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listProfilesOptionsModel.ProfileType = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ListProfiles(listProfilesOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(compliancev2.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listProfilesOptionsModel.ProfileType = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ListProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(compliancev2.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listProfilesOptionsModel.ProfileType = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(compliancev2.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listProfilesOptionsModel.ProfileType = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ListProfiles(listProfilesOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(compliancev2.ListProfilesOptions)
				listProfilesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listProfilesOptionsModel.XRequestID = core.StringPtr("testString")
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listProfilesOptionsModel.ProfileType = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ListProfiles(listProfilesOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(compliancev2.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.CreateProfile(createProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(compliancev2.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.CreateProfileWithContext(ctx, createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.CreateProfileWithContext(ctx, createProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.CreateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(compliancev2.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfile with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(compliancev2.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProfileOptions model with no property values
				createProfileOptionsModelNew := new(compliancev2.CreateProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.CreateProfile(createProfileOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(compliancev2.CreateProfileOptions)
				createProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				createProfileOptionsModel.ProfileType = core.StringPtr("custom")
				createProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				createProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				createProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				createProfileOptionsModel.XRequestID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.CreateProfile(createProfileOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(compliancev2.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.GetProfile(getProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(compliancev2.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.GetProfileWithContext(ctx, getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.GetProfileWithContext(ctx, getProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.GetProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(compliancev2.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(compliancev2.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfileOptionsModelNew := new(compliancev2.GetProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.GetProfile(getProfileOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(compliancev2.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.GetProfile(getProfileOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(compliancev2.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ReplaceProfile(replaceProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke ReplaceProfile successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(compliancev2.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ReplaceProfileWithContext(ctx, replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ReplaceProfileWithContext(ctx, replaceProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke ReplaceProfile successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ReplaceProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(compliancev2.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfile with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(compliancev2.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ReplaceProfile(replaceProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileOptions model with no property values
				replaceProfileOptionsModelNew := new(compliancev2.ReplaceProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.ReplaceProfile(replaceProfileOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
				defaultParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				defaultParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				defaultParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				defaultParametersPrototypeModel.ParameterDefaultValue = core.StringPtr("120")
				defaultParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				defaultParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileOptions model
				replaceProfileOptionsModel := new(compliancev2.ReplaceProfileOptions)
				replaceProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileOptionsModel.ProfileName = core.StringPtr("test_profile1")
				replaceProfileOptionsModel.ProfileDescription = core.StringPtr("test_description1")
				replaceProfileOptionsModel.ProfileType = core.StringPtr("custom")
				replaceProfileOptionsModel.Controls = []compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}
				replaceProfileOptionsModel.DefaultParameters = []compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}
				replaceProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ReplaceProfile(replaceProfileOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(compliancev2.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(compliancev2.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "predefined", "profile_version": "ProfileVersion", "version_group_label": "e0923045-f00d-44de-b49b-6f1f0e8033cc", "instance_id": "InstanceID", "latest": true, "hierarchy_enabled": true, "created_by": "CreatedBy", "created_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "controls_count": 13, "control_parents_count": 19, "attachments_count": 16, "controls": [{"control_library_id": "e98a56ff-dc24-41d4-9875-1e188e2da6cd", "control_id": "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_requirement": true, "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications_count": 26, "control_specifications": [{"control_specification_id": "f3517159-889e-4781-819a-89d89b747c85", "responsibility": "user", "component_id": "f3517159-889e-4781-819a-89d89b747c85", "componenet_name": "ComponenetName", "environment": "Environment", "control_specification_description": "ControlSpecificationDescription", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.DeleteCustomProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(compliancev2.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomProfile with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(compliancev2.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomProfileOptions model with no property values
				deleteCustomProfileOptionsModelNew := new(compliancev2.DeleteCustomProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(compliancev2.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.XRequestID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachments with error: Operation response processing error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(compliancev2.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ListAttachments(listAttachmentsOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(compliancev2.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ListAttachmentsWithContext(ctx, listAttachmentsOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachments successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ListAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(compliancev2.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachments with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(compliancev2.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ListAttachments(listAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAttachmentsOptions model with no property values
				listAttachmentsOptionsModelNew := new(compliancev2.ListAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.ListAttachments(listAttachmentsOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsOptions model
				listAttachmentsOptionsModel := new(compliancev2.ListAttachmentsOptions)
				listAttachmentsOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ListAttachments(listAttachmentsOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(compliancev2.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.CreateAttachment(createAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(compliancev2.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.CreateAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(compliancev2.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAttachment with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(compliancev2.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAttachmentOptions model with no property values
				createAttachmentOptionsModelNew := new(compliancev2.CreateAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.CreateAttachment(createAttachmentOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(compliancev2.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}
				createAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				createAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				createAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.CreateAttachment(createAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(compliancev2.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke GetProfileAttachment successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(compliancev2.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.GetProfileAttachmentWithContext(ctx, getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.GetProfileAttachmentWithContext(ctx, getProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke GetProfileAttachment successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.GetProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(compliancev2.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileAttachment with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(compliancev2.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileAttachmentOptions model with no property values
				getProfileAttachmentOptionsModelNew := new(compliancev2.GetProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.GetProfileAttachment(getProfileAttachmentOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmentOptions model
				getProfileAttachmentOptionsModel := new(compliancev2.GetProfileAttachmentOptions)
				getProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				getProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.GetProfileAttachment(getProfileAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(compliancev2.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("daily")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(compliancev2.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("daily")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ReplaceProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(compliancev2.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("daily")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfileAttachment with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(compliancev2.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("daily")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileAttachmentOptions model with no property values
				replaceProfileAttachmentOptionsModelNew := new(compliancev2.ReplaceProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(compliancev2.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ProfileID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				replaceProfileAttachmentOptionsModel.CreatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.Schedule = core.StringPtr("daily")
				replaceProfileAttachmentOptionsModel.Notifications = attachmentsNotificationsPrototypeModel
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				replaceProfileAttachmentOptionsModel.LastScan = lastScanModel
				replaceProfileAttachmentOptionsModel.NextScanTime = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				replaceProfileAttachmentOptionsModel.Description = core.StringPtr("Test description")
				replaceProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(compliancev2.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke DeleteProfileAttachment successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(compliancev2.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.DeleteProfileAttachmentWithContext(ctx, deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.DeleteProfileAttachmentWithContext(ctx, deleteProfileAttachmentOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}`)
				}))
			})
			It(`Invoke DeleteProfileAttachment successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.DeleteProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(compliancev2.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProfileAttachment with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(compliancev2.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProfileAttachmentOptions model with no property values
				deleteProfileAttachmentOptionsModelNew := new(compliancev2.DeleteProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmentOptions model
				deleteProfileAttachmentOptionsModel := new(compliancev2.DeleteProfileAttachmentOptions)
				deleteProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.XRequestID = core.StringPtr("testString")
				deleteProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.DeleteProfileAttachment(deleteProfileAttachmentOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(compliancev2.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.CreateScan(createScanOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(compliancev2.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.CreateScanWithContext(ctx, createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.CreateScanWithContext(ctx, createScanOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.CreateScan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(compliancev2.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateScan with error: Operation validation and request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(compliancev2.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateScanOptions model with no property values
				createScanOptionsModelNew := new(compliancev2.CreateScanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = complianceService.CreateScan(createScanOptionsModelNew)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(compliancev2.CreateScanOptions)
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.XCorrelationID = core.StringPtr("testString")
				createScanOptionsModel.XRequestID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.CreateScan(createScanOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachmentsAccount with error: Operation response processing error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(compliancev2.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				complianceService.EnableRetries(0, 0)
				result, response, operationErr = complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachmentsAccount successfully with retries`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())
				complianceService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(compliancev2.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := complianceService.ListAttachmentsAccountWithContext(ctx, listAttachmentsAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				complianceService.DisableRetries()
				result, response, operationErr := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = complianceService.ListAttachmentsAccountWithContext(ctx, listAttachmentsAccountOptionsModel)
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "attachments": [{"id": "130003ea8bfa43c5aacea07a86da3000", "profile_id": "7ec45986-54fc-4b66-a303-d9577b078c65", "account_id": "130003ea8bfa43c5aacea07a86da3000", "instance_id": "edf9524f-406c-412c-acbb-ee371a5cabda", "scope": [{"environment": "Environment", "properties": [{"name": "Name", "value": "Value"}]}], "created_on": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_on": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "status": "enabled", "schedule": "daily", "notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"]}}, "attachment_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_value": "ParameterValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "string"}], "last_scan": {"id": "e8a39d25-0051-4328-8462-988ad321f49a", "status": "in_progress", "time": "Time"}, "next_scan_time": "NextScanTime", "name": "account-130003ea8bfa43c5aacea07a86da3000", "description": "Test description"}]}`)
				}))
			})
			It(`Invoke ListAttachmentsAccount successfully`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := complianceService.ListAttachmentsAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(compliancev2.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachmentsAccount with error: Operation request error`, func() {
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(compliancev2.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := complianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
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
				complianceService, serviceErr := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(complianceService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := new(compliancev2.ListAttachmentsAccountOptions)
				listAttachmentsAccountOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.XRequestID = core.StringPtr("testString")
				listAttachmentsAccountOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAttachmentsAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := complianceService.ListAttachmentsAccount(listAttachmentsAccountOptionsModel)
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
			complianceService, _ := compliancev2.NewComplianceV2(&compliancev2.ComplianceV2Options{
				URL:           "http://compliancev2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAttachmentPrototype successfully`, func() {
				attachments := []compliancev2.AttachmentsPrototype{}
				_model, err := complianceService.NewAttachmentPrototype(attachments)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAttachmentsNotificationsPrototype successfully`, func() {
				enabled := true
				var controls *compliancev2.FailedControls = nil
				_, err := complianceService.NewAttachmentsNotificationsPrototype(enabled, controls)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAttachmentsPrototype successfully`, func() {
				name := "account-130003ea8bfa43c5aacea07a86da3000"
				scope := []compliancev2.MultiCloudScope{}
				status := "enabled"
				schedule := "daily"
				attachmentParameters := []compliancev2.AttachmentParametersPrototype{}
				_model, err := complianceService.NewAttachmentsPrototype(name, scope, status, schedule, attachmentParameters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateAttachmentOptions successfully`, func() {
				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				Expect(propertyModel).ToNot(BeNil())
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")
				Expect(propertyModel.Name).To(Equal(core.StringPtr("scope_id")))
				Expect(propertyModel.Value).To(Equal(core.StringPtr("cg3335893hh1428692d6747cf300yeb5")))

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				Expect(multiCloudScopeModel).ToNot(BeNil())
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}
				Expect(multiCloudScopeModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(multiCloudScopeModel.Properties).To(Equal([]compliancev2.Property{*propertyModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(15))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{}))

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				Expect(attachmentsNotificationsPrototypeModel).ToNot(BeNil())
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPrototypeModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(attachmentsNotificationsPrototypeModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				Expect(attachmentParametersPrototypeModel).ToNot(BeNil())
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(attachmentParametersPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(attachmentParametersPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(attachmentParametersPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(attachmentParametersPrototypeModel.ParameterValue).To(Equal(core.StringPtr("120")))
				Expect(attachmentParametersPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(attachmentParametersPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the AttachmentsPrototype model
				attachmentsPrototypeModel := new(compliancev2.AttachmentsPrototype)
				Expect(attachmentsPrototypeModel).ToNot(BeNil())
				attachmentsPrototypeModel.ID = core.StringPtr("130003ea8bfa43c5aacea07a86da3000")
				attachmentsPrototypeModel.Name = core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")
				attachmentsPrototypeModel.Description = core.StringPtr("Test description")
				attachmentsPrototypeModel.Scope = []compliancev2.MultiCloudScope{*multiCloudScopeModel}
				attachmentsPrototypeModel.Status = core.StringPtr("enabled")
				attachmentsPrototypeModel.Schedule = core.StringPtr("daily")
				attachmentsPrototypeModel.Notifications = attachmentsNotificationsPrototypeModel
				attachmentsPrototypeModel.AttachmentParameters = []compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}
				Expect(attachmentsPrototypeModel.ID).To(Equal(core.StringPtr("130003ea8bfa43c5aacea07a86da3000")))
				Expect(attachmentsPrototypeModel.Name).To(Equal(core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")))
				Expect(attachmentsPrototypeModel.Description).To(Equal(core.StringPtr("Test description")))
				Expect(attachmentsPrototypeModel.Scope).To(Equal([]compliancev2.MultiCloudScope{*multiCloudScopeModel}))
				Expect(attachmentsPrototypeModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(attachmentsPrototypeModel.Schedule).To(Equal(core.StringPtr("daily")))
				Expect(attachmentsPrototypeModel.Notifications).To(Equal(attachmentsNotificationsPrototypeModel))
				Expect(attachmentsPrototypeModel.AttachmentParameters).To(Equal([]compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}))

				// Construct an instance of the CreateAttachmentOptions model
				profilesID := "testString"
				createAttachmentOptionsAttachments := []compliancev2.AttachmentsPrototype{}
				createAttachmentOptionsModel := complianceService.NewCreateAttachmentOptions(profilesID, createAttachmentOptionsAttachments)
				createAttachmentOptionsModel.SetProfilesID("testString")
				createAttachmentOptionsModel.SetAttachments([]compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel})
				createAttachmentOptionsModel.SetProfileID("testString")
				createAttachmentOptionsModel.SetXCorrelationID("testString")
				createAttachmentOptionsModel.SetXRequestID("testString")
				createAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAttachmentOptionsModel).ToNot(BeNil())
				Expect(createAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Attachments).To(Equal([]compliancev2.AttachmentsPrototype{*attachmentsPrototypeModel}))
				Expect(createAttachmentOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				Expect(implementationModel).ToNot(BeNil())
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}
				Expect(implementationModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(implementationModel.AssessmentMethod).To(Equal(core.StringPtr("ibm-cloud-rule")))
				Expect(implementationModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(implementationModel.AssessmentDescription).To(Equal(core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")))
				Expect(implementationModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationModel.Parameters).To(Equal([]compliancev2.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}
				Expect(controlSpecificationsModel.ControlSpecificationID).To(Equal(core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("iam-identity")))
				Expect(controlSpecificationsModel.ComponenetName).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(controlSpecificationsModel.ControlSpecificationDescription).To(Equal(core.StringPtr("IBM cloud")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]compliancev2.Implementation{*implementationModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("sc-7")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("ibm-cloud")))

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				Expect(controlsInControlLibModel).ToNot(BeNil())
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibModel.ControlName).To(Equal(core.StringPtr("SC-7")))
				Expect(controlsInControlLibModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))
				Expect(controlsInControlLibModel.ControlDescription).To(Equal(core.StringPtr("Boundary Protection")))
				Expect(controlsInControlLibModel.ControlCategory).To(Equal(core.StringPtr("System and Communications Protection")))
				Expect(controlsInControlLibModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibModel.ControlRequirement).To(Equal(core.BoolPtr(false)))
				Expect(controlsInControlLibModel.ControlTags).To(Equal([]string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}))
				Expect(controlsInControlLibModel.ControlSpecifications).To(Equal([]compliancev2.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsControlLibraryName := "IBM Cloud for Financial Services"
				createCustomControlLibraryOptionsControlLibraryDescription := "IBM Cloud for Financial Services"
				createCustomControlLibraryOptionsControlLibraryType := "custom"
				createCustomControlLibraryOptionsControls := []compliancev2.ControlsInControlLib{}
				createCustomControlLibraryOptionsModel := complianceService.NewCreateCustomControlLibraryOptions(createCustomControlLibraryOptionsControlLibraryName, createCustomControlLibraryOptionsControlLibraryDescription, createCustomControlLibraryOptionsControlLibraryType, createCustomControlLibraryOptionsControls)
				createCustomControlLibraryOptionsModel.SetControlLibraryName("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.SetControlLibraryDescription("IBM Cloud for Financial Services")
				createCustomControlLibraryOptionsModel.SetControlLibraryType("custom")
				createCustomControlLibraryOptionsModel.SetControls([]compliancev2.ControlsInControlLib{*controlsInControlLibModel})
				createCustomControlLibraryOptionsModel.SetVersionGroupLabel("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")
				createCustomControlLibraryOptionsModel.SetControlLibraryVersion("1.1.0")
				createCustomControlLibraryOptionsModel.SetLatest(true)
				createCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				createCustomControlLibraryOptionsModel.SetXCorrelationID("testString")
				createCustomControlLibraryOptionsModel.SetXRequestID("testString")
				createCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryName).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryDescription).To(Equal(core.StringPtr("IBM Cloud for Financial Services")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("custom")))
				Expect(createCustomControlLibraryOptionsModel.Controls).To(Equal([]compliancev2.ControlsInControlLib{*controlsInControlLibModel}))
				Expect(createCustomControlLibraryOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("33fc7b80-0fa5-4f16-bbba-1f293f660f0d")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("1.1.0")))
				Expect(createCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(createCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createCustomControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				Expect(profileControlsPrototypeModel).ToNot(BeNil())
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				Expect(profileControlsPrototypeModel.ControlLibraryID).To(Equal(core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")))
				Expect(profileControlsPrototypeModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
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
				createProfileOptionsControls := []compliancev2.ProfileControlsPrototype{}
				createProfileOptionsDefaultParameters := []compliancev2.DefaultParametersPrototype{}
				createProfileOptionsModel := complianceService.NewCreateProfileOptions(createProfileOptionsProfileName, createProfileOptionsProfileDescription, createProfileOptionsProfileType, createProfileOptionsControls, createProfileOptionsDefaultParameters)
				createProfileOptionsModel.SetProfileName("test_profile1")
				createProfileOptionsModel.SetProfileDescription("test_description1")
				createProfileOptionsModel.SetProfileType("custom")
				createProfileOptionsModel.SetControls([]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel})
				createProfileOptionsModel.SetDefaultParameters([]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel})
				createProfileOptionsModel.SetXCorrelationID("testString")
				createProfileOptionsModel.SetXRequestID("testString")
				createProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileOptionsModel).ToNot(BeNil())
				Expect(createProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("test_profile1")))
				Expect(createProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("test_description1")))
				Expect(createProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("custom")))
				Expect(createProfileOptionsModel.Controls).To(Equal([]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}))
				Expect(createProfileOptionsModel.DefaultParameters).To(Equal([]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}))
				Expect(createProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateScanOptions successfully`, func() {
				// Construct an instance of the CreateScanOptions model
				createScanOptionsAttachmentID := "testString"
				createScanOptionsModel := complianceService.NewCreateScanOptions(createScanOptionsAttachmentID)
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
				deleteCustomControlLibraryOptionsModel := complianceService.NewDeleteCustomControlLibraryOptions(controlLibrariesID)
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
				deleteCustomProfileOptionsModel := complianceService.NewDeleteCustomProfileOptions(profilesID)
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
				deleteProfileAttachmentOptionsModel := complianceService.NewDeleteProfileAttachmentOptions(attachmentID, profilesID)
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
			It(`Invoke NewGetControlLibraryOptions successfully`, func() {
				// Construct an instance of the GetControlLibraryOptions model
				controlLibrariesID := "testString"
				getControlLibraryOptionsModel := complianceService.NewGetControlLibraryOptions(controlLibrariesID)
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
			It(`Invoke NewGetProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the GetProfileAttachmentOptions model
				attachmentID := "testString"
				profilesID := "testString"
				getProfileAttachmentOptionsModel := complianceService.NewGetProfileAttachmentOptions(attachmentID, profilesID)
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
				getProfileOptionsModel := complianceService.NewGetProfileOptions(profilesID)
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
			It(`Invoke NewListAttachmentsAccountOptions successfully`, func() {
				// Construct an instance of the ListAttachmentsAccountOptions model
				listAttachmentsAccountOptionsModel := complianceService.NewListAttachmentsAccountOptions()
				listAttachmentsAccountOptionsModel.SetXCorrelationID("testString")
				listAttachmentsAccountOptionsModel.SetXRequestID("testString")
				listAttachmentsAccountOptionsModel.SetLimit(int64(50))
				listAttachmentsAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentsAccountOptionsModel).ToNot(BeNil())
				Expect(listAttachmentsAccountOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsAccountOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsAccountOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listAttachmentsAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAttachmentsOptions successfully`, func() {
				// Construct an instance of the ListAttachmentsOptions model
				profilesID := "testString"
				listAttachmentsOptionsModel := complianceService.NewListAttachmentsOptions(profilesID)
				listAttachmentsOptionsModel.SetProfilesID("testString")
				listAttachmentsOptionsModel.SetXCorrelationID("testString")
				listAttachmentsOptionsModel.SetXRequestID("testString")
				listAttachmentsOptionsModel.SetLimit(int64(50))
				listAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentsOptionsModel).ToNot(BeNil())
				Expect(listAttachmentsOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListControlLibrariesOptions successfully`, func() {
				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := complianceService.NewListControlLibrariesOptions()
				listControlLibrariesOptionsModel.SetXCorrelationID("testString")
				listControlLibrariesOptionsModel.SetXRequestID("testString")
				listControlLibrariesOptionsModel.SetLimit(int64(50))
				listControlLibrariesOptionsModel.SetControlLibraryType("testString")
				listControlLibrariesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listControlLibrariesOptionsModel).ToNot(BeNil())
				Expect(listControlLibrariesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listControlLibrariesOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := complianceService.NewListProfilesOptions()
				listProfilesOptionsModel.SetXCorrelationID("testString")
				listProfilesOptionsModel.SetXRequestID("testString")
				listProfilesOptionsModel.SetLimit(int64(50))
				listProfilesOptionsModel.SetProfileType("testString")
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listProfilesOptionsModel.ProfileType).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMultiCloudScope successfully`, func() {
				environment := "testString"
				properties := []compliancev2.Property{}
				_model, err := complianceService.NewMultiCloudScope(environment, properties)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(compliancev2.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the Implementation model
				implementationModel := new(compliancev2.Implementation)
				Expect(implementationModel).ToNot(BeNil())
				implementationModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				implementationModel.AssessmentMethod = core.StringPtr("ibm-cloud-rule")
				implementationModel.AssessmentType = core.StringPtr("Automated")
				implementationModel.AssessmentDescription = core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")
				implementationModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationModel.Parameters = []compliancev2.ParameterInfo{*parameterInfoModel}
				Expect(implementationModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(implementationModel.AssessmentMethod).To(Equal(core.StringPtr("ibm-cloud-rule")))
				Expect(implementationModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(implementationModel.AssessmentDescription).To(Equal(core.StringPtr("Check that there is an Activity Tracker event route defined to collect global events generated by IBM Cloud services")))
				Expect(implementationModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationModel.Parameters).To(Equal([]compliancev2.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(compliancev2.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ControlSpecificationID = core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("iam-identity")
				controlSpecificationsModel.ComponenetName = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("ibm-cloud")
				controlSpecificationsModel.ControlSpecificationDescription = core.StringPtr("IBM cloud")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []compliancev2.Implementation{*implementationModel}
				Expect(controlSpecificationsModel.ControlSpecificationID).To(Equal(core.StringPtr("5c7d6f88-a92f-4734-9b49-bd22b0900184")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("iam-identity")))
				Expect(controlSpecificationsModel.ComponenetName).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(controlSpecificationsModel.ControlSpecificationDescription).To(Equal(core.StringPtr("IBM cloud")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]compliancev2.Implementation{*implementationModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(compliancev2.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("sc-7")
				controlDocsModel.ControlDocsType = core.StringPtr("ibm-cloud")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("sc-7")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("ibm-cloud")))

				// Construct an instance of the ControlsInControlLib model
				controlsInControlLibModel := new(compliancev2.ControlsInControlLib)
				Expect(controlsInControlLibModel).ToNot(BeNil())
				controlsInControlLibModel.ControlName = core.StringPtr("SC-7")
				controlsInControlLibModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				controlsInControlLibModel.ControlDescription = core.StringPtr("Boundary Protection")
				controlsInControlLibModel.ControlCategory = core.StringPtr("System and Communications Protection")
				controlsInControlLibModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibModel.ControlRequirement = core.BoolPtr(false)
				controlsInControlLibModel.ControlTags = []string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}
				controlsInControlLibModel.ControlSpecifications = []compliancev2.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibModel.ControlDocs = controlDocsModel
				controlsInControlLibModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibModel.ControlName).To(Equal(core.StringPtr("SC-7")))
				Expect(controlsInControlLibModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))
				Expect(controlsInControlLibModel.ControlDescription).To(Equal(core.StringPtr("Boundary Protection")))
				Expect(controlsInControlLibModel.ControlCategory).To(Equal(core.StringPtr("System and Communications Protection")))
				Expect(controlsInControlLibModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibModel.ControlRequirement).To(Equal(core.BoolPtr(false)))
				Expect(controlsInControlLibModel.ControlTags).To(Equal([]string{"1fa45e17-9322-4e6c-bbd6-1c51db08e790"}))
				Expect(controlsInControlLibModel.ControlSpecifications).To(Equal([]compliancev2.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				controlLibrariesID := "testString"
				replaceCustomControlLibraryOptionsModel := complianceService.NewReplaceCustomControlLibraryOptions(controlLibrariesID)
				replaceCustomControlLibraryOptionsModel.SetControlLibrariesID("testString")
				replaceCustomControlLibraryOptionsModel.SetID("testString")
				replaceCustomControlLibraryOptionsModel.SetAccountID("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryName("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryDescription("IBM Cloud for Financial Services")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryType("custom")
				replaceCustomControlLibraryOptionsModel.SetVersionGroupLabel("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryVersion("testString")
				replaceCustomControlLibraryOptionsModel.SetCreatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceCustomControlLibraryOptionsModel.SetCreatedBy("testString")
				replaceCustomControlLibraryOptionsModel.SetUpdatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceCustomControlLibraryOptionsModel.SetUpdatedBy("testString")
				replaceCustomControlLibraryOptionsModel.SetLatest(true)
				replaceCustomControlLibraryOptionsModel.SetHierarchyEnabled(true)
				replaceCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				replaceCustomControlLibraryOptionsModel.SetControlParentsCount(int64(38))
				replaceCustomControlLibraryOptionsModel.SetControls([]compliancev2.ControlsInControlLib{*controlsInControlLibModel})
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
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.CreatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceCustomControlLibraryOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceCustomControlLibraryOptionsModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(replaceCustomControlLibraryOptionsModel.HierarchyEnabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceCustomControlLibraryOptionsModel.ControlParentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceCustomControlLibraryOptionsModel.Controls).To(Equal([]compliancev2.ControlsInControlLib{*controlsInControlLibModel}))
				Expect(replaceCustomControlLibraryOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the Property model
				propertyModel := new(compliancev2.Property)
				Expect(propertyModel).ToNot(BeNil())
				propertyModel.Name = core.StringPtr("scope_id")
				propertyModel.Value = core.StringPtr("cg3335893hh1428692d6747cf300yeb5")
				Expect(propertyModel.Name).To(Equal(core.StringPtr("scope_id")))
				Expect(propertyModel.Value).To(Equal(core.StringPtr("cg3335893hh1428692d6747cf300yeb5")))

				// Construct an instance of the MultiCloudScope model
				multiCloudScopeModel := new(compliancev2.MultiCloudScope)
				Expect(multiCloudScopeModel).ToNot(BeNil())
				multiCloudScopeModel.Environment = core.StringPtr("ibm-cloud")
				multiCloudScopeModel.Properties = []compliancev2.Property{*propertyModel}
				Expect(multiCloudScopeModel.Environment).To(Equal(core.StringPtr("ibm-cloud")))
				Expect(multiCloudScopeModel.Properties).To(Equal([]compliancev2.Property{*propertyModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(compliancev2.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(15))
				failedControlsModel.FailedControlIds = []string{}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(15))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{}))

				// Construct an instance of the AttachmentsNotificationsPrototype model
				attachmentsNotificationsPrototypeModel := new(compliancev2.AttachmentsNotificationsPrototype)
				Expect(attachmentsNotificationsPrototypeModel).ToNot(BeNil())
				attachmentsNotificationsPrototypeModel.Enabled = core.BoolPtr(false)
				attachmentsNotificationsPrototypeModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPrototypeModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(attachmentsNotificationsPrototypeModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the AttachmentParametersPrototype model
				attachmentParametersPrototypeModel := new(compliancev2.AttachmentParametersPrototype)
				Expect(attachmentParametersPrototypeModel).ToNot(BeNil())
				attachmentParametersPrototypeModel.AssessmentType = core.StringPtr("Automated")
				attachmentParametersPrototypeModel.AssessmentID = core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")
				attachmentParametersPrototypeModel.ParameterName = core.StringPtr("session_invalidation_in_seconds")
				attachmentParametersPrototypeModel.ParameterValue = core.StringPtr("120")
				attachmentParametersPrototypeModel.ParameterDisplayName = core.StringPtr("Sign out due to inactivity in seconds")
				attachmentParametersPrototypeModel.ParameterType = core.StringPtr("numeric")
				Expect(attachmentParametersPrototypeModel.AssessmentType).To(Equal(core.StringPtr("Automated")))
				Expect(attachmentParametersPrototypeModel.AssessmentID).To(Equal(core.StringPtr("rule-a637949b-7e51-46c4-afd4-b96619001bf1")))
				Expect(attachmentParametersPrototypeModel.ParameterName).To(Equal(core.StringPtr("session_invalidation_in_seconds")))
				Expect(attachmentParametersPrototypeModel.ParameterValue).To(Equal(core.StringPtr("120")))
				Expect(attachmentParametersPrototypeModel.ParameterDisplayName).To(Equal(core.StringPtr("Sign out due to inactivity in seconds")))
				Expect(attachmentParametersPrototypeModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the LastScan model
				lastScanModel := new(compliancev2.LastScan)
				Expect(lastScanModel).ToNot(BeNil())
				lastScanModel.ID = core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")
				lastScanModel.Status = core.StringPtr("in_progress")
				lastScanModel.Time = core.StringPtr("testString")
				Expect(lastScanModel.ID).To(Equal(core.StringPtr("e8a39d25-0051-4328-8462-988ad321f49a")))
				Expect(lastScanModel.Status).To(Equal(core.StringPtr("in_progress")))
				Expect(lastScanModel.Time).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				attachmentID := "testString"
				profilesID := "testString"
				replaceProfileAttachmentOptionsModel := complianceService.NewReplaceProfileAttachmentOptions(attachmentID, profilesID)
				replaceProfileAttachmentOptionsModel.SetAttachmentID("testString")
				replaceProfileAttachmentOptionsModel.SetProfilesID("testString")
				replaceProfileAttachmentOptionsModel.SetID("testString")
				replaceProfileAttachmentOptionsModel.SetProfileID("testString")
				replaceProfileAttachmentOptionsModel.SetAccountID("testString")
				replaceProfileAttachmentOptionsModel.SetInstanceID("testString")
				replaceProfileAttachmentOptionsModel.SetScope([]compliancev2.MultiCloudScope{*multiCloudScopeModel})
				replaceProfileAttachmentOptionsModel.SetCreatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceProfileAttachmentOptionsModel.SetCreatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetUpdatedOn(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				replaceProfileAttachmentOptionsModel.SetUpdatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetStatus("enabled")
				replaceProfileAttachmentOptionsModel.SetSchedule("daily")
				replaceProfileAttachmentOptionsModel.SetNotifications(attachmentsNotificationsPrototypeModel)
				replaceProfileAttachmentOptionsModel.SetAttachmentParameters([]compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel})
				replaceProfileAttachmentOptionsModel.SetLastScan(lastScanModel)
				replaceProfileAttachmentOptionsModel.SetNextScanTime("testString")
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
				Expect(replaceProfileAttachmentOptionsModel.Scope).To(Equal([]compliancev2.MultiCloudScope{*multiCloudScopeModel}))
				Expect(replaceProfileAttachmentOptionsModel.CreatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceProfileAttachmentOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedOn).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(replaceProfileAttachmentOptionsModel.Schedule).To(Equal(core.StringPtr("daily")))
				Expect(replaceProfileAttachmentOptionsModel.Notifications).To(Equal(attachmentsNotificationsPrototypeModel))
				Expect(replaceProfileAttachmentOptionsModel.AttachmentParameters).To(Equal([]compliancev2.AttachmentParametersPrototype{*attachmentParametersPrototypeModel}))
				Expect(replaceProfileAttachmentOptionsModel.LastScan).To(Equal(lastScanModel))
				Expect(replaceProfileAttachmentOptionsModel.NextScanTime).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Name).To(Equal(core.StringPtr("account-0d8c3805dfea40aa8ad02265a18eb12b")))
				Expect(replaceProfileAttachmentOptionsModel.Description).To(Equal(core.StringPtr("Test description")))
				Expect(replaceProfileAttachmentOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsPrototype model
				profileControlsPrototypeModel := new(compliancev2.ProfileControlsPrototype)
				Expect(profileControlsPrototypeModel).ToNot(BeNil())
				profileControlsPrototypeModel.ControlLibraryID = core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")
				profileControlsPrototypeModel.ControlID = core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")
				Expect(profileControlsPrototypeModel.ControlLibraryID).To(Equal(core.StringPtr("e98a56ff-dc24-41d4-9875-1e188e2da6cd")))
				Expect(profileControlsPrototypeModel.ControlID).To(Equal(core.StringPtr("1fa45e17-9322-4e6c-bbd6-1c51db08e790")))

				// Construct an instance of the DefaultParametersPrototype model
				defaultParametersPrototypeModel := new(compliancev2.DefaultParametersPrototype)
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
				replaceProfileOptionsControls := []compliancev2.ProfileControlsPrototype{}
				replaceProfileOptionsDefaultParameters := []compliancev2.DefaultParametersPrototype{}
				replaceProfileOptionsModel := complianceService.NewReplaceProfileOptions(profilesID, replaceProfileOptionsProfileName, replaceProfileOptionsProfileDescription, replaceProfileOptionsProfileType, replaceProfileOptionsControls, replaceProfileOptionsDefaultParameters)
				replaceProfileOptionsModel.SetProfilesID("testString")
				replaceProfileOptionsModel.SetProfileName("test_profile1")
				replaceProfileOptionsModel.SetProfileDescription("test_description1")
				replaceProfileOptionsModel.SetProfileType("custom")
				replaceProfileOptionsModel.SetControls([]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel})
				replaceProfileOptionsModel.SetDefaultParameters([]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel})
				replaceProfileOptionsModel.SetXCorrelationID("testString")
				replaceProfileOptionsModel.SetXRequestID("testString")
				replaceProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProfileOptionsModel).ToNot(BeNil())
				Expect(replaceProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("test_profile1")))
				Expect(replaceProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("test_description1")))
				Expect(replaceProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("custom")))
				Expect(replaceProfileOptionsModel.Controls).To(Equal([]compliancev2.ProfileControlsPrototype{*profileControlsPrototypeModel}))
				Expect(replaceProfileOptionsModel.DefaultParameters).To(Equal([]compliancev2.DefaultParametersPrototype{*defaultParametersPrototypeModel}))
				Expect(replaceProfileOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.XRequestID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

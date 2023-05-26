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

package sccphoenixcomplianceapisv1_test

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
	"github.com/IBM/scc-go-sdk/v4/sccphoenixcomplianceapisv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SccPhoenixComplianceApisV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(sccPhoenixComplianceApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				URL: "https://sccphoenixcomplianceapisv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(sccPhoenixComplianceApisService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCC_PHOENIX_COMPLIANCE_APIS_URL":       "https://sccphoenixcomplianceapisv1/api",
				"SCC_PHOENIX_COMPLIANCE_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{})
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := sccPhoenixComplianceApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sccPhoenixComplianceApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sccPhoenixComplianceApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sccPhoenixComplianceApisService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL: "https://testService/api",
				})
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sccPhoenixComplianceApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sccPhoenixComplianceApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sccPhoenixComplianceApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sccPhoenixComplianceApisService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{})
				err := sccPhoenixComplianceApisService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := sccPhoenixComplianceApisService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != sccPhoenixComplianceApisService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(sccPhoenixComplianceApisService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(sccPhoenixComplianceApisService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCC_PHOENIX_COMPLIANCE_APIS_URL":       "https://sccphoenixcomplianceapisv1/api",
				"SCC_PHOENIX_COMPLIANCE_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(sccPhoenixComplianceApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCC_PHOENIX_COMPLIANCE_APIS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1UsingExternalConfig(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(sccPhoenixComplianceApisService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = sccphoenixcomplianceapisv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions) - Operation response error`, func() {
		createProfilePath := "/instances/testString/v3/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProfilePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProfile with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				createProfileOptionsModel.InstanceID = core.StringPtr("testString")
				createProfileOptionsModel.ProfileName = core.StringPtr("testString")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				createProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				createProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				createProfileOptionsModel.Latest = core.BoolPtr(true)
				createProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				createProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				createProfileOptionsModel.TransactionID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
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
		createProfilePath := "/instances/testString/v3/profiles"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				createProfileOptionsModel.InstanceID = core.StringPtr("testString")
				createProfileOptionsModel.ProfileName = core.StringPtr("testString")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				createProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				createProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				createProfileOptionsModel.Latest = core.BoolPtr(true)
				createProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				createProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				createProfileOptionsModel.TransactionID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.CreateProfileWithContext(ctx, createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.CreateProfileWithContext(ctx, createProfileOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.CreateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				createProfileOptionsModel.InstanceID = core.StringPtr("testString")
				createProfileOptionsModel.ProfileName = core.StringPtr("testString")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				createProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				createProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				createProfileOptionsModel.Latest = core.BoolPtr(true)
				createProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				createProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				createProfileOptionsModel.TransactionID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProfile with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				createProfileOptionsModel.InstanceID = core.StringPtr("testString")
				createProfileOptionsModel.ProfileName = core.StringPtr("testString")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				createProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				createProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				createProfileOptionsModel.Latest = core.BoolPtr(true)
				createProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				createProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				createProfileOptionsModel.TransactionID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProfileOptions model with no property values
				createProfileOptionsModelNew := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModelNew)
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
			It(`Invoke CreateProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(sccphoenixcomplianceapisv1.CreateProfileOptions)
				createProfileOptionsModel.InstanceID = core.StringPtr("testString")
				createProfileOptionsModel.ProfileName = core.StringPtr("testString")
				createProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				createProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				createProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				createProfileOptionsModel.Latest = core.BoolPtr(true)
				createProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				createProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				createProfileOptionsModel.TransactionID = core.StringPtr("testString")
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.CreateProfile(createProfileOptionsModel)
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
		listProfilesPath := "/instances/testString/v3/profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				listProfilesOptionsModel.InstanceID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
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
		listProfilesPath := "/instances/testString/v3/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				listProfilesOptionsModel.InstanceID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ListProfilesWithContext(ctx, listProfilesOptionsModel)
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profiles": [{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ListProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				listProfilesOptionsModel.InstanceID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				listProfilesOptionsModel.InstanceID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProfilesOptions model with no property values
				listProfilesOptionsModelNew := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(sccphoenixcomplianceapisv1.ListProfilesOptions)
				listProfilesOptionsModel.InstanceID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ListProfiles(listProfilesOptionsModel)
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
	Describe(`AddProfile(addProfileOptions *AddProfileOptions) - Operation response error`, func() {
		addProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addProfilePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddProfile with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AddProfileOptions model
				addProfileOptionsModel := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				addProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				addProfileOptionsModel.InstanceID = core.StringPtr("testString")
				addProfileOptionsModel.ProfileName = core.StringPtr("testString")
				addProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				addProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				addProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				addProfileOptionsModel.Latest = core.BoolPtr(true)
				addProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				addProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				addProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				addProfileOptionsModel.TransactionID = core.StringPtr("testString")
				addProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddProfile(addProfileOptions *AddProfileOptions)`, func() {
		addProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addProfilePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke AddProfile successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AddProfileOptions model
				addProfileOptionsModel := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				addProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				addProfileOptionsModel.InstanceID = core.StringPtr("testString")
				addProfileOptionsModel.ProfileName = core.StringPtr("testString")
				addProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				addProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				addProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				addProfileOptionsModel.Latest = core.BoolPtr(true)
				addProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				addProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				addProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				addProfileOptionsModel.TransactionID = core.StringPtr("testString")
				addProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.AddProfileWithContext(ctx, addProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.AddProfileWithContext(ctx, addProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addProfilePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke AddProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.AddProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AddProfileOptions model
				addProfileOptionsModel := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				addProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				addProfileOptionsModel.InstanceID = core.StringPtr("testString")
				addProfileOptionsModel.ProfileName = core.StringPtr("testString")
				addProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				addProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				addProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				addProfileOptionsModel.Latest = core.BoolPtr(true)
				addProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				addProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				addProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				addProfileOptionsModel.TransactionID = core.StringPtr("testString")
				addProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddProfile with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AddProfileOptions model
				addProfileOptionsModel := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				addProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				addProfileOptionsModel.InstanceID = core.StringPtr("testString")
				addProfileOptionsModel.ProfileName = core.StringPtr("testString")
				addProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				addProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				addProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				addProfileOptionsModel.Latest = core.BoolPtr(true)
				addProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				addProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				addProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				addProfileOptionsModel.TransactionID = core.StringPtr("testString")
				addProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddProfileOptions model with no property values
				addProfileOptionsModelNew := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModelNew)
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
			It(`Invoke AddProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the AddProfileOptions model
				addProfileOptionsModel := new(sccphoenixcomplianceapisv1.AddProfileOptions)
				addProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				addProfileOptionsModel.InstanceID = core.StringPtr("testString")
				addProfileOptionsModel.ProfileName = core.StringPtr("testString")
				addProfileOptionsModel.ProfileDescription = core.StringPtr("testString")
				addProfileOptionsModel.ProfileType = core.StringPtr("predefined")
				addProfileOptionsModel.ProfileVersion = core.StringPtr("testString")
				addProfileOptionsModel.Latest = core.BoolPtr(true)
				addProfileOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				addProfileOptionsModel.Controls = []sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}
				addProfileOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				addProfileOptionsModel.TransactionID = core.StringPtr("testString")
				addProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.AddProfile(addProfileOptionsModel)
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
		getProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfile with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
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
		getProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.GetProfileWithContext(ctx, getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.GetProfileWithContext(ctx, getProfileOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfileOptionsModelNew := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileOptions)
				getProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfile(getProfileOptionsModel)
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
		deleteCustomProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomProfilePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCustomProfile with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
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
		deleteCustomProfilePath := "/instances/testString/v3/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.DeleteCustomProfileWithContext(ctx, deleteCustomProfileOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "profile_name": "ProfileName", "profile_description": "ProfileDescription", "profile_type": "ProfileType", "profile_version": "ProfileVersion", "version_group_label": "VersionGroupLabel", "latest": true, "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "controls_count": 13, "attachments_count": 16, "controls": [{"control_library_id": "ControlLibraryID", "control_id": "ControlID", "control_library_version": "ControlLibraryVersion", "control_name": "ControlName", "control_description": "ControlDescription", "control_severity": "ControlSeverity", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}]}], "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke DeleteCustomProfile successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomProfile with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomProfileOptions model with no property values
				deleteCustomProfileOptionsModelNew := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomProfileOptions model
				deleteCustomProfileOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomProfileOptions)
				deleteCustomProfileOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomProfile(deleteCustomProfileOptionsModel)
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
	Describe(`ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions) - Operation response error`, func() {
		replaceProfileParametersPath := "/instances/testString/v3/profiles/testString/parameters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileParametersPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceProfileParameters with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileParametersOptions model
				replaceProfileParametersOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				replaceProfileParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.ID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				replaceProfileParametersOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceProfileParameters(replaceProfileParametersOptions *ReplaceProfileParametersOptions)`, func() {
		replaceProfileParametersPath := "/instances/testString/v3/profiles/testString/parameters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileParametersPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceProfileParameters successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileParametersOptions model
				replaceProfileParametersOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				replaceProfileParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.ID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				replaceProfileParametersOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParametersWithContext(ctx, replaceProfileParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ReplaceProfileParametersWithContext(ctx, replaceProfileParametersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileParametersPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "default_parameters": [{"assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameter_name": "ParameterName", "parameter_default_value": "ParameterDefaultValue", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceProfileParameters successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParameters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileParametersOptions model
				replaceProfileParametersOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				replaceProfileParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.ID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				replaceProfileParametersOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfileParameters with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileParametersOptions model
				replaceProfileParametersOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				replaceProfileParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.ID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				replaceProfileParametersOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileParametersOptions model with no property values
				replaceProfileParametersOptionsModelNew := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModelNew)
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
			It(`Invoke ReplaceProfileParameters successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceProfileParametersOptions model
				replaceProfileParametersOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileParametersOptions)
				replaceProfileParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.ID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.DefaultParameters = []sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}
				replaceProfileParametersOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileParameters(replaceProfileParametersOptionsModel)
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
		createAttachmentPath := "/instances/testString/v3/profiles/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAttachmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAttachment with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}
				createAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
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
		createAttachmentPath := "/instances/testString/v3/profiles/testString/attachments"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profile_id": "ProfileID", "attachments": [{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "Status", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "last_scan": "LastScan", "last_scan_status": "LastScanStatus", "last_scan_time": "LastScanTime"}]}`)
				}))
			})
			It(`Invoke CreateAttachment successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}
				createAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.CreateAttachmentWithContext(ctx, createAttachmentOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profile_id": "ProfileID", "attachments": [{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "Status", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "last_scan": "LastScan", "last_scan_status": "LastScanStatus", "last_scan_time": "LastScanTime"}]}`)
				}))
			})
			It(`Invoke CreateAttachment successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.CreateAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}
				createAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAttachment with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}
				createAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAttachmentOptions model with no property values
				createAttachmentOptionsModelNew := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModelNew)
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
			It(`Invoke CreateAttachment successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel

				// Construct an instance of the CreateAttachmentOptions model
				createAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.CreateAttachmentOptions)
				createAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				createAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				createAttachmentOptionsModel.Attachments = []sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}
				createAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.CreateAttachment(createAttachmentOptionsModel)
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
	Describe(`CheckProfileAttachmnets(checkProfileAttachmnetsOptions *CheckProfileAttachmnetsOptions) - Operation response error`, func() {
		checkProfileAttachmnetsPath := "/instances/testString/v3/profiles/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkProfileAttachmnetsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CheckProfileAttachmnets with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CheckProfileAttachmnetsOptions model
				checkProfileAttachmnetsOptionsModel := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				checkProfileAttachmnetsOptionsModel.ProfilesID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.InstanceID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.TransactionID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CheckProfileAttachmnets(checkProfileAttachmnetsOptions *CheckProfileAttachmnetsOptions)`, func() {
		checkProfileAttachmnetsPath := "/instances/testString/v3/profiles/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkProfileAttachmnetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profile_id": "ProfileID", "account_id": "AccountID", "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "CreatedOn", "created_by": "CreatedBy", "updated_on": "UpdatedOn", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}], "attachments": [{"attachments": [{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}]}]}`)
				}))
			})
			It(`Invoke CheckProfileAttachmnets successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the CheckProfileAttachmnetsOptions model
				checkProfileAttachmnetsOptionsModel := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				checkProfileAttachmnetsOptionsModel.ProfilesID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.InstanceID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.TransactionID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnetsWithContext(ctx, checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.CheckProfileAttachmnetsWithContext(ctx, checkProfileAttachmnetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(checkProfileAttachmnetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "profile_id": "ProfileID", "account_id": "AccountID", "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "CreatedOn", "created_by": "CreatedBy", "updated_on": "UpdatedOn", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}], "attachments": [{"attachments": [{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}]}]}`)
				}))
			})
			It(`Invoke CheckProfileAttachmnets successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CheckProfileAttachmnetsOptions model
				checkProfileAttachmnetsOptionsModel := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				checkProfileAttachmnetsOptionsModel.ProfilesID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.InstanceID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.TransactionID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CheckProfileAttachmnets with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CheckProfileAttachmnetsOptions model
				checkProfileAttachmnetsOptionsModel := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				checkProfileAttachmnetsOptionsModel.ProfilesID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.InstanceID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.TransactionID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CheckProfileAttachmnetsOptions model with no property values
				checkProfileAttachmnetsOptionsModelNew := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModelNew)
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
			It(`Invoke CheckProfileAttachmnets successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CheckProfileAttachmnetsOptions model
				checkProfileAttachmnetsOptionsModel := new(sccphoenixcomplianceapisv1.CheckProfileAttachmnetsOptions)
				checkProfileAttachmnetsOptionsModel.ProfilesID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.InstanceID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.TransactionID = core.StringPtr("testString")
				checkProfileAttachmnetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.CheckProfileAttachmnets(checkProfileAttachmnetsOptionsModel)
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
	Describe(`GetProfileAttachmnet(getProfileAttachmnetOptions *GetProfileAttachmnetOptions) - Operation response error`, func() {
		getProfileAttachmnetPath := "/instances/testString/v3/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmnetPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileAttachmnet with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmnetOptions model
				getProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				getProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileAttachmnet(getProfileAttachmnetOptions *GetProfileAttachmnetOptions)`, func() {
		getProfileAttachmnetPath := "/instances/testString/v3/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmnetPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke GetProfileAttachmnet successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileAttachmnetOptions model
				getProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				getProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnetWithContext(ctx, getProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.GetProfileAttachmnetWithContext(ctx, getProfileAttachmnetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileAttachmnetPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke GetProfileAttachmnet successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileAttachmnetOptions model
				getProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				getProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileAttachmnet with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmnetOptions model
				getProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				getProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileAttachmnetOptions model with no property values
				getProfileAttachmnetOptionsModelNew := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModelNew)
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
			It(`Invoke GetProfileAttachmnet successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetProfileAttachmnetOptions model
				getProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.GetProfileAttachmnetOptions)
				getProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.GetProfileAttachmnet(getProfileAttachmnetOptionsModel)
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
		replaceProfileAttachmentPath := "/instances/testString/v3/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceProfileAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.IncludedScope = scopePayloadModel
				replaceProfileAttachmentOptionsModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.CreatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				replaceProfileAttachmentOptionsModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				replaceProfileAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
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
		replaceProfileAttachmentPath := "/instances/testString/v3/profiles/testString/attachments/testString"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.IncludedScope = scopePayloadModel
				replaceProfileAttachmentOptionsModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.CreatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				replaceProfileAttachmentOptionsModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				replaceProfileAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ReplaceProfileAttachmentWithContext(ctx, replaceProfileAttachmentOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke ReplaceProfileAttachment successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.IncludedScope = scopePayloadModel
				replaceProfileAttachmentOptionsModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.CreatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				replaceProfileAttachmentOptionsModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				replaceProfileAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceProfileAttachment with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.IncludedScope = scopePayloadModel
				replaceProfileAttachmentOptionsModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.CreatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				replaceProfileAttachmentOptionsModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				replaceProfileAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceProfileAttachmentOptions model with no property values
				replaceProfileAttachmentOptionsModelNew := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				replaceProfileAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceProfileAttachmentOptions)
				replaceProfileAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.ID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.AccountID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.IncludedScope = scopePayloadModel
				replaceProfileAttachmentOptionsModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				replaceProfileAttachmentOptionsModel.CreatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.CreatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedBy = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.UpdatedOn = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Status = core.StringPtr("enabled")
				replaceProfileAttachmentOptionsModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				replaceProfileAttachmentOptionsModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				replaceProfileAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceProfileAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceProfileAttachment(replaceProfileAttachmentOptionsModel)
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
	Describe(`DeleteProfileAttachmnet(deleteProfileAttachmnetOptions *DeleteProfileAttachmnetOptions) - Operation response error`, func() {
		deleteProfileAttachmnetPath := "/instances/testString/v3/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmnetPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteProfileAttachmnet with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmnetOptions model
				deleteProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				deleteProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProfileAttachmnet(deleteProfileAttachmnetOptions *DeleteProfileAttachmnetOptions)`, func() {
		deleteProfileAttachmnetPath := "/instances/testString/v3/profiles/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmnetPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke DeleteProfileAttachmnet successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProfileAttachmnetOptions model
				deleteProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				deleteProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnetWithContext(ctx, deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.DeleteProfileAttachmnetWithContext(ctx, deleteProfileAttachmnetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfileAttachmnetPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "included_scope": {"scope_id": "ScopeID", "scope_type": "ScopeType"}, "exclusions": [{"scope_id": "ScopeID", "scope_type": "ScopeType"}], "created_by": "CreatedBy", "created_on": "CreatedOn", "updated_by": "UpdatedBy", "updated_on": "UpdatedOn", "status": "enabled", "attachment_parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}], "attachment_notifications": {"enabled": false, "controls": {"threshold_limit": 14, "failed_control_ids": ["FailedControlIds"]}}}`)
				}))
			})
			It(`Invoke DeleteProfileAttachmnet successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProfileAttachmnetOptions model
				deleteProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				deleteProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProfileAttachmnet with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmnetOptions model
				deleteProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				deleteProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProfileAttachmnetOptions model with no property values
				deleteProfileAttachmnetOptionsModelNew := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModelNew)
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
			It(`Invoke DeleteProfileAttachmnet successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileAttachmnetOptions model
				deleteProfileAttachmnetOptionsModel := new(sccphoenixcomplianceapisv1.DeleteProfileAttachmnetOptions)
				deleteProfileAttachmnetOptionsModel.ProfilesID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.InstanceID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileAttachmnetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteProfileAttachmnet(deleteProfileAttachmnetOptionsModel)
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
	Describe(`ListAttachmentParameters(listAttachmentParametersOptions *ListAttachmentParametersOptions) - Operation response error`, func() {
		listAttachmentParametersPath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentParametersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAttachmentParameters with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentParametersOptions model
				listAttachmentParametersOptionsModel := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				listAttachmentParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.AttachmentID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.InstanceID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAttachmentParameters(listAttachmentParametersOptions *ListAttachmentParametersOptions)`, func() {
		listAttachmentParametersPath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentParametersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ListAttachmentParameters successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ListAttachmentParametersOptions model
				listAttachmentParametersOptionsModel := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				listAttachmentParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.AttachmentID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.InstanceID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ListAttachmentParametersWithContext(ctx, listAttachmentParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ListAttachmentParametersWithContext(ctx, listAttachmentParametersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAttachmentParametersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ListAttachmentParameters successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ListAttachmentParameters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAttachmentParametersOptions model
				listAttachmentParametersOptionsModel := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				listAttachmentParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.AttachmentID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.InstanceID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAttachmentParameters with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentParametersOptions model
				listAttachmentParametersOptionsModel := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				listAttachmentParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.AttachmentID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.InstanceID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAttachmentParametersOptions model with no property values
				listAttachmentParametersOptionsModelNew := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModelNew)
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
			It(`Invoke ListAttachmentParameters successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListAttachmentParametersOptions model
				listAttachmentParametersOptionsModel := new(sccphoenixcomplianceapisv1.ListAttachmentParametersOptions)
				listAttachmentParametersOptionsModel.ProfilesID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.AttachmentID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.InstanceID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.TransactionID = core.StringPtr("testString")
				listAttachmentParametersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ListAttachmentParameters(listAttachmentParametersOptionsModel)
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
	Describe(`ReplaceAttachment(replaceAttachmentOptions *ReplaceAttachmentOptions) - Operation response error`, func() {
		replaceAttachmentPath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceAttachment with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmentOptions model
				replaceAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				replaceAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterDisplayName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterType = core.StringPtr("numeric")
				replaceAttachmentOptionsModel.ParameterValue = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentType = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceAttachment(replaceAttachmentOptions *ReplaceAttachmentOptions)`, func() {
		replaceAttachmentPath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmentPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceAttachment successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmentOptions model
				replaceAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				replaceAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterDisplayName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterType = core.StringPtr("numeric")
				replaceAttachmentOptionsModel.ParameterValue = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentType = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmentWithContext(ctx, replaceAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ReplaceAttachmentWithContext(ctx, replaceAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmentPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceAttachment successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmentOptions model
				replaceAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				replaceAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterDisplayName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterType = core.StringPtr("numeric")
				replaceAttachmentOptionsModel.ParameterValue = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentType = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceAttachment with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmentOptions model
				replaceAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				replaceAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterDisplayName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterType = core.StringPtr("numeric")
				replaceAttachmentOptionsModel.ParameterValue = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentType = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceAttachmentOptions model with no property values
				replaceAttachmentOptionsModelNew := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModelNew)
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
			It(`Invoke ReplaceAttachment successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmentOptions model
				replaceAttachmentOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmentOptions)
				replaceAttachmentOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterDisplayName = core.StringPtr("testString")
				replaceAttachmentOptionsModel.ParameterType = core.StringPtr("numeric")
				replaceAttachmentOptionsModel.ParameterValue = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentType = core.StringPtr("testString")
				replaceAttachmentOptionsModel.AssessmentID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachment(replaceAttachmentOptionsModel)
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
	Describe(`GetParametersByName(getParametersByNameOptions *GetParametersByNameOptions) - Operation response error`, func() {
		getParametersByNamePath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getParametersByNamePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetParametersByName with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetParametersByNameOptions model
				getParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				getParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				getParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				getParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				getParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				getParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				getParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetParametersByName(getParametersByNameOptions *GetParametersByNameOptions)`, func() {
		getParametersByNamePath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getParametersByNamePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke GetParametersByName successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the GetParametersByNameOptions model
				getParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				getParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				getParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				getParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				getParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				getParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				getParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.GetParametersByNameWithContext(ctx, getParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.GetParametersByNameWithContext(ctx, getParametersByNameOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getParametersByNamePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke GetParametersByName successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.GetParametersByName(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetParametersByNameOptions model
				getParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				getParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				getParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				getParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				getParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				getParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				getParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetParametersByName with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetParametersByNameOptions model
				getParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				getParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				getParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				getParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				getParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				getParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				getParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetParametersByNameOptions model with no property values
				getParametersByNameOptionsModelNew := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModelNew)
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
			It(`Invoke GetParametersByName successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetParametersByNameOptions model
				getParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.GetParametersByNameOptions)
				getParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				getParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				getParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				getParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				getParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				getParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.GetParametersByName(getParametersByNameOptionsModel)
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
	Describe(`ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions *ReplaceAttachmnetParametersByNameOptions) - Operation response error`, func() {
		replaceAttachmnetParametersByNamePath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmnetParametersByNamePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceAttachmnetParametersByName with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				replaceAttachmnetParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				replaceAttachmnetParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterType = core.StringPtr("numeric")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterValue = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmnetParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptions *ReplaceAttachmnetParametersByNameOptions)`, func() {
		replaceAttachmnetParametersByNamePath := "/instances/testString/v3/profiles/testString/attachments/testString/parameters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmnetParametersByNamePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceAttachmnetParametersByName successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				replaceAttachmnetParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				replaceAttachmnetParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterType = core.StringPtr("numeric")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterValue = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmnetParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByNameWithContext(ctx, replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByNameWithContext(ctx, replaceAttachmnetParametersByNameOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceAttachmnetParametersByNamePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric", "parameter_value": "ParameterValue", "assessment_type": "AssessmentType", "assessment_id": "AssessmentID", "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}`)
				}))
			})
			It(`Invoke ReplaceAttachmnetParametersByName successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				replaceAttachmnetParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				replaceAttachmnetParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterType = core.StringPtr("numeric")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterValue = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmnetParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceAttachmnetParametersByName with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				replaceAttachmnetParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				replaceAttachmnetParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterType = core.StringPtr("numeric")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterValue = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmnetParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceAttachmnetParametersByNameOptions model with no property values
				replaceAttachmnetParametersByNameOptionsModelNew := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModelNew)
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
			It(`Invoke ReplaceAttachmnetParametersByName successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				replaceAttachmnetParametersByNameOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceAttachmnetParametersByNameOptions)
				replaceAttachmnetParametersByNameOptionsModel.ProfilesID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.AttachmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.ParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.InstanceID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterType = core.StringPtr("numeric")
				replaceAttachmnetParametersByNameOptionsModel.NewParameterValue = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.NewParameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				replaceAttachmnetParametersByNameOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAttachmnetParametersByNameOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceAttachmnetParametersByName(replaceAttachmnetParametersByNameOptionsModel)
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
		createCustomControlLibraryPath := "/instances/testString/v3/control_libraries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomControlLibraryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				createCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
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
		createCustomControlLibraryPath := "/instances/testString/v3/control_libraries"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				createCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.CreateCustomControlLibraryWithContext(ctx, createCustomControlLibraryOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke CreateCustomControlLibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				createCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomControlLibrary with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				createCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomControlLibraryOptions model with no property values
				createCustomControlLibraryOptionsModelNew := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModelNew)
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
			It(`Invoke CreateCustomControlLibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the CreateCustomControlLibraryOptions model
				createCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.CreateCustomControlLibraryOptions)
				createCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				createCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				createCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				createCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				createCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				createCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.CreateCustomControlLibrary(createCustomControlLibraryOptionsModel)
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
		listControlLibrariesPath := "/instances/testString/v3/control_libraries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListControlLibraries with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.InstanceID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.TransactionID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
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
		listControlLibrariesPath := "/instances/testString/v3/control_libraries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listControlLibrariesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "CreatedOn", "created_by": "CreatedBy", "updated_on": "UpdatedOn", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.InstanceID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.TransactionID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ListControlLibrariesWithContext(ctx, listControlLibrariesOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 1, "limit": 20, "first": {"href": "Href"}, "next": {"href": "Href", "start": "Start"}, "control_libraries": [{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "ControlLibraryType", "created_on": "CreatedOn", "created_by": "CreatedBy", "updated_on": "UpdatedOn", "updated_by": "UpdatedBy", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13}]}`)
				}))
			})
			It(`Invoke ListControlLibraries successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ListControlLibraries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.InstanceID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.TransactionID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListControlLibraries with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.InstanceID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.TransactionID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListControlLibrariesOptions model with no property values
				listControlLibrariesOptionsModelNew := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModelNew)
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
			It(`Invoke ListControlLibraries successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ListControlLibrariesOptions model
				listControlLibrariesOptionsModel := new(sccphoenixcomplianceapisv1.ListControlLibrariesOptions)
				listControlLibrariesOptionsModel.InstanceID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.TransactionID = core.StringPtr("testString")
				listControlLibrariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ListControlLibraries(listControlLibrariesOptionsModel)
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
		replaceCustomControlLibraryPath := "/instances/testString/v3/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceCustomControlLibraryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				replaceCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
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
		replaceCustomControlLibraryPath := "/instances/testString/v3/control_libraries/testString"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				replaceCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.ReplaceCustomControlLibraryWithContext(ctx, replaceCustomControlLibraryOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke ReplaceCustomControlLibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				replaceCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceCustomControlLibrary with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				replaceCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceCustomControlLibraryOptions model with no property values
				replaceCustomControlLibraryOptionsModelNew := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				replaceCustomControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.ReplaceCustomControlLibraryOptions)
				replaceCustomControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.AccountID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryName = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryDescription = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryType = core.StringPtr("predefined")
				replaceCustomControlLibraryOptionsModel.VersionGroupLabel = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.ControlLibraryVersion = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Latest = core.BoolPtr(true)
				replaceCustomControlLibraryOptionsModel.ControlsCount = core.Int64Ptr(int64(38))
				replaceCustomControlLibraryOptionsModel.Controls = []sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}
				replaceCustomControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				replaceCustomControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.ReplaceCustomControlLibrary(replaceCustomControlLibraryOptionsModel)
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
		getControlLibraryPath := "/instances/testString/v3/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getControlLibraryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetControlLibrary with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				getControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
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
		getControlLibraryPath := "/instances/testString/v3/control_libraries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getControlLibraryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				getControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.GetControlLibraryWithContext(ctx, getControlLibraryOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke GetControlLibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.GetControlLibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				getControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetControlLibrary with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				getControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetControlLibraryOptions model with no property values
				getControlLibraryOptionsModelNew := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModelNew)
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
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the GetControlLibraryOptions model
				getControlLibraryOptionsModel := new(sccphoenixcomplianceapisv1.GetControlLibraryOptions)
				getControlLibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				getControlLibraryOptionsModel.InstanceID = core.StringPtr("testString")
				getControlLibraryOptionsModel.TransactionID = core.StringPtr("testString")
				getControlLibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.GetControlLibrary(getControlLibraryOptionsModel)
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
	Describe(`DeleteCustomControllibrary(deleteCustomControllibraryOptions *DeleteCustomControllibraryOptions) - Operation response error`, func() {
		deleteCustomControllibraryPath := "/instances/testString/v3/control_libraries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControllibraryPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCustomControllibrary with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControllibraryOptions model
				deleteCustomControllibraryOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				deleteCustomControllibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomControllibrary(deleteCustomControllibraryOptions *DeleteCustomControllibraryOptions)`, func() {
		deleteCustomControllibraryPath := "/instances/testString/v3/control_libraries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControllibraryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke DeleteCustomControllibrary successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCustomControllibraryOptions model
				deleteCustomControllibraryOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				deleteCustomControllibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibraryWithContext(ctx, deleteCustomControllibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.DeleteCustomControllibraryWithContext(ctx, deleteCustomControllibraryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomControllibraryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "control_library_name": "ControlLibraryName", "control_library_description": "ControlLibraryDescription", "control_library_type": "predefined", "version_group_label": "VersionGroupLabel", "control_library_version": "ControlLibraryVersion", "latest": true, "controls_count": 13, "controls": [{"control_name": "ControlName", "control_id": "ControlID", "control_description": "ControlDescription", "control_category": "ControlCategory", "control_parent": "ControlParent", "control_severity": "ControlSeverity", "control_tags": ["ControlTags"], "control_specifications": [{"id": "ID", "responsibility": "user", "component_id": "ComponentID", "environment": "Environment", "description": "Description", "assessments_count": 16, "assessments": [{"assessment_id": "AssessmentID", "assessment_method": "AssessmentMethod", "assessment_type": "AssessmentType", "assessment_description": "AssessmentDescription", "parameter_count": 14, "parameters": [{"parameter_name": "ParameterName", "parameter_display_name": "ParameterDisplayName", "parameter_type": "numeric"}]}]}], "control_docs": {"control_docs_id": "ControlDocsID", "control_docs_type": "ControlDocsType"}, "status": "enabled"}]}`)
				}))
			})
			It(`Invoke DeleteCustomControllibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibrary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCustomControllibraryOptions model
				deleteCustomControllibraryOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				deleteCustomControllibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCustomControllibrary with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControllibraryOptions model
				deleteCustomControllibraryOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				deleteCustomControllibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCustomControllibraryOptions model with no property values
				deleteCustomControllibraryOptionsModelNew := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModelNew)
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
			It(`Invoke DeleteCustomControllibrary successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomControllibraryOptions model
				deleteCustomControllibraryOptionsModel := new(sccphoenixcomplianceapisv1.DeleteCustomControllibraryOptions)
				deleteCustomControllibraryOptionsModel.ControlLibrariesID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.InstanceID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCustomControllibraryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.DeleteCustomControllibrary(deleteCustomControllibraryOptionsModel)
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
		createScanPath := "/instances/testString/v3/scans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createScanPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateScan with error: Operation response processing error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				createScanOptionsModel.InstanceID = core.StringPtr("testString")
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.TransactionID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				sccPhoenixComplianceApisService.EnableRetries(0, 0)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
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
		createScanPath := "/instances/testString/v3/scans"
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "attachment_id": "AttachmentID", "report_id": "ReportID", "status": "Status", "last_scan_time": "LastScanTime", "next_scan_time": "NextScanTime", "scan_type": "ScanType", "occurence": 9}`)
				}))
			})
			It(`Invoke CreateScan successfully with retries`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())
				sccPhoenixComplianceApisService.EnableRetries(0, 0)

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				createScanOptionsModel.InstanceID = core.StringPtr("testString")
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.TransactionID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := sccPhoenixComplianceApisService.CreateScanWithContext(ctx, createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				sccPhoenixComplianceApisService.DisableRetries()
				result, response, operationErr := sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = sccPhoenixComplianceApisService.CreateScanWithContext(ctx, createScanOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "attachment_id": "AttachmentID", "report_id": "ReportID", "status": "Status", "last_scan_time": "LastScanTime", "next_scan_time": "NextScanTime", "scan_type": "ScanType", "occurence": 9}`)
				}))
			})
			It(`Invoke CreateScan successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := sccPhoenixComplianceApisService.CreateScan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				createScanOptionsModel.InstanceID = core.StringPtr("testString")
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.TransactionID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateScan with error: Operation validation and request error`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				createScanOptionsModel.InstanceID = core.StringPtr("testString")
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.TransactionID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := sccPhoenixComplianceApisService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateScanOptions model with no property values
				createScanOptionsModelNew := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = sccPhoenixComplianceApisService.CreateScan(createScanOptionsModelNew)
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
			It(`Invoke CreateScan successfully`, func() {
				sccPhoenixComplianceApisService, serviceErr := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(sccPhoenixComplianceApisService).ToNot(BeNil())

				// Construct an instance of the CreateScanOptions model
				createScanOptionsModel := new(sccphoenixcomplianceapisv1.CreateScanOptions)
				createScanOptionsModel.InstanceID = core.StringPtr("testString")
				createScanOptionsModel.AttachmentID = core.StringPtr("testString")
				createScanOptionsModel.TransactionID = core.StringPtr("testString")
				createScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := sccPhoenixComplianceApisService.CreateScan(createScanOptionsModel)
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
			sccPhoenixComplianceApisService, _ := sccphoenixcomplianceapisv1.NewSccPhoenixComplianceApisV1(&sccphoenixcomplianceapisv1.SccPhoenixComplianceApisV1Options{
				URL:           "http://sccphoenixcomplianceapisv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				Expect(profileControlsInRequestModel).ToNot(BeNil())
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")
				Expect(profileControlsInRequestModel.ControlLibraryID).To(Equal(core.StringPtr("testString")))
				Expect(profileControlsInRequestModel.ControlID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				Expect(defaultParametersModel).ToNot(BeNil())
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")
				Expect(defaultParametersModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the AddProfileOptions model
				profilesID := "testString"
				instanceID := "testString"
				addProfileOptionsModel := sccPhoenixComplianceApisService.NewAddProfileOptions(profilesID, instanceID)
				addProfileOptionsModel.SetProfilesID("testString")
				addProfileOptionsModel.SetInstanceID("testString")
				addProfileOptionsModel.SetProfileName("testString")
				addProfileOptionsModel.SetProfileDescription("testString")
				addProfileOptionsModel.SetProfileType("predefined")
				addProfileOptionsModel.SetProfileVersion("testString")
				addProfileOptionsModel.SetLatest(true)
				addProfileOptionsModel.SetVersionGroupLabel("testString")
				addProfileOptionsModel.SetControls([]sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel})
				addProfileOptionsModel.SetDefaultParameters([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel})
				addProfileOptionsModel.SetTransactionID("testString")
				addProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addProfileOptionsModel).ToNot(BeNil())
				Expect(addProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("predefined")))
				Expect(addProfileOptionsModel.ProfileVersion).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(addProfileOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.Controls).To(Equal([]sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}))
				Expect(addProfileOptionsModel.DefaultParameters).To(Equal([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}))
				Expect(addProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCheckProfileAttachmnetsOptions successfully`, func() {
				// Construct an instance of the CheckProfileAttachmnetsOptions model
				profilesID := "testString"
				instanceID := "testString"
				checkProfileAttachmnetsOptionsModel := sccPhoenixComplianceApisService.NewCheckProfileAttachmnetsOptions(profilesID, instanceID)
				checkProfileAttachmnetsOptionsModel.SetProfilesID("testString")
				checkProfileAttachmnetsOptionsModel.SetInstanceID("testString")
				checkProfileAttachmnetsOptionsModel.SetTransactionID("testString")
				checkProfileAttachmnetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkProfileAttachmnetsOptionsModel).ToNot(BeNil())
				Expect(checkProfileAttachmnetsOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(checkProfileAttachmnetsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(checkProfileAttachmnetsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(checkProfileAttachmnetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAttachmentOptions successfully`, func() {
				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				Expect(scopePayloadModel).ToNot(BeNil())
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")
				Expect(scopePayloadModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(scopePayloadModel.ScopeType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				Expect(parameterDetailsModel).ToNot(BeNil())
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				Expect(parameterDetailsModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.ParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(parameterDetailsModel.ParameterValue).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.Parameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{"testString"}))

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				Expect(attachmentsNotificationsPayloadModel).ToNot(BeNil())
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPayloadModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(attachmentsNotificationsPayloadModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the AttachmentPayload model
				attachmentPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentPayload)
				Expect(attachmentPayloadModel).ToNot(BeNil())
				attachmentPayloadModel.ID = core.StringPtr("testString")
				attachmentPayloadModel.AccountID = core.StringPtr("testString")
				attachmentPayloadModel.IncludedScope = scopePayloadModel
				attachmentPayloadModel.Exclusions = []sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}
				attachmentPayloadModel.CreatedBy = core.StringPtr("testString")
				attachmentPayloadModel.CreatedOn = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedBy = core.StringPtr("testString")
				attachmentPayloadModel.UpdatedOn = core.StringPtr("testString")
				attachmentPayloadModel.Status = core.StringPtr("enabled")
				attachmentPayloadModel.AttachmentParameters = []sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}
				attachmentPayloadModel.AttachmentNotifications = attachmentsNotificationsPayloadModel
				Expect(attachmentPayloadModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.IncludedScope).To(Equal(scopePayloadModel))
				Expect(attachmentPayloadModel.Exclusions).To(Equal([]sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}))
				Expect(attachmentPayloadModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.CreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.UpdatedOn).To(Equal(core.StringPtr("testString")))
				Expect(attachmentPayloadModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(attachmentPayloadModel.AttachmentParameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}))
				Expect(attachmentPayloadModel.AttachmentNotifications).To(Equal(attachmentsNotificationsPayloadModel))

				// Construct an instance of the CreateAttachmentOptions model
				profilesID := "testString"
				instanceID := "testString"
				createAttachmentOptionsModel := sccPhoenixComplianceApisService.NewCreateAttachmentOptions(profilesID, instanceID)
				createAttachmentOptionsModel.SetProfilesID("testString")
				createAttachmentOptionsModel.SetInstanceID("testString")
				createAttachmentOptionsModel.SetAttachments([]sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel})
				createAttachmentOptionsModel.SetTransactionID("testString")
				createAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAttachmentOptionsModel).ToNot(BeNil())
				Expect(createAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Attachments).To(Equal([]sccphoenixcomplianceapisv1.AttachmentPayload{*attachmentPayloadModel}))
				Expect(createAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				Expect(implementationPayloadModel).ToNot(BeNil())
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				Expect(implementationPayloadModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentMethod).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentDescription).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationPayloadModel.Parameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}
				Expect(controlSpecificationsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("testString")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				Expect(controlsInControlLibRequestPayloadModel).ToNot(BeNil())
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibRequestPayloadModel.ControlName).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlID).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlDescription).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlCategory).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlSeverity).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlTags).To(Equal([]string{"testString"}))
				Expect(controlsInControlLibRequestPayloadModel.ControlSpecifications).To(Equal([]sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibRequestPayloadModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibRequestPayloadModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the CreateCustomControlLibraryOptions model
				instanceID := "testString"
				createCustomControlLibraryOptionsModel := sccPhoenixComplianceApisService.NewCreateCustomControlLibraryOptions(instanceID)
				createCustomControlLibraryOptionsModel.SetInstanceID("testString")
				createCustomControlLibraryOptionsModel.SetID("testString")
				createCustomControlLibraryOptionsModel.SetAccountID("testString")
				createCustomControlLibraryOptionsModel.SetControlLibraryName("testString")
				createCustomControlLibraryOptionsModel.SetControlLibraryDescription("testString")
				createCustomControlLibraryOptionsModel.SetControlLibraryType("predefined")
				createCustomControlLibraryOptionsModel.SetVersionGroupLabel("testString")
				createCustomControlLibraryOptionsModel.SetControlLibraryVersion("testString")
				createCustomControlLibraryOptionsModel.SetLatest(true)
				createCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				createCustomControlLibraryOptionsModel.SetControls([]sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel})
				createCustomControlLibraryOptionsModel.SetTransactionID("testString")
				createCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(createCustomControlLibraryOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryName).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryDescription).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("predefined")))
				Expect(createCustomControlLibraryOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(createCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createCustomControlLibraryOptionsModel.Controls).To(Equal([]sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}))
				Expect(createCustomControlLibraryOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileOptions successfully`, func() {
				// Construct an instance of the ProfileControlsInRequest model
				profileControlsInRequestModel := new(sccphoenixcomplianceapisv1.ProfileControlsInRequest)
				Expect(profileControlsInRequestModel).ToNot(BeNil())
				profileControlsInRequestModel.ControlLibraryID = core.StringPtr("testString")
				profileControlsInRequestModel.ControlID = core.StringPtr("testString")
				Expect(profileControlsInRequestModel.ControlLibraryID).To(Equal(core.StringPtr("testString")))
				Expect(profileControlsInRequestModel.ControlID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				Expect(defaultParametersModel).ToNot(BeNil())
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")
				Expect(defaultParametersModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the CreateProfileOptions model
				instanceID := "testString"
				createProfileOptionsModel := sccPhoenixComplianceApisService.NewCreateProfileOptions(instanceID)
				createProfileOptionsModel.SetInstanceID("testString")
				createProfileOptionsModel.SetProfileName("testString")
				createProfileOptionsModel.SetProfileDescription("testString")
				createProfileOptionsModel.SetProfileType("predefined")
				createProfileOptionsModel.SetProfileVersion("testString")
				createProfileOptionsModel.SetLatest(true)
				createProfileOptionsModel.SetVersionGroupLabel("testString")
				createProfileOptionsModel.SetControls([]sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel})
				createProfileOptionsModel.SetDefaultParameters([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel})
				createProfileOptionsModel.SetTransactionID("testString")
				createProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileOptionsModel).ToNot(BeNil())
				Expect(createProfileOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.ProfileName).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.ProfileDescription).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("predefined")))
				Expect(createProfileOptionsModel.ProfileVersion).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(createProfileOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Controls).To(Equal([]sccphoenixcomplianceapisv1.ProfileControlsInRequest{*profileControlsInRequestModel}))
				Expect(createProfileOptionsModel.DefaultParameters).To(Equal([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}))
				Expect(createProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateScanOptions successfully`, func() {
				// Construct an instance of the CreateScanOptions model
				instanceID := "testString"
				createScanOptionsModel := sccPhoenixComplianceApisService.NewCreateScanOptions(instanceID)
				createScanOptionsModel.SetInstanceID("testString")
				createScanOptionsModel.SetAttachmentID("testString")
				createScanOptionsModel.SetTransactionID("testString")
				createScanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createScanOptionsModel).ToNot(BeNil())
				Expect(createScanOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createScanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomControllibraryOptions successfully`, func() {
				// Construct an instance of the DeleteCustomControllibraryOptions model
				controlLibrariesID := "testString"
				instanceID := "testString"
				deleteCustomControllibraryOptionsModel := sccPhoenixComplianceApisService.NewDeleteCustomControllibraryOptions(controlLibrariesID, instanceID)
				deleteCustomControllibraryOptionsModel.SetControlLibrariesID("testString")
				deleteCustomControllibraryOptionsModel.SetInstanceID("testString")
				deleteCustomControllibraryOptionsModel.SetTransactionID("testString")
				deleteCustomControllibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomControllibraryOptionsModel).ToNot(BeNil())
				Expect(deleteCustomControllibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControllibraryOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControllibraryOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomControllibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomProfileOptions successfully`, func() {
				// Construct an instance of the DeleteCustomProfileOptions model
				profilesID := "testString"
				instanceID := "testString"
				deleteCustomProfileOptionsModel := sccPhoenixComplianceApisService.NewDeleteCustomProfileOptions(profilesID, instanceID)
				deleteCustomProfileOptionsModel.SetProfilesID("testString")
				deleteCustomProfileOptionsModel.SetInstanceID("testString")
				deleteCustomProfileOptionsModel.SetTransactionID("testString")
				deleteCustomProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomProfileOptionsModel).ToNot(BeNil())
				Expect(deleteCustomProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileAttachmnetOptions successfully`, func() {
				// Construct an instance of the DeleteProfileAttachmnetOptions model
				profilesID := "testString"
				attachmentID := "testString"
				instanceID := "testString"
				deleteProfileAttachmnetOptionsModel := sccPhoenixComplianceApisService.NewDeleteProfileAttachmnetOptions(profilesID, attachmentID, instanceID)
				deleteProfileAttachmnetOptionsModel.SetProfilesID("testString")
				deleteProfileAttachmnetOptionsModel.SetAttachmentID("testString")
				deleteProfileAttachmnetOptionsModel.SetInstanceID("testString")
				deleteProfileAttachmnetOptionsModel.SetTransactionID("testString")
				deleteProfileAttachmnetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileAttachmnetOptionsModel).ToNot(BeNil())
				Expect(deleteProfileAttachmnetOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmnetOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmnetOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmnetOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileAttachmnetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetControlLibraryOptions successfully`, func() {
				// Construct an instance of the GetControlLibraryOptions model
				controlLibrariesID := "testString"
				instanceID := "testString"
				getControlLibraryOptionsModel := sccPhoenixComplianceApisService.NewGetControlLibraryOptions(controlLibrariesID, instanceID)
				getControlLibraryOptionsModel.SetControlLibrariesID("testString")
				getControlLibraryOptionsModel.SetInstanceID("testString")
				getControlLibraryOptionsModel.SetTransactionID("testString")
				getControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getControlLibraryOptionsModel).ToNot(BeNil())
				Expect(getControlLibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetParametersByNameOptions successfully`, func() {
				// Construct an instance of the GetParametersByNameOptions model
				profilesID := "testString"
				attachmentID := "testString"
				parameterName := "testString"
				instanceID := "testString"
				getParametersByNameOptionsModel := sccPhoenixComplianceApisService.NewGetParametersByNameOptions(profilesID, attachmentID, parameterName, instanceID)
				getParametersByNameOptionsModel.SetProfilesID("testString")
				getParametersByNameOptionsModel.SetAttachmentID("testString")
				getParametersByNameOptionsModel.SetParameterName("testString")
				getParametersByNameOptionsModel.SetInstanceID("testString")
				getParametersByNameOptionsModel.SetTransactionID("testString")
				getParametersByNameOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getParametersByNameOptionsModel).ToNot(BeNil())
				Expect(getParametersByNameOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(getParametersByNameOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(getParametersByNameOptionsModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(getParametersByNameOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getParametersByNameOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getParametersByNameOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileAttachmnetOptions successfully`, func() {
				// Construct an instance of the GetProfileAttachmnetOptions model
				profilesID := "testString"
				attachmentID := "testString"
				instanceID := "testString"
				getProfileAttachmnetOptionsModel := sccPhoenixComplianceApisService.NewGetProfileAttachmnetOptions(profilesID, attachmentID, instanceID)
				getProfileAttachmnetOptionsModel.SetProfilesID("testString")
				getProfileAttachmnetOptionsModel.SetAttachmentID("testString")
				getProfileAttachmnetOptionsModel.SetInstanceID("testString")
				getProfileAttachmnetOptionsModel.SetTransactionID("testString")
				getProfileAttachmnetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileAttachmnetOptionsModel).ToNot(BeNil())
				Expect(getProfileAttachmnetOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmnetOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmnetOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmnetOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileAttachmnetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileOptions successfully`, func() {
				// Construct an instance of the GetProfileOptions model
				profilesID := "testString"
				instanceID := "testString"
				getProfileOptionsModel := sccPhoenixComplianceApisService.NewGetProfileOptions(profilesID, instanceID)
				getProfileOptionsModel.SetProfilesID("testString")
				getProfileOptionsModel.SetInstanceID("testString")
				getProfileOptionsModel.SetTransactionID("testString")
				getProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileOptionsModel).ToNot(BeNil())
				Expect(getProfileOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAttachmentParametersOptions successfully`, func() {
				// Construct an instance of the ListAttachmentParametersOptions model
				profilesID := "testString"
				attachmentID := "testString"
				instanceID := "testString"
				listAttachmentParametersOptionsModel := sccPhoenixComplianceApisService.NewListAttachmentParametersOptions(profilesID, attachmentID, instanceID)
				listAttachmentParametersOptionsModel.SetProfilesID("testString")
				listAttachmentParametersOptionsModel.SetAttachmentID("testString")
				listAttachmentParametersOptionsModel.SetInstanceID("testString")
				listAttachmentParametersOptionsModel.SetTransactionID("testString")
				listAttachmentParametersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAttachmentParametersOptionsModel).ToNot(BeNil())
				Expect(listAttachmentParametersOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentParametersOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentParametersOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentParametersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAttachmentParametersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListControlLibrariesOptions successfully`, func() {
				// Construct an instance of the ListControlLibrariesOptions model
				instanceID := "testString"
				listControlLibrariesOptionsModel := sccPhoenixComplianceApisService.NewListControlLibrariesOptions(instanceID)
				listControlLibrariesOptionsModel.SetInstanceID("testString")
				listControlLibrariesOptionsModel.SetTransactionID("testString")
				listControlLibrariesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listControlLibrariesOptionsModel).ToNot(BeNil())
				Expect(listControlLibrariesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listControlLibrariesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				instanceID := "testString"
				listProfilesOptionsModel := sccPhoenixComplianceApisService.NewListProfilesOptions(instanceID)
				listProfilesOptionsModel.SetInstanceID("testString")
				listProfilesOptionsModel.SetTransactionID("testString")
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceAttachmentOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ReplaceAttachmentOptions model
				profilesID := "testString"
				attachmentID := "testString"
				instanceID := "testString"
				replaceAttachmentOptionsModel := sccPhoenixComplianceApisService.NewReplaceAttachmentOptions(profilesID, attachmentID, instanceID)
				replaceAttachmentOptionsModel.SetProfilesID("testString")
				replaceAttachmentOptionsModel.SetAttachmentID("testString")
				replaceAttachmentOptionsModel.SetInstanceID("testString")
				replaceAttachmentOptionsModel.SetParameterName("testString")
				replaceAttachmentOptionsModel.SetParameterDisplayName("testString")
				replaceAttachmentOptionsModel.SetParameterType("numeric")
				replaceAttachmentOptionsModel.SetParameterValue("testString")
				replaceAttachmentOptionsModel.SetAssessmentType("testString")
				replaceAttachmentOptionsModel.SetAssessmentID("testString")
				replaceAttachmentOptionsModel.SetParameters([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel})
				replaceAttachmentOptionsModel.SetTransactionID("testString")
				replaceAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAttachmentOptionsModel).ToNot(BeNil())
				Expect(replaceAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.ParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(replaceAttachmentOptionsModel.ParameterValue).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.Parameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))
				Expect(replaceAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceAttachmnetParametersByNameOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ReplaceAttachmnetParametersByNameOptions model
				profilesID := "testString"
				attachmentID := "testString"
				parameterName := "testString"
				instanceID := "testString"
				replaceAttachmnetParametersByNameOptionsModel := sccPhoenixComplianceApisService.NewReplaceAttachmnetParametersByNameOptions(profilesID, attachmentID, parameterName, instanceID)
				replaceAttachmnetParametersByNameOptionsModel.SetProfilesID("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetAttachmentID("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetParameterName("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetInstanceID("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewParameterName("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewParameterDisplayName("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewParameterType("numeric")
				replaceAttachmnetParametersByNameOptionsModel.SetNewParameterValue("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewAssessmentType("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewAssessmentID("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetNewParameters([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel})
				replaceAttachmnetParametersByNameOptionsModel.SetTransactionID("testString")
				replaceAttachmnetParametersByNameOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAttachmnetParametersByNameOptionsModel).ToNot(BeNil())
				Expect(replaceAttachmnetParametersByNameOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewParameterName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewParameterValue).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewAssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewAssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.NewParameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))
				Expect(replaceAttachmnetParametersByNameOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAttachmnetParametersByNameOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceCustomControlLibraryOptions successfully`, func() {
				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ImplementationPayload model
				implementationPayloadModel := new(sccphoenixcomplianceapisv1.ImplementationPayload)
				Expect(implementationPayloadModel).ToNot(BeNil())
				implementationPayloadModel.AssessmentID = core.StringPtr("testString")
				implementationPayloadModel.AssessmentMethod = core.StringPtr("testString")
				implementationPayloadModel.AssessmentType = core.StringPtr("testString")
				implementationPayloadModel.AssessmentDescription = core.StringPtr("testString")
				implementationPayloadModel.ParameterCount = core.Int64Ptr(int64(38))
				implementationPayloadModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				Expect(implementationPayloadModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentMethod).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.AssessmentDescription).To(Equal(core.StringPtr("testString")))
				Expect(implementationPayloadModel.ParameterCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(implementationPayloadModel.Parameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the ControlSpecifications model
				controlSpecificationsModel := new(sccphoenixcomplianceapisv1.ControlSpecifications)
				Expect(controlSpecificationsModel).ToNot(BeNil())
				controlSpecificationsModel.ID = core.StringPtr("testString")
				controlSpecificationsModel.Responsibility = core.StringPtr("user")
				controlSpecificationsModel.ComponentID = core.StringPtr("testString")
				controlSpecificationsModel.Environment = core.StringPtr("testString")
				controlSpecificationsModel.Description = core.StringPtr("testString")
				controlSpecificationsModel.AssessmentsCount = core.Int64Ptr(int64(38))
				controlSpecificationsModel.Assessments = []sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}
				Expect(controlSpecificationsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Responsibility).To(Equal(core.StringPtr("user")))
				Expect(controlSpecificationsModel.ComponentID).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Environment).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(controlSpecificationsModel.AssessmentsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(controlSpecificationsModel.Assessments).To(Equal([]sccphoenixcomplianceapisv1.ImplementationPayload{*implementationPayloadModel}))

				// Construct an instance of the ControlDocs model
				controlDocsModel := new(sccphoenixcomplianceapisv1.ControlDocs)
				Expect(controlDocsModel).ToNot(BeNil())
				controlDocsModel.ControlDocsID = core.StringPtr("testString")
				controlDocsModel.ControlDocsType = core.StringPtr("testString")
				Expect(controlDocsModel.ControlDocsID).To(Equal(core.StringPtr("testString")))
				Expect(controlDocsModel.ControlDocsType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ControlsInControlLibRequestPayload model
				controlsInControlLibRequestPayloadModel := new(sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload)
				Expect(controlsInControlLibRequestPayloadModel).ToNot(BeNil())
				controlsInControlLibRequestPayloadModel.ControlName = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlID = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlDescription = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlCategory = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlParent = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlSeverity = core.StringPtr("testString")
				controlsInControlLibRequestPayloadModel.ControlTags = []string{"testString"}
				controlsInControlLibRequestPayloadModel.ControlSpecifications = []sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}
				controlsInControlLibRequestPayloadModel.ControlDocs = controlDocsModel
				controlsInControlLibRequestPayloadModel.Status = core.StringPtr("enabled")
				Expect(controlsInControlLibRequestPayloadModel.ControlName).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlID).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlDescription).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlCategory).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlParent).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlSeverity).To(Equal(core.StringPtr("testString")))
				Expect(controlsInControlLibRequestPayloadModel.ControlTags).To(Equal([]string{"testString"}))
				Expect(controlsInControlLibRequestPayloadModel.ControlSpecifications).To(Equal([]sccphoenixcomplianceapisv1.ControlSpecifications{*controlSpecificationsModel}))
				Expect(controlsInControlLibRequestPayloadModel.ControlDocs).To(Equal(controlDocsModel))
				Expect(controlsInControlLibRequestPayloadModel.Status).To(Equal(core.StringPtr("enabled")))

				// Construct an instance of the ReplaceCustomControlLibraryOptions model
				controlLibrariesID := "testString"
				instanceID := "testString"
				replaceCustomControlLibraryOptionsModel := sccPhoenixComplianceApisService.NewReplaceCustomControlLibraryOptions(controlLibrariesID, instanceID)
				replaceCustomControlLibraryOptionsModel.SetControlLibrariesID("testString")
				replaceCustomControlLibraryOptionsModel.SetInstanceID("testString")
				replaceCustomControlLibraryOptionsModel.SetID("testString")
				replaceCustomControlLibraryOptionsModel.SetAccountID("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryName("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryDescription("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryType("predefined")
				replaceCustomControlLibraryOptionsModel.SetVersionGroupLabel("testString")
				replaceCustomControlLibraryOptionsModel.SetControlLibraryVersion("testString")
				replaceCustomControlLibraryOptionsModel.SetLatest(true)
				replaceCustomControlLibraryOptionsModel.SetControlsCount(int64(38))
				replaceCustomControlLibraryOptionsModel.SetControls([]sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel})
				replaceCustomControlLibraryOptionsModel.SetTransactionID("testString")
				replaceCustomControlLibraryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceCustomControlLibraryOptionsModel).ToNot(BeNil())
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibrariesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryName).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryDescription).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryType).To(Equal(core.StringPtr("predefined")))
				Expect(replaceCustomControlLibraryOptionsModel.VersionGroupLabel).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.ControlLibraryVersion).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(replaceCustomControlLibraryOptionsModel.ControlsCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceCustomControlLibraryOptionsModel.Controls).To(Equal([]sccphoenixcomplianceapisv1.ControlsInControlLibRequestPayload{*controlsInControlLibRequestPayloadModel}))
				Expect(replaceCustomControlLibraryOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCustomControlLibraryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileAttachmentOptions successfully`, func() {
				// Construct an instance of the ScopePayload model
				scopePayloadModel := new(sccphoenixcomplianceapisv1.ScopePayload)
				Expect(scopePayloadModel).ToNot(BeNil())
				scopePayloadModel.ScopeID = core.StringPtr("testString")
				scopePayloadModel.ScopeType = core.StringPtr("testString")
				Expect(scopePayloadModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(scopePayloadModel.ScopeType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ParameterInfo model
				parameterInfoModel := new(sccphoenixcomplianceapisv1.ParameterInfo)
				Expect(parameterInfoModel).ToNot(BeNil())
				parameterInfoModel.ParameterName = core.StringPtr("testString")
				parameterInfoModel.ParameterDisplayName = core.StringPtr("testString")
				parameterInfoModel.ParameterType = core.StringPtr("numeric")
				Expect(parameterInfoModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterInfoModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ParameterDetails model
				parameterDetailsModel := new(sccphoenixcomplianceapisv1.ParameterDetails)
				Expect(parameterDetailsModel).ToNot(BeNil())
				parameterDetailsModel.ParameterName = core.StringPtr("testString")
				parameterDetailsModel.ParameterDisplayName = core.StringPtr("testString")
				parameterDetailsModel.ParameterType = core.StringPtr("numeric")
				parameterDetailsModel.ParameterValue = core.StringPtr("testString")
				parameterDetailsModel.AssessmentType = core.StringPtr("testString")
				parameterDetailsModel.AssessmentID = core.StringPtr("testString")
				parameterDetailsModel.Parameters = []sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}
				Expect(parameterDetailsModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.ParameterType).To(Equal(core.StringPtr("numeric")))
				Expect(parameterDetailsModel.ParameterValue).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(parameterDetailsModel.Parameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterInfo{*parameterInfoModel}))

				// Construct an instance of the FailedControls model
				failedControlsModel := new(sccphoenixcomplianceapisv1.FailedControls)
				Expect(failedControlsModel).ToNot(BeNil())
				failedControlsModel.ThresholdLimit = core.Int64Ptr(int64(38))
				failedControlsModel.FailedControlIds = []string{"testString"}
				Expect(failedControlsModel.ThresholdLimit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(failedControlsModel.FailedControlIds).To(Equal([]string{"testString"}))

				// Construct an instance of the AttachmentsNotificationsPayload model
				attachmentsNotificationsPayloadModel := new(sccphoenixcomplianceapisv1.AttachmentsNotificationsPayload)
				Expect(attachmentsNotificationsPayloadModel).ToNot(BeNil())
				attachmentsNotificationsPayloadModel.Enabled = core.BoolPtr(true)
				attachmentsNotificationsPayloadModel.Controls = failedControlsModel
				Expect(attachmentsNotificationsPayloadModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(attachmentsNotificationsPayloadModel.Controls).To(Equal(failedControlsModel))

				// Construct an instance of the ReplaceProfileAttachmentOptions model
				profilesID := "testString"
				attachmentID := "testString"
				instanceID := "testString"
				replaceProfileAttachmentOptionsModel := sccPhoenixComplianceApisService.NewReplaceProfileAttachmentOptions(profilesID, attachmentID, instanceID)
				replaceProfileAttachmentOptionsModel.SetProfilesID("testString")
				replaceProfileAttachmentOptionsModel.SetAttachmentID("testString")
				replaceProfileAttachmentOptionsModel.SetInstanceID("testString")
				replaceProfileAttachmentOptionsModel.SetID("testString")
				replaceProfileAttachmentOptionsModel.SetAccountID("testString")
				replaceProfileAttachmentOptionsModel.SetIncludedScope(scopePayloadModel)
				replaceProfileAttachmentOptionsModel.SetExclusions([]sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel})
				replaceProfileAttachmentOptionsModel.SetCreatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetCreatedOn("testString")
				replaceProfileAttachmentOptionsModel.SetUpdatedBy("testString")
				replaceProfileAttachmentOptionsModel.SetUpdatedOn("testString")
				replaceProfileAttachmentOptionsModel.SetStatus("enabled")
				replaceProfileAttachmentOptionsModel.SetAttachmentParameters([]sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel})
				replaceProfileAttachmentOptionsModel.SetAttachmentNotifications(attachmentsNotificationsPayloadModel)
				replaceProfileAttachmentOptionsModel.SetTransactionID("testString")
				replaceProfileAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProfileAttachmentOptionsModel).ToNot(BeNil())
				Expect(replaceProfileAttachmentOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.IncludedScope).To(Equal(scopePayloadModel))
				Expect(replaceProfileAttachmentOptionsModel.Exclusions).To(Equal([]sccphoenixcomplianceapisv1.ScopePayload{*scopePayloadModel}))
				Expect(replaceProfileAttachmentOptionsModel.CreatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.CreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedBy).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.UpdatedOn).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Status).To(Equal(core.StringPtr("enabled")))
				Expect(replaceProfileAttachmentOptionsModel.AttachmentParameters).To(Equal([]sccphoenixcomplianceapisv1.ParameterDetails{*parameterDetailsModel}))
				Expect(replaceProfileAttachmentOptionsModel.AttachmentNotifications).To(Equal(attachmentsNotificationsPayloadModel))
				Expect(replaceProfileAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceProfileParametersOptions successfully`, func() {
				// Construct an instance of the DefaultParameters model
				defaultParametersModel := new(sccphoenixcomplianceapisv1.DefaultParameters)
				Expect(defaultParametersModel).ToNot(BeNil())
				defaultParametersModel.AssessmentType = core.StringPtr("testString")
				defaultParametersModel.AssessmentID = core.StringPtr("testString")
				defaultParametersModel.ParameterName = core.StringPtr("testString")
				defaultParametersModel.ParameterDefaultValue = core.StringPtr("testString")
				defaultParametersModel.ParameterDisplayName = core.StringPtr("testString")
				defaultParametersModel.ParameterType = core.StringPtr("numeric")
				Expect(defaultParametersModel.AssessmentType).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.AssessmentID).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterDisplayName).To(Equal(core.StringPtr("testString")))
				Expect(defaultParametersModel.ParameterType).To(Equal(core.StringPtr("numeric")))

				// Construct an instance of the ReplaceProfileParametersOptions model
				profilesID := "testString"
				instanceID := "testString"
				replaceProfileParametersOptionsModel := sccPhoenixComplianceApisService.NewReplaceProfileParametersOptions(profilesID, instanceID)
				replaceProfileParametersOptionsModel.SetProfilesID("testString")
				replaceProfileParametersOptionsModel.SetInstanceID("testString")
				replaceProfileParametersOptionsModel.SetID("testString")
				replaceProfileParametersOptionsModel.SetDefaultParameters([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel})
				replaceProfileParametersOptionsModel.SetTransactionID("testString")
				replaceProfileParametersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceProfileParametersOptionsModel).ToNot(BeNil())
				Expect(replaceProfileParametersOptionsModel.ProfilesID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileParametersOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileParametersOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileParametersOptionsModel.DefaultParameters).To(Equal([]sccphoenixcomplianceapisv1.DefaultParameters{*defaultParametersModel}))
				Expect(replaceProfileParametersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceProfileParametersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

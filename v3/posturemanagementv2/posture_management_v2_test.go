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

package posturemanagementv2_test

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
	"github.com/IBM/scc-go-sdk/v3/posturemanagementv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PostureManagementV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(postureManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(postureManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
				URL: "https://posturemanagementv2/api",
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
				"POSTURE_MANAGEMENT_URL":       "https://posturemanagementv2/api",
				"POSTURE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2UsingExternalConfig(&posturemanagementv2.PostureManagementV2Options{})
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2UsingExternalConfig(&posturemanagementv2.PostureManagementV2Options{
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2UsingExternalConfig(&posturemanagementv2.PostureManagementV2Options{})
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
				"POSTURE_MANAGEMENT_URL":       "https://posturemanagementv2/api",
				"POSTURE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2UsingExternalConfig(&posturemanagementv2.PostureManagementV2Options{})

			It(`Instantiate service client with error`, func() {
				Expect(postureManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POSTURE_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2UsingExternalConfig(&posturemanagementv2.PostureManagementV2Options{
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
			url, err = posturemanagementv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := posturemanagementv2.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us.compliance.cloud.ibm.com"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := posturemanagementv2.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`CreateCredential(createCredentialOptions *CreateCredentialOptions) - Operation response error`, func() {
		createCredentialPath := "/posture/v2/credentials"
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv2.CreateCredentialOptions)
				createCredentialOptionsModel.Enabled = core.BoolPtr(true)
				createCredentialOptionsModel.Type = core.StringPtr("username_password")
				createCredentialOptionsModel.Name = core.StringPtr("test_create")
				createCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				createCredentialOptionsModel.DisplayFields = newCredentialDisplayFieldsModel
				createCredentialOptionsModel.Group = credentialGroupModel
				createCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
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
		createCredentialPath := "/posture/v2/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialPath))
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
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke CreateCredential successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv2.CreateCredentialOptions)
				createCredentialOptionsModel.Enabled = core.BoolPtr(true)
				createCredentialOptionsModel.Type = core.StringPtr("username_password")
				createCredentialOptionsModel.Name = core.StringPtr("test_create")
				createCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				createCredentialOptionsModel.DisplayFields = newCredentialDisplayFieldsModel
				createCredentialOptionsModel.Group = credentialGroupModel
				createCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke CreateCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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

				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv2.CreateCredentialOptions)
				createCredentialOptionsModel.Enabled = core.BoolPtr(true)
				createCredentialOptionsModel.Type = core.StringPtr("username_password")
				createCredentialOptionsModel.Name = core.StringPtr("test_create")
				createCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				createCredentialOptionsModel.DisplayFields = newCredentialDisplayFieldsModel
				createCredentialOptionsModel.Group = credentialGroupModel
				createCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
				createCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				createCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateCredential(createCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCredential with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv2.CreateCredentialOptions)
				createCredentialOptionsModel.Enabled = core.BoolPtr(true)
				createCredentialOptionsModel.Type = core.StringPtr("username_password")
				createCredentialOptionsModel.Name = core.StringPtr("test_create")
				createCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				createCredentialOptionsModel.DisplayFields = newCredentialDisplayFieldsModel
				createCredentialOptionsModel.Group = credentialGroupModel
				createCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
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
				createCredentialOptionsModelNew := new(posturemanagementv2.CreateCredentialOptions)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsModel := new(posturemanagementv2.CreateCredentialOptions)
				createCredentialOptionsModel.Enabled = core.BoolPtr(true)
				createCredentialOptionsModel.Type = core.StringPtr("username_password")
				createCredentialOptionsModel.Name = core.StringPtr("test_create")
				createCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				createCredentialOptionsModel.DisplayFields = newCredentialDisplayFieldsModel
				createCredentialOptionsModel.Group = credentialGroupModel
				createCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				createCredentialOptionsModel.AccountID = core.StringPtr("testString")
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
	Describe(`ListCredentials(listCredentialsOptions *ListCredentialsOptions) - Operation response error`, func() {
		listCredentialsPath := "/posture/v2/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCredentials with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(posturemanagementv2.ListCredentialsOptions)
				listCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				listCredentialsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listCredentialsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCredentials(listCredentialsOptions *ListCredentialsOptions)`, func() {
		listCredentialsPath := "/posture/v2/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "credentials": [{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}]}`)
				}))
			})
			It(`Invoke ListCredentials successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(posturemanagementv2.ListCredentialsOptions)
				listCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				listCredentialsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listCredentialsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListCredentialsWithContext(ctx, listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListCredentialsWithContext(ctx, listCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "credentials": [{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}]}`)
				}))
			})
			It(`Invoke ListCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(posturemanagementv2.ListCredentialsOptions)
				listCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				listCredentialsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listCredentialsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCredentials with error: Operation request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(posturemanagementv2.ListCredentialsOptions)
				listCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				listCredentialsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listCredentialsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListCredentials(listCredentialsOptionsModel)
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
			It(`Invoke ListCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(posturemanagementv2.ListCredentialsOptions)
				listCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				listCredentialsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listCredentialsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListCredentials(listCredentialsOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.CredentialList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.CredentialList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.CredentialList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.CredentialList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`GetCredential(getCredentialOptions *GetCredentialOptions) - Operation response error`, func() {
		getCredentialPath := "/posture/v2/credentials/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCredential with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCredentialOptions model
				getCredentialOptionsModel := new(posturemanagementv2.GetCredentialOptions)
				getCredentialOptionsModel.ID = core.StringPtr("testString")
				getCredentialOptionsModel.AccountID = core.StringPtr("testString")
				getCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				getCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetCredential(getCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetCredential(getCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCredential(getCredentialOptions *GetCredentialOptions)`, func() {
		getCredentialPath := "/posture/v2/credentials/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke GetCredential successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCredentialOptions model
				getCredentialOptionsModel := new(posturemanagementv2.GetCredentialOptions)
				getCredentialOptionsModel.ID = core.StringPtr("testString")
				getCredentialOptionsModel.AccountID = core.StringPtr("testString")
				getCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				getCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetCredentialWithContext(ctx, getCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetCredential(getCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetCredentialWithContext(ctx, getCredentialOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke GetCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetCredential(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCredentialOptions model
				getCredentialOptionsModel := new(posturemanagementv2.GetCredentialOptions)
				getCredentialOptionsModel.ID = core.StringPtr("testString")
				getCredentialOptionsModel.AccountID = core.StringPtr("testString")
				getCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				getCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetCredential(getCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCredential with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCredentialOptions model
				getCredentialOptionsModel := new(posturemanagementv2.GetCredentialOptions)
				getCredentialOptionsModel.ID = core.StringPtr("testString")
				getCredentialOptionsModel.AccountID = core.StringPtr("testString")
				getCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				getCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetCredential(getCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCredentialOptions model with no property values
				getCredentialOptionsModelNew := new(posturemanagementv2.GetCredentialOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetCredential(getCredentialOptionsModelNew)
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
			It(`Invoke GetCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCredentialOptions model
				getCredentialOptionsModel := new(posturemanagementv2.GetCredentialOptions)
				getCredentialOptionsModel.ID = core.StringPtr("testString")
				getCredentialOptionsModel.AccountID = core.StringPtr("testString")
				getCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				getCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetCredential(getCredentialOptionsModel)
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
	Describe(`UpdateCredential(updateCredentialOptions *UpdateCredentialOptions) - Operation response error`, func() {
		updateCredentialPath := "/posture/v2/credentials/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCredential with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the UpdateCredentialOptions model
				updateCredentialOptionsModel := new(posturemanagementv2.UpdateCredentialOptions)
				updateCredentialOptionsModel.ID = core.StringPtr("testString")
				updateCredentialOptionsModel.Enabled = core.BoolPtr(true)
				updateCredentialOptionsModel.Type = core.StringPtr("username_password")
				updateCredentialOptionsModel.Name = core.StringPtr("test_create")
				updateCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				updateCredentialOptionsModel.DisplayFields = updateCredentialDisplayFieldsModel
				updateCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.AccountID = core.StringPtr("testString")
				updateCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				updateCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.UpdateCredential(updateCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.UpdateCredential(updateCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCredential(updateCredentialOptions *UpdateCredentialOptions)`, func() {
		updateCredentialPath := "/posture/v2/credentials/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke UpdateCredential successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the UpdateCredentialOptions model
				updateCredentialOptionsModel := new(posturemanagementv2.UpdateCredentialOptions)
				updateCredentialOptionsModel.ID = core.StringPtr("testString")
				updateCredentialOptionsModel.Enabled = core.BoolPtr(true)
				updateCredentialOptionsModel.Type = core.StringPtr("username_password")
				updateCredentialOptionsModel.Name = core.StringPtr("test_create")
				updateCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				updateCredentialOptionsModel.DisplayFields = updateCredentialDisplayFieldsModel
				updateCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.AccountID = core.StringPtr("testString")
				updateCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				updateCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.UpdateCredentialWithContext(ctx, updateCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.UpdateCredential(updateCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.UpdateCredentialWithContext(ctx, updateCredentialOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"enabled": true, "id": "57", "type": "username_password", "name": "test_username", "description": "test_description", "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "created_by": "IBMid-5500081P5M", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "updated_by": "IBMid-5500081P5M", "group": {"id": "ID", "passphrase": "Passphrase"}, "purpose": "discovery_fact_collection_remediation"}`)
				}))
			})
			It(`Invoke UpdateCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.UpdateCredential(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the UpdateCredentialOptions model
				updateCredentialOptionsModel := new(posturemanagementv2.UpdateCredentialOptions)
				updateCredentialOptionsModel.ID = core.StringPtr("testString")
				updateCredentialOptionsModel.Enabled = core.BoolPtr(true)
				updateCredentialOptionsModel.Type = core.StringPtr("username_password")
				updateCredentialOptionsModel.Name = core.StringPtr("test_create")
				updateCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				updateCredentialOptionsModel.DisplayFields = updateCredentialDisplayFieldsModel
				updateCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.AccountID = core.StringPtr("testString")
				updateCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				updateCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.UpdateCredential(updateCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCredential with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the UpdateCredentialOptions model
				updateCredentialOptionsModel := new(posturemanagementv2.UpdateCredentialOptions)
				updateCredentialOptionsModel.ID = core.StringPtr("testString")
				updateCredentialOptionsModel.Enabled = core.BoolPtr(true)
				updateCredentialOptionsModel.Type = core.StringPtr("username_password")
				updateCredentialOptionsModel.Name = core.StringPtr("test_create")
				updateCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				updateCredentialOptionsModel.DisplayFields = updateCredentialDisplayFieldsModel
				updateCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.AccountID = core.StringPtr("testString")
				updateCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				updateCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.UpdateCredential(updateCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCredentialOptions model with no property values
				updateCredentialOptionsModelNew := new(posturemanagementv2.UpdateCredentialOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.UpdateCredential(updateCredentialOptionsModelNew)
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
			It(`Invoke UpdateCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")

				// Construct an instance of the UpdateCredentialOptions model
				updateCredentialOptionsModel := new(posturemanagementv2.UpdateCredentialOptions)
				updateCredentialOptionsModel.ID = core.StringPtr("testString")
				updateCredentialOptionsModel.Enabled = core.BoolPtr(true)
				updateCredentialOptionsModel.Type = core.StringPtr("username_password")
				updateCredentialOptionsModel.Name = core.StringPtr("test_create")
				updateCredentialOptionsModel.Description = core.StringPtr("This credential is used for testing.")
				updateCredentialOptionsModel.DisplayFields = updateCredentialDisplayFieldsModel
				updateCredentialOptionsModel.Purpose = core.StringPtr("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.AccountID = core.StringPtr("testString")
				updateCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				updateCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.UpdateCredential(updateCredentialOptionsModel)
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
	Describe(`DeleteCredential(deleteCredentialOptions *DeleteCredentialOptions)`, func() {
		deleteCredentialPath := "/posture/v2/credentials/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCredentialPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCredential successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := postureManagementService.DeleteCredential(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCredentialOptions model
				deleteCredentialOptionsModel := new(posturemanagementv2.DeleteCredentialOptions)
				deleteCredentialOptionsModel.ID = core.StringPtr("testString")
				deleteCredentialOptionsModel.AccountID = core.StringPtr("testString")
				deleteCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = postureManagementService.DeleteCredential(deleteCredentialOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCredential with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteCredentialOptions model
				deleteCredentialOptionsModel := new(posturemanagementv2.DeleteCredentialOptions)
				deleteCredentialOptionsModel.ID = core.StringPtr("testString")
				deleteCredentialOptionsModel.AccountID = core.StringPtr("testString")
				deleteCredentialOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCredentialOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := postureManagementService.DeleteCredential(deleteCredentialOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCredentialOptions model with no property values
				deleteCredentialOptionsModelNew := new(posturemanagementv2.DeleteCredentialOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = postureManagementService.DeleteCredential(deleteCredentialOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCollector(createCollectorOptions *CreateCollectorOptions) - Operation response error`, func() {
		createCollectorPath := "/posture/v2/collectors"
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv2.CreateCollectorOptions)
				createCollectorOptionsModel.Name = core.StringPtr("IBM-collector-sample")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("ibm")
				createCollectorOptionsModel.Description = core.StringPtr("sample collector")
				createCollectorOptionsModel.Passphrase = core.StringPtr("secret")
				createCollectorOptionsModel.IsUbiImage = core.BoolPtr(true)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
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
		createCollectorPath := "/posture/v2/collectors"
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke CreateCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv2.CreateCollectorOptions)
				createCollectorOptionsModel.Name = core.StringPtr("IBM-collector-sample")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("ibm")
				createCollectorOptionsModel.Description = core.StringPtr("sample collector")
				createCollectorOptionsModel.Passphrase = core.StringPtr("secret")
				createCollectorOptionsModel.IsUbiImage = core.BoolPtr(true)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke CreateCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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
				createCollectorOptionsModel := new(posturemanagementv2.CreateCollectorOptions)
				createCollectorOptionsModel.Name = core.StringPtr("IBM-collector-sample")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("ibm")
				createCollectorOptionsModel.Description = core.StringPtr("sample collector")
				createCollectorOptionsModel.Passphrase = core.StringPtr("secret")
				createCollectorOptionsModel.IsUbiImage = core.BoolPtr(true)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
				createCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				createCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateCollector(createCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv2.CreateCollectorOptions)
				createCollectorOptionsModel.Name = core.StringPtr("IBM-collector-sample")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("ibm")
				createCollectorOptionsModel.Description = core.StringPtr("sample collector")
				createCollectorOptionsModel.Passphrase = core.StringPtr("secret")
				createCollectorOptionsModel.IsUbiImage = core.BoolPtr(true)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
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
				createCollectorOptionsModelNew := new(posturemanagementv2.CreateCollectorOptions)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsModel := new(posturemanagementv2.CreateCollectorOptions)
				createCollectorOptionsModel.Name = core.StringPtr("IBM-collector-sample")
				createCollectorOptionsModel.IsPublic = core.BoolPtr(true)
				createCollectorOptionsModel.ManagedBy = core.StringPtr("ibm")
				createCollectorOptionsModel.Description = core.StringPtr("sample collector")
				createCollectorOptionsModel.Passphrase = core.StringPtr("secret")
				createCollectorOptionsModel.IsUbiImage = core.BoolPtr(true)
				createCollectorOptionsModel.AccountID = core.StringPtr("testString")
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
	Describe(`ListCollectors(listCollectorsOptions *ListCollectorsOptions) - Operation response error`, func() {
		listCollectorsPath := "/posture/v2/collectors"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollectors with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := new(posturemanagementv2.ListCollectorsOptions)
				listCollectorsOptionsModel.AccountID = core.StringPtr("testString")
				listCollectorsOptionsModel.TransactionID = core.StringPtr("testString")
				listCollectorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListCollectors(listCollectorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListCollectors(listCollectorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCollectors(listCollectorsOptions *ListCollectorsOptions)`, func() {
		listCollectorsPath := "/posture/v2/collectors"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectorsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}`)
				}))
			})
			It(`Invoke ListCollectors successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := new(posturemanagementv2.ListCollectorsOptions)
				listCollectorsOptionsModel.AccountID = core.StringPtr("testString")
				listCollectorsOptionsModel.TransactionID = core.StringPtr("testString")
				listCollectorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListCollectorsWithContext(ctx, listCollectorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListCollectors(listCollectorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListCollectorsWithContext(ctx, listCollectorsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCollectorsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}`)
				}))
			})
			It(`Invoke ListCollectors successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListCollectors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := new(posturemanagementv2.ListCollectorsOptions)
				listCollectorsOptionsModel.AccountID = core.StringPtr("testString")
				listCollectorsOptionsModel.TransactionID = core.StringPtr("testString")
				listCollectorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListCollectors(listCollectorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCollectors with error: Operation request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := new(posturemanagementv2.ListCollectorsOptions)
				listCollectorsOptionsModel.AccountID = core.StringPtr("testString")
				listCollectorsOptionsModel.TransactionID = core.StringPtr("testString")
				listCollectorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListCollectors(listCollectorsOptionsModel)
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
			It(`Invoke ListCollectors successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := new(posturemanagementv2.ListCollectorsOptions)
				listCollectorsOptionsModel.AccountID = core.StringPtr("testString")
				listCollectorsOptionsModel.TransactionID = core.StringPtr("testString")
				listCollectorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListCollectors(listCollectorsOptionsModel)
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
	Describe(`GetCollector(getCollectorOptions *GetCollectorOptions) - Operation response error`, func() {
		getCollectorPath := "/posture/v2/collectors/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCollector with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCollectorOptions model
				getCollectorOptionsModel := new(posturemanagementv2.GetCollectorOptions)
				getCollectorOptionsModel.ID = core.StringPtr("testString")
				getCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetCollector(getCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetCollector(getCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCollector(getCollectorOptions *GetCollectorOptions)`, func() {
		getCollectorPath := "/posture/v2/collectors/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectorPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke GetCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCollectorOptions model
				getCollectorOptionsModel := new(posturemanagementv2.GetCollectorOptions)
				getCollectorOptionsModel.ID = core.StringPtr("testString")
				getCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetCollectorWithContext(ctx, getCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetCollector(getCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetCollectorWithContext(ctx, getCollectorOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCollectorPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke GetCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectorOptions model
				getCollectorOptionsModel := new(posturemanagementv2.GetCollectorOptions)
				getCollectorOptionsModel.ID = core.StringPtr("testString")
				getCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetCollector(getCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCollectorOptions model
				getCollectorOptionsModel := new(posturemanagementv2.GetCollectorOptions)
				getCollectorOptionsModel.ID = core.StringPtr("testString")
				getCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetCollector(getCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCollectorOptions model with no property values
				getCollectorOptionsModelNew := new(posturemanagementv2.GetCollectorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetCollector(getCollectorOptionsModelNew)
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
			It(`Invoke GetCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCollectorOptions model
				getCollectorOptionsModel := new(posturemanagementv2.GetCollectorOptions)
				getCollectorOptionsModel.ID = core.StringPtr("testString")
				getCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetCollector(getCollectorOptionsModel)
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
	Describe(`UpdateCollector(updateCollectorOptions *UpdateCollectorOptions) - Operation response error`, func() {
		updateCollectorPath := "/posture/v2/collectors/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectorPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCollector with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CollectorUpdate model
				collectorUpdateModel := new(posturemanagementv2.CollectorUpdate)
				collectorUpdateModel.DisplayName = core.StringPtr("test-0112-collector_jj")
				collectorUpdateModel.Description = core.StringPtr("This collector is used for testing.")
				collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCollectorOptions model
				updateCollectorOptionsModel := new(posturemanagementv2.UpdateCollectorOptions)
				updateCollectorOptionsModel.ID = core.StringPtr("testString")
				updateCollectorOptionsModel.Collector = collectorUpdateModelAsPatch
				updateCollectorOptionsModel.AccountID = core.StringPtr("testString")
				updateCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				updateCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.UpdateCollector(updateCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.UpdateCollector(updateCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollector(updateCollectorOptions *UpdateCollectorOptions)`, func() {
		updateCollectorPath := "/posture/v2/collectors/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke UpdateCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CollectorUpdate model
				collectorUpdateModel := new(posturemanagementv2.CollectorUpdate)
				collectorUpdateModel.DisplayName = core.StringPtr("test-0112-collector_jj")
				collectorUpdateModel.Description = core.StringPtr("This collector is used for testing.")
				collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCollectorOptions model
				updateCollectorOptionsModel := new(posturemanagementv2.UpdateCollectorOptions)
				updateCollectorOptionsModel.ID = core.StringPtr("testString")
				updateCollectorOptionsModel.Collector = collectorUpdateModelAsPatch
				updateCollectorOptionsModel.AccountID = core.StringPtr("testString")
				updateCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				updateCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.UpdateCollectorWithContext(ctx, updateCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.UpdateCollector(updateCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.UpdateCollectorWithContext(ctx, updateCollectorOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}`)
				}))
			})
			It(`Invoke UpdateCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.UpdateCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CollectorUpdate model
				collectorUpdateModel := new(posturemanagementv2.CollectorUpdate)
				collectorUpdateModel.DisplayName = core.StringPtr("test-0112-collector_jj")
				collectorUpdateModel.Description = core.StringPtr("This collector is used for testing.")
				collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCollectorOptions model
				updateCollectorOptionsModel := new(posturemanagementv2.UpdateCollectorOptions)
				updateCollectorOptionsModel.ID = core.StringPtr("testString")
				updateCollectorOptionsModel.Collector = collectorUpdateModelAsPatch
				updateCollectorOptionsModel.AccountID = core.StringPtr("testString")
				updateCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				updateCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.UpdateCollector(updateCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CollectorUpdate model
				collectorUpdateModel := new(posturemanagementv2.CollectorUpdate)
				collectorUpdateModel.DisplayName = core.StringPtr("test-0112-collector_jj")
				collectorUpdateModel.Description = core.StringPtr("This collector is used for testing.")
				collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCollectorOptions model
				updateCollectorOptionsModel := new(posturemanagementv2.UpdateCollectorOptions)
				updateCollectorOptionsModel.ID = core.StringPtr("testString")
				updateCollectorOptionsModel.Collector = collectorUpdateModelAsPatch
				updateCollectorOptionsModel.AccountID = core.StringPtr("testString")
				updateCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				updateCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.UpdateCollector(updateCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCollectorOptions model with no property values
				updateCollectorOptionsModelNew := new(posturemanagementv2.UpdateCollectorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.UpdateCollector(updateCollectorOptionsModelNew)
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
			It(`Invoke UpdateCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CollectorUpdate model
				collectorUpdateModel := new(posturemanagementv2.CollectorUpdate)
				collectorUpdateModel.DisplayName = core.StringPtr("test-0112-collector_jj")
				collectorUpdateModel.Description = core.StringPtr("This collector is used for testing.")
				collectorUpdateModelAsPatch, asPatchErr := collectorUpdateModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCollectorOptions model
				updateCollectorOptionsModel := new(posturemanagementv2.UpdateCollectorOptions)
				updateCollectorOptionsModel.ID = core.StringPtr("testString")
				updateCollectorOptionsModel.Collector = collectorUpdateModelAsPatch
				updateCollectorOptionsModel.AccountID = core.StringPtr("testString")
				updateCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				updateCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.UpdateCollector(updateCollectorOptionsModel)
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
	Describe(`DeleteCollector(deleteCollectorOptions *DeleteCollectorOptions)`, func() {
		deleteCollectorPath := "/posture/v2/collectors/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectorPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := postureManagementService.DeleteCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCollectorOptions model
				deleteCollectorOptionsModel := new(posturemanagementv2.DeleteCollectorOptions)
				deleteCollectorOptionsModel.ID = core.StringPtr("testString")
				deleteCollectorOptionsModel.AccountID = core.StringPtr("testString")
				deleteCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = postureManagementService.DeleteCollector(deleteCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectorOptions model
				deleteCollectorOptionsModel := new(posturemanagementv2.DeleteCollectorOptions)
				deleteCollectorOptionsModel.ID = core.StringPtr("testString")
				deleteCollectorOptionsModel.AccountID = core.StringPtr("testString")
				deleteCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				deleteCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := postureManagementService.DeleteCollector(deleteCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCollectorOptions model with no property values
				deleteCollectorOptionsModelNew := new(posturemanagementv2.DeleteCollectorOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = postureManagementService.DeleteCollector(deleteCollectorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportProfiles(importProfilesOptions *ImportProfilesOptions) - Operation response error`, func() {
		importProfilesPath := "/posture/v2/profiles/import"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importProfilesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportProfiles with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ImportProfilesOptions model
				importProfilesOptionsModel := new(posturemanagementv2.ImportProfilesOptions)
				importProfilesOptionsModel.File = CreateMockReader("This is a mock file.")
				importProfilesOptionsModel.AccountID = core.StringPtr("testString")
				importProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				importProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ImportProfiles(importProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ImportProfiles(importProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportProfiles(importProfilesOptions *ImportProfilesOptions)`, func() {
		importProfilesPath := "/posture/v2/profiles/import"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importProfilesPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The API processed successfully.", "result": true}`)
				}))
			})
			It(`Invoke ImportProfiles successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ImportProfilesOptions model
				importProfilesOptionsModel := new(posturemanagementv2.ImportProfilesOptions)
				importProfilesOptionsModel.File = CreateMockReader("This is a mock file.")
				importProfilesOptionsModel.AccountID = core.StringPtr("testString")
				importProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				importProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ImportProfilesWithContext(ctx, importProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ImportProfiles(importProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ImportProfilesWithContext(ctx, importProfilesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(importProfilesPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The API processed successfully.", "result": true}`)
				}))
			})
			It(`Invoke ImportProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ImportProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportProfilesOptions model
				importProfilesOptionsModel := new(posturemanagementv2.ImportProfilesOptions)
				importProfilesOptionsModel.File = CreateMockReader("This is a mock file.")
				importProfilesOptionsModel.AccountID = core.StringPtr("testString")
				importProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				importProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ImportProfiles(importProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ImportProfiles with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ImportProfilesOptions model
				importProfilesOptionsModel := new(posturemanagementv2.ImportProfilesOptions)
				importProfilesOptionsModel.File = CreateMockReader("This is a mock file.")
				importProfilesOptionsModel.AccountID = core.StringPtr("testString")
				importProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				importProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ImportProfiles(importProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportProfilesOptions model with no property values
				importProfilesOptionsModelNew := new(posturemanagementv2.ImportProfilesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ImportProfiles(importProfilesOptionsModelNew)
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
			It(`Invoke ImportProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ImportProfilesOptions model
				importProfilesOptionsModel := new(posturemanagementv2.ImportProfilesOptions)
				importProfilesOptionsModel.File = CreateMockReader("This is a mock file.")
				importProfilesOptionsModel.AccountID = core.StringPtr("testString")
				importProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				importProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ImportProfiles(importProfilesOptionsModel)
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
		listProfilesPath := "/posture/v2/profiles"
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
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfiles with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv2.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(100))
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
		listProfilesPath := "/posture/v2/profiles"
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
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv2.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(100))
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
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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
				listProfilesOptionsModel := new(posturemanagementv2.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListProfiles(listProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfiles with error: Operation request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv2.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListProfiles(listProfilesOptionsModel)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := new(posturemanagementv2.ListProfilesOptions)
				listProfilesOptionsModel.AccountID = core.StringPtr("testString")
				listProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				listProfilesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listProfilesOptionsModel.Limit = core.Int64Ptr(int64(100))
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
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.ProfileList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.ProfileList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ProfileList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ProfileList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`GetProfile(getProfileOptions *GetProfileOptions) - Operation response error`, func() {
		getProfilePath := "/posture/v2/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfile with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(posturemanagementv2.GetProfileOptions)
				getProfileOptionsModel.ID = core.StringPtr("testString")
				getProfileOptionsModel.ProfileType = core.StringPtr("testString")
				getProfileOptionsModel.AccountID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetProfile(getProfileOptionsModel)
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
		getProfilePath := "/posture/v2/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}`)
				}))
			})
			It(`Invoke GetProfile successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(posturemanagementv2.GetProfileOptions)
				getProfileOptionsModel.ID = core.StringPtr("testString")
				getProfileOptionsModel.ProfileType = core.StringPtr("testString")
				getProfileOptionsModel.AccountID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetProfileWithContext(ctx, getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetProfileWithContext(ctx, getProfileOptionsModel)
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
					Expect(req.URL.Query()["profile_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(posturemanagementv2.GetProfileOptions)
				getProfileOptionsModel.ID = core.StringPtr("testString")
				getProfileOptionsModel.ProfileType = core.StringPtr("testString")
				getProfileOptionsModel.AccountID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(posturemanagementv2.GetProfileOptions)
				getProfileOptionsModel.ID = core.StringPtr("testString")
				getProfileOptionsModel.ProfileType = core.StringPtr("testString")
				getProfileOptionsModel.AccountID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetProfile(getProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfileOptionsModelNew := new(posturemanagementv2.GetProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetProfile(getProfileOptionsModelNew)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfileOptionsModel := new(posturemanagementv2.GetProfileOptions)
				getProfileOptionsModel.ID = core.StringPtr("testString")
				getProfileOptionsModel.ProfileType = core.StringPtr("testString")
				getProfileOptionsModel.AccountID = core.StringPtr("testString")
				getProfileOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetProfile(getProfileOptionsModel)
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
	Describe(`UpdateProfiles(updateProfilesOptions *UpdateProfilesOptions) - Operation response error`, func() {
		updateProfilesPath := "/posture/v2/profiles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilesPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProfiles with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateProfilesOptions model
				updateProfilesOptionsModel := new(posturemanagementv2.UpdateProfilesOptions)
				updateProfilesOptionsModel.ID = core.StringPtr("testString")
				updateProfilesOptionsModel.Name = core.StringPtr("AT_Controls_Testing")
				updateProfilesOptionsModel.Description = core.StringPtr("AT Controls")
				updateProfilesOptionsModel.BaseProfile = core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.Type = core.StringPtr("predefined")
				updateProfilesOptionsModel.IsEnabled = core.BoolPtr(true)
				updateProfilesOptionsModel.ControlIds = []string{"9980", "9979", "9994"}
				updateProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				updateProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProfiles(updateProfilesOptions *UpdateProfilesOptions)`, func() {
		updateProfilesPath := "/posture/v2/profiles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilesPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}`)
				}))
			})
			It(`Invoke UpdateProfiles successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProfilesOptions model
				updateProfilesOptionsModel := new(posturemanagementv2.UpdateProfilesOptions)
				updateProfilesOptionsModel.ID = core.StringPtr("testString")
				updateProfilesOptionsModel.Name = core.StringPtr("AT_Controls_Testing")
				updateProfilesOptionsModel.Description = core.StringPtr("AT Controls")
				updateProfilesOptionsModel.BaseProfile = core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.Type = core.StringPtr("predefined")
				updateProfilesOptionsModel.IsEnabled = core.BoolPtr(true)
				updateProfilesOptionsModel.ControlIds = []string{"9980", "9979", "9994"}
				updateProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				updateProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.UpdateProfilesWithContext(ctx, updateProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.UpdateProfilesWithContext(ctx, updateProfilesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProfilesPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "CIS IBM Foundations Benchmark 1.0.0", "description": "This profile contains controls for the CIS IBM Foundations Benchmark 1.0.0.", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "id": "3045", "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "type": "predefined", "no_of_controls": 1, "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "enabled": true}`)
				}))
			})
			It(`Invoke UpdateProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.UpdateProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProfilesOptions model
				updateProfilesOptionsModel := new(posturemanagementv2.UpdateProfilesOptions)
				updateProfilesOptionsModel.ID = core.StringPtr("testString")
				updateProfilesOptionsModel.Name = core.StringPtr("AT_Controls_Testing")
				updateProfilesOptionsModel.Description = core.StringPtr("AT Controls")
				updateProfilesOptionsModel.BaseProfile = core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.Type = core.StringPtr("predefined")
				updateProfilesOptionsModel.IsEnabled = core.BoolPtr(true)
				updateProfilesOptionsModel.ControlIds = []string{"9980", "9979", "9994"}
				updateProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				updateProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProfiles with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateProfilesOptions model
				updateProfilesOptionsModel := new(posturemanagementv2.UpdateProfilesOptions)
				updateProfilesOptionsModel.ID = core.StringPtr("testString")
				updateProfilesOptionsModel.Name = core.StringPtr("AT_Controls_Testing")
				updateProfilesOptionsModel.Description = core.StringPtr("AT Controls")
				updateProfilesOptionsModel.BaseProfile = core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.Type = core.StringPtr("predefined")
				updateProfilesOptionsModel.IsEnabled = core.BoolPtr(true)
				updateProfilesOptionsModel.ControlIds = []string{"9980", "9979", "9994"}
				updateProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				updateProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProfilesOptions model with no property values
				updateProfilesOptionsModelNew := new(posturemanagementv2.UpdateProfilesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.UpdateProfiles(updateProfilesOptionsModelNew)
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
			It(`Invoke UpdateProfiles successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateProfilesOptions model
				updateProfilesOptionsModel := new(posturemanagementv2.UpdateProfilesOptions)
				updateProfilesOptionsModel.ID = core.StringPtr("testString")
				updateProfilesOptionsModel.Name = core.StringPtr("AT_Controls_Testing")
				updateProfilesOptionsModel.Description = core.StringPtr("AT Controls")
				updateProfilesOptionsModel.BaseProfile = core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.Type = core.StringPtr("predefined")
				updateProfilesOptionsModel.IsEnabled = core.BoolPtr(true)
				updateProfilesOptionsModel.ControlIds = []string{"9980", "9979", "9994"}
				updateProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateProfilesOptionsModel.TransactionID = core.StringPtr("testString")
				updateProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.UpdateProfiles(updateProfilesOptionsModel)
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
	Describe(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {
		deleteProfilePath := "/posture/v2/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProfilePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProfile successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := postureManagementService.DeleteProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(posturemanagementv2.DeleteProfileOptions)
				deleteProfileOptionsModel.ID = core.StringPtr("testString")
				deleteProfileOptionsModel.AccountID = core.StringPtr("testString")
				deleteProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = postureManagementService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProfile with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(posturemanagementv2.DeleteProfileOptions)
				deleteProfileOptionsModel.ID = core.StringPtr("testString")
				deleteProfileOptionsModel.AccountID = core.StringPtr("testString")
				deleteProfileOptionsModel.TransactionID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := postureManagementService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProfileOptions model with no property values
				deleteProfileOptionsModelNew := new(posturemanagementv2.DeleteProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = postureManagementService.DeleteProfile(deleteProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileControls(getProfileControlsOptions *GetProfileControlsOptions) - Operation response error`, func() {
		getProfileControlsPath := "/posture/v2/profiles/testString/controls"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfileControls with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileControlsOptions model
				getProfileControlsOptionsModel := new(posturemanagementv2.GetProfileControlsOptions)
				getProfileControlsOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfileControls(getProfileControlsOptions *GetProfileControlsOptions)`, func() {
		getProfileControlsPath := "/posture/v2/profiles/testString/controls"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "controls": [{"id": "9979", "description": "Identity and Access Management", "external_control_id": "1.2", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "severity": "Medium", "is_manual": false, "is_remediable": false, "is_reversible": false, "is_automatable": false, "is_auto_remediable": false}]}]}`)
				}))
			})
			It(`Invoke GetProfileControls successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetProfileControlsOptions model
				getProfileControlsOptionsModel := new(posturemanagementv2.GetProfileControlsOptions)
				getProfileControlsOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetProfileControlsWithContext(ctx, getProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetProfileControlsWithContext(ctx, getProfileControlsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "controls": [{"id": "9979", "description": "Identity and Access Management", "external_control_id": "1.2", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "severity": "Medium", "is_manual": false, "is_remediable": false, "is_reversible": false, "is_automatable": false, "is_auto_remediable": false}]}]}`)
				}))
			})
			It(`Invoke GetProfileControls successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetProfileControls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileControlsOptions model
				getProfileControlsOptionsModel := new(posturemanagementv2.GetProfileControlsOptions)
				getProfileControlsOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProfileControls with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileControlsOptions model
				getProfileControlsOptionsModel := new(posturemanagementv2.GetProfileControlsOptions)
				getProfileControlsOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileControlsOptions model with no property values
				getProfileControlsOptionsModelNew := new(posturemanagementv2.GetProfileControlsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetProfileControls(getProfileControlsOptionsModelNew)
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
			It(`Invoke GetProfileControls successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetProfileControlsOptions model
				getProfileControlsOptionsModel := new(posturemanagementv2.GetProfileControlsOptions)
				getProfileControlsOptionsModel.ProfileID = core.StringPtr("testString")
				getProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetProfileControls(getProfileControlsOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.ControlList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`GetGroupProfileControls(getGroupProfileControlsOptions *GetGroupProfileControlsOptions) - Operation response error`, func() {
		getGroupProfileControlsPath := "/posture/v2/profiles/groups/testString/controls"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGroupProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGroupProfileControls with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetGroupProfileControlsOptions model
				getGroupProfileControlsOptionsModel := new(posturemanagementv2.GetGroupProfileControlsOptions)
				getGroupProfileControlsOptionsModel.GroupID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getGroupProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getGroupProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGroupProfileControls(getGroupProfileControlsOptions *GetGroupProfileControlsOptions)`, func() {
		getGroupProfileControlsPath := "/posture/v2/profiles/groups/testString/controls"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGroupProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "controls": [{"id": "9979", "description": "Identity and Access Management", "external_control_id": "1.2", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "severity": "Medium", "is_manual": false, "is_remediable": false, "is_reversible": false, "is_automatable": false, "is_auto_remediable": false}]}]}`)
				}))
			})
			It(`Invoke GetGroupProfileControls successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetGroupProfileControlsOptions model
				getGroupProfileControlsOptionsModel := new(posturemanagementv2.GetGroupProfileControlsOptions)
				getGroupProfileControlsOptionsModel.GroupID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getGroupProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getGroupProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetGroupProfileControlsWithContext(ctx, getGroupProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetGroupProfileControlsWithContext(ctx, getGroupProfileControlsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getGroupProfileControlsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "controls": [{"id": "9979", "description": "Identity and Access Management", "external_control_id": "1.2", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "severity": "Medium", "is_manual": false, "is_remediable": false, "is_reversible": false, "is_automatable": false, "is_auto_remediable": false}]}]}`)
				}))
			})
			It(`Invoke GetGroupProfileControls successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetGroupProfileControls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGroupProfileControlsOptions model
				getGroupProfileControlsOptionsModel := new(posturemanagementv2.GetGroupProfileControlsOptions)
				getGroupProfileControlsOptionsModel.GroupID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getGroupProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getGroupProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGroupProfileControls with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetGroupProfileControlsOptions model
				getGroupProfileControlsOptionsModel := new(posturemanagementv2.GetGroupProfileControlsOptions)
				getGroupProfileControlsOptionsModel.GroupID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getGroupProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getGroupProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGroupProfileControlsOptions model with no property values
				getGroupProfileControlsOptionsModelNew := new(posturemanagementv2.GetGroupProfileControlsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModelNew)
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
			It(`Invoke GetGroupProfileControls successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetGroupProfileControlsOptions model
				getGroupProfileControlsOptionsModel := new(posturemanagementv2.GetGroupProfileControlsOptions)
				getGroupProfileControlsOptionsModel.GroupID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.AccountID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.TransactionID = core.StringPtr("testString")
				getGroupProfileControlsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getGroupProfileControlsOptionsModel.Limit = core.Int64Ptr(int64(100))
				getGroupProfileControlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetGroupProfileControls(getGroupProfileControlsOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.ControlList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ControlList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`CreateScope(createScopeOptions *CreateScopeOptions) - Operation response error`, func() {
		createScopePath := "/posture/v2/scopes"
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv2.CreateScopeOptions)
				createScopeOptionsModel.Name = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.Description = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.CredentialType = core.StringPtr("on_premise")
				createScopeOptionsModel.Interval = core.Int64Ptr(int64(10))
				createScopeOptionsModel.IsDiscoveryScheduled = core.BoolPtr(true)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
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
		createScopePath := "/posture/v2/scopes"
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke CreateScope successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv2.CreateScopeOptions)
				createScopeOptionsModel.Name = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.Description = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.CredentialType = core.StringPtr("on_premise")
				createScopeOptionsModel.Interval = core.Int64Ptr(int64(10))
				createScopeOptionsModel.IsDiscoveryScheduled = core.BoolPtr(true)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke CreateScope successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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
				createScopeOptionsModel := new(posturemanagementv2.CreateScopeOptions)
				createScopeOptionsModel.Name = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.Description = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.CredentialType = core.StringPtr("on_premise")
				createScopeOptionsModel.Interval = core.Int64Ptr(int64(10))
				createScopeOptionsModel.IsDiscoveryScheduled = core.BoolPtr(true)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
				createScopeOptionsModel.TransactionID = core.StringPtr("testString")
				createScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateScope(createScopeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateScope with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv2.CreateScopeOptions)
				createScopeOptionsModel.Name = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.Description = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.CredentialType = core.StringPtr("on_premise")
				createScopeOptionsModel.Interval = core.Int64Ptr(int64(10))
				createScopeOptionsModel.IsDiscoveryScheduled = core.BoolPtr(true)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
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
				createScopeOptionsModelNew := new(posturemanagementv2.CreateScopeOptions)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsModel := new(posturemanagementv2.CreateScopeOptions)
				createScopeOptionsModel.Name = core.StringPtr("IBMSchema-new-048-test")
				createScopeOptionsModel.Description = core.StringPtr("IBMSchema")
				createScopeOptionsModel.CollectorIds = []string{"20"}
				createScopeOptionsModel.CredentialID = core.StringPtr("5")
				createScopeOptionsModel.CredentialType = core.StringPtr("on_premise")
				createScopeOptionsModel.Interval = core.Int64Ptr(int64(10))
				createScopeOptionsModel.IsDiscoveryScheduled = core.BoolPtr(true)
				createScopeOptionsModel.AccountID = core.StringPtr("testString")
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
		listScopesPath := "/posture/v2/scopes"
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListScopes with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv2.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
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
		listScopesPath := "/posture/v2/scopes"
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "id": "1", "uuid": "1", "name": "My_Example_Scope", "enabled": true, "credential_type": "ibm", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv2.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "id": "1", "uuid": "1", "name": "My_Example_Scope", "enabled": true, "credential_type": "ibm", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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
				listScopesOptionsModel := new(posturemanagementv2.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListScopes with error: Operation request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv2.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListScopes(listScopesOptionsModel)
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
			It(`Invoke ListScopes successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(posturemanagementv2.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.TransactionID = core.StringPtr("testString")
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
	Describe(`GetScopeDetails(getScopeDetailsOptions *GetScopeDetailsOptions) - Operation response error`, func() {
		getScopeDetailsPath := "/posture/v2/scopes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetScopeDetails with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsOptions model
				getScopeDetailsOptionsModel := new(posturemanagementv2.GetScopeDetailsOptions)
				getScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScopeDetails(getScopeDetailsOptions *GetScopeDetailsOptions)`, func() {
		getScopeDetailsPath := "/posture/v2/scopes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke GetScopeDetails successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetScopeDetailsOptions model
				getScopeDetailsOptionsModel := new(posturemanagementv2.GetScopeDetailsOptions)
				getScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetScopeDetailsWithContext(ctx, getScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetScopeDetailsWithContext(ctx, getScopeDetailsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke GetScopeDetails successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetScopeDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetScopeDetailsOptions model
				getScopeDetailsOptionsModel := new(posturemanagementv2.GetScopeDetailsOptions)
				getScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetScopeDetails with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsOptions model
				getScopeDetailsOptionsModel := new(posturemanagementv2.GetScopeDetailsOptions)
				getScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetScopeDetailsOptions model with no property values
				getScopeDetailsOptionsModelNew := new(posturemanagementv2.GetScopeDetailsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetScopeDetails(getScopeDetailsOptionsModelNew)
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
			It(`Invoke GetScopeDetails successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsOptions model
				getScopeDetailsOptionsModel := new(posturemanagementv2.GetScopeDetailsOptions)
				getScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetScopeDetails(getScopeDetailsOptionsModel)
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
	Describe(`UpdateScopeDetails(updateScopeDetailsOptions *UpdateScopeDetailsOptions) - Operation response error`, func() {
		updateScopeDetailsPath := "/posture/v2/scopes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateScopeDetailsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateScopeDetails with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateScopeDetailsOptions model
				updateScopeDetailsOptionsModel := new(posturemanagementv2.UpdateScopeDetailsOptions)
				updateScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Name = core.StringPtr("Scope Test1")
				updateScopeDetailsOptionsModel.Description = core.StringPtr("Scope Description")
				updateScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateScopeDetails(updateScopeDetailsOptions *UpdateScopeDetailsOptions)`, func() {
		updateScopeDetailsPath := "/posture/v2/scopes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateScopeDetailsPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke UpdateScopeDetails successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateScopeDetailsOptions model
				updateScopeDetailsOptionsModel := new(posturemanagementv2.UpdateScopeDetailsOptions)
				updateScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Name = core.StringPtr("Scope Test1")
				updateScopeDetailsOptionsModel.Description = core.StringPtr("Scope Description")
				updateScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.UpdateScopeDetailsWithContext(ctx, updateScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.UpdateScopeDetailsWithContext(ctx, updateScopeDetailsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateScopeDetailsPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "uuid": "UUID", "partner_uuid": "PartnerUUID", "description": "Description", "org_id": 5, "cloud_type_id": 11, "tld_credential_id": 15, "status": "pending", "status_msg": "StatusMsg", "subset_selected": true, "enabled": false, "last_discover_start_time": "LastDiscoverStartTime", "last_discover_completed_time": "LastDiscoverCompletedTime", "last_successful_discover_start_time": "LastSuccessfulDiscoverStartTime", "last_successful_discover_completed_time": "LastSuccessfulDiscoverCompletedTime", "task_type": "nop", "tasks": [{"task_logs": [{}], "task_id": 6, "task_gateway_id": 13, "task_gateway_name": "TaskGatewayName", "task_task_type": "nop", "task_gateway_schema_id": 19, "task_schema_name": "TaskSchemaName", "task_discover_id": 14, "task_status": "pending", "task_status_msg": "TaskStatusMsg", "task_start_time": 13, "task_updated_time": 15, "task_derived_status": "pending", "task_created_by": "TaskCreatedBy"}], "status_updated_time": "StatusUpdatedTime", "collectors_by_type": {"mapKey": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}]}, "credentials_by_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "credentials_by_sub_categeory_type": {"mapKey": [{"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}]}, "sub_categories_by_type": {"mapKey": ["ms_365"]}, "resource_groups": "ResourceGroups", "region_names": "RegionNames", "cloud_type": "CloudType", "env_sub_category": "EnvSubCategory", "tld_credentail": {"id": "ID", "name": "Name", "uuid": "UUID", "type": "Type", "data": {"anyKey": "anyValue"}, "display_fields": {"ibm_api_key": "sample_api_key", "aws_client_id": "sample_client_id", "aws_client_secret": "*********", "aws_region": "test_region", "aws_arn": "sample_arn", "username": "sample_username", "password": "************", "azure_client_id": "azure_124", "azure_client_secret": "************", "azure_subscription_id": "A32432890", "azure_resource_group": "azure_res_type", "database_name": "sample_db_name", "winrm_authtype": "sample_auth_type", "winrm_usessl": "ssl_test", "winrm_port": "80", "ms_365_client_id": "ms_client_id", "ms_365_client_secret": "ms_client_secret", "ms_365_tenant_id": "ms_tenant_id", "auth_url": "test.example.com", "project_name": "test_proj", "user_domain_name": "user_domain", "project_domain_name": "stack_domain_name"}, "version_timestamp": {"anyKey": "anyValue"}, "description": "Description", "is_enabled": false, "gateway_key": "GatewayKey", "credential_group": {"anyKey": "anyValue"}, "enabled_credential_group": true, "groups": [{"id": "ID", "passphrase": "Passphrase"}], "purpose": "Purpose"}, "collectors": [{"id": "ID", "display_name": "DisplayName", "name": "Name", "public_key": "PublicKey", "last_heartbeat": "2019-01-01T12:00:00.000Z", "status": "ready_to_install", "collector_version": "1.0.0_06141eef8d03443c4c6544bac3ea192c54c83e70", "image_version": "1.0.0_ef6c07f1c622bc5472d29b9b82e9885d23d12e8b", "description": "Description", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "enabled": false, "registration_code": "RegistrationCode", "type": "restricted", "credential_public_key": "CredentialPublicKey", "failure_count": 12, "approved_local_gateway_ip": "ApprovedLocalGatewayIP", "approved_internet_gateway_ip": "ApprovedInternetGatewayIP", "last_failed_local_gateway_ip": "LastFailedLocalGatewayIP", "reset_reason": "ResetReason", "hostname": "Test-MacBook-Pro.local", "install_path": "/Users/test/project/containers/collector1", "use_private_endpoint": true, "managed_by": "ibm", "trial_expiry": "2019-01-01T12:00:00.000Z", "last_failed_internet_gateway_ip": "LastFailedInternetGatewayIP", "status_description": "StatusDescription", "reset_time": "2019-01-01T12:00:00.000Z", "is_public": true, "is_ubi_image": true}], "first_level_scoped_data": [{"scope_object": "ScopeObject", "scope_init_scope": "ScopeInitScope", "scope": "Scope", "scope_changed": true, "scope_id": "ScopeID", "scope_properties": "ScopeProperties", "scope_overlay": "ScopeOverlay", "scope_new_found": false, "scope_discovery_status": {"anyKey": "anyValue"}, "scope_fact_status": {"anyKey": "anyValue"}, "scope_facts": "ScopeFacts", "scope_list_members": {"anyKey": "anyValue"}, "scope_children": {"anyKey": "anyValue"}, "scope_resource_category": "ScopeResourceCategory", "scope_resource_type": "ScopeResourceType", "scope_resource": "ScopeResource", "scope_resource_attributes": {"anyKey": "anyValue"}, "scope_drift": "ScopeDrift", "scope_parse_status": "ScopeParseStatus", "scope_transformed_facts": {"anyKey": "anyValue"}, "scope_collector_id": 16}], "discovery_methods": ["DiscoveryMethods"], "discovery_method": "DiscoveryMethod", "file_type": "FileType", "file_format": "FileFormat", "created_by": "CreatedBy", "created_at": "CreatedAt", "modified_by": "ModifiedBy", "modified_at": "ModifiedAt", "is_discovery_scheduled": true, "interval": 8, "discovery_setting_id": 18, "include_new_eagerly": false, "type": "validation", "correlation_id": "CorrelationID", "credential_attributes": "CredentialAttributes"}`)
				}))
			})
			It(`Invoke UpdateScopeDetails successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.UpdateScopeDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateScopeDetailsOptions model
				updateScopeDetailsOptionsModel := new(posturemanagementv2.UpdateScopeDetailsOptions)
				updateScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Name = core.StringPtr("Scope Test1")
				updateScopeDetailsOptionsModel.Description = core.StringPtr("Scope Description")
				updateScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateScopeDetails with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateScopeDetailsOptions model
				updateScopeDetailsOptionsModel := new(posturemanagementv2.UpdateScopeDetailsOptions)
				updateScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Name = core.StringPtr("Scope Test1")
				updateScopeDetailsOptionsModel.Description = core.StringPtr("Scope Description")
				updateScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateScopeDetailsOptions model with no property values
				updateScopeDetailsOptionsModelNew := new(posturemanagementv2.UpdateScopeDetailsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModelNew)
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
			It(`Invoke UpdateScopeDetails successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateScopeDetailsOptions model
				updateScopeDetailsOptionsModel := new(posturemanagementv2.UpdateScopeDetailsOptions)
				updateScopeDetailsOptionsModel.ID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Name = core.StringPtr("Scope Test1")
				updateScopeDetailsOptionsModel.Description = core.StringPtr("Scope Description")
				updateScopeDetailsOptionsModel.AccountID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.TransactionID = core.StringPtr("testString")
				updateScopeDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.UpdateScopeDetails(updateScopeDetailsOptionsModel)
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
	Describe(`DeleteScope(deleteScopeOptions *DeleteScopeOptions)`, func() {
		deleteScopePath := "/posture/v2/scopes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteScopePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteScope successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := postureManagementService.DeleteScope(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteScopeOptions model
				deleteScopeOptionsModel := new(posturemanagementv2.DeleteScopeOptions)
				deleteScopeOptionsModel.ID = core.StringPtr("testString")
				deleteScopeOptionsModel.AccountID = core.StringPtr("testString")
				deleteScopeOptionsModel.TransactionID = core.StringPtr("testString")
				deleteScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = postureManagementService.DeleteScope(deleteScopeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteScope with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteScopeOptions model
				deleteScopeOptionsModel := new(posturemanagementv2.DeleteScopeOptions)
				deleteScopeOptionsModel.ID = core.StringPtr("testString")
				deleteScopeOptionsModel.AccountID = core.StringPtr("testString")
				deleteScopeOptionsModel.TransactionID = core.StringPtr("testString")
				deleteScopeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := postureManagementService.DeleteScope(deleteScopeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteScopeOptions model with no property values
				deleteScopeOptionsModelNew := new(posturemanagementv2.DeleteScopeOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = postureManagementService.DeleteScope(deleteScopeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScopeTimeline(getScopeTimelineOptions *GetScopeTimelineOptions) - Operation response error`, func() {
		getScopeTimelinePath := "/posture/v2/scopes/testString/events"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeTimelinePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetScopeTimeline with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeTimelineOptions model
				getScopeTimelineOptionsModel := new(posturemanagementv2.GetScopeTimelineOptions)
				getScopeTimelineOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.AccountID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScopeTimeline(getScopeTimelineOptions *GetScopeTimelineOptions)`, func() {
		getScopeTimelinePath := "/posture/v2/scopes/testString/events"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeTimelinePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"events": [{"id": "ID", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "task_type": "discovery", "status": "discovery_completed", "data_available": true, "status_message": "Discovery is completed, please check the report."}]}`)
				}))
			})
			It(`Invoke GetScopeTimeline successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetScopeTimelineOptions model
				getScopeTimelineOptionsModel := new(posturemanagementv2.GetScopeTimelineOptions)
				getScopeTimelineOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.AccountID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetScopeTimelineWithContext(ctx, getScopeTimelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetScopeTimelineWithContext(ctx, getScopeTimelineOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getScopeTimelinePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"events": [{"id": "ID", "created_at": "2021-02-26T04:07:25.000Z", "updated_at": "2021-02-26T04:07:25.000Z", "task_type": "discovery", "status": "discovery_completed", "data_available": true, "status_message": "Discovery is completed, please check the report."}]}`)
				}))
			})
			It(`Invoke GetScopeTimeline successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetScopeTimeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetScopeTimelineOptions model
				getScopeTimelineOptionsModel := new(posturemanagementv2.GetScopeTimelineOptions)
				getScopeTimelineOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.AccountID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetScopeTimeline with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeTimelineOptions model
				getScopeTimelineOptionsModel := new(posturemanagementv2.GetScopeTimelineOptions)
				getScopeTimelineOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.AccountID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetScopeTimelineOptions model with no property values
				getScopeTimelineOptionsModelNew := new(posturemanagementv2.GetScopeTimelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModelNew)
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
			It(`Invoke GetScopeTimeline successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeTimelineOptions model
				getScopeTimelineOptionsModel := new(posturemanagementv2.GetScopeTimelineOptions)
				getScopeTimelineOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.AccountID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeTimelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetScopeTimeline(getScopeTimelineOptionsModel)
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
	Describe(`GetScopeDetailsCredentials(getScopeDetailsCredentialsOptions *GetScopeDetailsCredentialsOptions) - Operation response error`, func() {
		getScopeDetailsCredentialsPath := "/posture/v2/scopes/testString/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetScopeDetailsCredentials with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				getScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				getScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScopeDetailsCredentials(getScopeDetailsCredentialsOptions *GetScopeDetailsCredentialsOptions)`, func() {
		getScopeDetailsCredentialsPath := "/posture/v2/scopes/testString/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_attribute": "CredentialAttribute", "credential_id": "CredentialID"}`)
				}))
			})
			It(`Invoke GetScopeDetailsCredentials successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				getScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				getScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetScopeDetailsCredentialsWithContext(ctx, getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetScopeDetailsCredentialsWithContext(ctx, getScopeDetailsCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_attribute": "CredentialAttribute", "credential_id": "CredentialID"}`)
				}))
			})
			It(`Invoke GetScopeDetailsCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetScopeDetailsCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				getScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				getScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetScopeDetailsCredentials with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				getScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				getScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetScopeDetailsCredentialsOptions model with no property values
				getScopeDetailsCredentialsOptionsModelNew := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModelNew)
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
			It(`Invoke GetScopeDetailsCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				getScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.GetScopeDetailsCredentialsOptions)
				getScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetScopeDetailsCredentials(getScopeDetailsCredentialsOptionsModel)
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
	Describe(`ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptions *ReplaceScopeDetailsCredentialsOptions) - Operation response error`, func() {
		replaceScopeDetailsCredentialsPath := "/posture/v2/scopes/testString/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCredentialsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCredentials with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				replaceScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				replaceScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.CredentialID = core.StringPtr("1")
				replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute = core.StringPtr("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptions *ReplaceScopeDetailsCredentialsOptions)`, func() {
		replaceScopeDetailsCredentialsPath := "/posture/v2/scopes/testString/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCredentialsPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_attribute": "CredentialAttribute", "credential_id": "CredentialID"}`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCredentials successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				replaceScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				replaceScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.CredentialID = core.StringPtr("1")
				replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute = core.StringPtr("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ReplaceScopeDetailsCredentialsWithContext(ctx, replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ReplaceScopeDetailsCredentialsWithContext(ctx, replaceScopeDetailsCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCredentialsPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_attribute": "CredentialAttribute", "credential_id": "CredentialID"}`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				replaceScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				replaceScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.CredentialID = core.StringPtr("1")
				replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute = core.StringPtr("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceScopeDetailsCredentials with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				replaceScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				replaceScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.CredentialID = core.StringPtr("1")
				replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute = core.StringPtr("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceScopeDetailsCredentialsOptions model with no property values
				replaceScopeDetailsCredentialsOptionsModelNew := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModelNew)
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
			It(`Invoke ReplaceScopeDetailsCredentials successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				replaceScopeDetailsCredentialsOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCredentialsOptions)
				replaceScopeDetailsCredentialsOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.CredentialID = core.StringPtr("1")
				replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute = core.StringPtr("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCredentials(replaceScopeDetailsCredentialsOptionsModel)
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
	Describe(`GetScopeDetailsCollector(getScopeDetailsCollectorOptions *GetScopeDetailsCollectorOptions) - Operation response error`, func() {
		getScopeDetailsCollectorPath := "/posture/v2/scopes/testString/collectors"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCollectorPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetScopeDetailsCollector with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCollectorOptions model
				getScopeDetailsCollectorOptionsModel := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				getScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScopeDetailsCollector(getScopeDetailsCollectorOptions *GetScopeDetailsCollectorOptions)`, func() {
		getScopeDetailsCollectorPath := "/posture/v2/scopes/testString/collectors"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCollectorPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collector_ids": ["CollectorIds"]}`)
				}))
			})
			It(`Invoke GetScopeDetailsCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetScopeDetailsCollectorOptions model
				getScopeDetailsCollectorOptionsModel := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				getScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetScopeDetailsCollectorWithContext(ctx, getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetScopeDetailsCollectorWithContext(ctx, getScopeDetailsCollectorOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getScopeDetailsCollectorPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collector_ids": ["CollectorIds"]}`)
				}))
			})
			It(`Invoke GetScopeDetailsCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetScopeDetailsCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetScopeDetailsCollectorOptions model
				getScopeDetailsCollectorOptionsModel := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				getScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetScopeDetailsCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCollectorOptions model
				getScopeDetailsCollectorOptionsModel := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				getScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetScopeDetailsCollectorOptions model with no property values
				getScopeDetailsCollectorOptionsModelNew := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModelNew)
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
			It(`Invoke GetScopeDetailsCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetScopeDetailsCollectorOptions model
				getScopeDetailsCollectorOptionsModel := new(posturemanagementv2.GetScopeDetailsCollectorOptions)
				getScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				getScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetScopeDetailsCollector(getScopeDetailsCollectorOptionsModel)
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
	Describe(`ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptions *ReplaceScopeDetailsCollectorOptions) - Operation response error`, func() {
		replaceScopeDetailsCollectorPath := "/posture/v2/scopes/testString/collectors"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCollectorPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCollector with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				replaceScopeDetailsCollectorOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				replaceScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.CollectorIds = []string{"7"}
				replaceScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptions *ReplaceScopeDetailsCollectorOptions)`, func() {
		replaceScopeDetailsCollectorPath := "/posture/v2/scopes/testString/collectors"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collector_ids": ["CollectorIds"]}`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCollector successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				replaceScopeDetailsCollectorOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				replaceScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.CollectorIds = []string{"7"}
				replaceScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ReplaceScopeDetailsCollectorWithContext(ctx, replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ReplaceScopeDetailsCollectorWithContext(ctx, replaceScopeDetailsCollectorOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceScopeDetailsCollectorPath))
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
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collector_ids": ["CollectorIds"]}`)
				}))
			})
			It(`Invoke ReplaceScopeDetailsCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCollector(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				replaceScopeDetailsCollectorOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				replaceScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.CollectorIds = []string{"7"}
				replaceScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceScopeDetailsCollector with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				replaceScopeDetailsCollectorOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				replaceScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.CollectorIds = []string{"7"}
				replaceScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceScopeDetailsCollectorOptions model with no property values
				replaceScopeDetailsCollectorOptionsModelNew := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModelNew)
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
			It(`Invoke ReplaceScopeDetailsCollector successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				replaceScopeDetailsCollectorOptionsModel := new(posturemanagementv2.ReplaceScopeDetailsCollectorOptions)
				replaceScopeDetailsCollectorOptionsModel.ScopeID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.CollectorIds = []string{"7"}
				replaceScopeDetailsCollectorOptionsModel.AccountID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.TransactionID = core.StringPtr("testString")
				replaceScopeDetailsCollectorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ReplaceScopeDetailsCollector(replaceScopeDetailsCollectorOptionsModel)
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
	Describe(`GetCorrelationID(getCorrelationIDOptions *GetCorrelationIDOptions) - Operation response error`, func() {
		getCorrelationIDPath := "/posture/v2/scope/status/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorrelationIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCorrelationID with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCorrelationIDOptions model
				getCorrelationIDOptionsModel := new(posturemanagementv2.GetCorrelationIDOptions)
				getCorrelationIDOptionsModel.CorrelationID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.AccountID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.TransactionID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCorrelationID(getCorrelationIDOptions *GetCorrelationIDOptions)`, func() {
		getCorrelationIDPath := "/posture/v2/scope/status/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorrelationIDPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"correlation_id": "CorrelationID", "status": "Status", "start_time": "StartTime", "last_heartbeat": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCorrelationID successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCorrelationIDOptions model
				getCorrelationIDOptionsModel := new(posturemanagementv2.GetCorrelationIDOptions)
				getCorrelationIDOptionsModel.CorrelationID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.AccountID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.TransactionID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.GetCorrelationIDWithContext(ctx, getCorrelationIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.GetCorrelationIDWithContext(ctx, getCorrelationIDOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCorrelationIDPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"correlation_id": "CorrelationID", "status": "Status", "start_time": "StartTime", "last_heartbeat": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCorrelationID successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.GetCorrelationID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCorrelationIDOptions model
				getCorrelationIDOptionsModel := new(posturemanagementv2.GetCorrelationIDOptions)
				getCorrelationIDOptionsModel.CorrelationID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.AccountID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.TransactionID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCorrelationID with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCorrelationIDOptions model
				getCorrelationIDOptionsModel := new(posturemanagementv2.GetCorrelationIDOptions)
				getCorrelationIDOptionsModel.CorrelationID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.AccountID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.TransactionID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCorrelationIDOptions model with no property values
				getCorrelationIDOptionsModelNew := new(posturemanagementv2.GetCorrelationIDOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.GetCorrelationID(getCorrelationIDOptionsModelNew)
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
			It(`Invoke GetCorrelationID successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the GetCorrelationIDOptions model
				getCorrelationIDOptionsModel := new(posturemanagementv2.GetCorrelationIDOptions)
				getCorrelationIDOptionsModel.CorrelationID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.AccountID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.TransactionID = core.StringPtr("testString")
				getCorrelationIDOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.GetCorrelationID(getCorrelationIDOptionsModel)
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
	Describe(`ListLatestScans(listLatestScansOptions *ListLatestScansOptions) - Operation response error`, func() {
		listLatestScansPath := "/posture/v2/scans/validations/latest_scans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLatestScansPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLatestScans with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := new(posturemanagementv2.ListLatestScansOptions)
				listLatestScansOptionsModel.AccountID = core.StringPtr("testString")
				listLatestScansOptionsModel.TransactionID = core.StringPtr("testString")
				listLatestScansOptionsModel.Offset = core.Int64Ptr(int64(38))
				listLatestScansOptionsModel.Limit = core.Int64Ptr(int64(100))
				listLatestScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ListLatestScans(listLatestScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ListLatestScans(listLatestScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLatestScans(listLatestScansOptions *ListLatestScansOptions)`, func() {
		listLatestScansPath := "/posture/v2/scans/validations/latest_scans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLatestScansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "latest_scans": [{"scan_id": "262", "scan_name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "id": "1", "type": "predefined"}], "group_profile_id": "1", "group_profile_name": "CIS Windows Server Benchmarks", "report_run_by": "controller", "start_time": "2020-09-23T12:45:24.000Z", "report_setting_id": "66", "end_time": "2020-09-23T12:45:24.000Z", "result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}]}`)
				}))
			})
			It(`Invoke ListLatestScans successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := new(posturemanagementv2.ListLatestScansOptions)
				listLatestScansOptionsModel.AccountID = core.StringPtr("testString")
				listLatestScansOptionsModel.TransactionID = core.StringPtr("testString")
				listLatestScansOptionsModel.Offset = core.Int64Ptr(int64(38))
				listLatestScansOptionsModel.Limit = core.Int64Ptr(int64(100))
				listLatestScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ListLatestScansWithContext(ctx, listLatestScansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ListLatestScans(listLatestScansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ListLatestScansWithContext(ctx, listLatestScansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listLatestScansPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "latest_scans": [{"scan_id": "262", "scan_name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "id": "1", "type": "predefined"}], "group_profile_id": "1", "group_profile_name": "CIS Windows Server Benchmarks", "report_run_by": "controller", "start_time": "2020-09-23T12:45:24.000Z", "report_setting_id": "66", "end_time": "2020-09-23T12:45:24.000Z", "result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}]}`)
				}))
			})
			It(`Invoke ListLatestScans successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ListLatestScans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := new(posturemanagementv2.ListLatestScansOptions)
				listLatestScansOptionsModel.AccountID = core.StringPtr("testString")
				listLatestScansOptionsModel.TransactionID = core.StringPtr("testString")
				listLatestScansOptionsModel.Offset = core.Int64Ptr(int64(38))
				listLatestScansOptionsModel.Limit = core.Int64Ptr(int64(100))
				listLatestScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ListLatestScans(listLatestScansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListLatestScans with error: Operation request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := new(posturemanagementv2.ListLatestScansOptions)
				listLatestScansOptionsModel.AccountID = core.StringPtr("testString")
				listLatestScansOptionsModel.TransactionID = core.StringPtr("testString")
				listLatestScansOptionsModel.Offset = core.Int64Ptr(int64(38))
				listLatestScansOptionsModel.Limit = core.Int64Ptr(int64(100))
				listLatestScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ListLatestScans(listLatestScansOptionsModel)
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
			It(`Invoke ListLatestScans successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := new(posturemanagementv2.ListLatestScansOptions)
				listLatestScansOptionsModel.AccountID = core.StringPtr("testString")
				listLatestScansOptionsModel.TransactionID = core.StringPtr("testString")
				listLatestScansOptionsModel.Offset = core.Int64Ptr(int64(38))
				listLatestScansOptionsModel.Limit = core.Int64Ptr(int64(100))
				listLatestScansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ListLatestScans(listLatestScansOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.ScanList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.ScanList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ScanList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.ScanList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`CreateValidation(createValidationOptions *CreateValidationOptions) - Operation response error`, func() {
		createValidationPath := "/posture/v2/scans/validations"
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv2.CreateValidationOptions)
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
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
		createValidationPath := "/posture/v2/scans/validations"
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv2.CreateValidationOptions)
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
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
				createValidationOptionsModel := new(posturemanagementv2.CreateValidationOptions)
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
				createValidationOptionsModel.TransactionID = core.StringPtr("testString")
				createValidationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.CreateValidation(createValidationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateValidation with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv2.CreateValidationOptions)
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
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
				createValidationOptionsModelNew := new(posturemanagementv2.CreateValidationOptions)
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
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsModel := new(posturemanagementv2.CreateValidationOptions)
				createValidationOptionsModel.ScopeID = core.StringPtr("1")
				createValidationOptionsModel.ProfileID = core.StringPtr("6")
				createValidationOptionsModel.GroupProfileID = core.StringPtr("13")
				createValidationOptionsModel.AccountID = core.StringPtr("testString")
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
	Describe(`ScansSummary(scansSummaryOptions *ScansSummaryOptions) - Operation response error`, func() {
		scansSummaryPath := "/posture/v2/scans/validations/testString/summary"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scansSummaryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ScansSummary with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScansSummaryOptions model
				scansSummaryOptionsModel := new(posturemanagementv2.ScansSummaryOptions)
				scansSummaryOptionsModel.ScanID = core.StringPtr("testString")
				scansSummaryOptionsModel.ProfileID = core.StringPtr("testString")
				scansSummaryOptionsModel.AccountID = core.StringPtr("testString")
				scansSummaryOptionsModel.TransactionID = core.StringPtr("testString")
				scansSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ScansSummary(scansSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ScansSummary(scansSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ScansSummary(scansSummaryOptions *ScansSummaryOptions)`, func() {
		scansSummaryPath := "/posture/v2/scans/validations/testString/summary"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scansSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "4", "discover_id": "8", "profile_id": "50", "profile_name": "CIS IBM Foundations Benchmark 1.0.0", "scope_id": "2", "controls": [{"id": "7907", "status": "pass", "external_control_id": "50", "desciption": "CIS IBM Foundations Benchmark 1.0.0", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "status": "pass", "severity": "Medium", "completed_time": "2021-04-16T05:32:44.000Z", "error": "N/A", "resource_result": [{"name": "PasswordPolicy", "types": "Identity and Access Management", "status": "pass", "display_expected_value": "IBMid password policy should be required at least one lowercase letter", "actual_value": "{\"2fa\": \"Enabled\"}", "results_info": "IBMid password policy minimum password length is set to 8", "not_applicable_reason": "{\"isBootVolume\": \"True\"}"}], "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}}], "resource_statistics": {"pass_count": 4, "fail_count": 4, "unable_to_perform_count": 4, "not_applicable_count": 4}}]}`)
				}))
			})
			It(`Invoke ScansSummary successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ScansSummaryOptions model
				scansSummaryOptionsModel := new(posturemanagementv2.ScansSummaryOptions)
				scansSummaryOptionsModel.ScanID = core.StringPtr("testString")
				scansSummaryOptionsModel.ProfileID = core.StringPtr("testString")
				scansSummaryOptionsModel.AccountID = core.StringPtr("testString")
				scansSummaryOptionsModel.TransactionID = core.StringPtr("testString")
				scansSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ScansSummaryWithContext(ctx, scansSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ScansSummary(scansSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ScansSummaryWithContext(ctx, scansSummaryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(scansSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["profile_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "4", "discover_id": "8", "profile_id": "50", "profile_name": "CIS IBM Foundations Benchmark 1.0.0", "scope_id": "2", "controls": [{"id": "7907", "status": "pass", "external_control_id": "50", "desciption": "CIS IBM Foundations Benchmark 1.0.0", "goals": [{"description": "Check whether API keys unused for 180 days are detected and optionally disabled", "id": "3000039", "status": "pass", "severity": "Medium", "completed_time": "2021-04-16T05:32:44.000Z", "error": "N/A", "resource_result": [{"name": "PasswordPolicy", "types": "Identity and Access Management", "status": "pass", "display_expected_value": "IBMid password policy should be required at least one lowercase letter", "actual_value": "{\"2fa\": \"Enabled\"}", "results_info": "IBMid password policy minimum password length is set to 8", "not_applicable_reason": "{\"isBootVolume\": \"True\"}"}], "applicability_criteria": {"environment": ["ibm"], "resource": ["cloud_object_storage"], "environment_category": ["cloud_platform"], "resource_category": ["xaas"], "resource_type": ["storage"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}}], "resource_statistics": {"pass_count": 4, "fail_count": 4, "unable_to_perform_count": 4, "not_applicable_count": 4}}]}`)
				}))
			})
			It(`Invoke ScansSummary successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ScansSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ScansSummaryOptions model
				scansSummaryOptionsModel := new(posturemanagementv2.ScansSummaryOptions)
				scansSummaryOptionsModel.ScanID = core.StringPtr("testString")
				scansSummaryOptionsModel.ProfileID = core.StringPtr("testString")
				scansSummaryOptionsModel.AccountID = core.StringPtr("testString")
				scansSummaryOptionsModel.TransactionID = core.StringPtr("testString")
				scansSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ScansSummary(scansSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ScansSummary with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScansSummaryOptions model
				scansSummaryOptionsModel := new(posturemanagementv2.ScansSummaryOptions)
				scansSummaryOptionsModel.ScanID = core.StringPtr("testString")
				scansSummaryOptionsModel.ProfileID = core.StringPtr("testString")
				scansSummaryOptionsModel.AccountID = core.StringPtr("testString")
				scansSummaryOptionsModel.TransactionID = core.StringPtr("testString")
				scansSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ScansSummary(scansSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ScansSummaryOptions model with no property values
				scansSummaryOptionsModelNew := new(posturemanagementv2.ScansSummaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ScansSummary(scansSummaryOptionsModelNew)
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
			It(`Invoke ScansSummary successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScansSummaryOptions model
				scansSummaryOptionsModel := new(posturemanagementv2.ScansSummaryOptions)
				scansSummaryOptionsModel.ScanID = core.StringPtr("testString")
				scansSummaryOptionsModel.ProfileID = core.StringPtr("testString")
				scansSummaryOptionsModel.AccountID = core.StringPtr("testString")
				scansSummaryOptionsModel.TransactionID = core.StringPtr("testString")
				scansSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ScansSummary(scansSummaryOptionsModel)
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
	Describe(`ScanSummaries(scanSummariesOptions *ScanSummariesOptions) - Operation response error`, func() {
		scanSummariesPath := "/posture/v2/scans/validations/summaries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scanSummariesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_setting_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ScanSummaries with error: Operation response processing error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScanSummariesOptions model
				scanSummariesOptionsModel := new(posturemanagementv2.ScanSummariesOptions)
				scanSummariesOptionsModel.ReportSettingID = core.StringPtr("testString")
				scanSummariesOptionsModel.AccountID = core.StringPtr("testString")
				scanSummariesOptionsModel.TransactionID = core.StringPtr("testString")
				scanSummariesOptionsModel.Offset = core.Int64Ptr(int64(38))
				scanSummariesOptionsModel.Limit = core.Int64Ptr(int64(100))
				scanSummariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := postureManagementService.ScanSummaries(scanSummariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				postureManagementService.EnableRetries(0, 0)
				result, response, operationErr = postureManagementService.ScanSummaries(scanSummariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ScanSummaries(scanSummariesOptions *ScanSummariesOptions)`, func() {
		scanSummariesPath := "/posture/v2/scans/validations/summaries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scanSummariesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_setting_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "summaries": [{"id": "262", "name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "report_run_by": "controller", "start_time": "2020-09-23T12:45:24.000Z", "end_time": "2020-09-23T12:45:24.000Z", "status": "validation_completed", "profiles": [{"id": "48", "name": "IBM Cloud Best Practices Controls 1.0", "type": "standard", "validation_result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}], "group_profiles": [{"id": "48", "name": "IBM Cloud Best Practices Controls 1.0", "type": "standard", "validation_result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}]}]}`)
				}))
			})
			It(`Invoke ScanSummaries successfully with retries`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())
				postureManagementService.EnableRetries(0, 0)

				// Construct an instance of the ScanSummariesOptions model
				scanSummariesOptionsModel := new(posturemanagementv2.ScanSummariesOptions)
				scanSummariesOptionsModel.ReportSettingID = core.StringPtr("testString")
				scanSummariesOptionsModel.AccountID = core.StringPtr("testString")
				scanSummariesOptionsModel.TransactionID = core.StringPtr("testString")
				scanSummariesOptionsModel.Offset = core.Int64Ptr(int64(38))
				scanSummariesOptionsModel.Limit = core.Int64Ptr(int64(100))
				scanSummariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := postureManagementService.ScanSummariesWithContext(ctx, scanSummariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				postureManagementService.DisableRetries()
				result, response, operationErr := postureManagementService.ScanSummaries(scanSummariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = postureManagementService.ScanSummariesWithContext(ctx, scanSummariesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(scanSummariesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["report_setting_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 10, "limit": 15, "total_count": 50, "first": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "last": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "previous": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "next": {"href": "https://us-south.cloud.ibm.com/posture/v2/profiles/groups/1/controls?account_id=d194db5f52544a8f953aa539ced9b570&offset=10&limit=15"}, "summaries": [{"id": "262", "name": "IBM_Schema_Full - IBMCloudBestPractice", "scope_id": "21", "scope_name": "IBM_Schema_Full", "report_run_by": "controller", "start_time": "2020-09-23T12:45:24.000Z", "end_time": "2020-09-23T12:45:24.000Z", "status": "validation_completed", "profiles": [{"id": "48", "name": "IBM Cloud Best Practices Controls 1.0", "type": "standard", "validation_result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}], "group_profiles": [{"id": "48", "name": "IBM Cloud Best Practices Controls 1.0", "type": "standard", "validation_result": {"goals_pass_count": 118, "goals_unable_to_perform_count": 16, "goals_not_applicable_count": 6, "goals_fail_count": 154, "goals_total_count": 294, "controls_pass_count": 117, "controls_fail_count": 154, "controls_not_applicable_count": 6, "controls_unable_to_perform_count": 16, "controls_total_count": 293}}]}]}`)
				}))
			})
			It(`Invoke ScanSummaries successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := postureManagementService.ScanSummaries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ScanSummariesOptions model
				scanSummariesOptionsModel := new(posturemanagementv2.ScanSummariesOptions)
				scanSummariesOptionsModel.ReportSettingID = core.StringPtr("testString")
				scanSummariesOptionsModel.AccountID = core.StringPtr("testString")
				scanSummariesOptionsModel.TransactionID = core.StringPtr("testString")
				scanSummariesOptionsModel.Offset = core.Int64Ptr(int64(38))
				scanSummariesOptionsModel.Limit = core.Int64Ptr(int64(100))
				scanSummariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = postureManagementService.ScanSummaries(scanSummariesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ScanSummaries with error: Operation validation and request error`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScanSummariesOptions model
				scanSummariesOptionsModel := new(posturemanagementv2.ScanSummariesOptions)
				scanSummariesOptionsModel.ReportSettingID = core.StringPtr("testString")
				scanSummariesOptionsModel.AccountID = core.StringPtr("testString")
				scanSummariesOptionsModel.TransactionID = core.StringPtr("testString")
				scanSummariesOptionsModel.Offset = core.Int64Ptr(int64(38))
				scanSummariesOptionsModel.Limit = core.Int64Ptr(int64(100))
				scanSummariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := postureManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := postureManagementService.ScanSummaries(scanSummariesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ScanSummariesOptions model with no property values
				scanSummariesOptionsModelNew := new(posturemanagementv2.ScanSummariesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = postureManagementService.ScanSummaries(scanSummariesOptionsModelNew)
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
			It(`Invoke ScanSummaries successfully`, func() {
				postureManagementService, serviceErr := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(postureManagementService).ToNot(BeNil())

				// Construct an instance of the ScanSummariesOptions model
				scanSummariesOptionsModel := new(posturemanagementv2.ScanSummariesOptions)
				scanSummariesOptionsModel.ReportSettingID = core.StringPtr("testString")
				scanSummariesOptionsModel.AccountID = core.StringPtr("testString")
				scanSummariesOptionsModel.TransactionID = core.StringPtr("testString")
				scanSummariesOptionsModel.Offset = core.Int64Ptr(int64(38))
				scanSummariesOptionsModel.Limit = core.Int64Ptr(int64(100))
				scanSummariesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := postureManagementService.ScanSummaries(scanSummariesOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(posturemanagementv2.SummaryList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(posturemanagementv2.SummaryList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.SummaryList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(posturemanagementv2.SummaryList)
				nextObject := new(posturemanagementv2.PageLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			postureManagementService, _ := posturemanagementv2.NewPostureManagementV2(&posturemanagementv2.PostureManagementV2Options{
				URL:           "http://posturemanagementv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateCollectorOptions successfully`, func() {
				// Construct an instance of the CreateCollectorOptions model
				createCollectorOptionsName := "IBM-collector-sample"
				createCollectorOptionsIsPublic := true
				createCollectorOptionsManagedBy := "ibm"
				createCollectorOptionsModel := postureManagementService.NewCreateCollectorOptions(createCollectorOptionsName, createCollectorOptionsIsPublic, createCollectorOptionsManagedBy)
				createCollectorOptionsModel.SetName("IBM-collector-sample")
				createCollectorOptionsModel.SetIsPublic(true)
				createCollectorOptionsModel.SetManagedBy("ibm")
				createCollectorOptionsModel.SetDescription("sample collector")
				createCollectorOptionsModel.SetPassphrase("secret")
				createCollectorOptionsModel.SetIsUbiImage(true)
				createCollectorOptionsModel.SetAccountID("testString")
				createCollectorOptionsModel.SetTransactionID("testString")
				createCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectorOptionsModel).ToNot(BeNil())
				Expect(createCollectorOptionsModel.Name).To(Equal(core.StringPtr("IBM-collector-sample")))
				Expect(createCollectorOptionsModel.IsPublic).To(Equal(core.BoolPtr(true)))
				Expect(createCollectorOptionsModel.ManagedBy).To(Equal(core.StringPtr("ibm")))
				Expect(createCollectorOptionsModel.Description).To(Equal(core.StringPtr("sample collector")))
				Expect(createCollectorOptionsModel.Passphrase).To(Equal(core.StringPtr("secret")))
				Expect(createCollectorOptionsModel.IsUbiImage).To(Equal(core.BoolPtr(true)))
				Expect(createCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCredentialOptions successfully`, func() {
				// Construct an instance of the NewCredentialDisplayFields model
				newCredentialDisplayFieldsModel := new(posturemanagementv2.NewCredentialDisplayFields)
				Expect(newCredentialDisplayFieldsModel).ToNot(BeNil())
				newCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				newCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				newCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				newCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				newCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				newCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				newCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				newCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				newCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				newCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				newCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				newCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				newCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				newCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				newCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				newCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				newCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				newCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ns_tenant_id")
				newCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				newCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				newCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				newCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")
				Expect(newCredentialDisplayFieldsModel.IBMAPIKey).To(Equal(core.StringPtr("sample_api_key")))
				Expect(newCredentialDisplayFieldsModel.AwsClientID).To(Equal(core.StringPtr("sample_client_id")))
				Expect(newCredentialDisplayFieldsModel.AwsClientSecret).To(Equal(core.StringPtr("*********")))
				Expect(newCredentialDisplayFieldsModel.AwsRegion).To(Equal(core.StringPtr("test_region")))
				Expect(newCredentialDisplayFieldsModel.AwsArn).To(Equal(core.StringPtr("sample_arn")))
				Expect(newCredentialDisplayFieldsModel.Username).To(Equal(core.StringPtr("test")))
				Expect(newCredentialDisplayFieldsModel.Password).To(Equal(core.StringPtr("**********")))
				Expect(newCredentialDisplayFieldsModel.AzureClientID).To(Equal(core.StringPtr("azure_124")))
				Expect(newCredentialDisplayFieldsModel.AzureClientSecret).To(Equal(core.StringPtr("************")))
				Expect(newCredentialDisplayFieldsModel.AzureSubscriptionID).To(Equal(core.StringPtr("A32432890")))
				Expect(newCredentialDisplayFieldsModel.AzureResourceGroup).To(Equal(core.StringPtr("azure_res_type")))
				Expect(newCredentialDisplayFieldsModel.DatabaseName).To(Equal(core.StringPtr("sample_db_name")))
				Expect(newCredentialDisplayFieldsModel.WinrmAuthtype).To(Equal(core.StringPtr("sample_auth_type")))
				Expect(newCredentialDisplayFieldsModel.WinrmUsessl).To(Equal(core.StringPtr("ssl_test")))
				Expect(newCredentialDisplayFieldsModel.WinrmPort).To(Equal(core.StringPtr("80")))
				Expect(newCredentialDisplayFieldsModel.Ms365ClientID).To(Equal(core.StringPtr("ms_client_id")))
				Expect(newCredentialDisplayFieldsModel.Ms365ClientSecret).To(Equal(core.StringPtr("ms_client_secret")))
				Expect(newCredentialDisplayFieldsModel.Ms365TenantID).To(Equal(core.StringPtr("ns_tenant_id")))
				Expect(newCredentialDisplayFieldsModel.AuthURL).To(Equal(core.StringPtr("test.example.com")))
				Expect(newCredentialDisplayFieldsModel.ProjectName).To(Equal(core.StringPtr("test_proj")))
				Expect(newCredentialDisplayFieldsModel.UserDomainName).To(Equal(core.StringPtr("user_domain")))
				Expect(newCredentialDisplayFieldsModel.ProjectDomainName).To(Equal(core.StringPtr("stack_domain_name")))

				// Construct an instance of the CredentialGroup model
				credentialGroupModel := new(posturemanagementv2.CredentialGroup)
				Expect(credentialGroupModel).ToNot(BeNil())
				credentialGroupModel.ID = core.StringPtr("1")
				credentialGroupModel.Passphrase = core.StringPtr("passphrase")
				Expect(credentialGroupModel.ID).To(Equal(core.StringPtr("1")))
				Expect(credentialGroupModel.Passphrase).To(Equal(core.StringPtr("passphrase")))

				// Construct an instance of the CreateCredentialOptions model
				createCredentialOptionsEnabled := true
				createCredentialOptionsType := "username_password"
				createCredentialOptionsName := "test_create"
				createCredentialOptionsDescription := "This credential is used for testing."
				var createCredentialOptionsDisplayFields *posturemanagementv2.NewCredentialDisplayFields = nil
				var createCredentialOptionsGroup *posturemanagementv2.CredentialGroup = nil
				createCredentialOptionsPurpose := "discovery_fact_collection_remediation"
				createCredentialOptionsModel := postureManagementService.NewCreateCredentialOptions(createCredentialOptionsEnabled, createCredentialOptionsType, createCredentialOptionsName, createCredentialOptionsDescription, createCredentialOptionsDisplayFields, createCredentialOptionsGroup, createCredentialOptionsPurpose)
				createCredentialOptionsModel.SetEnabled(true)
				createCredentialOptionsModel.SetType("username_password")
				createCredentialOptionsModel.SetName("test_create")
				createCredentialOptionsModel.SetDescription("This credential is used for testing.")
				createCredentialOptionsModel.SetDisplayFields(newCredentialDisplayFieldsModel)
				createCredentialOptionsModel.SetGroup(credentialGroupModel)
				createCredentialOptionsModel.SetPurpose("discovery_fact_collection_remediation")
				createCredentialOptionsModel.SetAccountID("testString")
				createCredentialOptionsModel.SetTransactionID("testString")
				createCredentialOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCredentialOptionsModel).ToNot(BeNil())
				Expect(createCredentialOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createCredentialOptionsModel.Type).To(Equal(core.StringPtr("username_password")))
				Expect(createCredentialOptionsModel.Name).To(Equal(core.StringPtr("test_create")))
				Expect(createCredentialOptionsModel.Description).To(Equal(core.StringPtr("This credential is used for testing.")))
				Expect(createCredentialOptionsModel.DisplayFields).To(Equal(newCredentialDisplayFieldsModel))
				Expect(createCredentialOptionsModel.Group).To(Equal(credentialGroupModel))
				Expect(createCredentialOptionsModel.Purpose).To(Equal(core.StringPtr("discovery_fact_collection_remediation")))
				Expect(createCredentialOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createCredentialOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createCredentialOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateScopeOptions successfully`, func() {
				// Construct an instance of the CreateScopeOptions model
				createScopeOptionsName := "IBMSchema-new-048-test"
				createScopeOptionsDescription := "IBMSchema"
				createScopeOptionsCollectorIds := []string{"20"}
				createScopeOptionsCredentialID := "5"
				createScopeOptionsCredentialType := "on_premise"
				createScopeOptionsModel := postureManagementService.NewCreateScopeOptions(createScopeOptionsName, createScopeOptionsDescription, createScopeOptionsCollectorIds, createScopeOptionsCredentialID, createScopeOptionsCredentialType)
				createScopeOptionsModel.SetName("IBMSchema-new-048-test")
				createScopeOptionsModel.SetDescription("IBMSchema")
				createScopeOptionsModel.SetCollectorIds([]string{"20"})
				createScopeOptionsModel.SetCredentialID("5")
				createScopeOptionsModel.SetCredentialType("on_premise")
				createScopeOptionsModel.SetInterval(int64(10))
				createScopeOptionsModel.SetIsDiscoveryScheduled(true)
				createScopeOptionsModel.SetAccountID("testString")
				createScopeOptionsModel.SetTransactionID("testString")
				createScopeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createScopeOptionsModel).ToNot(BeNil())
				Expect(createScopeOptionsModel.Name).To(Equal(core.StringPtr("IBMSchema-new-048-test")))
				Expect(createScopeOptionsModel.Description).To(Equal(core.StringPtr("IBMSchema")))
				Expect(createScopeOptionsModel.CollectorIds).To(Equal([]string{"20"}))
				Expect(createScopeOptionsModel.CredentialID).To(Equal(core.StringPtr("5")))
				Expect(createScopeOptionsModel.CredentialType).To(Equal(core.StringPtr("on_premise")))
				Expect(createScopeOptionsModel.Interval).To(Equal(core.Int64Ptr(int64(10))))
				Expect(createScopeOptionsModel.IsDiscoveryScheduled).To(Equal(core.BoolPtr(true)))
				Expect(createScopeOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createScopeOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createScopeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateValidationOptions successfully`, func() {
				// Construct an instance of the CreateValidationOptions model
				createValidationOptionsScopeID := "1"
				createValidationOptionsProfileID := "6"
				createValidationOptionsModel := postureManagementService.NewCreateValidationOptions(createValidationOptionsScopeID, createValidationOptionsProfileID)
				createValidationOptionsModel.SetScopeID("1")
				createValidationOptionsModel.SetProfileID("6")
				createValidationOptionsModel.SetGroupProfileID("13")
				createValidationOptionsModel.SetAccountID("testString")
				createValidationOptionsModel.SetTransactionID("testString")
				createValidationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createValidationOptionsModel).ToNot(BeNil())
				Expect(createValidationOptionsModel.ScopeID).To(Equal(core.StringPtr("1")))
				Expect(createValidationOptionsModel.ProfileID).To(Equal(core.StringPtr("6")))
				Expect(createValidationOptionsModel.GroupProfileID).To(Equal(core.StringPtr("13")))
				Expect(createValidationOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createValidationOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createValidationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCredentialGroup successfully`, func() {
				id := "testString"
				passphrase := "testString"
				_model, err := postureManagementService.NewCredentialGroup(id, passphrase)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteCollectorOptions successfully`, func() {
				// Construct an instance of the DeleteCollectorOptions model
				id := "testString"
				deleteCollectorOptionsModel := postureManagementService.NewDeleteCollectorOptions(id)
				deleteCollectorOptionsModel.SetID("testString")
				deleteCollectorOptionsModel.SetAccountID("testString")
				deleteCollectorOptionsModel.SetTransactionID("testString")
				deleteCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectorOptionsModel).ToNot(BeNil())
				Expect(deleteCollectorOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCredentialOptions successfully`, func() {
				// Construct an instance of the DeleteCredentialOptions model
				id := "testString"
				deleteCredentialOptionsModel := postureManagementService.NewDeleteCredentialOptions(id)
				deleteCredentialOptionsModel.SetID("testString")
				deleteCredentialOptionsModel.SetAccountID("testString")
				deleteCredentialOptionsModel.SetTransactionID("testString")
				deleteCredentialOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCredentialOptionsModel).ToNot(BeNil())
				Expect(deleteCredentialOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCredentialOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCredentialOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCredentialOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileOptions successfully`, func() {
				// Construct an instance of the DeleteProfileOptions model
				id := "testString"
				deleteProfileOptionsModel := postureManagementService.NewDeleteProfileOptions(id)
				deleteProfileOptionsModel.SetID("testString")
				deleteProfileOptionsModel.SetAccountID("testString")
				deleteProfileOptionsModel.SetTransactionID("testString")
				deleteProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileOptionsModel).ToNot(BeNil())
				Expect(deleteProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteScopeOptions successfully`, func() {
				// Construct an instance of the DeleteScopeOptions model
				id := "testString"
				deleteScopeOptionsModel := postureManagementService.NewDeleteScopeOptions(id)
				deleteScopeOptionsModel.SetID("testString")
				deleteScopeOptionsModel.SetAccountID("testString")
				deleteScopeOptionsModel.SetTransactionID("testString")
				deleteScopeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteScopeOptionsModel).ToNot(BeNil())
				Expect(deleteScopeOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteScopeOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteScopeOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteScopeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCollectorOptions successfully`, func() {
				// Construct an instance of the GetCollectorOptions model
				id := "testString"
				getCollectorOptionsModel := postureManagementService.NewGetCollectorOptions(id)
				getCollectorOptionsModel.SetID("testString")
				getCollectorOptionsModel.SetAccountID("testString")
				getCollectorOptionsModel.SetTransactionID("testString")
				getCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectorOptionsModel).ToNot(BeNil())
				Expect(getCollectorOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCorrelationIDOptions successfully`, func() {
				// Construct an instance of the GetCorrelationIDOptions model
				correlationID := "testString"
				getCorrelationIDOptionsModel := postureManagementService.NewGetCorrelationIDOptions(correlationID)
				getCorrelationIDOptionsModel.SetCorrelationID("testString")
				getCorrelationIDOptionsModel.SetAccountID("testString")
				getCorrelationIDOptionsModel.SetTransactionID("testString")
				getCorrelationIDOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCorrelationIDOptionsModel).ToNot(BeNil())
				Expect(getCorrelationIDOptionsModel.CorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getCorrelationIDOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getCorrelationIDOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getCorrelationIDOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCredentialOptions successfully`, func() {
				// Construct an instance of the GetCredentialOptions model
				id := "testString"
				getCredentialOptionsModel := postureManagementService.NewGetCredentialOptions(id)
				getCredentialOptionsModel.SetID("testString")
				getCredentialOptionsModel.SetAccountID("testString")
				getCredentialOptionsModel.SetTransactionID("testString")
				getCredentialOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCredentialOptionsModel).ToNot(BeNil())
				Expect(getCredentialOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCredentialOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getCredentialOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getCredentialOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGroupProfileControlsOptions successfully`, func() {
				// Construct an instance of the GetGroupProfileControlsOptions model
				groupID := "testString"
				getGroupProfileControlsOptionsModel := postureManagementService.NewGetGroupProfileControlsOptions(groupID)
				getGroupProfileControlsOptionsModel.SetGroupID("testString")
				getGroupProfileControlsOptionsModel.SetAccountID("testString")
				getGroupProfileControlsOptionsModel.SetTransactionID("testString")
				getGroupProfileControlsOptionsModel.SetOffset(int64(38))
				getGroupProfileControlsOptionsModel.SetLimit(int64(100))
				getGroupProfileControlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGroupProfileControlsOptionsModel).ToNot(BeNil())
				Expect(getGroupProfileControlsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(getGroupProfileControlsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getGroupProfileControlsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getGroupProfileControlsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getGroupProfileControlsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(getGroupProfileControlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileControlsOptions successfully`, func() {
				// Construct an instance of the GetProfileControlsOptions model
				profileID := "testString"
				getProfileControlsOptionsModel := postureManagementService.NewGetProfileControlsOptions(profileID)
				getProfileControlsOptionsModel.SetProfileID("testString")
				getProfileControlsOptionsModel.SetAccountID("testString")
				getProfileControlsOptionsModel.SetTransactionID("testString")
				getProfileControlsOptionsModel.SetOffset(int64(38))
				getProfileControlsOptionsModel.SetLimit(int64(100))
				getProfileControlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileControlsOptionsModel).ToNot(BeNil())
				Expect(getProfileControlsOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileControlsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileControlsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileControlsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getProfileControlsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(getProfileControlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileOptions successfully`, func() {
				// Construct an instance of the GetProfileOptions model
				id := "testString"
				profileType := "testString"
				getProfileOptionsModel := postureManagementService.NewGetProfileOptions(id, profileType)
				getProfileOptionsModel.SetID("testString")
				getProfileOptionsModel.SetProfileType("testString")
				getProfileOptionsModel.SetAccountID("testString")
				getProfileOptionsModel.SetTransactionID("testString")
				getProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileOptionsModel).ToNot(BeNil())
				Expect(getProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.ProfileType).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetScopeDetailsCollectorOptions successfully`, func() {
				// Construct an instance of the GetScopeDetailsCollectorOptions model
				scopeID := "testString"
				getScopeDetailsCollectorOptionsModel := postureManagementService.NewGetScopeDetailsCollectorOptions(scopeID)
				getScopeDetailsCollectorOptionsModel.SetScopeID("testString")
				getScopeDetailsCollectorOptionsModel.SetAccountID("testString")
				getScopeDetailsCollectorOptionsModel.SetTransactionID("testString")
				getScopeDetailsCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getScopeDetailsCollectorOptionsModel).ToNot(BeNil())
				Expect(getScopeDetailsCollectorOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetScopeDetailsCredentialsOptions successfully`, func() {
				// Construct an instance of the GetScopeDetailsCredentialsOptions model
				scopeID := "testString"
				getScopeDetailsCredentialsOptionsModel := postureManagementService.NewGetScopeDetailsCredentialsOptions(scopeID)
				getScopeDetailsCredentialsOptionsModel.SetScopeID("testString")
				getScopeDetailsCredentialsOptionsModel.SetAccountID("testString")
				getScopeDetailsCredentialsOptionsModel.SetTransactionID("testString")
				getScopeDetailsCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getScopeDetailsCredentialsOptionsModel).ToNot(BeNil())
				Expect(getScopeDetailsCredentialsOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCredentialsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCredentialsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetScopeDetailsOptions successfully`, func() {
				// Construct an instance of the GetScopeDetailsOptions model
				id := "testString"
				getScopeDetailsOptionsModel := postureManagementService.NewGetScopeDetailsOptions(id)
				getScopeDetailsOptionsModel.SetID("testString")
				getScopeDetailsOptionsModel.SetAccountID("testString")
				getScopeDetailsOptionsModel.SetTransactionID("testString")
				getScopeDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getScopeDetailsOptionsModel).ToNot(BeNil())
				Expect(getScopeDetailsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetScopeTimelineOptions successfully`, func() {
				// Construct an instance of the GetScopeTimelineOptions model
				scopeID := "testString"
				getScopeTimelineOptionsModel := postureManagementService.NewGetScopeTimelineOptions(scopeID)
				getScopeTimelineOptionsModel.SetScopeID("testString")
				getScopeTimelineOptionsModel.SetAccountID("testString")
				getScopeTimelineOptionsModel.SetTransactionID("testString")
				getScopeTimelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getScopeTimelineOptionsModel).ToNot(BeNil())
				Expect(getScopeTimelineOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeTimelineOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeTimelineOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getScopeTimelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportProfilesOptions successfully`, func() {
				// Construct an instance of the ImportProfilesOptions model
				file := CreateMockReader("This is a mock file.")
				importProfilesOptionsModel := postureManagementService.NewImportProfilesOptions(file)
				importProfilesOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				importProfilesOptionsModel.SetAccountID("testString")
				importProfilesOptionsModel.SetTransactionID("testString")
				importProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importProfilesOptionsModel).ToNot(BeNil())
				Expect(importProfilesOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(importProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(importProfilesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(importProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectorsOptions successfully`, func() {
				// Construct an instance of the ListCollectorsOptions model
				listCollectorsOptionsModel := postureManagementService.NewListCollectorsOptions()
				listCollectorsOptionsModel.SetAccountID("testString")
				listCollectorsOptionsModel.SetTransactionID("testString")
				listCollectorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectorsOptionsModel).ToNot(BeNil())
				Expect(listCollectorsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectorsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCredentialsOptions successfully`, func() {
				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := postureManagementService.NewListCredentialsOptions()
				listCredentialsOptionsModel.SetAccountID("testString")
				listCredentialsOptionsModel.SetOffset(int64(38))
				listCredentialsOptionsModel.SetLimit(int64(100))
				listCredentialsOptionsModel.SetTransactionID("testString")
				listCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCredentialsOptionsModel).ToNot(BeNil())
				Expect(listCredentialsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listCredentialsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listCredentialsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listCredentialsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLatestScansOptions successfully`, func() {
				// Construct an instance of the ListLatestScansOptions model
				listLatestScansOptionsModel := postureManagementService.NewListLatestScansOptions()
				listLatestScansOptionsModel.SetAccountID("testString")
				listLatestScansOptionsModel.SetTransactionID("testString")
				listLatestScansOptionsModel.SetOffset(int64(38))
				listLatestScansOptionsModel.SetLimit(int64(100))
				listLatestScansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLatestScansOptionsModel).ToNot(BeNil())
				Expect(listLatestScansOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listLatestScansOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listLatestScansOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listLatestScansOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listLatestScansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfilesOptions successfully`, func() {
				// Construct an instance of the ListProfilesOptions model
				listProfilesOptionsModel := postureManagementService.NewListProfilesOptions()
				listProfilesOptionsModel.SetAccountID("testString")
				listProfilesOptionsModel.SetTransactionID("testString")
				listProfilesOptionsModel.SetOffset(int64(38))
				listProfilesOptionsModel.SetLimit(int64(100))
				listProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfilesOptionsModel).ToNot(BeNil())
				Expect(listProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listProfilesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listProfilesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListScopesOptions successfully`, func() {
				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := postureManagementService.NewListScopesOptions()
				listScopesOptionsModel.SetAccountID("testString")
				listScopesOptionsModel.SetTransactionID("testString")
				listScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listScopesOptionsModel).ToNot(BeNil())
				Expect(listScopesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceScopeDetailsCollectorOptions successfully`, func() {
				// Construct an instance of the ReplaceScopeDetailsCollectorOptions model
				scopeID := "testString"
				replaceScopeDetailsCollectorOptionsCollectorIds := []string{"7"}
				replaceScopeDetailsCollectorOptionsModel := postureManagementService.NewReplaceScopeDetailsCollectorOptions(scopeID, replaceScopeDetailsCollectorOptionsCollectorIds)
				replaceScopeDetailsCollectorOptionsModel.SetScopeID("testString")
				replaceScopeDetailsCollectorOptionsModel.SetCollectorIds([]string{"7"})
				replaceScopeDetailsCollectorOptionsModel.SetAccountID("testString")
				replaceScopeDetailsCollectorOptionsModel.SetTransactionID("testString")
				replaceScopeDetailsCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceScopeDetailsCollectorOptionsModel).ToNot(BeNil())
				Expect(replaceScopeDetailsCollectorOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCollectorOptionsModel.CollectorIds).To(Equal([]string{"7"}))
				Expect(replaceScopeDetailsCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceScopeDetailsCredentialsOptions successfully`, func() {
				// Construct an instance of the ReplaceScopeDetailsCredentialsOptions model
				scopeID := "testString"
				replaceScopeDetailsCredentialsOptionsCredentialID := "1"
				replaceScopeDetailsCredentialsOptionsModel := postureManagementService.NewReplaceScopeDetailsCredentialsOptions(scopeID, replaceScopeDetailsCredentialsOptionsCredentialID)
				replaceScopeDetailsCredentialsOptionsModel.SetScopeID("testString")
				replaceScopeDetailsCredentialsOptionsModel.SetCredentialID("1")
				replaceScopeDetailsCredentialsOptionsModel.SetCredentialAttribute("Credentials attribute")
				replaceScopeDetailsCredentialsOptionsModel.SetAccountID("testString")
				replaceScopeDetailsCredentialsOptionsModel.SetTransactionID("testString")
				replaceScopeDetailsCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceScopeDetailsCredentialsOptionsModel).ToNot(BeNil())
				Expect(replaceScopeDetailsCredentialsOptionsModel.ScopeID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCredentialsOptionsModel.CredentialID).To(Equal(core.StringPtr("1")))
				Expect(replaceScopeDetailsCredentialsOptionsModel.CredentialAttribute).To(Equal(core.StringPtr("Credentials attribute")))
				Expect(replaceScopeDetailsCredentialsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCredentialsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceScopeDetailsCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewScanSummariesOptions successfully`, func() {
				// Construct an instance of the ScanSummariesOptions model
				reportSettingID := "testString"
				scanSummariesOptionsModel := postureManagementService.NewScanSummariesOptions(reportSettingID)
				scanSummariesOptionsModel.SetReportSettingID("testString")
				scanSummariesOptionsModel.SetAccountID("testString")
				scanSummariesOptionsModel.SetTransactionID("testString")
				scanSummariesOptionsModel.SetOffset(int64(38))
				scanSummariesOptionsModel.SetLimit(int64(100))
				scanSummariesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(scanSummariesOptionsModel).ToNot(BeNil())
				Expect(scanSummariesOptionsModel.ReportSettingID).To(Equal(core.StringPtr("testString")))
				Expect(scanSummariesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(scanSummariesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(scanSummariesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(scanSummariesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(scanSummariesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewScansSummaryOptions successfully`, func() {
				// Construct an instance of the ScansSummaryOptions model
				scanID := "testString"
				profileID := "testString"
				scansSummaryOptionsModel := postureManagementService.NewScansSummaryOptions(scanID, profileID)
				scansSummaryOptionsModel.SetScanID("testString")
				scansSummaryOptionsModel.SetProfileID("testString")
				scansSummaryOptionsModel.SetAccountID("testString")
				scansSummaryOptionsModel.SetTransactionID("testString")
				scansSummaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(scansSummaryOptionsModel).ToNot(BeNil())
				Expect(scansSummaryOptionsModel.ScanID).To(Equal(core.StringPtr("testString")))
				Expect(scansSummaryOptionsModel.ProfileID).To(Equal(core.StringPtr("testString")))
				Expect(scansSummaryOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(scansSummaryOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(scansSummaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewScopeCollector successfully`, func() {
				collectorIds := []string{"testString"}
				_model, err := postureManagementService.NewScopeCollector(collectorIds)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewScopeCredential successfully`, func() {
				credentialID := "testString"
				_model, err := postureManagementService.NewScopeCredential(credentialID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCollectorOptions successfully`, func() {
				// Construct an instance of the UpdateCollectorOptions model
				id := "testString"
				collector := make(map[string]interface{})
				updateCollectorOptionsModel := postureManagementService.NewUpdateCollectorOptions(id, collector)
				updateCollectorOptionsModel.SetID("testString")
				updateCollectorOptionsModel.SetCollector(make(map[string]interface{}))
				updateCollectorOptionsModel.SetAccountID("testString")
				updateCollectorOptionsModel.SetTransactionID("testString")
				updateCollectorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectorOptionsModel).ToNot(BeNil())
				Expect(updateCollectorOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectorOptionsModel.Collector).To(Equal(make(map[string]interface{})))
				Expect(updateCollectorOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectorOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCredentialOptions successfully`, func() {
				// Construct an instance of the UpdateCredentialDisplayFields model
				updateCredentialDisplayFieldsModel := new(posturemanagementv2.UpdateCredentialDisplayFields)
				Expect(updateCredentialDisplayFieldsModel).ToNot(BeNil())
				updateCredentialDisplayFieldsModel.IBMAPIKey = core.StringPtr("sample_api_key")
				updateCredentialDisplayFieldsModel.AwsClientID = core.StringPtr("sample_client_id")
				updateCredentialDisplayFieldsModel.AwsClientSecret = core.StringPtr("*********")
				updateCredentialDisplayFieldsModel.AwsRegion = core.StringPtr("test_region")
				updateCredentialDisplayFieldsModel.AwsArn = core.StringPtr("sample_arn")
				updateCredentialDisplayFieldsModel.Username = core.StringPtr("test")
				updateCredentialDisplayFieldsModel.Password = core.StringPtr("**********")
				updateCredentialDisplayFieldsModel.AzureClientID = core.StringPtr("azure_124")
				updateCredentialDisplayFieldsModel.AzureClientSecret = core.StringPtr("************")
				updateCredentialDisplayFieldsModel.AzureSubscriptionID = core.StringPtr("A32432890")
				updateCredentialDisplayFieldsModel.AzureResourceGroup = core.StringPtr("azure_res_type")
				updateCredentialDisplayFieldsModel.DatabaseName = core.StringPtr("sample_db_name")
				updateCredentialDisplayFieldsModel.WinrmAuthtype = core.StringPtr("sample_auth_type")
				updateCredentialDisplayFieldsModel.WinrmUsessl = core.StringPtr("ssl_test")
				updateCredentialDisplayFieldsModel.WinrmPort = core.StringPtr("80")
				updateCredentialDisplayFieldsModel.Ms365ClientID = core.StringPtr("ms_client_id")
				updateCredentialDisplayFieldsModel.Ms365ClientSecret = core.StringPtr("ms_client_secret")
				updateCredentialDisplayFieldsModel.Ms365TenantID = core.StringPtr("ms_tenant_id")
				updateCredentialDisplayFieldsModel.AuthURL = core.StringPtr("test.example.com")
				updateCredentialDisplayFieldsModel.ProjectName = core.StringPtr("test_proj")
				updateCredentialDisplayFieldsModel.UserDomainName = core.StringPtr("user_domain")
				updateCredentialDisplayFieldsModel.ProjectDomainName = core.StringPtr("stack_domain_name")
				Expect(updateCredentialDisplayFieldsModel.IBMAPIKey).To(Equal(core.StringPtr("sample_api_key")))
				Expect(updateCredentialDisplayFieldsModel.AwsClientID).To(Equal(core.StringPtr("sample_client_id")))
				Expect(updateCredentialDisplayFieldsModel.AwsClientSecret).To(Equal(core.StringPtr("*********")))
				Expect(updateCredentialDisplayFieldsModel.AwsRegion).To(Equal(core.StringPtr("test_region")))
				Expect(updateCredentialDisplayFieldsModel.AwsArn).To(Equal(core.StringPtr("sample_arn")))
				Expect(updateCredentialDisplayFieldsModel.Username).To(Equal(core.StringPtr("test")))
				Expect(updateCredentialDisplayFieldsModel.Password).To(Equal(core.StringPtr("**********")))
				Expect(updateCredentialDisplayFieldsModel.AzureClientID).To(Equal(core.StringPtr("azure_124")))
				Expect(updateCredentialDisplayFieldsModel.AzureClientSecret).To(Equal(core.StringPtr("************")))
				Expect(updateCredentialDisplayFieldsModel.AzureSubscriptionID).To(Equal(core.StringPtr("A32432890")))
				Expect(updateCredentialDisplayFieldsModel.AzureResourceGroup).To(Equal(core.StringPtr("azure_res_type")))
				Expect(updateCredentialDisplayFieldsModel.DatabaseName).To(Equal(core.StringPtr("sample_db_name")))
				Expect(updateCredentialDisplayFieldsModel.WinrmAuthtype).To(Equal(core.StringPtr("sample_auth_type")))
				Expect(updateCredentialDisplayFieldsModel.WinrmUsessl).To(Equal(core.StringPtr("ssl_test")))
				Expect(updateCredentialDisplayFieldsModel.WinrmPort).To(Equal(core.StringPtr("80")))
				Expect(updateCredentialDisplayFieldsModel.Ms365ClientID).To(Equal(core.StringPtr("ms_client_id")))
				Expect(updateCredentialDisplayFieldsModel.Ms365ClientSecret).To(Equal(core.StringPtr("ms_client_secret")))
				Expect(updateCredentialDisplayFieldsModel.Ms365TenantID).To(Equal(core.StringPtr("ms_tenant_id")))
				Expect(updateCredentialDisplayFieldsModel.AuthURL).To(Equal(core.StringPtr("test.example.com")))
				Expect(updateCredentialDisplayFieldsModel.ProjectName).To(Equal(core.StringPtr("test_proj")))
				Expect(updateCredentialDisplayFieldsModel.UserDomainName).To(Equal(core.StringPtr("user_domain")))
				Expect(updateCredentialDisplayFieldsModel.ProjectDomainName).To(Equal(core.StringPtr("stack_domain_name")))

				// Construct an instance of the UpdateCredentialOptions model
				id := "testString"
				updateCredentialOptionsModel := postureManagementService.NewUpdateCredentialOptions(id)
				updateCredentialOptionsModel.SetID("testString")
				updateCredentialOptionsModel.SetEnabled(true)
				updateCredentialOptionsModel.SetType("username_password")
				updateCredentialOptionsModel.SetName("test_create")
				updateCredentialOptionsModel.SetDescription("This credential is used for testing.")
				updateCredentialOptionsModel.SetDisplayFields(updateCredentialDisplayFieldsModel)
				updateCredentialOptionsModel.SetPurpose("discovery_fact_collection_remediation")
				updateCredentialOptionsModel.SetAccountID("testString")
				updateCredentialOptionsModel.SetTransactionID("testString")
				updateCredentialOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCredentialOptionsModel).ToNot(BeNil())
				Expect(updateCredentialOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCredentialOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateCredentialOptionsModel.Type).To(Equal(core.StringPtr("username_password")))
				Expect(updateCredentialOptionsModel.Name).To(Equal(core.StringPtr("test_create")))
				Expect(updateCredentialOptionsModel.Description).To(Equal(core.StringPtr("This credential is used for testing.")))
				Expect(updateCredentialOptionsModel.DisplayFields).To(Equal(updateCredentialDisplayFieldsModel))
				Expect(updateCredentialOptionsModel.Purpose).To(Equal(core.StringPtr("discovery_fact_collection_remediation")))
				Expect(updateCredentialOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateCredentialOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCredentialOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProfilesOptions successfully`, func() {
				// Construct an instance of the UpdateProfilesOptions model
				id := "testString"
				updateProfilesOptionsModel := postureManagementService.NewUpdateProfilesOptions(id)
				updateProfilesOptionsModel.SetID("testString")
				updateProfilesOptionsModel.SetName("AT_Controls_Testing")
				updateProfilesOptionsModel.SetDescription("AT Controls")
				updateProfilesOptionsModel.SetBaseProfile("CIS IBM Foundations Benchmark 1.0.0")
				updateProfilesOptionsModel.SetType("predefined")
				updateProfilesOptionsModel.SetIsEnabled(true)
				updateProfilesOptionsModel.SetControlIds([]string{"9980", "9979", "9994"})
				updateProfilesOptionsModel.SetAccountID("testString")
				updateProfilesOptionsModel.SetTransactionID("testString")
				updateProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProfilesOptionsModel).ToNot(BeNil())
				Expect(updateProfilesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfilesOptionsModel.Name).To(Equal(core.StringPtr("AT_Controls_Testing")))
				Expect(updateProfilesOptionsModel.Description).To(Equal(core.StringPtr("AT Controls")))
				Expect(updateProfilesOptionsModel.BaseProfile).To(Equal(core.StringPtr("CIS IBM Foundations Benchmark 1.0.0")))
				Expect(updateProfilesOptionsModel.Type).To(Equal(core.StringPtr("predefined")))
				Expect(updateProfilesOptionsModel.IsEnabled).To(Equal(core.BoolPtr(true)))
				Expect(updateProfilesOptionsModel.ControlIds).To(Equal([]string{"9980", "9979", "9994"}))
				Expect(updateProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfilesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateScopeDetailsOptions successfully`, func() {
				// Construct an instance of the UpdateScopeDetailsOptions model
				id := "testString"
				updateScopeDetailsOptionsModel := postureManagementService.NewUpdateScopeDetailsOptions(id)
				updateScopeDetailsOptionsModel.SetID("testString")
				updateScopeDetailsOptionsModel.SetName("Scope Test1")
				updateScopeDetailsOptionsModel.SetDescription("Scope Description")
				updateScopeDetailsOptionsModel.SetAccountID("testString")
				updateScopeDetailsOptionsModel.SetTransactionID("testString")
				updateScopeDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateScopeDetailsOptionsModel).ToNot(BeNil())
				Expect(updateScopeDetailsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateScopeDetailsOptionsModel.Name).To(Equal(core.StringPtr("Scope Test1")))
				Expect(updateScopeDetailsOptionsModel.Description).To(Equal(core.StringPtr("Scope Description")))
				Expect(updateScopeDetailsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateScopeDetailsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateScopeDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

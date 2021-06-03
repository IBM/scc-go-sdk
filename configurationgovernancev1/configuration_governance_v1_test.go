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

package configurationgovernancev1_test

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
	"github.com/ibm/scc-go-sdk/configurationgovernancev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ConfigurationGovernanceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(configurationGovernanceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(configurationGovernanceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "https://configurationgovernancev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(configurationGovernanceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_URL":       "https://configurationgovernancev1/api",
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL: "https://testService/api",
				})
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})
				err := configurationGovernanceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := configurationGovernanceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != configurationGovernanceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(configurationGovernanceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(configurationGovernanceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_URL":       "https://configurationgovernancev1/api",
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(configurationGovernanceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONFIGURATION_GOVERNANCE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1UsingExternalConfig(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(configurationGovernanceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = configurationgovernancev1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configurationgovernancev1.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configurationgovernancev1.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu.compliance.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = configurationgovernancev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateRules(createRulesOptions *CreateRulesOptions) - Operation response error`, func() {
		createRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRules with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRules(createRulesOptions *CreateRulesOptions)`, func() {
		createRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"rules": [{"request_id": "3cebc877-58e7-44a5-a292-32114fa73558", "status_code": 201, "rule": {"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}, "errors": [{"code": "bad_request", "message": "The rule is missing an account ID"}], "trace": "861263b4-cee3-4514-8d8c-05d17308e6eb"}]}`)
				}))
			})
			It(`Invoke CreateRules successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.CreateRulesWithContext(ctx, createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.CreateRulesWithContext(ctx, createRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRulesPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"rules": [{"request_id": "3cebc877-58e7-44a5-a292-32114fa73558", "status_code": 201, "rule": {"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}, "errors": [{"code": "bad_request", "message": "The rule is missing an account ID"}], "trace": "861263b4-cee3-4514-8d8c-05d17308e6eb"}]}`)
				}))
			})
			It(`Invoke CreateRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.CreateRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRules with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRulesOptions model with no property values
				createRulesOptionsModelNew := new(configurationgovernancev1.CreateRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.CreateRules(createRulesOptionsModelNew)
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
			It(`Invoke CreateRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("disallow")

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsModel := new(configurationgovernancev1.CreateRulesOptions)
				createRulesOptionsModel.Rules = []configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}
				createRulesOptionsModel.TransactionID = core.StringPtr("testString")
				createRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.CreateRules(createRulesOptionsModel)
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
		listRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))
					// TODO: Add check for attached query parameter
					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))
					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRules with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModel)
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
		listRulesPath := "/config/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))
					// TODO: Add check for attached query parameter
					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))
					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "rules": [{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.ListRulesWithContext(ctx, listRulesOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"531fc3e28bfc43c5a2cea07786d93f5c"}))
					// TODO: Add check for attached query parameter
					Expect(req.URL.Query()["labels"]).To(Equal([]string{"SOC2,ITCS300"}))
					Expect(req.URL.Query()["scopes"]).To(Equal([]string{"scope_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "rules": [{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRulesOptions model with no property values
				listRulesOptionsModelNew := new(configurationgovernancev1.ListRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.ListRules(listRulesOptionsModelNew)
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
			It(`Invoke ListRules successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(configurationgovernancev1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Attached = core.BoolPtr(true)
				listRulesOptionsModel.Labels = core.StringPtr("SOC2,ITCS300")
				listRulesOptionsModel.Scopes = core.StringPtr("scope_id")
				listRulesOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRulesOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.ListRules(listRulesOptionsModel)
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
	Describe(`GetRule(getRuleOptions *GetRuleOptions) - Operation response error`, func() {
		getRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRule with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModel)
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
		getRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.GetRuleWithContext(ctx, getRuleOptionsModel)
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(configurationgovernancev1.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.GetRule(getRuleOptionsModelNew)
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
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(configurationgovernancev1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.GetRule(getRuleOptionsModel)
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
	Describe(`UpdateRule(updateRuleOptions *UpdateRuleOptions) - Operation response error`, func() {
		updateRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRule with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"SOC2", "ITCS300"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRule(updateRuleOptions *UpdateRuleOptions)`, func() {
		updateRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
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
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke UpdateRule successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"SOC2", "ITCS300"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.UpdateRuleWithContext(ctx, updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.UpdateRuleWithContext(ctx, updateRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRulePath))
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
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "name": "Name", "description": "Description", "rule_type": "user_defined", "target": {"service_name": "cloud-object-storage", "resource_kind": "bucket", "additional_target_attributes": [{"name": "Name", "value": "Value", "operator": "string_equals"}]}, "required_config": {"description": "Description", "property": "public_access_enabled", "operator": "is_true", "value": "Value"}, "enforcement_actions": [{"action": "audit_log"}], "labels": ["Label"], "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "creation_date": "2020-01-10T05:23:19.000Z", "created_by": "CreatedBy", "modification_date": "2020-01-10T05:23:19.000Z", "modified_by": "ModifiedBy", "number_of_attachments": 3}`)
				}))
			})
			It(`Invoke UpdateRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.UpdateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"SOC2", "ITCS300"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"SOC2", "ITCS300"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRuleOptions model with no property values
				updateRuleOptionsModelNew := new(configurationgovernancev1.UpdateRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.UpdateRule(updateRuleOptionsModelNew)
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
			It(`Invoke UpdateRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				enforcementActionModel.Action = core.StringPtr("audit_log")

				// Construct an instance of the UpdateRuleOptions model
				updateRuleOptionsModel := new(configurationgovernancev1.UpdateRuleOptions)
				updateRuleOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleOptionsModel.Name = core.StringPtr("Disable public access")
				updateRuleOptionsModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.Target = targetResourceModel
				updateRuleOptionsModel.RequiredConfig = ruleRequiredConfigModel
				updateRuleOptionsModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				updateRuleOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.RuleType = core.StringPtr("user_defined")
				updateRuleOptionsModel.Labels = []string{"SOC2", "ITCS300"}
				updateRuleOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.UpdateRule(updateRuleOptionsModel)
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
		deleteRulePath := "/config/v1/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRule successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configurationGovernanceService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configurationgovernancev1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configurationGovernanceService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(configurationgovernancev1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configurationGovernanceService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(configurationgovernancev1.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configurationGovernanceService.DeleteRule(deleteRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRuleAttachments(createRuleAttachmentsOptions *CreateRuleAttachmentsOptions) - Operation response error`, func() {
		createRuleAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRuleAttachmentsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRuleAttachments with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateRuleAttachmentsOptions model
				createRuleAttachmentsOptionsModel := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				createRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Attachments = []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}
				createRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRuleAttachments(createRuleAttachmentsOptions *CreateRuleAttachmentsOptions)`, func() {
		createRuleAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRuleAttachmentsPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke CreateRuleAttachments successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateRuleAttachmentsOptions model
				createRuleAttachmentsOptionsModel := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				createRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Attachments = []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}
				createRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.CreateRuleAttachmentsWithContext(ctx, createRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.CreateRuleAttachmentsWithContext(ctx, createRuleAttachmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRuleAttachmentsPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke CreateRuleAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.CreateRuleAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateRuleAttachmentsOptions model
				createRuleAttachmentsOptionsModel := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				createRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Attachments = []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}
				createRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRuleAttachments with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateRuleAttachmentsOptions model
				createRuleAttachmentsOptionsModel := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				createRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Attachments = []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}
				createRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRuleAttachmentsOptions model with no property values
				createRuleAttachmentsOptionsModelNew := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModelNew)
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
			It(`Invoke CreateRuleAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}

				// Construct an instance of the CreateRuleAttachmentsOptions model
				createRuleAttachmentsOptionsModel := new(configurationgovernancev1.CreateRuleAttachmentsOptions)
				createRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Attachments = []configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}
				createRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.CreateRuleAttachments(createRuleAttachmentsOptionsModel)
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
	Describe(`ListRuleAttachments(listRuleAttachmentsOptions *ListRuleAttachmentsOptions) - Operation response error`, func() {
		listRuleAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRuleAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRuleAttachments with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRuleAttachmentsOptions model
				listRuleAttachmentsOptionsModel := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				listRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRuleAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRuleAttachments(listRuleAttachmentsOptions *ListRuleAttachmentsOptions)`, func() {
		listRuleAttachmentsPath := "/config/v1/rules/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRuleAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke ListRuleAttachments successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the ListRuleAttachmentsOptions model
				listRuleAttachmentsOptionsModel := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				listRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRuleAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.ListRuleAttachmentsWithContext(ctx, listRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.ListRuleAttachmentsWithContext(ctx, listRuleAttachmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRuleAttachmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1000))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 1000, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "attachments": [{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}]}`)
				}))
			})
			It(`Invoke ListRuleAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.ListRuleAttachments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRuleAttachmentsOptions model
				listRuleAttachmentsOptionsModel := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				listRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRuleAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRuleAttachments with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRuleAttachmentsOptions model
				listRuleAttachmentsOptionsModel := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				listRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRuleAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRuleAttachmentsOptions model with no property values
				listRuleAttachmentsOptionsModelNew := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModelNew)
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
			It(`Invoke ListRuleAttachments successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the ListRuleAttachmentsOptions model
				listRuleAttachmentsOptionsModel := new(configurationgovernancev1.ListRuleAttachmentsOptions)
				listRuleAttachmentsOptionsModel.RuleID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listRuleAttachmentsOptionsModel.Limit = core.Int64Ptr(int64(1000))
				listRuleAttachmentsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listRuleAttachmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.ListRuleAttachments(listRuleAttachmentsOptionsModel)
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
	Describe(`GetRuleAttachment(getRuleAttachmentOptions *GetRuleAttachmentOptions) - Operation response error`, func() {
		getRuleAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRuleAttachmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRuleAttachment with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleAttachmentOptions model
				getRuleAttachmentOptionsModel := new(configurationgovernancev1.GetRuleAttachmentOptions)
				getRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRuleAttachment(getRuleAttachmentOptions *GetRuleAttachmentOptions)`, func() {
		getRuleAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRuleAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke GetRuleAttachment successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleAttachmentOptions model
				getRuleAttachmentOptionsModel := new(configurationgovernancev1.GetRuleAttachmentOptions)
				getRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.GetRuleAttachmentWithContext(ctx, getRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.GetRuleAttachmentWithContext(ctx, getRuleAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRuleAttachmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "attachment-fc7b9a77-1c85-406c-b346-f3f5bb9aa7e2", "rule_id": "rule-81f3db5e-f9db-4c46-9de3-a4a76e66adbf", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke GetRuleAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.GetRuleAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleAttachmentOptions model
				getRuleAttachmentOptionsModel := new(configurationgovernancev1.GetRuleAttachmentOptions)
				getRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRuleAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleAttachmentOptions model
				getRuleAttachmentOptionsModel := new(configurationgovernancev1.GetRuleAttachmentOptions)
				getRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleAttachmentOptions model with no property values
				getRuleAttachmentOptionsModelNew := new(configurationgovernancev1.GetRuleAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModelNew)
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
			It(`Invoke GetRuleAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the GetRuleAttachmentOptions model
				getRuleAttachmentOptionsModel := new(configurationgovernancev1.GetRuleAttachmentOptions)
				getRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.GetRuleAttachment(getRuleAttachmentOptionsModel)
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
	Describe(`UpdateRuleAttachment(updateRuleAttachmentOptions *UpdateRuleAttachmentOptions) - Operation response error`, func() {
		updateRuleAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRuleAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRuleAttachment with error: Operation response processing error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateRuleAttachmentOptions model
				updateRuleAttachmentOptionsModel := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				updateRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateRuleAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				configurationGovernanceService.EnableRetries(0, 0)
				result, response, operationErr = configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRuleAttachment(updateRuleAttachmentOptions *UpdateRuleAttachmentOptions)`, func() {
		updateRuleAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRuleAttachmentPath))
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
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "AttachmentID", "template_id": "TemplateID", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke UpdateRuleAttachment successfully with retries`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())
				configurationGovernanceService.EnableRetries(0, 0)

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateRuleAttachmentOptions model
				updateRuleAttachmentOptionsModel := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				updateRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateRuleAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := configurationGovernanceService.UpdateRuleAttachmentWithContext(ctx, updateRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				configurationGovernanceService.DisableRetries()
				result, response, operationErr := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = configurationGovernanceService.UpdateRuleAttachmentWithContext(ctx, updateRuleAttachmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRuleAttachmentPath))
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
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachment_id": "AttachmentID", "template_id": "TemplateID", "account_id": "AccountID", "included_scope": {"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}, "excluded_scopes": [{"note": "Note", "scope_id": "ScopeID", "scope_type": "enterprise"}]}`)
				}))
			})
			It(`Invoke UpdateRuleAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := configurationGovernanceService.UpdateRuleAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateRuleAttachmentOptions model
				updateRuleAttachmentOptionsModel := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				updateRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateRuleAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRuleAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateRuleAttachmentOptions model
				updateRuleAttachmentOptionsModel := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				updateRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateRuleAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRuleAttachmentOptions model with no property values
				updateRuleAttachmentOptionsModelNew := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModelNew)
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
			It(`Invoke UpdateRuleAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")

				// Construct an instance of the UpdateRuleAttachmentOptions model
				updateRuleAttachmentOptionsModel := new(configurationgovernancev1.UpdateRuleAttachmentOptions)
				updateRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.IncludedScope = ruleScopeModel
				updateRuleAttachmentOptionsModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				updateRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				updateRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := configurationGovernanceService.UpdateRuleAttachment(updateRuleAttachmentOptionsModel)
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
	Describe(`DeleteRuleAttachment(deleteRuleAttachmentOptions *DeleteRuleAttachmentOptions)`, func() {
		deleteRuleAttachmentPath := "/config/v1/rules/testString/attachments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRuleAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRuleAttachment successfully`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := configurationGovernanceService.DeleteRuleAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleAttachmentOptions model
				deleteRuleAttachmentOptionsModel := new(configurationgovernancev1.DeleteRuleAttachmentOptions)
				deleteRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = configurationGovernanceService.DeleteRuleAttachment(deleteRuleAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRuleAttachment with error: Operation validation and request error`, func() {
				configurationGovernanceService, serviceErr := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(configurationGovernanceService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleAttachmentOptions model
				deleteRuleAttachmentOptionsModel := new(configurationgovernancev1.DeleteRuleAttachmentOptions)
				deleteRuleAttachmentOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.AttachmentID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := configurationGovernanceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := configurationGovernanceService.DeleteRuleAttachment(deleteRuleAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleAttachmentOptions model with no property values
				deleteRuleAttachmentOptionsModelNew := new(configurationgovernancev1.DeleteRuleAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = configurationGovernanceService.DeleteRuleAttachment(deleteRuleAttachmentOptionsModelNew)
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
			configurationGovernanceService, _ := configurationgovernancev1.NewConfigurationGovernanceV1(&configurationgovernancev1.ConfigurationGovernanceV1Options{
				URL:           "http://configurationgovernancev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateRuleAttachmentsOptions successfully`, func() {
				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				Expect(ruleScopeModel).ToNot(BeNil())
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")
				Expect(ruleScopeModel.Note).To(Equal(core.StringPtr("My enterprise")))
				Expect(ruleScopeModel.ScopeID).To(Equal(core.StringPtr("282cf433ac91493ba860480d92519990")))
				Expect(ruleScopeModel.ScopeType).To(Equal(core.StringPtr("enterprise")))

				// Construct an instance of the RuleAttachmentRequest model
				ruleAttachmentRequestModel := new(configurationgovernancev1.RuleAttachmentRequest)
				Expect(ruleAttachmentRequestModel).ToNot(BeNil())
				ruleAttachmentRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleAttachmentRequestModel.IncludedScope = ruleScopeModel
				ruleAttachmentRequestModel.ExcludedScopes = []configurationgovernancev1.RuleScope{*ruleScopeModel}
				Expect(ruleAttachmentRequestModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(ruleAttachmentRequestModel.IncludedScope).To(Equal(ruleScopeModel))
				Expect(ruleAttachmentRequestModel.ExcludedScopes).To(Equal([]configurationgovernancev1.RuleScope{*ruleScopeModel}))

				// Construct an instance of the CreateRuleAttachmentsOptions model
				ruleID := "testString"
				createRuleAttachmentsOptionsAttachments := []configurationgovernancev1.RuleAttachmentRequest{}
				createRuleAttachmentsOptionsModel := configurationGovernanceService.NewCreateRuleAttachmentsOptions(ruleID, createRuleAttachmentsOptionsAttachments)
				createRuleAttachmentsOptionsModel.SetRuleID("testString")
				createRuleAttachmentsOptionsModel.SetAttachments([]configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel})
				createRuleAttachmentsOptionsModel.SetTransactionID("testString")
				createRuleAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRuleAttachmentsOptionsModel).ToNot(BeNil())
				Expect(createRuleAttachmentsOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleAttachmentsOptionsModel.Attachments).To(Equal([]configurationgovernancev1.RuleAttachmentRequest{*ruleAttachmentRequestModel}))
				Expect(createRuleAttachmentsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRuleRequest successfully`, func() {
				var rule *configurationgovernancev1.RuleRequest = nil
				_, err := configurationGovernanceService.NewCreateRuleRequest(rule)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateRulesOptions successfully`, func() {
				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				Expect(targetResourceAdditionalTargetAttributesItemModel).ToNot(BeNil())
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("resource_id")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")
				Expect(targetResourceAdditionalTargetAttributesItemModel.Name).To(Equal(core.StringPtr("resource_id")))
				Expect(targetResourceAdditionalTargetAttributesItemModel.Value).To(Equal(core.StringPtr("81f3db5e-f9db-4c46-9de3-a4a76e66adbf")))
				Expect(targetResourceAdditionalTargetAttributesItemModel.Operator).To(Equal(core.StringPtr("string_equals")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}
				Expect(targetResourceModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(targetResourceModel.ResourceKind).To(Equal(core.StringPtr("service")))
				Expect(targetResourceModel.AdditionalTargetAttributes).To(Equal([]configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}))

				// Construct an instance of the RuleConditionSingleProperty model
				ruleConditionModel := new(configurationgovernancev1.RuleConditionSingleProperty)
				Expect(ruleConditionModel).ToNot(BeNil())
				ruleConditionModel.Description = core.StringPtr("testString")
				ruleConditionModel.Property = core.StringPtr("public_access_enabled")
				ruleConditionModel.Operator = core.StringPtr("is_false")
				ruleConditionModel.Value = core.StringPtr("testString")
				Expect(ruleConditionModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionModel.Property).To(Equal(core.StringPtr("public_access_enabled")))
				Expect(ruleConditionModel.Operator).To(Equal(core.StringPtr("is_false")))
				Expect(ruleConditionModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RuleRequiredConfigMultiplePropertiesConditionAnd model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigMultiplePropertiesConditionAnd)
				Expect(ruleRequiredConfigModel).ToNot(BeNil())
				ruleRequiredConfigModel.Description = core.StringPtr("Public access check")
				ruleRequiredConfigModel.And = []configurationgovernancev1.RuleConditionIntf{ruleConditionModel}
				Expect(ruleRequiredConfigModel.Description).To(Equal(core.StringPtr("Public access check")))
				Expect(ruleRequiredConfigModel.And).To(Equal([]configurationgovernancev1.RuleConditionIntf{ruleConditionModel}))

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				Expect(enforcementActionModel).ToNot(BeNil())
				enforcementActionModel.Action = core.StringPtr("disallow")
				Expect(enforcementActionModel.Action).To(Equal(core.StringPtr("disallow")))

				// Construct an instance of the RuleRequest model
				ruleRequestModel := new(configurationgovernancev1.RuleRequest)
				Expect(ruleRequestModel).ToNot(BeNil())
				ruleRequestModel.AccountID = core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")
				ruleRequestModel.Name = core.StringPtr("Disable public access")
				ruleRequestModel.Description = core.StringPtr("Ensure that public access to account resources is disabled.")
				ruleRequestModel.RuleType = core.StringPtr("user_defined")
				ruleRequestModel.Target = targetResourceModel
				ruleRequestModel.RequiredConfig = ruleRequiredConfigModel
				ruleRequestModel.EnforcementActions = []configurationgovernancev1.EnforcementAction{*enforcementActionModel}
				ruleRequestModel.Labels = []string{"Access", "IAM"}
				Expect(ruleRequestModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(ruleRequestModel.Name).To(Equal(core.StringPtr("Disable public access")))
				Expect(ruleRequestModel.Description).To(Equal(core.StringPtr("Ensure that public access to account resources is disabled.")))
				Expect(ruleRequestModel.RuleType).To(Equal(core.StringPtr("user_defined")))
				Expect(ruleRequestModel.Target).To(Equal(targetResourceModel))
				Expect(ruleRequestModel.RequiredConfig).To(Equal(ruleRequiredConfigModel))
				Expect(ruleRequestModel.EnforcementActions).To(Equal([]configurationgovernancev1.EnforcementAction{*enforcementActionModel}))
				Expect(ruleRequestModel.Labels).To(Equal([]string{"Access", "IAM"}))

				// Construct an instance of the CreateRuleRequest model
				createRuleRequestModel := new(configurationgovernancev1.CreateRuleRequest)
				Expect(createRuleRequestModel).ToNot(BeNil())
				createRuleRequestModel.RequestID = core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")
				createRuleRequestModel.Rule = ruleRequestModel
				Expect(createRuleRequestModel.RequestID).To(Equal(core.StringPtr("3cebc877-58e7-44a5-a292-32114fa73558")))
				Expect(createRuleRequestModel.Rule).To(Equal(ruleRequestModel))

				// Construct an instance of the CreateRulesOptions model
				createRulesOptionsRules := []configurationgovernancev1.CreateRuleRequest{}
				createRulesOptionsModel := configurationGovernanceService.NewCreateRulesOptions(createRulesOptionsRules)
				createRulesOptionsModel.SetRules([]configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel})
				createRulesOptionsModel.SetTransactionID("testString")
				createRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRulesOptionsModel).ToNot(BeNil())
				Expect(createRulesOptionsModel.Rules).To(Equal([]configurationgovernancev1.CreateRuleRequest{*createRuleRequestModel}))
				Expect(createRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleAttachmentOptions successfully`, func() {
				// Construct an instance of the DeleteRuleAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				deleteRuleAttachmentOptionsModel := configurationGovernanceService.NewDeleteRuleAttachmentOptions(ruleID, attachmentID)
				deleteRuleAttachmentOptionsModel.SetRuleID("testString")
				deleteRuleAttachmentOptionsModel.SetAttachmentID("testString")
				deleteRuleAttachmentOptionsModel.SetTransactionID("testString")
				deleteRuleAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleAttachmentOptionsModel).ToNot(BeNil())
				Expect(deleteRuleAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				deleteRuleOptionsModel := configurationGovernanceService.NewDeleteRuleOptions(ruleID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetTransactionID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnforcementAction successfully`, func() {
				action := "audit_log"
				model, err := configurationGovernanceService.NewEnforcementAction(action)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetRuleAttachmentOptions successfully`, func() {
				// Construct an instance of the GetRuleAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				getRuleAttachmentOptionsModel := configurationGovernanceService.NewGetRuleAttachmentOptions(ruleID, attachmentID)
				getRuleAttachmentOptionsModel.SetRuleID("testString")
				getRuleAttachmentOptionsModel.SetAttachmentID("testString")
				getRuleAttachmentOptionsModel.SetTransactionID("testString")
				getRuleAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleAttachmentOptionsModel).ToNot(BeNil())
				Expect(getRuleAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				getRuleOptionsModel := configurationGovernanceService.NewGetRuleOptions(ruleID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetTransactionID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRuleAttachmentsOptions successfully`, func() {
				// Construct an instance of the ListRuleAttachmentsOptions model
				ruleID := "testString"
				listRuleAttachmentsOptionsModel := configurationGovernanceService.NewListRuleAttachmentsOptions(ruleID)
				listRuleAttachmentsOptionsModel.SetRuleID("testString")
				listRuleAttachmentsOptionsModel.SetTransactionID("testString")
				listRuleAttachmentsOptionsModel.SetLimit(int64(1000))
				listRuleAttachmentsOptionsModel.SetOffset(int64(38))
				listRuleAttachmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRuleAttachmentsOptionsModel).ToNot(BeNil())
				Expect(listRuleAttachmentsOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(listRuleAttachmentsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listRuleAttachmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(listRuleAttachmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listRuleAttachmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				accountID := "531fc3e28bfc43c5a2cea07786d93f5c"
				listRulesOptionsModel := configurationGovernanceService.NewListRulesOptions(accountID)
				listRulesOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				listRulesOptionsModel.SetTransactionID("testString")
				listRulesOptionsModel.SetAttached(true)
				listRulesOptionsModel.SetLabels("SOC2,ITCS300")
				listRulesOptionsModel.SetScopes("scope_id")
				listRulesOptionsModel.SetLimit(int64(1000))
				listRulesOptionsModel.SetOffset(int64(38))
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(listRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Attached).To(Equal(core.BoolPtr(true)))
				Expect(listRulesOptionsModel.Labels).To(Equal(core.StringPtr("SOC2,ITCS300")))
				Expect(listRulesOptionsModel.Scopes).To(Equal(core.StringPtr("scope_id")))
				Expect(listRulesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(listRulesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleAttachmentRequest successfully`, func() {
				accountID := "testString"
				var includedScope *configurationgovernancev1.RuleScope = nil
				_, err := configurationGovernanceService.NewRuleAttachmentRequest(accountID, includedScope)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRuleRequest successfully`, func() {
				name := "testString"
				description := "testString"
				var target *configurationgovernancev1.TargetResource = nil
				var requiredConfig configurationgovernancev1.RuleRequiredConfigIntf = nil
				enforcementActions := []configurationgovernancev1.EnforcementAction{}
				_, err := configurationGovernanceService.NewRuleRequest(name, description, target, requiredConfig, enforcementActions)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRuleScope successfully`, func() {
				scopeID := "testString"
				scopeType := "enterprise"
				model, err := configurationGovernanceService.NewRuleScope(scopeID, scopeType)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetResource successfully`, func() {
				serviceName := "cloud-object-storage"
				resourceKind := "bucket"
				model, err := configurationGovernanceService.NewTargetResource(serviceName, resourceKind)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetResourceAdditionalTargetAttributesItem successfully`, func() {
				name := "testString"
				value := "testString"
				operator := "string_equals"
				model, err := configurationGovernanceService.NewTargetResourceAdditionalTargetAttributesItem(name, value, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateRuleAttachmentOptions successfully`, func() {
				// Construct an instance of the RuleScope model
				ruleScopeModel := new(configurationgovernancev1.RuleScope)
				Expect(ruleScopeModel).ToNot(BeNil())
				ruleScopeModel.Note = core.StringPtr("My enterprise")
				ruleScopeModel.ScopeID = core.StringPtr("282cf433ac91493ba860480d92519990")
				ruleScopeModel.ScopeType = core.StringPtr("enterprise")
				Expect(ruleScopeModel.Note).To(Equal(core.StringPtr("My enterprise")))
				Expect(ruleScopeModel.ScopeID).To(Equal(core.StringPtr("282cf433ac91493ba860480d92519990")))
				Expect(ruleScopeModel.ScopeType).To(Equal(core.StringPtr("enterprise")))

				// Construct an instance of the UpdateRuleAttachmentOptions model
				ruleID := "testString"
				attachmentID := "testString"
				ifMatch := "testString"
				updateRuleAttachmentOptionsAccountID := "531fc3e28bfc43c5a2cea07786d93f5c"
				var updateRuleAttachmentOptionsIncludedScope *configurationgovernancev1.RuleScope = nil
				updateRuleAttachmentOptionsModel := configurationGovernanceService.NewUpdateRuleAttachmentOptions(ruleID, attachmentID, ifMatch, updateRuleAttachmentOptionsAccountID, updateRuleAttachmentOptionsIncludedScope)
				updateRuleAttachmentOptionsModel.SetRuleID("testString")
				updateRuleAttachmentOptionsModel.SetAttachmentID("testString")
				updateRuleAttachmentOptionsModel.SetIfMatch("testString")
				updateRuleAttachmentOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleAttachmentOptionsModel.SetIncludedScope(ruleScopeModel)
				updateRuleAttachmentOptionsModel.SetExcludedScopes([]configurationgovernancev1.RuleScope{*ruleScopeModel})
				updateRuleAttachmentOptionsModel.SetTransactionID("testString")
				updateRuleAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRuleAttachmentOptionsModel).ToNot(BeNil())
				Expect(updateRuleAttachmentOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleAttachmentOptionsModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleAttachmentOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(updateRuleAttachmentOptionsModel.IncludedScope).To(Equal(ruleScopeModel))
				Expect(updateRuleAttachmentOptionsModel.ExcludedScopes).To(Equal([]configurationgovernancev1.RuleScope{*ruleScopeModel}))
				Expect(updateRuleAttachmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRuleOptions successfully`, func() {
				// Construct an instance of the TargetResourceAdditionalTargetAttributesItem model
				targetResourceAdditionalTargetAttributesItemModel := new(configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem)
				Expect(targetResourceAdditionalTargetAttributesItemModel).ToNot(BeNil())
				targetResourceAdditionalTargetAttributesItemModel.Name = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Value = core.StringPtr("testString")
				targetResourceAdditionalTargetAttributesItemModel.Operator = core.StringPtr("string_equals")
				Expect(targetResourceAdditionalTargetAttributesItemModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceAdditionalTargetAttributesItemModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceAdditionalTargetAttributesItemModel.Operator).To(Equal(core.StringPtr("string_equals")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(configurationgovernancev1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ServiceName = core.StringPtr("iam-groups")
				targetResourceModel.ResourceKind = core.StringPtr("service")
				targetResourceModel.AdditionalTargetAttributes = []configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}
				Expect(targetResourceModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(targetResourceModel.ResourceKind).To(Equal(core.StringPtr("service")))
				Expect(targetResourceModel.AdditionalTargetAttributes).To(Equal([]configurationgovernancev1.TargetResourceAdditionalTargetAttributesItem{*targetResourceAdditionalTargetAttributesItemModel}))

				// Construct an instance of the RuleRequiredConfigSingleProperty model
				ruleRequiredConfigModel := new(configurationgovernancev1.RuleRequiredConfigSingleProperty)
				Expect(ruleRequiredConfigModel).ToNot(BeNil())
				ruleRequiredConfigModel.Description = core.StringPtr("testString")
				ruleRequiredConfigModel.Property = core.StringPtr("public_access_enabled")
				ruleRequiredConfigModel.Operator = core.StringPtr("is_false")
				ruleRequiredConfigModel.Value = core.StringPtr("testString")
				Expect(ruleRequiredConfigModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(ruleRequiredConfigModel.Property).To(Equal(core.StringPtr("public_access_enabled")))
				Expect(ruleRequiredConfigModel.Operator).To(Equal(core.StringPtr("is_false")))
				Expect(ruleRequiredConfigModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the EnforcementAction model
				enforcementActionModel := new(configurationgovernancev1.EnforcementAction)
				Expect(enforcementActionModel).ToNot(BeNil())
				enforcementActionModel.Action = core.StringPtr("audit_log")
				Expect(enforcementActionModel.Action).To(Equal(core.StringPtr("audit_log")))

				// Construct an instance of the UpdateRuleOptions model
				ruleID := "testString"
				ifMatch := "testString"
				updateRuleOptionsName := "Disable public access"
				updateRuleOptionsDescription := "Ensure that public access to account resources is disabled."
				var updateRuleOptionsTarget *configurationgovernancev1.TargetResource = nil
				var updateRuleOptionsRequiredConfig configurationgovernancev1.RuleRequiredConfigIntf = nil
				updateRuleOptionsEnforcementActions := []configurationgovernancev1.EnforcementAction{}
				updateRuleOptionsModel := configurationGovernanceService.NewUpdateRuleOptions(ruleID, ifMatch, updateRuleOptionsName, updateRuleOptionsDescription, updateRuleOptionsTarget, updateRuleOptionsRequiredConfig, updateRuleOptionsEnforcementActions)
				updateRuleOptionsModel.SetRuleID("testString")
				updateRuleOptionsModel.SetIfMatch("testString")
				updateRuleOptionsModel.SetName("Disable public access")
				updateRuleOptionsModel.SetDescription("Ensure that public access to account resources is disabled.")
				updateRuleOptionsModel.SetTarget(targetResourceModel)
				updateRuleOptionsModel.SetRequiredConfig(ruleRequiredConfigModel)
				updateRuleOptionsModel.SetEnforcementActions([]configurationgovernancev1.EnforcementAction{*enforcementActionModel})
				updateRuleOptionsModel.SetAccountID("531fc3e28bfc43c5a2cea07786d93f5c")
				updateRuleOptionsModel.SetRuleType("user_defined")
				updateRuleOptionsModel.SetLabels([]string{"SOC2", "ITCS300"})
				updateRuleOptionsModel.SetTransactionID("testString")
				updateRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRuleOptionsModel).ToNot(BeNil())
				Expect(updateRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.Name).To(Equal(core.StringPtr("Disable public access")))
				Expect(updateRuleOptionsModel.Description).To(Equal(core.StringPtr("Ensure that public access to account resources is disabled.")))
				Expect(updateRuleOptionsModel.Target).To(Equal(targetResourceModel))
				Expect(updateRuleOptionsModel.RequiredConfig).To(Equal(ruleRequiredConfigModel))
				Expect(updateRuleOptionsModel.EnforcementActions).To(Equal([]configurationgovernancev1.EnforcementAction{*enforcementActionModel}))
				Expect(updateRuleOptionsModel.AccountID).To(Equal(core.StringPtr("531fc3e28bfc43c5a2cea07786d93f5c")))
				Expect(updateRuleOptionsModel.RuleType).To(Equal(core.StringPtr("user_defined")))
				Expect(updateRuleOptionsModel.Labels).To(Equal([]string{"SOC2", "ITCS300"}))
				Expect(updateRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleConditionAndLvl2 successfully`, func() {
				and := []configurationgovernancev1.RuleSingleProperty{}
				model, err := configurationGovernanceService.NewRuleConditionAndLvl2(and)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleConditionOrLvl2 successfully`, func() {
				or := []configurationgovernancev1.RuleSingleProperty{}
				model, err := configurationGovernanceService.NewRuleConditionOrLvl2(or)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleConditionSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleConditionSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigSingleProperty successfully`, func() {
				property := "public_access_enabled"
				operator := "is_true"
				model, err := configurationGovernanceService.NewRuleRequiredConfigSingleProperty(property, operator)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigMultiplePropertiesConditionAnd successfully`, func() {
				and := []configurationgovernancev1.RuleConditionIntf{}
				model, err := configurationGovernanceService.NewRuleRequiredConfigMultiplePropertiesConditionAnd(and)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleRequiredConfigMultiplePropertiesConditionOr successfully`, func() {
				or := []configurationgovernancev1.RuleConditionIntf{}
				model, err := configurationGovernanceService.NewRuleRequiredConfigMultiplePropertiesConditionOr(or)
				Expect(model).ToNot(BeNil())
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

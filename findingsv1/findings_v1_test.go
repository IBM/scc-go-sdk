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

package findingsv1_test

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
	"github.com/ibm/scc-go-sdk/findingsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`FindingsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(findingsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(findingsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
				URL: "https://findingsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(findingsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FINDINGS_URL":       "https://findingsv1/api",
				"FINDINGS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				findingsService, serviceErr := findingsv1.NewFindingsV1UsingExternalConfig(&findingsv1.FindingsV1Options{})
				Expect(findingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := findingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != findingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(findingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(findingsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				findingsService, serviceErr := findingsv1.NewFindingsV1UsingExternalConfig(&findingsv1.FindingsV1Options{
					URL: "https://testService/api",
				})
				Expect(findingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(findingsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := findingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != findingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(findingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(findingsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				findingsService, serviceErr := findingsv1.NewFindingsV1UsingExternalConfig(&findingsv1.FindingsV1Options{})
				err := findingsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(findingsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := findingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != findingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(findingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(findingsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FINDINGS_URL":       "https://findingsv1/api",
				"FINDINGS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			findingsService, serviceErr := findingsv1.NewFindingsV1UsingExternalConfig(&findingsv1.FindingsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(findingsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"FINDINGS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			findingsService, serviceErr := findingsv1.NewFindingsV1UsingExternalConfig(&findingsv1.FindingsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(findingsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = findingsv1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us-south.secadvisor.cloud.ibm.com/findings"))
			Expect(err).To(BeNil())

			url, err = findingsv1.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us-south.secadvisor.cloud.ibm.com/findings"))
			Expect(err).To(BeNil())

			url, err = findingsv1.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://eu-gb.secadvisor.cloud.ibm.com/findings"))
			Expect(err).To(BeNil())

			url, err = findingsv1.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu.compliance.cloud.ibm.com/si/findings"))
			Expect(err).To(BeNil())

			url, err = findingsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`PostGraph(postGraphOptions *PostGraphOptions)`, func() {
		postGraphPath := "/v1/testString/graph"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postGraphPath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostGraph successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := findingsService.PostGraph(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PostGraphOptions model
				postGraphOptionsModel := new(findingsv1.PostGraphOptions)
				postGraphOptionsModel.AccountID = core.StringPtr("testString")
				postGraphOptionsModel.Body = CreateMockReader("This is a mock file.")
				postGraphOptionsModel.ContentType = core.StringPtr("application/json")
				postGraphOptionsModel.TransactionID = core.StringPtr("testString")
				postGraphOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = findingsService.PostGraph(postGraphOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PostGraph with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the PostGraphOptions model
				postGraphOptionsModel := new(findingsv1.PostGraphOptions)
				postGraphOptionsModel.AccountID = core.StringPtr("testString")
				postGraphOptionsModel.Body = CreateMockReader("This is a mock file.")
				postGraphOptionsModel.ContentType = core.StringPtr("application/json")
				postGraphOptionsModel.TransactionID = core.StringPtr("testString")
				postGraphOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := findingsService.PostGraph(postGraphOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PostGraphOptions model with no property values
				postGraphOptionsModelNew := new(findingsv1.PostGraphOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = findingsService.PostGraph(postGraphOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateNote(createNoteOptions *CreateNoteOptions) - Operation response error`, func() {
		createNotePath := "/v1/testString/providers/testString/notes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNotePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateNote with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the CreateNoteOptions model
				createNoteOptionsModel := new(findingsv1.CreateNoteOptions)
				createNoteOptionsModel.AccountID = core.StringPtr("testString")
				createNoteOptionsModel.ProviderID = core.StringPtr("testString")
				createNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				createNoteOptionsModel.LongDescription = core.StringPtr("testString")
				createNoteOptionsModel.Kind = core.StringPtr("FINDING")
				createNoteOptionsModel.ID = core.StringPtr("testString")
				createNoteOptionsModel.ReportedBy = reporterModel
				createNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				createNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				createNoteOptionsModel.Shared = core.BoolPtr(true)
				createNoteOptionsModel.Finding = findingTypeModel
				createNoteOptionsModel.Kpi = kpiTypeModel
				createNoteOptionsModel.Card = cardModel
				createNoteOptionsModel.Section = sectionModel
				createNoteOptionsModel.TransactionID = core.StringPtr("testString")
				createNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.CreateNote(createNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.CreateNote(createNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {
		createNotePath := "/v1/testString/providers/testString/notes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createNotePath))
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
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke CreateNote successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the CreateNoteOptions model
				createNoteOptionsModel := new(findingsv1.CreateNoteOptions)
				createNoteOptionsModel.AccountID = core.StringPtr("testString")
				createNoteOptionsModel.ProviderID = core.StringPtr("testString")
				createNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				createNoteOptionsModel.LongDescription = core.StringPtr("testString")
				createNoteOptionsModel.Kind = core.StringPtr("FINDING")
				createNoteOptionsModel.ID = core.StringPtr("testString")
				createNoteOptionsModel.ReportedBy = reporterModel
				createNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				createNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				createNoteOptionsModel.Shared = core.BoolPtr(true)
				createNoteOptionsModel.Finding = findingTypeModel
				createNoteOptionsModel.Kpi = kpiTypeModel
				createNoteOptionsModel.Card = cardModel
				createNoteOptionsModel.Section = sectionModel
				createNoteOptionsModel.TransactionID = core.StringPtr("testString")
				createNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.CreateNoteWithContext(ctx, createNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.CreateNote(createNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.CreateNoteWithContext(ctx, createNoteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createNotePath))
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
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke CreateNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.CreateNote(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the CreateNoteOptions model
				createNoteOptionsModel := new(findingsv1.CreateNoteOptions)
				createNoteOptionsModel.AccountID = core.StringPtr("testString")
				createNoteOptionsModel.ProviderID = core.StringPtr("testString")
				createNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				createNoteOptionsModel.LongDescription = core.StringPtr("testString")
				createNoteOptionsModel.Kind = core.StringPtr("FINDING")
				createNoteOptionsModel.ID = core.StringPtr("testString")
				createNoteOptionsModel.ReportedBy = reporterModel
				createNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				createNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				createNoteOptionsModel.Shared = core.BoolPtr(true)
				createNoteOptionsModel.Finding = findingTypeModel
				createNoteOptionsModel.Kpi = kpiTypeModel
				createNoteOptionsModel.Card = cardModel
				createNoteOptionsModel.Section = sectionModel
				createNoteOptionsModel.TransactionID = core.StringPtr("testString")
				createNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.CreateNote(createNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateNote with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the CreateNoteOptions model
				createNoteOptionsModel := new(findingsv1.CreateNoteOptions)
				createNoteOptionsModel.AccountID = core.StringPtr("testString")
				createNoteOptionsModel.ProviderID = core.StringPtr("testString")
				createNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				createNoteOptionsModel.LongDescription = core.StringPtr("testString")
				createNoteOptionsModel.Kind = core.StringPtr("FINDING")
				createNoteOptionsModel.ID = core.StringPtr("testString")
				createNoteOptionsModel.ReportedBy = reporterModel
				createNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				createNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				createNoteOptionsModel.Shared = core.BoolPtr(true)
				createNoteOptionsModel.Finding = findingTypeModel
				createNoteOptionsModel.Kpi = kpiTypeModel
				createNoteOptionsModel.Card = cardModel
				createNoteOptionsModel.Section = sectionModel
				createNoteOptionsModel.TransactionID = core.StringPtr("testString")
				createNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.CreateNote(createNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateNoteOptions model with no property values
				createNoteOptionsModelNew := new(findingsv1.CreateNoteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.CreateNote(createNoteOptionsModelNew)
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
			It(`Invoke CreateNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the CreateNoteOptions model
				createNoteOptionsModel := new(findingsv1.CreateNoteOptions)
				createNoteOptionsModel.AccountID = core.StringPtr("testString")
				createNoteOptionsModel.ProviderID = core.StringPtr("testString")
				createNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				createNoteOptionsModel.LongDescription = core.StringPtr("testString")
				createNoteOptionsModel.Kind = core.StringPtr("FINDING")
				createNoteOptionsModel.ID = core.StringPtr("testString")
				createNoteOptionsModel.ReportedBy = reporterModel
				createNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				createNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				createNoteOptionsModel.Shared = core.BoolPtr(true)
				createNoteOptionsModel.Finding = findingTypeModel
				createNoteOptionsModel.Kpi = kpiTypeModel
				createNoteOptionsModel.Card = cardModel
				createNoteOptionsModel.Section = sectionModel
				createNoteOptionsModel.TransactionID = core.StringPtr("testString")
				createNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.CreateNote(createNoteOptionsModel)
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
	Describe(`ListNotes(listNotesOptions *ListNotesOptions) - Operation response error`, func() {
		listNotesPath := "/v1/testString/providers/testString/notes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNotes with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNotesOptions model
				listNotesOptionsModel := new(findingsv1.ListNotesOptions)
				listNotesOptionsModel.AccountID = core.StringPtr("testString")
				listNotesOptionsModel.ProviderID = core.StringPtr("testString")
				listNotesOptionsModel.TransactionID = core.StringPtr("testString")
				listNotesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNotesOptionsModel.PageToken = core.StringPtr("testString")
				listNotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.ListNotes(listNotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.ListNotes(listNotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNotes(listNotesOptions *ListNotesOptions)`, func() {
		listNotesPath := "/v1/testString/providers/testString/notes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notes": [{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListNotes successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the ListNotesOptions model
				listNotesOptionsModel := new(findingsv1.ListNotesOptions)
				listNotesOptionsModel.AccountID = core.StringPtr("testString")
				listNotesOptionsModel.ProviderID = core.StringPtr("testString")
				listNotesOptionsModel.TransactionID = core.StringPtr("testString")
				listNotesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNotesOptionsModel.PageToken = core.StringPtr("testString")
				listNotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.ListNotesWithContext(ctx, listNotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.ListNotes(listNotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.ListNotesWithContext(ctx, listNotesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listNotesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notes": [{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListNotes successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.ListNotes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNotesOptions model
				listNotesOptionsModel := new(findingsv1.ListNotesOptions)
				listNotesOptionsModel.AccountID = core.StringPtr("testString")
				listNotesOptionsModel.ProviderID = core.StringPtr("testString")
				listNotesOptionsModel.TransactionID = core.StringPtr("testString")
				listNotesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNotesOptionsModel.PageToken = core.StringPtr("testString")
				listNotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.ListNotes(listNotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNotes with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNotesOptions model
				listNotesOptionsModel := new(findingsv1.ListNotesOptions)
				listNotesOptionsModel.AccountID = core.StringPtr("testString")
				listNotesOptionsModel.ProviderID = core.StringPtr("testString")
				listNotesOptionsModel.TransactionID = core.StringPtr("testString")
				listNotesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNotesOptionsModel.PageToken = core.StringPtr("testString")
				listNotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.ListNotes(listNotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListNotesOptions model with no property values
				listNotesOptionsModelNew := new(findingsv1.ListNotesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.ListNotes(listNotesOptionsModelNew)
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
			It(`Invoke ListNotes successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNotesOptions model
				listNotesOptionsModel := new(findingsv1.ListNotesOptions)
				listNotesOptionsModel.AccountID = core.StringPtr("testString")
				listNotesOptionsModel.ProviderID = core.StringPtr("testString")
				listNotesOptionsModel.TransactionID = core.StringPtr("testString")
				listNotesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNotesOptionsModel.PageToken = core.StringPtr("testString")
				listNotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.ListNotes(listNotesOptionsModel)
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
	Describe(`GetNote(getNoteOptions *GetNoteOptions) - Operation response error`, func() {
		getNotePath := "/v1/testString/providers/testString/notes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNotePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetNote with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetNoteOptions model
				getNoteOptionsModel := new(findingsv1.GetNoteOptions)
				getNoteOptionsModel.AccountID = core.StringPtr("testString")
				getNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getNoteOptionsModel.NoteID = core.StringPtr("testString")
				getNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.GetNote(getNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.GetNote(getNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetNote(getNoteOptions *GetNoteOptions)`, func() {
		getNotePath := "/v1/testString/providers/testString/notes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNotePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke GetNote successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the GetNoteOptions model
				getNoteOptionsModel := new(findingsv1.GetNoteOptions)
				getNoteOptionsModel.AccountID = core.StringPtr("testString")
				getNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getNoteOptionsModel.NoteID = core.StringPtr("testString")
				getNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.GetNoteWithContext(ctx, getNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.GetNote(getNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.GetNoteWithContext(ctx, getNoteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getNotePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke GetNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.GetNote(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNoteOptions model
				getNoteOptionsModel := new(findingsv1.GetNoteOptions)
				getNoteOptionsModel.AccountID = core.StringPtr("testString")
				getNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getNoteOptionsModel.NoteID = core.StringPtr("testString")
				getNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.GetNote(getNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetNote with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetNoteOptions model
				getNoteOptionsModel := new(findingsv1.GetNoteOptions)
				getNoteOptionsModel.AccountID = core.StringPtr("testString")
				getNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getNoteOptionsModel.NoteID = core.StringPtr("testString")
				getNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.GetNote(getNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetNoteOptions model with no property values
				getNoteOptionsModelNew := new(findingsv1.GetNoteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.GetNote(getNoteOptionsModelNew)
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
			It(`Invoke GetNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetNoteOptions model
				getNoteOptionsModel := new(findingsv1.GetNoteOptions)
				getNoteOptionsModel.AccountID = core.StringPtr("testString")
				getNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getNoteOptionsModel.NoteID = core.StringPtr("testString")
				getNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.GetNote(getNoteOptionsModel)
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
	Describe(`UpdateNote(updateNoteOptions *UpdateNoteOptions) - Operation response error`, func() {
		updateNotePath := "/v1/testString/providers/testString/notes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateNotePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateNote with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the UpdateNoteOptions model
				updateNoteOptionsModel := new(findingsv1.UpdateNoteOptions)
				updateNoteOptionsModel.AccountID = core.StringPtr("testString")
				updateNoteOptionsModel.ProviderID = core.StringPtr("testString")
				updateNoteOptionsModel.NoteID = core.StringPtr("testString")
				updateNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				updateNoteOptionsModel.LongDescription = core.StringPtr("testString")
				updateNoteOptionsModel.Kind = core.StringPtr("FINDING")
				updateNoteOptionsModel.ID = core.StringPtr("testString")
				updateNoteOptionsModel.ReportedBy = reporterModel
				updateNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				updateNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				updateNoteOptionsModel.Shared = core.BoolPtr(true)
				updateNoteOptionsModel.Finding = findingTypeModel
				updateNoteOptionsModel.Kpi = kpiTypeModel
				updateNoteOptionsModel.Card = cardModel
				updateNoteOptionsModel.Section = sectionModel
				updateNoteOptionsModel.TransactionID = core.StringPtr("testString")
				updateNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.UpdateNote(updateNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.UpdateNote(updateNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {
		updateNotePath := "/v1/testString/providers/testString/notes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateNotePath))
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
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke UpdateNote successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the UpdateNoteOptions model
				updateNoteOptionsModel := new(findingsv1.UpdateNoteOptions)
				updateNoteOptionsModel.AccountID = core.StringPtr("testString")
				updateNoteOptionsModel.ProviderID = core.StringPtr("testString")
				updateNoteOptionsModel.NoteID = core.StringPtr("testString")
				updateNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				updateNoteOptionsModel.LongDescription = core.StringPtr("testString")
				updateNoteOptionsModel.Kind = core.StringPtr("FINDING")
				updateNoteOptionsModel.ID = core.StringPtr("testString")
				updateNoteOptionsModel.ReportedBy = reporterModel
				updateNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				updateNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				updateNoteOptionsModel.Shared = core.BoolPtr(true)
				updateNoteOptionsModel.Finding = findingTypeModel
				updateNoteOptionsModel.Kpi = kpiTypeModel
				updateNoteOptionsModel.Card = cardModel
				updateNoteOptionsModel.Section = sectionModel
				updateNoteOptionsModel.TransactionID = core.StringPtr("testString")
				updateNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.UpdateNoteWithContext(ctx, updateNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.UpdateNote(updateNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.UpdateNoteWithContext(ctx, updateNoteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateNotePath))
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
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke UpdateNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.UpdateNote(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the UpdateNoteOptions model
				updateNoteOptionsModel := new(findingsv1.UpdateNoteOptions)
				updateNoteOptionsModel.AccountID = core.StringPtr("testString")
				updateNoteOptionsModel.ProviderID = core.StringPtr("testString")
				updateNoteOptionsModel.NoteID = core.StringPtr("testString")
				updateNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				updateNoteOptionsModel.LongDescription = core.StringPtr("testString")
				updateNoteOptionsModel.Kind = core.StringPtr("FINDING")
				updateNoteOptionsModel.ID = core.StringPtr("testString")
				updateNoteOptionsModel.ReportedBy = reporterModel
				updateNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				updateNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				updateNoteOptionsModel.Shared = core.BoolPtr(true)
				updateNoteOptionsModel.Finding = findingTypeModel
				updateNoteOptionsModel.Kpi = kpiTypeModel
				updateNoteOptionsModel.Card = cardModel
				updateNoteOptionsModel.Section = sectionModel
				updateNoteOptionsModel.TransactionID = core.StringPtr("testString")
				updateNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.UpdateNote(updateNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateNote with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the UpdateNoteOptions model
				updateNoteOptionsModel := new(findingsv1.UpdateNoteOptions)
				updateNoteOptionsModel.AccountID = core.StringPtr("testString")
				updateNoteOptionsModel.ProviderID = core.StringPtr("testString")
				updateNoteOptionsModel.NoteID = core.StringPtr("testString")
				updateNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				updateNoteOptionsModel.LongDescription = core.StringPtr("testString")
				updateNoteOptionsModel.Kind = core.StringPtr("FINDING")
				updateNoteOptionsModel.ID = core.StringPtr("testString")
				updateNoteOptionsModel.ReportedBy = reporterModel
				updateNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				updateNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				updateNoteOptionsModel.Shared = core.BoolPtr(true)
				updateNoteOptionsModel.Finding = findingTypeModel
				updateNoteOptionsModel.Kpi = kpiTypeModel
				updateNoteOptionsModel.Card = cardModel
				updateNoteOptionsModel.Section = sectionModel
				updateNoteOptionsModel.TransactionID = core.StringPtr("testString")
				updateNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.UpdateNote(updateNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateNoteOptions model with no property values
				updateNoteOptionsModelNew := new(findingsv1.UpdateNoteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.UpdateNote(updateNoteOptionsModelNew)
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
			It(`Invoke UpdateNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				kpiTypeModel.AggregationType = core.StringPtr("SUM")

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")

				// Construct an instance of the UpdateNoteOptions model
				updateNoteOptionsModel := new(findingsv1.UpdateNoteOptions)
				updateNoteOptionsModel.AccountID = core.StringPtr("testString")
				updateNoteOptionsModel.ProviderID = core.StringPtr("testString")
				updateNoteOptionsModel.NoteID = core.StringPtr("testString")
				updateNoteOptionsModel.ShortDescription = core.StringPtr("testString")
				updateNoteOptionsModel.LongDescription = core.StringPtr("testString")
				updateNoteOptionsModel.Kind = core.StringPtr("FINDING")
				updateNoteOptionsModel.ID = core.StringPtr("testString")
				updateNoteOptionsModel.ReportedBy = reporterModel
				updateNoteOptionsModel.RelatedURL = []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}
				updateNoteOptionsModel.ExpirationTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				updateNoteOptionsModel.Shared = core.BoolPtr(true)
				updateNoteOptionsModel.Finding = findingTypeModel
				updateNoteOptionsModel.Kpi = kpiTypeModel
				updateNoteOptionsModel.Card = cardModel
				updateNoteOptionsModel.Section = sectionModel
				updateNoteOptionsModel.TransactionID = core.StringPtr("testString")
				updateNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.UpdateNote(updateNoteOptionsModel)
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
	Describe(`DeleteNote(deleteNoteOptions *DeleteNoteOptions)`, func() {
		deleteNotePath := "/v1/testString/providers/testString/notes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNotePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := findingsService.DeleteNote(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNoteOptions model
				deleteNoteOptionsModel := new(findingsv1.DeleteNoteOptions)
				deleteNoteOptionsModel.AccountID = core.StringPtr("testString")
				deleteNoteOptionsModel.ProviderID = core.StringPtr("testString")
				deleteNoteOptionsModel.NoteID = core.StringPtr("testString")
				deleteNoteOptionsModel.TransactionID = core.StringPtr("testString")
				deleteNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = findingsService.DeleteNote(deleteNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNote with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the DeleteNoteOptions model
				deleteNoteOptionsModel := new(findingsv1.DeleteNoteOptions)
				deleteNoteOptionsModel.AccountID = core.StringPtr("testString")
				deleteNoteOptionsModel.ProviderID = core.StringPtr("testString")
				deleteNoteOptionsModel.NoteID = core.StringPtr("testString")
				deleteNoteOptionsModel.TransactionID = core.StringPtr("testString")
				deleteNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := findingsService.DeleteNote(deleteNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteNoteOptions model with no property values
				deleteNoteOptionsModelNew := new(findingsv1.DeleteNoteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = findingsService.DeleteNote(deleteNoteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOccurrenceNote(getOccurrenceNoteOptions *GetOccurrenceNoteOptions) - Operation response error`, func() {
		getOccurrenceNotePath := "/v1/testString/providers/testString/occurrences/testString/note"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrenceNotePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOccurrenceNote with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceNoteOptions model
				getOccurrenceNoteOptionsModel := new(findingsv1.GetOccurrenceNoteOptions)
				getOccurrenceNoteOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOccurrenceNote(getOccurrenceNoteOptions *GetOccurrenceNoteOptions)`, func() {
		getOccurrenceNotePath := "/v1/testString/providers/testString/occurrences/testString/note"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrenceNotePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke GetOccurrenceNote successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the GetOccurrenceNoteOptions model
				getOccurrenceNoteOptionsModel := new(findingsv1.GetOccurrenceNoteOptions)
				getOccurrenceNoteOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.GetOccurrenceNoteWithContext(ctx, getOccurrenceNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.GetOccurrenceNoteWithContext(ctx, getOccurrenceNoteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrenceNotePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"short_description": "ShortDescription", "long_description": "LongDescription", "kind": "FINDING", "related_url": [{"label": "Label", "url": "URL"}], "expiration_time": "2019-01-01T12:00:00.000Z", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "shared": true, "reported_by": {"id": "ID", "title": "Title", "url": "URL"}, "finding": {"severity": "LOW", "next_steps": [{"title": "Title", "url": "URL"}]}, "kpi": {"aggregation_type": "SUM"}, "card": {"section": "Section", "title": "Title", "subtitle": "Subtitle", "order": 1, "finding_note_names": ["FindingNoteNames"], "requires_configuration": false, "badge_text": "BadgeText", "badge_image": "BadgeImage", "elements": [{"text": "Text", "default_interval": "DefaultInterval", "kind": "TIME_SERIES", "default_time_range": "1d", "value_types": [{"kind": "FINDING_COUNT", "finding_note_names": ["FindingNoteNames"], "text": "Text"}]}]}, "section": {"title": "Title", "image": "Image"}}`)
				}))
			})
			It(`Invoke GetOccurrenceNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.GetOccurrenceNote(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOccurrenceNoteOptions model
				getOccurrenceNoteOptionsModel := new(findingsv1.GetOccurrenceNoteOptions)
				getOccurrenceNoteOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetOccurrenceNote with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceNoteOptions model
				getOccurrenceNoteOptionsModel := new(findingsv1.GetOccurrenceNoteOptions)
				getOccurrenceNoteOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOccurrenceNoteOptions model with no property values
				getOccurrenceNoteOptionsModelNew := new(findingsv1.GetOccurrenceNoteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModelNew)
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
			It(`Invoke GetOccurrenceNote successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceNoteOptions model
				getOccurrenceNoteOptionsModel := new(findingsv1.GetOccurrenceNoteOptions)
				getOccurrenceNoteOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceNoteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.GetOccurrenceNote(getOccurrenceNoteOptionsModel)
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
	Describe(`CreateOccurrence(createOccurrenceOptions *CreateOccurrenceOptions) - Operation response error`, func() {
		createOccurrencePath := "/v1/testString/providers/testString/occurrences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOccurrencePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Replace-If-Exists"]).ToNot(BeNil())
					Expect(req.Header["Replace-If-Exists"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOccurrence with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CreateOccurrenceOptions model
				createOccurrenceOptionsModel := new(findingsv1.CreateOccurrenceOptions)
				createOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				createOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				createOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				createOccurrenceOptionsModel.ID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				createOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				createOccurrenceOptionsModel.Context = contextModel
				createOccurrenceOptionsModel.Finding = findingModel
				createOccurrenceOptionsModel.Kpi = kpiModel
				createOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				createOccurrenceOptionsModel.ReplaceIfExists = core.BoolPtr(true)
				createOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				createOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.CreateOccurrence(createOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.CreateOccurrence(createOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOccurrence(createOccurrenceOptions *CreateOccurrenceOptions)`, func() {
		createOccurrencePath := "/v1/testString/providers/testString/occurrences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOccurrencePath))
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

					Expect(req.Header["Replace-If-Exists"]).ToNot(BeNil())
					Expect(req.Header["Replace-If-Exists"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CreateOccurrence successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CreateOccurrenceOptions model
				createOccurrenceOptionsModel := new(findingsv1.CreateOccurrenceOptions)
				createOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				createOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				createOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				createOccurrenceOptionsModel.ID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				createOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				createOccurrenceOptionsModel.Context = contextModel
				createOccurrenceOptionsModel.Finding = findingModel
				createOccurrenceOptionsModel.Kpi = kpiModel
				createOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				createOccurrenceOptionsModel.ReplaceIfExists = core.BoolPtr(true)
				createOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				createOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.CreateOccurrenceWithContext(ctx, createOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.CreateOccurrence(createOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.CreateOccurrenceWithContext(ctx, createOccurrenceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createOccurrencePath))
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

					Expect(req.Header["Replace-If-Exists"]).ToNot(BeNil())
					Expect(req.Header["Replace-If-Exists"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CreateOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.CreateOccurrence(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CreateOccurrenceOptions model
				createOccurrenceOptionsModel := new(findingsv1.CreateOccurrenceOptions)
				createOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				createOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				createOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				createOccurrenceOptionsModel.ID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				createOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				createOccurrenceOptionsModel.Context = contextModel
				createOccurrenceOptionsModel.Finding = findingModel
				createOccurrenceOptionsModel.Kpi = kpiModel
				createOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				createOccurrenceOptionsModel.ReplaceIfExists = core.BoolPtr(true)
				createOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				createOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.CreateOccurrence(createOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateOccurrence with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CreateOccurrenceOptions model
				createOccurrenceOptionsModel := new(findingsv1.CreateOccurrenceOptions)
				createOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				createOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				createOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				createOccurrenceOptionsModel.ID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				createOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				createOccurrenceOptionsModel.Context = contextModel
				createOccurrenceOptionsModel.Finding = findingModel
				createOccurrenceOptionsModel.Kpi = kpiModel
				createOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				createOccurrenceOptionsModel.ReplaceIfExists = core.BoolPtr(true)
				createOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				createOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.CreateOccurrence(createOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOccurrenceOptions model with no property values
				createOccurrenceOptionsModelNew := new(findingsv1.CreateOccurrenceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.CreateOccurrence(createOccurrenceOptionsModelNew)
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
			It(`Invoke CreateOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the CreateOccurrenceOptions model
				createOccurrenceOptionsModel := new(findingsv1.CreateOccurrenceOptions)
				createOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				createOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				createOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				createOccurrenceOptionsModel.ID = core.StringPtr("testString")
				createOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				createOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				createOccurrenceOptionsModel.Context = contextModel
				createOccurrenceOptionsModel.Finding = findingModel
				createOccurrenceOptionsModel.Kpi = kpiModel
				createOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				createOccurrenceOptionsModel.ReplaceIfExists = core.BoolPtr(true)
				createOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				createOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.CreateOccurrence(createOccurrenceOptionsModel)
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
	Describe(`ListOccurrences(listOccurrencesOptions *ListOccurrencesOptions) - Operation response error`, func() {
		listOccurrencesPath := "/v1/testString/providers/testString/occurrences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOccurrences with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListOccurrencesOptions model
				listOccurrencesOptionsModel := new(findingsv1.ListOccurrencesOptions)
				listOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.ListOccurrences(listOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.ListOccurrences(listOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOccurrences(listOccurrencesOptions *ListOccurrencesOptions)`, func() {
		listOccurrencesPath := "/v1/testString/providers/testString/occurrences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListOccurrences successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the ListOccurrencesOptions model
				listOccurrencesOptionsModel := new(findingsv1.ListOccurrencesOptions)
				listOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.ListOccurrencesWithContext(ctx, listOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.ListOccurrences(listOccurrencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.ListOccurrencesWithContext(ctx, listOccurrencesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListOccurrences successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.ListOccurrences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOccurrencesOptions model
				listOccurrencesOptionsModel := new(findingsv1.ListOccurrencesOptions)
				listOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.ListOccurrences(listOccurrencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOccurrences with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListOccurrencesOptions model
				listOccurrencesOptionsModel := new(findingsv1.ListOccurrencesOptions)
				listOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.ListOccurrences(listOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOccurrencesOptions model with no property values
				listOccurrencesOptionsModelNew := new(findingsv1.ListOccurrencesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.ListOccurrences(listOccurrencesOptionsModelNew)
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
			It(`Invoke ListOccurrences successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListOccurrencesOptions model
				listOccurrencesOptionsModel := new(findingsv1.ListOccurrencesOptions)
				listOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.ListOccurrences(listOccurrencesOptionsModel)
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
	Describe(`ListNoteOccurrences(listNoteOccurrencesOptions *ListNoteOccurrencesOptions) - Operation response error`, func() {
		listNoteOccurrencesPath := "/v1/testString/providers/testString/notes/testString/occurrences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNoteOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNoteOccurrences with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNoteOccurrencesOptions model
				listNoteOccurrencesOptionsModel := new(findingsv1.ListNoteOccurrencesOptions)
				listNoteOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.NoteID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNoteOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNoteOccurrences(listNoteOccurrencesOptions *ListNoteOccurrencesOptions)`, func() {
		listNoteOccurrencesPath := "/v1/testString/providers/testString/notes/testString/occurrences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNoteOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListNoteOccurrences successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the ListNoteOccurrencesOptions model
				listNoteOccurrencesOptionsModel := new(findingsv1.ListNoteOccurrencesOptions)
				listNoteOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.NoteID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNoteOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.ListNoteOccurrencesWithContext(ctx, listNoteOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.ListNoteOccurrencesWithContext(ctx, listNoteOccurrencesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listNoteOccurrencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["page_size"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["page_token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke ListNoteOccurrences successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.ListNoteOccurrences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNoteOccurrencesOptions model
				listNoteOccurrencesOptionsModel := new(findingsv1.ListNoteOccurrencesOptions)
				listNoteOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.NoteID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNoteOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNoteOccurrences with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNoteOccurrencesOptions model
				listNoteOccurrencesOptionsModel := new(findingsv1.ListNoteOccurrencesOptions)
				listNoteOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.NoteID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNoteOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListNoteOccurrencesOptions model with no property values
				listNoteOccurrencesOptionsModelNew := new(findingsv1.ListNoteOccurrencesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModelNew)
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
			It(`Invoke ListNoteOccurrences successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListNoteOccurrencesOptions model
				listNoteOccurrencesOptionsModel := new(findingsv1.ListNoteOccurrencesOptions)
				listNoteOccurrencesOptionsModel.AccountID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.ProviderID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.NoteID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.TransactionID = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.PageSize = core.Int64Ptr(int64(2))
				listNoteOccurrencesOptionsModel.PageToken = core.StringPtr("testString")
				listNoteOccurrencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.ListNoteOccurrences(listNoteOccurrencesOptionsModel)
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
	Describe(`GetOccurrence(getOccurrenceOptions *GetOccurrenceOptions) - Operation response error`, func() {
		getOccurrencePath := "/v1/testString/providers/testString/occurrences/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrencePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOccurrence with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceOptions model
				getOccurrenceOptionsModel := new(findingsv1.GetOccurrenceOptions)
				getOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.GetOccurrence(getOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.GetOccurrence(getOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOccurrence(getOccurrenceOptions *GetOccurrenceOptions)`, func() {
		getOccurrencePath := "/v1/testString/providers/testString/occurrences/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrencePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke GetOccurrence successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the GetOccurrenceOptions model
				getOccurrenceOptionsModel := new(findingsv1.GetOccurrenceOptions)
				getOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.GetOccurrenceWithContext(ctx, getOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.GetOccurrence(getOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.GetOccurrenceWithContext(ctx, getOccurrenceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getOccurrencePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"occurrences": [{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}], "next_page_token": "NextPageToken"}`)
				}))
			})
			It(`Invoke GetOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.GetOccurrence(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOccurrenceOptions model
				getOccurrenceOptionsModel := new(findingsv1.GetOccurrenceOptions)
				getOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.GetOccurrence(getOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetOccurrence with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceOptions model
				getOccurrenceOptionsModel := new(findingsv1.GetOccurrenceOptions)
				getOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.GetOccurrence(getOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOccurrenceOptions model with no property values
				getOccurrenceOptionsModelNew := new(findingsv1.GetOccurrenceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.GetOccurrence(getOccurrenceOptionsModelNew)
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
			It(`Invoke GetOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the GetOccurrenceOptions model
				getOccurrenceOptionsModel := new(findingsv1.GetOccurrenceOptions)
				getOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				getOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				getOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				getOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				getOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.GetOccurrence(getOccurrenceOptionsModel)
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
	Describe(`UpdateOccurrence(updateOccurrenceOptions *UpdateOccurrenceOptions) - Operation response error`, func() {
		updateOccurrencePath := "/v1/testString/providers/testString/occurrences/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOccurrencePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOccurrence with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the UpdateOccurrenceOptions model
				updateOccurrenceOptionsModel := new(findingsv1.UpdateOccurrenceOptions)
				updateOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				updateOccurrenceOptionsModel.ID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Context = contextModel
				updateOccurrenceOptionsModel.Finding = findingModel
				updateOccurrenceOptionsModel.Kpi = kpiModel
				updateOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				updateOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOccurrence(updateOccurrenceOptions *UpdateOccurrenceOptions)`, func() {
		updateOccurrencePath := "/v1/testString/providers/testString/occurrences/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOccurrencePath))
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
					fmt.Fprintf(res, "%s", `{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateOccurrence successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the UpdateOccurrenceOptions model
				updateOccurrenceOptionsModel := new(findingsv1.UpdateOccurrenceOptions)
				updateOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				updateOccurrenceOptionsModel.ID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Context = contextModel
				updateOccurrenceOptionsModel.Finding = findingModel
				updateOccurrenceOptionsModel.Kpi = kpiModel
				updateOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				updateOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.UpdateOccurrenceWithContext(ctx, updateOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.UpdateOccurrenceWithContext(ctx, updateOccurrenceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateOccurrencePath))
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
					fmt.Fprintf(res, "%s", `{"resource_url": "ResourceURL", "note_name": "NoteName", "kind": "FINDING", "remediation": "Remediation", "create_time": "2019-01-01T12:00:00.000Z", "update_time": "2019-01-01T12:00:00.000Z", "id": "ID", "context": {"region": "Region", "resource_crn": "ResourceCRN", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_type": "ResourceType", "service_crn": "ServiceCRN", "service_name": "ServiceName", "environment_name": "EnvironmentName", "component_name": "ComponentName", "toolchain_id": "ToolchainID"}, "finding": {"severity": "LOW", "certainty": "LOW", "next_steps": [{"title": "Title", "url": "URL"}], "network_connection": {"direction": "Direction", "protocol": "Protocol", "client": {"address": "Address", "port": 4}, "server": {"address": "Address", "port": 4}}, "data_transferred": {"client_bytes": 11, "server_bytes": 11, "client_packets": 13, "server_packets": 13}}, "kpi": {"value": 5, "total": 5}, "reference_data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.UpdateOccurrence(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the UpdateOccurrenceOptions model
				updateOccurrenceOptionsModel := new(findingsv1.UpdateOccurrenceOptions)
				updateOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				updateOccurrenceOptionsModel.ID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Context = contextModel
				updateOccurrenceOptionsModel.Finding = findingModel
				updateOccurrenceOptionsModel.Kpi = kpiModel
				updateOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				updateOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateOccurrence with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the UpdateOccurrenceOptions model
				updateOccurrenceOptionsModel := new(findingsv1.UpdateOccurrenceOptions)
				updateOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				updateOccurrenceOptionsModel.ID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Context = contextModel
				updateOccurrenceOptionsModel.Finding = findingModel
				updateOccurrenceOptionsModel.Kpi = kpiModel
				updateOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				updateOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOccurrenceOptions model with no property values
				updateOccurrenceOptionsModelNew := new(findingsv1.UpdateOccurrenceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.UpdateOccurrence(updateOccurrenceOptionsModelNew)
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
			It(`Invoke UpdateOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))

				// Construct an instance of the UpdateOccurrenceOptions model
				updateOccurrenceOptionsModel := new(findingsv1.UpdateOccurrenceOptions)
				updateOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.NoteName = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Kind = core.StringPtr("FINDING")
				updateOccurrenceOptionsModel.ID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.ResourceURL = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Remediation = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Context = contextModel
				updateOccurrenceOptionsModel.Finding = findingModel
				updateOccurrenceOptionsModel.Kpi = kpiModel
				updateOccurrenceOptionsModel.ReferenceData = map[string]interface{}{"anyKey": "anyValue"}
				updateOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				updateOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.UpdateOccurrence(updateOccurrenceOptionsModel)
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
	Describe(`DeleteOccurrence(deleteOccurrenceOptions *DeleteOccurrenceOptions)`, func() {
		deleteOccurrencePath := "/v1/testString/providers/testString/occurrences/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteOccurrencePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteOccurrence successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := findingsService.DeleteOccurrence(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteOccurrenceOptions model
				deleteOccurrenceOptionsModel := new(findingsv1.DeleteOccurrenceOptions)
				deleteOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = findingsService.DeleteOccurrence(deleteOccurrenceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteOccurrence with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the DeleteOccurrenceOptions model
				deleteOccurrenceOptionsModel := new(findingsv1.DeleteOccurrenceOptions)
				deleteOccurrenceOptionsModel.AccountID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.ProviderID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.OccurrenceID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.TransactionID = core.StringPtr("testString")
				deleteOccurrenceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := findingsService.DeleteOccurrence(deleteOccurrenceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteOccurrenceOptions model with no property values
				deleteOccurrenceOptionsModelNew := new(findingsv1.DeleteOccurrenceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = findingsService.DeleteOccurrence(deleteOccurrenceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProviders(listProvidersOptions *ListProvidersOptions) - Operation response error`, func() {
		listProvidersPath := "/v1/testString/providers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProvidersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["skip"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["start_provider_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["end_provider_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProviders with error: Operation response processing error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListProvidersOptions model
				listProvidersOptionsModel := new(findingsv1.ListProvidersOptions)
				listProvidersOptionsModel.AccountID = core.StringPtr("testString")
				listProvidersOptionsModel.TransactionID = core.StringPtr("testString")
				listProvidersOptionsModel.Limit = core.Int64Ptr(int64(2))
				listProvidersOptionsModel.Skip = core.Int64Ptr(int64(38))
				listProvidersOptionsModel.StartProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.EndProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := findingsService.ListProviders(listProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				findingsService.EnableRetries(0, 0)
				result, response, operationErr = findingsService.ListProviders(listProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProviders(listProvidersOptions *ListProvidersOptions)`, func() {
		listProvidersPath := "/v1/testString/providers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProvidersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["skip"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["start_provider_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["end_provider_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"providers": [{"name": "Name", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListProviders successfully with retries`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())
				findingsService.EnableRetries(0, 0)

				// Construct an instance of the ListProvidersOptions model
				listProvidersOptionsModel := new(findingsv1.ListProvidersOptions)
				listProvidersOptionsModel.AccountID = core.StringPtr("testString")
				listProvidersOptionsModel.TransactionID = core.StringPtr("testString")
				listProvidersOptionsModel.Limit = core.Int64Ptr(int64(2))
				listProvidersOptionsModel.Skip = core.Int64Ptr(int64(38))
				listProvidersOptionsModel.StartProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.EndProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := findingsService.ListProvidersWithContext(ctx, listProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				findingsService.DisableRetries()
				result, response, operationErr := findingsService.ListProviders(listProvidersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = findingsService.ListProvidersWithContext(ctx, listProvidersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProvidersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(2))}))
					Expect(req.URL.Query()["skip"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["start_provider_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["end_provider_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"providers": [{"name": "Name", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListProviders successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := findingsService.ListProviders(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProvidersOptions model
				listProvidersOptionsModel := new(findingsv1.ListProvidersOptions)
				listProvidersOptionsModel.AccountID = core.StringPtr("testString")
				listProvidersOptionsModel.TransactionID = core.StringPtr("testString")
				listProvidersOptionsModel.Limit = core.Int64Ptr(int64(2))
				listProvidersOptionsModel.Skip = core.Int64Ptr(int64(38))
				listProvidersOptionsModel.StartProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.EndProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = findingsService.ListProviders(listProvidersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProviders with error: Operation validation and request error`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListProvidersOptions model
				listProvidersOptionsModel := new(findingsv1.ListProvidersOptions)
				listProvidersOptionsModel.AccountID = core.StringPtr("testString")
				listProvidersOptionsModel.TransactionID = core.StringPtr("testString")
				listProvidersOptionsModel.Limit = core.Int64Ptr(int64(2))
				listProvidersOptionsModel.Skip = core.Int64Ptr(int64(38))
				listProvidersOptionsModel.StartProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.EndProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := findingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := findingsService.ListProviders(listProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProvidersOptions model with no property values
				listProvidersOptionsModelNew := new(findingsv1.ListProvidersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = findingsService.ListProviders(listProvidersOptionsModelNew)
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
			It(`Invoke ListProviders successfully`, func() {
				findingsService, serviceErr := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(findingsService).ToNot(BeNil())

				// Construct an instance of the ListProvidersOptions model
				listProvidersOptionsModel := new(findingsv1.ListProvidersOptions)
				listProvidersOptionsModel.AccountID = core.StringPtr("testString")
				listProvidersOptionsModel.TransactionID = core.StringPtr("testString")
				listProvidersOptionsModel.Limit = core.Int64Ptr(int64(2))
				listProvidersOptionsModel.Skip = core.Int64Ptr(int64(38))
				listProvidersOptionsModel.StartProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.EndProviderID = core.StringPtr("testString")
				listProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := findingsService.ListProviders(listProvidersOptionsModel)
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
			findingsService, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
				URL:           "http://findingsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCard successfully`, func() {
				section := "testString"
				title := "testString"
				subtitle := "testString"
				findingNoteNames := []string{"testString"}
				elements := []findingsv1.CardElementIntf{}
				model, err := findingsService.NewCard(section, title, subtitle, findingNoteNames, elements)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateNoteOptions successfully`, func() {
				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				Expect(reporterModel).ToNot(BeNil())
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")
				Expect(reporterModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(reporterModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(reporterModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				Expect(apiNoteRelatedURLModel).ToNot(BeNil())
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")
				Expect(apiNoteRelatedURLModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(apiNoteRelatedURLModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				Expect(remediationStepModel).ToNot(BeNil())
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")
				Expect(remediationStepModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(remediationStepModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				Expect(findingTypeModel).ToNot(BeNil())
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				Expect(findingTypeModel.Severity).To(Equal(core.StringPtr("LOW")))
				Expect(findingTypeModel.NextSteps).To(Equal([]findingsv1.RemediationStep{*remediationStepModel}))

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				Expect(kpiTypeModel).ToNot(BeNil())
				kpiTypeModel.AggregationType = core.StringPtr("SUM")
				Expect(kpiTypeModel.AggregationType).To(Equal(core.StringPtr("SUM")))

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				Expect(valueTypeModel).ToNot(BeNil())
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")
				Expect(valueTypeModel.Kind).To(Equal(core.StringPtr("FINDING_COUNT")))
				Expect(valueTypeModel.FindingNoteNames).To(Equal([]string{"testString"}))
				Expect(valueTypeModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				Expect(cardElementModel).ToNot(BeNil())
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}
				Expect(cardElementModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(cardElementModel.DefaultInterval).To(Equal(core.StringPtr("testString")))
				Expect(cardElementModel.Kind).To(Equal(core.StringPtr("TIME_SERIES")))
				Expect(cardElementModel.DefaultTimeRange).To(Equal(core.StringPtr("1d")))
				Expect(cardElementModel.ValueTypes).To(Equal([]findingsv1.ValueTypeIntf{valueTypeModel}))

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				Expect(cardModel).ToNot(BeNil())
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}
				Expect(cardModel.Section).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Subtitle).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Order).To(Equal(core.Int64Ptr(int64(1))))
				Expect(cardModel.FindingNoteNames).To(Equal([]string{"testString"}))
				Expect(cardModel.RequiresConfiguration).To(Equal(core.BoolPtr(true)))
				Expect(cardModel.BadgeText).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.BadgeImage).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Elements).To(Equal([]findingsv1.CardElementIntf{cardElementModel}))

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				Expect(sectionModel).ToNot(BeNil())
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")
				Expect(sectionModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(sectionModel.Image).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateNoteOptions model
				accountID := "testString"
				providerID := "testString"
				createNoteOptionsShortDescription := "testString"
				createNoteOptionsLongDescription := "testString"
				createNoteOptionsKind := "FINDING"
				createNoteOptionsID := "testString"
				var createNoteOptionsReportedBy *findingsv1.Reporter = nil
				createNoteOptionsModel := findingsService.NewCreateNoteOptions(accountID, providerID, createNoteOptionsShortDescription, createNoteOptionsLongDescription, createNoteOptionsKind, createNoteOptionsID, createNoteOptionsReportedBy)
				createNoteOptionsModel.SetAccountID("testString")
				createNoteOptionsModel.SetProviderID("testString")
				createNoteOptionsModel.SetShortDescription("testString")
				createNoteOptionsModel.SetLongDescription("testString")
				createNoteOptionsModel.SetKind("FINDING")
				createNoteOptionsModel.SetID("testString")
				createNoteOptionsModel.SetReportedBy(reporterModel)
				createNoteOptionsModel.SetRelatedURL([]findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel})
				createNoteOptionsModel.SetExpirationTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				createNoteOptionsModel.SetShared(true)
				createNoteOptionsModel.SetFinding(findingTypeModel)
				createNoteOptionsModel.SetKpi(kpiTypeModel)
				createNoteOptionsModel.SetCard(cardModel)
				createNoteOptionsModel.SetSection(sectionModel)
				createNoteOptionsModel.SetTransactionID("testString")
				createNoteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createNoteOptionsModel).ToNot(BeNil())
				Expect(createNoteOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.Kind).To(Equal(core.StringPtr("FINDING")))
				Expect(createNoteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.ReportedBy).To(Equal(reporterModel))
				Expect(createNoteOptionsModel.RelatedURL).To(Equal([]findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}))
				Expect(createNoteOptionsModel.ExpirationTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(createNoteOptionsModel.Shared).To(Equal(core.BoolPtr(true)))
				Expect(createNoteOptionsModel.Finding).To(Equal(findingTypeModel))
				Expect(createNoteOptionsModel.Kpi).To(Equal(kpiTypeModel))
				Expect(createNoteOptionsModel.Card).To(Equal(cardModel))
				Expect(createNoteOptionsModel.Section).To(Equal(sectionModel))
				Expect(createNoteOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createNoteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOccurrenceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")
				Expect(contextModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceCRN).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ServiceCRN).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.EnvironmentName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ComponentName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ToolchainID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				Expect(remediationStepModel).ToNot(BeNil())
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")
				Expect(remediationStepModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(remediationStepModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				Expect(socketAddressModel).ToNot(BeNil())
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))
				Expect(socketAddressModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(socketAddressModel.Port).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				Expect(networkConnectionModel).ToNot(BeNil())
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel
				Expect(networkConnectionModel.Direction).To(Equal(core.StringPtr("testString")))
				Expect(networkConnectionModel.Protocol).To(Equal(core.StringPtr("testString")))
				Expect(networkConnectionModel.Client).To(Equal(socketAddressModel))
				Expect(networkConnectionModel.Server).To(Equal(socketAddressModel))

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				Expect(dataTransferredModel).ToNot(BeNil())
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))
				Expect(dataTransferredModel.ClientBytes).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ServerBytes).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ClientPackets).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ServerPackets).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				Expect(findingModel).ToNot(BeNil())
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel
				Expect(findingModel.Severity).To(Equal(core.StringPtr("LOW")))
				Expect(findingModel.Certainty).To(Equal(core.StringPtr("LOW")))
				Expect(findingModel.NextSteps).To(Equal([]findingsv1.RemediationStep{*remediationStepModel}))
				Expect(findingModel.NetworkConnection).To(Equal(networkConnectionModel))
				Expect(findingModel.DataTransferred).To(Equal(dataTransferredModel))

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				Expect(kpiModel).ToNot(BeNil())
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))
				Expect(kpiModel.Value).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(kpiModel.Total).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the CreateOccurrenceOptions model
				accountID := "testString"
				providerID := "testString"
				createOccurrenceOptionsNoteName := "testString"
				createOccurrenceOptionsKind := "FINDING"
				createOccurrenceOptionsID := "testString"
				createOccurrenceOptionsModel := findingsService.NewCreateOccurrenceOptions(accountID, providerID, createOccurrenceOptionsNoteName, createOccurrenceOptionsKind, createOccurrenceOptionsID)
				createOccurrenceOptionsModel.SetAccountID("testString")
				createOccurrenceOptionsModel.SetProviderID("testString")
				createOccurrenceOptionsModel.SetNoteName("testString")
				createOccurrenceOptionsModel.SetKind("FINDING")
				createOccurrenceOptionsModel.SetID("testString")
				createOccurrenceOptionsModel.SetResourceURL("testString")
				createOccurrenceOptionsModel.SetRemediation("testString")
				createOccurrenceOptionsModel.SetContext(contextModel)
				createOccurrenceOptionsModel.SetFinding(findingModel)
				createOccurrenceOptionsModel.SetKpi(kpiModel)
				createOccurrenceOptionsModel.SetReferenceData(map[string]interface{}{"anyKey": "anyValue"})
				createOccurrenceOptionsModel.SetReplaceIfExists(true)
				createOccurrenceOptionsModel.SetTransactionID("testString")
				createOccurrenceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOccurrenceOptionsModel).ToNot(BeNil())
				Expect(createOccurrenceOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.NoteName).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.Kind).To(Equal(core.StringPtr("FINDING")))
				Expect(createOccurrenceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.ResourceURL).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.Remediation).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.Context).To(Equal(contextModel))
				Expect(createOccurrenceOptionsModel.Finding).To(Equal(findingModel))
				Expect(createOccurrenceOptionsModel.Kpi).To(Equal(kpiModel))
				Expect(createOccurrenceOptionsModel.ReferenceData).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createOccurrenceOptionsModel.ReplaceIfExists).To(Equal(core.BoolPtr(true)))
				Expect(createOccurrenceOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createOccurrenceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteNoteOptions successfully`, func() {
				// Construct an instance of the DeleteNoteOptions model
				accountID := "testString"
				providerID := "testString"
				noteID := "testString"
				deleteNoteOptionsModel := findingsService.NewDeleteNoteOptions(accountID, providerID, noteID)
				deleteNoteOptionsModel.SetAccountID("testString")
				deleteNoteOptionsModel.SetProviderID("testString")
				deleteNoteOptionsModel.SetNoteID("testString")
				deleteNoteOptionsModel.SetTransactionID("testString")
				deleteNoteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNoteOptionsModel).ToNot(BeNil())
				Expect(deleteNoteOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNoteOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNoteOptionsModel.NoteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNoteOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteNoteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteOccurrenceOptions successfully`, func() {
				// Construct an instance of the DeleteOccurrenceOptions model
				accountID := "testString"
				providerID := "testString"
				occurrenceID := "testString"
				deleteOccurrenceOptionsModel := findingsService.NewDeleteOccurrenceOptions(accountID, providerID, occurrenceID)
				deleteOccurrenceOptionsModel.SetAccountID("testString")
				deleteOccurrenceOptionsModel.SetProviderID("testString")
				deleteOccurrenceOptionsModel.SetOccurrenceID("testString")
				deleteOccurrenceOptionsModel.SetTransactionID("testString")
				deleteOccurrenceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteOccurrenceOptionsModel).ToNot(BeNil())
				Expect(deleteOccurrenceOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOccurrenceOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOccurrenceOptionsModel.OccurrenceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOccurrenceOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOccurrenceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFindingType successfully`, func() {
				severity := "LOW"
				model, err := findingsService.NewFindingType(severity)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetNoteOptions successfully`, func() {
				// Construct an instance of the GetNoteOptions model
				accountID := "testString"
				providerID := "testString"
				noteID := "testString"
				getNoteOptionsModel := findingsService.NewGetNoteOptions(accountID, providerID, noteID)
				getNoteOptionsModel.SetAccountID("testString")
				getNoteOptionsModel.SetProviderID("testString")
				getNoteOptionsModel.SetNoteID("testString")
				getNoteOptionsModel.SetTransactionID("testString")
				getNoteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getNoteOptionsModel).ToNot(BeNil())
				Expect(getNoteOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getNoteOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(getNoteOptionsModel.NoteID).To(Equal(core.StringPtr("testString")))
				Expect(getNoteOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getNoteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOccurrenceNoteOptions successfully`, func() {
				// Construct an instance of the GetOccurrenceNoteOptions model
				accountID := "testString"
				providerID := "testString"
				occurrenceID := "testString"
				getOccurrenceNoteOptionsModel := findingsService.NewGetOccurrenceNoteOptions(accountID, providerID, occurrenceID)
				getOccurrenceNoteOptionsModel.SetAccountID("testString")
				getOccurrenceNoteOptionsModel.SetProviderID("testString")
				getOccurrenceNoteOptionsModel.SetOccurrenceID("testString")
				getOccurrenceNoteOptionsModel.SetTransactionID("testString")
				getOccurrenceNoteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOccurrenceNoteOptionsModel).ToNot(BeNil())
				Expect(getOccurrenceNoteOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceNoteOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceNoteOptionsModel.OccurrenceID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceNoteOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceNoteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOccurrenceOptions successfully`, func() {
				// Construct an instance of the GetOccurrenceOptions model
				accountID := "testString"
				providerID := "testString"
				occurrenceID := "testString"
				getOccurrenceOptionsModel := findingsService.NewGetOccurrenceOptions(accountID, providerID, occurrenceID)
				getOccurrenceOptionsModel.SetAccountID("testString")
				getOccurrenceOptionsModel.SetProviderID("testString")
				getOccurrenceOptionsModel.SetOccurrenceID("testString")
				getOccurrenceOptionsModel.SetTransactionID("testString")
				getOccurrenceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOccurrenceOptionsModel).ToNot(BeNil())
				Expect(getOccurrenceOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceOptionsModel.OccurrenceID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getOccurrenceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewKpi successfully`, func() {
				value := float64(72.5)
				model, err := findingsService.NewKpi(value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewKpiType successfully`, func() {
				aggregationType := "SUM"
				model, err := findingsService.NewKpiType(aggregationType)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListNoteOccurrencesOptions successfully`, func() {
				// Construct an instance of the ListNoteOccurrencesOptions model
				accountID := "testString"
				providerID := "testString"
				noteID := "testString"
				listNoteOccurrencesOptionsModel := findingsService.NewListNoteOccurrencesOptions(accountID, providerID, noteID)
				listNoteOccurrencesOptionsModel.SetAccountID("testString")
				listNoteOccurrencesOptionsModel.SetProviderID("testString")
				listNoteOccurrencesOptionsModel.SetNoteID("testString")
				listNoteOccurrencesOptionsModel.SetTransactionID("testString")
				listNoteOccurrencesOptionsModel.SetPageSize(int64(2))
				listNoteOccurrencesOptionsModel.SetPageToken("testString")
				listNoteOccurrencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNoteOccurrencesOptionsModel).ToNot(BeNil())
				Expect(listNoteOccurrencesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listNoteOccurrencesOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(listNoteOccurrencesOptionsModel.NoteID).To(Equal(core.StringPtr("testString")))
				Expect(listNoteOccurrencesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listNoteOccurrencesOptionsModel.PageSize).To(Equal(core.Int64Ptr(int64(2))))
				Expect(listNoteOccurrencesOptionsModel.PageToken).To(Equal(core.StringPtr("testString")))
				Expect(listNoteOccurrencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNotesOptions successfully`, func() {
				// Construct an instance of the ListNotesOptions model
				accountID := "testString"
				providerID := "testString"
				listNotesOptionsModel := findingsService.NewListNotesOptions(accountID, providerID)
				listNotesOptionsModel.SetAccountID("testString")
				listNotesOptionsModel.SetProviderID("testString")
				listNotesOptionsModel.SetTransactionID("testString")
				listNotesOptionsModel.SetPageSize(int64(2))
				listNotesOptionsModel.SetPageToken("testString")
				listNotesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNotesOptionsModel).ToNot(BeNil())
				Expect(listNotesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listNotesOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(listNotesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listNotesOptionsModel.PageSize).To(Equal(core.Int64Ptr(int64(2))))
				Expect(listNotesOptionsModel.PageToken).To(Equal(core.StringPtr("testString")))
				Expect(listNotesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOccurrencesOptions successfully`, func() {
				// Construct an instance of the ListOccurrencesOptions model
				accountID := "testString"
				providerID := "testString"
				listOccurrencesOptionsModel := findingsService.NewListOccurrencesOptions(accountID, providerID)
				listOccurrencesOptionsModel.SetAccountID("testString")
				listOccurrencesOptionsModel.SetProviderID("testString")
				listOccurrencesOptionsModel.SetTransactionID("testString")
				listOccurrencesOptionsModel.SetPageSize(int64(2))
				listOccurrencesOptionsModel.SetPageToken("testString")
				listOccurrencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOccurrencesOptionsModel).ToNot(BeNil())
				Expect(listOccurrencesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listOccurrencesOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(listOccurrencesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listOccurrencesOptionsModel.PageSize).To(Equal(core.Int64Ptr(int64(2))))
				Expect(listOccurrencesOptionsModel.PageToken).To(Equal(core.StringPtr("testString")))
				Expect(listOccurrencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProvidersOptions successfully`, func() {
				// Construct an instance of the ListProvidersOptions model
				accountID := "testString"
				listProvidersOptionsModel := findingsService.NewListProvidersOptions(accountID)
				listProvidersOptionsModel.SetAccountID("testString")
				listProvidersOptionsModel.SetTransactionID("testString")
				listProvidersOptionsModel.SetLimit(int64(2))
				listProvidersOptionsModel.SetSkip(int64(38))
				listProvidersOptionsModel.SetStartProviderID("testString")
				listProvidersOptionsModel.SetEndProviderID("testString")
				listProvidersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProvidersOptionsModel).ToNot(BeNil())
				Expect(listProvidersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProvidersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listProvidersOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(2))))
				Expect(listProvidersOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listProvidersOptionsModel.StartProviderID).To(Equal(core.StringPtr("testString")))
				Expect(listProvidersOptionsModel.EndProviderID).To(Equal(core.StringPtr("testString")))
				Expect(listProvidersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostGraphOptions successfully`, func() {
				// Construct an instance of the PostGraphOptions model
				accountID := "testString"
				postGraphOptionsModel := findingsService.NewPostGraphOptions(accountID)
				postGraphOptionsModel.SetAccountID("testString")
				postGraphOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				postGraphOptionsModel.SetContentType("application/json")
				postGraphOptionsModel.SetTransactionID("testString")
				postGraphOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postGraphOptionsModel).ToNot(BeNil())
				Expect(postGraphOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(postGraphOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(postGraphOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(postGraphOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(postGraphOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReporter successfully`, func() {
				id := "testString"
				title := "testString"
				model, err := findingsService.NewReporter(id, title)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSection successfully`, func() {
				title := "testString"
				image := "testString"
				model, err := findingsService.NewSection(title, image)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSocketAddress successfully`, func() {
				address := "testString"
				model, err := findingsService.NewSocketAddress(address)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateNoteOptions successfully`, func() {
				// Construct an instance of the Reporter model
				reporterModel := new(findingsv1.Reporter)
				Expect(reporterModel).ToNot(BeNil())
				reporterModel.ID = core.StringPtr("testString")
				reporterModel.Title = core.StringPtr("testString")
				reporterModel.URL = core.StringPtr("testString")
				Expect(reporterModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(reporterModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(reporterModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the APINoteRelatedURL model
				apiNoteRelatedURLModel := new(findingsv1.APINoteRelatedURL)
				Expect(apiNoteRelatedURLModel).ToNot(BeNil())
				apiNoteRelatedURLModel.Label = core.StringPtr("testString")
				apiNoteRelatedURLModel.URL = core.StringPtr("testString")
				Expect(apiNoteRelatedURLModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(apiNoteRelatedURLModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				Expect(remediationStepModel).ToNot(BeNil())
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")
				Expect(remediationStepModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(remediationStepModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the FindingType model
				findingTypeModel := new(findingsv1.FindingType)
				Expect(findingTypeModel).ToNot(BeNil())
				findingTypeModel.Severity = core.StringPtr("LOW")
				findingTypeModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				Expect(findingTypeModel.Severity).To(Equal(core.StringPtr("LOW")))
				Expect(findingTypeModel.NextSteps).To(Equal([]findingsv1.RemediationStep{*remediationStepModel}))

				// Construct an instance of the KpiType model
				kpiTypeModel := new(findingsv1.KpiType)
				Expect(kpiTypeModel).ToNot(BeNil())
				kpiTypeModel.AggregationType = core.StringPtr("SUM")
				Expect(kpiTypeModel.AggregationType).To(Equal(core.StringPtr("SUM")))

				// Construct an instance of the ValueTypeFindingCountValueType model
				valueTypeModel := new(findingsv1.ValueTypeFindingCountValueType)
				Expect(valueTypeModel).ToNot(BeNil())
				valueTypeModel.Kind = core.StringPtr("FINDING_COUNT")
				valueTypeModel.FindingNoteNames = []string{"testString"}
				valueTypeModel.Text = core.StringPtr("testString")
				Expect(valueTypeModel.Kind).To(Equal(core.StringPtr("FINDING_COUNT")))
				Expect(valueTypeModel.FindingNoteNames).To(Equal([]string{"testString"}))
				Expect(valueTypeModel.Text).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CardElementTimeSeriesCardElement model
				cardElementModel := new(findingsv1.CardElementTimeSeriesCardElement)
				Expect(cardElementModel).ToNot(BeNil())
				cardElementModel.Text = core.StringPtr("testString")
				cardElementModel.DefaultInterval = core.StringPtr("testString")
				cardElementModel.Kind = core.StringPtr("TIME_SERIES")
				cardElementModel.DefaultTimeRange = core.StringPtr("1d")
				cardElementModel.ValueTypes = []findingsv1.ValueTypeIntf{valueTypeModel}
				Expect(cardElementModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(cardElementModel.DefaultInterval).To(Equal(core.StringPtr("testString")))
				Expect(cardElementModel.Kind).To(Equal(core.StringPtr("TIME_SERIES")))
				Expect(cardElementModel.DefaultTimeRange).To(Equal(core.StringPtr("1d")))
				Expect(cardElementModel.ValueTypes).To(Equal([]findingsv1.ValueTypeIntf{valueTypeModel}))

				// Construct an instance of the Card model
				cardModel := new(findingsv1.Card)
				Expect(cardModel).ToNot(BeNil())
				cardModel.Section = core.StringPtr("testString")
				cardModel.Title = core.StringPtr("testString")
				cardModel.Subtitle = core.StringPtr("testString")
				cardModel.Order = core.Int64Ptr(int64(1))
				cardModel.FindingNoteNames = []string{"testString"}
				cardModel.RequiresConfiguration = core.BoolPtr(true)
				cardModel.BadgeText = core.StringPtr("testString")
				cardModel.BadgeImage = core.StringPtr("testString")
				cardModel.Elements = []findingsv1.CardElementIntf{cardElementModel}
				Expect(cardModel.Section).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Subtitle).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Order).To(Equal(core.Int64Ptr(int64(1))))
				Expect(cardModel.FindingNoteNames).To(Equal([]string{"testString"}))
				Expect(cardModel.RequiresConfiguration).To(Equal(core.BoolPtr(true)))
				Expect(cardModel.BadgeText).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.BadgeImage).To(Equal(core.StringPtr("testString")))
				Expect(cardModel.Elements).To(Equal([]findingsv1.CardElementIntf{cardElementModel}))

				// Construct an instance of the Section model
				sectionModel := new(findingsv1.Section)
				Expect(sectionModel).ToNot(BeNil())
				sectionModel.Title = core.StringPtr("testString")
				sectionModel.Image = core.StringPtr("testString")
				Expect(sectionModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(sectionModel.Image).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateNoteOptions model
				accountID := "testString"
				providerID := "testString"
				noteID := "testString"
				updateNoteOptionsShortDescription := "testString"
				updateNoteOptionsLongDescription := "testString"
				updateNoteOptionsKind := "FINDING"
				updateNoteOptionsID := "testString"
				var updateNoteOptionsReportedBy *findingsv1.Reporter = nil
				updateNoteOptionsModel := findingsService.NewUpdateNoteOptions(accountID, providerID, noteID, updateNoteOptionsShortDescription, updateNoteOptionsLongDescription, updateNoteOptionsKind, updateNoteOptionsID, updateNoteOptionsReportedBy)
				updateNoteOptionsModel.SetAccountID("testString")
				updateNoteOptionsModel.SetProviderID("testString")
				updateNoteOptionsModel.SetNoteID("testString")
				updateNoteOptionsModel.SetShortDescription("testString")
				updateNoteOptionsModel.SetLongDescription("testString")
				updateNoteOptionsModel.SetKind("FINDING")
				updateNoteOptionsModel.SetID("testString")
				updateNoteOptionsModel.SetReportedBy(reporterModel)
				updateNoteOptionsModel.SetRelatedURL([]findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel})
				updateNoteOptionsModel.SetExpirationTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				updateNoteOptionsModel.SetShared(true)
				updateNoteOptionsModel.SetFinding(findingTypeModel)
				updateNoteOptionsModel.SetKpi(kpiTypeModel)
				updateNoteOptionsModel.SetCard(cardModel)
				updateNoteOptionsModel.SetSection(sectionModel)
				updateNoteOptionsModel.SetTransactionID("testString")
				updateNoteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateNoteOptionsModel).ToNot(BeNil())
				Expect(updateNoteOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.NoteID).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.Kind).To(Equal(core.StringPtr("FINDING")))
				Expect(updateNoteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.ReportedBy).To(Equal(reporterModel))
				Expect(updateNoteOptionsModel.RelatedURL).To(Equal([]findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel}))
				Expect(updateNoteOptionsModel.ExpirationTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(updateNoteOptionsModel.Shared).To(Equal(core.BoolPtr(true)))
				Expect(updateNoteOptionsModel.Finding).To(Equal(findingTypeModel))
				Expect(updateNoteOptionsModel.Kpi).To(Equal(kpiTypeModel))
				Expect(updateNoteOptionsModel.Card).To(Equal(cardModel))
				Expect(updateNoteOptionsModel.Section).To(Equal(sectionModel))
				Expect(updateNoteOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateNoteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOccurrenceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(findingsv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.Region = core.StringPtr("testString")
				contextModel.ResourceCRN = core.StringPtr("testString")
				contextModel.ResourceID = core.StringPtr("testString")
				contextModel.ResourceName = core.StringPtr("testString")
				contextModel.ResourceType = core.StringPtr("testString")
				contextModel.ServiceCRN = core.StringPtr("testString")
				contextModel.ServiceName = core.StringPtr("testString")
				contextModel.EnvironmentName = core.StringPtr("testString")
				contextModel.ComponentName = core.StringPtr("testString")
				contextModel.ToolchainID = core.StringPtr("testString")
				Expect(contextModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceCRN).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ServiceCRN).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.EnvironmentName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ComponentName).To(Equal(core.StringPtr("testString")))
				Expect(contextModel.ToolchainID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RemediationStep model
				remediationStepModel := new(findingsv1.RemediationStep)
				Expect(remediationStepModel).ToNot(BeNil())
				remediationStepModel.Title = core.StringPtr("testString")
				remediationStepModel.URL = core.StringPtr("testString")
				Expect(remediationStepModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(remediationStepModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SocketAddress model
				socketAddressModel := new(findingsv1.SocketAddress)
				Expect(socketAddressModel).ToNot(BeNil())
				socketAddressModel.Address = core.StringPtr("testString")
				socketAddressModel.Port = core.Int64Ptr(int64(38))
				Expect(socketAddressModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(socketAddressModel.Port).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NetworkConnection model
				networkConnectionModel := new(findingsv1.NetworkConnection)
				Expect(networkConnectionModel).ToNot(BeNil())
				networkConnectionModel.Direction = core.StringPtr("testString")
				networkConnectionModel.Protocol = core.StringPtr("testString")
				networkConnectionModel.Client = socketAddressModel
				networkConnectionModel.Server = socketAddressModel
				Expect(networkConnectionModel.Direction).To(Equal(core.StringPtr("testString")))
				Expect(networkConnectionModel.Protocol).To(Equal(core.StringPtr("testString")))
				Expect(networkConnectionModel.Client).To(Equal(socketAddressModel))
				Expect(networkConnectionModel.Server).To(Equal(socketAddressModel))

				// Construct an instance of the DataTransferred model
				dataTransferredModel := new(findingsv1.DataTransferred)
				Expect(dataTransferredModel).ToNot(BeNil())
				dataTransferredModel.ClientBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerBytes = core.Int64Ptr(int64(38))
				dataTransferredModel.ClientPackets = core.Int64Ptr(int64(38))
				dataTransferredModel.ServerPackets = core.Int64Ptr(int64(38))
				Expect(dataTransferredModel.ClientBytes).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ServerBytes).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ClientPackets).To(Equal(core.Int64Ptr(int64(38))))
				Expect(dataTransferredModel.ServerPackets).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Finding model
				findingModel := new(findingsv1.Finding)
				Expect(findingModel).ToNot(BeNil())
				findingModel.Severity = core.StringPtr("LOW")
				findingModel.Certainty = core.StringPtr("LOW")
				findingModel.NextSteps = []findingsv1.RemediationStep{*remediationStepModel}
				findingModel.NetworkConnection = networkConnectionModel
				findingModel.DataTransferred = dataTransferredModel
				Expect(findingModel.Severity).To(Equal(core.StringPtr("LOW")))
				Expect(findingModel.Certainty).To(Equal(core.StringPtr("LOW")))
				Expect(findingModel.NextSteps).To(Equal([]findingsv1.RemediationStep{*remediationStepModel}))
				Expect(findingModel.NetworkConnection).To(Equal(networkConnectionModel))
				Expect(findingModel.DataTransferred).To(Equal(dataTransferredModel))

				// Construct an instance of the Kpi model
				kpiModel := new(findingsv1.Kpi)
				Expect(kpiModel).ToNot(BeNil())
				kpiModel.Value = core.Float64Ptr(float64(72.5))
				kpiModel.Total = core.Float64Ptr(float64(72.5))
				Expect(kpiModel.Value).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(kpiModel.Total).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the UpdateOccurrenceOptions model
				accountID := "testString"
				providerID := "testString"
				occurrenceID := "testString"
				updateOccurrenceOptionsNoteName := "testString"
				updateOccurrenceOptionsKind := "FINDING"
				updateOccurrenceOptionsID := "testString"
				updateOccurrenceOptionsModel := findingsService.NewUpdateOccurrenceOptions(accountID, providerID, occurrenceID, updateOccurrenceOptionsNoteName, updateOccurrenceOptionsKind, updateOccurrenceOptionsID)
				updateOccurrenceOptionsModel.SetAccountID("testString")
				updateOccurrenceOptionsModel.SetProviderID("testString")
				updateOccurrenceOptionsModel.SetOccurrenceID("testString")
				updateOccurrenceOptionsModel.SetNoteName("testString")
				updateOccurrenceOptionsModel.SetKind("FINDING")
				updateOccurrenceOptionsModel.SetID("testString")
				updateOccurrenceOptionsModel.SetResourceURL("testString")
				updateOccurrenceOptionsModel.SetRemediation("testString")
				updateOccurrenceOptionsModel.SetContext(contextModel)
				updateOccurrenceOptionsModel.SetFinding(findingModel)
				updateOccurrenceOptionsModel.SetKpi(kpiModel)
				updateOccurrenceOptionsModel.SetReferenceData(map[string]interface{}{"anyKey": "anyValue"})
				updateOccurrenceOptionsModel.SetTransactionID("testString")
				updateOccurrenceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOccurrenceOptionsModel).ToNot(BeNil())
				Expect(updateOccurrenceOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.OccurrenceID).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.NoteName).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.Kind).To(Equal(core.StringPtr("FINDING")))
				Expect(updateOccurrenceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.ResourceURL).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.Remediation).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.Context).To(Equal(contextModel))
				Expect(updateOccurrenceOptionsModel.Finding).To(Equal(findingModel))
				Expect(updateOccurrenceOptionsModel.Kpi).To(Equal(kpiModel))
				Expect(updateOccurrenceOptionsModel.ReferenceData).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateOccurrenceOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateOccurrenceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAPINote successfully`, func() {
				shortDescription := "testString"
				longDescription := "testString"
				kind := "FINDING"
				id := "testString"
				var reportedBy *findingsv1.Reporter = nil
				_, err := findingsService.NewAPINote(shortDescription, longDescription, kind, id, reportedBy)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAPINoteRelatedURL successfully`, func() {
				label := "testString"
				url := "testString"
				model, err := findingsService.NewAPINoteRelatedURL(label, url)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAPIOccurrence successfully`, func() {
				noteName := "testString"
				kind := "FINDING"
				id := "testString"
				model, err := findingsService.NewAPIOccurrence(noteName, kind, id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCardElementBreakdownCardElement successfully`, func() {
				text := "testString"
				kind := "BREAKDOWN"
				valueTypes := []findingsv1.ValueTypeIntf{}
				model, err := findingsService.NewCardElementBreakdownCardElement(text, kind, valueTypes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCardElementNumericCardElement successfully`, func() {
				text := "testString"
				kind := "NUMERIC"
				var valueType *findingsv1.NumericCardElementValueType = nil
				_, err := findingsService.NewCardElementNumericCardElement(text, kind, valueType)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCardElementTimeSeriesCardElement successfully`, func() {
				text := "testString"
				kind := "TIME_SERIES"
				valueTypes := []findingsv1.ValueTypeIntf{}
				model, err := findingsService.NewCardElementTimeSeriesCardElement(text, kind, valueTypes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewValueTypeFindingCountValueType successfully`, func() {
				kind := "FINDING_COUNT"
				findingNoteNames := []string{"testString"}
				text := "testString"
				model, err := findingsService.NewValueTypeFindingCountValueType(kind, findingNoteNames, text)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewValueTypeKpiValueType successfully`, func() {
				kind := "KPI"
				kpiNoteName := "testString"
				text := "testString"
				model, err := findingsService.NewValueTypeKpiValueType(kind, kpiNoteName, text)
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

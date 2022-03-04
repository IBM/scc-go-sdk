//go:build integration
// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/addonmgrv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the addonmgrv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AddonMgrV1 Integration Tests`, func() {

	const externalConfigFile = "../addon_mgr_v1.env"

	var (
		err             error
		addonMgrService *addonmgrv1.AddonMgrV1
		serviceURL      string
		config          map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(addonmgrv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			addonMgrServiceOptions := &addonmgrv1.AddonMgrV1Options{
				AccountID: core.StringPtr(config["ACCOUNT_ID"]),
			}

			addonMgrService, err = addonmgrv1.NewAddonMgrV1UsingExternalConfig(addonMgrServiceOptions)

			Expect(err).To(BeNil())
			Expect(addonMgrService).ToNot(BeNil())
			Expect(addonMgrService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			addonMgrService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`AddNetworkInsightsCosDetailsV2 - Add cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options)`, func() {

			cosDetailsV2CosDetailsItemModel := &addonmgrv1.CosDetailsV2CosDetailsItem{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				Type:         core.StringPtr("network-insights"),
				CosBucketURL: core.StringPtr("testString"),
			}

			addNetworkInsightsCosDetailsV2Options := &addonmgrv1.AddNetworkInsightsCosDetailsV2Options{
				RegionID:   core.StringPtr(config["REGION_ID"]),
				CosDetails: []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel},
			}

			response, err := addonMgrService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`AddActivityInsightsCosDetailsV2 - Add cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options)`, func() {

			cosDetailsV2CosDetailsItemModel := &addonmgrv1.CosDetailsV2CosDetailsItem{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				Type:         core.StringPtr("network-insights"),
				CosBucketURL: core.StringPtr("testString"),
			}

			addActivityInsightsCosDetailsV2Options := &addonmgrv1.AddActivityInsightsCosDetailsV2Options{
				RegionID:   core.StringPtr(config["REGION_ID"]),
				CosDetails: []addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel},
			}

			response, err := addonMgrService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`DisableInsightsV2 - Disable add-on`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DisableInsightsV2(disableInsightsV2Options *DisableInsightsV2Options)`, func() {

			disableInsightsV2Options := &addonmgrv1.DisableInsightsV2Options{
				RegionID:         core.StringPtr(config["REGION_ID"]),
				NetworkInsights:  core.BoolPtr(true),
				ActivityInsights: core.BoolPtr(true),
			}

			response, err := addonMgrService.DisableInsightsV2(disableInsightsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`EnableInsightsV2 - Enable add-on`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`EnableInsightsV2(enableInsightsV2Options *EnableInsightsV2Options)`, func() {

			enableInsightsV2Options := &addonmgrv1.EnableInsightsV2Options{
				RegionID:         core.StringPtr(config["REGION_ID"]),
				NetworkInsights:  core.BoolPtr(true),
				ActivityInsights: core.BoolPtr(true),
			}

			response, err := addonMgrService.EnableInsightsV2(enableInsightsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetSupportedInsightsV2 - Fetch supported insights`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options)`, func() {

			getSupportedInsightsV2Options := &addonmgrv1.GetSupportedInsightsV2Options{}

			allInsights, response, err := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allInsights).ToNot(BeNil())

		})
	})

	Describe(`DeleteNetworkInsightsCosDetailsV2 - Delete cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options)`, func() {

			deleteNetworkInsightsCosDetailsV2Options := &addonmgrv1.DeleteNetworkInsightsCosDetailsV2Options{
				Ids: []string{"testString"},
			}

			response, err := addonMgrService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteActivityInsightsCosDetailsV2 - Delete cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options)`, func() {

			deleteActivityInsightsCosDetailsV2Options := &addonmgrv1.DeleteActivityInsightsCosDetailsV2Options{
				Ids: []string{"testString"},
			}

			response, err := addonMgrService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

//
// Utility functions are declared in the unit test file
//

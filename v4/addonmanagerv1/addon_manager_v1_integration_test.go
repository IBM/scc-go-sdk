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

package addonmanagerv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the addonmanagerv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var accountID = os.Getenv("ACCOUNT_ID")

var _ = Describe(`AddonManagerV1 Integration Tests`, func() {

	const externalConfigFile = "../addon_manager_v1.env"

	var (
		err                 error
		addonManagerService *addonmanagerv1.AddonManagerV1
		serviceURL          string
		config              map[string]string
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
			config, err = core.GetServiceProperties(addonmanagerv1.DefaultServiceName)
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

			addonManagerServiceOptions := &addonmanagerv1.AddonManagerV1Options{
				AccountID: &accountID,
			}

			addonManagerService, err = addonmanagerv1.NewAddonManagerV1UsingExternalConfig(addonManagerServiceOptions)

			Expect(err).To(BeNil())
			Expect(addonManagerService).ToNot(BeNil())
			Expect(addonManagerService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			addonManagerService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSupportedInsightsV2 - Fetch supported insights`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupportedInsightsV2(getSupportedInsightsV2Options *GetSupportedInsightsV2Options)`, func() {

			getSupportedInsightsV2Options := &addonmanagerv1.GetSupportedInsightsV2Options{}

			allInsights, response, err := addonManagerService.GetSupportedInsightsV2(getSupportedInsightsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allInsights).ToNot(BeNil())

		})
	})

	Describe(`GetNetworkInsightStatusV2 - Get Network Insight status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNetworkInsightStatusV2(getNetworkInsightStatusV2Options *GetNetworkInsightStatusV2Options)`, func() {

			getNetworkInsightStatusV2Options := &addonmanagerv1.GetNetworkInsightStatusV2Options{}

			_, response, err := addonManagerService.GetNetworkInsightStatusV2(getNetworkInsightStatusV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetActivityInsightStatusV2 - Get Activity Insight status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActivityInsightStatusV2(getActivityInsightStatusV2Options *GetActivityInsightStatusV2Options)`, func() {

			getActivityInsightStatusV2Options := &addonmanagerv1.GetActivityInsightStatusV2Options{}

			_, response, err := addonManagerService.GetActivityInsightStatusV2(getActivityInsightStatusV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`UpdateNetworkInsightStatusV2 - Enable/Disable Network Insight`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNetworkInsightStatusV2(updateNetworkInsightStatusV2Options *UpdateNetworkInsightStatusV2Options)`, func() {

			updateNetworkInsightStatusV2Options := &addonmanagerv1.UpdateNetworkInsightStatusV2Options{
				RegionID: core.StringPtr(config["REGION_ID"]),
				Status:   core.StringPtr("enable"),
			}

			response, err := addonManagerService.UpdateNetworkInsightStatusV2(updateNetworkInsightStatusV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`UpdateActivityInsightStatusV2 - Enable/Disable Activity Insight`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateActivityInsightStatusV2(updateActivityInsightStatusV2Options *UpdateActivityInsightStatusV2Options)`, func() {

			updateActivityInsightStatusV2Options := &addonmanagerv1.UpdateActivityInsightStatusV2Options{
				RegionID: core.StringPtr(config["REGION_ID"]),
				Status:   core.StringPtr("enable"),
			}

			response, err := addonManagerService.UpdateActivityInsightStatusV2(updateActivityInsightStatusV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`AddNetworkInsightsCosDetailsV2 - Add cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options *AddNetworkInsightsCosDetailsV2Options)`, func() {

			cosDetailsV2CosDetailsItemModel := &addonmanagerv1.CosDetails{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				CosBucketURL: core.StringPtr("testString"),
				Type:         core.StringPtr("network_insights"),
			}

			addNetworkInsightsCosDetailsV2Options := &addonmanagerv1.AddNetworkInsightsCosDetailsV2Options{
				RegionID:   core.StringPtr(config["REGION_ID"]),
				CosDetails: []addonmanagerv1.CosDetails{*cosDetailsV2CosDetailsItemModel},
			}

			_, response, err := addonManagerService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`AddActivityInsightsCosDetailsV2 - Add cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options *AddActivityInsightsCosDetailsV2Options)`, func() {

			cosDetailsV2CosDetailsItemModel := &addonmanagerv1.CosDetails{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				CosBucketURL: core.StringPtr("testString"),
				Type:         core.StringPtr("activity_insights"),
			}

			addActivityInsightsCosDetailsV2Options := &addonmanagerv1.AddActivityInsightsCosDetailsV2Options{
				RegionID:   core.StringPtr(config["REGION_ID"]),
				CosDetails: []addonmanagerv1.CosDetails{*cosDetailsV2CosDetailsItemModel},
			}

			_, response, err := addonManagerService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`GetNetworkInsightsCosDetailsV2 - Get cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options *GetNetworkInsightsCosDetailsV2Options)`, func() {

			getNetworkInsightsCosDetailsV2Options := &addonmanagerv1.GetNetworkInsightsCosDetailsV2Options{}

			_, response, err := addonManagerService.GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetActivityInsightsCosDetailsV2 - Get cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options *GetActivityInsightsCosDetailsV2Options)`, func() {

			getActivityInsightsCosDetailsV2Options := &addonmanagerv1.GetActivityInsightsCosDetailsV2Options{}

			_, response, err := addonManagerService.GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteNetworkInsightsCosDetailsV2 - Delete cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options *DeleteNetworkInsightsCosDetailsV2Options)`, func() {

			deleteNetworkInsightsCosDetailsV2Options := &addonmanagerv1.DeleteNetworkInsightsCosDetailsV2Options{
				Ids: []string{"e0b05667-407f-4a7c-9b45-bc5332e9ad7b"},
			}

			response, err := addonManagerService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteActivityInsightsCosDetailsV2 - Delete cos details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options *DeleteActivityInsightsCosDetailsV2Options)`, func() {

			deleteActivityInsightsCosDetailsV2Options := &addonmanagerv1.DeleteActivityInsightsCosDetailsV2Options{
				Ids: []string{"e0b05667-407f-4a7c-9b45-bc5332e9ad7b"},
			}

			response, err := addonManagerService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

//
// Utility functions are declared in the unit test file
//

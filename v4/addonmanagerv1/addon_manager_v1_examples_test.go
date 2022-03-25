// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Addon Manager service.
//
// The following configuration properties are assumed to be defined:
// ADDON_MANAGER_URL=<service base url>
// ADDON_MANAGER_AUTH_TYPE=iam
// ADDON_MANAGER_APIKEY=<IAM apikey>
// ADDON_MANAGER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`AddonManagerV1 Examples Tests`, func() {

	const externalConfigFile = "../addon_manager_v1.env"

	var (
		addonManagerService *addonmanagerv1.AddonManagerV1
		config              map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(addonmanagerv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			addonManagerServiceOptions := &addonmanagerv1.AddonManagerV1Options{
				AccountID: core.StringPtr("testString"),
			}

			addonManagerService, err = addonmanagerv1.NewAddonManagerV1UsingExternalConfig(addonManagerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(addonManagerService).ToNot(BeNil())
		})
	})

	Describe(`AddonManagerV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupportedInsightsV2 request example`, func() {
			fmt.Println("\nGetSupportedInsightsV2() result:")
			// begin-get_supported_insights_v2

			getSupportedInsightsV2Options := addonManagerService.NewGetSupportedInsightsV2Options()

			allInsights, response, err := addonManagerService.GetSupportedInsightsV2(getSupportedInsightsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(allInsights, "", "  ")
			fmt.Println(string(b))

			// end-get_supported_insights_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allInsights).ToNot(BeNil())

		})
		It(`AddNetworkInsightsCosDetailsV2 request example`, func() {
			fmt.Println("\nAddNetworkInsightsCosDetailsV2() result:")
			// begin-add_network_insights_cos_details_v2

			cosDetailsModel := &addonmanagerv1.CosDetails{}

			addNetworkInsightsCosDetailsV2Options := addonManagerService.NewAddNetworkInsightsCosDetailsV2Options(
				"testString",
				[]addonmanagerv1.CosDetails{*cosDetailsModel},
			)

			networkInsightsCosDetailsOutput, response, err := addonManagerService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInsightsCosDetailsOutput, "", "  ")
			fmt.Println(string(b))

			// end-add_network_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInsightsCosDetailsOutput).ToNot(BeNil())

		})
		It(`GetNetworkInsightsCosDetailsV2 request example`, func() {
			fmt.Println("\nGetNetworkInsightsCosDetailsV2() result:")
			// begin-get_network_insights_cos_details_v2

			getNetworkInsightsCosDetailsV2Options := addonManagerService.NewGetNetworkInsightsCosDetailsV2Options()

			networkInsightsCosDetailsOutput, response, err := addonManagerService.GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInsightsCosDetailsOutput, "", "  ")
			fmt.Println(string(b))

			// end-get_network_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInsightsCosDetailsOutput).ToNot(BeNil())

		})
		It(`GetNetworkInsightStatusV2 request example`, func() {
			fmt.Println("\nGetNetworkInsightStatusV2() result:")
			// begin-get_network_insight_status_v2

			getNetworkInsightStatusV2Options := addonManagerService.NewGetNetworkInsightStatusV2Options()

			networkInsightsStatusConfigOutput, response, err := addonManagerService.GetNetworkInsightStatusV2(getNetworkInsightStatusV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInsightsStatusConfigOutput, "", "  ")
			fmt.Println(string(b))

			// end-get_network_insight_status_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInsightsStatusConfigOutput).ToNot(BeNil())

		})
		It(`UpdateNetworkInsightStatusV2 request example`, func() {
			// begin-update_network_insight_status_v2

			updateNetworkInsightStatusV2Options := addonManagerService.NewUpdateNetworkInsightStatusV2Options(
				"testString",
				"enable",
			)

			response, err := addonManagerService.UpdateNetworkInsightStatusV2(updateNetworkInsightStatusV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from UpdateNetworkInsightStatusV2(): %d\n", response.StatusCode)
			}

			// end-update_network_insight_status_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`AddActivityInsightsCosDetailsV2 request example`, func() {
			fmt.Println("\nAddActivityInsightsCosDetailsV2() result:")
			// begin-add_activity_insights_cos_details_v2

			cosDetailsModel := &addonmanagerv1.CosDetails{}

			addActivityInsightsCosDetailsV2Options := addonManagerService.NewAddActivityInsightsCosDetailsV2Options(
				"testString",
				[]addonmanagerv1.CosDetails{*cosDetailsModel},
			)

			activityInsightsCosDetailsOutput, response, err := addonManagerService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(activityInsightsCosDetailsOutput, "", "  ")
			fmt.Println(string(b))

			// end-add_activity_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(activityInsightsCosDetailsOutput).ToNot(BeNil())

		})
		It(`GetActivityInsightsCosDetailsV2 request example`, func() {
			fmt.Println("\nGetActivityInsightsCosDetailsV2() result:")
			// begin-get_activity_insights_cos_details_v2

			getActivityInsightsCosDetailsV2Options := addonManagerService.NewGetActivityInsightsCosDetailsV2Options()

			activityInsightsCosDetailsOutput, response, err := addonManagerService.GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(activityInsightsCosDetailsOutput, "", "  ")
			fmt.Println(string(b))

			// end-get_activity_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(activityInsightsCosDetailsOutput).ToNot(BeNil())

		})
		It(`GetActivityInsightStatusV2 request example`, func() {
			fmt.Println("\nGetActivityInsightStatusV2() result:")
			// begin-get_activity_insight_status_v2

			getActivityInsightStatusV2Options := addonManagerService.NewGetActivityInsightStatusV2Options()

			activityInsightsStatusConfigOutput, response, err := addonManagerService.GetActivityInsightStatusV2(getActivityInsightStatusV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(activityInsightsStatusConfigOutput, "", "  ")
			fmt.Println(string(b))

			// end-get_activity_insight_status_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(activityInsightsStatusConfigOutput).ToNot(BeNil())

		})
		It(`UpdateActivityInsightStatusV2 request example`, func() {
			// begin-update_activity_insight_status_v2

			updateActivityInsightStatusV2Options := addonManagerService.NewUpdateActivityInsightStatusV2Options(
				"testString",
				"enable",
			)

			response, err := addonManagerService.UpdateActivityInsightStatusV2(updateActivityInsightStatusV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from UpdateActivityInsightStatusV2(): %d\n", response.StatusCode)
			}

			// end-update_activity_insight_status_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteNetworkInsightsCosDetailsV2 request example`, func() {
			// begin-delete_network_insights_cos_details_v2

			deleteNetworkInsightsCosDetailsV2Options := addonManagerService.NewDeleteNetworkInsightsCosDetailsV2Options()

			response, err := addonManagerService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteNetworkInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-delete_network_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteActivityInsightsCosDetailsV2 request example`, func() {
			// begin-delete_activity_insights_cos_details_v2

			deleteActivityInsightsCosDetailsV2Options := addonManagerService.NewDeleteActivityInsightsCosDetailsV2Options()

			response, err := addonManagerService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteActivityInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-delete_activity_insights_cos_details_v2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

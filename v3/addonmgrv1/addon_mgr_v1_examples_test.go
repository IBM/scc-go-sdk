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

package addonmgrv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/addonmgrv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the AddonMgr service.
//
// The following configuration properties are assumed to be defined:
// ADDON_MGR_URL=<service base url>
// ADDON_MGR_AUTH_TYPE=iam
// ADDON_MGR_APIKEY=<IAM apikey>
// ADDON_MGR_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`AddonMgrV1 Examples Tests`, func() {

	const externalConfigFile = "../addon_mgr_v1.env"

	var (
		addonMgrService *addonmgrv1.AddonMgrV1
		config          map[string]string
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
			config, err = core.GetServiceProperties(addonmgrv1.DefaultServiceName)
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

			addonMgrServiceOptions := &addonmgrv1.AddonMgrV1Options{
				AccountID: core.StringPtr("testString"),
			}

			addonMgrService, err = addonmgrv1.NewAddonMgrV1UsingExternalConfig(addonMgrServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(addonMgrService).ToNot(BeNil())
		})
	})

	Describe(`AddonMgrV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddNetworkInsightsCosDetailsV2 request example`, func() {
			// begin-addNetworkInsightsCosDetailsV2

			cosDetailsV2CosDetailsItemModel := &addonmgrv1.CosDetailsV2CosDetailsItem{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				Type:         core.StringPtr("network-insights"),
				CosBucketURL: core.StringPtr("testString"),
			}

			addNetworkInsightsCosDetailsV2Options := addonMgrService.NewAddNetworkInsightsCosDetailsV2Options(
				"testString",
				[]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel},
			)

			response, err := addonMgrService.AddNetworkInsightsCosDetailsV2(addNetworkInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from AddNetworkInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-addNetworkInsightsCosDetailsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`AddActivityInsightsCosDetailsV2 request example`, func() {
			// begin-addActivityInsightsCosDetailsV2

			cosDetailsV2CosDetailsItemModel := &addonmgrv1.CosDetailsV2CosDetailsItem{
				CosInstance:  core.StringPtr("testString"),
				BucketName:   core.StringPtr("testString"),
				Description:  core.StringPtr("testString"),
				Type:         core.StringPtr("network-insights"),
				CosBucketURL: core.StringPtr("testString"),
			}

			addActivityInsightsCosDetailsV2Options := addonMgrService.NewAddActivityInsightsCosDetailsV2Options(
				"testString",
				[]addonmgrv1.CosDetailsV2CosDetailsItem{*cosDetailsV2CosDetailsItemModel},
			)

			response, err := addonMgrService.AddActivityInsightsCosDetailsV2(addActivityInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from AddActivityInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-addActivityInsightsCosDetailsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DisableInsightsV2 request example`, func() {
			// begin-disableInsightsV2

			disableInsightsV2Options := addonMgrService.NewDisableInsightsV2Options(
				"testString",
			)

			response, err := addonMgrService.DisableInsightsV2(disableInsightsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DisableInsightsV2(): %d\n", response.StatusCode)
			}

			// end-disableInsightsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`EnableInsightsV2 request example`, func() {
			// begin-enableInsightsV2

			enableInsightsV2Options := addonMgrService.NewEnableInsightsV2Options(
				"testString",
			)

			response, err := addonMgrService.EnableInsightsV2(enableInsightsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from EnableInsightsV2(): %d\n", response.StatusCode)
			}

			// end-enableInsightsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetSupportedInsightsV2 request example`, func() {
			fmt.Println("\nGetSupportedInsightsV2() result:")
			// begin-getSupportedInsightsV2

			getSupportedInsightsV2Options := addonMgrService.NewGetSupportedInsightsV2Options()

			allInsights, response, err := addonMgrService.GetSupportedInsightsV2(getSupportedInsightsV2Options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(allInsights, "", "  ")
			fmt.Println(string(b))

			// end-getSupportedInsightsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allInsights).ToNot(BeNil())

		})
		It(`DeleteNetworkInsightsCosDetailsV2 request example`, func() {
			// begin-deleteNetworkInsightsCosDetailsV2

			deleteNetworkInsightsCosDetailsV2Options := addonMgrService.NewDeleteNetworkInsightsCosDetailsV2Options()

			response, err := addonMgrService.DeleteNetworkInsightsCosDetailsV2(deleteNetworkInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteNetworkInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-deleteNetworkInsightsCosDetailsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteActivityInsightsCosDetailsV2 request example`, func() {
			// begin-deleteActivityInsightsCosDetailsV2

			deleteActivityInsightsCosDetailsV2Options := addonMgrService.NewDeleteActivityInsightsCosDetailsV2Options()

			response, err := addonMgrService.DeleteActivityInsightsCosDetailsV2(deleteActivityInsightsCosDetailsV2Options)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteActivityInsightsCosDetailsV2(): %d\n", response.StatusCode)
			}

			// end-deleteActivityInsightsCosDetailsV2

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

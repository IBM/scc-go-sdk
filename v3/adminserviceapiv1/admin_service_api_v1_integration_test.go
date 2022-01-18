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

package adminserviceapiv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/adminserviceapiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the adminserviceapiv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var accountID = os.Getenv("ACCOUNT_ID")

var _ = Describe(`AdminServiceApiV1 Integration Tests`, func() {

	const externalConfigFile = "../admin_service_api_v1.env"

	var (
		err                    error
		adminServiceApiService *adminserviceapiv1.AdminServiceApiV1
		serviceURL             string
		config                 map[string]string
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
			config, err = core.GetServiceProperties(adminserviceapiv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			adminServiceApiServiceOptions := &adminserviceapiv1.AdminServiceApiV1Options{}

			adminServiceApiService, err = adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(adminServiceApiServiceOptions)

			Expect(err).To(BeNil())
			Expect(adminServiceApiService).ToNot(BeNil())
			Expect(adminServiceApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			adminServiceApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetSettings - View account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {

			getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{
				AccountID: &accountID,
			}

			accountSettings, response, err := adminServiceApiService.GetSettings(getSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 500
			//
		})
	})

	Describe(`PatchAccountSettings - Update account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PatchAccountSettings(patchAccountSettingsOptions *PatchAccountSettingsOptions)`, func() {

			locationIdModel := &adminserviceapiv1.LocationID{
				ID: core.StringPtr("us"),
			}

			patchAccountSettingsOptions := &adminserviceapiv1.PatchAccountSettingsOptions{
				AccountID:          &accountID,
				Location:           locationIdModel,
			}

			accountSettings, response, err := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			if response.StatusCode == 204 {
				Expect(accountSettings).To(BeNil())
			} else {
				Expect(accountSettings).ToNot(BeNil())
			}
			Expect(accountSettings).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 403
			// 500
			//
		})
	})

	Describe(`ListLocations - View available locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLocations(listLocationsOptions *ListLocationsOptions)`, func() {

			listLocationsOptions := &adminserviceapiv1.ListLocationsOptions{}

			locations, response, err := adminServiceApiService.ListLocations(listLocationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(locations).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 403
			// 500
			//
		})
	})

	Describe(`GetLocation - View the details of a location`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLocation(getLocationOptions *GetLocationOptions)`, func() {

			getLocationOptions := &adminserviceapiv1.GetLocationOptions{
				LocationID: core.StringPtr("us"),
			}

			location, response, err := adminServiceApiService.GetLocation(getLocationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(location).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 403
			// 404
			// 500
			//
		})
	})

	Describe(`SendTestEvent - Send test event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SendTestEvent(sendTestEventOptions *SendTestEventOptions)`, func() {

			sendTestEventOptions := &adminserviceapiv1.SendTestEventOptions{
				AccountID: &accountID,
			}

			testEvent, response, err := adminServiceApiService.SendTestEvent(sendTestEventOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testEvent).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 403
			// 500
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//

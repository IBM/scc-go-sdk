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

package adminserviceapiv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/scc-go-sdk/v3/adminserviceapiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Admin Service API service.
//
// The following configuration properties are assumed to be defined:
// ADMIN_SERVICE_API_URL=<service base url>
// ADMIN_SERVICE_API_AUTH_TYPE=iam
// ADMIN_SERVICE_API_APIKEY=<IAM apikey>
// ADMIN_SERVICE_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../admin_service_api_v1.env"

var (
	adminServiceApiService *adminserviceapiv1.AdminServiceApiV1
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`AdminServiceApiV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(adminserviceapiv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			adminServiceApiServiceOptions := &adminserviceapiv1.AdminServiceApiV1Options{}

			adminServiceApiService, err = adminserviceapiv1.NewAdminServiceApiV1UsingExternalConfig(adminServiceApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(adminServiceApiService).ToNot(BeNil())
		})
	})

	Describe(`AdminServiceApiV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-GetSettings

			getSettingsOptions := adminServiceApiService.NewGetSettingsOptions(
				"testString",
			)

			accountSettings, response, err := adminServiceApiService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-GetSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`PatchAccountSettings request example`, func() {
			fmt.Println("\nPatchAccountSettings() result:")
			// begin-PatchAccountSettings

			locationIdModel := &adminserviceapiv1.LocationID{
				ID: core.StringPtr("us"),
			}

			notificationsRegistrationModel := &adminserviceapiv1.NotificationsRegistration{
				InstanceCrn: core.StringPtr("testString"),
			}

			patchAccountSettingsOptions := adminServiceApiService.NewPatchAccountSettingsOptions(
				"testString",
				locationIdModel,
				notificationsRegistrationModel,
			)

			accountSettings, response, err := adminServiceApiService.PatchAccountSettings(patchAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-PatchAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`ListLocations request example`, func() {
			fmt.Println("\nListLocations() result:")
			// begin-ListLocations

			listLocationsOptions := adminServiceApiService.NewListLocationsOptions()

			locations, response, err := adminServiceApiService.ListLocations(listLocationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(locations, "", "  ")
			fmt.Println(string(b))

			// end-ListLocations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(locations).ToNot(BeNil())

		})
		It(`GetLocation request example`, func() {
			fmt.Println("\nGetLocation() result:")
			// begin-GetLocation

			getLocationOptions := adminServiceApiService.NewGetLocationOptions(
				"us",
			)

			location, response, err := adminServiceApiService.GetLocation(getLocationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(location, "", "  ")
			fmt.Println(string(b))

			// end-GetLocation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(location).ToNot(BeNil())

		})
		It(`SendTestEvent request example`, func() {
			fmt.Println("\nSendTestEvent() result:")
			// begin-SendTestEvent

			sendTestEventOptions := adminServiceApiService.NewSendTestEventOptions(
				"testString",
			)

			testEvent, response, err := adminServiceApiService.SendTestEvent(sendTestEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testEvent, "", "  ")
			fmt.Println(string(b))

			// end-SendTestEvent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testEvent).ToNot(BeNil())

		})
	})
})

// +build examples

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

package notificationsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm/scc-go-sdk/notificationsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Notifications service.
//
// The following configuration properties are assumed to be defined:
// NOTIFICATIONS_URL=<service base url>
// NOTIFICATIONS_AUTH_TYPE=iam
// NOTIFICATIONS_APIKEY=<IAM apikey>
// NOTIFICATIONS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../notifications_v1.env"

var (
	notificationsService *notificationsv1.NotificationsV1
	config               map[string]string
	configLoaded         bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`NotificationsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(notificationsv1.DefaultServiceName)
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

			notificationsServiceOptions := &notificationsv1.NotificationsV1Options{}

			notificationsService, err = notificationsv1.NewNotificationsV1UsingExternalConfig(notificationsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(notificationsService).ToNot(BeNil())
		})
	})

	Describe(`NotificationsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAllChannels request example`, func() {
			fmt.Println("\nListAllChannels() result:")
			// begin-listAllChannels

			listAllChannelsOptions := notificationsService.NewListAllChannelsOptions(
				"testString",
			)

			channelsList, response, err := notificationsService.ListAllChannels(listAllChannelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelsList, "", "  ")
			fmt.Println(string(b))

			// end-listAllChannels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsList).ToNot(BeNil())

		})
		It(`CreateNotificationChannel request example`, func() {
			fmt.Println("\nCreateNotificationChannel() result:")
			// begin-createNotificationChannel

			createNotificationChannelOptions := notificationsService.NewCreateNotificationChannelOptions(
				"testString",
				"testString",
				"Webhook",
				"testString",
			)

			channelInfo, response, err := notificationsService.CreateNotificationChannel(createNotificationChannelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelInfo, "", "  ")
			fmt.Println(string(b))

			// end-createNotificationChannel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelInfo).ToNot(BeNil())

		})
		It(`GetNotificationChannel request example`, func() {
			fmt.Println("\nGetNotificationChannel() result:")
			// begin-getNotificationChannel

			getNotificationChannelOptions := notificationsService.NewGetNotificationChannelOptions(
				"testString",
				"testString",
			)

			channelGet, response, err := notificationsService.GetNotificationChannel(getNotificationChannelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelGet, "", "  ")
			fmt.Println(string(b))

			// end-getNotificationChannel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelGet).ToNot(BeNil())

		})
		It(`UpdateNotificationChannel request example`, func() {
			fmt.Println("\nUpdateNotificationChannel() result:")
			// begin-updateNotificationChannel

			updateNotificationChannelOptions := notificationsService.NewUpdateNotificationChannelOptions(
				"testString",
				"testString",
				"testString",
				"Webhook",
				"testString",
			)

			channelInfo, response, err := notificationsService.UpdateNotificationChannel(updateNotificationChannelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelInfo, "", "  ")
			fmt.Println(string(b))

			// end-updateNotificationChannel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelInfo).ToNot(BeNil())

		})
		It(`TestNotificationChannel request example`, func() {
			fmt.Println("\nTestNotificationChannel() result:")
			// begin-testNotificationChannel

			testNotificationChannelOptions := notificationsService.NewTestNotificationChannelOptions(
				"testString",
				"testString",
			)

			testChannel, response, err := notificationsService.TestNotificationChannel(testNotificationChannelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(testChannel, "", "  ")
			fmt.Println(string(b))

			// end-testNotificationChannel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testChannel).ToNot(BeNil())

		})
		It(`GetPublicKey request example`, func() {
			fmt.Println("\nGetPublicKey() result:")
			// begin-getPublicKey

			getPublicKeyOptions := notificationsService.NewGetPublicKeyOptions(
				"testString",
			)

			publicKeyGet, response, err := notificationsService.GetPublicKey(getPublicKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicKeyGet, "", "  ")
			fmt.Println(string(b))

			// end-getPublicKey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicKeyGet).ToNot(BeNil())

		})
		It(`DeleteNotificationChannels request example`, func() {
			fmt.Println("\nDeleteNotificationChannels() result:")
			// begin-deleteNotificationChannels

			deleteNotificationChannelsOptions := notificationsService.NewDeleteNotificationChannelsOptions(
				"testString",
				[]string{"testString"},
			)

			channelsDelete, response, err := notificationsService.DeleteNotificationChannels(deleteNotificationChannelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelsDelete, "", "  ")
			fmt.Println(string(b))

			// end-deleteNotificationChannels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsDelete).ToNot(BeNil())

		})
		It(`DeleteNotificationChannel request example`, func() {
			fmt.Println("\nDeleteNotificationChannel() result:")
			// begin-deleteNotificationChannel

			deleteNotificationChannelOptions := notificationsService.NewDeleteNotificationChannelOptions(
				"testString",
				"testString",
			)

			channelDelete, response, err := notificationsService.DeleteNotificationChannel(deleteNotificationChannelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(channelDelete, "", "  ")
			fmt.Println(string(b))

			// end-deleteNotificationChannel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelDelete).ToNot(BeNil())

		})
	})
})

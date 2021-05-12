// +build integration

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

package posturemanagementv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/posturemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the Posture Management v1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`posturemanagementv1 Integration Tests`, func() {

	const externalConfigFile = "../posturemanagement_v1.env"

	var (
		err          error
		postureManagementService *posturemanagementv1.PostureManagementV1
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(posturemanagementv1.DefaultServiceName)
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

			notificationsServiceOptions := &posturemanagementv1.NotificationsV1Options{}

			notificationsService, err = posturemanagementv1.NewNotificationsV1UsingExternalConfig(notificationsServiceOptions)

			Expect(err).To(BeNil())
			Expect(notificationsService).ToNot(BeNil())
			Expect(notificationsService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListAllChannels - list all channels`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAllChannels(listAllChannelsOptions *ListAllChannelsOptions)`, func() {

			listAllChannelsOptions := &posturemanagementv1.ListAllChannelsOptions{
				AccountID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(38)),
				Skip: core.Int64Ptr(int64(38)),
			}

			channelsList, response, err := notificationsService.ListAllChannels(listAllChannelsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsList).ToNot(BeNil())

		})
	})

	Describe(`CreateNotificationChannel - create notification channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNotificationChannel(createNotificationChannelOptions *CreateNotificationChannelOptions)`, func() {

			notificationChannelAlertSourceItemModel := &posturemanagementv1.NotificationChannelAlertSourceItem{
				ProviderName: core.StringPtr("testString"),
				FindingTypes: []string{"testString"},
			}

			createNotificationChannelOptions := &posturemanagementv1.CreateNotificationChannelOptions{
				AccountID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Type: core.StringPtr("Webhook"),
				Endpoint: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Severity: []string{"low"},
				Enabled: core.BoolPtr(true),
				AlertSource: []posturemanagementv1.NotificationChannelAlertSourceItem{*notificationChannelAlertSourceItemModel},
				TransactionID: core.StringPtr("testString"),
			}

			channelInfo, response, err := notificationsService.CreateNotificationChannel(createNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelInfo).ToNot(BeNil())

		})
	})

	Describe(`GetNotificationChannel - get the details of a specific channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNotificationChannel(getNotificationChannelOptions *GetNotificationChannelOptions)`, func() {

			getNotificationChannelOptions := &posturemanagementv1.GetNotificationChannelOptions{
				AccountID: core.StringPtr("testString"),
				ChannelID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			channelGet, response, err := notificationsService.GetNotificationChannel(getNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelGet).ToNot(BeNil())

		})
	})

	Describe(`UpdateNotificationChannel - update notification channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNotificationChannel(updateNotificationChannelOptions *UpdateNotificationChannelOptions)`, func() {

			notificationChannelAlertSourceItemModel := &posturemanagementv1.NotificationChannelAlertSourceItem{
				ProviderName: core.StringPtr("testString"),
				FindingTypes: []string{"testString"},
			}

			updateNotificationChannelOptions := &posturemanagementv1.UpdateNotificationChannelOptions{
				AccountID: core.StringPtr("testString"),
				ChannelID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Type: core.StringPtr("Webhook"),
				Endpoint: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Severity: []string{"low"},
				Enabled: core.BoolPtr(true),
				AlertSource: []posturemanagementv1.NotificationChannelAlertSourceItem{*notificationChannelAlertSourceItemModel},
				TransactionID: core.StringPtr("testString"),
			}

			channelInfo, response, err := notificationsService.UpdateNotificationChannel(updateNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelInfo).ToNot(BeNil())

		})
	})

	Describe(`TestNotificationChannel - test notification channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestNotificationChannel(testNotificationChannelOptions *TestNotificationChannelOptions)`, func() {

			testNotificationChannelOptions := &posturemanagementv1.TestNotificationChannelOptions{
				AccountID: core.StringPtr("testString"),
				ChannelID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			testChannel, response, err := notificationsService.TestNotificationChannel(testNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(testChannel).ToNot(BeNil())

		})
	})

	Describe(`GetPublicKey - fetch notifications public key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPublicKey(getPublicKeyOptions *GetPublicKeyOptions)`, func() {

			getPublicKeyOptions := &posturemanagementv1.GetPublicKeyOptions{
				AccountID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			publicKeyGet, response, err := notificationsService.GetPublicKey(getPublicKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicKeyGet).ToNot(BeNil())

		})
	})

	Describe(`DeleteNotificationChannels - bulk delete of channels`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNotificationChannels(deleteNotificationChannelsOptions *DeleteNotificationChannelsOptions)`, func() {

			deleteNotificationChannelsOptions := &posturemanagementv1.DeleteNotificationChannelsOptions{
				AccountID: core.StringPtr("testString"),
				Body: []string{"testString"},
				TransactionID: core.StringPtr("testString"),
			}

			channelsDelete, response, err := notificationsService.DeleteNotificationChannels(deleteNotificationChannelsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsDelete).ToNot(BeNil())

		})
	})

	Describe(`DeleteNotificationChannel - delete the details of a specific channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNotificationChannel(deleteNotificationChannelOptions *DeleteNotificationChannelOptions)`, func() {

			deleteNotificationChannelOptions := &posturemanagementv1.DeleteNotificationChannelOptions{
				AccountID: core.StringPtr("testString"),
				ChannelID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			channelDelete, response, err := notificationsService.DeleteNotificationChannel(deleteNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelDelete).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//

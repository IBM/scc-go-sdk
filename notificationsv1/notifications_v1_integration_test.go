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

package notificationsv1_test

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm/scc-go-sdk/notificationsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the notificationsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var accountID = os.Getenv("ACCOUNT_ID")
var testString = "testString"
var channelID = ""
var identifier = os.Getenv("TRAVIS_JOB_ID")

var _ = Describe(`NotificationsV1 Integration Tests`, func() {

	if identifier == "" {
		identifier = fmt.Sprintf("%d", time.Now().Unix())
	}

	const externalConfigFile = "../notifications_v1.env"

	var (
		err                  error
		notificationsService *notificationsv1.NotificationsV1
		serviceURL           string
		config               map[string]string
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
			config, err = core.GetServiceProperties(notificationsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			fmt.Printf("Running Integration Tests using AccountID: %s\n", accountID)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			notificationsServiceOptions := &notificationsv1.NotificationsV1Options{}

			notificationsService, err = notificationsv1.NewNotificationsV1UsingExternalConfig(notificationsServiceOptions)

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

			listAllChannelsOptions := &notificationsv1.ListAllChannelsOptions{
				AccountID: &accountID,
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

			notificationChannelAlertSourceItemModel := &notificationsv1.NotificationChannelAlertSourceItem{
				ProviderName: core.StringPtr("VA"),
				FindingTypes: []string{"image_with_vulnerabilities"},
			}

			createNotificationChannelOptions := &notificationsv1.CreateNotificationChannelOptions{
				AccountID:   &accountID,
				Name:        core.StringPtr(fmt.Sprintf("%s-%s", testString, identifier)),
				Type:        core.StringPtr("Webhook"),
				Endpoint:    core.StringPtr("https://webhook.site/136fe1e2-3c3f-4bff-925f-391fbb202546"),
				Description: &testString,
				Severity:    []string{"low"},
				Enabled:     core.BoolPtr(true),
				AlertSource: []notificationsv1.NotificationChannelAlertSourceItem{*notificationChannelAlertSourceItemModel},
			}

			channelInfo, response, err := notificationsService.CreateNotificationChannel(createNotificationChannelOptions)

			channelID = *channelInfo.ChannelID

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

			getNotificationChannelOptions := &notificationsv1.GetNotificationChannelOptions{
				AccountID: &accountID,
				ChannelID: &channelID,
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

			notificationChannelAlertSourceItemModel := &notificationsv1.NotificationChannelAlertSourceItem{
				ProviderName: core.StringPtr("VA"),
				FindingTypes: []string{"image_with_vulnerabilities"},
			}

			updateNotificationChannelOptions := &notificationsv1.UpdateNotificationChannelOptions{
				AccountID:   &accountID,
				ChannelID:   &channelID,
				Name:        core.StringPtr(fmt.Sprintf("%s-%s", testString, identifier)),
				Type:        core.StringPtr("Webhook"),
				Endpoint:    core.StringPtr("https://webhook.site/136fe1e2-3c3f-4bff-925f-391fbb202546"),
				Description: &testString,
				Severity:    []string{"low"},
				Enabled:     core.BoolPtr(true),
				AlertSource: []notificationsv1.NotificationChannelAlertSourceItem{*notificationChannelAlertSourceItemModel},
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

			testNotificationChannelOptions := &notificationsv1.TestNotificationChannelOptions{
				AccountID: &accountID,
				ChannelID: &channelID,
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

			getPublicKeyOptions := &notificationsv1.GetPublicKeyOptions{
				AccountID: &accountID,
			}

			publicKeyGet, response, err := notificationsService.GetPublicKey(getPublicKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicKeyGet).ToNot(BeNil())

		})
	})

	Describe(`DeleteNotificationChannel - delete the details of a specific channel`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNotificationChannel(deleteNotificationChannelOptions *DeleteNotificationChannelOptions)`, func() {

			deleteNotificationChannelOptions := &notificationsv1.DeleteNotificationChannelOptions{
				AccountID: &accountID,
				ChannelID: &channelID,
			}

			channelDelete, response, err := notificationsService.DeleteNotificationChannel(deleteNotificationChannelOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelDelete).ToNot(BeNil())

		})
	})

	Describe(`DeleteNotificationChannels - bulk delete of channels`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNotificationChannels(deleteNotificationChannelsOptions *DeleteNotificationChannelsOptions)`, func() {

			createNotificationChannelOptions := &notificationsv1.CreateNotificationChannelOptions{
				AccountID:   &accountID,
				Name:        core.StringPtr(fmt.Sprintf("%s-%s", testString, identifier)),
				Type:        core.StringPtr("Webhook"),
				Endpoint:    core.StringPtr("https://webhook.site/136fe1e2-3c3f-4bff-925f-391fbb202546"),
				Description: &testString,
				Severity:    []string{"low"},
				Enabled:     core.BoolPtr(true),
			}

			channelInfo, response, err := notificationsService.CreateNotificationChannel(createNotificationChannelOptions)

			channelID = *channelInfo.ChannelID

			deleteNotificationChannelsOptions := &notificationsv1.DeleteNotificationChannelsOptions{
				AccountID: &accountID,
				Body:      []string{channelID},
			}

			channelsDelete, response, err := notificationsService.DeleteNotificationChannels(deleteNotificationChannelsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsDelete).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//

// Cleanup: Delete all notes and occurrences under provider 'testString'
var _ = AfterSuite(func() {
	const externalConfigFile = "../notifications_v1.env"
	_, err := os.Stat(externalConfigFile)
	if err != nil {
		Skip("External configuration file not found, skipping tests: " + err.Error())
	}
	os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
	config, err := core.GetServiceProperties(notificationsv1.DefaultServiceName)
	if err != nil {
		Skip("Error loading service properties, skipping tests: " + err.Error())
	}
	serviceURL := config["URL"]
	if serviceURL == "" {
		Skip("Unable to load service URL configuration property, skipping tests")
	}

	fmt.Printf("cleaning up account: %s\n", accountID)

	notificationsServiceOptions := &notificationsv1.NotificationsV1Options{}

	notificationsService, err := notificationsv1.NewNotificationsV1UsingExternalConfig(notificationsServiceOptions)

	listAllChannelsOptions := &notificationsv1.ListAllChannelsOptions{
		AccountID: &accountID,
	}

	channels, _, err := notificationsService.ListAllChannels(listAllChannelsOptions)
	if err != nil {
		Skip("Error occurred while listing channels for cleanup: " + err.Error())
	}

	for _, channel := range channels.Channels {
		if *channel.ChannelID == channelID {
			deleteNotificationChannelOptions := &notificationsv1.DeleteNotificationChannelOptions{
				AccountID: &accountID,
				ChannelID: channel.ChannelID,
			}

			_, _, err := notificationsService.DeleteNotificationChannel(deleteNotificationChannelOptions)
			if err != nil {
				Skip("Error occurred while deleting channel for cleanup: " + err.Error())
			}
		}
	}

	fmt.Printf("cleanup was successful\n")
})

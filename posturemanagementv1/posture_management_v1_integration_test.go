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

var _ = Describe(`Jason test`, func() {

	const externalConfigFile = "../posturemanagementv1.env"

	var (
		err                      error
		postureManagementService *posturemanagementv1.PostureManagementV1
		serviceURL               string
		config                   map[string]string
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

			notificationsServiceOptions := &posturemanagementv1.PostureManagementV1Options{}

			postureManagementService, err = posturemanagementv1.NewPostureManagementV1UsingExternalConfig(notificationsServiceOptions)

			Expect(err).To(BeNil())
			Expect(postureManagementService).ToNot(BeNil())
			Expect(postureManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`List Scans - list all scans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListScans(listScansOptions *ListScansOptions)`, func() {

			listScansOptions := &posturemanagementv1.ListScansOptions{
				AccountID:     core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Name:          core.StringPtr("testString"),
			}

			channelsList, response, err := postureManagementService.ListAllChannels(listScansOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(channelsList).ToNot(BeNil())

		})
	})

})

//
// Utility functions are declared in the unit test file
//

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

	Describe(`Integration test`, func() {
		It(`Create collector`, func() {
			apiKey := os.Getenv("IAM_API_KEY")
			url := os.Getenv("IAM_APIKEY_URL")
			accountId := os.Getenv("ACCOUNT_ID")
			authenticator := &core.IamAuthenticator{
				ApiKey: apiKey,
				URL:    url, //use for dev/preprod env
			}
			service, _ := posturemanagementv1.NewPostureManagementV1(&posturemanagementv1.PostureManagementV1Options{
				Authenticator: authenticator,
				URL:           "https://asap-dev.compliance.test.cloud.ibm.com", //Specify url or use default
			})

			source := service.NewCreateCollectorOptions(accountId)
			source.SetCollectorName("jason-test-collector-05")
			source.SetCollectorDescription("jason scope")
			source.SetManagedBy("CUSTOMER")
			source.SetIsPublic(true)
			source.SetPassPhrase("secret")

			_, response, err := service.CreateCollector(source)

			if err != nil {
				fmt.Println(response.Result)
				fmt.Println("Failed to create collector: ", err)
				return
			}
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})

//
// Utility functions are declared in the unit test file
//

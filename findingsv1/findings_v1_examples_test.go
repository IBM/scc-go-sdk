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

package findingsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm/scc-go-sdk/findingsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Findings service.
//
// The following configuration properties are assumed to be defined:
// FINDINGS_URL=<service base url>
// FINDINGS_AUTH_TYPE=iam
// FINDINGS_APIKEY=<IAM apikey>
// FINDINGS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../findings_v1.env"

var (
	findingsService *findingsv1.FindingsV1
	config          map[string]string
	configLoaded    bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`FindingsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(findingsv1.DefaultServiceName)
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

			findingsServiceOptions := &findingsv1.FindingsV1Options{}

			findingsService, err = findingsv1.NewFindingsV1UsingExternalConfig(findingsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(findingsService).ToNot(BeNil())
		})
	})

	Describe(`FindingsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostGraph request example`, func() {
			// begin-postGraph

			postGraphOptions := findingsService.NewPostGraphOptions(
				"testString",
			)
			postGraphOptions.SetBody(CreateMockReader("This is a mock file."))

			response, err := findingsService.PostGraph(postGraphOptions)
			if err != nil {
				panic(err)
			}

			// end-postGraph
			fmt.Printf("\nPostGraph() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`CreateNote request example`, func() {
			fmt.Println("\nCreateNote() result:")
			// begin-createNote

			reporterModel := &findingsv1.Reporter{
				ID:    core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
			}

			createNoteOptions := findingsService.NewCreateNoteOptions(
				"testString",
				"testString",
				"testString",
				"testString",
				"FINDING",
				"testString",
				reporterModel,
			)

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiNote, "", "  ")
			fmt.Println(string(b))

			// end-createNote

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
		It(`ListNotes request example`, func() {
			fmt.Println("\nListNotes() result:")
			// begin-listNotes

			listNotesOptions := findingsService.NewListNotesOptions(
				"testString",
				"testString",
			)

			apiListNotesResponse, response, err := findingsService.ListNotes(listNotesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiListNotesResponse, "", "  ")
			fmt.Println(string(b))

			// end-listNotes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNotesResponse).ToNot(BeNil())

		})
		It(`GetNote request example`, func() {
			fmt.Println("\nGetNote() result:")
			// begin-getNote

			getNoteOptions := findingsService.NewGetNoteOptions(
				"testString",
				"testString",
				"testString",
			)

			apiNote, response, err := findingsService.GetNote(getNoteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiNote, "", "  ")
			fmt.Println(string(b))

			// end-getNote

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
		It(`UpdateNote request example`, func() {
			fmt.Println("\nUpdateNote() result:")
			// begin-updateNote

			reporterModel := &findingsv1.Reporter{
				ID:    core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
			}

			updateNoteOptions := findingsService.NewUpdateNoteOptions(
				"testString",
				"testString",
				"testString",
				"testString",
				"testString",
				"FINDING",
				"testString",
				reporterModel,
			)

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiNote, "", "  ")
			fmt.Println(string(b))

			// end-updateNote

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
		It(`GetOccurrenceNote request example`, func() {
			fmt.Println("\nGetOccurrenceNote() result:")
			// begin-getOccurrenceNote

			getOccurrenceNoteOptions := findingsService.NewGetOccurrenceNoteOptions(
				"testString",
				"testString",
				"testString",
			)

			apiNote, response, err := findingsService.GetOccurrenceNote(getOccurrenceNoteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiNote, "", "  ")
			fmt.Println(string(b))

			// end-getOccurrenceNote

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
		It(`CreateOccurrence request example`, func() {
			fmt.Println("\nCreateOccurrence() result:")
			// begin-createOccurrence

			createOccurrenceOptions := findingsService.NewCreateOccurrenceOptions(
				"testString",
				"testString",
				"testString",
				"FINDING",
				"testString",
			)

			apiOccurrence, response, err := findingsService.CreateOccurrence(createOccurrenceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiOccurrence, "", "  ")
			fmt.Println(string(b))

			// end-createOccurrence

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
		It(`ListOccurrences request example`, func() {
			fmt.Println("\nListOccurrences() result:")
			// begin-listOccurrences

			listOccurrencesOptions := findingsService.NewListOccurrencesOptions(
				"testString",
				"testString",
			)

			apiListOccurrencesResponse, response, err := findingsService.ListOccurrences(listOccurrencesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiListOccurrencesResponse, "", "  ")
			fmt.Println(string(b))

			// end-listOccurrences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
		It(`ListNoteOccurrences request example`, func() {
			fmt.Println("\nListNoteOccurrences() result:")
			// begin-listNoteOccurrences

			listNoteOccurrencesOptions := findingsService.NewListNoteOccurrencesOptions(
				"testString",
				"testString",
				"testString",
			)

			apiListNoteOccurrencesResponse, response, err := findingsService.ListNoteOccurrences(listNoteOccurrencesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiListNoteOccurrencesResponse, "", "  ")
			fmt.Println(string(b))

			// end-listNoteOccurrences

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNoteOccurrencesResponse).ToNot(BeNil())

		})
		It(`GetOccurrence request example`, func() {
			fmt.Println("\nGetOccurrence() result:")
			// begin-getOccurrence

			getOccurrenceOptions := findingsService.NewGetOccurrenceOptions(
				"testString",
				"testString",
				"testString",
			)

			apiListOccurrencesResponse, response, err := findingsService.GetOccurrence(getOccurrenceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiListOccurrencesResponse, "", "  ")
			fmt.Println(string(b))

			// end-getOccurrence

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
		It(`UpdateOccurrence request example`, func() {
			fmt.Println("\nUpdateOccurrence() result:")
			// begin-updateOccurrence

			updateOccurrenceOptions := findingsService.NewUpdateOccurrenceOptions(
				"testString",
				"testString",
				"testString",
				"testString",
				"FINDING",
				"testString",
			)

			apiOccurrence, response, err := findingsService.UpdateOccurrence(updateOccurrenceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiOccurrence, "", "  ")
			fmt.Println(string(b))

			// end-updateOccurrence

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
		It(`ListProviders request example`, func() {
			fmt.Println("\nListProviders() result:")
			// begin-listProviders

			listProvidersOptions := findingsService.NewListProvidersOptions(
				"testString",
			)

			apiListProvidersResponse, response, err := findingsService.ListProviders(listProvidersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiListProvidersResponse, "", "  ")
			fmt.Println(string(b))

			// end-listProviders

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListProvidersResponse).ToNot(BeNil())

		})
		It(`DeleteOccurrence request example`, func() {
			// begin-deleteOccurrence

			deleteOccurrenceOptions := findingsService.NewDeleteOccurrenceOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := findingsService.DeleteOccurrence(deleteOccurrenceOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteOccurrence
			fmt.Printf("\nDeleteOccurrence() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteNote request example`, func() {
			// begin-deleteNote

			deleteNoteOptions := findingsService.NewDeleteNoteOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := findingsService.DeleteNote(deleteNoteOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteNote
			fmt.Printf("\nDeleteNote() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

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

package findingsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/findingsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the findingsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`FindingsV1 Integration Tests`, func() {

	const externalConfigFile = "../findings_v1.env"

	var (
		err          error
		findingsService *findingsv1.FindingsV1
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
			config, err = core.GetServiceProperties(findingsv1.DefaultServiceName)
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

			findingsServiceOptions := &findingsv1.FindingsV1Options{}

			findingsService, err = findingsv1.NewFindingsV1UsingExternalConfig(findingsServiceOptions)

			Expect(err).To(BeNil())
			Expect(findingsService).ToNot(BeNil())
			Expect(findingsService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`PostGraph - query findings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostGraph(postGraphOptions *PostGraphOptions)`, func() {

			postGraphOptions := &findingsv1.PostGraphOptions{
				AccountID: core.StringPtr("testString"),
				Body: CreateMockReader("This is a mock file."),
				ContentType: core.StringPtr("application/json"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := findingsService.PostGraph(postGraphOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`CreateNote - Creates a new `Note``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID: core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			apiNoteRelatedURLModel := &findingsv1.APINoteRelatedURL{
				Label: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			findingTypeModel := &findingsv1.FindingType{
				Severity: core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
			}

			kpiTypeModel := &findingsv1.KpiType{
				AggregationType: core.StringPtr("SUM"),
			}

			valueTypeModel := &findingsv1.ValueTypeFindingCountValueType{
				Kind: core.StringPtr("FINDING_COUNT"),
				FindingNoteNames: []string{"testString"},
				Text: core.StringPtr("testString"),
			}

			cardElementModel := &findingsv1.CardElementTimeSeriesCardElement{
				Text: core.StringPtr("testString"),
				DefaultInterval: core.StringPtr("testString"),
				Kind: core.StringPtr("TIME_SERIES"),
				DefaultTimeRange: core.StringPtr("1d"),
				ValueTypes: []findingsv1.ValueTypeIntf{valueTypeModel},
			}

			cardModel := &findingsv1.Card{
				Section: core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
				Subtitle: core.StringPtr("testString"),
				Order: core.Int64Ptr(int64(1)),
				FindingNoteNames: []string{"testString"},
				RequiresConfiguration: core.BoolPtr(true),
				BadgeText: core.StringPtr("testString"),
				BadgeImage: core.StringPtr("testString"),
				Elements: []findingsv1.CardElementIntf{cardElementModel},
			}

			sectionModel := &findingsv1.Section{
				Title: core.StringPtr("testString"),
				Image: core.StringPtr("testString"),
			}

			createNoteOptions := &findingsv1.CreateNoteOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Kind: core.StringPtr("FINDING"),
				ID: core.StringPtr("testString"),
				ReportedBy: reporterModel,
				RelatedURL: []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel},
				ExpirationTime: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared: core.BoolPtr(true),
				Finding: findingTypeModel,
				Kpi: kpiTypeModel,
				Card: cardModel,
				Section: sectionModel,
				TransactionID: core.StringPtr("testString"),
			}

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`ListNotes - Lists all `Notes` for a given provider`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNotes(listNotesOptions *ListNotesOptions)`, func() {

			listNotesOptions := &findingsv1.ListNotesOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				PageSize: core.Int64Ptr(int64(2)),
				PageToken: core.StringPtr("testString"),
			}

			apiListNotesResponse, response, err := findingsService.ListNotes(listNotesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNotesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetNote - Returns the requested `Note``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNote(getNoteOptions *GetNoteOptions)`, func() {

			getNoteOptions := &findingsv1.GetNoteOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				NoteID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			apiNote, response, err := findingsService.GetNote(getNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`UpdateNote - Updates an existing `Note``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID: core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			apiNoteRelatedURLModel := &findingsv1.APINoteRelatedURL{
				Label: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			findingTypeModel := &findingsv1.FindingType{
				Severity: core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
			}

			kpiTypeModel := &findingsv1.KpiType{
				AggregationType: core.StringPtr("SUM"),
			}

			valueTypeModel := &findingsv1.ValueTypeFindingCountValueType{
				Kind: core.StringPtr("FINDING_COUNT"),
				FindingNoteNames: []string{"testString"},
				Text: core.StringPtr("testString"),
			}

			cardElementModel := &findingsv1.CardElementTimeSeriesCardElement{
				Text: core.StringPtr("testString"),
				DefaultInterval: core.StringPtr("testString"),
				Kind: core.StringPtr("TIME_SERIES"),
				DefaultTimeRange: core.StringPtr("1d"),
				ValueTypes: []findingsv1.ValueTypeIntf{valueTypeModel},
			}

			cardModel := &findingsv1.Card{
				Section: core.StringPtr("testString"),
				Title: core.StringPtr("testString"),
				Subtitle: core.StringPtr("testString"),
				Order: core.Int64Ptr(int64(1)),
				FindingNoteNames: []string{"testString"},
				RequiresConfiguration: core.BoolPtr(true),
				BadgeText: core.StringPtr("testString"),
				BadgeImage: core.StringPtr("testString"),
				Elements: []findingsv1.CardElementIntf{cardElementModel},
			}

			sectionModel := &findingsv1.Section{
				Title: core.StringPtr("testString"),
				Image: core.StringPtr("testString"),
			}

			updateNoteOptions := &findingsv1.UpdateNoteOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				NoteID: core.StringPtr("testString"),
				ShortDescription: core.StringPtr("testString"),
				LongDescription: core.StringPtr("testString"),
				Kind: core.StringPtr("FINDING"),
				ID: core.StringPtr("testString"),
				ReportedBy: reporterModel,
				RelatedURL: []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel},
				ExpirationTime: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared: core.BoolPtr(true),
				Finding: findingTypeModel,
				Kpi: kpiTypeModel,
				Card: cardModel,
				Section: sectionModel,
				TransactionID: core.StringPtr("testString"),
			}

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`GetOccurrenceNote - Gets the `Note` attached to the given `Occurrence``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOccurrenceNote(getOccurrenceNoteOptions *GetOccurrenceNoteOptions)`, func() {

			getOccurrenceNoteOptions := &findingsv1.GetOccurrenceNoteOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				OccurrenceID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			apiNote, response, err := findingsService.GetOccurrenceNote(getOccurrenceNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`CreateOccurrence - Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOccurrence(createOccurrenceOptions *CreateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region: core.StringPtr("testString"),
				ResourceCRN: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				ResourceType: core.StringPtr("testString"),
				ServiceCRN: core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				EnvironmentName: core.StringPtr("testString"),
				ComponentName: core.StringPtr("testString"),
				ToolchainID: core.StringPtr("testString"),
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			socketAddressModel := &findingsv1.SocketAddress{
				Address: core.StringPtr("testString"),
				Port: core.Int64Ptr(int64(38)),
			}

			networkConnectionModel := &findingsv1.NetworkConnection{
				Direction: core.StringPtr("testString"),
				Protocol: core.StringPtr("testString"),
				Client: socketAddressModel,
				Server: socketAddressModel,
			}

			dataTransferredModel := &findingsv1.DataTransferred{
				ClientBytes: core.Int64Ptr(int64(38)),
				ServerBytes: core.Int64Ptr(int64(38)),
				ClientPackets: core.Int64Ptr(int64(38)),
				ServerPackets: core.Int64Ptr(int64(38)),
			}

			findingModel := &findingsv1.Finding{
				Severity: core.StringPtr("LOW"),
				Certainty: core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
				NetworkConnection: networkConnectionModel,
				DataTransferred: dataTransferredModel,
			}

			kpiModel := &findingsv1.Kpi{
				Value: core.Float64Ptr(float64(72.5)),
				Total: core.Float64Ptr(float64(72.5)),
			}

			createOccurrenceOptions := &findingsv1.CreateOccurrenceOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				NoteName: core.StringPtr("testString"),
				Kind: core.StringPtr("FINDING"),
				ID: core.StringPtr("testString"),
				ResourceURL: core.StringPtr("testString"),
				Remediation: core.StringPtr("testString"),
				Context: contextModel,
				Finding: findingModel,
				Kpi: kpiModel,
				ReferenceData: map[string]interface{}{"anyKey": "anyValue"},
				ReplaceIfExists: core.BoolPtr(true),
				TransactionID: core.StringPtr("testString"),
			}

			apiOccurrence, response, err := findingsService.CreateOccurrence(createOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`ListOccurrences - Lists active `Occurrences` for a given provider matching the filters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOccurrences(listOccurrencesOptions *ListOccurrencesOptions)`, func() {

			listOccurrencesOptions := &findingsv1.ListOccurrencesOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				PageSize: core.Int64Ptr(int64(2)),
				PageToken: core.StringPtr("testString"),
			}

			apiListOccurrencesResponse, response, err := findingsService.ListOccurrences(listOccurrencesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`ListNoteOccurrences - Lists `Occurrences` referencing the specified `Note`. Use this method to get all occurrences referencing your `Note` across all your customer providers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNoteOccurrences(listNoteOccurrencesOptions *ListNoteOccurrencesOptions)`, func() {

			listNoteOccurrencesOptions := &findingsv1.ListNoteOccurrencesOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				NoteID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				PageSize: core.Int64Ptr(int64(2)),
				PageToken: core.StringPtr("testString"),
			}

			apiListNoteOccurrencesResponse, response, err := findingsService.ListNoteOccurrences(listNoteOccurrencesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNoteOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetOccurrence - Returns the requested `Occurrence``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOccurrence(getOccurrenceOptions *GetOccurrenceOptions)`, func() {

			getOccurrenceOptions := &findingsv1.GetOccurrenceOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				OccurrenceID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			apiListOccurrencesResponse, response, err := findingsService.GetOccurrence(getOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`UpdateOccurrence - Updates an existing `Occurrence``, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOccurrence(updateOccurrenceOptions *UpdateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region: core.StringPtr("testString"),
				ResourceCRN: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				ResourceName: core.StringPtr("testString"),
				ResourceType: core.StringPtr("testString"),
				ServiceCRN: core.StringPtr("testString"),
				ServiceName: core.StringPtr("testString"),
				EnvironmentName: core.StringPtr("testString"),
				ComponentName: core.StringPtr("testString"),
				ToolchainID: core.StringPtr("testString"),
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: core.StringPtr("testString"),
				URL: core.StringPtr("testString"),
			}

			socketAddressModel := &findingsv1.SocketAddress{
				Address: core.StringPtr("testString"),
				Port: core.Int64Ptr(int64(38)),
			}

			networkConnectionModel := &findingsv1.NetworkConnection{
				Direction: core.StringPtr("testString"),
				Protocol: core.StringPtr("testString"),
				Client: socketAddressModel,
				Server: socketAddressModel,
			}

			dataTransferredModel := &findingsv1.DataTransferred{
				ClientBytes: core.Int64Ptr(int64(38)),
				ServerBytes: core.Int64Ptr(int64(38)),
				ClientPackets: core.Int64Ptr(int64(38)),
				ServerPackets: core.Int64Ptr(int64(38)),
			}

			findingModel := &findingsv1.Finding{
				Severity: core.StringPtr("LOW"),
				Certainty: core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
				NetworkConnection: networkConnectionModel,
				DataTransferred: dataTransferredModel,
			}

			kpiModel := &findingsv1.Kpi{
				Value: core.Float64Ptr(float64(72.5)),
				Total: core.Float64Ptr(float64(72.5)),
			}

			updateOccurrenceOptions := &findingsv1.UpdateOccurrenceOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				OccurrenceID: core.StringPtr("testString"),
				NoteName: core.StringPtr("testString"),
				Kind: core.StringPtr("FINDING"),
				ID: core.StringPtr("testString"),
				ResourceURL: core.StringPtr("testString"),
				Remediation: core.StringPtr("testString"),
				Context: contextModel,
				Finding: findingModel,
				Kpi: kpiModel,
				ReferenceData: map[string]interface{}{"anyKey": "anyValue"},
				TransactionID: core.StringPtr("testString"),
			}

			apiOccurrence, response, err := findingsService.UpdateOccurrence(updateOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`ListProviders - Lists all `Providers` for a given account id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviders(listProvidersOptions *ListProvidersOptions)`, func() {

			listProvidersOptions := &findingsv1.ListProvidersOptions{
				AccountID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(2)),
				Skip: core.Int64Ptr(int64(38)),
				StartProviderID: core.StringPtr("testString"),
				EndProviderID: core.StringPtr("testString"),
			}

			apiListProvidersResponse, response, err := findingsService.ListProviders(listProvidersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListProvidersResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteOccurrence - Deletes the given `Occurrence` from the system`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOccurrence(deleteOccurrenceOptions *DeleteOccurrenceOptions)`, func() {

			deleteOccurrenceOptions := &findingsv1.DeleteOccurrenceOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				OccurrenceID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := findingsService.DeleteOccurrence(deleteOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteNote - Deletes the given `Note` from the system`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNote(deleteNoteOptions *DeleteNoteOptions)`, func() {

			deleteNoteOptions := &findingsv1.DeleteNoteOptions{
				AccountID: core.StringPtr("testString"),
				ProviderID: core.StringPtr("testString"),
				NoteID: core.StringPtr("testString"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := findingsService.DeleteNote(deleteNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})

//
// Utility functions are declared in the unit test file
//

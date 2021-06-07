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
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm/scc-go-sdk/findingsv1"
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

var accountID = os.Getenv("ACCOUNT_ID")
var providerID = os.Getenv("PROVIDER_ID")
var testString = "testString"
var identifier = os.Getenv("TRAVIS_JOB_ID")

var _ = Describe(`FindingsV1 Integration Tests`, func() {

	if identifier == "" {
		identifier = fmt.Sprintf("%d", time.Now().Unix())
	}

	if providerID == "" {
		providerID = "sdk-it"
	}

	const externalConfigFile = "../findings_v1.env"

	var (
		err             error
		findingsService *findingsv1.FindingsV1
		serviceURL      string
		config          map[string]string
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
			fmt.Printf("Running Integration Tests using AccountID: %s and ProviderID: %s\n", accountID, providerID)
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
				AccountID:   &accountID,
				Body:        CreateMockReader(`{notes{id}}`),
				ContentType: core.StringPtr("application/graphql"),
			}

			response, err := findingsService.PostGraph(postGraphOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`CreateNote - Creates a new 'Note' (FINDING)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			apiNoteRelatedURLModel := &findingsv1.APINoteRelatedURL{
				Label: &testString,
				URL:   &testString,
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: &testString,
				URL:   &testString,
			}

			findingTypeModel := &findingsv1.FindingType{
				Severity:  core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
			}

			createNoteOptions := &findingsv1.CreateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("FINDING"),
				ID:               core.StringPtr(fmt.Sprintf("finding-note-%s", identifier)),
				ReportedBy:       reporterModel,
				RelatedURL:       []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel},
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Finding:          findingTypeModel,
			}

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`CreateNote - Creates a new 'Note' (KPI)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			kpiTypeModel := &findingsv1.KpiType{
				AggregationType: core.StringPtr("SUM"),
			}

			createNoteOptions := &findingsv1.CreateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("KPI"),
				ID:               core.StringPtr(fmt.Sprintf("kpi-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Kpi:              kpiTypeModel,
			}

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`CreateNote - Creates a new 'Note' (CARD)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			valueTypeModel := &findingsv1.ValueTypeFindingCountValueType{
				Kind:             core.StringPtr("FINDING_COUNT"),
				FindingNoteNames: []string{testString},
				Text:             &testString,
			}

			cardElementModel := &findingsv1.CardElementTimeSeriesCardElement{
				Text:             &testString,
				DefaultInterval:  &testString,
				Kind:             core.StringPtr("TIME_SERIES"),
				DefaultTimeRange: core.StringPtr("1d"),
				ValueTypes:       []findingsv1.ValueTypeIntf{valueTypeModel},
			}

			cardModel := &findingsv1.Card{
				Section:               &testString,
				Title:                 &testString,
				Subtitle:              &testString,
				Order:                 core.Int64Ptr(int64(1)),
				FindingNoteNames:      []string{fmt.Sprintf("%s/providers/%s/notes/finding-note-%s", accountID, providerID, identifier)},
				RequiresConfiguration: core.BoolPtr(true),
				BadgeText:             &testString,
				BadgeImage:            &testString,
				Elements:              []findingsv1.CardElementIntf{cardElementModel},
			}

			createNoteOptions := &findingsv1.CreateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("CARD"),
				ID:               core.StringPtr(fmt.Sprintf("card-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Card:             cardModel,
			}

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`CreateNote - Creates a new 'Note' (SECTION)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNote(createNoteOptions *CreateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			sectionModel := &findingsv1.Section{
				Title: &testString,
				Image: &testString,
			}

			createNoteOptions := &findingsv1.CreateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("SECTION"),
				ID:               core.StringPtr(fmt.Sprintf("section-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Section:          sectionModel,
			}

			apiNote, response, err := findingsService.CreateNote(createNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`ListNotes - Lists all 'Notes' for a given provider`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNotes(listNotesOptions *ListNotesOptions)`, func() {

			listNotesOptions := &findingsv1.ListNotesOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
			}

			apiListNotesResponse, response, err := findingsService.ListNotes(listNotesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNotesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetNote - Returns the requested 'Note'`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNote(getNoteOptions *GetNoteOptions)`, func() {

			getNoteOptions := &findingsv1.GetNoteOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
				NoteID:     core.StringPtr(fmt.Sprintf("finding-note-%s", identifier)),
			}

			apiNote, response, err := findingsService.GetNote(getNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`UpdateNote - Updates an existing 'Note' (FINDING)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			apiNoteRelatedURLModel := &findingsv1.APINoteRelatedURL{
				Label: &testString,
				URL:   &testString,
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: &testString,
				URL:   &testString,
			}

			findingTypeModel := &findingsv1.FindingType{
				Severity:  core.StringPtr("LOW"),
				NextSteps: []findingsv1.RemediationStep{*remediationStepModel},
			}

			updateNoteOptions := &findingsv1.UpdateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				NoteID:           core.StringPtr(fmt.Sprintf("finding-note-%s", identifier)),
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("FINDING"),
				ID:               core.StringPtr(fmt.Sprintf("finding-note-%s", identifier)),
				ReportedBy:       reporterModel,
				RelatedURL:       []findingsv1.APINoteRelatedURL{*apiNoteRelatedURLModel},
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Finding:          findingTypeModel,
			}

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`UpdateNote - Updates an existing 'Note' (KPI)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			kpiTypeModel := &findingsv1.KpiType{
				AggregationType: core.StringPtr("SUM"),
			}

			updateNoteOptions := &findingsv1.UpdateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				NoteID:           core.StringPtr(fmt.Sprintf("kpi-note-%s", identifier)),
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("KPI"),
				ID:               core.StringPtr(fmt.Sprintf("kpi-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Kpi:              kpiTypeModel,
			}

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`UpdateNote - Updates an existing 'Note' (CARD)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			valueTypeModel := &findingsv1.ValueTypeFindingCountValueType{
				Kind:             core.StringPtr("FINDING_COUNT"),
				FindingNoteNames: []string{testString},
				Text:             &testString,
			}

			cardElementModel := &findingsv1.CardElementTimeSeriesCardElement{
				Text:             &testString,
				DefaultInterval:  &testString,
				Kind:             core.StringPtr("TIME_SERIES"),
				DefaultTimeRange: core.StringPtr("1d"),
				ValueTypes:       []findingsv1.ValueTypeIntf{valueTypeModel},
			}

			cardModel := &findingsv1.Card{
				Section:               &testString,
				Title:                 &testString,
				Subtitle:              &testString,
				FindingNoteNames:      []string{fmt.Sprintf("%s/providers/%s/notes/finding-note-%s", accountID, providerID, identifier)},
				RequiresConfiguration: core.BoolPtr(true),
				BadgeText:             &testString,
				BadgeImage:            &testString,
				Elements:              []findingsv1.CardElementIntf{cardElementModel},
			}

			updateNoteOptions := &findingsv1.UpdateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				NoteID:           core.StringPtr(fmt.Sprintf("card-note-%s", identifier)),
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("CARD"),
				ID:               core.StringPtr(fmt.Sprintf("card-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Card:             cardModel,
			}

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`UpdateNote - Updates an existing 'Note' (SECTION)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNote(updateNoteOptions *UpdateNoteOptions)`, func() {

			reporterModel := &findingsv1.Reporter{
				ID:    &testString,
				Title: &testString,
				URL:   &testString,
			}

			sectionModel := &findingsv1.Section{
				Title: &testString,
				Image: &testString,
			}

			updateNoteOptions := &findingsv1.UpdateNoteOptions{
				AccountID:        &accountID,
				ProviderID:       &providerID,
				NoteID:           core.StringPtr(fmt.Sprintf("section-note-%s", identifier)),
				ShortDescription: &testString,
				LongDescription:  &testString,
				Kind:             core.StringPtr("SECTION"),
				ID:               core.StringPtr(fmt.Sprintf("section-note-%s", identifier)),
				ReportedBy:       reporterModel,
				ExpirationTime:   CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Shared:           core.BoolPtr(true),
				Section:          sectionModel,
			}

			apiNote, response, err := findingsService.UpdateNote(updateNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`CreateOccurrence - Creates a new 'Occurrence' (FINDING). Use this method to create 'Occurrences' for a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOccurrence(createOccurrenceOptions *CreateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region:          &testString,
				ResourceCRN:     &testString,
				ResourceID:      &testString,
				ResourceName:    &testString,
				ResourceType:    &testString,
				ServiceCRN:      &testString,
				ServiceName:     &testString,
				EnvironmentName: &testString,
				ComponentName:   &testString,
				ToolchainID:     &testString,
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: &testString,
				URL:   &testString,
			}

			socketAddressModel := &findingsv1.SocketAddress{
				Address: &testString,
				Port:    core.Int64Ptr(int64(38)),
			}

			networkConnectionModel := &findingsv1.NetworkConnection{
				Direction: &testString,
				Protocol:  &testString,
				Client:    socketAddressModel,
				Server:    socketAddressModel,
			}

			dataTransferredModel := &findingsv1.DataTransferred{
				ClientBytes:   core.Int64Ptr(int64(38)),
				ServerBytes:   core.Int64Ptr(int64(38)),
				ClientPackets: core.Int64Ptr(int64(38)),
				ServerPackets: core.Int64Ptr(int64(38)),
			}

			findingModel := &findingsv1.Finding{
				Severity:          core.StringPtr("LOW"),
				Certainty:         core.StringPtr("LOW"),
				NextSteps:         []findingsv1.RemediationStep{*remediationStepModel},
				NetworkConnection: networkConnectionModel,
				DataTransferred:   dataTransferredModel,
			}

			createOccurrenceOptions := &findingsv1.CreateOccurrenceOptions{
				AccountID:       &accountID,
				ProviderID:      &providerID,
				NoteName:        core.StringPtr(fmt.Sprintf("%s/providers/%s/notes/finding-note-%s", accountID, providerID, identifier)),
				Kind:            core.StringPtr("FINDING"),
				ID:              core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
				ResourceURL:     &testString,
				Remediation:     &testString,
				Context:         contextModel,
				Finding:         findingModel,
				ReplaceIfExists: core.BoolPtr(true),
			}

			apiOccurrence, response, err := findingsService.CreateOccurrence(createOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`CreateOccurrence - Creates a new 'Occurrence' (KPI). Use this method to create 'Occurrences' for a resource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOccurrence(createOccurrenceOptions *CreateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region:          &testString,
				ResourceCRN:     &testString,
				ResourceID:      &testString,
				ResourceName:    &testString,
				ResourceType:    &testString,
				ServiceCRN:      &testString,
				ServiceName:     &testString,
				EnvironmentName: &testString,
				ComponentName:   &testString,
				ToolchainID:     &testString,
			}

			kpiModel := &findingsv1.Kpi{
				Value: core.Float64Ptr(float64(72.5)),
				Total: core.Float64Ptr(float64(72.5)),
			}

			createOccurrenceOptions := &findingsv1.CreateOccurrenceOptions{
				AccountID:       &accountID,
				ProviderID:      &providerID,
				NoteName:        core.StringPtr(fmt.Sprintf("%s/providers/%s/notes/kpi-note-%s", accountID, providerID, identifier)),
				Kind:            core.StringPtr("KPI"),
				ID:              core.StringPtr(fmt.Sprintf("kpi-occurrence-%s", identifier)),
				ResourceURL:     &testString,
				Remediation:     &testString,
				Context:         contextModel,
				Kpi:             kpiModel,
				ReplaceIfExists: core.BoolPtr(true),
			}

			apiOccurrence, response, err := findingsService.CreateOccurrence(createOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`GetOccurrenceNote - Gets the 'Note' attached to the given 'Occurrence'`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOccurrenceNote(getOccurrenceNoteOptions *GetOccurrenceNoteOptions)`, func() {

			getOccurrenceNoteOptions := &findingsv1.GetOccurrenceNoteOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
			}

			apiNote, response, err := findingsService.GetOccurrenceNote(getOccurrenceNoteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiNote).ToNot(BeNil())

		})
	})

	Describe(`ListOccurrences - Lists active 'Occurrences' for a given provider matching the filters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOccurrences(listOccurrencesOptions *ListOccurrencesOptions)`, func() {

			listOccurrencesOptions := &findingsv1.ListOccurrencesOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
			}

			apiListOccurrencesResponse, response, err := findingsService.ListOccurrences(listOccurrencesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`ListNoteOccurrences - Lists 'Occurrences' referencing the specified 'Note'. Use this method to get all occurrences referencing your 'Note' across all your customer providers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNoteOccurrences(listNoteOccurrencesOptions *ListNoteOccurrencesOptions)`, func() {

			listNoteOccurrencesOptions := &findingsv1.ListNoteOccurrencesOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
				NoteID:     core.StringPtr(fmt.Sprintf("finding-note-%s", identifier)),
			}

			apiListNoteOccurrencesResponse, response, err := findingsService.ListNoteOccurrences(listNoteOccurrencesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListNoteOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetOccurrence - Returns the requested 'Occurrence'`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOccurrence(getOccurrenceOptions *GetOccurrenceOptions)`, func() {

			getOccurrenceOptions := &findingsv1.GetOccurrenceOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
			}

			apiListOccurrencesResponse, response, err := findingsService.GetOccurrence(getOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListOccurrencesResponse).ToNot(BeNil())

		})
	})

	Describe(`UpdateOccurrence - Updates an existing 'Occurrence' (FINDING)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOccurrence(updateOccurrenceOptions *UpdateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region:          &testString,
				ResourceCRN:     &testString,
				ResourceID:      &testString,
				ResourceName:    &testString,
				ResourceType:    &testString,
				ServiceCRN:      &testString,
				ServiceName:     &testString,
				EnvironmentName: &testString,
				ComponentName:   &testString,
				ToolchainID:     &testString,
			}

			remediationStepModel := &findingsv1.RemediationStep{
				Title: &testString,
				URL:   &testString,
			}

			socketAddressModel := &findingsv1.SocketAddress{
				Address: &testString,
				Port:    core.Int64Ptr(int64(38)),
			}

			networkConnectionModel := &findingsv1.NetworkConnection{
				Direction: &testString,
				Protocol:  &testString,
				Client:    socketAddressModel,
				Server:    socketAddressModel,
			}

			dataTransferredModel := &findingsv1.DataTransferred{
				ClientBytes:   core.Int64Ptr(int64(38)),
				ServerBytes:   core.Int64Ptr(int64(38)),
				ClientPackets: core.Int64Ptr(int64(38)),
				ServerPackets: core.Int64Ptr(int64(38)),
			}

			findingModel := &findingsv1.Finding{
				Severity:          core.StringPtr("LOW"),
				Certainty:         core.StringPtr("LOW"),
				NextSteps:         []findingsv1.RemediationStep{*remediationStepModel},
				NetworkConnection: networkConnectionModel,
				DataTransferred:   dataTransferredModel,
			}

			updateOccurrenceOptions := &findingsv1.UpdateOccurrenceOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
				NoteName:     core.StringPtr(fmt.Sprintf("%s/providers/%s/notes/finding-note-%s", accountID, providerID, identifier)),
				Kind:         core.StringPtr("FINDING"),
				ID:           core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
				ResourceURL:  &testString,
				Remediation:  &testString,
				Context:      contextModel,
				Finding:      findingModel,
			}

			apiOccurrence, response, err := findingsService.UpdateOccurrence(updateOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`UpdateOccurrence - Updates an existing 'Occurrence' (KPI)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOccurrence(updateOccurrenceOptions *UpdateOccurrenceOptions)`, func() {

			contextModel := &findingsv1.Context{
				Region:          &testString,
				ResourceCRN:     &testString,
				ResourceID:      &testString,
				ResourceName:    &testString,
				ResourceType:    &testString,
				ServiceCRN:      &testString,
				ServiceName:     &testString,
				EnvironmentName: &testString,
				ComponentName:   &testString,
				ToolchainID:     &testString,
			}

			kpiModel := &findingsv1.Kpi{
				Value: core.Float64Ptr(float64(72.5)),
				Total: core.Float64Ptr(float64(72.5)),
			}

			updateOccurrenceOptions := &findingsv1.UpdateOccurrenceOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: core.StringPtr(fmt.Sprintf("kpi-occurrence-%s", identifier)),
				NoteName:     core.StringPtr(fmt.Sprintf("%s/providers/%s/notes/kpi-note-%s", accountID, providerID, identifier)),
				Kind:         core.StringPtr("KPI"),
				ID:           core.StringPtr(fmt.Sprintf("kpi-occurrence-%s", identifier)),
				ResourceURL:  &testString,
				Remediation:  &testString,
				Context:      contextModel,
				Kpi:          kpiModel,
			}

			apiOccurrence, response, err := findingsService.UpdateOccurrence(updateOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiOccurrence).ToNot(BeNil())

		})
	})

	Describe(`ListProviders - Lists all 'Providers' for a given account id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProviders(listProvidersOptions *ListProvidersOptions)`, func() {

			listProvidersOptions := &findingsv1.ListProvidersOptions{
				AccountID: &accountID,
			}

			apiListProvidersResponse, response, err := findingsService.ListProviders(listProvidersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiListProvidersResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteOccurrence - Deletes the given 'Occurrence' from the system`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOccurrence(deleteOccurrenceOptions *DeleteOccurrenceOptions)`, func() {

			deleteOccurrenceOptions := &findingsv1.DeleteOccurrenceOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: core.StringPtr(fmt.Sprintf("finding-occurrence-%s", identifier)),
			}

			response, err := findingsService.DeleteOccurrence(deleteOccurrenceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`DeleteNote - Deletes the given 'Note' from the system`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNote(deleteNoteOptions *DeleteNoteOptions)`, func() {

			deleteNoteOptions := &findingsv1.DeleteNoteOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
				NoteID:     core.StringPtr(fmt.Sprintf("section-note-%s", identifier)),
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

// Cleanup: Delete all notes and occurrences created during testing
var _ = AfterSuite(func() {
	const externalConfigFile = "../findings_v1.env"
	_, err := os.Stat(externalConfigFile)
	if err != nil {
		Skip("External configuration file not found, skipping tests: " + err.Error())
	}
	os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
	config, err := core.GetServiceProperties(findingsv1.DefaultServiceName)
	if err != nil {
		Skip("Error loading service properties, skipping tests: " + err.Error())
	}
	serviceURL := config["URL"]
	if serviceURL == "" {
		Skip("Unable to load service URL configuration property, skipping tests")
	}

	fmt.Printf("cleaning up account: %s with provider: %s\n", accountID, providerID)

	findingsServiceOptions := &findingsv1.FindingsV1Options{}

	findingsService, err := findingsv1.NewFindingsV1UsingExternalConfig(findingsServiceOptions)

	// list notes and delete one by one
	listNotesOptions := &findingsv1.ListNotesOptions{
		AccountID:  &accountID,
		ProviderID: &providerID,
	}

	apiListNotesResponse, _, err := findingsService.ListNotes(listNotesOptions)
	if err != nil {
		Skip("Error occurred while listing notes for cleanup: " + err.Error())
	}

	for _, note := range apiListNotesResponse.Notes {
		parts := strings.Split(*note.ID, "-")
		if parts[len(parts)-1] == identifier {
			deleteNoteOptions := &findingsv1.DeleteNoteOptions{
				AccountID:  &accountID,
				ProviderID: &providerID,
				NoteID:     note.ID,
			}
			_, err := findingsService.DeleteNote(deleteNoteOptions)
			if err != nil {
				Skip("Error occurred while deleting note for cleanup: " + err.Error())
			}
		}
	}

	// list occurrences and delete one by one
	listOccurrencesOptions := &findingsv1.ListOccurrencesOptions{
		AccountID:  &accountID,
		ProviderID: &providerID,
	}

	apiListOccurrencesResponse, _, err := findingsService.ListOccurrences(listOccurrencesOptions)
	if err != nil {
		Skip("Error occurred while listing occurrences for cleanup: " + err.Error())
	}

	for _, occurrence := range apiListOccurrencesResponse.Occurrences {
		parts := strings.Split(*occurrence.ID, "-")
		if parts[len(parts)-1] == identifier {
			deleteOccurrenceOptions := &findingsv1.DeleteOccurrenceOptions{
				AccountID:    &accountID,
				ProviderID:   &providerID,
				OccurrenceID: occurrence.ID,
			}
			_, err := findingsService.DeleteOccurrence(deleteOccurrenceOptions)
			if err != nil {
				Skip("Error occurred while deleting occurrence for cleanup: " + err.Error())
			}
		}
	}

	fmt.Printf("cleanup was successful\n")

	// cross checking if the provider is there or not
	listProvidersOptions := &findingsv1.ListProvidersOptions{
		AccountID: &accountID,
	}

	apiListProvidersResponse, _, err := findingsService.ListProviders(listProvidersOptions)
	for _, provider := range apiListProvidersResponse.Providers {
		if *provider.ID == providerID {
			fmt.Printf("seems like account has some resources left even after a successful cleanup, please consider manual cleanup for account: %s and provider: %s\n", accountID, providerID)
		}
	}
})

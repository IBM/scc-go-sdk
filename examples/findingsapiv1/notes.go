package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/findingsapiv1"
)

//CreateFindingNote creates a note of kind FINDING
func CreateFindingNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "exampleProvider"
	shortDescription := "exampleString"
	longDescription := "exampleString"
	kind := "FINDING"
	id := "exampleNote"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	nextStep := []findingsapiv1.RemediationStep{{Title: core.StringPtr("title"), URL: core.StringPtr("https://hello.world")}}
	finding := findingsapiv1.FindingType{Severity: core.StringPtr("CRITICAL"), NextSteps: nextStep}

	var createNoteOptions = service.NewCreateNoteOptions(accountID, providerID, shortDescription, longDescription, kind, id, reportedBy)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetFinding(&finding) //finding is required for type FINDING

	result, response, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		fmt.Println(operationErr)
		fmt.Println("Failed to create note: ", operationErr)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
		fmt.Println(*result.Kind)
		fmt.Println(*result.ShortDescription)
	}
}

//CreateKPINote creates a note of kind KPI
func CreateKPINote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "exampleProvider"
	shortDescription := "exampleString"
	longDescription := "exampleString"
	kind := "KPI"
	id := "examplekpiNote"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	kpi, _ := service.NewKpiType("SUM")

	var createNoteOptions = service.NewCreateNoteOptions(accountID, providerID, shortDescription, longDescription, kind, id, reportedBy)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetKpi(kpi)

	result, response, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		fmt.Println("Failed to create note: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
		fmt.Println(*result.Kind)
		fmt.Println(*result.ShortDescription)
	}

}

//CreateSectionNote creates a note of kind SECTION
func CreateSectionNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "exampleProvider"
	shortDescription := "exampleString"
	longDescription := "exampleString"
	kind := "SECTION"
	id := "exampleSectionNote"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	section, _ := service.NewSection("Test Section", "s1img")

	var createNoteOptions = service.NewCreateNoteOptions(accountID, providerID, shortDescription, longDescription, kind, id, reportedBy)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetSection(section)

	result, response, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		fmt.Println("Failed to create note: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
		fmt.Println(*result.Kind)
		fmt.Println(*result.ShortDescription)
	}

}

//GetNote gets a note
func GetNote() {
	providerID := "custom-provider"
	noteID := "custom-note"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	getNotesOptions := service.NewGetNoteOptions(accountID, providerID, noteID)

	result, response, err := service.GetNote(getNotesOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(response.Result)
		return
	}

	fmt.Println(response.Result)
	fmt.Println(*result.Kind)
	fmt.Println(*result.ID)

}

//CreateCardNote creates a note of king CARD
func CreateCardNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "exampleProvider"
	shortDescription := "exampleString"
	longDescription := "exampleString"
	kind := "CARD"
	id := "exampleCardNote"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	cardValueKind := "FINDING_COUNT"
	cardElementKind := "NUMERIC"
	cardElementText := "text"
	cardValueType := findingsapiv1.CardValueType{Kind: &cardValueKind, FindingNoteNames: []string{"providers/sdktest/notes/sdk_note_id1"}}
	cardElement := []findingsapiv1.CardElement{{
		Kind:      &cardElementKind,
		Text:      &cardElementText,
		ValueType: &cardValueType, //ValueType required for kind NUMERIC
	}}
	card, _ := service.NewCard("My Security Tools", "Card posted from Go SDK", "subtitle", []string{"providers/sdktest/notes/sdk_note_id1"}, cardElement)

	var createNoteOptions = service.NewCreateNoteOptions(accountID, providerID, shortDescription, longDescription, kind, id, reportedBy)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetCard(card)

	result, response, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		fmt.Println("Failed to create note: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
		fmt.Println(*result.Kind)
		fmt.Println(*result.ShortDescription)
	}

}

//ListNotes Lists notes under a provider
func ListNotes() {
	providerID := "custom-provider"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
	})

	listNotesOptions := service.NewListNotesOptions(accountID, providerID)
	listNotesOptions.SetPageSize(2)

	listNotesResult, listNotesResponse, err := service.ListNotes(listNotesOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(listNotesResponse.Result)
		return
	}

	fmt.Println(listNotesResponse.Result)
	fmt.Println("Result: ", listNotesResult.Notes[0])
}

//UpdateNote Updates a note
func UpdateNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
	})

	providerID := "custom-provider"
	noteID := accountID + "/providers/custom-provider/notes/custom-note"
	shortDescription := "hello 3rd world"
	longDescription := "hello world"
	kind := "FINDING"
	id := "custom-note"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	remediationTitle := "title"
	remediationURL := "https://hello.world"
	nextStep := []findingsapiv1.RemediationStep{{Title: &remediationTitle, URL: &remediationURL}}
	severity := "MEDIUM"
	finding := findingsapiv1.FindingType{Severity: &severity, NextSteps: nextStep}

	var updateNoteOptions = service.NewUpdateNoteOptions(accountID, providerID, noteID, shortDescription, longDescription, kind, id, reportedBy)
	updateNoteOptions.SetHeaders(headers)
	updateNoteOptions.SetAccountID(accountID)
	updateNoteOptions.SetFinding(&finding) //finding is required for type FINDING

	result, response, operationErr := service.UpdateNote(updateNoteOptions)
	if operationErr != nil {
		fmt.Println("Failed to edit note: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(*result.ShortDescription)
		fmt.Println(result)
		fmt.Println(response.StatusCode)

	}

}

//DeleteNote deletes a note
func DeleteNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "custom-provider"
	noteID := "custom-note-kpi"
	var deleteOptions = service.NewDeleteNoteOptions(accountID, providerID, noteID)
	response, err := service.DeleteNote(deleteOptions)
	if err != nil {
		fmt.Println("Failed to delete newly created note with error: ", err)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
	}

}

//GetOccurrenceNote  gets details of the note associated with the specified occurrence
func GetOccurrenceNote() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "custom-provider"
	occID := "1234567"
	var getOccurrenceNoteOptions = service.NewGetOccurrenceNoteOptions(accountID, providerID, occID)

	res, response, err := service.GetOccurrenceNote(getOccurrenceNoteOptions)
	if err != nil {
		fmt.Println("Failed to GetOccurrenceNote, err: ", err)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*res.ID)
		fmt.Println(*res.Kind)
		fmt.Println(*res.ShortDescription)
	}

}

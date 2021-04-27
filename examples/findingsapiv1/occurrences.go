package examples

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/ibm-cloud-security/scc-go-sdk/findingsv1"
)

//CreateFindingOccurrence creates a new occurrence of type finding
func CreateFindingOccurrence() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "custom-provider"
	ID := "test-finding"
	noteName := accountID + "/providers/custom-provider/notes/custom-note"
	kind := "FINDING"
	nextStep := []findingsv1.RemediationStep{{Title: core.StringPtr("title"), URL: core.StringPtr("https://hello.world")}}
	finding := findingsv1.Finding{Severity: core.StringPtr("CRITICAL"), Certainty: core.StringPtr("LOW"), NextSteps: nextStep}
	context := findingsv1.Context{Region: core.StringPtr("us-south"), ResourceType: core.StringPtr("my_cluster"), ResourceName: core.StringPtr("test")}

	var createOccurrenceOptions = service.NewCreateOccurrenceOptions(accountID, providerID, noteName, kind, ID)
	createOccurrenceOptions.SetHeaders(headers)
	createOccurrenceOptions.SetAccountID(accountID)
	createOccurrenceOptions.SetFinding(&finding) //finding field is required for type "FINDING"
	createOccurrenceOptions.SetContext(&context)

	result, response, operationErr := service.CreateOccurrence(createOccurrenceOptions)
	if operationErr != nil {
		fmt.Println("Failed to create occurrence: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
	}

}

//CreateKPIOccurrence creates a new occurrence of type kpi
func CreateKPIOccurrence() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings", //Specify url or use default
	})

	providerID := "custom-provider"
	ID := "test-kpi"
	noteName := accountID + "/providers/custom-provider/notes/custom-note-kpi"
	kind := "KPI"
	kpi, _ := service.NewKpi(2.0)

	var createOccurrenceOptions = service.NewCreateOccurrenceOptions(accountID, providerID, noteName, kind, ID)
	createOccurrenceOptions.SetHeaders(headers)
	createOccurrenceOptions.SetKpi(kpi)

	result, response, operationErr := service.CreateOccurrence(createOccurrenceOptions)
	if operationErr != nil {
		fmt.Println("Failed to create occurrence: ", operationErr)
		fmt.Println(response.Result)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(*result.ID)
	}

}

//UpdateOccurrence Updates an existing occurrence
func UpdateOccurrence() {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
		URL:           "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings",
	})

	providerID := "custom-provider"
	ID := "test-kpi"
	occID := accountID + "/providers/custom-provider/occurrences/test-kpi"
	noteName := accountID + "/providers/custom-provider/notes/custom-note-kpi"
	kind := "KPI"
	kpiValue := 3.0
	kpiTotal := 3.0
	kpi := findingsv1.Kpi{Value: &kpiValue, Total: &kpiTotal}

	var updateOccurrenceOptions = service.NewUpdateOccurrenceOptions(accountID, providerID, occID, noteName, kind, ID)
	updateOccurrenceOptions.SetHeaders(headers)
	updateOccurrenceOptions.SetKpi(&kpi) //Kpi is required only for type "KPI"

	result, response, operationErr := service.UpdateOccurrence(updateOccurrenceOptions)
	if operationErr != nil {
		fmt.Println("Failed to edit occurrence: ", operationErr)
		fmt.Println(response.Result)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Kpi.Value)

}

//GetOccurrence gets details of an occurrence by occurrence id
func GetOccurrence() {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
		URL:           "https://us-south.secadvisor.cloud.ibm.com/findings",
	})

	providerID := "sdktest"
	occID := "12345"

	getOccurrenceOptions := service.NewGetOccurrenceOptions(accountID, providerID, occID)
	res, detailedResponse, err := service.GetOccurrence(getOccurrenceOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(detailedResponse.Result)
	}

	fmt.Println(res)
	fmt.Println(detailedResponse.Result) //type interface {}
	fmt.Println(*res.Occurrences[0].ID)

}

//DeleteOccurrence deletes a occurrence by id
func DeleteOccurrence() {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
	})

	providerID := "custom-provider"
	occID := "test-kpi"

	var deleteOptions = service.NewDeleteOccurrenceOptions(accountID, providerID, occID)
	response, err := service.DeleteOccurrence(deleteOptions)
	if err != nil || response.StatusCode != 200 {
		fmt.Println("Failed to delete occurrence with error: ", err)
		fmt.Println(response.Result)
	}

	fmt.Println(response.StatusCode)

}

//GetNoteOccurrences get all occurrences associated with a note. supports paging
func GetNoteOccurrences() {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
	})

	providerID := "custom-provider"
	noteID := "custom-note"

	listNoteOccurrencesOptions := service.NewListNoteOccurrencesOptions(accountID, providerID, noteID)
	listNoteOccurrencesOptions.SetPageSize(2)
	res, detailedResponse, err := service.ListNoteOccurrences(listNoteOccurrencesOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(detailedResponse.Result)
	}

	fmt.Println(len(res.Occurrences))
	fmt.Println(*res.Occurrences[0].ID)
	fmt.Println(*res.Occurrences[1].ID)
	// fmt.Println(*res.NextPageToken)
	fmt.Println(detailedResponse.StatusCode)

}

//GetProviderOccurrences gets all occurrences associated with a provider
func GetProviderOccurrences() {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    url, //use for dev/preprod env
	}
	service, _ := findingsv1.NewFindingsV1(&findingsv1.FindingsV1Options{
		Authenticator: authenticator,
	})

	providerID := "custom-provider"

	listProviderOccurrencesOptions := service.NewListOccurrencesOptions(accountID, providerID)
	listProviderOccurrencesOptions.SetPageSize(2)
	res, detailedResponse, err := service.ListOccurrences(listProviderOccurrencesOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(detailedResponse.Result)
	}

	fmt.Println(len(res.Occurrences))
	fmt.Println(*res.Occurrences[0].ID)
	fmt.Println(*res.Occurrences[1].ID)
	// fmt.Println(*res.NextPageToken)
	fmt.Println(detailedResponse.StatusCode)

}

// +build integration

package findingsapiv1_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/ibm-cloud-security/scc-go-sdk/findingsapiv1"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var apiKey = os.Getenv("apiKey")
var accountID = os.Getenv("accountID")
var URL = os.Getenv("URL")
var findingsServiceURL = os.Getenv("findingsServiceURL")
var inputFilePath = "../testInput/json"

var (
	service     *findingsapiv1.FindingsApiV1
	shouldSkip  bool = false
	err         error
	errConflict error = errors.New("Conflict")
)

var inputEnvPath = "../testInput/env"

func shouldSkipTest(t *testing.T) {
	if !shouldSkip {
		t.Skip("External configuration is not available, skipping...")
	}
}

func createNoteHelper(t *testing.T, path string) (result *findingsapiv1.ApiNote, options *findingsapiv1.CreateNoteOptions) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	query, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	var createNoteOptions *findingsapiv1.CreateNoteOptions
	now := time.Now()
	timeStamp := now.UnixNano()
	json.Unmarshal([]byte(query), &createNoteOptions)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetAccountID(accountID)
	createNoteOptions.SetID(strconv.FormatInt(timeStamp, 10))
	result, _, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		t.Log("Failed to create note: ", operationErr)
	}

	return result, createNoteOptions

}

func createOccurrenceHelper(t *testing.T, path string, noteID string) (*findingsapiv1.ApiOccurrence, *findingsapiv1.CreateOccurrenceOptions) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	query, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	var createOccurrenceOptions *findingsapiv1.CreateOccurrenceOptions
	now := time.Now()
	timeStamp := now.UnixNano()
	json.Unmarshal([]byte(query), &createOccurrenceOptions)
	createOccurrenceOptions.SetHeaders(headers)
	createOccurrenceOptions.SetAccountID(accountID)
	createOccurrenceOptions.SetID(strconv.FormatInt(timeStamp, 10))
	noteName := accountID + "/providers/" + *createOccurrenceOptions.ProviderID + "/notes/" + noteID
	createOccurrenceOptions.NoteName = &noteName

	result, _, operationErr := service.CreateOccurrence(createOccurrenceOptions)
	if operationErr != nil {
		t.Log("Failed to create occurrence: ", operationErr)
	}

	return result, createOccurrenceOptions

}

func deleteNoteHelper(t *testing.T, createNoteOptions *findingsapiv1.CreateNoteOptions) {
	var deleteOptions = service.NewDeleteNoteOptions(*(createNoteOptions.AccountID), *(createNoteOptions.ProviderID), *(createNoteOptions.ID))
	response, err := service.DeleteNote(deleteOptions)
	if err != nil || response.StatusCode != 200 {
		t.Fatal("Failed to delete newly created note with error: ", err)
	}
}

func deleteOccurrenceHelper(t *testing.T, createOccurrenceOptions *findingsapiv1.CreateOccurrenceOptions) {
	var deleteOptions = service.NewDeleteOccurrenceOptions(accountID, *createOccurrenceOptions.ProviderID, *createOccurrenceOptions.ID)
	response, err := service.DeleteOccurrence(deleteOptions)
	if err != nil || response.StatusCode != 200 {
		t.Fatal("Failed to delete newly created occurrence with error: ", err)
	}
}

func TestServiceSetupWithExternalConfig(t *testing.T) {
	externalConfigFile := inputEnvPath + "/findingsapiv1.env"
	fmt.Println("Starting to Test Service Setup With External Config")
	err = godotenv.Load(externalConfigFile)
	if err != nil {
		fmt.Println("Failed to load env vars: ", err)
	}

	testService, err := findingsapiv1.NewFindingsApiV1UsingExternalConfig(&findingsapiv1.FindingsApiV1Options{})
	if err != nil {
		fmt.Println("Failed to load external config: ", err)
	}

	assert.NotNil(t, testService)
}

func TestServiceSetupAndSetUrl(t *testing.T) {
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    URL,
	}
	testService, err := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
	})
	if testService == nil {
		t.Fatal("Expected service to not be nil, but got: ", err)
	}

	testService.SetServiceURL("https://dev-dallas.secadvisor.test.cloud.ibm.com/findings")

	assert.Equal(t, testService.Service.Options.URL, "https://dev-dallas.secadvisor.test.cloud.ibm.com/findings")
	if err != nil {
		t.Fatal("expected testServiceErr to be nil, but got: ", err)
	}

}

func TestServiceSetup(t *testing.T) {
	fmt.Println("Creating findings api service instance...")
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
		URL:    URL,
	}
	service, err = findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
		Authenticator: authenticator,
		URL:           findingsServiceURL,
	})
	if service == nil {
		t.Fatal("Expected service to not be nil, but got: ", err)
	}

}

func TestPostGraph(t *testing.T) {
	fmt.Println("creating new note")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")
	fmt.Println("creating new occurrence")
	_, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/graphql"

	query, err := os.Open(inputFilePath + "/findingCount.graphql.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()

	//Test post graph with nil options
	res, operationErr := service.PostGraph(nil)
	assert.NotNil(t, operationErr)
	assert.Nil(t, res)

	//Test post graph without required options
	var options findingsapiv1.PostGraphOptions
	res, operationErr = service.PostGraph(&options)
	assert.NotNil(t, operationErr)
	assert.Nil(t, res)

	//Test post graph with porper data
	fmt.Println("Posting graph query...")
	postGraphOptions := service.NewPostGraphOptions(accountID)
	postGraphOptions.SetAccountID(accountID)
	postGraphOptions.SetBody(query)
	postGraphOptions.SetHeaders(headers)
	postGraphOptions.SetContentType("application/graphql")
	res, operationErr = service.PostGraph(postGraphOptions)
	if operationErr != nil {
		fmt.Println("Err", operationErr)
	}
	fmt.Println("Posted graph query...")

	findingCount := res.Result.(map[string]interface{})["data"].(map[string]interface{})["findingCount"]
	findingCount = int(findingCount.(float64))

	assert.GreaterOrEqual(t, findingCount, 1)

	assert.Nil(t, operationErr)
	assert.NotNil(t, res)

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)
	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestListProvider(t *testing.T) {
	fmt.Println("creating new note")
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	res, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")
	if res == nil {
		t.Fatal("failed to create note")
	} else {
		//Test List Providers with nil options
		res, _, err := service.ListProviders(nil)
		assert.NotNil(t, err)
		assert.Nil(t, res)

		//Test List Providers without required options
		var options findingsapiv1.ListProvidersOptions
		res, _, err = service.ListProviders(&options)
		assert.NotNil(t, err)
		assert.Nil(t, res)

		var listProvidersOptions = service.NewListProvidersOptions(accountID)
		listProvidersOptions.SetHeaders(headers)
		listProvidersOptions.SetAccountID(accountID)
		listProvidersOptions.SetLimit(2)
		listProvidersOptions.SetSkip(0)
		listProvidersOptions.SetStartProviderID("sec_")
		listProvidersOptions.SetEndProviderID("sf")

		res, _, err = service.ListProviders(listProvidersOptions)
		if err != nil {
			t.Fatal("Failed to get list of providers: ", err)
		}

		assert.Equal(t, *res.Providers[0].ID, *createNoteOptions.ProviderID)

		fmt.Println("Cleaning up note....")
		deleteNoteHelper(t, createNoteOptions)
	}

}

func TestListProviderOccurrences(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test ListProviderOccurrence with nil options
	res, _, err := service.ListOccurrences(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test ListProviderOccurrence without required options
	var options findingsapiv1.ListOccurrencesOptions
	res, _, err = service.ListOccurrences(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("Creating note....")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/providerNote.json")
	fmt.Println("Creating occurrence....")
	_, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/providerOccurrence.json", *createNoteOptions.ID)

	//Test using required only
	fmt.Println("listing occurrences....")
	listProviderOccurrencesOptions := service.NewListOccurrencesOptions(accountID, *(createOccurrenceOptions.ProviderID))
	res, _, err = service.ListOccurrences(listProviderOccurrencesOptions)

	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, *(res.Occurrences[0].ID), *(createOccurrenceOptions.ID))
	assert.Equal(t, *(res.Occurrences[0].NoteName), accountID+"/providers/"+*createNoteOptions.ProviderID+"/notes/"+*createNoteOptions.ID)

	fmt.Println("listing occurrences using setter functions")
	var listOptions findingsapiv1.ListOccurrencesOptions
	listOptions.SetAccountID(accountID)
	listOptions.SetHeaders(headers)
	listOptions.SetPageToken(*res.NextPageToken)
	listOptions.SetProviderID(*(createOccurrenceOptions.ProviderID))
	listOptions.SetPageSize(10)

	res, _, err = service.ListOccurrences(&listOptions)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)

}

func TestListNoteOccurrences(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test ListProviderOccurrence with nil options
	res, _, err := service.ListNoteOccurrences(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test ListProviderOccurrence without required options
	var options findingsapiv1.ListNoteOccurrencesOptions
	res, _, err = service.ListNoteOccurrences(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("Creating note....")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("Creating occurrence....")
	_, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/noteOccurrence.json", *createNoteOptions.ID)

	fmt.Println("listing occurrences using required only....")
	listNoteOccurrencesOptions := service.NewListNoteOccurrencesOptions(accountID, *(createOccurrenceOptions.ProviderID), *(createNoteOptions.ID))
	res, _, err = service.ListNoteOccurrences(listNoteOccurrencesOptions)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, *(res.Occurrences[0].ID), *(createOccurrenceOptions.ID))
	assert.Equal(t, *(res.Occurrences[0].NoteName), accountID+"/providers/"+*createNoteOptions.ProviderID+"/notes/"+*createNoteOptions.ID)

	//List occurrecnts using setter functions
	var listOptions findingsapiv1.ListNoteOccurrencesOptions
	listOptions.SetAccountID(accountID)
	listOptions.SetHeaders(headers)
	listOptions.SetProviderID(*createOccurrenceOptions.ProviderID)
	listOptions.SetNoteID(*createNoteOptions.ID)
	listOptions.SetPageSize(10)
	listOptions.SetPageToken(*res.NextPageToken)
	res, _, err = service.ListNoteOccurrences(&listOptions)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)

}

func TestGetOccurrence(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test ListProviderOccurrence with nil options
	res, _, err := service.GetOccurrence(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test ListProviderOccurrence without required options
	var options findingsapiv1.GetOccurrenceOptions
	res, _, err = service.GetOccurrence(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("creating new note for delete")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")
	fmt.Println("creating new occurrence for delete")
	result, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)

	assert.NotNil(t, result)

	if result != nil {
		getOccurrenceOptions := service.NewGetOccurrenceOptions(accountID, *(createOccurrenceOptions.ProviderID), *(createOccurrenceOptions.ID))
		res, _, err := service.GetOccurrence(getOccurrenceOptions)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, *(res.ID), *(createOccurrenceOptions.ID))
		assert.Equal(t, *res.Finding.Severity, *createOccurrenceOptions.Finding.Severity)
		assert.Equal(t, *res.ResourceURL, *createOccurrenceOptions.ResourceURL)
		assert.Equal(t, *res.Context.ResourceName, *createOccurrenceOptions.Context.ResourceName)

		//Test get occ using setter fucntions
		var getOptions findingsapiv1.GetOccurrenceOptions
		getOptions.SetAccountID(accountID)
		getOptions.SetHeaders(headers)
		getOptions.SetProviderID(*createOccurrenceOptions.ProviderID)
		getOptions.SetOccurrenceID(*createOccurrenceOptions.ID)
		res, _, err = service.GetOccurrence(&getOptions)
		assert.Nil(t, err)
		assert.NotNil(t, res)

		fmt.Println("Cleaning up occurrence....")
		deleteOccurrenceHelper(t, createOccurrenceOptions)

		fmt.Println("Cleaning up note....")
		deleteNoteHelper(t, createNoteOptions)
	}
}

func TestCreateOccurrence(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test ListProviderOccurrence with nil options
	res, _, err := service.CreateOccurrence(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test CreateOccurrence without required options
	var options findingsapiv1.CreateOccurrenceOptions
	res, _, err = service.CreateOccurrence(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("creating new note")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("Creating occurrence")
	result, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)

	assert.NotNil(t, result)
	fmt.Println("Created occurrence")

	assert.Equal(t, *(result.ID), *(createOccurrenceOptions.ID))
	assert.Equal(t, *result.Finding.Severity, *createOccurrenceOptions.Finding.Severity)
	assert.Equal(t, *result.ResourceURL, *createOccurrenceOptions.ResourceURL)
	assert.Equal(t, *result.Context.ResourceName, *createOccurrenceOptions.Context.ResourceName)
	assert.Equal(t, *result.Finding.NetworkConnection.Direction, *createOccurrenceOptions.Finding.NetworkConnection.Direction)
	assert.Equal(t, *result.Finding.NetworkConnection.Protocol, *createOccurrenceOptions.Finding.NetworkConnection.Protocol)

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)

	//Test create occurrence using setters
	occID := "sec_advisor_202X_occ_id1"
	noteID := *createNoteOptions.ID
	providerID := "sec_advisor_202X_provider"
	noteName := accountID + "/providers/" + providerID + "/notes/" + noteID
	createOptions := service.NewCreateOccurrenceOptions(accountID, *createNoteOptions.ProviderID, noteName, "FINDING", occID)
	createTime := strfmt.DateTime(time.Now())
	region := "us-south"
	context := findingsapiv1.Context{Region: &region}
	remediationTitle := "title"
	remediationURL := "https://hello.world"
	nextStep := []findingsapiv1.RemediationStep{{Title: &remediationTitle, URL: &remediationURL}}
	severity := "MEDIUM"
	certainity := "LOW"
	finding := findingsapiv1.Finding{Severity: &severity, Certainty: &certainity, NextSteps: nextStep}
	createOptions.SetAccountID(accountID)
	createOptions.SetProviderID(*createNoteOptions.ProviderID)
	createOptions.SetNoteName(noteName)
	createOptions.SetKind("FINDING")
	createOptions.SetID(occID)
	createOptions.SetResourceURL("https://ss.ss")
	createOptions.SetRemediation("steps for remediation")
	createOptions.SetCreateTime(&createTime)
	createOptions.SetContext(&context)
	createOptions.SetFinding(&finding)
	createOptions.SetUpdateTime(&createTime)
	createOptions.SetReplaceIfExists(true)
	fmt.Println("Creating occurrence")
	result, resp, operationErr := service.CreateOccurrence(createOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		t.Fatal("Failed to create occurrence: ", operationErr)
	}
	assert.Equal(t, *(result.ID), *(createOptions.ID))
	assert.Equal(t, *result.Finding.Severity, *createOptions.Finding.Severity)
	assert.Equal(t, *result.ResourceURL, *createOptions.ResourceURL)

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOptions)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestDeleteOccurrence(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test DeleteOccurrence with nil options
	res, err := service.DeleteOccurrence(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test DeleteOccurrence without required options
	var options findingsapiv1.DeleteOccurrenceOptions
	res, err = service.DeleteOccurrence(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("creating new note for delete")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("Creating new occurrence for delete....")
	result, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)
	if result == nil {
		t.Fatal("Failed to create occurrence")
	}

	fmt.Println("Created new occurrence for delete....")
	assert.Equal(t, *(result.ID), *(createOccurrenceOptions.ID))
	fmt.Println("Deleting occurrence....")
	var deleteOptions = service.NewDeleteOccurrenceOptions(*(createOccurrenceOptions.AccountID), *(createOccurrenceOptions.ProviderID), *(createOccurrenceOptions.ID))
	response, err := service.DeleteOccurrence(deleteOptions)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)

	//Test Delete occ using setters
	fmt.Println("Creating new occurrence for delete....")
	result, createOccurrenceOptions = createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)
	assert.NotNil(t, result)
	if result == nil {
		t.Fatal("Failed to create occurrence")
	}

	delOptions := service.NewDeleteOccurrenceOptions("test", "test", "test")
	assert.Equal(t, *delOptions.AccountID, "test")
	delOptions.SetAccountID(accountID)
	delOptions.SetHeaders(headers)
	delOptions.SetOccurrenceID(*createOccurrenceOptions.ID)
	delOptions.SetProviderID(*createOccurrenceOptions.ProviderID)
	response, err = service.DeleteOccurrence(delOptions)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestUpdateOccurrence(t *testing.T) {
	//Test UpdateOccurrence with nil options
	res, _, err := service.UpdateOccurrence(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test UpdateOccurrence without required options
	var options findingsapiv1.UpdateOccurrenceOptions
	res, _, err = service.UpdateOccurrence(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	fmt.Println("Creating new note for edit....")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/noteWithKpi.json")
	fmt.Println("Creating new occurrence for edit....")
	_, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occKpi.json", *createNoteOptions.ID)

	query, err := ioutil.ReadFile(inputFilePath + "/editedKpiOcc.json")
	if err != nil {
		t.Fatal(err)
	}

	var updateOccurrenceOptions *findingsapiv1.UpdateOccurrenceOptions
	json.Unmarshal([]byte(query), &updateOccurrenceOptions)
	updateOccurrenceOptions.SetHeaders(headers)
	updateOccurrenceOptions.SetAccountID(accountID)
	updateOccurrenceOptions.SetID(*createOccurrenceOptions.ID)
	updateOccurrenceOptions.SetOccurrenceID(*createOccurrenceOptions.ID)

	//Update kpis to verify kpi setters
	kpi, _ := service.NewKpi(float64(5))
	total := float64(10)
	kpi.Total = &total

	updateOccurrenceOptions.SetKpi(kpi)
	noteName := accountID + "/providers/" + *updateOccurrenceOptions.ProviderID + "/notes/" + *createNoteOptions.ID
	updateOccurrenceOptions.NoteName = &noteName

	fmt.Println("Editing occurrence....")
	result, resp, operationErr := service.UpdateOccurrence(updateOccurrenceOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		fmt.Println("Failed to edit occurrence: ", operationErr)
	}

	assert.Nil(t, operationErr)
	assert.NotNil(t, result)
	fmt.Println("Edited occurrence....")

	assert.Equal(t, *(result.ID), *(updateOccurrenceOptions.ID))
	assert.Equal(t, *(result.Kpi.Value), *(updateOccurrenceOptions.Kpi.Value))
	assert.Equal(t, *(result.Kpi.Total), *(updateOccurrenceOptions.Kpi.Total))

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestUpdateOccurrenceUsingSetters(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	fmt.Println("Creating new note for edit....")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("Creating new occurrence for edit....")
	_, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)

	//Setup updated data
	noteID := *createNoteOptions.ID
	providerID := *createNoteOptions.ProviderID
	occID := *createOccurrenceOptions.ID
	noteName := accountID + "/providers/" + providerID + "/notes/" + noteID
	updateOccurrenceOptions := service.NewUpdateOccurrenceOptions(accountID, providerID, occID, noteName, "FINDING", occID)
	time := strfmt.DateTime(time.Now())
	remediationTitle := "title"
	remediationURL := "https://hello.world"
	nextStep := []findingsapiv1.RemediationStep{{Title: &remediationTitle, URL: &remediationURL}}
	severity := "MEDIUM"
	certainity := "LOW"
	finding := findingsapiv1.Finding{Severity: &severity, Certainty: &certainity, NextSteps: nextStep}
	resourceName := "us-south"
	context := findingsapiv1.Context{ResourceName: &resourceName}
	updateOccurrenceOptions.SetAccountID(accountID)
	updateOccurrenceOptions.SetProviderID(providerID)
	updateOccurrenceOptions.SetOccurrenceID(occID)
	updateOccurrenceOptions.SetNoteName(noteName)
	updateOccurrenceOptions.SetKind("FINDING")
	updateOccurrenceOptions.SetID(occID)
	updateOccurrenceOptions.SetResourceURL("https://ss.ss")
	updateOccurrenceOptions.SetRemediation("remediation")
	updateOccurrenceOptions.SetCreateTime(&time)
	updateOccurrenceOptions.SetUpdateTime(&time)
	updateOccurrenceOptions.SetContext(&context)
	updateOccurrenceOptions.SetFinding(&finding)

	fmt.Println("Editing occurrence....")
	result, resp, operationErr := service.UpdateOccurrence(updateOccurrenceOptions)
	if operationErr != nil {
		fmt.Println(resp.Result)
		fmt.Println("Failed to edit occurrence: ", operationErr)
	}

	assert.Nil(t, operationErr)
	assert.NotNil(t, result)
	fmt.Println("Edited occurrence....")
	assert.Equal(t, *(result.ID), *(updateOccurrenceOptions.ID))
	assert.Equal(t, *(result.ResourceURL), *(updateOccurrenceOptions.ResourceURL))

	fmt.Println("Cleaning up occurrence....")
	deleteOccurrenceHelper(t, createOccurrenceOptions)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestGetAllNote(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test GetAllNote with nil options
	res, _, err := service.ListNotes(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test GetAllNote without required options
	var options findingsapiv1.ListNotesOptions
	res, _, err = service.ListNotes(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("creating notes")
	_, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("created notes")
	//Test using required fileds
	getNotesOptions := service.NewListNotesOptions(accountID, *(createNoteOptions.ProviderID))
	res, _, err = service.ListNotes(getNotesOptions)

	found := 0
	for i := 0; i < len(res.Notes); i++ {
		if *res.Notes[i].ID == *createNoteOptions.ID {
			found++
		}
	}

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, found, 1)

	//Test using setters
	var listOptions findingsapiv1.ListNotesOptions
	listOptions.SetAccountID(accountID)
	listOptions.SetProviderID(*createNoteOptions.ProviderID)
	listOptions.SetHeaders(headers)
	listOptions.SetPageSize(10)
	listOptions.SetPageToken(*res.NextPageToken)
	res, _, err = service.ListNotes(&listOptions)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	fmt.Println("Cleaning up notes....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestGetNote(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test GetNote with nil options
	res, _, err := service.GetNote(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test GetNote without required options
	var options findingsapiv1.GetNoteOptions
	res, _, err = service.GetNote(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	result, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")
	assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
	assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))

	getNotesOptions := service.NewGetNoteOptions(accountID, *(createNoteOptions.ProviderID), *(createNoteOptions.ID))
	res, _, err = service.GetNote(getNotesOptions)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, *(res.ID), *(createNoteOptions.ID))
	assert.Equal(t, *(res.ShortDescription), *(createNoteOptions.ShortDescription))
	assert.Equal(t, *(res.Kind), *(createNoteOptions.Kind))

	//Test get with setters
	getOptions := service.NewGetNoteOptions("test", "test", "test")
	assert.Equal(t, *getOptions.AccountID, "test")
	getOptions.SetHeaders(headers)
	getOptions.SetAccountID(accountID)
	getOptions.SetNoteID(*createNoteOptions.ID)
	getOptions.SetProviderID(*createNoteOptions.ProviderID)
	res, _, err = service.GetNote(getOptions)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, *(res.ID), *(createNoteOptions.ID))
	assert.Equal(t, *(res.ShortDescription), *(createNoteOptions.ShortDescription))
	assert.Equal(t, *(res.Kind), *(createNoteOptions.Kind))

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestPostNote(t *testing.T) {
	result, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	assert.NotNil(t, result)
	if result != nil {
		assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))
		fmt.Println("Cleaning up note....")
		deleteNoteHelper(t, createNoteOptions)
	}
}

func TestDeleteNote(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test delete Note with nil options
	res, err := service.DeleteNote(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test delete Note without required options
	var options findingsapiv1.DeleteNoteOptions
	res, err = service.DeleteNote(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test delete noe using required
	fmt.Println("Creating new note for delete....")
	result, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")
	assert.NotNil(t, result)

	if result != nil {
		fmt.Println("Created new note for delete....")
		assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))
		fmt.Println("Deleting note....")
		var deleteOptions = service.NewDeleteNoteOptions(*(createNoteOptions.AccountID), *(createNoteOptions.ProviderID), *(createNoteOptions.ID))
		response, err := service.DeleteNote(deleteOptions)
		assert.Nil(t, err)
		assert.Equal(t, response.StatusCode, 200)
	}

	//Test delete note using setters
	fmt.Println("Creating new note for delete....")
	result, createNoteOptions = createNoteHelper(t, inputFilePath+"/note.json")
	assert.NotNil(t, result)

	if result != nil {
		fmt.Println("Deleting note....")
		var deleteOptions = service.NewDeleteNoteOptions("test", "test", "test")
		assert.Equal(t, *deleteOptions.AccountID, "test")
		deleteOptions.SetProviderID(*(createNoteOptions.ProviderID))
		deleteOptions.SetAccountID(accountID)
		deleteOptions.SetHeaders(headers)
		deleteOptions.SetNoteID(*(createNoteOptions.ID))
		response, err := service.DeleteNote(deleteOptions)
		assert.Nil(t, err)
		assert.Equal(t, response.StatusCode, 200)
	}
}

func TestPostCard(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	providerID := "exampleProvider"
	shortDescription := "exampleString"
	longDescription := "exampleString"
	kind := "CARD"
	id := "exampleCardNote"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	cardValueKind := "FINDING_COUNT"
	cardElementKind := "TIME_SERIES"
	cardElementText := "text TIME_SERIES"
	var cardValueTypes []findingsapiv1.ValueTypeIntf
	cardValueTypes = append(cardValueTypes, &findingsapiv1.ValueTypeFindingCountValueType{Kind: &cardValueKind, Text: &cardElementKind, FindingNoteNames: []string{"providers/sdktest/notes/sdk_note_id1"}})
	cardValueTypes = append(cardValueTypes, &findingsapiv1.ValueTypeFindingCountValueType{Kind: core.StringPtr("FINDING_COUNT"), Text: &cardElementText, FindingNoteNames: []string{"providers/sdktest/notes/sdk_note_id2"}})

	var cardElement []findingsapiv1.CardElementIntf
	cardElement = append(cardElement, &findingsapiv1.CardElementTimeSeriesCardElement{
		Kind:       &cardElementKind,
		Text:       &cardElementText,
		ValueTypes: cardValueTypes, //ValueType required for kind NUMERIC
	})
	card, _ := service.NewCard("My Security Tools", "Card from Go SDK time series", "subtitle", []string{"providers/sdktest/notes/sdk_note_id1"}, cardElement)
	var createNoteOptions = service.NewCreateNoteOptions("fd139bf514294d76914c05c11b6a175a", providerID, shortDescription, longDescription, kind, id, reportedBy)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetCard(card)
	result, _, operationErr := service.CreateNote(createNoteOptions)
	if operationErr != nil {
		fmt.Println("Failed to create Card: ", operationErr)
		panic(operationErr)
	}
	assert.NotNil(t, result)
	if result != nil {
		assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))
		assert.Equal(t, *result.Card.Title, *createNoteOptions.Card.Title)

		fmt.Println("Cleaning up note....")
		deleteNoteHelper(t, createNoteOptions)
	}
}

func TestPostCardUsingSetters(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	createNoteOptions := service.NewCreateNoteOptions("test", "test", "test", "test", "test", "test", new(findingsapiv1.Reporter))
	createTime := strfmt.DateTime(time.Now())
	kind := "CARD"
	label := "label"
	url := "https://ss.ss"
	reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
	cardValueKind := "FINDING_COUNT"
	cardElementKind := findingsapiv1.CardElement_Kind_Numeric
	cardElementText := "text"
	cardValueType := findingsapiv1.NumericCardElementValueType{Kind: &cardValueKind, FindingNoteNames: []string{"providers/sdktest/notes/sdk_note_id1"}}
	var cardElement []findingsapiv1.CardElementIntf
	cardElement = append(cardElement, &findingsapiv1.CardElementNumericCardElement{
		Kind:      &cardElementKind,
		Text:      &cardElementText,
		ValueType: &cardValueType, //ValueType required for kind NUMERIC
	})

	card, _ := service.NewCard("My Security Tools", "Card posted from Go SDK", "subtitle", []string{"providers/sdktest/notes/sdk_note_id1"}, cardElement)
	assert.Equal(t, *createNoteOptions.AccountID, "test")
	createNoteOptions.SetAccountID(accountID)
	createNoteOptions.SetCreateTime(&createTime)
	createNoteOptions.SetExpirationTime(&createTime)
	createNoteOptions.SetHeaders(headers)
	createNoteOptions.SetID("sdk_note_id1_card")
	createNoteOptions.SetProviderID("sdktest")
	createNoteOptions.SetShortDescription("short desc")
	createNoteOptions.SetLongDescription("long desc")
	createNoteOptions.SetKind(kind)
	createNoteOptions.SetReportedBy(reportedBy)
	createNoteOptions.SetRelatedURL([]findingsapiv1.ApiNoteRelatedURL{{Label: &label, URL: &url}})
	createNoteOptions.SetUpdateTime(&createTime)
	createNoteOptions.SetCard(card)
	createNoteOptions.SetShared(false)

	result, _, err := service.CreateNote(createNoteOptions)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	fmt.Println("Cleaning up note....")
	deleteNoteHelper(t, createNoteOptions)
}

func TestEditNote(t *testing.T) {

	//Test edit Note with nil options
	res, _, err := service.UpdateNote(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test edit Note without required options
	var options findingsapiv1.UpdateNoteOptions
	res, _, err = service.UpdateNote(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	fmt.Println("Creating new note for edit....")
	result, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	if result != nil {
		fmt.Println("Created new note for edit....")
		assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))

		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"
		time := strfmt.DateTime(time.Now())
		reportedBy, _ := service.NewReporter("hello", "https://ss.ss")
		remediationTitle := "title"
		remediationURL := "https://hello.world"
		nextStep := []findingsapiv1.RemediationStep{{Title: &remediationTitle, URL: &remediationURL}}
		severity := "MEDIUM"
		finding, _ := service.NewFindingType(severity)
		finding.NextSteps = nextStep
		// findingsapiv1.FindingType{Severity: &severity, NextSteps: nextStep}
		label := "label"
		relatedUrl := "https://ss.ss"
		related := findingsapiv1.ApiNoteRelatedURL{Label: &label, URL: &relatedUrl}

		var updateNoteOptions findingsapiv1.UpdateNoteOptions
		fmt.Println("Created new note for edit....")
		updateNoteOptions.SetHeaders(headers)
		updateNoteOptions.SetAccountID(accountID)
		updateNoteOptions.SetProviderID("sec_advisor_202X_provider")
		updateNoteOptions.SetNoteID("providers/sec_advisor_202X_provider/notes/" + *createNoteOptions.ID)
		updateNoteOptions.SetShortDescription("sdk test findings edited")
		updateNoteOptions.SetLongDescription("sdk test findings edited")
		updateNoteOptions.SetKind("FINDING")
		updateNoteOptions.SetID(*createNoteOptions.ID)
		updateNoteOptions.SetNoteID(*createNoteOptions.ID)
		updateNoteOptions.SetRelatedURL([]findingsapiv1.ApiNoteRelatedURL{related})
		updateNoteOptions.SetReportedBy(reportedBy)
		updateNoteOptions.SetFinding(finding)
		updateNoteOptions.SetExpirationTime(&time)
		updateNoteOptions.SetCreateTime(&time)
		updateNoteOptions.SetUpdateTime(&time)
		updateNoteOptions.SetShared(false)

		fmt.Println("Editing note....")
		result, _, operationErr := service.UpdateNote(&updateNoteOptions)
		if operationErr != nil {
			fmt.Println("Failed to edit note: ", operationErr)
		}

		assert.Nil(t, operationErr)
		assert.NotNil(t, result)
		fmt.Println("Edited note....")

		assert.Equal(t, *(result.ID), *(updateNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(updateNoteOptions.ShortDescription))
		assert.Equal(t, *(result.LongDescription), *(updateNoteOptions.LongDescription))
		assert.Equal(t, *(result.Kind), *(updateNoteOptions.Kind))

	}
	fmt.Println("Deleting note....")
	var deleteOptions = service.NewDeleteNoteOptions(*(createNoteOptions.AccountID), *(createNoteOptions.ProviderID), *(createNoteOptions.ID))
	service.DeleteNote(deleteOptions)
}

func TestEditKpiNote(t *testing.T) {
	fmt.Println("Creating new note for edit....")
	result, createNoteOptions := createNoteHelper(t, inputFilePath+"/noteWithKpi.json")

	if result != nil {
		fmt.Println("Created new note for edit....")
		assert.Equal(t, *(result.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(createNoteOptions.ShortDescription))

		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"

		providerId := "sec_advisor_202X_provider"
		noteId := *createNoteOptions.ID
		ID := *createNoteOptions.ID
		longDesc := "long desc"
		shortDesc := "short desc"
		kind := "KPI"
		reportedBy, _ := service.NewReporter("id", "title")
		aggType := findingsapiv1.KpiType_AggregationType_Sum
		kpi, _ := service.NewKpiType(aggType)
		var updateNoteOptions = service.NewUpdateNoteOptions(accountID, providerId, noteId, shortDesc, longDesc, kind, ID, reportedBy)
		updateNoteOptions.SetHeaders(headers)
		updateNoteOptions.SetKpi(kpi)

		fmt.Println("Editing note....")
		result, _, operationErr := service.UpdateNote(updateNoteOptions)
		if operationErr != nil {
			fmt.Println("Failed to edit note: ", operationErr)
		}

		assert.Nil(t, err)
		assert.NotNil(t, result)
		fmt.Println("Edited note....")

		assert.Equal(t, *(result.ID), *(updateNoteOptions.ID))
		assert.Equal(t, *(result.ShortDescription), *(updateNoteOptions.ShortDescription))
		assert.Equal(t, *(result.LongDescription), *(updateNoteOptions.LongDescription))
		assert.Equal(t, *(result.Kind), *(updateNoteOptions.Kind))

	}
	fmt.Println("Deleting note....")
	var deleteOptions = service.NewDeleteNoteOptions(*(createNoteOptions.AccountID), *(createNoteOptions.ProviderID), *(createNoteOptions.ID))
	service.DeleteNote(deleteOptions)
}

func TestGetOccurrenceNote(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	//Test GetOccurrenceNote with nil options
	res, _, err := service.GetOccurrenceNote(nil)
	assert.NotNil(t, err)
	assert.Nil(t, res)

	//Test GetOccurrenceNote without required options
	var options findingsapiv1.GetOccurrenceNoteOptions
	res, _, err = service.GetOccurrenceNote(&options)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	fmt.Println("creating new note")
	res, createNoteOptions := createNoteHelper(t, inputFilePath+"/note.json")

	fmt.Println("Creating occurrence")
	result, createOccurrenceOptions := createOccurrenceHelper(t, inputFilePath+"/occurrence.json", *createNoteOptions.ID)

	assert.NotNil(t, result)
	fmt.Println("Created occurrence")

	if res != nil || result != nil {

		var getOccurrenceNoteOptions = service.NewGetOccurrenceNoteOptions(accountID, *createNoteOptions.ProviderID, *createOccurrenceOptions.ID)

		res, _, err := service.GetOccurrenceNote(getOccurrenceNoteOptions)
		if err != nil {
			t.Fatal("Failed to GetOccurrenceNote, err: ", err)
		}

		assert.Equal(t, *(res.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(res.Kind), *(createNoteOptions.Kind))
		assert.Equal(t, *res.Finding.Severity, *createNoteOptions.Finding.Severity)
		assert.Equal(t, *res.ShortDescription, *createNoteOptions.ShortDescription)

		//Test get occ note using setters
		getOccurrenceNoteOptions = service.NewGetOccurrenceNoteOptions("test", "test", "test")
		assert.Equal(t, *(getOccurrenceNoteOptions.AccountID), "test")
		getOccurrenceNoteOptions.SetAccountID(accountID)
		getOccurrenceNoteOptions.SetProviderID(*createNoteOptions.ProviderID)
		getOccurrenceNoteOptions.SetHeaders(headers)
		getOccurrenceNoteOptions.SetOccurrenceID(*createOccurrenceOptions.ID)
		res, _, err = service.GetOccurrenceNote(getOccurrenceNoteOptions)
		assert.Equal(t, *(res.ID), *(createNoteOptions.ID))
		assert.Equal(t, *(res.Kind), *(createNoteOptions.Kind))
		assert.Equal(t, *res.Finding.Severity, *createNoteOptions.Finding.Severity)
		assert.Equal(t, *res.ShortDescription, *createNoteOptions.ShortDescription)

		fmt.Println("Cleaning up occurrence....")
		deleteOccurrenceHelper(t, createOccurrenceOptions)

		fmt.Println("Cleaning up note....")
		deleteNoteHelper(t, createNoteOptions)
	}
}

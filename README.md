[![Build Status](https://api.travis-ci.org/ibm-cloud-security/scc-go-sdk.svg?branch=master)](https://travis-ci.org/github/ibm-cloud-security/scc-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Security and Compliance Center GO SDK

This repository contains the released GO client SDK for IBM Cloud Security and Compliance Center services.  Check out our listed below for more details.  
  
* Findings API : https://cloud.ibm.com/apidocs/security-advisor/findings
* Notifications API : https://cloud.ibm.com/apidocs/security-advisor/notifications
* Configuration Governance API : (https://cloud.ibm.com/apidocs/security-compliance/config)

<details>
<summary>Table of Contents</summary>

- [IBM Cloud Security and Compliance Center GO SDK](#ibm-cloud-security-and-compliance-center-go-sdk)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
      - [`go get` command](#go-get-command)
      - [Go modules](#go-modules)
      - [`dep` dependency manager](#dep-dependency-manager)
  - [Authentication](#authentication)
    - [Supplying the IAM API key:](#supplying-the-iam-api-key)
  - [Sending request headers](#sending-request-headers)
  - [Error Handling](#error-handling)
      - [Findings](#findings)
      - [Notifications](#notifications)
  - [Sample Code](#sample-code)
      - [Findings](#findings-1)
      - [Notifications](#notifications-1)
      - [Configuration Governance](#configuration-governance)
  - [Tests](#tests)
    - [Run unit tests:](#run-unit-tests)
    - [Run integration tests:](#run-integration-tests)
  - [License](#license)
  - [Open Issues](#open-issues)
</details>

## Overview

The  IBM Cloud Security and Compliance Center GO SDK allows developers to programmatically interact with the ibm cloud Security and Compliance Center findings and notifications api.

Service Name | Package name 
--- | --- 
[Findings Service](https://cloud.ibm.com/apidocs/security-advisor/findings) | findingsv1
[Notifications Service](https://cloud.ibm.com/apidocs/security-advisor/notifications) | notificationsv1
[Configuration Governance Service](https://cloud.ibm.com/apidocs/security-compliance/config) | configurationgovernancev1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.1

There are a few different ways to download and install the Findings API Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.com/ibm-cloud-security/scc-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/ibm-cloud-security/scc-go-sdk/findingsv1"
  "github.com/ibm-cloud-security/scc-go-sdk/notificationsv1"
  "github.com/ibm-cloud-security/scc-go-sdk/configurationgovernancev1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/ibm-cloud-security/scc-go-sdk"
  version = "v0.0.1"

```

then run `dep ensure`.

## Authentication

 Security and Compliance Center Findings API GO SDK uses token-based [Identity and Access Management (IAM) authentication](https://cloud.ibm.com/docs/iam?topic=iam-getstarted).

IAM authentication uses a service API key to get an access token that is passed with the call.
Access tokens are valid for a limited amount of time and must be regenerated.

To provide credentials to the SDK, you supply either an IAM service **API key** or an **access token**:

- Use the API key to have the SDK manage the lifecycle of the access token. The SDK requests an access token, ensures that the access token is valid, and refreshes it if necessary.
- Use the access token if you want to manage the lifecycle yourself. For details, see [Authenticating with IAM tokens](https://cloud.ibm.com/docs/services/watson/getting-started-iam.html).


### Supplying the IAM API key:


```go
import (
  "github.com/IBM/go-sdk-core/v5/core"
  "github.com/ibm-cloud-security/scc-go-sdk/findingsapiv1"
  "github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

authenticator := &core.IamAuthenticator{
	ApiKey: apiKey,
	URL:    URL, //Required only if you are targetting Dev/Preprod environment
}

findingsService, err := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
  Authenticator: authenticator, 
})
```



## Sending request headers

Custom headers can be passed in any request in the form of a `map[string]string` as:

```go
headers := make(map[string]string)
	headers["Custom-Header"] = "custom_value"
```

For example, to send a header called `Custom-Header` to a call in  `notificationsapiv1`, pass the headers parameter as:

```go
import (
  "github.com/IBM/go-sdk-core/v5/core"
  "github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

authenticator := &core.IamAuthenticator{
	ApiKey: apiKey,
	URL:    URL, //Required only if you are targetting Dev/Preprod environment
}
service, err := notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
  Authenticator: authenticator,
})

var deleteOptions = service.NewDeleteNotificationChannelOptions(accountID, channelID)
deleteOptions.SetHeaders(headers)
```

## Error Handling

The  scc-go-sdk generates an **error** for any unsuccessful method invocation.
If the method receives an error response from an API call to the service, it will generate an **error** which is sent has the last return value of the function. It also returns a **DetailedResponse** structure which consists further details about the response.

`Error` can be handled in the following way.  

#### Findings
```go
import (
 "fmt"

 "github.com/IBM/go-sdk-core/v5/core"
 "github.com/ibm-cloud-security/scc-go-sdk/findingsapiv1"
)

providerID := "providerID" //Invalid provider id
noteID := "custom-note" //invalid note id

authenticator := &core.IamAuthenticator{
  ApiKey: apiKey,
  URL:    url, //use for dev/preprod env
}
service, _ := findingsapiv1.NewFindingsApiV1(&findingsapiv1.FindingsApiV1Options{
  Authenticator: authenticator,
  URL:           "https://us-south.secadvisor.cloud.ibm.com/findings", //Specify url or use default
})

getNotesOptions := service.NewGetNoteOptions(accountID, "providerID", noteID)

result, response, err := service.GetNote(getNotesOptions)

if err != nil {
  fmt.Println(err) //Prints: "Not Found"
  fmt.Println(response.StatusCode) //Prints: 404
  fmt.Println(response.Result) //See Expected Response section below for details
}
```
***Expected Response*** for the above case case would be. This is of type map[string]interface {}-
```
map[
  detail: Document not found: <AccountID>/providers/providerID/notes/custom-note 
  instance: <AccountID>/providers/providerID/notes/custom-note 
  status: 404 
  title: Not Found 
  type: about:blank
]

```

#### Notifications

```go
import (
 "fmt"

 "github.com/IBM/go-sdk-core/v5/core"
 "github.com/ibm-cloud-security/scc-go-sdk/notificationsapiv1"
)

channelID := "channel" //invalid channel id

authenticator := &core.IamAuthenticator{
  ApiKey: apiKey,
  URL:    url, //use for dev/preprod env
}
service, _ := notificationsapiv1.NewNotificationsApiV1(&notificationsapiv1.NotificationsApiV1Options{
  Authenticator: authenticator,
  URL:           "https://us-south.secadvisor.cloud.ibm.com/notifications", //Specify url or use default
})

getChannelOptions := service.NewGetNotificationChannelOptions(accountID, channelID)

result, response, err := service.GetNotificationChannel(getChannelOptions)

if err != nil {
  fmt.Println(err) //Prints: "Internal Server Error"
  fmt.Println(response.StatusCode) //Prints: 500
  fmt.Println(response.Result) //See Expected Response section below for details
}
```
***Expected Response*** for the above case case would be. This is of type map[string]interface {}-
```
map[
  code:NOTIFICATIONS-CHANNELS-API-ERR500-01 
  message:Internal Server Error
]

```

## Sample Code

#### Findings
Example | http method  
------------ | ------------- 
[***post_graph***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/postGraph.go) | POST /v1/{account_id}/graph 
[***list_providers***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/providers.go) | GET /v1/{account_id}/providers
[***create_finding_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L11) | POST /v1/{account_id}/providers/{provider_id}/notes
[***create_card_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L158) | POST /v1/{account_id}/providers/{provider_id}/notes
[***create_note_with_kpi***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L53) | POST /v1/{account_id}/providers/{provider_id}/notes
[***create_note_with_section***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L91) | POST /v1/{account_id}/providers/{provider_id}/notes
[***list_provider_notes***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L205) | GET /v1/{account_id}/providers/{provider_id}/notes
[***get_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L129) | GET /v1/{account_id}/providers/{provider_id}/notes
[***get_occurrence_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L300) | GET /v1/{account_id}/providers/{provider_id}/occurrences/{occurrence_id}/note
[***update_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L231) | PUT /v1/{account_id}/providers/{provider_id}/notes/{note_id}
[***delete_note***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/notes.go#L274) | DELETE /v1/{account_id}/providers/{provider_id}/notes/{note_id}
[***create_finding_occurrence***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L11) | POST /v1/{account_id}/providers/{provider_id}/occurrences
[***create_kpi_occurrence***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L57) | POST /v1/{account_id}/providers/{provider_id}/occurrences
[***update_occurrence***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L92) | PUT /v1/{account_id}/providers/{provider_id}/occurrences/{occurrence_id}
[***delete_occurrence***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L156) | DELETE /v1/{account_id}/providers/{provider_id}/occurrences/{occurrence_id}
[***get_occurrence***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L129) | GET /v1/{account_id}/providers/{provider_id}/occurrences/{occurrence_id}
[***list_note_occurrences***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L180) | GET /v1/{account_id}/providers/{provider_id}/notes/{note_id}/occurrences
[***list_provider_occurrences***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/findingsapiv1/occurrences.go#L209) | GET /v1/{account_id}/providers/{provider_id}/occurrences

#### Notifications
Example | http method  
------------ | ------------- 
[***list_channels***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/listChannels.go) | GET /v1/{account_id}/notifications/channels
[***create_channel***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/createChannel.go) | POST /v1/{account_id}/notifications/channels
[***bulk_delete_channel***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/bulkDeleteChannel.go) | DELETE /v1/{account_id}/notifications/channels
[***delete_channel***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/deleteChannel.go) | DELETE /v1/{account_id}/notifications/channels/{channel_id}
[***get_channel***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/getChannel.go) | GET /v1/{account_id}/notifications/channels/{channel_id}
[***update_channel***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/updateChannel.go) | PUT /v1/{account_id}/notifications/channels/{channel_id}
[***test_connection***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/testConnection.go) | GET /v1/{account_id}/notifications/channels/{channel_id}/test
[***get_public_key***](https://github.com/ibm-cloud-security/scc-go-sdk/blob/master/examples/notificationsapiv1/getPublicKey.go) | GET v1/{account_id}/notifications/public_key

#### Configuration Governance
Example | http method  
------------ | ------------- 


## Tests
### Run unit tests:
```shell
go test ./...
```  

Get code coverage for each test suite:
```shell
go test -coverprofile=unit.out ./...
go tool cover -html=unit.out
```  

### Run integration tests:
```shell
go test ./... -tags=integration
```  

Get code coverage for each test suite:
```shell
go test -coverprofile=integration.out ./... -tags=integration
go tool cover -html=integration.out
```  


## License

The ibm-cloud-scc-go-sdk is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).


## Open Issues

Currently  if `go get` is used as mode to download the module - GOPATH might face issues related to IBM SDK core module.
Advised method is to use  go mdoules.

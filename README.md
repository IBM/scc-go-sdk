[![Build Status](https://travis-ci.com/IBM/scc-go-sdk.svg?branch=main)](https://travis-ci.com/github/IBM/scc-go-sdk)
[![Release](https://img.shields.io/github/v/release/IBM/scc-go-sdk)](https://img.shields.io/github/v/release/IBM/scc-go-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/IBM/scc-go-sdk.svg)](https://pkg.go.dev/github.com/IBM/scc-go-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/scc-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codecov](https://codecov.io/gh/IBM/scc-go-sdk/branch/main/graph/badge.svg?token=59EXPRL5V5)](https://codecov.io/gh/IBM/scc-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)


# IBM Cloud Security & Compliance Center Go SDK Version v0.0.6

Go client library to interact with various
[IBM Cloud Security & Compliance Center APIs](https://cloud.ibm.com/docs?tab=api-docs&category=platform_services%2Csecurity).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Security & Compliance Center Go SDK allows developers to programmatically interact with the following IBM Cloud services:

| Service Name                                                                         | Package name              |
| ------------------------------------------------------------------------------------ | ------------------------- |
| [Findings](https://cloud.ibm.com/apidocs/security-advisor/findings)                  | findingsv1                |
| [Notifications](https://cloud.ibm.com/apidocs/security-advisor/notifications)        | notificationsv1           |
| [Configuration Governance](https://cloud.ibm.com/apidocs/security-compliance/config) | configurationgovernancev1 |

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one
[here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: v0.0.6

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/ibm/scc-go-sdk/findingsv1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `findingsv1` part of the import path is the package name
associated with the Findings service.
See the service table above to find the appropriate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/ibm/scc-go-sdk/findingsv1
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see
[this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/ibm/scc-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

The IBM Cloud Security & Compliance Center Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: test lint tidy
travis-ci: test-cov lint tidy

test:
	go test `go list ./...`
	find . -name "*.env" -type f -delete

test-cov:
	go test `go list ./...` ${COVERAGE}

test-int: test-dep
	go test `go list ./...` -tags=integration

test-int-cov: test-dep
	go test `go list ./...` -tags=integration ${COVERAGE}

test-dep:
	echo "${FINDINGS_ENV}" | base64 -d >> findings_v1.env
	echo "${NOTIFICATIONS_ENV}" | base64 -d >> notifications_v1.env
	# echo "${CONFIGURATION_GOVERNANCE_ENV}" | base64 -d >> configuration_governance_v1.env

lint:
	golangci-lint run

tidy:
	go mod tidy

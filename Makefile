# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic
SHELL := /usr/bin/env bash

all: test lint fmtcheck tidy
travis-ci: test-cov lint tidy

test:
	go test `go list ./...`

test-cov:
	go test `go list ./...` ${COVERAGE}

test-int:
	go test `go list ./...` -tags=integration

test-int-cov:
	go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	golangci-lint run

fmtcheck:
	diff -u <(echo -n) <(gofmt -d -s .)

tidy:
	go mod tidy

# Makefile to build the project

VDIR=v5
COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test lint tidy
travis-ci: test-cov lint tidy

.PHONY: test
test:
	cd ${VDIR} && go test `go list ./...`

.PHONY: test-cov
test-cov:
	cd ${VDIR} && go test `go list ./...` ${COVERAGE}

.PHONY: test-int
test-int:
	cd ${VDIR} && go test `go list ./...` -tags=integration

.PHONY: test-int-verbose
test-int-verbose:
	cd ${VDIR} && go test -v `go list ./...` -tags=integration

.PHONY: test-int-cov
test-int-cov:
	cd ${VDIR} && go test `go list ./...` -tags=integration ${COVERAGE}

.PHONY: test-examples
test-examples:
	cd ${VDIR} && go test -failfast `go list ./...` -tags=examples

.PHONY: test-examples-verbose
test-examples-verbose:
	cd ${VDIR} && go test -v -failfast `go list ./...` -tags=examples

.PHONY: lint
lint:
	diff -u <(echo -n) <(gofmt -d -s .)

.PHONY: tidy
tidy:
	cd ${VDIR} && go mod tidy

.PHONY: install
install:
	cd ${VDIR} && go mod download

.PHONY: docs
docs:
	cd ${VDIR} && go get github.com/johnstarich/go/gopages@v0.1.8 && go run github.com/johnstarich/go/gopages -base https://ibm.github.io/scc-go-sdk

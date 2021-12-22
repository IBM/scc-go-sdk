# Makefile to build the project

VDIR=v3
COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test lint fmtcheck tidy
travis-ci: test-cov lint tidy

test:
	cd ${VDIR} && go test `go list ./...`

test-cov:
	cd ${VDIR} && go test `go list ./...` ${COVERAGE}

test-int:
	cd ${VDIR} && go test `go list ./...` -tags=integration

test-int-cov:
	cd ${VDIR} && go test `go list ./...` -tags=integration ${COVERAGE}

fmtcheck:
	diff -u <(echo -n) <(gofmt -d -s .)

tidy:
	cd ${VDIR} && go mod tidy

install:
	cd ${VDIR} && go mod download

docs:
	cd ${VDIR} && go run github.com/johnstarich/go/gopages -base https://ibm.github.io/scc-go-sdk
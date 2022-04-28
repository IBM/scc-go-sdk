# Makefile to build the project

VDIR=v4
COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test lint lint tidy
travis-ci: test-cov lint tidy

test:
	cd ${VDIR} && go test `go list ./...`

test-cov:
	cd ${VDIR} && go test `go list ./...` ${COVERAGE}

test-int:
	cd ${VDIR} && go test `go list ./...` -tags=integration

test-int-cov:
	cd ${VDIR} && go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	diff -u <(echo -n) <(gofmt -d -s .)

tidy:
	cd ${VDIR} && go mod tidy

install:
	cd ${VDIR} && go mod download

docs:
	cd ${VDIR} && go get github.com/johnstarich/go/gopages@v0.1.8 && go run github.com/johnstarich/go/gopages -base https://ibm.github.io/scc-go-sdk
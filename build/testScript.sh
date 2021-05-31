#!/bin/bash

if [[ $TRAVIS_BRANCH == "main" && $TRAVIS_PULL_REQUEST == "false" ]]; then
  curl https://us-south.functions.appdomain.cloud/api/v1/web/e6b54af6-ab44-4149-a8e4-e906dcc58136/default/secadvstg-location-shift.json
  make test-int-cov
  find . -name "*.env" -type f -delete
fi

#!/bin/bash

set -euo pipefail
echo "${SCC_ENV}" | base64 -d >> v5/security_and_compliance_center_api_v3.env
make test-int

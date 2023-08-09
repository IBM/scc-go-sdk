#!/bin/bash

set -euo pipefail

echo "${POSTURE_MANAGEMENT_ENV}" | base64 -d >> v4/posture_management_v1.env
echo "${POSTURE_MANAGEMENT_V2_ENV}" | base64 -d >> v4/posture_management_v2.env
echo "${ADMIN_SERVICE_API_ENV}" | base64 -d >> v4/admin_service_api_v1.env
echo "${SCC_ENV}" | base64 -d >> v5/security_and_compliance_center_api_v3.env
make test-int

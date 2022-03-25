#!/bin/bash

set -euo pipefail

curl https://us-south.functions.appdomain.cloud/api/v1/web/e6b54af6-ab44-4149-a8e4-e906dcc58136/default/secadvstg-location-shift.json
echo "${FINDINGS_ENV}" | base64 -d >> v4/findings_v1.env
echo "${CONFIGURATION_GOVERNANCE_ENV}" | base64 -d >> v4/configuration_governance_v1.env
echo "${POSTURE_MANAGEMENT_ENV}" | base64 -d >> v4/posture_management_v1.env
echo "${POSTURE_MANAGEMENT_V2_ENV}" | base64 -d >> v4/posture_management_v2.env
echo "${ADMIN_SERVICE_API_ENV}" | base64 -d >> v4/admin_service_api_v1.env
echo "${ADDON_MANAGER_API_ENV}" | base64 -d >> v4/addon_manager_v1.env
make test-int

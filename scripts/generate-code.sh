#!/usr/bin/env bash

set -euo pipefail

# Remove generated files
if [ -f .openapi-generator/files ]; then
  while read -r file; do
    rm -f "$file"
  done < .openapi-generator/files
fi

# Generate library
docker run --rm --env JAVA_OPTS=-DmaxYamlCodePoints=9999999 -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.9.0 \
    generate \
    --config /local/.openapi-generator/config.yaml \
    --input-spec /local/api/openapi.yaml \
    --output /local \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --http-user-agent go-netbox/"$(cat api/netbox_version)"

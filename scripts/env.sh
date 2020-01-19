#!/bin/bash

# code path
export ROOT_PACKAGE="github.com/yaoice/autotz"

# API Group
export CUSTOM_RESOURCE_NAME="autotz"

# API Version
export CUSTOM_RESOURCE_VERSION="v1alpha1"

/Users/iceyao/Documents/src/k8s.io/code-generator/generate-groups.sh all "$ROOT_PACKAGE/pkg/generated" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

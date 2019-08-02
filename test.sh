#!/bin/bash

set -ex

hack/dind hack/test/unit

hack/dind hack/make.sh dynbinary test-integration

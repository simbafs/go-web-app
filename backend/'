#!/usr/bin/env bash

# https://belief-driven-design.com/build-time-variables-in-go-51439b26ef9/

# STEP 1: Determinate the required values

# PACKAGE="backend"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2>/dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

# STEP 2: Build the ldflags

LDFLAGS=(
	"-X main.Mode=release"
	"-X main.Version=${VERSION}"
	"-X main.CommitHash=${COMMIT_HASH}"
	"-X main.BuildTime=${BUILD_TIMESTAMP}"
)

# STEP 3: Actual Go build process

go build -ldflags="${LDFLAGS[*]}" -o ../main

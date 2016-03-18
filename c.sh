#!/bin/sh

VERSION="$(git describe --tags $1)"
GOOS=$OS CGO_ENABLED=0 GOARCH=$ARCH go build -ldflags "-X main.Version=$VERSION" -o vegeta

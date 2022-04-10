#!/bin/sh

CGO_ENABLED=0 go build -o build/server -tags gingonic,release ./cmd/server

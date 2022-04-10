#!/bin/sh

CGO_ENABLED=0 go build -o build/importer ./cmd/importer

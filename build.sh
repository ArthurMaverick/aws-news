#!/bin/bash

set -e

cd ./
CGO_ENABLED=0
go build -tags lambda.norpc -o main ./main.go | jq -Rs '{"status": "success", "output": .}'

#!/bin/bash
set -euo pipefail
cd $(dirname $0)/../

for package in $(go list ./...); do
    echo ======== $package ========
    echo + go test -failfast $package
           go test -failfast $package
done
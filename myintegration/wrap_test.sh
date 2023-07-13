#!/bin/bash
PKGARGS="$*"
rm -rf covdatafiles; mkdir covdatafiles
set -x

GOCOVERDIR=covdatafiles ./test.sh -cover $PKGARGS
go tool covdata percent -i=covdatafiles -o=cov.txt
go tool cover -func=cov.txt
#!/bin/bash
BUILDARGS="$*"
go build $BUILDARGS -o myintegration .

NUMS=(-1 2 hello)
for NUM in "${NUMS[@]}"; do
    ./myintegration -num=$NUM > /dev/null
done
echo "finished processing ${#NUMS[@]} cases, no crashes"
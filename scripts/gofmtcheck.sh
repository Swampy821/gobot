#!/usr/bin/env bash
# Taken from https://raw.githubusercontent.com/hashicorp/terraform/master/scripts/gofmtcheck.sh

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
GOFMT_FILES=$(gofmt -l $(find . -name '*.go' | grep -v vendor))
if [[ -n ${GOFMT_FILES} ]]; then
    printf "gofmt needs running on the following files:\n"
    printf "${GOFMT_FILES}\n"
    printf "You can use the command: \`make fmt\` to reformat code.\n"
    exit 1
fi

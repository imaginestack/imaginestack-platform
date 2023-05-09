#!/bin/bash

set -e

GV="$1"

rm -rf ./pkg/client
./hack/generate_group.sh "client,lister,informer" imaginekube.com/imaginekube/pkg/client imaginekube.com/api "${GV}" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv imaginekube.com/imaginekube/pkg/client ./pkg/
rm -rf ./imaginekube.com

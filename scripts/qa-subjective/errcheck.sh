#!/bin/bash

# finds unchecked errors

# find the dir we exist within...
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
# and cd into root project dir
echo $DIR
cd ${DIR}/../..
go get -u  github.com/kisielk/errcheck
errcheck -ignoregenerated ./...

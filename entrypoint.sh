#!/bin/bash

# pushd `dirname $0` > /dev/null

export GOPATH=/go:$GOPATH
# export GIN_MODE=debug
# /go/src/vc.cli/vc-cli --registry=kubernetes
/go/src/vc.cli/vc-cli --selector=static --server_address=0.0.0.0:8080 --broker_address=0.0.0.0:10001 --registry=kubernetes
# /go/src/vc.cli/vc-cli --server_name=vincross --server_address=0.0.0.0:8080 --broker_address=0.0.0.0:10001 --registry=kubernetes
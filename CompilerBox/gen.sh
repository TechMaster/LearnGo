#!/usr/bin/env bash
protoc --proto_path=proto/ -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogoslick_out=proto/. \
--micro_out=proto/.  \
compiler.proto
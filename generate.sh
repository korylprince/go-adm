#!/bin/bash

go install ./cmd/yamlschemagen

pushd ./schema/
go generate
popd

go install ./cmd/cmdgen
go install ./cmd/declgen
go install ./cmd/profilegen

pushd ./commands/
go generate
popd

pushd ./declarative/
go generate
popd

pushd ./profiles/
go generate
popd

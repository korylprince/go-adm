#!/bin/bash

go install ./cmd/yamlschemagen

pushd ./schema/
go generate
popd

go install ./cmd/cmdgen
go install ./cmd/declgen
go install ./cmd/profilegen
go install ./cmd/structgen

pushd ./generated/mdm/
go generate
popd

pushd ./generated/declarative/
go generate
popd

pushd ./generated/mdm/checkin/
go generate
popd

pushd ./generated/mdm/errors/
go generate
popd

pushd ./generated/other/
go generate
popd

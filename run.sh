#!/bin/bash
cd lang/token
go generate
cd ../..
go build -o bin/neonc
# go test ./tests
bin/neonc $1
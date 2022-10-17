#!/bin/bash
cd lib
go generate
cd ..
go build -o bin/neonc
go test ./tests
bin/neonc $1
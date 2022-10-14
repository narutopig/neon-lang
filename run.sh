#!/bin/bash
cd lib
go generate
cd ..
go build -o bin/neonc
bin/neonc $1
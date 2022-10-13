#!/bin/bash
cd lib
go generate
cd ..
go build -o bin
bin/neon-lang $1
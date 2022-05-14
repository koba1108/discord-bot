#!/bin/sh

echo Running static check for golang
go install honnef.co/go/tools/cmd/staticcheck@2020.2.1
staticcheck ./internals/...

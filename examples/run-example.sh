#!/bin/bash

# ./run-result.sh <example file>

file=$1
code_file=$(mktemp /tmp/go-genned-code.XXX.go)

echo "CODE:"
go run $file | tee $code_file
echo "RUN RESULT:"
go run $code_file

rm $code_file

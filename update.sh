#!/bin/bash
echo "read only update contracts"
cd app
go get -u github.com/i-b8o/regulations_contracts@$1

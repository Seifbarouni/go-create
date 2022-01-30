#!/bin/bash
cd cmd/go-create
go build -o ../../bin/
sudo cp ../../bin/go-create /usr/local/bin/
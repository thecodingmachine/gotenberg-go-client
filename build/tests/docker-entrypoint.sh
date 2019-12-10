#!/bin/bash

set -xe

# Testing Go client.
gotenberg &
sleep 10
go test -race -cover -covermode=atomic github.com/thecodingmachine/gotenberg-go-client/v7
sleep 10 # allows Gotenberg to remove generated files.
#!/usr/bin/env bash
make init                      # Initializes 
make                           # Builds the binary
make test                      # Runs unit tests with coverage 
./test.sh                      # Runs the test cases

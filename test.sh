#!/usr/bin/env bash

COUNT=1
for i in $(seq 1 $COUNT); do
    # Run the test cases
    INPUT_FILE="./testcases/input$i.txt"
    OUTPUT_FILE="./testcases/output$i.txt"
    TEMP_FILE="/tmp/output($i).txt"

    cat $INPUT_FILE | ./build/main > $TEMP_FILE
    if diff -w $TEMP_FILE $OUTPUT_FILE; then
        echo "Test case $i passed"
    else
        echo "Test case $i failed"
    fi
done

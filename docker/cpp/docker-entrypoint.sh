#!/bin/bash

# set -e

submission_directory="/sandbox/submission"
testcases_directory="/sandbox/testcases"

mkdir -p "$submission_directory/output" "$submission_directory/error"

cd $submission_directory
g++ code.cpp -o executable 2>"$submission_directory/compilation_error.err"

if [[ $? -ne 0 ]]; then
    exit 145
fi

# TODO: Restrict container network access

timelimit=${timelimit:-1}

for testcase in $(ls $testcases_directory); do
    timeout -s 9 -t $timelimit ./executable 0<"$testcases_directory/$testcase" \
        1>"$submission_directory/output/$testcase" \
        2>"$submission_directory/error/$testcase"
    if [[ $? -eq 137 ]]; then
        exit 146
    fi
done

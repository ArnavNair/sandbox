#!/bin/bash

submission_directory="/sandbox/submission"
testcases_directory="/sandbox/testcases"

mkdir -p "$submission_directory/output" "$submission_directory/error"

cd $submission_directory
g++ code.cpp -o executable

for testcase in $(ls $testcases_directory); do
    filename="${testcase%.*}"
    ./executable 0<"$testcases_directory/$testcase" 1>"$submission_directory/output/$filename.out" 2>"$submission_directory/error/$filename.err"
done

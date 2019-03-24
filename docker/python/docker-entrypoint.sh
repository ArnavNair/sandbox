#!/bin/bash

submission_directory="/sandbox/submission"
testcases_directory="/sandbox/testcases"
mkdir -p $submission_directory/output $submission_directory/error

cd $submission_directory
mv code code.py

# TODO: Restrict container network access

for testcase in $(ls $testcases_directory); do
    python3 code.py 0<"$testcases_directory/$testcase" \
        1>"$submission_directory/output/$testcase" \
        2>"$submission_directory/error/$testcase"
done

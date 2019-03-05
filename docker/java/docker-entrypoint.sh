#!/bin/bash

submission_directory="/sandbox/submission"
testcases_directory="/sandbox/testcases"
mkdir -p $submission_directory/output $submission_directory/error

cd $submission_directory
mv code Code.java

# Compile code. Redirect errors to compilation_error.err file.
javac Code.java 2> error/compilation_error.err
if [[ $? -ne 0 ]]; then
    exit 145 # Compilation error. Terminate.
fi

# TODO: Restrict container network access

# Run code for each test case.
for testcase in $(ls $testcases_directory); do
    java Code 0<"$testcases_directory/$testcase" \
        1>"$submission_directory/output/$testcase" \
        2>"$submission_directory/error/$testcase"
done

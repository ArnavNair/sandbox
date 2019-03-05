#!/bin/bash

submission_directory="/sandbox/submission"
testcases_directory="/sandbox/testcases"
mkdir -p $submission_directory/output $submission_directory/error

cd $submission_directory
mv code code.py

# TODO: Restrict container network access

for testcase in $(ls testcases); do
    python3 ./submission/code.py < testcases/$testcase \
        1>./submission/output/$testcase \
        2>./submission/error/$testcase
done

for testcase in $(ls testcases); do
    python3 code.py < testcases/$testcase 1>output/$testcase.out 2>output/$testcase.err
done

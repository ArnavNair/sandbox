for testcase in $(ls testcases); do
    ./executable < testcases/$testcase 1>output/$testcase.out 2>output/$testcase.err
done
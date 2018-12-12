for testcase in $(ls testcases); do
    ./executable < testcases/$testcase
done

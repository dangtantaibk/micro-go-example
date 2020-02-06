#!/bin/bash
GOPATH=${PWD}
RED=`tput setaf 1`
GRN=`tput setaf 2`
PUR=`tput setaf 13`
RESET=`tput sgr0`
TEST_RESULT_DIR=./test-results
mkdir -p ${TEST_RESULT_DIR}
for var in "$@"
do
    echo "${GRN}Checking src/$var${RESET}"
    PKG=$var
    PKG_LIST+="$(go list ${PKG}/... | grep -v /vendor/ | grep -v migrations) "
done
echo "----------------"
echo "${GRN}Test:${RESET}"

go test -v -covermode=count ${PKG_LIST} -coverprofile ${TEST_RESULT_DIR}/.testCoverage.txt | tee ${TEST_RESULT_DIR}/test.log; echo ${PIPESTATUS[0]} > ${TEST_RESULT_DIR}/test.out

cat ${TEST_RESULT_DIR}/test.log | go-junit-report > ${TEST_RESULT_DIR}/report.xml

echo "----------------"
echo "${GRN}Result:${RESET}"
go tool cover -func ${TEST_RESULT_DIR}/.testCoverage.txt
exit $(cat ${TEST_RESULT_DIR}/test.out)

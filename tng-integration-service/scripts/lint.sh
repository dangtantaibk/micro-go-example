#!/bin/bash
GOPATH=${PWD}
# PKG_LIST=$(go list $1/... | grep -v /vendor/) ; \
# golint -set_exit_status ${PKG_LIST} ; \
# PKG_LIST=$(go list $2/... | grep -v /vendor/) ; \
# golint -set_exit_status ${PKG_LIST} ; \
EXIT_CODE=0
cd src
for var in "$@"
do
    echo "Linting src/$var"
    PKG_LIST=$(go list $var/... | grep -v /vendor/ | grep -v migrations) ; \
    golint -set_exit_status ${PKG_LIST} ; \
    EXIT_CODE=$(( $EXIT_CODE + $? ))
done
cd ..
exit $EXIT_CODE
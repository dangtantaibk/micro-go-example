#!/bin/bash
export PATH=$PATH:$GOPATH/bin
cd src/tng/common
make migrate
# If you want to add more command here, please save the exit code of `make migrate` and return it
#     Example:
#           tmp=$?
#           echo "Do you code here"
#           exit $tmp

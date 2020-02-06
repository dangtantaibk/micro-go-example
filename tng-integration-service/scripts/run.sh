#!/bin/bash
make migrate
[ $? -eq 0 ] || exit $?;
make run
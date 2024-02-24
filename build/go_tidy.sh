#!/bin/bash
#
# Copyright © 2024 honeysense.com All rights reserved.
# Author: sunrui
# Date: 2024-02-25 01:22:45
#

cd ../internal || exit
pwd
go mod tidy -v
go get -u all
cd ../pkg || exit
pwd
go mod tidy -v
go get -u all
cd ../cmd/client/ || exit
pwd
go mod tidy -v
go get -u all
cd ../server/ || exit
pwd
go mod tidy -v
go get -u all

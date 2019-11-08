#!/bin/bash

rm -rf testdata/commits-on-master
rm -rf testdata/commits-on-different-branches

cd testdata

git clone commits-on-master.bundle

git clone commits-on-different-branches.bundle

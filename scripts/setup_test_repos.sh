#!/bin/bash

rm -rf testdata/commits_on_branch

if [ ! -d './testdata/commits_on_branch' ]; then
  echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
  cd testdata
  git clone commits_on_branch.bundle
  echo 'done'
  
else
  echo 'testdata/commits_on_branch_test directory already exists; aborting.'
  exit -1
fi

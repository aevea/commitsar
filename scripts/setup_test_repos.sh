#!/bin/bash

rm -rf testdata/commits_on_branch
rm -rf testdata/git_tags

if [ ! -d './testdata/commits_on_branch' ]; then
  cd testdata

  echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
  git clone commits_on_branch.bundle
  echo 'done'

  echo 'testdata/git_tags/ directory does not exist at the root; creating...'
  git clone git_tags.bundle
  echo 'done'
  
else
  echo 'testdata/commits_on_branch_test directory already exists; aborting.'
  exit -1
fi

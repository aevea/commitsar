#!/bin/bash

if [ ! -d './testdata/commits_on_branch_test' ]; then
  echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
  mkdir -p testdata/commits_on_branch_test
  echo 'done'
  
  cd testdata/commits_on_branch_test

  echo 'initialising test repository...'
  git init
  touch first
  git add .
  git commit -m "first commit on master"
  touch second
  git add .
  git commit -m "second commit on master"
  git checkout -b behind-master
  touch third
  git add .
  git commit -m "first commit on branch"
  git checkout master
  # Because this script is so fast it will actual provide false positives without this sleep
  sleep 1
  touch four
  git add .
  git commit -m "third commit on master"
  git checkout behind-master
  git merge master --no-edit
  echo 'done'
  
else
  echo 'testdata/commits_on_branch_test directory already exists; aborting.'
  exit -1
fi

#!/bin/bash

if [ ! -d './testdata/commits_on_branch_test' ]; then
  echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
  mkdir -p testdata/commits_on_branch_test
  echo 'done'
  
  cd testdata/commits_on_branch_test

  echo 'initialising test repository...'
  git config --global user.email "bot@commitsar.tch"
  git config --global user.name "Commitsar Bot"
  git init
  touch first
  git add .
  git commit -m "first commit on master"
  touch second
  git add .
  git commit -m "second commit on master"
  git checkout -b behind-master
  git reset HEAD~1
  git checkout -- .
  touch third
  git add .
  git commit -m "first commit on behind-master branch"
  git merge master --no-edit
  echo 'done'
  
else
  echo 'testdata/commits_on_branch_test directory already exists; aborting.'
  exit -1
fi

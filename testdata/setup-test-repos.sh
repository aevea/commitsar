#!/bin/bash

cd $(dirname $0)

for bundle in *.bundle; do
  repo="${bundle%.*}"
  rm -rf ./$repo;
  echo "testdata/$repo directory does not exist at the root; creating...";
  git clone $bundle;
  echo "done";
done

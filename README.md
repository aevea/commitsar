# Commitsar

[![Go Report Card](https://goreportcard.com/badge/github.com/commitsar-app/commitsar)](https://goreportcard.com/report/github.com/commitsar-app/commitsar)
[![Build Status](https://cloud.drone.io/api/badges/commitsar-app/commitsar/status.svg)](https://cloud.drone.io/commitsar-app/commitsar)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Tool to make sure your commits are compliant with [conventional commits](https://www.conventionalcommits.org). It is aimed mainly at CIs to prevent branches with commits that don't comply. Usage as a pre-commit hook is also under consideration.

## Table of contents

1. [Usage](#usage)

## Usage

Commitsar is shipped as a Dockerfile. This is the easiest way to add it to your CI.

**Important: Commitsar currently needs to be run in the same folder as the git repository you want checked, currently no override is provided for setting path to git repo see https://github.com/commitsar-app/commitsar/issues/93**

#### Github action

Checkout git in order to get commits and master branch
```
- name: Check out code into the Go module directory
        uses: actions/checkout@v1
```

Run the Commitsar action
```
- name: Commitsar Action
  uses: docker://commitsar/commitsar:latest
```


##### Example for CircleCI:

```
validate-commits:
	    docker:
	      - image: commitsar/commitsar
	    steps:
	      - checkout
	      - run: commitsar
```

##### From binary

Adjust for version and distribution. Please check [Releases](https://github.com/commitsar-app/commitsar/releases).

```
- curl -L -O https://github.com/commitsar-app/commitsar/releases/download/v0.0.2/commitsar_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf commitsar_v0.0.2_Linux_x86_64.tar.gz
- ./commitsar
```

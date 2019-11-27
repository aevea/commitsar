---
title: "Introduction"
weight: 10
---

# Introduction

Tool to make sure your commits are compliant with [conventional commits](https://www.conventionalcommits.org). It is aimed mainly at CIs to prevent branches with commits that don't comply. Usage as a pre-commit hook is also under consideration.

# Usage

Commitsar is shipped as a Dockerfile. This is the easiest way to add it to your CI.

**Important: Commitsar currently needs to be run in the same folder as the git repository you want checked, currently no override is provided for setting path to git repo see https://github.com/outillage/commitsar/issues/93**

#### Github action

> Github actions

Checkout git in order to get commits and master branch

```
- name: Check out code into the Go module directory
        uses: actions/checkout@v1
```

Run the Commitsar action

```
- name: Commitsar Action
  uses: commitsar-app/commitsar@v0.6.3 (substitute for current version)
```

## Example for CircleCI

> CircleCI

```yaml
validate-commits:
	    docker:
	      - image: commitsar/commitsar
	    steps:
	      - checkout
	      - run: commitsar
```

## From binary

> Binary

```shell
- curl -L -O https://github.com/outillage/commitsar/releases/download/v0.0.2/commitsar_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf commitsar_v0.0.2_Linux_x86_64.tar.gz
- ./commitsar
```

> **Important** adjust for current version

---
title: "Introduction"
weight: 10
---

# Introduction

Tool to make sure your commits are compliant with [conventional commits](https://www.conventionalcommits.org). It is aimed mainly at CIs to prevent branches with commits that don't comply. Usage as a pre-commit hook is also under consideration.

# Usage

Commitsar is shipped as a Dockerfile. This is the easiest way to add it to your CI.

**Important: Commitsar currently needs to be in the same folder as the git repository you want checked**

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
- curl -L -O https://github.com/commitsar-app/commitsar/releases/download/v0.0.2/commitsar_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf commitsar_v0.0.2_Linux_x86_64.tar.gz
- ./commitsar
```

> **Important** adjust for current version

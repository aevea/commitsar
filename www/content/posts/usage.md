---
title: "Usage"
date: 2019-08-14T15:47:52+02:00
draft: false
weight: 2
---

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

---
id: github
title: Github Actions
---

## Using the Github Action

A minimal example:
```yaml
name: Linters

on: [pull_request]

jobs:
  validate-commits:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code into the Go module directory
        uses: actions/checkout@v1
      - name: Commitsar check
        uses: aevea/commitsar@v0.11.0 (substitute for the current version)
```

This will run `commitsar` on every pull request and validate the commits for it.

## Using Github Actions + Docker

This is a faster method since you don't have to build the Docker image in your Github action. If you need maximum security provided by Github actions please the [Github Action Flow](#github-action).

```yaml
name: Linters

on: [pull_request]

jobs:
  validate-commits:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code into the Go module directory
        uses: actions/checkout@v1
      - name: Commitsar check
        uses: docker://aevea/commitsar
```
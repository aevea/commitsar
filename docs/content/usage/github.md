---
id: github
title: Github Actions
---

## Important prefaces

### JIRA pipeline

When using [JIRA](https://commitsar.aevea.ee/configuration/config-file) make sure to set `actions/checkout@v2` to pull from the pull_request HEAD. Github uses a single merge commit by default which is not referenced in git. This will cause commitsar JIRA check to fail as this commit will not be found by the API getting queried for PR by commit.

### actions/checkout@v2

When using `actions/checkout@v2` please set `fetch_depth` to 0. Currently commitsar needs full git objects to work correctly. This will be fixed in an upcoming release.

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

## Using with JIRA pipeline

This pipeline example uses the checkout at PR HEAD. <https://github.com/actions/checkout#Checkout-pull-request-HEAD-commit-instead-of-merge-commit>

```yaml
validate-commits:
  runs-on: ubuntu-latest
  steps:
    - name: Set up Go 1.13
      uses: actions/setup-go@v1
      with:
        go-version: 1.13
    - name: Check out code into the Go module directory
      uses: actions/checkout@v2
      with:
        fetch-depth: 0
        ref: ${{ github.event.pull_request.head.sha }}
    - name: Commitsar check
      uses: docker://aevea/commitsar
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

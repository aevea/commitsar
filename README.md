# Commitsar

[![Go Report Card](https://goreportcard.com/badge/github.com/aevea/commitsar)](https://goreportcard.com/report/github.com/aevea/commitsar)
![Test](https://github.com/aevea/commitsar/workflows/Test/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/aevea/commitsar?style=flat-square)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/aevea/commitsar/latest?style=flat-square)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![Run on Repl.it](https://repl.it/badge/github/aevea/commitsar)](https://repl.it/github/aevea/commitsar)

Tool to make sure your commits are compliant with [conventional commits](https://www.conventionalcommits.org). It is aimed mainly at CIs to prevent branches with commits that don't comply. Usage as a pre-commit hook is also under consideration.

## Table of contents

1. [Usage](#usage)
2. [Flags](#flags)

## Usage

Please check [Documentation](https://commitsar.tech).

**Important: Commitsar currently needs to be run in the same folder as the git repository you want checked, currently no override is provided for setting path to git repo see https://github.com/aevea/commitsar/issues/93**

#### Running using https://gobinaries.com/

```sh
curl -sf https://gobinaries.com/aevea/commitsar | sh
```

Or a specific version:
```sh
curl -sf https://gobinaries.com/aevea/commitsar[@VERSION] | sh
```


#### Github action

Checkout git in order to get commits and master branch

```
- name: Check out code into the Go module directory
        uses: actions/checkout@v1
```

Run the Commitsar action

```
- name: Commitsar Action
  uses: docker://aevea/commitsar
```

##### Example for CircleCI:

```
validate-commits:
	    docker:
	      - image: aevea/commitsar
	    steps:
	      - checkout
	      - run: commitsar
```

##### From binary

Adjust for version and distribution. Please check [Releases](https://github.com/aevea/commitsar/releases).

```
- curl -L -O https://github.com/aevea/commitsar/releases/download/v0.0.2/commitsar_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf commitsar_v0.0.2_Linux_x86_64.tar.gz
- ./commitsar
```

# Flags

Commitsar allows the following flags:

| Name    | Flag  | Required | Default | Description                                                                        |
| ------- | ----- | -------- | ------- | ---------------------------------------------------------------------------------- |
| Verbose | --v   | false    | false   | Debug output into console                                                          |
| Strict  | --s   | false    | true    | Strict check of category types                                                     |
| All     | --all | false    | false   | Whether to check all commits on given branch. **Takes precedence over LIMIT flag** |

On top of that a single argument is allowed:

`commitsar <from commit>...<to commit>`

e.g. `commitsar 7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4`

If only one commit hash is used then commitsar will assume it to be the TO commit.

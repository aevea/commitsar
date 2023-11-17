---
id: get_started
title: Get started
---

## Requirements

- [Go](https://golang.org/) (See pinned version in go.mod)
- [git](https://git-scm.com/) (while git is not required to run commitsar, it is needed to set up tests)

This project also uses [magefiles](https://magefile.org/). To get up and running please run: `go install github.com/magefile/mage`. These allow us to set up templates and cross-platform commands.

## Running tests

Run `mage test`. This will git clone all git bundles and run tests in the entire repo. To run single tests use the VSCode functionality or provide the full path to `go test`.

## Git bundles

In order to test commitsar against complicated real-life examples we use [git bundles](https://git-scm.com/docs/git-bundle). These can be `git cloned` like real repositories. They allow us to create real history and prevent edge cases that can happen when creating git history using shell (all commits have the same timestamp, for example).

### Creating a new bundle

- create a new folder testdata/
- cd into it
- git initgitz
- commit away
- once ready run `git bundle create name_of_repo.bundle --all`
- copy the bundle to testdata/

All files outside of .bundle files are ignored in the testdata folder so you don't have to worry about cleanup.

## Commit style

This project uses the [conventional commit style](https://www.conventionalcommits.org/en/v1.0.0/), it is also the default style commitsar uses.

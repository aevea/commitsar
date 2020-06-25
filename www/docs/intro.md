---
id: intro
title: Introduction
---

Tool to make sure your commits are compliant with [conventional commits](https://www.conventionalcommits.org). It is aimed mainly at CIs to prevent branches with commits that don't comply. Usage as a pre-commit hook is also under consideration.

## Usage

Commitsar is shipped as a Dockerfile. This is the easiest way to add it to your CI.

**Important: Commitsar currently needs to be run in the same folder as the git repository you want checked, currently no override is provided for setting path to git repo see https://github.com/aevea/commitsar/issues/93**

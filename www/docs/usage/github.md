---
id: github
title: Github Actions
---

Checkout git in order to get commits and master branch

```
- name: Check out code into the Go module directory
        uses: actions/checkout@v1
```

Run the Commitsar action

```
- name: Commitsar Action
  uses: aevea/commitsar@v0.6.3 (substitute for current version)
```
---
id: circleci
title: CircleCI
---

Minimal usage example:

```yaml
validate-commits:
  docker:
    - image: aevea/commitsar
  steps:
    - checkout
    - run: commitsar
```

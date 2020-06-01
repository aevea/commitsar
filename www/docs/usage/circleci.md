---
id: circleci
title: CircleCI
---

```yaml
validate-commits:
    docker:
	    - image: aevea/commitsar
    steps:
	    - checkout
	    - run: commitsar
```
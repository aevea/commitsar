---
id: gitlab
title: Gitlab CI
---

For Gitlab usage you can include the following job:

```yaml
validate-commits:
  stage: test
  image: aevea/commitsar
  script:
    - git fetch origin master
    - commitsar
```

**Important** In case of an error such as: `reference not found` please set a higher `GIT_DEPTH` variable setting. Commitsar currently relies on full commit objects which do not get pulled in on the shallow clone that `GIT_DEPTH` uses.

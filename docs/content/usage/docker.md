---
id: docker
title: Docker
---

For running in docker just use the following command:

```sh
docker run --rm --name="commitsar" -w /src -v "$(pwd)":/src aevea/commitsar 
```

```sh
docker run --rm --name="commitsar" -w /src -v "$(pwd)":/src aevea/commitsar ./path-to-repo
```

Make sure to load the working directory where `.git` folder is present. Commitsar will not work without the `.git` folder. This can be overridden by setting the path argument.
